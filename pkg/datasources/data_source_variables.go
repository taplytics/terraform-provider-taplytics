package datasources

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	taplytics "github.com/taplytics/go-sdk"
	client2 "github.com/taplytics/terraform-provider-taplytics/pkg/client"
	"strconv"
	"time"
)

func DataSourceVariables() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVariablesRead,
		Schema: map[string]*schema.Schema{
			"userid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"variables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVariablesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	provider := meta.(*client2.Client)
	client := provider
	var userId = d.Get("userid").(string)
	resp, err := client.GetUserVariables(userId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get user variables",
			Detail:   "Failed to get user variables: " + err.Error(),
		})
	}
	enabledFlags := flattenVariablesData(&resp)
	if err = d.Set("variables", enabledFlags); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to set key variables",
			Detail:   "Failed to set key with variables flattened data.",
		})
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenVariablesData(variables *[]taplytics.Variable) []interface{} {
	if variables != nil {
		_flags := make([]interface{}, len(*variables), len(*variables))

		for i, flag := range *variables {
			f := make(map[string]interface{})

			f["name"] = flag.Name
			f["value"] = flag.Value

			_flags[i] = f
		}
		return _flags
	}
	return make([]interface{}, 0)
}
