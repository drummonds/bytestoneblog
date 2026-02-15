terraform {
  required_version = ">= 1.6"

  required_providers {
    ovh = {
      source  = "ovh/ovh"
      version = "~> 2.5"
    }
  }

  backend "local" {}
}

provider "ovh" {
  endpoint = "ovh-eu"
  # Credentials from environment:
  #   OVH_APPLICATION_KEY
  #   OVH_APPLICATION_SECRET
  #   OVH_CONSUMER_KEY
}
