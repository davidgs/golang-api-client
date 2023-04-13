# TunerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**TunerProperties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTunerConfig

`func NewTunerConfig(name string, ) *TunerConfig`

NewTunerConfig instantiates a new TunerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunerConfigWithDefaults

`func NewTunerConfigWithDefaults() *TunerConfig`

NewTunerConfigWithDefaults instantiates a new TunerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TunerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TunerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TunerConfig) SetName(v string)`

SetName sets Name field to given value.


### GetTunerProperties

`func (o *TunerConfig) GetTunerProperties() map[string]string`

GetTunerProperties returns the TunerProperties field if non-nil, zero value otherwise.

### GetTunerPropertiesOk

`func (o *TunerConfig) GetTunerPropertiesOk() (*map[string]string, bool)`

GetTunerPropertiesOk returns a tuple with the TunerProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunerProperties

`func (o *TunerConfig) SetTunerProperties(v map[string]string)`

SetTunerProperties sets TunerProperties field to given value.

### HasTunerProperties

`func (o *TunerConfig) HasTunerProperties() bool`

HasTunerProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


