package datasources

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	uapi "github.com/taplytics/terraform-provider-taplytics/pkg/uapi-client"
	"strconv"
	"time"
)

func DataSourceFeatureFlags() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFeatureFlagsRead,
		Schema: map[string]*schema.Schema{
			"userid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"featureflags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"keyname": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFeatureFlagsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	provider := meta.(*uapi.Client)
	client := provider
	var userId = d.Get("userid").(string)
	resp, err := client.UAPI_GetFeatureFlags(userId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get featureflags",
			Detail:   "Failed to get featureflags: " + err.Error(),
		})
	}
	enabledFlags := flattenFeatureFlagsData(&resp)
	if err = d.Set("featureflags", enabledFlags); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to set key featureflags",
			Detail:   "Failed to set key with featureflags flattened data.",
		})
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenFeatureFlagsData(featureFlags *[]uapi.FeatureFlag) []interface{} {
	if featureFlags != nil {
		_flags := make([]interface{}, len(*featureFlags), len(*featureFlags))

		for i, flag := range *featureFlags {
			f := make(map[string]interface{})

			f["name"] = flag.Name
			f["keyname"] = flag.KeyName

			_flags[i] = f
		}
		return _flags
	}
	return make([]interface{}, 0)
}
