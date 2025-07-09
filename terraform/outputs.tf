output "go_service_url" {
  value = google_cloud_run_service.go_service.status[0].url
}

output "python_service_url" {
  value = google_cloud_run_service.python_service.status[0].url
}
