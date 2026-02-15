package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "", "path to kubeconfig (default: infra/ovh/kubeconfig.yml)")
	registryPassword := flag.String("registry-password", "", "registry htpasswd password (default: $REGISTRY_PASSWORD or auto-generated)")
	manifestDir := flag.String("manifests", "", "path to k8s manifests dir (default: infra/ovh/k8s)")
	dryRun := flag.Bool("dry-run", false, "print commands without executing")
	flag.Parse()

	// Resolve defaults relative to repo root.
	repoRoot := repoRoot()

	if *kubeconfig == "" {
		*kubeconfig = filepath.Join(repoRoot, "infra", "ovh", "kubeconfig.yml")
	}
	if *manifestDir == "" {
		*manifestDir = filepath.Join(repoRoot, "infra", "ovh", "k8s")
	}
	if *registryPassword == "" {
		*registryPassword = os.Getenv("REGISTRY_PASSWORD")
	}
	if *registryPassword == "" {
		*registryPassword = generatePassword()
		log.Printf("Generated registry password: %s", *registryPassword)
		log.Println("Save this! Set REGISTRY_PASSWORD or -registry-password for future runs.")
	}

	os.Setenv("KUBECONFIG", *kubeconfig)
	log.Printf("kubeconfig=%s manifests=%s", *kubeconfig, *manifestDir)

	step("Checking prerequisites", func() {
		checkPrereqs(*dryRun)
	})

	step("Installing nginx-ingress controller", func() {
		helmUpgrade(*dryRun,
			"ingress-nginx", "ingress-nginx",
			"https://kubernetes.github.io/ingress-nginx",
			"ingress-nginx",
			"controller.resources.requests.memory=64Mi",
			"controller.resources.limits.memory=128Mi",
		)
	})

	step("Installing cert-manager", func() {
		helmUpgrade(*dryRun,
			"cert-manager", "cert-manager",
			"https://charts.jetstack.io",
			"cert-manager",
			"crds.enabled=true",
			"resources.requests.memory=32Mi",
			"resources.limits.memory=64Mi",
		)
	})

	step("Applying namespace", func() {
		kubectlApply(*dryRun, filepath.Join(*manifestDir, "namespace.yaml"))
	})

	step("Creating registry htpasswd secret", func() {
		createRegistrySecret(*dryRun, *registryPassword)
	})

	step("Applying ClusterIssuer", func() {
		kubectlApply(*dryRun, filepath.Join(*manifestDir, "cert-issuer.yaml"))
	})

	step("Applying registry manifests", func() {
		kubectlApply(*dryRun, filepath.Join(*manifestDir, "registry.yaml"))
	})

	step("Applying blog service", func() {
		kubectlApply(*dryRun, filepath.Join(*manifestDir, "service.yaml"))
	})

	step("Applying blog ingress", func() {
		kubectlApply(*dryRun, filepath.Join(*manifestDir, "ingress.yaml"))
	})

	step("Applying blog deployment", func() {
		kubectlApply(*dryRun, filepath.Join(*manifestDir, "deployment.yaml"))
	})

	step("Fetching LoadBalancer IP", func() {
		if *dryRun {
			log.Println("[dry-run] would query ingress-nginx-controller service")
			return
		}
		printLoadBalancerIP()
	})

	log.Println("bootstrap complete")
}

func step(name string, fn func()) {
	log.Printf("==> %s", name)
	start := time.Now()
	fn()
	log.Printf("    done (%s)", time.Since(start).Round(time.Millisecond))
}

func checkPrereqs(dryRun bool) {
	for _, bin := range []string{"helm", "kubectl"} {
		if _, err := exec.LookPath(bin); err != nil {
			log.Fatalf("prerequisite %q not found in PATH", bin)
		}
	}
	if dryRun {
		return
	}
	// Verify kubectl can reach the cluster.
	out, err := exec.Command("kubectl", "cluster-info").CombinedOutput()
	if err != nil {
		log.Fatalf("kubectl cannot reach cluster: %s\n%s", err, out)
	}
}

