# PartitionOffsetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentOffsetsMap** | Pointer to **map[string]string** |  | [optional] 
**LatestUpstreamOffsetMap** | Pointer to **map[string]string** |  | [optional] 
**RecordsLagMap** | Pointer to **map[string]string** |  | [optional] 
**AvailabilityLagMsMap** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPartitionOffsetInfo

`func NewPartitionOffsetInfo() *PartitionOffsetInfo`

NewPartitionOffsetInfo instantiates a new PartitionOffsetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionOffsetInfoWithDefaults

`func NewPartitionOffsetInfoWithDefaults() *PartitionOffsetInfo`

NewPartitionOffsetInfoWithDefaults instantiates a new PartitionOffsetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentOffsetsMap

`func (o *PartitionOffsetInfo) GetCurrentOffsetsMap() map[string]string`

GetCurrentOffsetsMap returns the CurrentOffsetsMap field if non-nil, zero value otherwise.

### GetCurrentOffsetsMapOk

`func (o *PartitionOffsetInfo) GetCurrentOffsetsMapOk() (*map[string]string, bool)`

GetCurrentOffsetsMapOk returns a tuple with the CurrentOffsetsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOffsetsMap

`func (o *PartitionOffsetInfo) SetCurrentOffsetsMap(v map[string]string)`

SetCurrentOffsetsMap sets CurrentOffsetsMap field to given value.

### HasCurrentOffsetsMap

`func (o *PartitionOffsetInfo) HasCurrentOffsetsMap() bool`

HasCurrentOffsetsMap returns a boolean if a field has been set.

### GetLatestUpstreamOffsetMap

`func (o *PartitionOffsetInfo) GetLatestUpstreamOffsetMap() map[string]string`

GetLatestUpstreamOffsetMap returns the LatestUpstreamOffsetMap field if non-nil, zero value otherwise.

### GetLatestUpstreamOffsetMapOk

`func (o *PartitionOffsetInfo) GetLatestUpstreamOffsetMapOk() (*map[string]string, bool)`

GetLatestUpstreamOffsetMapOk returns a tuple with the LatestUpstreamOffsetMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestUpstreamOffsetMap

`func (o *PartitionOffsetInfo) SetLatestUpstreamOffsetMap(v map[string]string)`

SetLatestUpstreamOffsetMap sets LatestUpstreamOffsetMap field to given value.

### HasLatestUpstreamOffsetMap

`func (o *PartitionOffsetInfo) HasLatestUpstreamOffsetMap() bool`

HasLatestUpstreamOffsetMap returns a boolean if a field has been set.

### GetRecordsLagMap

`func (o *PartitionOffsetInfo) GetRecordsLagMap() map[string]string`

GetRecordsLagMap returns the RecordsLagMap field if non-nil, zero value otherwise.

### GetRecordsLagMapOk

`func (o *PartitionOffsetInfo) GetRecordsLagMapOk() (*map[string]string, bool)`

GetRecordsLagMapOk returns a tuple with the RecordsLagMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsLagMap

`func (o *PartitionOffsetInfo) SetRecordsLagMap(v map[string]string)`

SetRecordsLagMap sets RecordsLagMap field to given value.

### HasRecordsLagMap

`func (o *PartitionOffsetInfo) HasRecordsLagMap() bool`

HasRecordsLagMap returns a boolean if a field has been set.

### GetAvailabilityLagMsMap

`func (o *PartitionOffsetInfo) GetAvailabilityLagMsMap() map[string]string`

GetAvailabilityLagMsMap returns the AvailabilityLagMsMap field if non-nil, zero value otherwise.

### GetAvailabilityLagMsMapOk

`func (o *PartitionOffsetInfo) GetAvailabilityLagMsMapOk() (*map[string]string, bool)`

GetAvailabilityLagMsMapOk returns a tuple with the AvailabilityLagMsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityLagMsMap

`func (o *PartitionOffsetInfo) SetAvailabilityLagMsMap(v map[string]string)`

SetAvailabilityLagMsMap sets AvailabilityLagMsMap field to given value.

### HasAvailabilityLagMsMap

`func (o *PartitionOffsetInfo) HasAvailabilityLagMsMap() bool`

HasAvailabilityLagMsMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


