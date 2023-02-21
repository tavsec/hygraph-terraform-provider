package hygraph

import (
	"context"
	hygraph "tavsec/hygraph/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &webhooksDataSource{}
	_ datasource.DataSourceWithConfigure = &webhooksDataSource{}
)

func NewWebhooksDataSource() datasource.DataSource {
	return &webhooksDataSource{}
}

type webhooksDataSource struct {
	client *hygraph.Client
}
type webhooksDataSourceModel struct {
	Webhooks []webhookModel `tfsdk:"webhooks"`
}

type webhookModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (d *webhooksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhooks"
}

func (d *webhooksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"webhooks": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *webhooksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state webhooksDataSourceModel

	webhooks, err := d.client.GetWebhooks()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read HyGraph Webhooks",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, webhook := range webhooks {
		webhookState := webhookModel{
			ID:   types.StringValue(webhook.ID),
			Name: types.StringValue(webhook.Name),
		}

		state.Webhooks = append(state.Webhooks, webhookState)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *webhooksDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*hygraph.Client)
}
