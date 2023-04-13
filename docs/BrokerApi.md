# \BrokerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBrokersForTable**](BrokerApi.md#GetBrokersForTable) | **Get** /brokers/tables/{tableName} | List brokers for a given table
[**GetBrokersForTableV2**](BrokerApi.md#GetBrokersForTableV2) | **Get** /v2/brokers/tables/{tableName} | List brokers for a given table
[**GetBrokersForTenant**](BrokerApi.md#GetBrokersForTenant) | **Get** /brokers/tenants/{tenantName} | List brokers for a given tenant
[**GetBrokersForTenantV2**](BrokerApi.md#GetBrokersForTenantV2) | **Get** /v2/brokers/tenants/{tenantName} | List brokers for a given tenant
[**GetTablesToBrokersMapping**](BrokerApi.md#GetTablesToBrokersMapping) | **Get** /brokers/tables | List tables to brokers mappings
[**GetTablesToBrokersMappingV2**](BrokerApi.md#GetTablesToBrokersMappingV2) | **Get** /v2/brokers/tables | List tables to brokers mappings
[**GetTenantsToBrokersMapping**](BrokerApi.md#GetTenantsToBrokersMapping) | **Get** /brokers/tenants | List tenants to brokers mappings
[**GetTenantsToBrokersMappingV2**](BrokerApi.md#GetTenantsToBrokersMappingV2) | **Get** /v2/brokers/tenants | List tenants to brokers mappings
[**ListBrokersMapping**](BrokerApi.md#ListBrokersMapping) | **Get** /brokers | List tenants and tables to brokers mappings
[**ListBrokersMappingV2**](BrokerApi.md#ListBrokersMappingV2) | **Get** /v2/brokers | List tenants and tables to brokers mappings
[**ToggleQueryRateLimiting**](BrokerApi.md#ToggleQueryRateLimiting) | **Post** /brokers/instances/{instanceName}/qps | Enable/disable the query rate limiting for a broker instance



## GetBrokersForTable

> []string GetBrokersForTable(ctx, tableName).Type_(type_).State(state).Execute()

List brokers for a given table



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
    type_ := "type__example" // string | OFFLINE|REALTIME (optional)
    state := "state_example" // string | ONLINE|OFFLINE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerApi.GetBrokersForTable(context.Background(), tableName).Type_(type_).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerApi.GetBrokersForTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrokersForTable`: []string
    fmt.Fprintf(os.Stdout, "Response from `BrokerApi.GetBrokersForTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrokersForTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |
 **state** | **string** | ONLINE|OFFLINE |

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


## GetBrokersForTableV2

> []InstanceInfo GetBrokersForTableV2(ctx, tableName).Type_(type_).State(state).Execute()

List brokers for a given table



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
    type_ := "type__example" // string | OFFLINE|REALTIME (optional)
    state := "state_example" // string | ONLINE|OFFLINE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerApi.GetBrokersForTableV2(context.Background(), tableName).Type_(type_).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerApi.GetBrokersForTableV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrokersForTableV2`: []InstanceInfo
    fmt.Fprintf(os.Stdout, "Response from `BrokerApi.GetBrokersForTableV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrokersForTableV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |
 **state** | **string** | ONLINE|OFFLINE |

### Return type

[**[]InstanceInfo**](InstanceInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrokersForTenant

> []string GetBrokersForTenant(ctx, tenantName).State(state).Execute()

List brokers for a given tenant



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
    tenantName := "tenantName_example" // string | Name of the tenant
    state := "state_example" // string | ONLINE|OFFLINE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerApi.GetBrokersForTenant(context.Background(), tenantName).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerApi.GetBrokersForTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrokersForTenant`: []string
    fmt.Fprintf(os.Stdout, "Response from `BrokerApi.GetBrokersForTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantName** | **string** | Name of the tenant |

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrokersForTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | ONLINE|OFFLINE |

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


## GetBrokersForTenantV2

> []InstanceInfo GetBrokersForTenantV2(ctx, tenantName).State(state).Execute()

List brokers for a given tenant



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
    tenantName := "tenantName_example" // string | Name of the tenant
    state := "state_example" // string | ONLINE|OFFLINE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerApi.GetBrokersForTenantV2(context.Background(), tenantName).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerApi.GetBrokersForTenantV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrokersForTenantV2`: []InstanceInfo
    fmt.Fprintf(os.Stdout, "Response from `BrokerApi.GetBrokersForTenantV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantName** | **string** | Name of the tenant |

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrokersForTenantV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | ONLINE|OFFLINE |

### Return type

[**[]InstanceInfo**](InstanceInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTablesToBrokersMapping

> map[string][]string GetTablesToBrokersMapping(ctx).State(state).Execute()

List tables to brokers mappings



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
    state := "state_example" // string | ONLINE|OFFLINE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerApi.GetTablesToBrokersMapping(context.Background()).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerApi.GetTablesToBrokersMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTablesToBrokersMapping`: map[string][]string
    fmt.Fprintf(os.Stdout, "Response from `BrokerApi.GetTablesToBrokersMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTablesToBrokersMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | ONLINE|OFFLINE |

### Return type

[**map[string][]string**](array.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTablesToBrokersMappingV2

> map[string][]InstanceInfo GetTablesToBrokersMappingV2(ctx).State(state).Execute()

List tables to brokers mappings



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
    state := "state_example" // string | ONLINE|OFFLINE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerApi.GetTablesToBrokersMappingV2(context.Background()).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerApi.GetTablesToBrokersMappingV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTablesToBrokersMappingV2`: map[string][]InstanceInfo
    fmt.Fprintf(os.Stdout, "Response from `BrokerApi.GetTablesToBrokersMappingV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTablesToBrokersMappingV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | ONLINE|OFFLINE |

### Return type

[**map[string][]InstanceInfo**](array.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantsToBrokersMapping

> map[string][]string GetTenantsToBrokersMapping(ctx).State(state).Execute()

List tenants to brokers mappings



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
    state := "state_example" // string | ONLINE|OFFLINE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerApi.GetTenantsToBrokersMapping(context.Background()).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerApi.GetTenantsToBrokersMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantsToBrokersMapping`: map[string][]string
    fmt.Fprintf(os.Stdout, "Response from `BrokerApi.GetTenantsToBrokersMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsToBrokersMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | ONLINE|OFFLINE |

### Return type

[**map[string][]string**](array.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantsToBrokersMappingV2

> map[string][]InstanceInfo GetTenantsToBrokersMappingV2(ctx).State(state).Execute()

List tenants to brokers mappings



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
    state := "state_example" // string | ONLINE|OFFLINE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerApi.GetTenantsToBrokersMappingV2(context.Background()).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerApi.GetTenantsToBrokersMappingV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantsToBrokersMappingV2`: map[string][]InstanceInfo
    fmt.Fprintf(os.Stdout, "Response from `BrokerApi.GetTenantsToBrokersMappingV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsToBrokersMappingV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | ONLINE|OFFLINE |

### Return type

[**map[string][]InstanceInfo**](array.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrokersMapping

> map[string]map[string][]string ListBrokersMapping(ctx).State(state).Execute()

List tenants and tables to brokers mappings



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
    state := "state_example" // string | ONLINE|OFFLINE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerApi.ListBrokersMapping(context.Background()).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerApi.ListBrokersMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrokersMapping`: map[string]map[string][]string
    fmt.Fprintf(os.Stdout, "Response from `BrokerApi.ListBrokersMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBrokersMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | ONLINE|OFFLINE |

### Return type

[**map[string]map[string][]string**](map.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrokersMappingV2

> map[string]map[string][]InstanceInfo ListBrokersMappingV2(ctx).State(state).Execute()

List tenants and tables to brokers mappings



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
    state := "state_example" // string | ONLINE|OFFLINE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerApi.ListBrokersMappingV2(context.Background()).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerApi.ListBrokersMappingV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrokersMappingV2`: map[string]map[string][]InstanceInfo
    fmt.Fprintf(os.Stdout, "Response from `BrokerApi.ListBrokersMappingV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBrokersMappingV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | ONLINE|OFFLINE |

### Return type

[**map[string]map[string][]InstanceInfo**](map.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleQueryRateLimiting

> ToggleQueryRateLimiting(ctx, instanceName).State(state).Execute()

Enable/disable the query rate limiting for a broker instance



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
    instanceName := "Broker_my.broker.com_30000" // string | Broker instance name
    state := "state_example" // string | ENABLE|DISABLE

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BrokerApi.ToggleQueryRateLimiting(context.Background(), instanceName).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerApi.ToggleQueryRateLimiting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Broker instance name |

### Other Parameters

Other parameters are passed through a pointer to a apiToggleQueryRateLimitingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | ENABLE|DISABLE |

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

