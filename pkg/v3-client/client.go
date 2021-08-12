package v3_client

import (
	"context"
	"github.com/antihax/optional"
	taplytics "github.com/taplytics/go-sdk-v3"
	"github.com/taplytics/terraform-provider-taplytics/pkg/clients"
	"net/http"
)

type Client clients.Client

func NewClient(token string) *taplytics.APIClient {
	configuration := taplytics.NewConfiguration()
	// Set API key
	configuration.AddDefaultHeader("Authorization", "Bearer "+token)
	apiClient := taplytics.NewAPIClient(configuration)
	return apiClient
}

func (c *Client) GetFeatureFlags(status, paginationToken string, perPage int32) (flags []taplytics.FeatureFlag, err error) {
	resp, _, err := c.V3API.FeatureFlagsApi.GetFeatureFlags(context.Background(), &taplytics.FeatureFlagsApiGetFeatureFlagsOpts{
		Status:          optional.NewInterface(status),
		PaginationToken: optional.NewString(paginationToken),
		PerPage:         optional.NewInt32(perPage),
	})

	if err != nil {
		return nil, err
	}
	flags = resp.FeatureFlags
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
	flag, _, err = c.V3API.FeatureFlagsApi.CreateFeatureFlag(context.Background(), taplytics.FeatureFlagCreationBody{
		Name:        "",
		Key:         "",
		Description: "",
		Tags:        nil,
		Audience:    nil,
		Rollout:     nil,
	})

	if err != nil {
		return taplytics.FeatureFlag{}, err
	}
	return
}

func (c *Client) UpdateFeatureFlag(newFlag taplytics.FeatureFlag, featureKey string, updateAction string) (flag taplytics.FeatureFlag, err error) {
	flag, _, err = c.V3API.FeatureFlagsApi.UpdateFeatureFlag(context.Background(), taplytics.FeatureFlagUpdateBody{
		Name:        "",
		Description: "",
		Tags:        nil,
		Audience:    nil,
		Rollout:     nil,
	}, featureKey)
	if err != nil {
		return taplytics.FeatureFlag{}, err
	}
	return
}

func (c *Client) DeleteFeatureFlag(featureKey string) (response *http.Response, err error) {
	response, err = c.V3API.FeatureFlagsApi.DeleteFeatureFlag(context.Background(), featureKey)
	if err != nil {
		return nil, err
	}
	return
}
