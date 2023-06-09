/*
Pinot Controller API

APIs for accessing Pinot Controller information

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BloomFilterConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BloomFilterConfig{}

// BloomFilterConfig struct for BloomFilterConfig
type BloomFilterConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
	Fpp *float64 `json:"fpp,omitempty"`
	MaxSizeInBytes *int32 `json:"maxSizeInBytes,omitempty"`
	LoadOnHeap *bool `json:"loadOnHeap,omitempty"`
}

// NewBloomFilterConfig instantiates a new BloomFilterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBloomFilterConfig() *BloomFilterConfig {
	this := BloomFilterConfig{}
	return &this
}

// NewBloomFilterConfigWithDefaults instantiates a new BloomFilterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBloomFilterConfigWithDefaults() *BloomFilterConfig {
	this := BloomFilterConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BloomFilterConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BloomFilterConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BloomFilterConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BloomFilterConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFpp returns the Fpp field value if set, zero value otherwise.
func (o *BloomFilterConfig) GetFpp() float64 {
	if o == nil || IsNil(o.Fpp) {
		var ret float64
		return ret
	}
	return *o.Fpp
}

// GetFppOk returns a tuple with the Fpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BloomFilterConfig) GetFppOk() (*float64, bool) {
	if o == nil || IsNil(o.Fpp) {
		return nil, false
	}
	return o.Fpp, true
}

// HasFpp returns a boolean if a field has been set.
func (o *BloomFilterConfig) HasFpp() bool {
	if o != nil && !IsNil(o.Fpp) {
		return true
	}

	return false
}

// SetFpp gets a reference to the given float64 and assigns it to the Fpp field.
func (o *BloomFilterConfig) SetFpp(v float64) {
	o.Fpp = &v
}

// GetMaxSizeInBytes returns the MaxSizeInBytes field value if set, zero value otherwise.
func (o *BloomFilterConfig) GetMaxSizeInBytes() int32 {
	if o == nil || IsNil(o.MaxSizeInBytes) {
		var ret int32
		return ret
	}
	return *o.MaxSizeInBytes
}

// GetMaxSizeInBytesOk returns a tuple with the MaxSizeInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BloomFilterConfig) GetMaxSizeInBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSizeInBytes) {
		return nil, false
	}
	return o.MaxSizeInBytes, true
}

// HasMaxSizeInBytes returns a boolean if a field has been set.
func (o *BloomFilterConfig) HasMaxSizeInBytes() bool {
	if o != nil && !IsNil(o.MaxSizeInBytes) {
		return true
	}

	return false
}

// SetMaxSizeInBytes gets a reference to the given int32 and assigns it to the MaxSizeInBytes field.
func (o *BloomFilterConfig) SetMaxSizeInBytes(v int32) {
	o.MaxSizeInBytes = &v
}

// GetLoadOnHeap returns the LoadOnHeap field value if set, zero value otherwise.
func (o *BloomFilterConfig) GetLoadOnHeap() bool {
	if o == nil || IsNil(o.LoadOnHeap) {
		var ret bool
		return ret
	}
	return *o.LoadOnHeap
}

// GetLoadOnHeapOk returns a tuple with the LoadOnHeap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BloomFilterConfig) GetLoadOnHeapOk() (*bool, bool) {
	if o == nil || IsNil(o.LoadOnHeap) {
		return nil, false
	}
	return o.LoadOnHeap, true
}

// HasLoadOnHeap returns a boolean if a field has been set.
func (o *BloomFilterConfig) HasLoadOnHeap() bool {
	if o != nil && !IsNil(o.LoadOnHeap) {
		return true
	}

	return false
}

// SetLoadOnHeap gets a reference to the given bool and assigns it to the LoadOnHeap field.
func (o *BloomFilterConfig) SetLoadOnHeap(v bool) {
	o.LoadOnHeap = &v
}

func (o BloomFilterConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BloomFilterConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: enabled is readOnly
	// skip: fpp is readOnly
	// skip: maxSizeInBytes is readOnly
	// skip: loadOnHeap is readOnly
	return toSerialize, nil
}

type NullableBloomFilterConfig struct {
	value *BloomFilterConfig
	isSet bool
}

func (v NullableBloomFilterConfig) Get() *BloomFilterConfig {
	return v.value
}

func (v *NullableBloomFilterConfig) Set(val *BloomFilterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBloomFilterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBloomFilterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBloomFilterConfig(val *BloomFilterConfig) *NullableBloomFilterConfig {
	return &NullableBloomFilterConfig{value: val, isSet: true}
}

func (v NullableBloomFilterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBloomFilterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


