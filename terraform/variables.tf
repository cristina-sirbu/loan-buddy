variable "project_id" {
  description = "GCP project ID"
  type        = string
  default     = "loan-buddy-demo"
}

variable "region" {
  type    = string
  default = "us-central1"
}

variable "payment_service_url" {
  type = string
}