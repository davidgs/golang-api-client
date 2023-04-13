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
)


// UpsertApiService UpsertApi service
type UpsertApiService service

type ApiEstimateHeapUsageRequest struct {
	ctx context.Context
	ApiService *UpsertApiService
	cardinality *int64
	primaryKeySize *int32
	numPartitions *int32
	body *string
}

// cardinality
func (r ApiEstimateHeapUsageRequest) Cardinality(cardinality int64) ApiEstimateHeapUsageRequest {
	r.cardinality = &cardinality
	return r
}

// primaryKeySize
func (r ApiEstimateHeapUsageRequest) PrimaryKeySize(primaryKeySize int32) ApiEstimateHeapUsageRequest {
	r.primaryKeySize = &primaryKeySize
	return r
}

// numPartitions
func (r ApiEstimateHeapUsageRequest) NumPartitions(numPartitions int32) ApiEstimateHeapUsageRequest {
	r.numPartitions = &numPartitions
	return r
}

func (r ApiEstimateHeapUsageRequest) Body(body string) ApiEstimateHeapUsageRequest {
	r.body = &body
	return r
}

func (r ApiEstimateHeapUsageRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.EstimateHeapUsageExecute(r)
}

/*
EstimateHeapUsage Estimate memory usage for an upsert table

This API returns the estimated heap usage based on primary key column stats. This allows us to estimate table size before onboarding.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEstimateHeapUsageRequest
*/
func (a *UpsertApiService) EstimateHeapUsage(ctx context.Context) ApiEstimateHeapUsageRequest {
	return ApiEstimateHeapUsageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *UpsertApiService) EstimateHeapUsageExecute(r ApiEstimateHeapUsageRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpsertApiService.EstimateHeapUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/upsert/estimateHeapUsage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cardinality == nil {
		return localVarReturnValue, nil, reportError("cardinality is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "cardinality", r.cardinality, "")
	if r.primaryKeySize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "primaryKeySize", r.primaryKeySize, "")
	}
	if r.numPartitions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "numPartitions", r.numPartitions, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
