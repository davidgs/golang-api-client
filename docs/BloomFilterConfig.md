# BloomFilterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [readonly] 
**Fpp** | Pointer to **float64** |  | [optional] [readonly] 
**MaxSizeInBytes** | Pointer to **int32** |  | [optional] [readonly] 
**LoadOnHeap** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewBloomFilterConfig

`func NewBloomFilterConfig() *BloomFilterConfig`

NewBloomFilterConfig instantiates a new BloomFilterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBloomFilterConfigWithDefaults

`func NewBloomFilterConfigWithDefaults() *BloomFilterConfig`

NewBloomFilterConfigWithDefaults instantiates a new BloomFilterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BloomFilterConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BloomFilterConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BloomFilterConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BloomFilterConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFpp

`func (o *BloomFilterConfig) GetFpp() float64`

GetFpp returns the Fpp field if non-nil, zero value otherwise.

### GetFppOk

`func (o *BloomFilterConfig) GetFppOk() (*float64, bool)`

GetFppOk returns a tuple with the Fpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpp

`func (o *BloomFilterConfig) SetFpp(v float64)`

SetFpp sets Fpp field to given value.

### HasFpp

`func (o *BloomFilterConfig) HasFpp() bool`

HasFpp returns a boolean if a field has been set.

### GetMaxSizeInBytes

`func (o *BloomFilterConfig) GetMaxSizeInBytes() int32`

GetMaxSizeInBytes returns the MaxSizeInBytes field if non-nil, zero value otherwise.

### GetMaxSizeInBytesOk

`func (o *BloomFilterConfig) GetMaxSizeInBytesOk() (*int32, bool)`

GetMaxSizeInBytesOk returns a tuple with the MaxSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSizeInBytes

`func (o *BloomFilterConfig) SetMaxSizeInBytes(v int32)`

SetMaxSizeInBytes sets MaxSizeInBytes field to given value.

### HasMaxSizeInBytes

`func (o *BloomFilterConfig) HasMaxSizeInBytes() bool`

HasMaxSizeInBytes returns a boolean if a field has been set.

### GetLoadOnHeap

`func (o *BloomFilterConfig) GetLoadOnHeap() bool`

GetLoadOnHeap returns the LoadOnHeap field if non-nil, zero value otherwise.

### GetLoadOnHeapOk

`func (o *BloomFilterConfig) GetLoadOnHeapOk() (*bool, bool)`

GetLoadOnHeapOk returns a tuple with the LoadOnHeap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOnHeap

`func (o *BloomFilterConfig) SetLoadOnHeap(v bool)`

SetLoadOnHeap sets LoadOnHeap field to given value.

### HasLoadOnHeap

`func (o *BloomFilterConfig) HasLoadOnHeap() bool`

HasLoadOnHeap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


