data "azurerm_client_config" "current" {}

resource "azurerm_resource_group" "rg_root" {
  name     = "rg-${var.project_prefix}-${var.environmentType}-${var.iteration}"
  location = var.locationName
}

resource "azurerm_shared_image_gallery" "sig_compute_gallery" {
  name                = "Azure Compute Gallery"
  resource_group_name = azurerm_resource_group.rg_root.name
  location            = var.locationName

}

resource "azurerm_shared_image" "si_base_image" {
  //az vm image list --publisher microsoftwindowsdesktop --offer Windows-11 --sku win11-22h2-avd --output table --all
  name                = "base_image"
  gallery_name        = azurerm_shared_image_gallery.sig_compute_gallery.name
  resource_group_name = azurerm_resource_group.rg_root.name
  location            = var.locationName
  os_type             = "Windows"

  identifier {
    publisher = "microsoftwindowsdesktop" //Product ID: windows-11 //Plan ID: win11-21h2-pro
    offer     = "Windows-11"
    sku       = "win11-22h2-avd"
  }
}
