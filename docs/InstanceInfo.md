# InstanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**InstanceName** | Pointer to **string** |  | [optional] 

## Methods

### NewInstanceInfo

`func NewInstanceInfo() *InstanceInfo`

NewInstanceInfo instantiates a new InstanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceInfoWithDefaults

`func NewInstanceInfoWithDefaults() *InstanceInfo`

NewInstanceInfoWithDefaults instantiates a new InstanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *InstanceInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InstanceInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InstanceInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InstanceInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetHost

`func (o *InstanceInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InstanceInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InstanceInfo) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InstanceInfo) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetInstanceName

`func (o *InstanceInfo) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *InstanceInfo) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *InstanceInfo) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *InstanceInfo) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


