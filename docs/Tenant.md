# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantRole** | **string** |  | [readonly] 
**TenantName** | **string** |  | [readonly] 
**NumberOfInstances** | Pointer to **int32** |  | [optional] [readonly] 
**OfflineInstances** | Pointer to **int32** |  | [optional] [readonly] 
**RealtimeInstances** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewTenant

`func NewTenant(tenantRole string, tenantName string, ) *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantRole

`func (o *Tenant) GetTenantRole() string`

GetTenantRole returns the TenantRole field if non-nil, zero value otherwise.

### GetTenantRoleOk

`func (o *Tenant) GetTenantRoleOk() (*string, bool)`

GetTenantRoleOk returns a tuple with the TenantRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantRole

`func (o *Tenant) SetTenantRole(v string)`

SetTenantRole sets TenantRole field to given value.


### GetTenantName

`func (o *Tenant) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *Tenant) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *Tenant) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.


### GetNumberOfInstances

`func (o *Tenant) GetNumberOfInstances() int32`

GetNumberOfInstances returns the NumberOfInstances field if non-nil, zero value otherwise.

### GetNumberOfInstancesOk

`func (o *Tenant) GetNumberOfInstancesOk() (*int32, bool)`

GetNumberOfInstancesOk returns a tuple with the NumberOfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfInstances

`func (o *Tenant) SetNumberOfInstances(v int32)`

SetNumberOfInstances sets NumberOfInstances field to given value.

### HasNumberOfInstances

`func (o *Tenant) HasNumberOfInstances() bool`

HasNumberOfInstances returns a boolean if a field has been set.

### GetOfflineInstances

`func (o *Tenant) GetOfflineInstances() int32`

GetOfflineInstances returns the OfflineInstances field if non-nil, zero value otherwise.

### GetOfflineInstancesOk

`func (o *Tenant) GetOfflineInstancesOk() (*int32, bool)`

GetOfflineInstancesOk returns a tuple with the OfflineInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineInstances

`func (o *Tenant) SetOfflineInstances(v int32)`

SetOfflineInstances sets OfflineInstances field to given value.

### HasOfflineInstances

`func (o *Tenant) HasOfflineInstances() bool`

HasOfflineInstances returns a boolean if a field has been set.

### GetRealtimeInstances

`func (o *Tenant) GetRealtimeInstances() int32`

GetRealtimeInstances returns the RealtimeInstances field if non-nil, zero value otherwise.

### GetRealtimeInstancesOk

`func (o *Tenant) GetRealtimeInstancesOk() (*int32, bool)`

GetRealtimeInstancesOk returns a tuple with the RealtimeInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtimeInstances

`func (o *Tenant) SetRealtimeInstances(v int32)`

SetRealtimeInstances sets RealtimeInstances field to given value.

### HasRealtimeInstances

`func (o *Tenant) HasRealtimeInstances() bool`

HasRealtimeInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


