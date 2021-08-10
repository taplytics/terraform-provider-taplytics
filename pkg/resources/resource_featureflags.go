package resources

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client "github.com/taplytics/terraform-provider-taplytics/pkg/v3-client"
	"strconv"
	"time"
)

func ResourceFeatureFlag() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFFCreate,
		ReadContext:   resourceFFRead,
		UpdateContext: resourceFFUpdate,
		DeleteContext: resourceFFDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Required: false,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"audience": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"comparator": {
							Type:     schema.TypeString,
							Required: false,
						},
						"dataKey": {
							Type:     schema.TypeString,
							Required: false,
						},
						"values": {
							Type:     schema.TypeList,
							Required: true,
							// Need to figure out a better solution for this, as the values list doesn't have a strict type contract for the API.
							//Elem: &schema.Schema{Type: schema.Type}
						},
					},
				},
			},
			"rollout": {
				Type: schema.TypeMap,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"startPercentage": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"startDate": {
							Type:     schema.TypeString,
							Required: true,
						},
						"endPercentage": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"endDate": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Required: true,
			},
			"featureflag": {
				Computed: true,
				Type:     schema.TypeMap,
			},
		},
	}
}

func resourceFFCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	provider := m.(*client.Client)
	client := provider
	client.GetFeatureFlag("")

	return diags
}

func resourceFFRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	provider := m.(*client.Client)
	client := provider
	var ffkey = d.Get("key").(string)

	featureflag, err := client.GetFeatureFlag(ffkey)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get featureflag",
			Detail:   "Failed to get featureflag: " + err.Error(),
		})
	}
	if err = d.Set("featureflag", featureflag); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to set key variation",
			Detail:   "Failed to set key with variation flattened data.",
		})
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func resourceFFUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	return resourceFFRead(ctx, d, m)
}

func resourceFFDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
