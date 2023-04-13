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

// checks if the PinotTaskConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PinotTaskConfig{}

// PinotTaskConfig struct for PinotTaskConfig
type PinotTaskConfig struct {
	TableName *string `json:"tableName,omitempty"`
	Configs *map[string]string `json:"configs,omitempty"`
	TaskId *string `json:"taskId,omitempty"`
	TaskType *string `json:"taskType,omitempty"`
}

// NewPinotTaskConfig instantiates a new PinotTaskConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinotTaskConfig() *PinotTaskConfig {
	this := PinotTaskConfig{}
	return &this
}

// NewPinotTaskConfigWithDefaults instantiates a new PinotTaskConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinotTaskConfigWithDefaults() *PinotTaskConfig {
	this := PinotTaskConfig{}
	return &this
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *PinotTaskConfig) GetTableName() string {
	if o == nil || IsNil(o.TableName) {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinotTaskConfig) GetTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.TableName) {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *PinotTaskConfig) HasTableName() bool {
	if o != nil && !IsNil(o.TableName) {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *PinotTaskConfig) SetTableName(v string) {
	o.TableName = &v
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *PinotTaskConfig) GetConfigs() map[string]string {
	if o == nil || IsNil(o.Configs) {
		var ret map[string]string
		return ret
	}
	return *o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinotTaskConfig) GetConfigsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Configs) {
		return nil, false
	}
	return o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *PinotTaskConfig) HasConfigs() bool {
	if o != nil && !IsNil(o.Configs) {
		return true
	}

	return false
}

// SetConfigs gets a reference to the given map[string]string and assigns it to the Configs field.
func (o *PinotTaskConfig) SetConfigs(v map[string]string) {
	o.Configs = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *PinotTaskConfig) GetTaskId() string {
	if o == nil || IsNil(o.TaskId) {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinotTaskConfig) GetTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaskId) {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *PinotTaskConfig) HasTaskId() bool {
	if o != nil && !IsNil(o.TaskId) {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *PinotTaskConfig) SetTaskId(v string) {
	o.TaskId = &v
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *PinotTaskConfig) GetTaskType() string {
	if o == nil || IsNil(o.TaskType) {
		var ret string
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinotTaskConfig) GetTaskTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaskType) {
		return nil, false
	}
	return o.TaskType, true
}

// HasTaskType returns a boolean if a field has been set.
func (o *PinotTaskConfig) HasTaskType() bool {
	if o != nil && !IsNil(o.TaskType) {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given string and assigns it to the TaskType field.
func (o *PinotTaskConfig) SetTaskType(v string) {
	o.TaskType = &v
}

func (o PinotTaskConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PinotTaskConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TableName) {
		toSerialize["tableName"] = o.TableName
	}
	if !IsNil(o.Configs) {
		toSerialize["configs"] = o.Configs
	}
	if !IsNil(o.TaskId) {
		toSerialize["taskId"] = o.TaskId
	}
	if !IsNil(o.TaskType) {
		toSerialize["taskType"] = o.TaskType
	}
	return toSerialize, nil
}

type NullablePinotTaskConfig struct {
	value *PinotTaskConfig
	isSet bool
}

func (v NullablePinotTaskConfig) Get() *PinotTaskConfig {
	return v.value
}

func (v *NullablePinotTaskConfig) Set(val *PinotTaskConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePinotTaskConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePinotTaskConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePinotTaskConfig(val *PinotTaskConfig) *NullablePinotTaskConfig {
	return &NullablePinotTaskConfig{value: val, isSet: true}
}

func (v NullablePinotTaskConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePinotTaskConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


