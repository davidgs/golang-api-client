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


// QueryApiService QueryApi service
type QueryApiService service

type ApiCancelQueryRequest struct {
	ctx context.Context
	ApiService *QueryApiService
	brokerId string
	queryId int64
	timeoutMs *int32
	verbose *bool
}

// Timeout for servers to respond the cancel request
func (r ApiCancelQueryRequest) TimeoutMs(timeoutMs int32) ApiCancelQueryRequest {
	r.timeoutMs = &timeoutMs
	return r
}

// Return verbose responses for troubleshooting
func (r ApiCancelQueryRequest) Verbose(verbose bool) ApiCancelQueryRequest {
	r.verbose = &verbose
	return r
}

func (r ApiCancelQueryRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelQueryExecute(r)
}

/*
CancelQuery Cancel a query as identified by the queryId

No effect if no query exists for the given queryId on the requested broker. Query may continue to run for a short while after calling cancel as it's done in a non-blocking manner. The cancel method can be called multiple times.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param brokerId Broker that's running the query
 @param queryId QueryId as assigned by the broker
 @return ApiCancelQueryRequest
*/
func (a *QueryApiService) CancelQuery(ctx context.Context, brokerId string, queryId int64) ApiCancelQueryRequest {
	return ApiCancelQueryRequest{
		ApiService: a,
		ctx: ctx,
		brokerId: brokerId,
		queryId: queryId,
	}
}

// Execute executes the request
func (a *QueryApiService) CancelQueryExecute(r ApiCancelQueryRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryApiService.CancelQuery")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query/{brokerId}/{queryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"brokerId"+"}", url.PathEscape(parameterValueToString(r.brokerId, "brokerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queryId"+"}", url.PathEscape(parameterValueToString(r.queryId, "queryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.timeoutMs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeoutMs", r.timeoutMs, "")
	}
	if r.verbose != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "verbose", r.verbose, "")
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

type ApiGetRunningQueriesRequest struct {
	ctx context.Context
	ApiService *QueryApiService
	timeoutMs *int32
}

// Timeout for brokers to return running queries
func (r ApiGetRunningQueriesRequest) TimeoutMs(timeoutMs int32) ApiGetRunningQueriesRequest {
	r.timeoutMs = &timeoutMs
	return r
}

func (r ApiGetRunningQueriesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetRunningQueriesExecute(r)
}

/*
GetRunningQueries Get running queries from all brokers

The queries are returned with brokers running them

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRunningQueriesRequest
*/
func (a *QueryApiService) GetRunningQueries(ctx context.Context) ApiGetRunningQueriesRequest {
	return ApiGetRunningQueriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *QueryApiService) GetRunningQueriesExecute(r ApiGetRunningQueriesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryApiService.GetRunningQueries")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/queries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.timeoutMs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeoutMs", r.timeoutMs, "")
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