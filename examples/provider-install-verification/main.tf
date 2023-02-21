terraform {
  required_providers {
    hygraph = {
      source = "registry.terraform.io/tavsec/hygraph"
    }
  }
}

provider "hygraph" {}

data "hygraph_webhook" "example" {}