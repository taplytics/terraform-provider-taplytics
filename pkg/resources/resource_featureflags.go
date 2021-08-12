package resources

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	taplytics "github.com/taplytics/go-sdk-v3"
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
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
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
							Optional: true,
						},
						"dataKey": {
							Type:     schema.TypeString,
							Optional: true,
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
	var draft = taplytics.DRAFT_FeatureFlagStatus

	v3 := m.(*client.Client)
	newFlag := taplytics.FeatureFlag{
		Rollout: &taplytics.Rollout{
			EndDate:         time.Time{},
			EndPercentage:   0,
			StartDate:       time.Time{},
			StartPercentage: 0,
		},
		Audience:    []taplytics.AudienceFilter{},
		Tags:        []string{},
		ShortId:     "",
		Status:      &draft,
		Description: "",
		Key:         "",
		Name:        "",
		Id:          "",
	}

	flag, err := v3.CreateFeatureFlag(newFlag)
	if err != nil {
		return nil
	}

	d.Set("featureflag", flag)

	if err != nil {
		return nil
	}
	return diags
}

func resourceFFRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	v3 := m.(*client.Client)
	var ffkey = d.Get("key").(string)

	featureflag, err := v3.GetFeatureFlag(ffkey)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get featureflag",
			Detail:   "Failed to get featureflag: " + err.Error(),
		})
	}
	if err = d.Set("featureflag", flattenFeatureFlag(&featureflag)); err != nil {
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
	provider := m.(*client.Client)

	resp, err := provider.DeleteFeatureFlag(d.Get("key").(string))
	if err != nil || resp.StatusCode != 200 {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Failed to delete featureflag.",
			Detail:        "An error occured in the request: " + err.Error(),
			AttributePath: nil,
		})
		return nil
	}

	return diags
}

func flattenFeatureFlag(flag *taplytics.FeatureFlag) (result map[string]interface{}) {
	if flag != nil {
		f := make(map[string]interface{})
		f["name"] = flag.Name
		f["key"] = flag.Key
		return f
	}

	return make(map[string]interface{}, 0)
}
