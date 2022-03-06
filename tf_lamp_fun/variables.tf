variable "project_id" {
  description = "The GCP Project where these resources will be deployed."
  type        = string
}

variable "zone" {
  description = "The default GCP zone"
  type        = string
}

variable "region" {
  description = "The default GCP region"
  type        = string
}
