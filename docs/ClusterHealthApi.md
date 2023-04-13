# \ClusterHealthApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterHealthDetails**](ClusterHealthApi.md#GetClusterHealthDetails) | **Get** /clusterHealth | Get cached cluster health details



## GetClusterHealthDetails

> ClusterHealthResponse GetClusterHealthDetails(ctx).Execute()

Get cached cluster health details

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
    resp, r, err := apiClient.ClusterHealthApi.GetClusterHealthDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterHealthApi.GetClusterHealthDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterHealthDetails`: ClusterHealthResponse
    fmt.Fprintf(os.Stdout, "Response from `ClusterHealthApi.GetClusterHealthDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterHealthDetailsRequest struct via the builder pattern


### Return type

[**ClusterHealthResponse**](ClusterHealthResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

