/*
 * Taplytics Universal API
 *
 * The Taplytics Universal API is an API to quickly use Taplytics features and functionality at edge. This API allows the utilization of the Taplytics experimentation  functionality anywhere in the stack, both client- and server-side.  Each call to the Universal API requires two main parameters: token and user_id. - token is unique to each Taplytics project and can be found under Settings -> Project Settings -> Taplytics SDK Key - user_id is your custom defined user ID.  To be able to utilize the Universal API, please ensure that your Taplytics project is setup for usage. You may contact your Account Manager or support@taplytics.com for any questions.
 *
 * API version: 1.1
 * Contact: support@taplytics.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package go_sdk

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// FeatureFlagsApiService FeatureFlagsApi service
type FeatureFlagsApiService service

type ApiFeatureflagsGetRequest struct {
	ctx        _context.Context
	ApiService *FeatureFlagsApiService
	token      *string
	userId     *string
	attributes *UserAttributes
	customData *map[string]interface{}
}

func (r ApiFeatureflagsGetRequest) Token(token string) ApiFeatureflagsGetRequest {
	r.token = &token
	return r
}
func (r ApiFeatureflagsGetRequest) UserId(userId string) ApiFeatureflagsGetRequest {
	r.userId = &userId
	return r
}
func (r ApiFeatureflagsGetRequest) Attributes(attributes UserAttributes) ApiFeatureflagsGetRequest {
	r.attributes = &attributes
	return r
}
func (r ApiFeatureflagsGetRequest) CustomData(customData map[string]interface{}) ApiFeatureflagsGetRequest {
	r.customData = &customData
	return r
}

func (r ApiFeatureflagsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FeatureflagsGetExecute(r)
}

/*
 * FeatureflagsGet Get enabled Feature Flags for the user
 * Returns a list of objects that have the name and the key name of the enabled Feature Flags for the user.  For more information on setting up feature flags, you may visit our Launch Control documentation: https://docs.taplytics.com/docs/guides-feature-flag
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFeatureflagsGetRequest
 */
