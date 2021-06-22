package taplytics_tf

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceFeatureFlag() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFeatureFlagRead,
		Schema: map[string]*schema.Schema{
			"userid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"featureflag_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceFeatureFlagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	provider := meta.(*Client)
	client := provider
	var userId = d.Get("userid").(string)
	var flagKey = d.Get("featureflag_key").(string)
	isEnabled, err := client.isFeatureFlagEnabled(userId, flagKey)
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