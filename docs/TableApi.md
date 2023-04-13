# \TableApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConfig**](TableApi.md#AddConfig) | **Post** /tableConfigs | Add the TableConfigs using the tableConfigsStr json
[**AddTable**](TableApi.md#AddTable) | **Post** /tables | Adds a table
[**AlterTableStateOrListTableConfig**](TableApi.md#AlterTableStateOrListTableConfig) | **Get** /tables/{tableName} | Get/Enable/Disable/Drop a table
[**AssignInstances**](TableApi.md#AssignInstances) | **Post** /tables/{tableName}/assignInstances | Assign server instances to a table
[**CheckTableConfig**](TableApi.md#CheckTableConfig) | **Post** /tables/validate | Validate table config for a table
[**DeleteConfig**](TableApi.md#DeleteConfig) | **Delete** /tableConfigs/{tableName} | Delete the TableConfigs
[**DeleteTable**](TableApi.md#DeleteTable) | **Delete** /tables/{tableName} | Deletes a table
[**DeleteTimeBoundary**](TableApi.md#DeleteTimeBoundary) | **Delete** /tables/{tableName}/timeBoundary | Delete hybrid table query time boundary
[**ForceCommit**](TableApi.md#ForceCommit) | **Post** /tables/{tableName}/forceCommit | Force commit the current consuming segments
[**GetConfig**](TableApi.md#GetConfig) | **Get** /tableConfigs/{tableName} | Get the TableConfigs for a given raw tableName
[**GetConsumingSegmentsInfo**](TableApi.md#GetConsumingSegmentsInfo) | **Get** /tables/{tableName}/consumingSegmentsInfo | Returns state of consuming segments
[**GetControllerJobs**](TableApi.md#GetControllerJobs) | **Get** /table/{tableName}/jobs | Get list of controller jobs for this table
[**GetExternalView**](TableApi.md#GetExternalView) | **Get** /tables/{tableName}/externalview | Get table external view
[**GetForceCommitJobStatus**](TableApi.md#GetForceCommitJobStatus) | **Get** /tables/forceCommitStatus/{jobId} | Get status for a submitted force commit operation
[**GetIdealState**](TableApi.md#GetIdealState) | **Get** /tables/{tableName}/idealstate | Get table ideal state
[**GetInstancePartitions**](TableApi.md#GetInstancePartitions) | **Get** /tables/{tableName}/instancePartitions | Get the instance partitions
[**GetLiveBrokers**](TableApi.md#GetLiveBrokers) | **Get** /tables/livebrokers | List tables to live brokers mappings
[**GetLiveBrokersForTable**](TableApi.md#GetLiveBrokersForTable) | **Get** /tables/{tableName}/livebrokers | List the brokers serving a table
[**GetPauseStatus**](TableApi.md#GetPauseStatus) | **Get** /tables/{tableName}/pauseStatus | Return pause status of a realtime table
[**GetTableAggregateMetadata**](TableApi.md#GetTableAggregateMetadata) | **Get** /tables/{tableName}/metadata | Get the aggregate metadata of all segments for a table
[**GetTableInstances**](TableApi.md#GetTableInstances) | **Get** /tables/{tableName}/instances | List table instances
[**GetTableSize**](TableApi.md#GetTableSize) | **Get** /tables/{tableName}/size | Read table sizes
[**GetTableState**](TableApi.md#GetTableState) | **Get** /tables/{tableName}/state | Get current table state
[**GetTableStats**](TableApi.md#GetTableStats) | **Get** /tables/{tableName}/stats | table stats
[**GetTableStatus**](TableApi.md#GetTableStatus) | **Get** /tables/{tableName}/status | table status
[**IngestFromFile**](TableApi.md#IngestFromFile) | **Post** /ingestFromFile | Ingest a file
[**IngestFromURI**](TableApi.md#IngestFromURI) | **Post** /ingestFromURI | Ingest from the given URI
[**ListConfigs**](TableApi.md#ListConfigs) | **Get** /tableConfigs | Lists all TableConfigs in cluster
[**ListTables**](TableApi.md#ListTables) | **Get** /tables | Lists all tables in cluster
[**PauseConsumption**](TableApi.md#PauseConsumption) | **Post** /tables/{tableName}/pauseConsumption | Pause consumption of a realtime table
[**Put**](TableApi.md#Put) | **Put** /tables/{tableName}/segmentConfigs | Update segments configuration
[**Rebalance**](TableApi.md#Rebalance) | **Post** /tables/{tableName}/rebalance | Rebalances a table (reassign instances and segments for a table)
[**RebalanceStatus**](TableApi.md#RebalanceStatus) | **Get** /rebalanceStatus/{jobId} | Gets detailed stats of a rebalance operation
[**RebuildBrokerResource**](TableApi.md#RebuildBrokerResource) | **Post** /tables/{tableName}/rebuildBrokerResourceFromHelixTags | Rebuild broker resource for table
[**RecommendConfig**](TableApi.md#RecommendConfig) | **Put** /tables/recommender | Recommend config
[**RemoveInstancePartitions**](TableApi.md#RemoveInstancePartitions) | **Delete** /tables/{tableName}/instancePartitions | Remove the instance partitions
[**ReplaceInstance**](TableApi.md#ReplaceInstance) | **Post** /tables/{tableName}/replaceInstance | Replace an instance in the instance partitions
[**ResumeConsumption**](TableApi.md#ResumeConsumption) | **Post** /tables/{tableName}/resumeConsumption | Resume consumption of a realtime table
[**SetInstancePartitions**](TableApi.md#SetInstancePartitions) | **Put** /tables/{tableName}/instancePartitions | Create/update the instance partitions
[**SetTimeBoundary**](TableApi.md#SetTimeBoundary) | **Post** /tables/{tableName}/timeBoundary | Set hybrid table query time boundary based on offline segments&#39; metadata
[**UpdateConfig**](TableApi.md#UpdateConfig) | **Put** /tableConfigs/{tableName} | Update the TableConfigs provided by the tableConfigsStr json
[**UpdateIndexingConfig**](TableApi.md#UpdateIndexingConfig) | **Put** /tables/{tableName}/indexingConfigs | Update table indexing configuration
[**UpdateTableConfig**](TableApi.md#UpdateTableConfig) | **Put** /tables/{tableName} | Updates table config for a table
[**UpdateTableMetadata**](TableApi.md#UpdateTableMetadata) | **Put** /tables/{tableName}/metadataConfigs | Update table metadata
[**ValidateConfig**](TableApi.md#ValidateConfig) | **Post** /tableConfigs/validate | Validate the TableConfigs
[**ValidateTableAndSchema**](TableApi.md#ValidateTableAndSchema) | **Post** /tables/validateTableAndSchema | Validate table config for a table along with specified schema



## AddConfig

> ConfigSuccessResponse AddConfig(ctx).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()

Add the TableConfigs using the tableConfigsStr json



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
    validationTypesToSkip := "validationTypesToSkip_example" // string | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) (optional)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.AddConfig(context.Background()).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.AddConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConfig`: ConfigSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TableApi.AddConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validationTypesToSkip** | **string** | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) | 
 **body** | **string** |  | 

### Return type

[**ConfigSuccessResponse**](ConfigSuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTable

> ConfigSuccessResponse AddTable(ctx).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()

Adds a table



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
    validationTypesToSkip := "validationTypesToSkip_example" // string | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) (optional)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.AddTable(context.Background()).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.AddTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTable`: ConfigSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TableApi.AddTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validationTypesToSkip** | **string** | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) | 
 **body** | **string** |  | 

### Return type

[**ConfigSuccessResponse**](ConfigSuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlterTableStateOrListTableConfig

> string AlterTableStateOrListTableConfig(ctx, tableName).State(state).Type_(type_).Execute()

Get/Enable/Disable/Drop a table



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
    tableName := "tableName_example" // string | Name of the table
    state := "state_example" // string | enable|disable|drop (optional)
    type_ := "type__example" // string | realtime|offline (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.AlterTableStateOrListTableConfig(context.Background(), tableName).State(state).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.AlterTableStateOrListTableConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlterTableStateOrListTableConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.AlterTableStateOrListTableConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlterTableStateOrListTableConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | enable|disable|drop | 
 **type_** | **string** | realtime|offline | 

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


## AssignInstances

> map[string]InstancePartitions AssignInstances(ctx, tableName).Type_(type_).DryRun(dryRun).Execute()

Assign server instances to a table

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
    tableName := "tableName_example" // string | Name of the table
    type_ := "type__example" // string | OFFLINE|CONSUMING|COMPLETED|tier name (optional)
    dryRun := true // bool | Whether to do dry-run (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.AssignInstances(context.Background(), tableName).Type_(type_).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.AssignInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignInstances`: map[string]InstancePartitions
    fmt.Fprintf(os.Stdout, "Response from `TableApi.AssignInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|CONSUMING|COMPLETED|tier name | 
 **dryRun** | **bool** | Whether to do dry-run | [default to false]

### Return type

[**map[string]InstancePartitions**](InstancePartitions.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckTableConfig

> map[string]interface{} CheckTableConfig(ctx).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()

Validate table config for a table



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
    validationTypesToSkip := "validationTypesToSkip_example" // string | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) (optional)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.CheckTableConfig(context.Background()).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.CheckTableConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckTableConfig`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TableApi.CheckTableConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckTableConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validationTypesToSkip** | **string** | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) | 
 **body** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfig

> SuccessResponse DeleteConfig(ctx, tableName).Execute()

Delete the TableConfigs



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
    tableName := "tableName_example" // string | TableConfigs name i.e. raw table name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.DeleteConfig(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.DeleteConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfig`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TableApi.DeleteConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | TableConfigs name i.e. raw table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTable

> SuccessResponse DeleteTable(ctx, tableName).Type_(type_).Retention(retention).Execute()

Deletes a table



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
    tableName := "tableName_example" // string | Name of the table to delete
    type_ := "type__example" // string | realtime|offline (optional)
    retention := "retention_example" // string | Retention period for the table segments (e.g. 12h, 3d); If not set, the retention period will default to the first config that's not null: the cluster setting, then '7d'. Using 0d or -1d will instantly delete segments without retention (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.DeleteTable(context.Background(), tableName).Type_(type_).Retention(retention).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.DeleteTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTable`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TableApi.DeleteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | realtime|offline | 
 **retention** | **string** | Retention period for the table segments (e.g. 12h, 3d); If not set, the retention period will default to the first config that&#39;s not null: the cluster setting, then &#39;7d&#39;. Using 0d or -1d will instantly delete segments without retention | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTimeBoundary

> SuccessResponse DeleteTimeBoundary(ctx, tableName).Execute()

Delete hybrid table query time boundary



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
    tableName := "tableName_example" // string | Name of the hybrid table (without type suffix)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.DeleteTimeBoundary(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.DeleteTimeBoundary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTimeBoundary`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TableApi.DeleteTimeBoundary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the hybrid table (without type suffix) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeBoundaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForceCommit

> map[string]string ForceCommit(ctx, tableName).Execute()

Force commit the current consuming segments



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
    tableName := "tableName_example" // string | Name of the table

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.ForceCommit(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.ForceCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ForceCommit`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.ForceCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiForceCommitRequest struct via the builder pattern


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


## GetConfig

> string GetConfig(ctx, tableName).Execute()

Get the TableConfigs for a given raw tableName



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
    tableName := "tableName_example" // string | Raw table name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.GetConfig(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.GetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Raw table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


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


## GetConsumingSegmentsInfo

> GetConsumingSegmentsInfo(ctx, tableName).Execute()

Returns state of consuming segments



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
    tableName := "myTable | myTable_REALTIME" // string | Realtime table name with or without type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.GetConsumingSegmentsInfo(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetConsumingSegmentsInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Realtime table name with or without type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumingSegmentsInfoRequest struct via the builder pattern


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


## GetControllerJobs

> map[string]map[string]string GetControllerJobs(ctx, tableName).Type_(type_).JobTypes(jobTypes).Execute()

Get list of controller jobs for this table



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
    tableName := "tableName_example" // string | Name of the table
    type_ := "type__example" // string | OFFLINE|REALTIME (optional)
    jobTypes := "jobTypes_example" // string | Comma separated list of job types (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.GetControllerJobs(context.Background(), tableName).Type_(type_).JobTypes(jobTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetControllerJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetControllerJobs`: map[string]map[string]string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.GetControllerJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetControllerJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME | 
 **jobTypes** | **string** | Comma separated list of job types | 

### Return type

[**map[string]map[string]string**](map.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalView

> TableView GetExternalView(ctx, tableName).TableType(tableType).Execute()

Get table external view



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
    tableName := "tableName_example" // string | Name of the table
    tableType := "tableType_example" // string | realtime|offline (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.GetExternalView(context.Background(), tableName).TableType(tableType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetExternalView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalView`: TableView
    fmt.Fprintf(os.Stdout, "Response from `TableApi.GetExternalView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tableType** | **string** | realtime|offline | 

### Return type

[**TableView**](TableView.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForceCommitJobStatus

> map[string]interface{} GetForceCommitJobStatus(ctx, jobId).Execute()

Get status for a submitted force commit operation



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
    jobId := "jobId_example" // string | Force commit job id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.GetForceCommitJobStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetForceCommitJobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetForceCommitJobStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TableApi.GetForceCommitJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Force commit job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForceCommitJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdealState

> TableView GetIdealState(ctx, tableName).TableType(tableType).Execute()

Get table ideal state



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
    tableName := "tableName_example" // string | Name of the table
    tableType := "tableType_example" // string | realtime|offline (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.GetIdealState(context.Background(), tableName).TableType(tableType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetIdealState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdealState`: TableView
    fmt.Fprintf(os.Stdout, "Response from `TableApi.GetIdealState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdealStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tableType** | **string** | realtime|offline | 

### Return type

[**TableView**](TableView.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstancePartitions

> map[string]InstancePartitions GetInstancePartitions(ctx, tableName).Type_(type_).Execute()

Get the instance partitions

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
    tableName := "tableName_example" // string | Name of the table
    type_ := "type__example" // string | OFFLINE|CONSUMING|COMPLETED|tier name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.GetInstancePartitions(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetInstancePartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstancePartitions`: map[string]InstancePartitions
    fmt.Fprintf(os.Stdout, "Response from `TableApi.GetInstancePartitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancePartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|CONSUMING|COMPLETED|tier name | 

### Return type

[**map[string]InstancePartitions**](InstancePartitions.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveBrokers

> GetLiveBrokers(ctx).Execute()

List tables to live brokers mappings



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
    r, err := apiClient.TableApi.GetLiveBrokers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetLiveBrokers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveBrokersRequest struct via the builder pattern


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


## GetLiveBrokersForTable

> GetLiveBrokersForTable(ctx, tableName).Execute()

List the brokers serving a table



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
    tableName := "tableName_example" // string | Table name (with or without type)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.GetLiveBrokersForTable(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetLiveBrokersForTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Table name (with or without type) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveBrokersForTableRequest struct via the builder pattern


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


## GetPauseStatus

> GetPauseStatus(ctx, tableName).Execute()

Return pause status of a realtime table



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
    tableName := "tableName_example" // string | Name of the table

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.GetPauseStatus(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetPauseStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPauseStatusRequest struct via the builder pattern


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


## GetTableAggregateMetadata

> string GetTableAggregateMetadata(ctx, tableName).Type_(type_).Columns(columns).Execute()

Get the aggregate metadata of all segments for a table



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
    tableName := "tableName_example" // string | Name of the table
    type_ := "type__example" // string | OFFLINE|REALTIME (optional)
    columns := []string{"Inner_example"} // []string | Columns name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.GetTableAggregateMetadata(context.Background(), tableName).Type_(type_).Columns(columns).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetTableAggregateMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableAggregateMetadata`: string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.GetTableAggregateMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableAggregateMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME | 
 **columns** | **[]string** | Columns name | 

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


## GetTableInstances

> GetTableInstances(ctx, tableName).Type_(type_).Execute()

List table instances



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
    tableName := "tableName_example" // string | Table name without type
    type_ := "broker" // string | Instance type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.GetTableInstances(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetTableInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Table name without type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Instance type | 

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


## GetTableSize

> GetTableSize(ctx, tableName).Detailed(detailed).Execute()

Read table sizes



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
    tableName := "myTable | myTable_OFFLINE" // string | Table name without type
    detailed := true // bool | Get detailed information (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.GetTableSize(context.Background(), tableName).Detailed(detailed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetTableSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Table name without type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detailed** | **bool** | Get detailed information | [default to true]

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


## GetTableState

> string GetTableState(ctx, tableName).Type_(type_).Execute()

Get current table state



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
    tableName := "tableName_example" // string | Name of the table to get its state
    type_ := "type__example" // string | realtime|offline

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.GetTableState(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetTableState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableState`: string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.GetTableState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table to get its state | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | realtime|offline | 

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


## GetTableStats

> string GetTableStats(ctx, tableName).Type_(type_).Execute()

table stats



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
    tableName := "tableName_example" // string | Name of the table
    type_ := "type__example" // string | realtime|offline (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.GetTableStats(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetTableStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableStats`: string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.GetTableStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | realtime|offline | 

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


## GetTableStatus

> string GetTableStatus(ctx, tableName).Type_(type_).Execute()

table status



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
    tableName := "tableName_example" // string | Name of the table
    type_ := "type__example" // string | realtime|offline (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.GetTableStatus(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.GetTableStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableStatus`: string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.GetTableStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | realtime|offline | 

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


## IngestFromFile

> IngestFromFile(ctx).TableNameWithType(tableNameWithType).BatchConfigMapStr(batchConfigMapStr).ContentDisposition(contentDisposition).Entity(entity).MediaType(mediaType).MessageBodyWorkers(messageBodyWorkers).Parent(parent).Providers(providers).BodyParts(bodyParts).Execute()

Ingest a file



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
    tableNameWithType := "tableNameWithType_example" // string | Name of the table to upload the file to
    batchConfigMapStr := "batchConfigMapStr_example" // string | Batch config Map as json string. Must pass inputFormat, and optionally record reader properties. e.g. {\"inputFormat\":\"json\"}
    contentDisposition := TODO // ContentDisposition |  (optional)
    entity := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
    mediaType := TODO // MediaType |  (optional)
    messageBodyWorkers := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
    parent := TODO // MultiPart |  (optional)
    providers := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
    bodyParts := []BodyPart{"TODO"} // []BodyPart |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.IngestFromFile(context.Background()).TableNameWithType(tableNameWithType).BatchConfigMapStr(batchConfigMapStr).ContentDisposition(contentDisposition).Entity(entity).MediaType(mediaType).MessageBodyWorkers(messageBodyWorkers).Parent(parent).Providers(providers).BodyParts(bodyParts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.IngestFromFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableNameWithType** | **string** | Name of the table to upload the file to | 
 **batchConfigMapStr** | **string** | Batch config Map as json string. Must pass inputFormat, and optionally record reader properties. e.g. {\&quot;inputFormat\&quot;:\&quot;json\&quot;} | 
 **contentDisposition** | [**ContentDisposition**](ContentDisposition.md) |  | 
 **entity** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **mediaType** | [**MediaType**](MediaType.md) |  | 
 **messageBodyWorkers** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **parent** | [**MultiPart**](MultiPart.md) |  | 
 **providers** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **bodyParts** | [**[]BodyPart**](BodyPart.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestFromURI

> IngestFromURI(ctx).TableNameWithType(tableNameWithType).BatchConfigMapStr(batchConfigMapStr).SourceURIStr(sourceURIStr).Execute()

Ingest from the given URI



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
    tableNameWithType := "tableNameWithType_example" // string | Name of the table to upload the file to
    batchConfigMapStr := "batchConfigMapStr_example" // string | Batch config Map as json string. Must pass inputFormat, and optionally input FS properties. e.g. {\"inputFormat\":\"json\"}
    sourceURIStr := "sourceURIStr_example" // string | URI of file to upload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.IngestFromURI(context.Background()).TableNameWithType(tableNameWithType).BatchConfigMapStr(batchConfigMapStr).SourceURIStr(sourceURIStr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.IngestFromURI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestFromURIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableNameWithType** | **string** | Name of the table to upload the file to | 
 **batchConfigMapStr** | **string** | Batch config Map as json string. Must pass inputFormat, and optionally input FS properties. e.g. {\&quot;inputFormat\&quot;:\&quot;json\&quot;} | 
 **sourceURIStr** | **string** | URI of file to upload | 

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


## ListConfigs

> string ListConfigs(ctx).Execute()

Lists all TableConfigs in cluster



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
    resp, r, err := apiClient.TableApi.ListConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.ListConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConfigs`: string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.ListConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConfigsRequest struct via the builder pattern


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


## ListTables

> string ListTables(ctx).Type_(type_).TaskType(taskType).SortType(sortType).SortAsc(sortAsc).Execute()

Lists all tables in cluster



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
    type_ := "type__example" // string | realtime|offline (optional)
    taskType := "taskType_example" // string | Task type (optional)
    sortType := "sortType_example" // string | name|creationTime|lastModifiedTime (optional)
    sortAsc := true // bool | true|false (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.ListTables(context.Background()).Type_(type_).TaskType(taskType).SortType(sortType).SortAsc(sortAsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.ListTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTables`: string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.ListTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | realtime|offline | 
 **taskType** | **string** | Task type | 
 **sortType** | **string** | name|creationTime|lastModifiedTime | 
 **sortAsc** | **bool** | true|false | [default to true]

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


## PauseConsumption

> PauseConsumption(ctx, tableName).Execute()

Pause consumption of a realtime table



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
    tableName := "tableName_example" // string | Name of the table

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.PauseConsumption(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.PauseConsumption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseConsumptionRequest struct via the builder pattern


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


## Put

> Put(ctx, tableName).Body(body).Execute()

Update segments configuration



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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.Put(context.Background(), tableName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.Put``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

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


## Rebalance

> RebalanceResult Rebalance(ctx, tableName).Type_(type_).DryRun(dryRun).ReassignInstances(reassignInstances).IncludeConsuming(includeConsuming).Bootstrap(bootstrap).Downtime(downtime).MinAvailableReplicas(minAvailableReplicas).BestEfforts(bestEfforts).ExternalViewCheckIntervalInMs(externalViewCheckIntervalInMs).ExternalViewStabilizationTimeoutInMs(externalViewStabilizationTimeoutInMs).Execute()

Rebalances a table (reassign instances and segments for a table)



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
    tableName := "tableName_example" // string | Name of the table to rebalance
    type_ := "type__example" // string | OFFLINE|REALTIME
    dryRun := true // bool | Whether to rebalance table in dry-run mode (optional) (default to false)
    reassignInstances := true // bool | Whether to reassign instances before reassigning segments (optional) (default to false)
    includeConsuming := true // bool | Whether to reassign CONSUMING segments for real-time table (optional) (default to false)
    bootstrap := true // bool | Whether to rebalance table in bootstrap mode (regardless of minimum segment movement, reassign all segments in a round-robin fashion as if adding new segments to an empty table) (optional) (default to false)
    downtime := true // bool | Whether to allow downtime for the rebalance (optional) (default to false)
    minAvailableReplicas := int32(56) // int32 | For no-downtime rebalance, minimum number of replicas to keep alive during rebalance, or maximum number of replicas allowed to be unavailable if value is negative (optional) (default to 1)
    bestEfforts := true // bool | Whether to use best-efforts to rebalance (not fail the rebalance when the no-downtime contract cannot be achieved) (optional) (default to false)
    externalViewCheckIntervalInMs := int64(789) // int64 | How often to check if external view converges with ideal states (optional) (default to 1000)
    externalViewStabilizationTimeoutInMs := int64(789) // int64 | How long to wait till external view converges with ideal states (optional) (default to 3600000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.Rebalance(context.Background(), tableName).Type_(type_).DryRun(dryRun).ReassignInstances(reassignInstances).IncludeConsuming(includeConsuming).Bootstrap(bootstrap).Downtime(downtime).MinAvailableReplicas(minAvailableReplicas).BestEfforts(bestEfforts).ExternalViewCheckIntervalInMs(externalViewCheckIntervalInMs).ExternalViewStabilizationTimeoutInMs(externalViewStabilizationTimeoutInMs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.Rebalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Rebalance`: RebalanceResult
    fmt.Fprintf(os.Stdout, "Response from `TableApi.Rebalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table to rebalance | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME | 
 **dryRun** | **bool** | Whether to rebalance table in dry-run mode | [default to false]
 **reassignInstances** | **bool** | Whether to reassign instances before reassigning segments | [default to false]
 **includeConsuming** | **bool** | Whether to reassign CONSUMING segments for real-time table | [default to false]
 **bootstrap** | **bool** | Whether to rebalance table in bootstrap mode (regardless of minimum segment movement, reassign all segments in a round-robin fashion as if adding new segments to an empty table) | [default to false]
 **downtime** | **bool** | Whether to allow downtime for the rebalance | [default to false]
 **minAvailableReplicas** | **int32** | For no-downtime rebalance, minimum number of replicas to keep alive during rebalance, or maximum number of replicas allowed to be unavailable if value is negative | [default to 1]
 **bestEfforts** | **bool** | Whether to use best-efforts to rebalance (not fail the rebalance when the no-downtime contract cannot be achieved) | [default to false]
 **externalViewCheckIntervalInMs** | **int64** | How often to check if external view converges with ideal states | [default to 1000]
 **externalViewStabilizationTimeoutInMs** | **int64** | How long to wait till external view converges with ideal states | [default to 3600000]

### Return type

[**RebalanceResult**](RebalanceResult.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebalanceStatus

> ServerRebalanceJobStatusResponse RebalanceStatus(ctx, jobId).Execute()

Gets detailed stats of a rebalance operation



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
    jobId := "jobId_example" // string | Rebalance Job Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.RebalanceStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.RebalanceStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebalanceStatus`: ServerRebalanceJobStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TableApi.RebalanceStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Rebalance Job Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebalanceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerRebalanceJobStatusResponse**](ServerRebalanceJobStatusResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    r, err := apiClient.TableApi.RebuildBrokerResource(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.RebuildBrokerResource``: %v\n", err)
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


## RecommendConfig

> string RecommendConfig(ctx).Body(body).Execute()

Recommend config



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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.RecommendConfig(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.RecommendConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecommendConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.RecommendConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecommendConfigRequest struct via the builder pattern


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


## RemoveInstancePartitions

> SuccessResponse RemoveInstancePartitions(ctx, tableName).Type_(type_).Execute()

Remove the instance partitions

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
    tableName := "tableName_example" // string | Name of the table
    type_ := "type__example" // string | OFFLINE|CONSUMING|COMPLETED|tier name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.RemoveInstancePartitions(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.RemoveInstancePartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveInstancePartitions`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TableApi.RemoveInstancePartitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInstancePartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|CONSUMING|COMPLETED|tier name | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceInstance

> map[string]InstancePartitions ReplaceInstance(ctx, tableName).OldInstanceId(oldInstanceId).NewInstanceId(newInstanceId).Type_(type_).Execute()

Replace an instance in the instance partitions

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
    tableName := "tableName_example" // string | Name of the table
    oldInstanceId := "oldInstanceId_example" // string | Old instance to be replaced
    newInstanceId := "newInstanceId_example" // string | New instance to replace with
    type_ := "type__example" // string | OFFLINE|CONSUMING|COMPLETED|tier name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.ReplaceInstance(context.Background(), tableName).OldInstanceId(oldInstanceId).NewInstanceId(newInstanceId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.ReplaceInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceInstance`: map[string]InstancePartitions
    fmt.Fprintf(os.Stdout, "Response from `TableApi.ReplaceInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oldInstanceId** | **string** | Old instance to be replaced | 
 **newInstanceId** | **string** | New instance to replace with | 
 **type_** | **string** | OFFLINE|CONSUMING|COMPLETED|tier name | 

### Return type

[**map[string]InstancePartitions**](InstancePartitions.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeConsumption

> ResumeConsumption(ctx, tableName).ConsumeFrom(consumeFrom).Execute()

Resume consumption of a realtime table



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
    tableName := "tableName_example" // string | Name of the table
    consumeFrom := "consumeFrom_example" // string | smallest | largest (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.ResumeConsumption(context.Background(), tableName).ConsumeFrom(consumeFrom).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.ResumeConsumption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeConsumptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumeFrom** | **string** | smallest | largest | 

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


## SetInstancePartitions

> map[string]InstancePartitions SetInstancePartitions(ctx, tableName).Body(body).Execute()

Create/update the instance partitions

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
    tableName := "tableName_example" // string | Name of the table
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.SetInstancePartitions(context.Background(), tableName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.SetInstancePartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetInstancePartitions`: map[string]InstancePartitions
    fmt.Fprintf(os.Stdout, "Response from `TableApi.SetInstancePartitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetInstancePartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**map[string]InstancePartitions**](InstancePartitions.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTimeBoundary

> SuccessResponse SetTimeBoundary(ctx, tableName).Execute()

Set hybrid table query time boundary based on offline segments' metadata



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
    tableName := "tableName_example" // string | Name of the hybrid table (without type suffix)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.SetTimeBoundary(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.SetTimeBoundary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetTimeBoundary`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TableApi.SetTimeBoundary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the hybrid table (without type suffix) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTimeBoundaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfig

> ConfigSuccessResponse UpdateConfig(ctx, tableName).ValidationTypesToSkip(validationTypesToSkip).Reload(reload).Body(body).Execute()

Update the TableConfigs provided by the tableConfigsStr json



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
    tableName := "tableName_example" // string | TableConfigs name i.e. raw table name
    validationTypesToSkip := "validationTypesToSkip_example" // string | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) (optional)
    reload := true // bool | Reload the table if the new schema is backward compatible (optional) (default to false)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.UpdateConfig(context.Background(), tableName).ValidationTypesToSkip(validationTypesToSkip).Reload(reload).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.UpdateConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfig`: ConfigSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TableApi.UpdateConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | TableConfigs name i.e. raw table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validationTypesToSkip** | **string** | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) | 
 **reload** | **bool** | Reload the table if the new schema is backward compatible | [default to false]
 **body** | **string** |  | 

### Return type

[**ConfigSuccessResponse**](ConfigSuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndexingConfig

> UpdateIndexingConfig(ctx, tableName).Body(body).Execute()

Update table indexing configuration

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
    tableName := "tableName_example" // string | Table name (without type)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.UpdateIndexingConfig(context.Background(), tableName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.UpdateIndexingConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Table name (without type) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndexingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

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


## UpdateTableConfig

> ConfigSuccessResponse UpdateTableConfig(ctx, tableName).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()

Updates table config for a table



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
    tableName := "tableName_example" // string | Name of the table to update
    validationTypesToSkip := "validationTypesToSkip_example" // string | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) (optional)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.UpdateTableConfig(context.Background(), tableName).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.UpdateTableConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTableConfig`: ConfigSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TableApi.UpdateTableConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTableConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validationTypesToSkip** | **string** | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) | 
 **body** | **string** |  | 

### Return type

[**ConfigSuccessResponse**](ConfigSuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTableMetadata

> UpdateTableMetadata(ctx, tableName).Body(body).Execute()

Update table metadata



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
    tableName := "tableName_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TableApi.UpdateTableMetadata(context.Background(), tableName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.UpdateTableMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTableMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

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


## ValidateConfig

> string ValidateConfig(ctx).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()

Validate the TableConfigs



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
    validationTypesToSkip := "validationTypesToSkip_example" // string | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) (optional)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.ValidateConfig(context.Background()).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.ValidateConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.ValidateConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validationTypesToSkip** | **string** | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) | 
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


## ValidateTableAndSchema

> string ValidateTableAndSchema(ctx).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()

Validate table config for a table along with specified schema



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
    validationTypesToSkip := "validationTypesToSkip_example" // string | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) (optional)
    body := *openapiclient.NewTableAndSchemaConfig(*openapiclient.NewTableConfig()) // TableAndSchemaConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TableApi.ValidateTableAndSchema(context.Background()).ValidationTypesToSkip(validationTypesToSkip).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableApi.ValidateTableAndSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateTableAndSchema`: string
    fmt.Fprintf(os.Stdout, "Response from `TableApi.ValidateTableAndSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateTableAndSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validationTypesToSkip** | **string** | comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT) | 
 **body** | [**TableAndSchemaConfig**](TableAndSchemaConfig.md) |  | 

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

