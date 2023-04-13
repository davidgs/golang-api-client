# InstanceReplicaGroupPartitionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplicaGroupBased** | Pointer to **bool** |  | [optional] [readonly] 
**NumInstances** | Pointer to **int32** |  | [optional] [readonly] 
**NumReplicaGroups** | Pointer to **int32** |  | [optional] [readonly] 
**NumInstancesPerReplicaGroup** | Pointer to **int32** |  | [optional] [readonly] 
**NumPartitions** | Pointer to **int32** |  | [optional] [readonly] 
**NumInstancesPerPartition** | Pointer to **int32** |  | [optional] [readonly] 
**MinimizeDataMovement** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewInstanceReplicaGroupPartitionConfig

`func NewInstanceReplicaGroupPartitionConfig() *InstanceReplicaGroupPartitionConfig`

NewInstanceReplicaGroupPartitionConfig instantiates a new InstanceReplicaGroupPartitionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceReplicaGroupPartitionConfigWithDefaults

`func NewInstanceReplicaGroupPartitionConfigWithDefaults() *InstanceReplicaGroupPartitionConfig`

NewInstanceReplicaGroupPartitionConfigWithDefaults instantiates a new InstanceReplicaGroupPartitionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicaGroupBased

`func (o *InstanceReplicaGroupPartitionConfig) GetReplicaGroupBased() bool`

GetReplicaGroupBased returns the ReplicaGroupBased field if non-nil, zero value otherwise.

### GetReplicaGroupBasedOk

`func (o *InstanceReplicaGroupPartitionConfig) GetReplicaGroupBasedOk() (*bool, bool)`

GetReplicaGroupBasedOk returns a tuple with the ReplicaGroupBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaGroupBased

`func (o *InstanceReplicaGroupPartitionConfig) SetReplicaGroupBased(v bool)`

SetReplicaGroupBased sets ReplicaGroupBased field to given value.

### HasReplicaGroupBased

`func (o *InstanceReplicaGroupPartitionConfig) HasReplicaGroupBased() bool`

HasReplicaGroupBased returns a boolean if a field has been set.

### GetNumInstances

`func (o *InstanceReplicaGroupPartitionConfig) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *InstanceReplicaGroupPartitionConfig) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *InstanceReplicaGroupPartitionConfig) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.

### HasNumInstances

`func (o *InstanceReplicaGroupPartitionConfig) HasNumInstances() bool`

HasNumInstances returns a boolean if a field has been set.

### GetNumReplicaGroups

`func (o *InstanceReplicaGroupPartitionConfig) GetNumReplicaGroups() int32`

GetNumReplicaGroups returns the NumReplicaGroups field if non-nil, zero value otherwise.

### GetNumReplicaGroupsOk

`func (o *InstanceReplicaGroupPartitionConfig) GetNumReplicaGroupsOk() (*int32, bool)`

GetNumReplicaGroupsOk returns a tuple with the NumReplicaGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicaGroups

`func (o *InstanceReplicaGroupPartitionConfig) SetNumReplicaGroups(v int32)`

SetNumReplicaGroups sets NumReplicaGroups field to given value.

### HasNumReplicaGroups

`func (o *InstanceReplicaGroupPartitionConfig) HasNumReplicaGroups() bool`

HasNumReplicaGroups returns a boolean if a field has been set.

### GetNumInstancesPerReplicaGroup

`func (o *InstanceReplicaGroupPartitionConfig) GetNumInstancesPerReplicaGroup() int32`

GetNumInstancesPerReplicaGroup returns the NumInstancesPerReplicaGroup field if non-nil, zero value otherwise.

### GetNumInstancesPerReplicaGroupOk

`func (o *InstanceReplicaGroupPartitionConfig) GetNumInstancesPerReplicaGroupOk() (*int32, bool)`

GetNumInstancesPerReplicaGroupOk returns a tuple with the NumInstancesPerReplicaGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstancesPerReplicaGroup

`func (o *InstanceReplicaGroupPartitionConfig) SetNumInstancesPerReplicaGroup(v int32)`

SetNumInstancesPerReplicaGroup sets NumInstancesPerReplicaGroup field to given value.

### HasNumInstancesPerReplicaGroup

`func (o *InstanceReplicaGroupPartitionConfig) HasNumInstancesPerReplicaGroup() bool`

HasNumInstancesPerReplicaGroup returns a boolean if a field has been set.

### GetNumPartitions

`func (o *InstanceReplicaGroupPartitionConfig) GetNumPartitions() int32`

GetNumPartitions returns the NumPartitions field if non-nil, zero value otherwise.

### GetNumPartitionsOk

`func (o *InstanceReplicaGroupPartitionConfig) GetNumPartitionsOk() (*int32, bool)`

GetNumPartitionsOk returns a tuple with the NumPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitions

`func (o *InstanceReplicaGroupPartitionConfig) SetNumPartitions(v int32)`

SetNumPartitions sets NumPartitions field to given value.

### HasNumPartitions

`func (o *InstanceReplicaGroupPartitionConfig) HasNumPartitions() bool`

HasNumPartitions returns a boolean if a field has been set.

### GetNumInstancesPerPartition

`func (o *InstanceReplicaGroupPartitionConfig) GetNumInstancesPerPartition() int32`

GetNumInstancesPerPartition returns the NumInstancesPerPartition field if non-nil, zero value otherwise.

### GetNumInstancesPerPartitionOk

`func (o *InstanceReplicaGroupPartitionConfig) GetNumInstancesPerPartitionOk() (*int32, bool)`

GetNumInstancesPerPartitionOk returns a tuple with the NumInstancesPerPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstancesPerPartition

`func (o *InstanceReplicaGroupPartitionConfig) SetNumInstancesPerPartition(v int32)`

SetNumInstancesPerPartition sets NumInstancesPerPartition field to given value.

### HasNumInstancesPerPartition

`func (o *InstanceReplicaGroupPartitionConfig) HasNumInstancesPerPartition() bool`

HasNumInstancesPerPartition returns a boolean if a field has been set.

### GetMinimizeDataMovement

`func (o *InstanceReplicaGroupPartitionConfig) GetMinimizeDataMovement() bool`

GetMinimizeDataMovement returns the MinimizeDataMovement field if non-nil, zero value otherwise.

### GetMinimizeDataMovementOk

`func (o *InstanceReplicaGroupPartitionConfig) GetMinimizeDataMovementOk() (*bool, bool)`

GetMinimizeDataMovementOk returns a tuple with the MinimizeDataMovement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimizeDataMovement

`func (o *InstanceReplicaGroupPartitionConfig) SetMinimizeDataMovement(v bool)`

SetMinimizeDataMovement sets MinimizeDataMovement field to given value.

### HasMinimizeDataMovement

`func (o *InstanceReplicaGroupPartitionConfig) HasMinimizeDataMovement() bool`

HasMinimizeDataMovement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


