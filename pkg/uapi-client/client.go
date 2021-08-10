package uapi_client

import (
	"context"
	"encoding/json"
	taplytics "github.com/taplytics/gosdk"
	_ "github.com/taplytics/terraform-provider-taplytics/pkg/client-all"
	client_all "github.com/taplytics/terraform-provider-taplytics/pkg/client-all"
	"io/ioutil"
	"strconv"
)

type FeatureFlag struct {
	Name    string `json:"name"`
	KeyName string `json:"keyName"`
}

type Client client_all.Client

func NewClient(token string) *taplytics.APIClient {
	configuration := taplytics.NewConfiguration()
	api_client := taplytics.NewAPIClient(configuration)
	return api_client
}

func (c *Client) GetFeatureFlags(userId string) (featureFlags []FeatureFlag, err error) {
	attributes := *taplytics.NewUserAttributes()
	customData := make(map[string]interface{}) // map[string]interface{} | serialized custom data object
	resp, err := c.UAPI.FeatureFlagsApi.FeatureflagsGet(context.Background()).
		Token(c.SDKToken).
		UserId(userId).
		CustomData(customData).
		Attributes(attributes).
		Execute()

	if err != nil {
		return nil, err
	}
	featureFlags = []FeatureFlag{}
	err = json.NewDecoder(resp.Body).Decode(&featureFlags)
	if err != nil {
		return nil, err
	}
	return
}

func (c *Client) IsFeatureFlagEnabled(userId string, key string) (enabled bool, err error) {
	attributes := *taplytics.NewUserAttributes()
	customData := make(map[string]interface{})
	resp, err := c.UAPI.FeatureFlagsApi.IsfeatureflagenabledGet(context.Background()).
		Token(c.SDKToken).
		UserId(userId).
		FeatureFlagKey(key).
		CustomData(customData).
		Attributes(attributes).
		Execute()

	if err != nil {
		return false, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	enabled, err = strconv.ParseBool(string(body))
	if err != nil {
		return false, err
	}
	return
}

func (c *Client) GetUserBucketing(userId string) (bucketing map[string]string, err error) {
	attributes := *taplytics.NewUserAttributes()
	customData := make(map[string]interface{})

	resp, err := c.UAPI.BucketingApi.BucketingGet(context.Background()).
		Token(c.SDKToken).
		UserId(userId).
		CustomData(customData).
		Attributes(attributes).
		Execute()

	if err != nil {
		return nil, err
	}

	bucketing = make(map[string]string)
	err = json.NewDecoder(resp.Body).Decode(&bucketing)
	if err != nil {
		return nil, err
	}
	return
}

func (c *Client) GetUserVariables(userId string) (variables []taplytics.Variable, err error) {
	attributes := *taplytics.NewUserAttributes()
	customData := make(map[string]interface{})
	variables, _, err = c.UAPI.VariablesApi.VariablesGet(context.Background()).Token(c.SDKToken).UserId(userId).CustomData(customData).Attributes(attributes).Execute()
	if err != nil {
		return nil, err
	}
	return
}
func (c *Client) GetUserVariable(userId, variableName string) (variable taplytics.Variable, err error) {

	attributes := *taplytics.NewUserAttributes()
	customData := make(map[string]interface{})
	resp, err := c.UAPI.VariablesApi.VariablevalueGet(context.Background()).Token(c.SDKToken).UserId(userId).CustomData(customData).Attributes(attributes).VarName(variableName).Execute()
	if err != nil {
		return variable, err
	}
	err = json.NewDecoder(resp.Body).Decode(&variable)
	if err != nil {
		return variable, err
	}
	return
}

func (c *Client) GetUserVariation(userId, experiment string) (variation string, err error) {
	attributes := *taplytics.NewUserAttributes()
	customData := make(map[string]interface{})
	resp, err := c.UAPI.VariationsApi.VariationGet(context.Background()).Token(c.SDKToken).UserId(userId).CustomData(customData).Attributes(attributes).ExperimentName(experiment).Execute()
	if err != nil {
		return "", err
	}
	_variation, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	variation = string(_variation)
	return
}
