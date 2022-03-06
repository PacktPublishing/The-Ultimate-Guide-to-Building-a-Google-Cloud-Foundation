output "vm_public_ip" {
    value = google_compute_instance.app_vm.network_interface.0.access_config.0.nat_ip
    description = "The public IP for the new VM"
}