package v3_client

import (
	"context"
	"github.com/antihax/optional"
	taplytics "github.com/taplytics/go-sdk-v3"
	client_all "github.com/taplytics/terraform-provider-taplytics/pkg/clients"
	"net/http"
)

type Client client_all.Client

func NewClient(token string) *taplytics.APIClient {
	configuration := taplytics.NewConfiguration()
	// Set API key
	configuration.AddDefaultHeader("Authorization", "Bearer "+token)
	apiClient := taplytics.NewAPIClient(configuration)
	return apiClient
}

func (c *Client) GetFeatureFlags(status string, pageNum int32) (flags []taplytics.FeatureFlag, err error) {
	flags, _, err = c.V3API.FeatureFlagsApi.GetFeatureFlags(context.Background(), &taplytics.FeatureFlagsApiGetFeatureFlagsOpts{
		Status:  optional.NewString(status),
		PageNum: optional.NewInt32(pageNum),
	})
	if err != nil {
		return nil, err
	}
	return
}

func (c *Client) GetFeatureFlag(featurekey string) (flag taplytics.FeatureFlag, err error) {
	flag, _, err = c.V3API.FeatureFlagsApi.GetFeatureFlag(context.Background(), featurekey)
	if err != nil {
		return taplytics.FeatureFlag{}, err
	}
	return
}

func (c *Client) CreateFeatureFlag(newFlag taplytics.FeatureFlag) (flag taplytics.FeatureFlag, err error) {
	flag, _, err = c.V3API.FeatureFlagsApi.CreateFeatureFlag(context.Background(), newFlag)

	if err != nil {
		return taplytics.FeatureFlag{}, err
	}
	return
}

func (c *Client) UpdateFeatureFlag(newFlag taplytics.FeatureFlag, featureKey string, updateAction string) (response *http.Response, err error) {
	response, err = c.V3API.FeatureFlagsApi.UpdateFeatureFlag(context.Background(), newFlag, featureKey, &taplytics.FeatureFlagsApiUpdateFeatureFlagOpts{Action: optional.NewString(updateAction)})
	if err != nil {
		return nil, err
	}
	return
}
