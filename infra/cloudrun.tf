resource "google_cloud_run_v2_service" "blog" {
  name     = "blog"
  location = var.region

  template {
    scaling {
      min_instance_count = 0
      max_instance_count = 2
    }

    containers {
      # Placeholder image — replaced by the deploy tool on first deploy.
      image = "us-docker.pkg.dev/cloudrun/container/hello"

      ports {
        container_port = 8080
      }

      resources {
        limits = {
          memory = "256Mi"
          cpu    = "1"
        }
        cpu_idle = true
      }

      startup_probe {
        http_get {
          path = "/healthz"
        }
        initial_delay_seconds = 0
        period_seconds        = 3
        failure_threshold     = 3
      }

      liveness_probe {
        http_get {
          path = "/healthz"
        }
        period_seconds = 30
      }
    }
  }

  # The deploy tool updates the image; don't let Terraform revert it.
  lifecycle {
    ignore_changes = [
      template[0].containers[0].image,
    ]
  }

  depends_on = [google_project_service.apis]
}

# Allow unauthenticated access (public blog).
resource "google_cloud_run_v2_service_iam_member" "public" {
  project  = google_cloud_run_v2_service.blog.project
  location = google_cloud_run_v2_service.blog.location
  name     = google_cloud_run_v2_service.blog.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}
