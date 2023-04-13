# SegmentPartitionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnPartitionMap** | [**map[string]ColumnPartitionConfig**](ColumnPartitionConfig.md) |  | [readonly] 

## Methods

### NewSegmentPartitionConfig

`func NewSegmentPartitionConfig(columnPartitionMap map[string]ColumnPartitionConfig, ) *SegmentPartitionConfig`

NewSegmentPartitionConfig instantiates a new SegmentPartitionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentPartitionConfigWithDefaults

`func NewSegmentPartitionConfigWithDefaults() *SegmentPartitionConfig`

NewSegmentPartitionConfigWithDefaults instantiates a new SegmentPartitionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnPartitionMap

`func (o *SegmentPartitionConfig) GetColumnPartitionMap() map[string]ColumnPartitionConfig`

GetColumnPartitionMap returns the ColumnPartitionMap field if non-nil, zero value otherwise.

### GetColumnPartitionMapOk

`func (o *SegmentPartitionConfig) GetColumnPartitionMapOk() (*map[string]ColumnPartitionConfig, bool)`

GetColumnPartitionMapOk returns a tuple with the ColumnPartitionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnPartitionMap

`func (o *SegmentPartitionConfig) SetColumnPartitionMap(v map[string]ColumnPartitionConfig)`

SetColumnPartitionMap sets ColumnPartitionMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


