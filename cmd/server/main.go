package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func main() {
	dir := flag.String("dir", "/site", "directory to serve")
	flag.Parse()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})
	mux.Handle("/", fileHandler(*dir))

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      loggingMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Graceful shutdown on SIGTERM.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		log.Printf("serving %s on :%s", *dir, port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down…")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("shutdown: %v", err)
	}
	log.Println("stopped")
}

// fileHandler returns a handler that serves static files with custom behaviour:
//   - directory requests serve index.html
//   - /foo redirects to /foo/ if foo is a directory
//   - unmatched paths return 404.html with 404 status
//   - cache headers set by file type
func fileHandler(root string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path

		// Clean the path to prevent directory traversal.
		clean := filepath.Clean(urlPath)
		filePath := filepath.Join(root, clean)

		info, err := os.Stat(filePath)
		if err != nil {
			serve404(w, r, root)
			return
		}

		// If it's a directory, redirect to trailing slash if needed,
		// then serve index.html.
		if info.IsDir() {
			if !strings.HasSuffix(urlPath, "/") {
				http.Redirect(w, r, urlPath+"/", http.StatusMovedPermanently)
				return
			}
			indexPath := filepath.Join(filePath, "index.html")
			if _, err := os.Stat(indexPath); err != nil {
				serve404(w, r, root)
				return
			}
			setCacheHeaders(w, "index.html")
			http.ServeFile(w, r, indexPath)
			return
		}

		setCacheHeaders(w, filePath)
		http.ServeFile(w, r, filePath)
	}
}

func serve404(w http.ResponseWriter, r *http.Request, root string) {
	page := filepath.Join(root, "404.html")
	content, err := os.ReadFile(page)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write(content)
}

func setCacheHeaders(w http.ResponseWriter, name string) {
	ext := strings.ToLower(filepath.Ext(name))
	var cc string
	switch ext {
	case ".css", ".js", ".woff", ".woff2", ".ttf", ".eot":
		cc = "public, max-age=31536000, immutable"
	case ".png", ".jpg", ".jpeg", ".gif", ".webp", ".svg", ".ico":
		cc = "public, max-age=86400"
	case ".html":
		cc = "public, max-age=300"
	case ".xml":
		cc = "public, max-age=3600"
	default:
		cc = "public, max-age=3600"
	}
	w.Header().Set("Cache-Control", cc)
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		log.Printf("method=%s path=%s status=%d duration=%s",
			r.Method, r.URL.Path, rec.status, time.Since(start).Round(time.Microsecond))
	})
}
