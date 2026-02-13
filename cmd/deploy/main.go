package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	run "cloud.google.com/go/run/apiv2"
	runpb "cloud.google.com/go/run/apiv2/runpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	project := flag.String("project", "", "GCP project ID (required)")
	region := flag.String("region", "europe-west2", "GCP region")
	tag := flag.String("tag", "", "image tag (default: git short SHA)")
	dryRun := flag.Bool("dry-run", false, "print commands without executing")
	skipBuild := flag.Bool("skip-build", false, "skip Hugo and container build steps")
	flag.Parse()

	if *project == "" {
		log.Fatal("flag -project is required")
	}

	if *tag == "" {
		*tag = gitTag()
	}

	registry := fmt.Sprintf("%s-docker.pkg.dev/%s/blog", *region, *project)
	image := fmt.Sprintf("%s/blog:%s", registry, *tag)
	serviceName := fmt.Sprintf("projects/%s/locations/%s/services/blog", *project, *region)

	log.Printf("project=%s region=%s tag=%s image=%s", *project, *region, *tag, image)

	if !*skipBuild {
		step("Checking prerequisites", func() { checkPrereqs() })
		step("Building Hugo site", func() {
			run_cmd(*dryRun, "hugo", "--minify")
		})
		step("Building container image", func() {
			run_cmd(*dryRun, "podman", "build", "-t", image, "-f", "Containerfile", ".")
		})
	}

	authFile := filepath.Join(os.TempDir(), "podman-gcp-auth.json")

	step("Logging in to Artifact Registry", func() {
		if *dryRun {
			log.Printf("[dry-run] podman login %s", registry)
			return
		}
		loginRegistry(registry, authFile)
	})

	step("Pushing image to Artifact Registry", func() {
		run_cmd(*dryRun, "podman", "push", "--authfile", authFile, image)
	})

	step("Updating Cloud Run service", func() {
		if *dryRun {
			log.Printf("[dry-run] would update Cloud Run service to image %s", image)
			return
		}
		if err := updateCloudRunService(context.Background(), serviceName, image); err != nil {
			log.Fatalf("update Cloud Run: %v", err)
		}
	})

	log.Println("deploy complete")
}

func step(name string, fn func()) {
	log.Printf("==> %s", name)
	start := time.Now()
	fn()
	log.Printf("    done (%s)", time.Since(start).Round(time.Millisecond))
}

func checkPrereqs() {
	for _, bin := range []string{"hugo", "podman", "gcloud"} {
		if _, err := exec.LookPath(bin); err != nil {
			log.Fatalf("prerequisite %q not found in PATH", bin)
		}
	}

	// Verify gcloud auth.
	out, err := exec.Command("gcloud", "auth", "print-access-token").CombinedOutput()
	if err != nil {
		log.Fatalf("gcloud not authenticated: %s\n%s", err, out)
	}
}

func gitTag() string {
	out, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		return fmt.Sprintf("t%d", time.Now().Unix())
	}
	return strings.TrimSpace(string(out))
}

func run_cmd(dryRun bool, name string, args ...string) {
	if dryRun {
		log.Printf("[dry-run] %s %s", name, strings.Join(args, " "))
		return
	}
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("%s failed: %v", name, err)
	}
}

func loginRegistry(registry, authFile string) {
	// Extract the hostname (e.g. "europe-west2-docker.pkg.dev") from the registry path.
	host := strings.SplitN(registry, "/", 2)[0]

	token, err := exec.Command("gcloud", "auth", "print-access-token").Output()
	if err != nil {
		log.Fatalf("gcloud auth print-access-token: %v", err)
	}

	cmd := exec.Command("podman", "login",
		"--authfile", authFile,
		"-u", "oauth2accesstoken",
		"--password-stdin",
		host,
	)
	cmd.Stdin = strings.NewReader(strings.TrimSpace(string(token)))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("podman login failed: %v", err)
	}
}

func updateCloudRunService(ctx context.Context, serviceName, image string) error {
	client, err := run.NewServicesClient(ctx)
	if err != nil {
		return fmt.Errorf("create client: %w", err)
	}
	defer client.Close()

	// Get current service.
	svc, err := client.GetService(ctx, &runpb.GetServiceRequest{
		Name: serviceName,
	})
	if err != nil {
		return fmt.Errorf("get service: %w", err)
	}

	// Update the container image.
	svc.Template.Containers[0].Image = image

	op, err := client.UpdateService(ctx, &runpb.UpdateServiceRequest{
		Service: svc,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"template.containers"},
		},
	})
	if err != nil {
		return fmt.Errorf("update service: %w", err)
	}

	// Wait for the new revision to be serving.
	log.Println("    waiting for new revision…")
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("wait for revision: %w", err)
	}

	return nil
}
