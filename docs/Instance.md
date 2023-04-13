# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | [readonly] 
**Port** | **int32** |  | [readonly] 
**Type** | **string** |  | [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] [readonly] 
**Pools** | Pointer to **map[string]int32** |  | [optional] [readonly] 
**GrpcPort** | Pointer to **int32** |  | [optional] [readonly] 
**AdminPort** | Pointer to **int32** |  | [optional] [readonly] 
**QueryServicePort** | Pointer to **int32** |  | [optional] [readonly] 
**QueryMailboxPort** | Pointer to **int32** |  | [optional] [readonly] 
**QueriesDisabled** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewInstance

`func NewInstance(host string, port int32, type_ string, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *Instance) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Instance) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Instance) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *Instance) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Instance) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Instance) SetPort(v int32)`

SetPort sets Port field to given value.


### GetType

`func (o *Instance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Instance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Instance) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *Instance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Instance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Instance) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Instance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPools

`func (o *Instance) GetPools() map[string]int32`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *Instance) GetPoolsOk() (*map[string]int32, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *Instance) SetPools(v map[string]int32)`

SetPools sets Pools field to given value.

### HasPools

`func (o *Instance) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetGrpcPort

`func (o *Instance) GetGrpcPort() int32`

GetGrpcPort returns the GrpcPort field if non-nil, zero value otherwise.

### GetGrpcPortOk

`func (o *Instance) GetGrpcPortOk() (*int32, bool)`

GetGrpcPortOk returns a tuple with the GrpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpcPort

`func (o *Instance) SetGrpcPort(v int32)`

SetGrpcPort sets GrpcPort field to given value.

### HasGrpcPort

`func (o *Instance) HasGrpcPort() bool`

HasGrpcPort returns a boolean if a field has been set.

### GetAdminPort

`func (o *Instance) GetAdminPort() int32`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *Instance) GetAdminPortOk() (*int32, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *Instance) SetAdminPort(v int32)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *Instance) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### GetQueryServicePort

`func (o *Instance) GetQueryServicePort() int32`

GetQueryServicePort returns the QueryServicePort field if non-nil, zero value otherwise.

### GetQueryServicePortOk

`func (o *Instance) GetQueryServicePortOk() (*int32, bool)`

GetQueryServicePortOk returns a tuple with the QueryServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryServicePort

`func (o *Instance) SetQueryServicePort(v int32)`

SetQueryServicePort sets QueryServicePort field to given value.

### HasQueryServicePort

`func (o *Instance) HasQueryServicePort() bool`

HasQueryServicePort returns a boolean if a field has been set.

### GetQueryMailboxPort

`func (o *Instance) GetQueryMailboxPort() int32`

GetQueryMailboxPort returns the QueryMailboxPort field if non-nil, zero value otherwise.

### GetQueryMailboxPortOk

`func (o *Instance) GetQueryMailboxPortOk() (*int32, bool)`

GetQueryMailboxPortOk returns a tuple with the QueryMailboxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryMailboxPort

`func (o *Instance) SetQueryMailboxPort(v int32)`

SetQueryMailboxPort sets QueryMailboxPort field to given value.

### HasQueryMailboxPort

`func (o *Instance) HasQueryMailboxPort() bool`

HasQueryMailboxPort returns a boolean if a field has been set.

### GetQueriesDisabled

`func (o *Instance) GetQueriesDisabled() bool`

GetQueriesDisabled returns the QueriesDisabled field if non-nil, zero value otherwise.

### GetQueriesDisabledOk

`func (o *Instance) GetQueriesDisabledOk() (*bool, bool)`

GetQueriesDisabledOk returns a tuple with the QueriesDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueriesDisabled

`func (o *Instance) SetQueriesDisabled(v bool)`

SetQueriesDisabled sets QueriesDisabled field to given value.

### HasQueriesDisabled

`func (o *Instance) HasQueriesDisabled() bool`

HasQueriesDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


