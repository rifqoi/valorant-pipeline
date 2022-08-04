locals {
  data_lake_bucket = "valorant_data_lake_${var.project}"
  landing_zone_bucket = "valorant_landing_bucket_${var.project}"
  process_zone_bucket = "valorant_process_bucket_${var.project}"
}

terraform {
  required_version = ">=1.0"
  # backend "local" {} # gcs for google cloud, s3 for aws but you can use provide keyword to init the backedn
  backend "gcs" {
    bucket  = "valorant_data_lake_erudite-bonbon-352111"
    prefix  = "terraform/state"
  }
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
}

provider "google" {
  project = var.project
  region  = var.region
  // credentials = file(var.credentials)  # Use this if you do not want to set env-var GOOGLE_APPLICATION_CREDENTIALS
}

resource "google_storage_bucket" "data-lake-bucket" {
  name     = local.data_lake_bucket
  location = var.region

  # Optional, but recommended settings:
  storage_class               = var.storage_class
  uniform_bucket_level_access = true

  versioning {
    enabled = true
  }

}

# Landing zone bucket
resource "google_storage_bucket" "landing-zone-bucket" {
  name     = local.landing_zone_bucket
  location = var.region

  # Optional, but recommended settings:
  storage_class               = var.storage_class
  uniform_bucket_level_access = true

  versioning {
    enabled = true
  }

}

# Process zone bucket
resource "google_storage_bucket" "process_zone_bucket" {
  name     = local.process_zone_bucket
  location = var.region

  # Optional, but recommended settings:
  storage_class               = var.storage_class
  uniform_bucket_level_access = true

  versioning {
    enabled = true
  }

}
# DWH
# Ref: https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/bigquery_dataset
resource "google_bigquery_dataset" "dataset" {
  dataset_id = var.BQ_DATASET
  project    = var.project
  location   = var.region
}

# Firewall
resource "google_compute_firewall" "rules" {
  project     = var.project
  name        = "airflow-firewall"
  network     = "default"
  description = "Airflow firewall targeting port 8080"

  allow {
    protocol  = "tcp"
    ports     = ["22", "8080", "5000"]
  }

  allow {
    protocol = "icmp"
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags = ["airflow"]
}

resource "google_compute_instance" "valorant-vm1" {
  name         = "valorant-vm1"
  machine_type = "e2-standard-4"
  zone         = var.zone
  project = var.project
  allow_stopping_for_update = true
  desired_status = "TERMINATED"

  tags = ["airflow-firewall", "http-server", "https-server"]

  boot_disk {
    initialize_params {
      image = "ubuntu-2004-lts"
    }
  }

  network_interface {
    network = "default"

    access_config {
      // Ephemeral public IP
    }
  }

  service_account {
    # Google recommends custom service accounts that have cloud-platform scope
    # and permissions granted via IAM Roles.
    email  = "de-projects@erudite-bonbon-352111.iam.gserviceaccount.com"
    scopes = ["cloud-platform"]
  }
}
