terraform {
  required_providers {
    hygraph = {
      source = "registry.terraform.io/tavsec/hygraph"
    }
  }
}

provider "hygraph" {
  host = "https://management.hygraph.com/graphql"
  auth_token = "test123"
}

data "hygraph_webhooks" "example" {}