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

// VariablesApiService VariablesApi service
type VariablesApiService service

type ApiVariablesGetRequest struct {
	ctx        _context.Context
	ApiService *VariablesApiService
	token      *string
	userId     *string
	attributes *UserAttributes
	customData *map[string]interface{}
}

func (r ApiVariablesGetRequest) Token(token string) ApiVariablesGetRequest {
	r.token = &token
	return r
}
func (r ApiVariablesGetRequest) UserId(userId string) ApiVariablesGetRequest {
	r.userId = &userId
	return r
}
func (r ApiVariablesGetRequest) Attributes(attributes UserAttributes) ApiVariablesGetRequest {
	r.attributes = &attributes
	return r
}
func (r ApiVariablesGetRequest) CustomData(customData map[string]interface{}) ApiVariablesGetRequest {
	r.customData = &customData
	return r
}

func (r ApiVariablesGetRequest) Execute() ([]Variable, *_nethttp.Response, error) {
	return r.ApiService.VariablesGetExecute(r)
}

/*
 * VariablesGet Get all active variables for user
 * Returns all code variables and their values for the given user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiVariablesGetRequest
 */
func (a *VariablesApiService) VariablesGet(ctx _context.Context) ApiVariablesGetRequest {
	return ApiVariablesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return []Variable
 */
func (a *VariablesApiService) VariablesGetExecute(r ApiVariablesGetRequest) ([]Variable, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Variable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesApiService.VariablesGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
	}
	if r.userId == nil {
		return localVarReturnValue, nil, reportError("userId is required and must be specified")
	}
	if r.attributes == nil {
		return localVarReturnValue, nil, reportError("attributes is required and must be specified")
	}
	if r.customData == nil {
		return localVarReturnValue, nil, reportError("customData is required and must be specified")
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiVariablesPostRequest struct {
	ctx                          _context.Context
	ApiService                   *VariablesApiService
	token                        *string
	userId                       *string
	userAttributesWithCustomData *UserAttributesWithCustomData
}

func (r ApiVariablesPostRequest) Token(token string) ApiVariablesPostRequest {
	r.token = &token
	return r
}
func (r ApiVariablesPostRequest) UserId(userId string) ApiVariablesPostRequest {
	r.userId = &userId
	return r
}
func (r ApiVariablesPostRequest) UserAttributesWithCustomData(userAttributesWithCustomData UserAttributesWithCustomData) ApiVariablesPostRequest {
	r.userAttributesWithCustomData = &userAttributesWithCustomData
	return r
}

func (r ApiVariablesPostRequest) Execute() ([]Variable, *_nethttp.Response, error) {
	return r.ApiService.VariablesPostExecute(r)
}

/*
 * VariablesPost Get all active variables for user
 * Returns all code variables and their values for the given user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiVariablesPostRequest
 */
func (a *VariablesApiService) VariablesPost(ctx _context.Context) ApiVariablesPostRequest {
	return ApiVariablesPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return []Variable
 */
func (a *VariablesApiService) VariablesPostExecute(r ApiVariablesPostRequest) ([]Variable, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Variable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesApiService.VariablesPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
	}
	if r.userId == nil {
		return localVarReturnValue, nil, reportError("userId is required and must be specified")
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiVariablevalueGetRequest struct {
	ctx          _context.Context
	ApiService   *VariablesApiService
	token        *string
	userId       *string
	varName      *string
	defaultValue *string
	attributes   *UserAttributes
	customData   *map[string]interface{}
}

func (r ApiVariablevalueGetRequest) Token(token string) ApiVariablevalueGetRequest {
	r.token = &token
	return r
}
func (r ApiVariablevalueGetRequest) UserId(userId string) ApiVariablevalueGetRequest {
	r.userId = &userId
	return r
}
func (r ApiVariablevalueGetRequest) VarName(varName string) ApiVariablevalueGetRequest {
	r.varName = &varName
	return r
}
func (r ApiVariablevalueGetRequest) DefaultValue(defaultValue string) ApiVariablevalueGetRequest {
	r.defaultValue = &defaultValue
	return r
}
func (r ApiVariablevalueGetRequest) Attributes(attributes UserAttributes) ApiVariablevalueGetRequest {
	r.attributes = &attributes
	return r
}
func (r ApiVariablevalueGetRequest) CustomData(customData map[string]interface{}) ApiVariablevalueGetRequest {
	r.customData = &customData
	return r
}

func (r ApiVariablevalueGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.VariablevalueGetExecute(r)
}

/*
 * VariablevalueGet Get value for a Taplytics Variable
 * Value for given Taplytics code variable.  For more information on setting up code variables, you may visit our code experiment setup documentation: https://docs.taplytics.com/docs/guides-code-based-experiments
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiVariablevalueGetRequest
 */
func (a *VariablesApiService) VariablevalueGet(ctx _context.Context) ApiVariablevalueGetRequest {
	return ApiVariablevalueGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *VariablesApiService) VariablevalueGetExecute(r ApiVariablevalueGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesApiService.VariablevalueGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variablevalue"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.token == nil {
		return nil, reportError("token is required and must be specified")
	}
	if r.userId == nil {
		return nil, reportError("userId is required and must be specified")
	}
	if r.varName == nil {
		return nil, reportError("varName is required and must be specified")
	}
	if r.defaultValue == nil {
		return nil, reportError("defaultValue is required and must be specified")
	}
	if r.attributes == nil {
		return nil, reportError("attributes is required and must be specified")
	}
	if r.customData == nil {
		return nil, reportError("customData is required and must be specified")
	}

	localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	localVarQueryParams.Add("user_id", parameterToString(*r.userId, ""))
	localVarQueryParams.Add("varName", parameterToString(*r.varName, ""))
	localVarQueryParams.Add("defaultValue", parameterToString(*r.defaultValue, ""))
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

type ApiVariablevaluePostRequest struct {
	ctx                          _context.Context
	ApiService                   *VariablesApiService
	token                        *string
	userId                       *string
	varName                      *string
	defaultValue                 *string
	userAttributesWithCustomData *UserAttributesWithCustomData
}

func (r ApiVariablevaluePostRequest) Token(token string) ApiVariablevaluePostRequest {
	r.token = &token
	return r
}
func (r ApiVariablevaluePostRequest) UserId(userId string) ApiVariablevaluePostRequest {
	r.userId = &userId
	return r
}
func (r ApiVariablevaluePostRequest) VarName(varName string) ApiVariablevaluePostRequest {
	r.varName = &varName
	return r
}
func (r ApiVariablevaluePostRequest) DefaultValue(defaultValue string) ApiVariablevaluePostRequest {
	r.defaultValue = &defaultValue
	return r
}
func (r ApiVariablevaluePostRequest) UserAttributesWithCustomData(userAttributesWithCustomData UserAttributesWithCustomData) ApiVariablevaluePostRequest {
	r.userAttributesWithCustomData = &userAttributesWithCustomData
	return r
}

func (r ApiVariablevaluePostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.VariablevaluePostExecute(r)
}

/*
 * VariablevaluePost Get value for a Taplytics Variable
 * Value for given Taplytics code variable.  For more information on setting up code variables, you may visit our code experiment setup documentation: https://docs.taplytics.com/docs/guides-code-based-experiments
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiVariablevaluePostRequest
 */
func (a *VariablesApiService) VariablevaluePost(ctx _context.Context) ApiVariablevaluePostRequest {
	return ApiVariablevaluePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *VariablesApiService) VariablevaluePostExecute(r ApiVariablevaluePostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesApiService.VariablevaluePost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variablevalue"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.token == nil {
		return nil, reportError("token is required and must be specified")
	}
	if r.userId == nil {
		return nil, reportError("userId is required and must be specified")
	}
	if r.varName == nil {
		return nil, reportError("varName is required and must be specified")
	}
	if r.defaultValue == nil {
		return nil, reportError("defaultValue is required and must be specified")
	}

	localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	localVarQueryParams.Add("user_id", parameterToString(*r.userId, ""))
	localVarQueryParams.Add("varName", parameterToString(*r.varName, ""))
	localVarQueryParams.Add("defaultValue", parameterToString(*r.defaultValue, ""))
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
