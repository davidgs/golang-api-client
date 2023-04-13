# SegmentConsumerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentName** | Pointer to **string** |  | [optional] [readonly] 
**ConsumerState** | Pointer to **string** |  | [optional] [readonly] 
**LastConsumedTimestamp** | Pointer to **int64** |  | [optional] [readonly] 
**PartitionToOffsetMap** | Pointer to **map[string]string** |  | [optional] [readonly] 
**PartitionOffsetInfo** | Pointer to [**PartitionOffsetInfo**](PartitionOffsetInfo.md) |  | [optional] 

## Methods

### NewSegmentConsumerInfo

`func NewSegmentConsumerInfo() *SegmentConsumerInfo`

NewSegmentConsumerInfo instantiates a new SegmentConsumerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentConsumerInfoWithDefaults

`func NewSegmentConsumerInfoWithDefaults() *SegmentConsumerInfo`

NewSegmentConsumerInfoWithDefaults instantiates a new SegmentConsumerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentName

`func (o *SegmentConsumerInfo) GetSegmentName() string`

GetSegmentName returns the SegmentName field if non-nil, zero value otherwise.

### GetSegmentNameOk

`func (o *SegmentConsumerInfo) GetSegmentNameOk() (*string, bool)`

GetSegmentNameOk returns a tuple with the SegmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentName

`func (o *SegmentConsumerInfo) SetSegmentName(v string)`

SetSegmentName sets SegmentName field to given value.

### HasSegmentName

`func (o *SegmentConsumerInfo) HasSegmentName() bool`

HasSegmentName returns a boolean if a field has been set.

### GetConsumerState

`func (o *SegmentConsumerInfo) GetConsumerState() string`

GetConsumerState returns the ConsumerState field if non-nil, zero value otherwise.

### GetConsumerStateOk

`func (o *SegmentConsumerInfo) GetConsumerStateOk() (*string, bool)`

GetConsumerStateOk returns a tuple with the ConsumerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerState

`func (o *SegmentConsumerInfo) SetConsumerState(v string)`

SetConsumerState sets ConsumerState field to given value.

### HasConsumerState

`func (o *SegmentConsumerInfo) HasConsumerState() bool`

HasConsumerState returns a boolean if a field has been set.

### GetLastConsumedTimestamp

`func (o *SegmentConsumerInfo) GetLastConsumedTimestamp() int64`

GetLastConsumedTimestamp returns the LastConsumedTimestamp field if non-nil, zero value otherwise.

### GetLastConsumedTimestampOk

`func (o *SegmentConsumerInfo) GetLastConsumedTimestampOk() (*int64, bool)`

GetLastConsumedTimestampOk returns a tuple with the LastConsumedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedTimestamp

`func (o *SegmentConsumerInfo) SetLastConsumedTimestamp(v int64)`

SetLastConsumedTimestamp sets LastConsumedTimestamp field to given value.

### HasLastConsumedTimestamp

`func (o *SegmentConsumerInfo) HasLastConsumedTimestamp() bool`

HasLastConsumedTimestamp returns a boolean if a field has been set.

### GetPartitionToOffsetMap

`func (o *SegmentConsumerInfo) GetPartitionToOffsetMap() map[string]string`

GetPartitionToOffsetMap returns the PartitionToOffsetMap field if non-nil, zero value otherwise.

### GetPartitionToOffsetMapOk

`func (o *SegmentConsumerInfo) GetPartitionToOffsetMapOk() (*map[string]string, bool)`

GetPartitionToOffsetMapOk returns a tuple with the PartitionToOffsetMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionToOffsetMap

`func (o *SegmentConsumerInfo) SetPartitionToOffsetMap(v map[string]string)`

SetPartitionToOffsetMap sets PartitionToOffsetMap field to given value.

### HasPartitionToOffsetMap

`func (o *SegmentConsumerInfo) HasPartitionToOffsetMap() bool`

HasPartitionToOffsetMap returns a boolean if a field has been set.

### GetPartitionOffsetInfo

`func (o *SegmentConsumerInfo) GetPartitionOffsetInfo() PartitionOffsetInfo`

GetPartitionOffsetInfo returns the PartitionOffsetInfo field if non-nil, zero value otherwise.

### GetPartitionOffsetInfoOk

`func (o *SegmentConsumerInfo) GetPartitionOffsetInfoOk() (*PartitionOffsetInfo, bool)`

GetPartitionOffsetInfoOk returns a tuple with the PartitionOffsetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionOffsetInfo

`func (o *SegmentConsumerInfo) SetPartitionOffsetInfo(v PartitionOffsetInfo)`

SetPartitionOffsetInfo sets PartitionOffsetInfo field to given value.

### HasPartitionOffsetInfo

`func (o *SegmentConsumerInfo) HasPartitionOffsetInfo() bool`

HasPartitionOffsetInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


