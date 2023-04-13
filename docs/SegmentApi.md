# \SegmentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllSegments**](SegmentApi.md#DeleteAllSegments) | **Delete** /segments/{tableName} | Delete all segments
[**DeleteSegment**](SegmentApi.md#DeleteSegment) | **Delete** /segments/{tableName}/{segmentName} | Delete a segment
[**DeleteSegments**](SegmentApi.md#DeleteSegments) | **Post** /segments/{tableName}/delete | Delete the segments in the JSON array payload
[**DownloadSegment**](SegmentApi.md#DownloadSegment) | **Get** /segments/{tableName}/{segmentName} | Download a segment
[**EndReplaceSegments**](SegmentApi.md#EndReplaceSegments) | **Post** /segments/{tableName}/endReplaceSegments | End to replace segments
[**GetReloadJobStatus**](SegmentApi.md#GetReloadJobStatus) | **Get** /segments/segmentReloadStatus/{jobId} | Get status for a submitted reload operation
[**GetSegmentMetadata**](SegmentApi.md#GetSegmentMetadata) | **Get** /segments/{tableName}/{segmentName}/metadata | Get the metadata for a segment
[**GetSegmentMetadataDeprecated1**](SegmentApi.md#GetSegmentMetadataDeprecated1) | **Get** /tables/{tableName}/segments/{segmentName}/metadata | Get the metadata for a segment (deprecated, use &#39;GET /segments/{tableName}/{segmentName}/metadata&#39; instead)
[**GetSegmentMetadataDeprecated2**](SegmentApi.md#GetSegmentMetadataDeprecated2) | **Get** /tables/{tableName}/segments/{segmentName} | Get the metadata for a segment (deprecated, use &#39;GET /segments/{tableName}/{segmentName}/metadata&#39; instead)
[**GetSegmentTiers**](SegmentApi.md#GetSegmentTiers) | **Get** /segments/{tableName}/{segmentName}/tiers | Get storage tiers for the given segment
[**GetSegmentToCrcMap**](SegmentApi.md#GetSegmentToCrcMap) | **Get** /segments/{tableName}/crc | Get a map from segment to CRC of the segment (only apply to OFFLINE table)
[**GetSegmentToCrcMapDeprecated**](SegmentApi.md#GetSegmentToCrcMapDeprecated) | **Get** /tables/{tableName}/segments/crc | Get a map from segment to CRC of the segment (deprecated, use &#39;GET /segments/{tableName}/crc&#39; instead)
[**GetSegments**](SegmentApi.md#GetSegments) | **Get** /segments/{tableName} | List all segments. An optional &#39;excludeReplacedSegments&#39; parameter is used to get the list of segments which has not yet been replaced (determined by segment lineage entries) and can be queried from the table. The value is false by default.
[**GetSelectedSegments**](SegmentApi.md#GetSelectedSegments) | **Get** /segments/{tableName}/select | Get the selected segments given the (inclusive) start and (exclusive) end timestamps in milliseconds. These timestamps will be compared against the minmax values of the time column in each segment. If the table is a refresh use case, the value of start and end timestamp is voided, since there is no time column for refresh use case; instead, the whole qualified segments will be returned. If no timestamps are provided, all the qualified segments will be returned. For the segments that partially belong to the time range, the boolean flag &#39;excludeOverlapping&#39; is introduced in order for user to determine whether to exclude this kind of segments in the response.
[**GetServerMetadata**](SegmentApi.md#GetServerMetadata) | **Get** /segments/{tableName}/metadata | Get the server metadata for all table segments
[**GetServerToSegmentsMap**](SegmentApi.md#GetServerToSegmentsMap) | **Get** /segments/{tableName}/servers | Get a map from server to segments hosted by the server
[**GetServerToSegmentsMapDeprecated1**](SegmentApi.md#GetServerToSegmentsMapDeprecated1) | **Get** /tables/{tableName}/segments | Get a map from server to segments hosted by the server (deprecated, use &#39;GET /segments/{tableName}/servers&#39; instead)
[**GetServerToSegmentsMapDeprecated2**](SegmentApi.md#GetServerToSegmentsMapDeprecated2) | **Get** /tables/{tableName}/segments/metadata | Get a map from server to segments hosted by the server (deprecated, use &#39;GET /segments/{tableName}/servers&#39; instead)
[**GetTableTiers**](SegmentApi.md#GetTableTiers) | **Get** /segments/{tableName}/tiers | Get storage tier for all segments in the given table
[**GetZookeeperMetadata**](SegmentApi.md#GetZookeeperMetadata) | **Get** /segments/{tableName}/zkmetadata | Get the zookeeper metadata for all table segments
[**ListSegmentLineage**](SegmentApi.md#ListSegmentLineage) | **Get** /segments/{tableName}/lineage | List segment lineage
[**ReloadAllSegments**](SegmentApi.md#ReloadAllSegments) | **Post** /segments/{tableName}/reload | Reload all segments
[**ReloadAllSegmentsDeprecated1**](SegmentApi.md#ReloadAllSegmentsDeprecated1) | **Post** /tables/{tableName}/segments/reload | Reload all segments (deprecated, use &#39;POST /segments/{tableName}/reload&#39; instead)
[**ReloadAllSegmentsDeprecated2**](SegmentApi.md#ReloadAllSegmentsDeprecated2) | **Get** /tables/{tableName}/segments/reload | Reload all segments (deprecated, use &#39;POST /segments/{tableName}/reload&#39; instead)
[**ReloadSegment**](SegmentApi.md#ReloadSegment) | **Post** /segments/{tableName}/{segmentName}/reload | Reload a segment
[**ReloadSegmentDeprecated1**](SegmentApi.md#ReloadSegmentDeprecated1) | **Post** /tables/{tableName}/segments/{segmentName}/reload | Reload a segment (deprecated, use &#39;POST /segments/{tableName}/{segmentName}/reload&#39; instead)
[**ReloadSegmentDeprecated2**](SegmentApi.md#ReloadSegmentDeprecated2) | **Get** /tables/{tableName}/segments/{segmentName}/reload | Reload a segment (deprecated, use &#39;POST /segments/{tableName}/{segmentName}/reload&#39; instead)
[**ResetSegment**](SegmentApi.md#ResetSegment) | **Post** /segments/{tableNameWithType}/{segmentName}/reset | Resets a segment by first disabling it, waiting for external view to stabilize, and finally enabling it again
[**ResetSegments**](SegmentApi.md#ResetSegments) | **Post** /segments/{tableNameWithType}/reset | Resets all segments (when errorSegmentsOnly &#x3D; false) or segments with Error state (when errorSegmentsOnly &#x3D; true) of the table, by first disabling them, waiting for external view to stabilize, and finally enabling them
[**RevertReplaceSegments**](SegmentApi.md#RevertReplaceSegments) | **Post** /segments/{tableName}/revertReplaceSegments | Revert segments replacement
[**StartReplaceSegments**](SegmentApi.md#StartReplaceSegments) | **Post** /segments/{tableName}/startReplaceSegments | Start to replace segments
[**UpdateTimeIntervalZK**](SegmentApi.md#UpdateTimeIntervalZK) | **Post** /segments/{tableNameWithType}/updateZKTimeInterval | Update the start and end time of the segments based on latest schema
[**UploadSegmentAsMultiPart**](SegmentApi.md#UploadSegmentAsMultiPart) | **Post** /segments | Upload a segment
[**UploadSegmentAsMultiPartV2**](SegmentApi.md#UploadSegmentAsMultiPartV2) | **Post** /v2/segments | Upload a segment



## DeleteAllSegments

> SuccessResponse DeleteAllSegments(ctx, tableName).Type_(type_).Retention(retention).Execute()

Delete all segments



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
    type_ := "type__example" // string | OFFLINE|REALTIME
    retention := "retention_example" // string | Retention period for the table segments (e.g. 12h, 3d); If not set, the retention period will default to the first config that's not null: the table config, then to cluster setting, then '7d'. Using 0d or -1d will instantly delete segments without retention (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.DeleteAllSegments(context.Background(), tableName).Type_(type_).Retention(retention).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.DeleteAllSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllSegments`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.DeleteAllSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |
 **retention** | **string** | Retention period for the table segments (e.g. 12h, 3d); If not set, the retention period will default to the first config that&#39;s not null: the table config, then to cluster setting, then &#39;7d&#39;. Using 0d or -1d will instantly delete segments without retention |

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


## DeleteSegment

> SuccessResponse DeleteSegment(ctx, tableName, segmentName).Retention(retention).Execute()

Delete a segment



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
    segmentName := "segmentName_example" // string | Name of the segment
    retention := "retention_example" // string | Retention period for the table segments (e.g. 12h, 3d); If not set, the retention period will default to the first config that's not null: the table config, then to cluster setting, then '7d'. Using 0d or -1d will instantly delete segments without retention (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.DeleteSegment(context.Background(), tableName, segmentName).Retention(retention).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.DeleteSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSegment`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.DeleteSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |
**segmentName** | **string** | Name of the segment |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **retention** | **string** | Retention period for the table segments (e.g. 12h, 3d); If not set, the retention period will default to the first config that&#39;s not null: the table config, then to cluster setting, then &#39;7d&#39;. Using 0d or -1d will instantly delete segments without retention |

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


## DeleteSegments

> SuccessResponse DeleteSegments(ctx, tableName).Retention(retention).Body(body).Execute()

Delete the segments in the JSON array payload



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
    retention := "retention_example" // string | Retention period for the table segments (e.g. 12h, 3d); If not set, the retention period will default to the first config that's not null: the table config, then to cluster setting, then '7d'. Using 0d or -1d will instantly delete segments without retention (optional)
    body := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.DeleteSegments(context.Background(), tableName).Retention(retention).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.DeleteSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSegments`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.DeleteSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **retention** | **string** | Retention period for the table segments (e.g. 12h, 3d); If not set, the retention period will default to the first config that&#39;s not null: the table config, then to cluster setting, then &#39;7d&#39;. Using 0d or -1d will instantly delete segments without retention |
 **body** | **[]string** |  |

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadSegment

> DownloadSegment(ctx, tableName, segmentName).Execute()

Download a segment



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
    segmentName := "segmentName_example" // string | Name of the segment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentApi.DownloadSegment(context.Background(), tableName, segmentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.DownloadSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |
**segmentName** | **string** | Name of the segment |

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSegmentRequest struct via the builder pattern


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


## EndReplaceSegments

> EndReplaceSegments(ctx, tableName).Type_(type_).SegmentLineageEntryId(segmentLineageEntryId).Execute()

End to replace segments



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
    type_ := "type__example" // string | OFFLINE|REALTIME
    segmentLineageEntryId := "segmentLineageEntryId_example" // string | Segment lineage entry id returned by startReplaceSegments API

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentApi.EndReplaceSegments(context.Background(), tableName).Type_(type_).SegmentLineageEntryId(segmentLineageEntryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.EndReplaceSegments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEndReplaceSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |
 **segmentLineageEntryId** | **string** | Segment lineage entry id returned by startReplaceSegments API |

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


## GetReloadJobStatus

> ServerReloadControllerJobStatusResponse GetReloadJobStatus(ctx, jobId).Execute()

Get status for a submitted reload operation



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
    jobId := "jobId_example" // string | Reload job id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.GetReloadJobStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetReloadJobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReloadJobStatus`: ServerReloadControllerJobStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetReloadJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Reload job id |

### Other Parameters

Other parameters are passed through a pointer to a apiGetReloadJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerReloadControllerJobStatusResponse**](ServerReloadControllerJobStatusResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentMetadata

> map[string]map[string]interface{} GetSegmentMetadata(ctx, tableName, segmentName).Columns(columns).Execute()

Get the metadata for a segment



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
    segmentName := "segmentName_example" // string | Name of the segment
    columns := []string{"Inner_example"} // []string | Columns name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.GetSegmentMetadata(context.Background(), tableName, segmentName).Columns(columns).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetSegmentMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegmentMetadata`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetSegmentMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |
**segmentName** | **string** | Name of the segment |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **columns** | **[]string** | Columns name |

### Return type

**map[string]map[string]interface{}**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentMetadataDeprecated1

> [][]map[string]map[string]interface{} GetSegmentMetadataDeprecated1(ctx, tableName, segmentName).Type_(type_).Execute()

Get the metadata for a segment (deprecated, use 'GET /segments/{tableName}/{segmentName}/metadata' instead)



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
    segmentName := "segmentName_example" // string | Name of the segment
    type_ := "type__example" // string | OFFLINE|REALTIME (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.GetSegmentMetadataDeprecated1(context.Background(), tableName, segmentName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetSegmentMetadataDeprecated1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegmentMetadataDeprecated1`: [][]map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetSegmentMetadataDeprecated1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |
**segmentName** | **string** | Name of the segment |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentMetadataDeprecated1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** | OFFLINE|REALTIME |

### Return type

[**[][]map[string]map[string]interface{}**](array.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentMetadataDeprecated2

> [][]map[string]map[string]interface{} GetSegmentMetadataDeprecated2(ctx, tableName, segmentName).State(state).Type_(type_).Execute()

Get the metadata for a segment (deprecated, use 'GET /segments/{tableName}/{segmentName}/metadata' instead)



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
    segmentName := "segmentName_example" // string | Name of the segment
    state := "state_example" // string | MUST be null (optional)
    type_ := "type__example" // string | OFFLINE|REALTIME (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.GetSegmentMetadataDeprecated2(context.Background(), tableName, segmentName).State(state).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetSegmentMetadataDeprecated2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegmentMetadataDeprecated2`: [][]map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetSegmentMetadataDeprecated2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |
**segmentName** | **string** | Name of the segment |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentMetadataDeprecated2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **state** | **string** | MUST be null |
 **type_** | **string** | OFFLINE|REALTIME |

### Return type

[**[][]map[string]map[string]interface{}**](array.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentTiers

> GetSegmentTiers(ctx, tableName, segmentName).Type_(type_).Execute()

Get storage tiers for the given segment



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
    segmentName := "segmentName_example" // string | Name of the segment
    type_ := "type__example" // string | OFFLINE|REALTIME

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentApi.GetSegmentTiers(context.Background(), tableName, segmentName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetSegmentTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |
**segmentName** | **string** | Name of the segment |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** | OFFLINE|REALTIME |

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


## GetSegmentToCrcMap

> map[string]string GetSegmentToCrcMap(ctx, tableName).Execute()

Get a map from segment to CRC of the segment (only apply to OFFLINE table)



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
    resp, r, err := apiClient.SegmentApi.GetSegmentToCrcMap(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetSegmentToCrcMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegmentToCrcMap`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetSegmentToCrcMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentToCrcMapRequest struct via the builder pattern


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


## GetSegmentToCrcMapDeprecated

> map[string]string GetSegmentToCrcMapDeprecated(ctx, tableName).Execute()

Get a map from segment to CRC of the segment (deprecated, use 'GET /segments/{tableName}/crc' instead)



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
    resp, r, err := apiClient.SegmentApi.GetSegmentToCrcMapDeprecated(context.Background(), tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetSegmentToCrcMapDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegmentToCrcMapDeprecated`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetSegmentToCrcMapDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentToCrcMapDeprecatedRequest struct via the builder pattern


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


## GetSegments

> []map[string][]string GetSegments(ctx, tableName).Type_(type_).ExcludeReplacedSegments(excludeReplacedSegments).Execute()

List all segments. An optional 'excludeReplacedSegments' parameter is used to get the list of segments which has not yet been replaced (determined by segment lineage entries) and can be queried from the table. The value is false by default.



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
    excludeReplacedSegments := "excludeReplacedSegments_example" // string | Whether to exclude replaced segments in the response, which have been replaced specified in the segment lineage entries and cannot be queried from the table (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.GetSegments(context.Background(), tableName).Type_(type_).ExcludeReplacedSegments(excludeReplacedSegments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegments`: []map[string][]string
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |
 **excludeReplacedSegments** | **string** | Whether to exclude replaced segments in the response, which have been replaced specified in the segment lineage entries and cannot be queried from the table |

### Return type

[**[]map[string][]string**](map.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectedSegments

> []map[string][]string GetSelectedSegments(ctx, tableName).Type_(type_).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).ExcludeOverlapping(excludeOverlapping).Execute()

Get the selected segments given the (inclusive) start and (exclusive) end timestamps in milliseconds. These timestamps will be compared against the minmax values of the time column in each segment. If the table is a refresh use case, the value of start and end timestamp is voided, since there is no time column for refresh use case; instead, the whole qualified segments will be returned. If no timestamps are provided, all the qualified segments will be returned. For the segments that partially belong to the time range, the boolean flag 'excludeOverlapping' is introduced in order for user to determine whether to exclude this kind of segments in the response.



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
    startTimestamp := "startTimestamp_example" // string | Start timestamp (inclusive) (optional)
    endTimestamp := "endTimestamp_example" // string | End timestamp (exclusive) (optional)
    excludeOverlapping := true // bool | Whether to exclude the segments overlapping with the timestamps, false by default (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.GetSelectedSegments(context.Background(), tableName).Type_(type_).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).ExcludeOverlapping(excludeOverlapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetSelectedSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelectedSegments`: []map[string][]string
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetSelectedSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectedSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |
 **startTimestamp** | **string** | Start timestamp (inclusive) |
 **endTimestamp** | **string** | End timestamp (exclusive) |
 **excludeOverlapping** | **bool** | Whether to exclude the segments overlapping with the timestamps, false by default | [default to false]

### Return type

[**[]map[string][]string**](map.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerMetadata

> string GetServerMetadata(ctx, tableName).Type_(type_).Columns(columns).Execute()

Get the server metadata for all table segments



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
    columns := []string{"Inner_example"} // []string | Columns name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.GetServerMetadata(context.Background(), tableName).Type_(type_).Columns(columns).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetServerMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerMetadata`: string
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetServerMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerMetadataRequest struct via the builder pattern


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


## GetServerToSegmentsMap

> []map[string]map[string]interface{} GetServerToSegmentsMap(ctx, tableName).Type_(type_).Execute()

Get a map from server to segments hosted by the server



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.GetServerToSegmentsMap(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetServerToSegmentsMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerToSegmentsMap`: []map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetServerToSegmentsMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerToSegmentsMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |

### Return type

[**[]map[string]map[string]interface{}**](map.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerToSegmentsMapDeprecated1

> []map[string]string GetServerToSegmentsMapDeprecated1(ctx, tableName).State(state).Type_(type_).Execute()

Get a map from server to segments hosted by the server (deprecated, use 'GET /segments/{tableName}/servers' instead)



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
    state := "state_example" // string | MUST be null (optional)
    type_ := "type__example" // string | OFFLINE|REALTIME (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.GetServerToSegmentsMapDeprecated1(context.Background(), tableName).State(state).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetServerToSegmentsMapDeprecated1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerToSegmentsMapDeprecated1`: []map[string]string
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetServerToSegmentsMapDeprecated1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerToSegmentsMapDeprecated1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | MUST be null |
 **type_** | **string** | OFFLINE|REALTIME |

### Return type

[**[]map[string]string**](map.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerToSegmentsMapDeprecated2

> []map[string]string GetServerToSegmentsMapDeprecated2(ctx, tableName).State(state).Type_(type_).Execute()

Get a map from server to segments hosted by the server (deprecated, use 'GET /segments/{tableName}/servers' instead)



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
    state := "state_example" // string | MUST be null (optional)
    type_ := "type__example" // string | OFFLINE|REALTIME (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.GetServerToSegmentsMapDeprecated2(context.Background(), tableName).State(state).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetServerToSegmentsMapDeprecated2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerToSegmentsMapDeprecated2`: []map[string]string
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetServerToSegmentsMapDeprecated2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerToSegmentsMapDeprecated2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | MUST be null |
 **type_** | **string** | OFFLINE|REALTIME |

### Return type

[**[]map[string]string**](map.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTableTiers

> GetTableTiers(ctx, tableName).Type_(type_).Execute()

Get storage tier for all segments in the given table



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
    type_ := "type__example" // string | OFFLINE|REALTIME

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentApi.GetTableTiers(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetTableTiers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetTableTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |

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


## GetZookeeperMetadata

> map[string]map[string]string GetZookeeperMetadata(ctx, tableName).Type_(type_).Execute()

Get the zookeeper metadata for all table segments



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.GetZookeeperMetadata(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.GetZookeeperMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZookeeperMetadata`: map[string]map[string]string
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.GetZookeeperMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiGetZookeeperMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |

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


## ListSegmentLineage

> ListSegmentLineage(ctx, tableName).Type_(type_).Execute()

List segment lineage



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
    type_ := "type__example" // string | OFFLINE|REALTIME

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentApi.ListSegmentLineage(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.ListSegmentLineage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiListSegmentLineageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |

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


## ReloadAllSegments

> SuccessResponse ReloadAllSegments(ctx, tableName).Type_(type_).ForceDownload(forceDownload).Execute()

Reload all segments



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
    forceDownload := true // bool | Whether to force server to download segment (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.ReloadAllSegments(context.Background(), tableName).Type_(type_).ForceDownload(forceDownload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.ReloadAllSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReloadAllSegments`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.ReloadAllSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiReloadAllSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |
 **forceDownload** | **bool** | Whether to force server to download segment | [default to false]

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


## ReloadAllSegmentsDeprecated1

> SuccessResponse ReloadAllSegmentsDeprecated1(ctx, tableName).Type_(type_).Execute()

Reload all segments (deprecated, use 'POST /segments/{tableName}/reload' instead)



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.ReloadAllSegmentsDeprecated1(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.ReloadAllSegmentsDeprecated1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReloadAllSegmentsDeprecated1`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.ReloadAllSegmentsDeprecated1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiReloadAllSegmentsDeprecated1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |

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


## ReloadAllSegmentsDeprecated2

> SuccessResponse ReloadAllSegmentsDeprecated2(ctx, tableName).Type_(type_).Execute()

Reload all segments (deprecated, use 'POST /segments/{tableName}/reload' instead)



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.ReloadAllSegmentsDeprecated2(context.Background(), tableName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.ReloadAllSegmentsDeprecated2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReloadAllSegmentsDeprecated2`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.ReloadAllSegmentsDeprecated2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |

### Other Parameters

Other parameters are passed through a pointer to a apiReloadAllSegmentsDeprecated2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |

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


## ReloadSegment

> SuccessResponse ReloadSegment(ctx, tableName, segmentName).ForceDownload(forceDownload).Execute()

Reload a segment



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
    segmentName := "segmentName_example" // string | Name of the segment
    forceDownload := true // bool | Whether to force server to download segment (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.ReloadSegment(context.Background(), tableName, segmentName).ForceDownload(forceDownload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.ReloadSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReloadSegment`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.ReloadSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |
**segmentName** | **string** | Name of the segment |

### Other Parameters

Other parameters are passed through a pointer to a apiReloadSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceDownload** | **bool** | Whether to force server to download segment | [default to false]

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


## ReloadSegmentDeprecated1

> SuccessResponse ReloadSegmentDeprecated1(ctx, tableName, segmentName).Type_(type_).Execute()

Reload a segment (deprecated, use 'POST /segments/{tableName}/{segmentName}/reload' instead)



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
    segmentName := "segmentName_example" // string | Name of the segment
    type_ := "type__example" // string | OFFLINE|REALTIME (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.ReloadSegmentDeprecated1(context.Background(), tableName, segmentName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.ReloadSegmentDeprecated1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReloadSegmentDeprecated1`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.ReloadSegmentDeprecated1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |
**segmentName** | **string** | Name of the segment |

### Other Parameters

Other parameters are passed through a pointer to a apiReloadSegmentDeprecated1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** | OFFLINE|REALTIME |

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


## ReloadSegmentDeprecated2

> SuccessResponse ReloadSegmentDeprecated2(ctx, tableName, segmentName).Type_(type_).Execute()

Reload a segment (deprecated, use 'POST /segments/{tableName}/{segmentName}/reload' instead)



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
    segmentName := "segmentName_example" // string | Name of the segment
    type_ := "type__example" // string | OFFLINE|REALTIME (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.ReloadSegmentDeprecated2(context.Background(), tableName, segmentName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.ReloadSegmentDeprecated2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReloadSegmentDeprecated2`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.ReloadSegmentDeprecated2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableName** | **string** | Name of the table |
**segmentName** | **string** | Name of the segment |

### Other Parameters

Other parameters are passed through a pointer to a apiReloadSegmentDeprecated2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** | OFFLINE|REALTIME |

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


## ResetSegment

> SuccessResponse ResetSegment(ctx, tableNameWithType, segmentName).TargetInstance(targetInstance).Execute()

Resets a segment by first disabling it, waiting for external view to stabilize, and finally enabling it again



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
    tableNameWithType := "tableNameWithType_example" // string | Name of the table with type
    segmentName := "segmentName_example" // string | Name of the segment
    targetInstance := "targetInstance_example" // string | Name of the target instance to reset (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.ResetSegment(context.Background(), tableNameWithType, segmentName).TargetInstance(targetInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.ResetSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetSegment`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.ResetSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableNameWithType** | **string** | Name of the table with type |
**segmentName** | **string** | Name of the segment |

### Other Parameters

Other parameters are passed through a pointer to a apiResetSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **targetInstance** | **string** | Name of the target instance to reset |

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


## ResetSegments

> SuccessResponse ResetSegments(ctx, tableNameWithType).TargetInstance(targetInstance).ErrorSegmentsOnly(errorSegmentsOnly).Execute()

Resets all segments (when errorSegmentsOnly = false) or segments with Error state (when errorSegmentsOnly = true) of the table, by first disabling them, waiting for external view to stabilize, and finally enabling them



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
    tableNameWithType := "tableNameWithType_example" // string | Name of the table with type
    targetInstance := "targetInstance_example" // string | Name of the target instance to reset (optional)
    errorSegmentsOnly := true // bool | Whether to reset only segments with error state (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentApi.ResetSegments(context.Background(), tableNameWithType).TargetInstance(targetInstance).ErrorSegmentsOnly(errorSegmentsOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.ResetSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetSegments`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentApi.ResetSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableNameWithType** | **string** | Name of the table with type |

### Other Parameters

Other parameters are passed through a pointer to a apiResetSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetInstance** | **string** | Name of the target instance to reset |
 **errorSegmentsOnly** | **bool** | Whether to reset only segments with error state | [default to false]

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


## RevertReplaceSegments

> RevertReplaceSegments(ctx, tableName).Type_(type_).SegmentLineageEntryId(segmentLineageEntryId).ForceRevert(forceRevert).Execute()

Revert segments replacement



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
    type_ := "type__example" // string | OFFLINE|REALTIME
    segmentLineageEntryId := "segmentLineageEntryId_example" // string | Segment lineage entry id to revert
    forceRevert := true // bool | Force revert in case the user knows that the lineage entry is interrupted (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentApi.RevertReplaceSegments(context.Background(), tableName).Type_(type_).SegmentLineageEntryId(segmentLineageEntryId).ForceRevert(forceRevert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.RevertReplaceSegments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRevertReplaceSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |
 **segmentLineageEntryId** | **string** | Segment lineage entry id to revert |
 **forceRevert** | **bool** | Force revert in case the user knows that the lineage entry is interrupted | [default to false]

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


## StartReplaceSegments

> StartReplaceSegments(ctx, tableName).Type_(type_).Body(body).ForceCleanup(forceCleanup).Execute()

Start to replace segments



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
    type_ := "type__example" // string | OFFLINE|REALTIME
    body := *openapiclient.NewStartReplaceSegmentsRequest() // StartReplaceSegmentsRequest | Fields belonging to start replace segment request
    forceCleanup := true // bool | Force cleanup (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentApi.StartReplaceSegments(context.Background(), tableName).Type_(type_).Body(body).ForceCleanup(forceCleanup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.StartReplaceSegments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStartReplaceSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | OFFLINE|REALTIME |
 **body** | [**StartReplaceSegmentsRequest**](StartReplaceSegmentsRequest.md) | Fields belonging to start replace segment request |
 **forceCleanup** | **bool** | Force cleanup | [default to false]

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


## UpdateTimeIntervalZK

> UpdateTimeIntervalZK(ctx, tableNameWithType).Execute()

Update the start and end time of the segments based on latest schema



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
    tableNameWithType := "myTable_REALTIME" // string | Table name with type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentApi.UpdateTimeIntervalZK(context.Background(), tableNameWithType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.UpdateTimeIntervalZK``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableNameWithType** | **string** | Table name with type |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTimeIntervalZKRequest struct via the builder pattern


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


## UploadSegmentAsMultiPart

> UploadSegmentAsMultiPart(ctx).TableName(tableName).TableType(tableType).EnableParallelPushProtection(enableParallelPushProtection).AllowRefresh(allowRefresh).ContentDisposition(contentDisposition).Entity(entity).MediaType(mediaType).MessageBodyWorkers(messageBodyWorkers).Parent(parent).Providers(providers).BodyParts(bodyParts).Execute()

Upload a segment



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
    tableName := "tableName_example" // string | Name of the table (optional)
    tableType := "tableType_example" // string | Type of the table (optional) (default to "OFFLINE")
    enableParallelPushProtection := true // bool | Whether to enable parallel push protection (optional) (default to false)
    allowRefresh := true // bool | Whether to refresh if the segment already exists (optional) (default to true)
    contentDisposition := TODO // ContentDisposition |  (optional)
    entity := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
    mediaType := TODO // MediaType |  (optional)
    messageBodyWorkers := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
    parent := TODO // MultiPart |  (optional)
    providers := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
    bodyParts := []BodyPart{"TODO"} // []BodyPart |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentApi.UploadSegmentAsMultiPart(context.Background()).TableName(tableName).TableType(tableType).EnableParallelPushProtection(enableParallelPushProtection).AllowRefresh(allowRefresh).ContentDisposition(contentDisposition).Entity(entity).MediaType(mediaType).MessageBodyWorkers(messageBodyWorkers).Parent(parent).Providers(providers).BodyParts(bodyParts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.UploadSegmentAsMultiPart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadSegmentAsMultiPartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableName** | **string** | Name of the table |
 **tableType** | **string** | Type of the table | [default to &quot;OFFLINE&quot;]
 **enableParallelPushProtection** | **bool** | Whether to enable parallel push protection | [default to false]
 **allowRefresh** | **bool** | Whether to refresh if the segment already exists | [default to true]
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


## UploadSegmentAsMultiPartV2

> UploadSegmentAsMultiPartV2(ctx).TableName(tableName).TableType(tableType).EnableParallelPushProtection(enableParallelPushProtection).AllowRefresh(allowRefresh).ContentDisposition(contentDisposition).Entity(entity).MediaType(mediaType).MessageBodyWorkers(messageBodyWorkers).Parent(parent).Providers(providers).BodyParts(bodyParts).Execute()

Upload a segment



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
    tableName := "tableName_example" // string | Name of the table (optional)
    tableType := "tableType_example" // string | Type of the table (optional) (default to "OFFLINE")
    enableParallelPushProtection := true // bool | Whether to enable parallel push protection (optional) (default to false)
    allowRefresh := true // bool | Whether to refresh if the segment already exists (optional) (default to true)
    contentDisposition := TODO // ContentDisposition |  (optional)
    entity := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
    mediaType := TODO // MediaType |  (optional)
    messageBodyWorkers := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
    parent := TODO // MultiPart |  (optional)
    providers := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
    bodyParts := []BodyPart{"TODO"} // []BodyPart |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentApi.UploadSegmentAsMultiPartV2(context.Background()).TableName(tableName).TableType(tableType).EnableParallelPushProtection(enableParallelPushProtection).AllowRefresh(allowRefresh).ContentDisposition(contentDisposition).Entity(entity).MediaType(mediaType).MessageBodyWorkers(messageBodyWorkers).Parent(parent).Providers(providers).BodyParts(bodyParts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentApi.UploadSegmentAsMultiPartV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadSegmentAsMultiPartV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableName** | **string** | Name of the table |
 **tableType** | **string** | Type of the table | [default to &quot;OFFLINE&quot;]
 **enableParallelPushProtection** | **bool** | Whether to enable parallel push protection | [default to false]
 **allowRefresh** | **bool** | Whether to refresh if the segment already exists | [default to true]
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

