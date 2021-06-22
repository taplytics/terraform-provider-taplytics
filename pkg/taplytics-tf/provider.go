package taplytics_tf

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"taplytics_featureflags": dataSourceFeatureFlags(),
			"taplytics_featureflag": dataSourceFeatureFlag(),
			"taplytics_bucketing": dataSourceBucketing(),
			"taplytics_variation": dataSourceVariation(),
			"taplytics_variables": dataSourceVariables(),
			"taplytics_variable": dataSourceVariable(),
			//"taplytics_config": dataSourceConfig(), Not enabled due to the dynamic nature of it, and not being able to dump the full data to terraform.
		},

		ResourcesMap: map[string]*schema.Resource{
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	apitoken := d.Get("api_token").(string)

	if apitoken != "" {
		c := NewClient(apitoken)
		return c, diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to create Taplytics client",
		Detail:   "API Token is required.",
	})
	return nil, diags
}
