# ReplicaGroupStrategyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartitionColumn** | Pointer to **string** |  | [optional] [readonly] 
**NumInstancesPerPartition** | **int32** |  | [readonly] 

## Methods

### NewReplicaGroupStrategyConfig

`func NewReplicaGroupStrategyConfig(numInstancesPerPartition int32, ) *ReplicaGroupStrategyConfig`

NewReplicaGroupStrategyConfig instantiates a new ReplicaGroupStrategyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaGroupStrategyConfigWithDefaults

`func NewReplicaGroupStrategyConfigWithDefaults() *ReplicaGroupStrategyConfig`

NewReplicaGroupStrategyConfigWithDefaults instantiates a new ReplicaGroupStrategyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitionColumn

`func (o *ReplicaGroupStrategyConfig) GetPartitionColumn() string`

GetPartitionColumn returns the PartitionColumn field if non-nil, zero value otherwise.

### GetPartitionColumnOk

`func (o *ReplicaGroupStrategyConfig) GetPartitionColumnOk() (*string, bool)`

GetPartitionColumnOk returns a tuple with the PartitionColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionColumn

`func (o *ReplicaGroupStrategyConfig) SetPartitionColumn(v string)`

SetPartitionColumn sets PartitionColumn field to given value.

### HasPartitionColumn

`func (o *ReplicaGroupStrategyConfig) HasPartitionColumn() bool`

HasPartitionColumn returns a boolean if a field has been set.

### GetNumInstancesPerPartition

`func (o *ReplicaGroupStrategyConfig) GetNumInstancesPerPartition() int32`

GetNumInstancesPerPartition returns the NumInstancesPerPartition field if non-nil, zero value otherwise.

### GetNumInstancesPerPartitionOk

`func (o *ReplicaGroupStrategyConfig) GetNumInstancesPerPartitionOk() (*int32, bool)`

GetNumInstancesPerPartitionOk returns a tuple with the NumInstancesPerPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstancesPerPartition

`func (o *ReplicaGroupStrategyConfig) SetNumInstancesPerPartition(v int32)`

SetNumInstancesPerPartition sets NumInstancesPerPartition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


