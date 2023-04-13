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

// checks if the TimestampConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimestampConfig{}

// TimestampConfig struct for TimestampConfig
type TimestampConfig struct {
	Granularities []string `json:"granularities,omitempty"`
}

// NewTimestampConfig instantiates a new TimestampConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimestampConfig() *TimestampConfig {
	this := TimestampConfig{}
	return &this
}

// NewTimestampConfigWithDefaults instantiates a new TimestampConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimestampConfigWithDefaults() *TimestampConfig {
	this := TimestampConfig{}
	return &this
}

// GetGranularities returns the Granularities field value if set, zero value otherwise.
func (o *TimestampConfig) GetGranularities() []string {
	if o == nil || IsNil(o.Granularities) {
		var ret []string
		return ret
	}
	return o.Granularities
}

// GetGranularitiesOk returns a tuple with the Granularities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimestampConfig) GetGranularitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Granularities) {
		return nil, false
	}
	return o.Granularities, true
}

// HasGranularities returns a boolean if a field has been set.
func (o *TimestampConfig) HasGranularities() bool {
	if o != nil && !IsNil(o.Granularities) {
		return true
	}

	return false
}

// SetGranularities gets a reference to the given []string and assigns it to the Granularities field.
func (o *TimestampConfig) SetGranularities(v []string) {
	o.Granularities = v
}

func (o TimestampConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimestampConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: granularities is readOnly
	return toSerialize, nil
}

type NullableTimestampConfig struct {
	value *TimestampConfig
	isSet bool
}

func (v NullableTimestampConfig) Get() *TimestampConfig {
	return v.value
}

func (v *NullableTimestampConfig) Set(val *TimestampConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTimestampConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTimestampConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimestampConfig(val *TimestampConfig) *NullableTimestampConfig {
	return &NullableTimestampConfig{value: val, isSet: true}
}

func (v NullableTimestampConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimestampConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


