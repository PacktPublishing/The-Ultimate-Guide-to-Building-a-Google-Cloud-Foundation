output "vm_public_ip" {
    value = "${module.compute.vm_public_ip}"
    description = "The public ip for the new VM"
}

output "bucket_name"{
    value = "${module.storage.bucket_name}"
}