func helmUpgrade(dryRun bool, release, chart, repo, namespace string, sets ...string) {
	args := []string{
		"upgrade", "--install", release, chart,
		"--repo", repo,
		"--namespace", namespace, "--create-namespace",
		"--wait",
	}
	for _, s := range sets {
		args = append(args, "--set", s)
	}
	runCmd(dryRun, "helm", args...)
}

func kubectlApply(dryRun bool, path string) {
	runCmd(dryRun, "kubectl", "apply", "-f", path)
}

func createRegistrySecret(dryRun bool, password string) {
	if dryRun {
		log.Println("[dry-run] kubectl create secret generic registry-htpasswd ...")
		return
	}

	// Generate htpasswd line using htpasswd or fallback.
	htpasswd, err := generateHtpasswd(password)
	if err != nil {
		log.Fatalf("generate htpasswd: %v", err)
	}

	// Idempotent: create with --dry-run=client then pipe to apply.
	create := exec.Command("kubectl", "create", "secret", "generic", "registry-htpasswd",
		"--from-literal=htpasswd="+htpasswd,
		"--namespace", "blog",
		"--dry-run=client", "-o", "yaml",
	)
	apply := exec.Command("kubectl", "apply", "-f", "-")

	pipe, err := create.StdoutPipe()
	if err != nil {
		log.Fatalf("pipe: %v", err)
	}
	apply.Stdin = pipe
	apply.Stdout = os.Stdout
	apply.Stderr = os.Stderr

	if err := create.Start(); err != nil {
		log.Fatalf("kubectl create: %v", err)
	}
	if err := apply.Start(); err != nil {
		log.Fatalf("kubectl apply: %v", err)
	}
	if err := create.Wait(); err != nil {
		log.Fatalf("kubectl create: %v", err)
	}
	if err := apply.Wait(); err != nil {
		log.Fatalf("kubectl apply: %v", err)
	}
}

func generateHtpasswd(password string) (string, error) {
	// Try htpasswd binary first.
	if _, err := exec.LookPath("htpasswd"); err == nil {
		out, err := exec.Command("htpasswd", "-Bbn", "admin", password).Output()
		if err != nil {
			return "", fmt.Errorf("htpasswd: %w", err)
		}
		return strings.TrimSpace(string(out)), nil
	}
	// Fallback: use openssl for bcrypt-style password.
	out, err := exec.Command("openssl", "passwd", "-apr1", password).Output()
	if err != nil {
		return "", fmt.Errorf("openssl passwd: %w", err)
	}
	return "admin:" + strings.TrimSpace(string(out)), nil
}

func printLoadBalancerIP() {
	for i := 0; i < 30; i++ {
		out, err := exec.Command("kubectl", "get", "svc", "ingress-nginx-controller",
			"-n", "ingress-nginx",
			"-o", "jsonpath={.status.loadBalancer.ingress[0].ip}",
		).Output()
		if err == nil && len(strings.TrimSpace(string(out))) > 0 {
			ip := strings.TrimSpace(string(out))
			log.Println("")
			log.Println("============================================")
			log.Printf("LoadBalancer IP: %s", ip)
			log.Println("")
			log.Println("Configure DNS A records:")
			log.Printf("  www.bytestone.uk      -> %s", ip)
			log.Printf("  registry.bytestone.uk -> %s", ip)
			log.Println("============================================")
			return
		}
		time.Sleep(5 * time.Second)
	}
	log.Println("LoadBalancer IP not yet assigned.")
	log.Println("Check with: kubectl get svc ingress-nginx-controller -n ingress-nginx")
}

func runCmd(dryRun bool, name string, args ...string) {
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

func repoRoot() string {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		log.Fatal("not in a git repository")
	}
	return strings.TrimSpace(string(out))
}

func generatePassword() string {
	b := make([]byte, 18)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("generate password: %v", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
