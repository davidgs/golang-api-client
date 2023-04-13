# InstanceTagPoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** |  | [readonly] 
**PoolBased** | Pointer to **bool** |  | [optional] [readonly] 
**NumPools** | Pointer to **int32** |  | [optional] [readonly] 
**Pools** | Pointer to **[]int32** |  | [optional] [readonly] 

## Methods

### NewInstanceTagPoolConfig

`func NewInstanceTagPoolConfig(tag string, ) *InstanceTagPoolConfig`

NewInstanceTagPoolConfig instantiates a new InstanceTagPoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTagPoolConfigWithDefaults

`func NewInstanceTagPoolConfigWithDefaults() *InstanceTagPoolConfig`

NewInstanceTagPoolConfigWithDefaults instantiates a new InstanceTagPoolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *InstanceTagPoolConfig) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *InstanceTagPoolConfig) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *InstanceTagPoolConfig) SetTag(v string)`

SetTag sets Tag field to given value.


### GetPoolBased

`func (o *InstanceTagPoolConfig) GetPoolBased() bool`

GetPoolBased returns the PoolBased field if non-nil, zero value otherwise.

### GetPoolBasedOk

`func (o *InstanceTagPoolConfig) GetPoolBasedOk() (*bool, bool)`

GetPoolBasedOk returns a tuple with the PoolBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolBased

`func (o *InstanceTagPoolConfig) SetPoolBased(v bool)`

SetPoolBased sets PoolBased field to given value.

### HasPoolBased

`func (o *InstanceTagPoolConfig) HasPoolBased() bool`

HasPoolBased returns a boolean if a field has been set.

### GetNumPools

`func (o *InstanceTagPoolConfig) GetNumPools() int32`

GetNumPools returns the NumPools field if non-nil, zero value otherwise.

### GetNumPoolsOk

`func (o *InstanceTagPoolConfig) GetNumPoolsOk() (*int32, bool)`

GetNumPoolsOk returns a tuple with the NumPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPools

`func (o *InstanceTagPoolConfig) SetNumPools(v int32)`

SetNumPools sets NumPools field to given value.

### HasNumPools

`func (o *InstanceTagPoolConfig) HasNumPools() bool`

HasNumPools returns a boolean if a field has been set.

### GetPools

`func (o *InstanceTagPoolConfig) GetPools() []int32`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *InstanceTagPoolConfig) GetPoolsOk() (*[]int32, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *InstanceTagPoolConfig) SetPools(v []int32)`

SetPools sets Pools field to given value.

### HasPools

`func (o *InstanceTagPoolConfig) HasPools() bool`

HasPools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


