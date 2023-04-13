# RebalanceStateStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentsMissing** | Pointer to **int32** |  | [optional] 
**SegmentsToRebalance** | Pointer to **int32** |  | [optional] 
**PercentSegmentsToRebalance** | Pointer to **float64** |  | [optional] 
**ReplicasToRebalance** | Pointer to **int32** |  | [optional] 

## Methods

### NewRebalanceStateStats

`func NewRebalanceStateStats() *RebalanceStateStats`

NewRebalanceStateStats instantiates a new RebalanceStateStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebalanceStateStatsWithDefaults

`func NewRebalanceStateStatsWithDefaults() *RebalanceStateStats`

NewRebalanceStateStatsWithDefaults instantiates a new RebalanceStateStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentsMissing

`func (o *RebalanceStateStats) GetSegmentsMissing() int32`

GetSegmentsMissing returns the SegmentsMissing field if non-nil, zero value otherwise.

### GetSegmentsMissingOk

`func (o *RebalanceStateStats) GetSegmentsMissingOk() (*int32, bool)`

GetSegmentsMissingOk returns a tuple with the SegmentsMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentsMissing

`func (o *RebalanceStateStats) SetSegmentsMissing(v int32)`

SetSegmentsMissing sets SegmentsMissing field to given value.

### HasSegmentsMissing

`func (o *RebalanceStateStats) HasSegmentsMissing() bool`

HasSegmentsMissing returns a boolean if a field has been set.

### GetSegmentsToRebalance

`func (o *RebalanceStateStats) GetSegmentsToRebalance() int32`

GetSegmentsToRebalance returns the SegmentsToRebalance field if non-nil, zero value otherwise.

### GetSegmentsToRebalanceOk

`func (o *RebalanceStateStats) GetSegmentsToRebalanceOk() (*int32, bool)`

GetSegmentsToRebalanceOk returns a tuple with the SegmentsToRebalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentsToRebalance

`func (o *RebalanceStateStats) SetSegmentsToRebalance(v int32)`

SetSegmentsToRebalance sets SegmentsToRebalance field to given value.

### HasSegmentsToRebalance

`func (o *RebalanceStateStats) HasSegmentsToRebalance() bool`

HasSegmentsToRebalance returns a boolean if a field has been set.

### GetPercentSegmentsToRebalance

`func (o *RebalanceStateStats) GetPercentSegmentsToRebalance() float64`

GetPercentSegmentsToRebalance returns the PercentSegmentsToRebalance field if non-nil, zero value otherwise.

### GetPercentSegmentsToRebalanceOk

`func (o *RebalanceStateStats) GetPercentSegmentsToRebalanceOk() (*float64, bool)`

GetPercentSegmentsToRebalanceOk returns a tuple with the PercentSegmentsToRebalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentSegmentsToRebalance

`func (o *RebalanceStateStats) SetPercentSegmentsToRebalance(v float64)`

SetPercentSegmentsToRebalance sets PercentSegmentsToRebalance field to given value.

### HasPercentSegmentsToRebalance

`func (o *RebalanceStateStats) HasPercentSegmentsToRebalance() bool`

HasPercentSegmentsToRebalance returns a boolean if a field has been set.

### GetReplicasToRebalance

`func (o *RebalanceStateStats) GetReplicasToRebalance() int32`

GetReplicasToRebalance returns the ReplicasToRebalance field if non-nil, zero value otherwise.

### GetReplicasToRebalanceOk

`func (o *RebalanceStateStats) GetReplicasToRebalanceOk() (*int32, bool)`

GetReplicasToRebalanceOk returns a tuple with the ReplicasToRebalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicasToRebalance

`func (o *RebalanceStateStats) SetReplicasToRebalance(v int32)`

SetReplicasToRebalance sets ReplicasToRebalance field to given value.

### HasReplicasToRebalance

`func (o *RebalanceStateStats) HasReplicasToRebalance() bool`

HasReplicasToRebalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


