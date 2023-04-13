# \TenantApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeTenantState**](TenantApi.md#ChangeTenantState) | **Post** /tenants/{tenantName}/metadata | Change tenant state
[**CreateTenant**](TenantApi.md#CreateTenant) | **Post** /tenants |  Create a tenant
[**DeleteTenant**](TenantApi.md#DeleteTenant) | **Delete** /tenants/{tenantName} | Delete a tenant
[**GetAllTenants**](TenantApi.md#GetAllTenants) | **Get** /tenants | List all tenants
[**GetTablesOnTenant**](TenantApi.md#GetTablesOnTenant) | **Get** /tenants/{tenantName}/tables | List tables on a a server tenant
[**GetTenantMetadata**](TenantApi.md#GetTenantMetadata) | **Get** /tenants/{tenantName}/metadata | Get tenant information
[**ListInstanceOrToggleTenantState**](TenantApi.md#ListInstanceOrToggleTenantState) | **Get** /tenants/{tenantName} | List instance for a tenant, or enable/disable/drop a tenant
[**RebuildBrokerResource**](TenantApi.md#RebuildBrokerResource) | **Post** /tables/{tableName}/rebuildBrokerResourceFromHelixTags | Rebuild broker resource for table
[**UpdateTenant**](TenantApi.md#UpdateTenant) | **Put** /tenants | Update a tenant



## ChangeTenantState

> string ChangeTenantState(ctx, tenantName).State(state).Type_(type_).Execute()

Change tenant state

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
    tenantName := "tenantName_example" // string | Tenant name
    state := "state_example" // string | state
    type_ := "type__example" // string | tenant type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantApi.ChangeTenantState(context.Background(), tenantName).State(state).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.ChangeTenantState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeTenantState`: string
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.ChangeTenantState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantName** | **string** | Tenant name | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTenantStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | state | 
 **type_** | **string** | tenant type | 

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


## CreateTenant

> CreateTenant(ctx).Body(body).Execute()

 Create a tenant

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
    body := *openapiclient.NewTenant("TenantRole_example", "TenantName_example") // Tenant |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TenantApi.CreateTenant(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.CreateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Tenant**](Tenant.md) |  | 

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


## DeleteTenant

> DeleteTenant(ctx, tenantName).Type_(type_).Execute()

Delete a tenant

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
    tenantName := "tenantName_example" // string | Tenant name
    type_ := "type__example" // string | Tenant type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TenantApi.DeleteTenant(context.Background(), tenantName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.DeleteTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantName** | **string** | Tenant name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Tenant type | 

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


## GetAllTenants

> GetAllTenants(ctx).Type_(type_).Execute()

List all tenants

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
    type_ := "type__example" // string | Tenant type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TenantApi.GetAllTenants(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.GetAllTenants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Tenant type | 

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


## GetTablesOnTenant

> GetTablesOnTenant(ctx, tenantName).Execute()

List tables on a a server tenant

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
    tenantName := "tenantName_example" // string | Tenant name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TenantApi.GetTablesOnTenant(context.Background(), tenantName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.GetTablesOnTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantName** | **string** | Tenant name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTablesOnTenantRequest struct via the builder pattern


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


## GetTenantMetadata

> TenantMetadata GetTenantMetadata(ctx, tenantName).Type_(type_).Execute()

Get tenant information

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
    tenantName := "tenantName_example" // string | Tenant name
    type_ := "type__example" // string | tenant type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantApi.GetTenantMetadata(context.Background(), tenantName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.GetTenantMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantMetadata`: TenantMetadata
    fmt.Fprintf(os.Stdout, "Response from `TenantApi.GetTenantMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantName** | **string** | Tenant name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | tenant type | 

### Return type

[**TenantMetadata**](TenantMetadata.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstanceOrToggleTenantState

> ListInstanceOrToggleTenantState(ctx, tenantName).Type_(type_).TableType(tableType).State(state).Execute()

List instance for a tenant, or enable/disable/drop a tenant

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
    tenantName := "tenantName_example" // string | Tenant name
    type_ := "type__example" // string | Tenant type (server|broker) (optional)
    tableType := "tableType_example" // string | Table type (offline|realtime) (optional)
    state := "state_example" // string | state (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TenantApi.ListInstanceOrToggleTenantState(context.Background(), tenantName).Type_(type_).TableType(tableType).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.ListInstanceOrToggleTenantState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantName** | **string** | Tenant name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceOrToggleTenantStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Tenant type (server|broker) | 
 **tableType** | **string** | Table type (offline|realtime) | 
 **state** | **string** | state | 

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


## RebuildBrokerResource

> RebuildBrokerResource(ctx, tableName).Execute()

Rebuild broker resource for table



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
    tableName := "tableName_example" // string | Table name (with type)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TenantApi.RebuildBrokerResource(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.RebuildBrokerResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Table name (with type) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildBrokerResourceRequest struct via the builder pattern


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


## UpdateTenant

> UpdateTenant(ctx).Body(body).Execute()

Update a tenant

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
    body := *openapiclient.NewTenant("TenantRole_example", "TenantName_example") // Tenant |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TenantApi.UpdateTenant(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantApi.UpdateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Tenant**](Tenant.md) |  | 

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

