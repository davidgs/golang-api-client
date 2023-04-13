# \TunerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TuneTable**](TunerApi.md#TuneTable) | **Get** /tuner/{tableName} | Apply tuner(s) to a table
[**TuneTable1**](TunerApi.md#TuneTable1) | **Post** /tuner/{tableName} | Apply specific tuner to a table



## TuneTable

> string TuneTable(ctx, tableName).Execute()

Apply tuner(s) to a table

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    tableName := "tableName_example" // string | Name of the table

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TunerApi.TuneTable(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TunerApi.TuneTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TuneTable`: string
    fmt.Fprintf(os.Stdout, "Response from `TunerApi.TuneTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiTuneTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TuneTable1

> string TuneTable1(ctx, tableName).Body(body).Execute()

Apply specific tuner to a table

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    tableName := "tableName_example" // string | Name of the table
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TunerApi.TuneTable1(context.Background(), tableName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TunerApi.TuneTable1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TuneTable1`: string
    fmt.Fprintf(os.Stdout, "Response from `TunerApi.TuneTable1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiTuneTable1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  |

### Return type

**string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

