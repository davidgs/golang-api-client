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

// checks if the AdhocTaskConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdhocTaskConfig{}

// AdhocTaskConfig struct for AdhocTaskConfig
type AdhocTaskConfig struct {
	TaskType string `json:"taskType"`
	TableName string `json:"tableName"`
	TaskName *string `json:"taskName,omitempty"`
	TaskConfigs *map[string]string `json:"taskConfigs,omitempty"`
}

// NewAdhocTaskConfig instantiates a new AdhocTaskConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdhocTaskConfig(taskType string, tableName string) *AdhocTaskConfig {
	this := AdhocTaskConfig{}
	this.TaskType = taskType
	this.TableName = tableName
	return &this
}

// NewAdhocTaskConfigWithDefaults instantiates a new AdhocTaskConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdhocTaskConfigWithDefaults() *AdhocTaskConfig {
	this := AdhocTaskConfig{}
	return &this
}

// GetTaskType returns the TaskType field value
func (o *AdhocTaskConfig) GetTaskType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value
// and a boolean to check if the value has been set.
func (o *AdhocTaskConfig) GetTaskTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskType, true
}

// SetTaskType sets field value
func (o *AdhocTaskConfig) SetTaskType(v string) {
	o.TaskType = v
}

// GetTableName returns the TableName field value
func (o *AdhocTaskConfig) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *AdhocTaskConfig) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value
func (o *AdhocTaskConfig) SetTableName(v string) {
	o.TableName = v
}

// GetTaskName returns the TaskName field value if set, zero value otherwise.
func (o *AdhocTaskConfig) GetTaskName() string {
	if o == nil || IsNil(o.TaskName) {
		var ret string
		return ret
	}
	return *o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdhocTaskConfig) GetTaskNameOk() (*string, bool) {
	if o == nil || IsNil(o.TaskName) {
		return nil, false
	}
	return o.TaskName, true
}

// HasTaskName returns a boolean if a field has been set.
func (o *AdhocTaskConfig) HasTaskName() bool {
	if o != nil && !IsNil(o.TaskName) {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given string and assigns it to the TaskName field.
func (o *AdhocTaskConfig) SetTaskName(v string) {
	o.TaskName = &v
}

// GetTaskConfigs returns the TaskConfigs field value if set, zero value otherwise.
func (o *AdhocTaskConfig) GetTaskConfigs() map[string]string {
	if o == nil || IsNil(o.TaskConfigs) {
		var ret map[string]string
		return ret
	}
	return *o.TaskConfigs
}

// GetTaskConfigsOk returns a tuple with the TaskConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdhocTaskConfig) GetTaskConfigsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.TaskConfigs) {
		return nil, false
	}
	return o.TaskConfigs, true
}

// HasTaskConfigs returns a boolean if a field has been set.
func (o *AdhocTaskConfig) HasTaskConfigs() bool {
	if o != nil && !IsNil(o.TaskConfigs) {
		return true
	}

	return false
}

// SetTaskConfigs gets a reference to the given map[string]string and assigns it to the TaskConfigs field.
func (o *AdhocTaskConfig) SetTaskConfigs(v map[string]string) {
	o.TaskConfigs = &v
}

func (o AdhocTaskConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdhocTaskConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: taskType is readOnly
	// skip: tableName is readOnly
	// skip: taskName is readOnly
	// skip: taskConfigs is readOnly
	return toSerialize, nil
}

type NullableAdhocTaskConfig struct {
	value *AdhocTaskConfig
	isSet bool
}

func (v NullableAdhocTaskConfig) Get() *AdhocTaskConfig {
	return v.value
}

func (v *NullableAdhocTaskConfig) Set(val *AdhocTaskConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAdhocTaskConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAdhocTaskConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdhocTaskConfig(val *AdhocTaskConfig) *NullableAdhocTaskConfig {
	return &NullableAdhocTaskConfig{value: val, isSet: true}
}

func (v NullableAdhocTaskConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdhocTaskConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

