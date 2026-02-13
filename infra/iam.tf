resource "google_service_account" "deployer" {
  account_id   = "blog-deployer"
  display_name = "Blog Deployer"
  description  = "Service account for deploying the blog to Cloud Run"

  depends_on = [google_project_service.apis]
}

# Push images to Artifact Registry.
resource "google_project_iam_member" "deployer_ar_writer" {
  project = var.project_id
  role    = "roles/artifactregistry.writer"
  member  = "serviceAccount:${google_service_account.deployer.email}"
}

# Deploy new revisions to Cloud Run.
resource "google_project_iam_member" "deployer_run_admin" {
  project = var.project_id
  role    = "roles/run.admin"
  member  = "serviceAccount:${google_service_account.deployer.email}"
}

# Allow the deployer to act as the Cloud Run service agent.
resource "google_project_iam_member" "deployer_sa_user" {
  project = var.project_id
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${google_service_account.deployer.email}"
}
