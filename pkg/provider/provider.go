package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	clientall "github.com/taplytics/terraform-provider-taplytics/pkg/client-all"
	v3_client "github.com/taplytics/terraform-provider-taplytics/pkg/v3-client"

	"github.com/taplytics/terraform-provider-taplytics/pkg/datasources"
	"github.com/taplytics/terraform-provider-taplytics/pkg/resources"
	"github.com/taplytics/terraform-provider-taplytics/pkg/uapi-client"
)

type Client clientall.Client

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: providerConfigure,
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:     schema.TypeString,
				Required: true,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"taplytics_featureflags": datasources.DataSourceFeatureFlags(),
			"taplytics_featureflag":  datasources.DataSourceFeatureFlag(),
			"taplytics_bucketing":    datasources.DataSourceBucketing(),
			"taplytics_variation":    datasources.DataSourceVariation(),
			"taplytics_variables":    datasources.DataSourceVariables(),
			"taplytics_variable":     datasources.DataSourceVariable(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"taplytics_featureflag": resources.ResourceFeatureFlag(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	apitoken := d.Get("api_token").(string)
	var c = Client{
		Token: apitoken,
	}
	if apitoken != "" {
		c.UAPI = uapi_client.NewClient(apitoken)
		// V3 API token might be different
		c.V3API = v3_client.NewClient(apitoken)
		return c, diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to create Taplytics client",
		Detail:   "API Token is required.",
	})
	return nil, diags
}
