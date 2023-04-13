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

// checks if the SegmentAssignmentConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentAssignmentConfig{}

// SegmentAssignmentConfig struct for SegmentAssignmentConfig
type SegmentAssignmentConfig struct {
	AssignmentStrategy *string `json:"assignmentStrategy,omitempty"`
}

// NewSegmentAssignmentConfig instantiates a new SegmentAssignmentConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentAssignmentConfig() *SegmentAssignmentConfig {
	this := SegmentAssignmentConfig{}
	return &this
}

// NewSegmentAssignmentConfigWithDefaults instantiates a new SegmentAssignmentConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentAssignmentConfigWithDefaults() *SegmentAssignmentConfig {
	this := SegmentAssignmentConfig{}
	return &this
}

// GetAssignmentStrategy returns the AssignmentStrategy field value if set, zero value otherwise.
func (o *SegmentAssignmentConfig) GetAssignmentStrategy() string {
	if o == nil || IsNil(o.AssignmentStrategy) {
		var ret string
		return ret
	}
	return *o.AssignmentStrategy
}

// GetAssignmentStrategyOk returns a tuple with the AssignmentStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentAssignmentConfig) GetAssignmentStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.AssignmentStrategy) {
		return nil, false
	}
	return o.AssignmentStrategy, true
}

// HasAssignmentStrategy returns a boolean if a field has been set.
func (o *SegmentAssignmentConfig) HasAssignmentStrategy() bool {
	if o != nil && !IsNil(o.AssignmentStrategy) {
		return true
	}

	return false
}

// SetAssignmentStrategy gets a reference to the given string and assigns it to the AssignmentStrategy field.
func (o *SegmentAssignmentConfig) SetAssignmentStrategy(v string) {
	o.AssignmentStrategy = &v
}

func (o SegmentAssignmentConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentAssignmentConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssignmentStrategy) {
		toSerialize["assignmentStrategy"] = o.AssignmentStrategy
	}
	return toSerialize, nil
}

type NullableSegmentAssignmentConfig struct {
	value *SegmentAssignmentConfig
	isSet bool
}

func (v NullableSegmentAssignmentConfig) Get() *SegmentAssignmentConfig {
	return v.value
}

func (v *NullableSegmentAssignmentConfig) Set(val *SegmentAssignmentConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentAssignmentConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentAssignmentConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentAssignmentConfig(val *SegmentAssignmentConfig) *NullableSegmentAssignmentConfig {
	return &NullableSegmentAssignmentConfig{value: val, isSet: true}
}

func (v NullableSegmentAssignmentConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentAssignmentConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


