variable "project_id" {
  description = "GCP project ID"
  type        = string
}

variable "region" {
  description = "GCP region for all resources"
  type        = string
  default     = "europe-west2"
}

variable "domain" {
  description = "Root domain for the blog"
  type        = string
  default     = "bytestone.uk"
}
