# DedupConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedupEnabled** | **bool** |  | [readonly] 
**HashFunction** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewDedupConfig

`func NewDedupConfig(dedupEnabled bool, ) *DedupConfig`

NewDedupConfig instantiates a new DedupConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedupConfigWithDefaults

`func NewDedupConfigWithDefaults() *DedupConfig`

NewDedupConfigWithDefaults instantiates a new DedupConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedupEnabled

`func (o *DedupConfig) GetDedupEnabled() bool`

GetDedupEnabled returns the DedupEnabled field if non-nil, zero value otherwise.

### GetDedupEnabledOk

`func (o *DedupConfig) GetDedupEnabledOk() (*bool, bool)`

GetDedupEnabledOk returns a tuple with the DedupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupEnabled

`func (o *DedupConfig) SetDedupEnabled(v bool)`

SetDedupEnabled sets DedupEnabled field to given value.


### GetHashFunction

`func (o *DedupConfig) GetHashFunction() string`

GetHashFunction returns the HashFunction field if non-nil, zero value otherwise.

### GetHashFunctionOk

`func (o *DedupConfig) GetHashFunctionOk() (*string, bool)`

GetHashFunctionOk returns a tuple with the HashFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashFunction

`func (o *DedupConfig) SetHashFunction(v string)`

SetHashFunction sets HashFunction field to given value.

### HasHashFunction

`func (o *DedupConfig) HasHashFunction() bool`

HasHashFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


