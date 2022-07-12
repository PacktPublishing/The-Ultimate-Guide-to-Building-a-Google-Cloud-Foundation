variable "zone" {
  description = "The zone where the compute resources will be created"
  type        = string
}

variable "machine_type" {
  description = "GCE machine type."
  type        = string
  default     = "n1-standard-4"
}
