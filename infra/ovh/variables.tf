variable "ovh_project_id" {
  description = "OVH Public Cloud project ID"
  type        = string
}

variable "region" {
  description = "OVH region for the Kubernetes cluster"
  type        = string
  default     = "GRA9"
}

variable "cluster_name" {
  description = "Name for the Managed Kubernetes cluster"
  type        = string
  default     = "blog"
}

variable "k8s_version" {
  description = "Kubernetes version"
  type        = string
  default     = "1.31"
}

variable "node_flavor" {
  description = "OVH instance flavor for worker nodes"
  type        = string
  default     = "d2-2"
}

variable "domain" {
  description = "Root domain for the blog"
  type        = string
  default     = "bytestone.uk"
}
