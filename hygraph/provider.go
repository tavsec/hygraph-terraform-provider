package hygraph

import (
	"context"
	"os"
	hygraph "tavsec/hygraph/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &hygraphProvider{}
)

type hygraphProviderModel struct {
	Host      types.String `tfsdk:"host"`
	AuthToken types.String `tfsdk:"auth_token"`
}

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &hygraphProvider{}
}

// hygraphProvider is the provider implementation.
type hygraphProvider struct{}

// Metadata returns the provider type name.
func (p *hygraphProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "hygraph"
}

// Schema defines the provider-level schema for configuration data.
func (p *hygraphProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Required: true,
			},
			"auth_token": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

// Configure prepares a HyGraph API client for data sources and resources.
func (p *hygraphProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config hygraphProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown HyGraph API Host",
			"The provider cannot create the HyGraph API client as there is an unknown configuration value for the HyGraph API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the HYGRAPH_HOST environment variable.",
		)
	}

	if config.AuthToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("auth_token"),
			"Unknown HyGraph API AuthToken",
			"The provider cannot create the HyGraph API client as there is an unknown configuration value for the HyGraph API Auth Token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the HYGRAPH_AUTH_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("HYGRAPH_HOST")
	auth_token := os.Getenv("HYGRAPH_AUTH_TOKEN")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.AuthToken.IsNull() {
		auth_token = config.AuthToken.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing HyGraph API Host",
			"The provider cannot create the HyGraph API client as there is a missing or empty value for the HyGraph API host. "+
				"Set the host value in the configuration or use the HASHICUPS_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if auth_token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("auth_token"),
			"Missing HyGraph API AuthToken",
			"The provider cannot create the HyGraph API client as there is a missing or empty value for the HyGraph API AuthToken. "+
				"Set the password value in the configuration or use the HASHICUPS_AUTH_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new HyGraph client using the configuration values
	client, err := hygraph.NewClient(&host, &auth_token)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create HyGraph API Client",
			"An unexpected error occurred when creating the HyGraph API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"HyGraph Client Error: "+err.Error(),
		)
		return
	}

	// Make the HyGraph client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *hygraphProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewWebhooksDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *hygraphProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
