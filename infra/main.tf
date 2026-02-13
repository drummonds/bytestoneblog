terraform {
  required_version = ">= 1.0"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 6.0"
    }
  }

  backend "gcs" {
    bucket = "bytestone-terraform-state"
    prefix = "blog"
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

# Enable required GCP APIs.
locals {
  apis = [
    "run.googleapis.com",
    "artifactregistry.googleapis.com",
    "dns.googleapis.com",
    "iam.googleapis.com",
    "compute.googleapis.com",
    "cloudresourcemanager.googleapis.com",
  ]
}

resource "google_project_service" "apis" {
  for_each = toset(local.apis)

  project            = var.project_id
  service            = each.value
  disable_on_destroy = false
}
