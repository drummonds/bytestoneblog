output "cloud_run_url" {
  description = "Cloud Run service URL"
  value       = google_cloud_run_v2_service.blog.uri
}

output "dns_nameservers" {
  description = "Nameservers for the DNS zone — set these at your registrar"
  value       = google_dns_managed_zone.blog.name_servers
}

output "deployer_email" {
  description = "Deployer service account email"
  value       = google_service_account.deployer.email
}

output "registry_path" {
  description = "Artifact Registry image path prefix"
  value       = "${var.region}-docker.pkg.dev/${var.project_id}/${google_artifact_registry_repository.blog.repository_id}"
}