func (a *FeatureFlagsApiService) FeatureflagsGet(ctx _context.Context) ApiFeatureflagsGetRequest {
	return ApiFeatureflagsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FeatureFlagsApiService) FeatureflagsGetExecute(r ApiFeatureflagsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagsApiService.FeatureflagsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/featureflags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.token == nil {
		return nil, reportError("token is required and must be specified")
	}
	if r.userId == nil {
		return nil, reportError("userId is required and must be specified")
	}
	if r.attributes == nil {
		return nil, reportError("attributes is required and must be specified")
	}
	if r.customData == nil {
		return nil, reportError("customData is required and must be specified")
	}

	localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	localVarQueryParams.Add("user_id", parameterToString(*r.userId, ""))
	localVarQueryParams.Add("attributes", parameterToString(*r.attributes, ""))
	localVarQueryParams.Add("customData", parameterToString(*r.customData, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFeatureflagsPostRequest struct {
	ctx                          _context.Context
	ApiService                   *FeatureFlagsApiService
	token                        *string
	userId                       *string
	userAttributesWithCustomData *UserAttributesWithCustomData
}

func (r ApiFeatureflagsPostRequest) Token(token string) ApiFeatureflagsPostRequest {
	r.token = &token
	return r
}
func (r ApiFeatureflagsPostRequest) UserId(userId string) ApiFeatureflagsPostRequest {
	r.userId = &userId
	return r
}
func (r ApiFeatureflagsPostRequest) UserAttributesWithCustomData(userAttributesWithCustomData UserAttributesWithCustomData) ApiFeatureflagsPostRequest {
	r.userAttributesWithCustomData = &userAttributesWithCustomData
	return r
}

func (r ApiFeatureflagsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FeatureflagsPostExecute(r)
}

/*
 * FeatureflagsPost Get enabled Feature Flags for the user
 * Returns a list of objects that have the name and the key name of the enabled Feature Flags for the user.  For more information on setting up feature flags, you may visit our Launch Control documentation: https://docs.taplytics.com/docs/guides-feature-flag
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFeatureflagsPostRequest
 */
func (a *FeatureFlagsApiService) FeatureflagsPost(ctx _context.Context) ApiFeatureflagsPostRequest {
	return ApiFeatureflagsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FeatureFlagsApiService) FeatureflagsPostExecute(r ApiFeatureflagsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagsApiService.FeatureflagsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/featureflags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.token == nil {
		return nil, reportError("token is required and must be specified")
	}
	if r.userId == nil {
		return nil, reportError("userId is required and must be specified")
	}

	localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	localVarQueryParams.Add("user_id", parameterToString(*r.userId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/jason"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.userAttributesWithCustomData
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIsfeatureflagenabledGetRequest struct {
	ctx            _context.Context
	ApiService     *FeatureFlagsApiService
	token          *string
	userId         *string
	featureFlagKey *string
	attributes     *UserAttributes
	customData     *map[string]interface{}
}

func (r ApiIsfeatureflagenabledGetRequest) Token(token string) ApiIsfeatureflagenabledGetRequest {
	r.token = &token
	return r
}
func (r ApiIsfeatureflagenabledGetRequest) UserId(userId string) ApiIsfeatureflagenabledGetRequest {
	r.userId = &userId
	return r
}
func (r ApiIsfeatureflagenabledGetRequest) FeatureFlagKey(featureFlagKey string) ApiIsfeatureflagenabledGetRequest {
	r.featureFlagKey = &featureFlagKey
	return r
}
func (r ApiIsfeatureflagenabledGetRequest) Attributes(attributes UserAttributes) ApiIsfeatureflagenabledGetRequest {
	r.attributes = &attributes
	return r
}
func (r ApiIsfeatureflagenabledGetRequest) CustomData(customData map[string]interface{}) ApiIsfeatureflagenabledGetRequest {
	r.customData = &customData
	return r
}

func (r ApiIsfeatureflagenabledGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IsfeatureflagenabledGetExecute(r)
}

/*
 * IsfeatureflagenabledGet Get if feature flag is enabled
 * Returns true or false based on if the Feature Flag key passed in maps to a Feature Flag that is currently enabled for the user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIsfeatureflagenabledGetRequest
 */
func (a *FeatureFlagsApiService) IsfeatureflagenabledGet(ctx _context.Context) ApiIsfeatureflagenabledGetRequest {
	return ApiIsfeatureflagenabledGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FeatureFlagsApiService) IsfeatureflagenabledGetExecute(r ApiIsfeatureflagenabledGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagsApiService.IsfeatureflagenabledGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/isfeatureflagenabled"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.token == nil {
		return nil, reportError("token is required and must be specified")
	}
	if r.userId == nil {
		return nil, reportError("userId is required and must be specified")
	}
	if r.featureFlagKey == nil {
		return nil, reportError("featureFlagKey is required and must be specified")
	}
	if r.attributes == nil {
		return nil, reportError("attributes is required and must be specified")
	}
	if r.customData == nil {
		return nil, reportError("customData is required and must be specified")
	}

	localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	localVarQueryParams.Add("user_id", parameterToString(*r.userId, ""))
	localVarQueryParams.Add("featureFlagKey", parameterToString(*r.featureFlagKey, ""))
	localVarQueryParams.Add("attributes", parameterToString(*r.attributes, ""))
	localVarQueryParams.Add("customData", parameterToString(*r.customData, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIsfeatureflagenabledPostRequest struct {
	ctx                          _context.Context
	ApiService                   *FeatureFlagsApiService
	token                        *string
	userId                       *string
	featureFlagKey               *string
	userAttributesWithCustomData *UserAttributesWithCustomData
}

func (r ApiIsfeatureflagenabledPostRequest) Token(token string) ApiIsfeatureflagenabledPostRequest {
	r.token = &token
	return r
}
func (r ApiIsfeatureflagenabledPostRequest) UserId(userId string) ApiIsfeatureflagenabledPostRequest {
	r.userId = &userId
	return r
}
func (r ApiIsfeatureflagenabledPostRequest) FeatureFlagKey(featureFlagKey string) ApiIsfeatureflagenabledPostRequest {
	r.featureFlagKey = &featureFlagKey
	return r
}
func (r ApiIsfeatureflagenabledPostRequest) UserAttributesWithCustomData(userAttributesWithCustomData UserAttributesWithCustomData) ApiIsfeatureflagenabledPostRequest {
	r.userAttributesWithCustomData = &userAttributesWithCustomData
	return r
}

func (r ApiIsfeatureflagenabledPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IsfeatureflagenabledPostExecute(r)
}

/*
 * IsfeatureflagenabledPost Get if feature flag is enabled
 * Returns true or false based on if the Feature Flag key passed in maps to a Feature Flag that is currently enabled for the user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIsfeatureflagenabledPostRequest
 */
func (a *FeatureFlagsApiService) IsfeatureflagenabledPost(ctx _context.Context) ApiIsfeatureflagenabledPostRequest {
	return ApiIsfeatureflagenabledPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *FeatureFlagsApiService) IsfeatureflagenabledPostExecute(r ApiIsfeatureflagenabledPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagsApiService.IsfeatureflagenabledPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/isfeatureflagenabled"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.token == nil {
		return nil, reportError("token is required and must be specified")
	}
	if r.userId == nil {
		return nil, reportError("userId is required and must be specified")
	}
	if r.featureFlagKey == nil {
		return nil, reportError("featureFlagKey is required and must be specified")
	}

	localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	localVarQueryParams.Add("user_id", parameterToString(*r.userId, ""))
	localVarQueryParams.Add("featureFlagKey", parameterToString(*r.featureFlagKey, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/jason"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.userAttributesWithCustomData
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
