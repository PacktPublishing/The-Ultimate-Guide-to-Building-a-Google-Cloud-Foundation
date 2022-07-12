resource "google_compute_instance" "app_vm" {
  name                      = "demo-vm"
  machine_type              = var.machine_type
  zone                      = var.zone
  allow_stopping_for_update = true

  boot_disk {
    initialize_params {
      image = "ubuntu-2110-impish-v20220204"
      # You can find available images in GCP on the 
      # Compute Engine | Images page.
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }
}
