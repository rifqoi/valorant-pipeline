locals {
  data_lake_bucket = "valorant_data_lake"
  landing_zone_bucket = "valorant_landing_bucket"
  process_zone_bucket = "valorant_process_bucket"
}

variable "project" {
  description = "GCP Project ID"
  default     = "erudite-bonbon-352111"
}

variable "region" {
  description = "Region for GCP resources. Choose as per your location: https://cloud.google.com/about/locations"
  default     = "asia-southeast2"
  type        = string
}

variable "storage_class" {
  description = "Storage class type for your bucket. Check official docs for more info."
  default     = "STANDARD"
}

variable "BQ_DATASET" {
  description = "BigQuery Dataset that raw data (from GCS) will be written to"
  type        = string
  default     = "valorant_datasets"
}

