# ClusterHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnhealthyServerCount** | Pointer to **int64** |  | [optional] 
**TableToMisconfiguredSegmentCount** | Pointer to **map[string]int32** |  | [optional] 
**TableToErrorSegmentsCount** | Pointer to **map[string]int32** |  | [optional] 
**TableToSegmentsWitHMissingColumnsCount** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewClusterHealthResponse

`func NewClusterHealthResponse() *ClusterHealthResponse`

NewClusterHealthResponse instantiates a new ClusterHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterHealthResponseWithDefaults

`func NewClusterHealthResponseWithDefaults() *ClusterHealthResponse`

NewClusterHealthResponseWithDefaults instantiates a new ClusterHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnhealthyServerCount

`func (o *ClusterHealthResponse) GetUnhealthyServerCount() int64`

GetUnhealthyServerCount returns the UnhealthyServerCount field if non-nil, zero value otherwise.

### GetUnhealthyServerCountOk

`func (o *ClusterHealthResponse) GetUnhealthyServerCountOk() (*int64, bool)`

GetUnhealthyServerCountOk returns a tuple with the UnhealthyServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyServerCount

`func (o *ClusterHealthResponse) SetUnhealthyServerCount(v int64)`

SetUnhealthyServerCount sets UnhealthyServerCount field to given value.

### HasUnhealthyServerCount

`func (o *ClusterHealthResponse) HasUnhealthyServerCount() bool`

HasUnhealthyServerCount returns a boolean if a field has been set.

### GetTableToMisconfiguredSegmentCount

`func (o *ClusterHealthResponse) GetTableToMisconfiguredSegmentCount() map[string]int32`

GetTableToMisconfiguredSegmentCount returns the TableToMisconfiguredSegmentCount field if non-nil, zero value otherwise.

### GetTableToMisconfiguredSegmentCountOk

`func (o *ClusterHealthResponse) GetTableToMisconfiguredSegmentCountOk() (*map[string]int32, bool)`

GetTableToMisconfiguredSegmentCountOk returns a tuple with the TableToMisconfiguredSegmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableToMisconfiguredSegmentCount

`func (o *ClusterHealthResponse) SetTableToMisconfiguredSegmentCount(v map[string]int32)`

SetTableToMisconfiguredSegmentCount sets TableToMisconfiguredSegmentCount field to given value.

### HasTableToMisconfiguredSegmentCount

`func (o *ClusterHealthResponse) HasTableToMisconfiguredSegmentCount() bool`

HasTableToMisconfiguredSegmentCount returns a boolean if a field has been set.

### GetTableToErrorSegmentsCount

`func (o *ClusterHealthResponse) GetTableToErrorSegmentsCount() map[string]int32`

GetTableToErrorSegmentsCount returns the TableToErrorSegmentsCount field if non-nil, zero value otherwise.

### GetTableToErrorSegmentsCountOk

`func (o *ClusterHealthResponse) GetTableToErrorSegmentsCountOk() (*map[string]int32, bool)`

GetTableToErrorSegmentsCountOk returns a tuple with the TableToErrorSegmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableToErrorSegmentsCount

`func (o *ClusterHealthResponse) SetTableToErrorSegmentsCount(v map[string]int32)`

SetTableToErrorSegmentsCount sets TableToErrorSegmentsCount field to given value.

### HasTableToErrorSegmentsCount

`func (o *ClusterHealthResponse) HasTableToErrorSegmentsCount() bool`

HasTableToErrorSegmentsCount returns a boolean if a field has been set.

### GetTableToSegmentsWitHMissingColumnsCount

`func (o *ClusterHealthResponse) GetTableToSegmentsWitHMissingColumnsCount() map[string]int32`

GetTableToSegmentsWitHMissingColumnsCount returns the TableToSegmentsWitHMissingColumnsCount field if non-nil, zero value otherwise.

### GetTableToSegmentsWitHMissingColumnsCountOk

`func (o *ClusterHealthResponse) GetTableToSegmentsWitHMissingColumnsCountOk() (*map[string]int32, bool)`

GetTableToSegmentsWitHMissingColumnsCountOk returns a tuple with the TableToSegmentsWitHMissingColumnsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableToSegmentsWitHMissingColumnsCount

`func (o *ClusterHealthResponse) SetTableToSegmentsWitHMissingColumnsCount(v map[string]int32)`

SetTableToSegmentsWitHMissingColumnsCount sets TableToSegmentsWitHMissingColumnsCount field to given value.

### HasTableToSegmentsWitHMissingColumnsCount

`func (o *ClusterHealthResponse) HasTableToSegmentsWitHMissingColumnsCount() bool`

HasTableToSegmentsWitHMissingColumnsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


