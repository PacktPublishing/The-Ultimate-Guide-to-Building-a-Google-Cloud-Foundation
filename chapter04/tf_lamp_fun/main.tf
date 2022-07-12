provider "google" {
  project = var.project_id
}

locals {
  cpu_count    = "1"
  chip_type    = "standard"
  chip_family  = "n1"
  machine_type = format("%s-%s-%s", local.chip_family, local.chip_type, local.cpu_count)
}

module "compute" {
  source       = "./modules/compute"
  zone         = var.zone
  machine_type = local.machine_type
  depends_on = [
    module.storage
  ]
}

module "storage" {
  source     = "./modules/storage"
  region     = var.region
  project_id = var.project_id
}
