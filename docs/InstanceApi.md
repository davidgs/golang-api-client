# \InstanceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInstance**](InstanceApi.md#AddInstance) | **Post** /instances | Create a new instance
[**DropInstance**](InstanceApi.md#DropInstance) | **Delete** /instances/{instanceName} | Drop an instance
[**GetAllInstances**](InstanceApi.md#GetAllInstances) | **Get** /instances | List all instances
[**GetInstance**](InstanceApi.md#GetInstance) | **Get** /instances/{instanceName} | Get instance information
[**ToggleInstanceState**](InstanceApi.md#ToggleInstanceState) | **Post** /instances/{instanceName}/state | Enable/disable/drop an instance
[**UpdateBrokerResource**](InstanceApi.md#UpdateBrokerResource) | **Post** /instances/{instanceName}/updateBrokerResource | Update the tables served by the specified broker instance in the broker resource
[**UpdateInstance**](InstanceApi.md#UpdateInstance) | **Put** /instances/{instanceName} | Update the specified instance
[**UpdateInstanceTags**](InstanceApi.md#UpdateInstanceTags) | **Put** /instances/{instanceName}/updateTags | Update the tags of the specified instance



## AddInstance

> AddInstance(ctx).UpdateBrokerResource(updateBrokerResource).Body(body).Execute()

Create a new instance



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
    updateBrokerResource := true // bool | Whether to update broker resource for broker instance (optional) (default to false)
    body := *openapiclient.NewInstance("Host_example", int32(123), "Type_example") // Instance |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstanceApi.AddInstance(context.Background()).UpdateBrokerResource(updateBrokerResource).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.AddInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateBrokerResource** | **bool** | Whether to update broker resource for broker instance | [default to false]
 **body** | [**Instance**](Instance.md) |  |

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DropInstance

> DropInstance(ctx, instanceName).Execute()

Drop an instance



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
    instanceName := "Server_a.b.com_20000 | Broker_my.broker.com_30000" // string | Instance name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstanceApi.DropInstance(context.Background(), instanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.DropInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Instance name |

### Other Parameters

Other parameters are passed through a pointer to a apiDropInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAllInstances

> GetAllInstances(ctx).Execute()

List all instances

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
    r, err := apiClient.InstanceApi.GetAllInstances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.GetAllInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllInstancesRequest struct via the builder pattern


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


## GetInstance

> GetInstance(ctx, instanceName).Execute()

Get instance information

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
    instanceName := "Server_a.b.com_20000 | Broker_my.broker.com_30000" // string | Instance name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstanceApi.GetInstance(context.Background(), instanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.GetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Instance name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ToggleInstanceState

> ToggleInstanceState(ctx, instanceName).Body(body).Execute()

Enable/disable/drop an instance



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
    instanceName := "Server_a.b.com_20000 | Broker_my.broker.com_30000" // string | Instance name
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstanceApi.ToggleInstanceState(context.Background(), instanceName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.ToggleInstanceState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Instance name |

### Other Parameters

Other parameters are passed through a pointer to a apiToggleInstanceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  |

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBrokerResource

> UpdateBrokerResource(ctx, instanceName).Execute()

Update the tables served by the specified broker instance in the broker resource



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
    instanceName := "Broker_my.broker.com_30000" // string | Instance name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstanceApi.UpdateBrokerResource(context.Background(), instanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.UpdateBrokerResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Instance name |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrokerResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateInstance

> UpdateInstance(ctx, instanceName).UpdateBrokerResource(updateBrokerResource).Body(body).Execute()

Update the specified instance



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
    instanceName := "Server_a.b.com_20000 | Broker_my.broker.com_30000" // string | Instance name
    updateBrokerResource := true // bool | Whether to update broker resource for broker instance (optional) (default to false)
    body := *openapiclient.NewInstance("Host_example", int32(123), "Type_example") // Instance |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstanceApi.UpdateInstance(context.Background(), instanceName).UpdateBrokerResource(updateBrokerResource).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.UpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Instance name |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBrokerResource** | **bool** | Whether to update broker resource for broker instance | [default to false]
 **body** | [**Instance**](Instance.md) |  |

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstanceTags

> UpdateInstanceTags(ctx, instanceName).Tags(tags).UpdateBrokerResource(updateBrokerResource).Execute()

Update the tags of the specified instance



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
    instanceName := "Server_a.b.com_20000 | Broker_my.broker.com_30000" // string | Instance name
    tags := "tags_example" // string | Comma separated tags list
    updateBrokerResource := true // bool | Whether to update broker resource for broker instance (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstanceApi.UpdateInstanceTags(context.Background(), instanceName).Tags(tags).UpdateBrokerResource(updateBrokerResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.UpdateInstanceTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Instance name |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tags** | **string** | Comma separated tags list |
 **updateBrokerResource** | **bool** | Whether to update broker resource for broker instance | [default to false]

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

