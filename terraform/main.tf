terraform {
  required_version = ">=1.0"
  backend "local" {} # gcs for google cloud, s3 for aws but you can use provide keyword to init the backedn
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
}

locals {
  data_lake_bucket = "valorant_data_lake_${var.project}"
  landing_zone_bucket = "valorant_landing_bucket_${var.project}"
  process_zone_bucket = "valorant_process_bucket_${var.project}"
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
