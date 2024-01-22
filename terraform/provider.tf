terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "< 3"
    }
    azuread = {
      source = "hashicorp/azuread"
      version = "2.14.0"
    }
  }
}

provider "azuread" {
  # Configuration options
}

provider "azurerm" {
  features {}
}
