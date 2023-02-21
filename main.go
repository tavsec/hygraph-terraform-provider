package main

import (
	"context"
	"tavsec/hygraph/hygraph"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	providerserver.Serve(context.Background(), hygraph.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/tavsec/hygraph",
	})
}
