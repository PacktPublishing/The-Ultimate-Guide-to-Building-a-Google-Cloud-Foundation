
resource "google_storage_bucket" "app_bucket" {
  name          = "bkt-${var.project_id}-demo-app"
  location      = var.region
  storage_class = "STANDARD"
  force_destroy = true
  uniform_bucket_level_access = true
}