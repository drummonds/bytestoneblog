resource "ovh_cloud_project_kube" "blog" {
  service_name = var.ovh_project_id
  name         = var.cluster_name
  region       = var.region
  version      = var.k8s_version
}

resource "ovh_cloud_project_kube_nodepool" "workers" {
  service_name  = var.ovh_project_id
  kube_id       = ovh_cloud_project_kube.blog.id
  name          = "workers"
  flavor_name   = var.node_flavor
  desired_nodes = 1
  min_nodes     = 1
  max_nodes     = 1
}
