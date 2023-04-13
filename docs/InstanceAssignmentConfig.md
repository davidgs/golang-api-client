# InstanceAssignmentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagPoolConfig** | [**InstanceTagPoolConfig**](InstanceTagPoolConfig.md) |  | 
**ConstraintConfig** | Pointer to [**InstanceConstraintConfig**](InstanceConstraintConfig.md) |  | [optional] 
**ReplicaGroupPartitionConfig** | [**InstanceReplicaGroupPartitionConfig**](InstanceReplicaGroupPartitionConfig.md) |  | 
**PartitionSelector** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewInstanceAssignmentConfig

`func NewInstanceAssignmentConfig(tagPoolConfig InstanceTagPoolConfig, replicaGroupPartitionConfig InstanceReplicaGroupPartitionConfig, ) *InstanceAssignmentConfig`

NewInstanceAssignmentConfig instantiates a new InstanceAssignmentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceAssignmentConfigWithDefaults

`func NewInstanceAssignmentConfigWithDefaults() *InstanceAssignmentConfig`

NewInstanceAssignmentConfigWithDefaults instantiates a new InstanceAssignmentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagPoolConfig

`func (o *InstanceAssignmentConfig) GetTagPoolConfig() InstanceTagPoolConfig`

GetTagPoolConfig returns the TagPoolConfig field if non-nil, zero value otherwise.

### GetTagPoolConfigOk

`func (o *InstanceAssignmentConfig) GetTagPoolConfigOk() (*InstanceTagPoolConfig, bool)`

GetTagPoolConfigOk returns a tuple with the TagPoolConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagPoolConfig

`func (o *InstanceAssignmentConfig) SetTagPoolConfig(v InstanceTagPoolConfig)`

SetTagPoolConfig sets TagPoolConfig field to given value.


### GetConstraintConfig

`func (o *InstanceAssignmentConfig) GetConstraintConfig() InstanceConstraintConfig`

GetConstraintConfig returns the ConstraintConfig field if non-nil, zero value otherwise.

### GetConstraintConfigOk

`func (o *InstanceAssignmentConfig) GetConstraintConfigOk() (*InstanceConstraintConfig, bool)`

GetConstraintConfigOk returns a tuple with the ConstraintConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintConfig

`func (o *InstanceAssignmentConfig) SetConstraintConfig(v InstanceConstraintConfig)`

SetConstraintConfig sets ConstraintConfig field to given value.

### HasConstraintConfig

`func (o *InstanceAssignmentConfig) HasConstraintConfig() bool`

HasConstraintConfig returns a boolean if a field has been set.

### GetReplicaGroupPartitionConfig

`func (o *InstanceAssignmentConfig) GetReplicaGroupPartitionConfig() InstanceReplicaGroupPartitionConfig`

GetReplicaGroupPartitionConfig returns the ReplicaGroupPartitionConfig field if non-nil, zero value otherwise.

### GetReplicaGroupPartitionConfigOk

`func (o *InstanceAssignmentConfig) GetReplicaGroupPartitionConfigOk() (*InstanceReplicaGroupPartitionConfig, bool)`

GetReplicaGroupPartitionConfigOk returns a tuple with the ReplicaGroupPartitionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaGroupPartitionConfig

`func (o *InstanceAssignmentConfig) SetReplicaGroupPartitionConfig(v InstanceReplicaGroupPartitionConfig)`

SetReplicaGroupPartitionConfig sets ReplicaGroupPartitionConfig field to given value.


### GetPartitionSelector

`func (o *InstanceAssignmentConfig) GetPartitionSelector() string`

GetPartitionSelector returns the PartitionSelector field if non-nil, zero value otherwise.

### GetPartitionSelectorOk

`func (o *InstanceAssignmentConfig) GetPartitionSelectorOk() (*string, bool)`

GetPartitionSelectorOk returns a tuple with the PartitionSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionSelector

`func (o *InstanceAssignmentConfig) SetPartitionSelector(v string)`

SetPartitionSelector sets PartitionSelector field to given value.

### HasPartitionSelector

`func (o *InstanceAssignmentConfig) HasPartitionSelector() bool`

HasPartitionSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


