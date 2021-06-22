package taplytics_tf

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceBucketing() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBucketingRead,
		Schema: map[string]*schema.Schema{
			"userid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucketing": {
				Type:     schema.TypeSet,
				Computed: true,
			},
		},
	}
}

func dataSourceBucketingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	provider := meta.(*Client)
	client := provider
	var userId = d.Get("userid").(string)
	bucketing, err := client.getUserBucketing(userId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get user bucketing",
			Detail:   "Failed to get user bucketing: " + err.Error(),
		})
	}
	if err = d.Set("bucketing", bucketing); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to set key enabled",
			Detail:   "Failed to set key with featureflags flattened data.",
		})
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
