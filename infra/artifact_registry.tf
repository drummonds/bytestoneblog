resource "google_artifact_registry_repository" "blog" {
  location      = var.region
  repository_id = "blog"
  format        = "DOCKER"
  description   = "Container images for bytestone blog"

  cleanup_policies {
    id     = "keep-recent"
    action = "KEEP"

    most_recent_versions {
      keep_count = 5
    }
  }

  depends_on = [google_project_service.apis]
}
