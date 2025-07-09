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

resource "google_cloud_run_service" "go_service" {
  name     = "loan-buddy-go"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/loan-buddy-go:latest"
        ports {
          container_port = 8080
        }
        env {
          name  = "PAYMENT_SERVICE_URL"
          value = var.payment_service_url
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service" "python_service" {
  name     = "loan-buddy-python"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/loan-buddy-python:latest"
        ports {
          container_port = 8000
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
