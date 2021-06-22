package taplytics_tf

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceVariation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVariableRead,
		Schema: map[string]*schema.Schema{
			"userid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"experiment": {
				Type: schema.TypeString,
				Required: true,
			},
			"variation": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVariationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	provider := meta.(*Client)
	client := provider
	var userId = d.Get("userid").(string)
	var experiment = d.Get("experiment").(string)
	variation, err := client.getUserVariation(userId, experiment)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get variation",
			Detail:   "Failed to get variation: " + err.Error(),
		})
	}
	if err = d.Set("variation", variation); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to set key variation",
			Detail:   "Failed to set key with variation flattened data.",
		})
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}