package datasources

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/taplytics/terraform-provider-taplytics/pkg/uapi-client"
	"strconv"
	"time"
)

func DataSourceBucketing() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBucketingRead,
		Schema: map[string]*schema.Schema{
			"userid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucketing": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"experiment": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"variation": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceBucketingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	provider := meta.(*uapi_client.Client)
	client := provider
	var userId = d.Get("userid").(string)
	bucketing, err := client.GetUserBucketing(userId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get user bucketing",
			Detail:   "Failed to get user bucketing: " + err.Error(),
		})
	}
	flattened := flattenBucketingData(&bucketing)
	if err = d.Set("bucketing", flattened); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to set key enabled",
			Detail:   "Failed to set key with featureflags flattened data.",
		})
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenBucketingData(bucketing *map[string]string) []interface{} {
	if bucketing != nil {
		_flags := make([]interface{}, len(*bucketing), len(*bucketing))
		i := 0
		for experiment, variation := range *bucketing {
			f := make(map[string]interface{})

			f["experiment"] = experiment
			f["variation"] = variation

			_flags[i] = f
			i++
		}
		return _flags
	}
	return make([]interface{}, 0)
}
