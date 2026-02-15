output "kubeconfig" {
  description = "Kubeconfig for the OVH Kubernetes cluster"
  value       = ovh_cloud_project_kube.blog.kubeconfig
  sensitive   = true
}

output "cluster_id" {
  description = "ID of the OVH Kubernetes cluster"
  value       = ovh_cloud_project_kube.blog.id
}
