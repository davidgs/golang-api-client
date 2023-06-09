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

// checks if the TableTaskConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TableTaskConfig{}

// TableTaskConfig struct for TableTaskConfig
type TableTaskConfig struct {
	TaskTypeConfigsMap *map[string]map[string]string `json:"taskTypeConfigsMap,omitempty"`
}

// NewTableTaskConfig instantiates a new TableTaskConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableTaskConfig() *TableTaskConfig {
	this := TableTaskConfig{}
	return &this
}

// NewTableTaskConfigWithDefaults instantiates a new TableTaskConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableTaskConfigWithDefaults() *TableTaskConfig {
	this := TableTaskConfig{}
	return &this
}

// GetTaskTypeConfigsMap returns the TaskTypeConfigsMap field value if set, zero value otherwise.
func (o *TableTaskConfig) GetTaskTypeConfigsMap() map[string]map[string]string {
	if o == nil || IsNil(o.TaskTypeConfigsMap) {
		var ret map[string]map[string]string
		return ret
	}
	return *o.TaskTypeConfigsMap
}

// GetTaskTypeConfigsMapOk returns a tuple with the TaskTypeConfigsMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableTaskConfig) GetTaskTypeConfigsMapOk() (*map[string]map[string]string, bool) {
	if o == nil || IsNil(o.TaskTypeConfigsMap) {
		return nil, false
	}
	return o.TaskTypeConfigsMap, true
}

// HasTaskTypeConfigsMap returns a boolean if a field has been set.
func (o *TableTaskConfig) HasTaskTypeConfigsMap() bool {
	if o != nil && !IsNil(o.TaskTypeConfigsMap) {
		return true
	}

	return false
}

// SetTaskTypeConfigsMap gets a reference to the given map[string]map[string]string and assigns it to the TaskTypeConfigsMap field.
func (o *TableTaskConfig) SetTaskTypeConfigsMap(v map[string]map[string]string) {
	o.TaskTypeConfigsMap = &v
}

func (o TableTaskConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TableTaskConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: taskTypeConfigsMap is readOnly
	return toSerialize, nil
}

type NullableTableTaskConfig struct {
	value *TableTaskConfig
	isSet bool
}

func (v NullableTableTaskConfig) Get() *TableTaskConfig {
	return v.value
}

func (v *NullableTableTaskConfig) Set(val *TableTaskConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTableTaskConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTableTaskConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableTaskConfig(val *TableTaskConfig) *NullableTableTaskConfig {
	return &NullableTableTaskConfig{value: val, isSet: true}
}

func (v NullableTableTaskConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableTaskConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


