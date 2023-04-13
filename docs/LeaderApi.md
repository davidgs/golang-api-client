# \LeaderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLeaderForTable**](LeaderApi.md#GetLeaderForTable) | **Get** /leader/tables/{tableName} | Gets leader for a given table
[**GetLeadersForAllTables**](LeaderApi.md#GetLeadersForAllTables) | **Get** /leader/tables | Gets leaders for all tables



## GetLeaderForTable

> LeadControllerResponse GetLeaderForTable(ctx, tableName).Execute()

Gets leader for a given table



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
    tableName := "tableName_example" // string | Table name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LeaderApi.GetLeaderForTable(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LeaderApi.GetLeaderForTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLeaderForTable`: LeadControllerResponse
    fmt.Fprintf(os.Stdout, "Response from `LeaderApi.GetLeaderForTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderForTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LeadControllerResponse**](LeadControllerResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeadersForAllTables

> LeadControllerResponse GetLeadersForAllTables(ctx).Execute()

Gets leaders for all tables



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LeaderApi.GetLeadersForAllTables(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LeaderApi.GetLeadersForAllTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLeadersForAllTables`: LeadControllerResponse
    fmt.Fprintf(os.Stdout, "Response from `LeaderApi.GetLeadersForAllTables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeadersForAllTablesRequest struct via the builder pattern


### Return type

[**LeadControllerResponse**](LeadControllerResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

