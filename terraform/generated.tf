## __generated__ by Terraform
## Please review these resources and move them into your main configuration files.
#
## __generated__ by Terraform from "79bedc6c-961b-4d6e-a2c1-a3c4061fe73c/appRoleAssignment/k4pKxsI7TUivkwYsxqRdfglkcEYW4mhCtR4Ib2sIaEk"
#resource "azuread_app_role_assignment" "example" {
#  app_role_id         = "6594cbb3-da49-4e74-a3bd-da5b30a5461b"
#  principal_object_id = "c64a8a93-3bc2-484d-af93-062cc6a45d7e"
#  resource_object_id  = "79bedc6c-961b-4d6e-a2c1-a3c4061fe73c"
#  timeouts {
#    create = null
#    delete = null
#    read   = null
#  }
#}
#
## __generated__ by Terraform
#resource "azuread_application" "example" {
#  device_only_auth_enabled       = false
#  display_name                   = "DHL eCommerce BNL JFrog Artifactory"
#  fallback_public_client_enabled = false
#  group_membership_claims        = ["All"]
#  identifier_uris                = ["https://dhlparcel.pe.jfrog.io/ui/login"]
#  logo_image                     = null
#  marketing_url                  = null
#  oauth2_post_response_required  = false
#  owners                         = ["084ba2b9-741c-4192-bae8-5e9e8c9b88df", "123e2afc-45a4-4a1b-9fa9-f0085673b3c9", "4057d968-8eff-4d16-864c-4a4daa14174a", "4c5db822-e1d1-45d8-b15a-4b638d5b1a48", "5276c33c-27f4-4e97-b214-89481348c6a9", "92f3bc42-cd56-4d08-800e-b14b31fedad4", "aff4ad65-8fea-4651-8a66-8beb99d13474", "bf3ec58e-9a2f-458d-8589-98f9f3ca6106", "c66cee6f-4c5d-4ba9-ab93-5e8ad95207b8", "cc6c7a76-fd31-403c-b142-8f4f3fb0c16a", "cd39f849-17a6-49fc-ac37-0840ae4a8283", "d0d88120-e2cd-43e6-b784-a71cb08cdda4", "de435233-a919-4d8d-a1e4-7b21f9ba0363", "ea74cde4-85fb-46c3-bc2e-e8ef6399b609", "f9a2f7f4-319e-471b-992b-d5b689b42118", "f9f8f958-9999-4985-8ecd-e9683e0483ca"]
#  prevent_duplicate_names        = false
#  privacy_statement_url          = null
#  sign_in_audience               = "AzureADMyOrg"
#  support_url                    = null
#  tags                           = []
#  template_id                    = "5da54e9e-f26d-4c5e-a38c-b410271496a8"
#  terms_of_service_url           = null
#  api {
#    known_client_applications      = []
#    mapped_claims_enabled          = false
#    requested_access_token_version = 1
#    oauth2_permission_scope {
#      admin_consent_description  = "Allow the application to access JFrog Artifactory on behalf of the signed-in user."
#      admin_consent_display_name = "Access JFrog Artifactory"
#      enabled                    = true
#      id                         = "3bb9b43b-6855-4623-85e1-7ca02d32051d"
#      type                       = "User"
#      user_consent_description   = "Allow the application to access JFrog Artifactory on your behalf."
#      user_consent_display_name  = "Access JFrog Artifactory"
#      value                      = "user_impersonation"
#    }
#  }
#  app_role {
#    allowed_member_types = ["User"]
#    description          = "msiam_access"
#    display_name         = "msiam_access"
#    enabled              = true
#    id                   = "6594cbb3-da49-4e74-a3bd-da5b30a5461b"
#    value                = null
#  }
#  feature_tags {
#    custom_single_sign_on = false
#    enterprise            = false
#    gallery               = false
#    hide                  = false
#  }
#  optional_claims {
#    saml2_token {
#      additional_properties = []
#      essential             = false
#      name                  = "groups"
#      source                = null
#    }
#  }
#  public_client {
#    redirect_uris = []
#  }
#  single_page_application {
#    redirect_uris = []
#  }
#  timeouts {
#    create = null
#    delete = null
#    read   = null
#    update = null
#  }
#  web {
#    homepage_url  = "https://*.jfrog.io/artifactory/webapp/saml/loginResponse?metadata=jfrogartifactory|ISV9.2|primary|z"
#    logout_url    = "https://login.microsoftonline.com/cd99fef8-1cd3-4a2a-9bdf-15531181d65e/oauth2/v2.0/logout"
#    redirect_uris = ["https://dhlparcel.pe.jfrog.io/artifactory/webapp/saml/loginResponse"]
#    implicit_grant {
#      access_token_issuance_enabled = false
#      id_token_issuance_enabled     = true
#    }
#  }
#}
