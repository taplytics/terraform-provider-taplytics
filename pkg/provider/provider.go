package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/taplytics/terraform-provider-taplytics/pkg/client"
	"github.com/taplytics/terraform-provider-taplytics/pkg/datasources"
)

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
			//"taplytics_config": dataSourceConfig(), Not enabled due to the dynamic nature of it, and not being able to dump the full data to terraform.
		},

		ResourcesMap: map[string]*schema.Resource{},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	apitoken := d.Get("api_token").(string)

	if apitoken != "" {
		c := client.NewClient(apitoken)
		return c, diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to create Taplytics client",
		Detail:   "API Token is required.",
	})
	return nil, diags
}
