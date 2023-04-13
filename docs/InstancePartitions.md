# InstancePartitions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstancePartitionsName** | Pointer to **string** |  | [optional] [readonly] 
**PartitionToInstancesMap** | Pointer to **map[string][]string** |  | [optional] [readonly] 

## Methods

### NewInstancePartitions

`func NewInstancePartitions() *InstancePartitions`

NewInstancePartitions instantiates a new InstancePartitions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancePartitionsWithDefaults

`func NewInstancePartitionsWithDefaults() *InstancePartitions`

NewInstancePartitionsWithDefaults instantiates a new InstancePartitions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstancePartitionsName

`func (o *InstancePartitions) GetInstancePartitionsName() string`

GetInstancePartitionsName returns the InstancePartitionsName field if non-nil, zero value otherwise.

### GetInstancePartitionsNameOk

`func (o *InstancePartitions) GetInstancePartitionsNameOk() (*string, bool)`

GetInstancePartitionsNameOk returns a tuple with the InstancePartitionsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePartitionsName

`func (o *InstancePartitions) SetInstancePartitionsName(v string)`

SetInstancePartitionsName sets InstancePartitionsName field to given value.

### HasInstancePartitionsName

`func (o *InstancePartitions) HasInstancePartitionsName() bool`

HasInstancePartitionsName returns a boolean if a field has been set.

### GetPartitionToInstancesMap

`func (o *InstancePartitions) GetPartitionToInstancesMap() map[string][]string`

GetPartitionToInstancesMap returns the PartitionToInstancesMap field if non-nil, zero value otherwise.

### GetPartitionToInstancesMapOk

`func (o *InstancePartitions) GetPartitionToInstancesMapOk() (*map[string][]string, bool)`

GetPartitionToInstancesMapOk returns a tuple with the PartitionToInstancesMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionToInstancesMap

`func (o *InstancePartitions) SetPartitionToInstancesMap(v map[string][]string)`

SetPartitionToInstancesMap sets PartitionToInstancesMap field to given value.

### HasPartitionToInstancesMap

`func (o *InstancePartitions) HasPartitionToInstancesMap() bool`

HasPartitionToInstancesMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


