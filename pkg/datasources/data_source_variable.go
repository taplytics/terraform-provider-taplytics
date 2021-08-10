package datasources

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client2 "github.com/taplytics/terraform-provider-taplytics/pkg/uapi-client"
	"strconv"
	"time"
)

func DataSourceVariable() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVariableRead,
		Schema: map[string]*schema.Schema{
			"userid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"variable_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVariableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	provider := meta.(*client2.Client)
	client := provider
	var userId = d.Get("userid").(string)
	var flagKey = d.Get("featureflag_key").(string)
	isEnabled, err := client.IsFeatureFlagEnabled(userId, flagKey)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get featureflag",
			Detail:   "Failed to get featureflag: " + err.Error(),
		})
	}
	if err = d.Set("enabled", isEnabled); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to set key enabled",
			Detail:   "Failed to set key with featureflags flattened data.",
		})
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
