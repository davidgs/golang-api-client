# \LoggerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadLogFile**](LoggerApi.md#DownloadLogFile) | **Get** /loggers/download | Download a log file
[**DownloadLogFileFromInstance**](LoggerApi.md#DownloadLogFileFromInstance) | **Get** /loggers/instances/{instanceName}/download | Download a log file from a given instance
[**GetLocalLogFiles**](LoggerApi.md#GetLocalLogFiles) | **Get** /loggers/files | Get all local log files
[**GetLogFilesFromAllInstances**](LoggerApi.md#GetLogFilesFromAllInstances) | **Get** /loggers/instances | Collect log files from all the instances
[**GetLogFilesFromInstance**](LoggerApi.md#GetLogFilesFromInstance) | **Get** /loggers/instances/{instanceName} | Collect log files from a given instance
[**GetLogger**](LoggerApi.md#GetLogger) | **Get** /loggers/{loggerName} | Get logger configs
[**GetLoggers**](LoggerApi.md#GetLoggers) | **Get** /loggers | Get all the loggers
[**SetLoggerLevel**](LoggerApi.md#SetLoggerLevel) | **Put** /loggers/{loggerName} | Set logger level



## DownloadLogFile

> DownloadLogFile(ctx).FilePath(filePath).Execute()

Download a log file

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
    filePath := "filePath_example" // string | Log file path

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LoggerApi.DownloadLogFile(context.Background()).FilePath(filePath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggerApi.DownloadLogFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLogFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filePath** | **string** | Log file path | 

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


## DownloadLogFileFromInstance

> DownloadLogFileFromInstance(ctx, instanceName).FilePath(filePath).Execute()

Download a log file from a given instance

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
    instanceName := "instanceName_example" // string | Instance Name
    filePath := "filePath_example" // string | Log file path

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LoggerApi.DownloadLogFileFromInstance(context.Background(), instanceName).FilePath(filePath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggerApi.DownloadLogFileFromInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Instance Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLogFileFromInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filePath** | **string** | Log file path | 

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


## GetLocalLogFiles

> []string GetLocalLogFiles(ctx).Execute()

Get all local log files

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
    resp, r, err := apiClient.LoggerApi.GetLocalLogFiles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggerApi.GetLocalLogFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocalLogFiles`: []string
    fmt.Fprintf(os.Stdout, "Response from `LoggerApi.GetLocalLogFiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalLogFilesRequest struct via the builder pattern


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


## GetLogFilesFromAllInstances

> map[string][]string GetLogFilesFromAllInstances(ctx).Execute()

Collect log files from all the instances

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
    resp, r, err := apiClient.LoggerApi.GetLogFilesFromAllInstances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggerApi.GetLogFilesFromAllInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogFilesFromAllInstances`: map[string][]string
    fmt.Fprintf(os.Stdout, "Response from `LoggerApi.GetLogFilesFromAllInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogFilesFromAllInstancesRequest struct via the builder pattern


### Return type

[**map[string][]string**](set.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogFilesFromInstance

> []string GetLogFilesFromInstance(ctx, instanceName).Execute()

Collect log files from a given instance

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
    instanceName := "instanceName_example" // string | Instance Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoggerApi.GetLogFilesFromInstance(context.Background(), instanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggerApi.GetLogFilesFromInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogFilesFromInstance`: []string
    fmt.Fprintf(os.Stdout, "Response from `LoggerApi.GetLogFilesFromInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Instance Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogFilesFromInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetLogger

> map[string]string GetLogger(ctx, loggerName).Execute()

Get logger configs



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
    loggerName := "loggerName_example" // string | Logger name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoggerApi.GetLogger(context.Background(), loggerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggerApi.GetLogger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogger`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `LoggerApi.GetLogger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loggerName** | **string** | Logger name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoggers

> []string GetLoggers(ctx).Execute()

Get all the loggers



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
    resp, r, err := apiClient.LoggerApi.GetLoggers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggerApi.GetLoggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoggers`: []string
    fmt.Fprintf(os.Stdout, "Response from `LoggerApi.GetLoggers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggersRequest struct via the builder pattern


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


## SetLoggerLevel

> map[string]string SetLoggerLevel(ctx, loggerName).Level(level).Execute()

Set logger level



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
    loggerName := "loggerName_example" // string | Logger name
    level := "level_example" // string | Logger level (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoggerApi.SetLoggerLevel(context.Background(), loggerName).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggerApi.SetLoggerLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetLoggerLevel`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `LoggerApi.SetLoggerLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loggerName** | **string** | Logger name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLoggerLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **level** | **string** | Logger level | 

### Return type

**map[string]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

