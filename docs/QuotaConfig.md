# QuotaConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Storage** | Pointer to **string** |  | [optional] [readonly] 
**MaxQueriesPerSecond** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewQuotaConfig

`func NewQuotaConfig() *QuotaConfig`

NewQuotaConfig instantiates a new QuotaConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaConfigWithDefaults

`func NewQuotaConfigWithDefaults() *QuotaConfig`

NewQuotaConfigWithDefaults instantiates a new QuotaConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorage

`func (o *QuotaConfig) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *QuotaConfig) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *QuotaConfig) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *QuotaConfig) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetMaxQueriesPerSecond

`func (o *QuotaConfig) GetMaxQueriesPerSecond() string`

GetMaxQueriesPerSecond returns the MaxQueriesPerSecond field if non-nil, zero value otherwise.

### GetMaxQueriesPerSecondOk

`func (o *QuotaConfig) GetMaxQueriesPerSecondOk() (*string, bool)`

GetMaxQueriesPerSecondOk returns a tuple with the MaxQueriesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueriesPerSecond

`func (o *QuotaConfig) SetMaxQueriesPerSecond(v string)`

SetMaxQueriesPerSecond sets MaxQueriesPerSecond field to given value.

### HasMaxQueriesPerSecond

`func (o *QuotaConfig) HasMaxQueriesPerSecond() bool`

HasMaxQueriesPerSecond returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


