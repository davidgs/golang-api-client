# \AtomicIngestionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndDataIngestRequest**](AtomicIngestionApi.md#EndDataIngestRequest) | **Post** /segments/{tableName}/endDataIngestRequest | Mark the end of data ingestion to upload multiple segments
[**StartDataIngestRequest**](AtomicIngestionApi.md#StartDataIngestRequest) | **Post** /segments/{tableName}/startDataIngestRequest | Mark the start of data ingestion to upload multiple segments



## EndDataIngestRequest

> EndDataIngestRequest(ctx, tableName).TableType(tableType).TaskType(taskType).CheckpointEntryKey(checkpointEntryKey).Execute()

Mark the end of data ingestion to upload multiple segments

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
    tableType := "tableType_example" // string | OFFLINE|REALTIME
    taskType := "taskType_example" // string | Task type
    checkpointEntryKey := "checkpointEntryKey_example" // string | Key of checkpoint entry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AtomicIngestionApi.EndDataIngestRequest(context.Background(), tableName).TableType(tableType).TaskType(taskType).CheckpointEntryKey(checkpointEntryKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtomicIngestionApi.EndDataIngestRequest``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEndDataIngestRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tableType** | **string** | OFFLINE|REALTIME |
 **taskType** | **string** | Task type |
 **checkpointEntryKey** | **string** | Key of checkpoint entry |

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


## StartDataIngestRequest

> StartDataIngestRequest(ctx, tableName).TableType(tableType).TaskType(taskType).Body(body).Execute()

Mark the start of data ingestion to upload multiple segments

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
    tableType := "tableType_example" // string | OFFLINE|REALTIME
    taskType := "taskType_example" // string | Task type
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AtomicIngestionApi.StartDataIngestRequest(context.Background(), tableName).TableType(tableType).TaskType(taskType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtomicIngestionApi.StartDataIngestRequest``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStartDataIngestRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tableType** | **string** | OFFLINE|REALTIME |
 **taskType** | **string** | Task type |
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

