# \PeriodicTaskApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPeriodicTaskNames**](PeriodicTaskApi.md#GetPeriodicTaskNames) | **Get** /periodictask/names | Get comma-delimited list of all available periodic task names.
[**RunPeriodicTask**](PeriodicTaskApi.md#RunPeriodicTask) | **Get** /periodictask/run | Run periodic task against table. If table name is missing, task will run against all tables.



## GetPeriodicTaskNames

> []string GetPeriodicTaskNames(ctx).Execute()

Get comma-delimited list of all available periodic task names.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeriodicTaskApi.GetPeriodicTaskNames(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeriodicTaskApi.GetPeriodicTaskNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPeriodicTaskNames`: []string
    fmt.Fprintf(os.Stdout, "Response from `PeriodicTaskApi.GetPeriodicTaskNames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeriodicTaskNamesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunPeriodicTask

> RunPeriodicTask(ctx).Taskname(taskname).TableName(tableName).Type_(type_).Execute()

Run periodic task against table. If table name is missing, task will run against all tables.

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
    taskname := "taskname_example" // string | Periodic task name
    tableName := "tableName_example" // string | Name of the table (optional)
    type_ := "type__example" // string | OFFLINE | REALTIME (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PeriodicTaskApi.RunPeriodicTask(context.Background()).Taskname(taskname).TableName(tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeriodicTaskApi.RunPeriodicTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunPeriodicTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskname** | **string** | Periodic task name |
 **tableName** | **string** | Name of the table |
 **type_** | **string** | OFFLINE | REALTIME |

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

