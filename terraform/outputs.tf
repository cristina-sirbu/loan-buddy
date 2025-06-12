output "cluster_name" {
  value = google_container_cluster.loan_buddy_gke.name
}

output "endpoint" {
  value = google_container_cluster.loan_buddy_gke.endpoint
}
