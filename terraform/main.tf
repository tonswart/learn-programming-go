data "azuread_application" "example" {
  display_name = "DHL eCommerce BNL JFrog Artifactory"
}

data "azuread_group" "example" {
  for_each = toset(local.aad_groups)

  display_name     = each.value
  security_enabled = true
}

locals {
  app_roles = { for item in data.azuread_application.example.app_roles : item.display_name => item }
  aad_groups = [
    "CDLAPP-DHL-Parcel-NL-InfraTeam-Users",
    "CDLAPP-DHL-Parcel-NL-Dadp-Users"
  ]
}

resource "azuread_app_role_assignment" "example" {
  for_each = toset(local.aad_groups)

  app_role_id         = local.app_roles["msiam_access"].id
  resource_object_id  = data.azuread_application.example.object_id

  principal_object_id = data.azuread_group.example[each.key].object_id
}
