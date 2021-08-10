package clients

import (
	v3sdk "github.com/taplytics/go-sdk-v3"
	uapisdk "github.com/taplytics/gosdk"
)

type Client struct {
	SDKToken  string
	RESTToken string
	UAPI      *uapisdk.APIClient
	V3API     *v3sdk.APIClient
}
