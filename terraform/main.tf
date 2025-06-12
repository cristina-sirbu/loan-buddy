terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }

  required_version = ">= 1.5.0"
}

provider "google" {
  project = var.project_id
  region  = var.region
  # This is just a placeholder.
  # credentials = file("fake-gcp-key.json")
}

resource "google_container_cluster" "loan_buddy_gke" {
  name     = "loan-buddy-cluster"
  location = var.region

  initial_node_count = 1

  node_config {
    machine_type = "e2-medium"
  }
}
