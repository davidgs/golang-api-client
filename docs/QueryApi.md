# \QueryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelQuery**](QueryApi.md#CancelQuery) | **Delete** /query/{brokerId}/{queryId} | Cancel a query as identified by the queryId
[**GetRunningQueries**](QueryApi.md#GetRunningQueries) | **Get** /queries | Get running queries from all brokers



## CancelQuery

> CancelQuery(ctx, brokerId, queryId).TimeoutMs(timeoutMs).Verbose(verbose).Execute()

Cancel a query as identified by the queryId



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    brokerId := "brokerId_example" // string | Broker that's running the query
    queryId := int64(789) // int64 | QueryId as assigned by the broker
    timeoutMs := int32(56) // int32 | Timeout for servers to respond the cancel request (optional) (default to 3000)
    verbose := true // bool | Return verbose responses for troubleshooting (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.QueryApi.CancelQuery(context.Background(), brokerId, queryId).TimeoutMs(timeoutMs).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.CancelQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brokerId** | **string** | Broker that&#39;s running the query | 
**queryId** | **int64** | QueryId as assigned by the broker | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeoutMs** | **int32** | Timeout for servers to respond the cancel request | [default to 3000]
 **verbose** | **bool** | Return verbose responses for troubleshooting | [default to false]

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunningQueries

> GetRunningQueries(ctx).TimeoutMs(timeoutMs).Execute()

Get running queries from all brokers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    timeoutMs := int32(56) // int32 | Timeout for brokers to return running queries (optional) (default to 3000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.QueryApi.GetRunningQueries(context.Background()).TimeoutMs(timeoutMs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.GetRunningQueries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRunningQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeoutMs** | **int32** | Timeout for brokers to return running queries | [default to 3000]

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

