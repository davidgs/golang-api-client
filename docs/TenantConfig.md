# TenantConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Broker** | Pointer to **string** |  | [optional] [readonly] 
**Server** | Pointer to **string** |  | [optional] [readonly] 
**TagOverrideConfig** | Pointer to [**TagOverrideConfig**](TagOverrideConfig.md) |  | [optional] 

## Methods

### NewTenantConfig

`func NewTenantConfig() *TenantConfig`

NewTenantConfig instantiates a new TenantConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantConfigWithDefaults

`func NewTenantConfigWithDefaults() *TenantConfig`

NewTenantConfigWithDefaults instantiates a new TenantConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBroker

`func (o *TenantConfig) GetBroker() string`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *TenantConfig) GetBrokerOk() (*string, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *TenantConfig) SetBroker(v string)`

SetBroker sets Broker field to given value.

### HasBroker

`func (o *TenantConfig) HasBroker() bool`

HasBroker returns a boolean if a field has been set.

### GetServer

`func (o *TenantConfig) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *TenantConfig) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *TenantConfig) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *TenantConfig) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetTagOverrideConfig

`func (o *TenantConfig) GetTagOverrideConfig() TagOverrideConfig`

GetTagOverrideConfig returns the TagOverrideConfig field if non-nil, zero value otherwise.

### GetTagOverrideConfigOk

`func (o *TenantConfig) GetTagOverrideConfigOk() (*TagOverrideConfig, bool)`

GetTagOverrideConfigOk returns a tuple with the TagOverrideConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagOverrideConfig

`func (o *TenantConfig) SetTagOverrideConfig(v TagOverrideConfig)`

SetTagOverrideConfig sets TagOverrideConfig field to given value.

### HasTagOverrideConfig

`func (o *TenantConfig) HasTagOverrideConfig() bool`

HasTagOverrideConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


