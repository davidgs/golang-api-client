/*
Pinot Controller API

APIs for accessing Pinot Controller information

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// TenantApiService TenantApi service
type TenantApiService service

type ApiChangeTenantStateRequest struct {
	ctx context.Context
	ApiService *TenantApiService
	tenantName string
	state *string
	type_ *string
}

// state
func (r ApiChangeTenantStateRequest) State(state string) ApiChangeTenantStateRequest {
	r.state = &state
	return r
}

// tenant type
func (r ApiChangeTenantStateRequest) Type_(type_ string) ApiChangeTenantStateRequest {
	r.type_ = &type_
	return r
}

func (r ApiChangeTenantStateRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.ChangeTenantStateExecute(r)
}

/*
ChangeTenantState Change tenant state

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantName Tenant name
 @return ApiChangeTenantStateRequest
*/
func (a *TenantApiService) ChangeTenantState(ctx context.Context, tenantName string) ApiChangeTenantStateRequest {
	return ApiChangeTenantStateRequest{
		ApiService: a,
		ctx: ctx,
		tenantName: tenantName,
	}
}

// Execute executes the request
//  @return string
func (a *TenantApiService) ChangeTenantStateExecute(r ApiChangeTenantStateRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantApiService.ChangeTenantState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tenants/{tenantName}/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.state == nil {
		return localVarReturnValue, nil, reportError("state is required and must be specified")
	}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["oauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateTenantRequest struct {
	ctx context.Context
	ApiService *TenantApiService
	body *Tenant
}

func (r ApiCreateTenantRequest) Body(body Tenant) ApiCreateTenantRequest {
	r.body = &body
	return r
}

func (r ApiCreateTenantRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateTenantExecute(r)
}

/*
CreateTenant  Create a tenant

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTenantRequest
*/
func (a *TenantApiService) CreateTenant(ctx context.Context) ApiCreateTenantRequest {
	return ApiCreateTenantRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TenantApiService) CreateTenantExecute(r ApiCreateTenantRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantApiService.CreateTenant")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tenants"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["oauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteTenantRequest struct {
	ctx context.Context
	ApiService *TenantApiService
	tenantName string
	type_ *string
}

// Tenant type
func (r ApiDeleteTenantRequest) Type_(type_ string) ApiDeleteTenantRequest {
	r.type_ = &type_
	return r
}

func (r ApiDeleteTenantRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTenantExecute(r)
}

/*
DeleteTenant Delete a tenant

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantName Tenant name
 @return ApiDeleteTenantRequest
*/
func (a *TenantApiService) DeleteTenant(ctx context.Context, tenantName string) ApiDeleteTenantRequest {
	return ApiDeleteTenantRequest{
		ApiService: a,
		ctx: ctx,
		tenantName: tenantName,
	}
}

// Execute executes the request
func (a *TenantApiService) DeleteTenantExecute(r ApiDeleteTenantRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantApiService.DeleteTenant")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tenants/{tenantName}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.type_ == nil {
		return nil, reportError("type_ is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["oauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetAllTenantsRequest struct {
	ctx context.Context
	ApiService *TenantApiService
	type_ *string
}

// Tenant type
func (r ApiGetAllTenantsRequest) Type_(type_ string) ApiGetAllTenantsRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetAllTenantsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetAllTenantsExecute(r)
}

/*
GetAllTenants List all tenants

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllTenantsRequest
*/
func (a *TenantApiService) GetAllTenants(ctx context.Context) ApiGetAllTenantsRequest {
	return ApiGetAllTenantsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TenantApiService) GetAllTenantsExecute(r ApiGetAllTenantsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantApiService.GetAllTenants")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tenants"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["oauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetTablesOnTenantRequest struct {
	ctx context.Context
	ApiService *TenantApiService
	tenantName string
}

func (r ApiGetTablesOnTenantRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetTablesOnTenantExecute(r)
}

/*
GetTablesOnTenant List tables on a a server tenant

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantName Tenant name
 @return ApiGetTablesOnTenantRequest
*/
func (a *TenantApiService) GetTablesOnTenant(ctx context.Context, tenantName string) ApiGetTablesOnTenantRequest {
	return ApiGetTablesOnTenantRequest{
		ApiService: a,
		ctx: ctx,
		tenantName: tenantName,
	}
}

// Execute executes the request
func (a *TenantApiService) GetTablesOnTenantExecute(r ApiGetTablesOnTenantRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantApiService.GetTablesOnTenant")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tenants/{tenantName}/tables"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["oauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetTenantMetadataRequest struct {
	ctx context.Context
	ApiService *TenantApiService
	tenantName string
	type_ *string
}

// tenant type
func (r ApiGetTenantMetadataRequest) Type_(type_ string) ApiGetTenantMetadataRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetTenantMetadataRequest) Execute() (*TenantMetadata, *http.Response, error) {
	return r.ApiService.GetTenantMetadataExecute(r)
}

/*
GetTenantMetadata Get tenant information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantName Tenant name
 @return ApiGetTenantMetadataRequest
*/
func (a *TenantApiService) GetTenantMetadata(ctx context.Context, tenantName string) ApiGetTenantMetadataRequest {
	return ApiGetTenantMetadataRequest{
		ApiService: a,
		ctx: ctx,
		tenantName: tenantName,
	}
}

// Execute executes the request
//  @return TenantMetadata
func (a *TenantApiService) GetTenantMetadataExecute(r ApiGetTenantMetadataRequest) (*TenantMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TenantMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantApiService.GetTenantMetadata")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tenants/{tenantName}/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["oauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListInstanceOrToggleTenantStateRequest struct {
	ctx context.Context
	ApiService *TenantApiService
	tenantName string
	type_ *string
	tableType *string
	state *string
}

// Tenant type (server|broker)
func (r ApiListInstanceOrToggleTenantStateRequest) Type_(type_ string) ApiListInstanceOrToggleTenantStateRequest {
	r.type_ = &type_
	return r
}

// Table type (offline|realtime)
func (r ApiListInstanceOrToggleTenantStateRequest) TableType(tableType string) ApiListInstanceOrToggleTenantStateRequest {
	r.tableType = &tableType
	return r
}

// state
func (r ApiListInstanceOrToggleTenantStateRequest) State(state string) ApiListInstanceOrToggleTenantStateRequest {
	r.state = &state
	return r
}

func (r ApiListInstanceOrToggleTenantStateRequest) Execute() (*http.Response, error) {
	return r.ApiService.ListInstanceOrToggleTenantStateExecute(r)
}

/*
ListInstanceOrToggleTenantState List instance for a tenant, or enable/disable/drop a tenant

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantName Tenant name
 @return ApiListInstanceOrToggleTenantStateRequest
*/
func (a *TenantApiService) ListInstanceOrToggleTenantState(ctx context.Context, tenantName string) ApiListInstanceOrToggleTenantStateRequest {
	return ApiListInstanceOrToggleTenantStateRequest{
		ApiService: a,
		ctx: ctx,
		tenantName: tenantName,
	}
}

// Execute executes the request
func (a *TenantApiService) ListInstanceOrToggleTenantStateExecute(r ApiListInstanceOrToggleTenantStateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantApiService.ListInstanceOrToggleTenantState")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tenants/{tenantName}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.tableType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tableType", r.tableType, "")
	}
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["oauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRebuildBrokerResourceRequest struct {
	ctx context.Context
	ApiService *TenantApiService
	tableName string
}

func (r ApiRebuildBrokerResourceRequest) Execute() (*http.Response, error) {
	return r.ApiService.RebuildBrokerResourceExecute(r)
}

/*
RebuildBrokerResource Rebuild broker resource for table

when new brokers are added

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tableName Table name (with type)
 @return ApiRebuildBrokerResourceRequest
*/
func (a *TenantApiService) RebuildBrokerResource(ctx context.Context, tableName string) ApiRebuildBrokerResourceRequest {
	return ApiRebuildBrokerResourceRequest{
		ApiService: a,
		ctx: ctx,
		tableName: tableName,
	}
}

// Execute executes the request
func (a *TenantApiService) RebuildBrokerResourceExecute(r ApiRebuildBrokerResourceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantApiService.RebuildBrokerResource")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tables/{tableName}/rebuildBrokerResourceFromHelixTags"
	localVarPath = strings.Replace(localVarPath, "{"+"tableName"+"}", url.PathEscape(parameterValueToString(r.tableName, "tableName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["oauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateTenantRequest struct {
	ctx context.Context
	ApiService *TenantApiService
	body *Tenant
}

func (r ApiUpdateTenantRequest) Body(body Tenant) ApiUpdateTenantRequest {
	r.body = &body
	return r
}

func (r ApiUpdateTenantRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateTenantExecute(r)
}

/*
UpdateTenant Update a tenant

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateTenantRequest
*/
func (a *TenantApiService) UpdateTenant(ctx context.Context) ApiUpdateTenantRequest {
	return ApiUpdateTenantRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TenantApiService) UpdateTenantExecute(r ApiUpdateTenantRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantApiService.UpdateTenant")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tenants"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["oauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
