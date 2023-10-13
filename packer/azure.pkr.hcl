source "azure-arm" "vm" {
  use_azure_cli_auth = true
  subscription_id    = var.subscription_id
  tenant_id          = var.tenant_id

  location                          = var.locationName
  managed_image_name                = "img-packer-${var.environmentType}-${var.iteration}"
  managed_image_resource_group_name = "rg-${var.project_prefix}-${var.environmentType}-${var.iteration}"

  shared_image_gallery_destination {
    subscription        = var.subscription_id
    resource_group      = "rg-${var.project_prefix}-${var.environmentType}-${var.iteration}"
    gallery_name        = "az_compute_gallery"
    image_name          = "avd_base_image"
    image_version       = "1.0.0"
    replication_regions = [var.locationName]
  }

  communicator   = "winrm"
  winrm_use_ssl  = true
  winrm_insecure = true
  winrm_timeout  = "5m"
  winrm_username = "packer"

  os_type         = "Windows"
  image_publisher = "microsoftwindowsdesktop" //Product ID: windows-11 //Plan ID: win11-21h2-pro
  image_offer     = "Windows-11"
  image_sku       = "win11-22h2-avd"

  vm_size = "Standard_D2s_v3"
}

