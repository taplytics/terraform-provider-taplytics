package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	clientall "github.com/taplytics/terraform-provider-taplytics/pkg/clients"
	v3client "github.com/taplytics/terraform-provider-taplytics/pkg/v3-client"

	"github.com/taplytics/terraform-provider-taplytics/pkg/datasources"
	"github.com/taplytics/terraform-provider-taplytics/pkg/resources"
	"github.com/taplytics/terraform-provider-taplytics/pkg/uapi-client"
)

type Client clientall.Client

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: providerConfigure,
		Schema: map[string]*schema.Schema{
			"sdk_token": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rest_token": {
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

	var c = Client{
		SDKToken:  d.Get("sdk_token").(string),
		RESTToken: d.Get("rest_token").(string),
	}
	if c.SDKToken != "" && c.RESTToken != "" {
		c.UAPI = uapi_client.NewClient(c.SDKToken)
		// V3 API token might be different
		c.V3API = v3client.NewClient(c.RESTToken)
		return c, diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to create Taplytics client",
		Detail:   "REST Token and SDK Token are required.",
	})
	return nil, diags
}
