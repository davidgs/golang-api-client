# ConsumingSegmentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | Pointer to **string** |  | [optional] 
**ConsumerState** | Pointer to **string** |  | [optional] 
**LastConsumedTimestamp** | Pointer to **int64** |  | [optional] 
**PartitionToOffsetMap** | Pointer to **map[string]string** |  | [optional] 
**PartitionOffsetInfo** | Pointer to [**PartitionOffsetInfo**](PartitionOffsetInfo.md) |  | [optional] 

## Methods

### NewConsumingSegmentInfo

`func NewConsumingSegmentInfo() *ConsumingSegmentInfo`

NewConsumingSegmentInfo instantiates a new ConsumingSegmentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumingSegmentInfoWithDefaults

`func NewConsumingSegmentInfoWithDefaults() *ConsumingSegmentInfo`

NewConsumingSegmentInfoWithDefaults instantiates a new ConsumingSegmentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *ConsumingSegmentInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ConsumingSegmentInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ConsumingSegmentInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ConsumingSegmentInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetConsumerState

`func (o *ConsumingSegmentInfo) GetConsumerState() string`

GetConsumerState returns the ConsumerState field if non-nil, zero value otherwise.

### GetConsumerStateOk

`func (o *ConsumingSegmentInfo) GetConsumerStateOk() (*string, bool)`

GetConsumerStateOk returns a tuple with the ConsumerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerState

`func (o *ConsumingSegmentInfo) SetConsumerState(v string)`

SetConsumerState sets ConsumerState field to given value.

### HasConsumerState

`func (o *ConsumingSegmentInfo) HasConsumerState() bool`

HasConsumerState returns a boolean if a field has been set.

### GetLastConsumedTimestamp

`func (o *ConsumingSegmentInfo) GetLastConsumedTimestamp() int64`

GetLastConsumedTimestamp returns the LastConsumedTimestamp field if non-nil, zero value otherwise.

### GetLastConsumedTimestampOk

`func (o *ConsumingSegmentInfo) GetLastConsumedTimestampOk() (*int64, bool)`

GetLastConsumedTimestampOk returns a tuple with the LastConsumedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedTimestamp

`func (o *ConsumingSegmentInfo) SetLastConsumedTimestamp(v int64)`

SetLastConsumedTimestamp sets LastConsumedTimestamp field to given value.

### HasLastConsumedTimestamp

`func (o *ConsumingSegmentInfo) HasLastConsumedTimestamp() bool`

HasLastConsumedTimestamp returns a boolean if a field has been set.

### GetPartitionToOffsetMap

`func (o *ConsumingSegmentInfo) GetPartitionToOffsetMap() map[string]string`

GetPartitionToOffsetMap returns the PartitionToOffsetMap field if non-nil, zero value otherwise.

### GetPartitionToOffsetMapOk

`func (o *ConsumingSegmentInfo) GetPartitionToOffsetMapOk() (*map[string]string, bool)`

GetPartitionToOffsetMapOk returns a tuple with the PartitionToOffsetMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionToOffsetMap

`func (o *ConsumingSegmentInfo) SetPartitionToOffsetMap(v map[string]string)`

SetPartitionToOffsetMap sets PartitionToOffsetMap field to given value.

### HasPartitionToOffsetMap

`func (o *ConsumingSegmentInfo) HasPartitionToOffsetMap() bool`

HasPartitionToOffsetMap returns a boolean if a field has been set.

### GetPartitionOffsetInfo

`func (o *ConsumingSegmentInfo) GetPartitionOffsetInfo() PartitionOffsetInfo`

GetPartitionOffsetInfo returns the PartitionOffsetInfo field if non-nil, zero value otherwise.

### GetPartitionOffsetInfoOk

`func (o *ConsumingSegmentInfo) GetPartitionOffsetInfoOk() (*PartitionOffsetInfo, bool)`

GetPartitionOffsetInfoOk returns a tuple with the PartitionOffsetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionOffsetInfo

`func (o *ConsumingSegmentInfo) SetPartitionOffsetInfo(v PartitionOffsetInfo)`

SetPartitionOffsetInfo sets PartitionOffsetInfo field to given value.

### HasPartitionOffsetInfo

`func (o *ConsumingSegmentInfo) HasPartitionOffsetInfo() bool`

HasPartitionOffsetInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


