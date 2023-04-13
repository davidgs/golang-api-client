# \UpsertApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EstimateHeapUsage**](UpsertApi.md#EstimateHeapUsage) | **Post** /upsert/estimateHeapUsage | Estimate memory usage for an upsert table



## EstimateHeapUsage

> string EstimateHeapUsage(ctx).Cardinality(cardinality).PrimaryKeySize(primaryKeySize).NumPartitions(numPartitions).Body(body).Execute()

Estimate memory usage for an upsert table



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
    cardinality := int64(789) // int64 | cardinality
    primaryKeySize := int32(56) // int32 | primaryKeySize (optional) (default to -1)
    numPartitions := int32(56) // int32 | numPartitions (optional) (default to -1)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpsertApi.EstimateHeapUsage(context.Background()).Cardinality(cardinality).PrimaryKeySize(primaryKeySize).NumPartitions(numPartitions).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpsertApi.EstimateHeapUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EstimateHeapUsage`: string
    fmt.Fprintf(os.Stdout, "Response from `UpsertApi.EstimateHeapUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateHeapUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardinality** | **int64** | cardinality | 
 **primaryKeySize** | **int32** | primaryKeySize | [default to -1]
 **numPartitions** | **int32** | numPartitions | [default to -1]
 **body** | **string** |  | 

### Return type

**string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

