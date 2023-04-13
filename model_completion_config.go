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

// checks if the CompletionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompletionConfig{}

// CompletionConfig struct for CompletionConfig
type CompletionConfig struct {
	CompletionMode string `json:"completionMode"`
}

// NewCompletionConfig instantiates a new CompletionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletionConfig(completionMode string) *CompletionConfig {
	this := CompletionConfig{}
	this.CompletionMode = completionMode
	return &this
}

// NewCompletionConfigWithDefaults instantiates a new CompletionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletionConfigWithDefaults() *CompletionConfig {
	this := CompletionConfig{}
	return &this
}

// GetCompletionMode returns the CompletionMode field value
func (o *CompletionConfig) GetCompletionMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompletionMode
}

// GetCompletionModeOk returns a tuple with the CompletionMode field value
// and a boolean to check if the value has been set.
func (o *CompletionConfig) GetCompletionModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletionMode, true
}

// SetCompletionMode sets field value
func (o *CompletionConfig) SetCompletionMode(v string) {
	o.CompletionMode = v
}

func (o CompletionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompletionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: completionMode is readOnly
	return toSerialize, nil
}

type NullableCompletionConfig struct {
	value *CompletionConfig
	isSet bool
}

func (v NullableCompletionConfig) Get() *CompletionConfig {
	return v.value
}

func (v *NullableCompletionConfig) Set(val *CompletionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletionConfig(val *CompletionConfig) *NullableCompletionConfig {
	return &NullableCompletionConfig{value: val, isSet: true}
}

func (v NullableCompletionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

