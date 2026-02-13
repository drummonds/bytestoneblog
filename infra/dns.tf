resource "google_dns_managed_zone" "blog" {
  name        = "bytestone-uk"
  dns_name    = "${var.domain}."
  description = "DNS zone for ${var.domain}"

  depends_on = [google_project_service.apis]
}

# CNAME for www → Cloud Run.
resource "google_dns_record_set" "www" {
  managed_zone = google_dns_managed_zone.blog.name
  name         = "www.${var.domain}."
  type         = "CNAME"
  ttl          = 300
  rrdatas      = ["ghs.googlehosted.com."]
}

# Domain mapping — uncomment after running:
#   gcloud domains verify bytestone.uk
#
# resource "google_cloud_run_domain_mapping" "www" {
#   location = var.region
#   name     = "www.${var.domain}"
#
#   metadata {
#     namespace = var.project_id
#   }
#
#   spec {
#     route_name = google_cloud_run_v2_service.blog.name
#   }
# }
