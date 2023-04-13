# RebalanceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**InstanceAssignment** | Pointer to [**map[string]InstancePartitions**](InstancePartitions.md) |  | [optional] [readonly] 
**TierInstanceAssignment** | Pointer to [**map[string]InstancePartitions**](InstancePartitions.md) |  | [optional] [readonly] 
**SegmentAssignment** | Pointer to **map[string]map[string]string** |  | [optional] [readonly] 

## Methods

### NewRebalanceResult

`func NewRebalanceResult() *RebalanceResult`

NewRebalanceResult instantiates a new RebalanceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebalanceResultWithDefaults

`func NewRebalanceResultWithDefaults() *RebalanceResult`

NewRebalanceResultWithDefaults instantiates a new RebalanceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *RebalanceResult) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *RebalanceResult) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *RebalanceResult) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *RebalanceResult) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStatus

`func (o *RebalanceResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RebalanceResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RebalanceResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RebalanceResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *RebalanceResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RebalanceResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RebalanceResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RebalanceResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstanceAssignment

`func (o *RebalanceResult) GetInstanceAssignment() map[string]InstancePartitions`

GetInstanceAssignment returns the InstanceAssignment field if non-nil, zero value otherwise.

### GetInstanceAssignmentOk

`func (o *RebalanceResult) GetInstanceAssignmentOk() (*map[string]InstancePartitions, bool)`

GetInstanceAssignmentOk returns a tuple with the InstanceAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAssignment

`func (o *RebalanceResult) SetInstanceAssignment(v map[string]InstancePartitions)`

SetInstanceAssignment sets InstanceAssignment field to given value.

### HasInstanceAssignment

`func (o *RebalanceResult) HasInstanceAssignment() bool`

HasInstanceAssignment returns a boolean if a field has been set.

### GetTierInstanceAssignment

`func (o *RebalanceResult) GetTierInstanceAssignment() map[string]InstancePartitions`

GetTierInstanceAssignment returns the TierInstanceAssignment field if non-nil, zero value otherwise.

### GetTierInstanceAssignmentOk

`func (o *RebalanceResult) GetTierInstanceAssignmentOk() (*map[string]InstancePartitions, bool)`

GetTierInstanceAssignmentOk returns a tuple with the TierInstanceAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierInstanceAssignment

`func (o *RebalanceResult) SetTierInstanceAssignment(v map[string]InstancePartitions)`

SetTierInstanceAssignment sets TierInstanceAssignment field to given value.

### HasTierInstanceAssignment

`func (o *RebalanceResult) HasTierInstanceAssignment() bool`

HasTierInstanceAssignment returns a boolean if a field has been set.

### GetSegmentAssignment

`func (o *RebalanceResult) GetSegmentAssignment() map[string]map[string]string`

GetSegmentAssignment returns the SegmentAssignment field if non-nil, zero value otherwise.

### GetSegmentAssignmentOk

`func (o *RebalanceResult) GetSegmentAssignmentOk() (*map[string]map[string]string, bool)`

GetSegmentAssignmentOk returns a tuple with the SegmentAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentAssignment

`func (o *RebalanceResult) SetSegmentAssignment(v map[string]map[string]string)`

SetSegmentAssignment sets SegmentAssignment field to given value.

### HasSegmentAssignment

`func (o *RebalanceResult) HasSegmentAssignment() bool`

HasSegmentAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


