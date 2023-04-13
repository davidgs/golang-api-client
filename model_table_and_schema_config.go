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

// checks if the TableAndSchemaConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TableAndSchemaConfig{}

// TableAndSchemaConfig struct for TableAndSchemaConfig
type TableAndSchemaConfig struct {
	TableConfig TableConfig `json:"tableConfig"`
	Schema *Schema `json:"schema,omitempty"`
}

// NewTableAndSchemaConfig instantiates a new TableAndSchemaConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableAndSchemaConfig(tableConfig TableConfig) *TableAndSchemaConfig {
	this := TableAndSchemaConfig{}
	this.TableConfig = tableConfig
	return &this
}

// NewTableAndSchemaConfigWithDefaults instantiates a new TableAndSchemaConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableAndSchemaConfigWithDefaults() *TableAndSchemaConfig {
	this := TableAndSchemaConfig{}
	return &this
}

// GetTableConfig returns the TableConfig field value
func (o *TableAndSchemaConfig) GetTableConfig() TableConfig {
	if o == nil {
		var ret TableConfig
		return ret
	}

	return o.TableConfig
}

// GetTableConfigOk returns a tuple with the TableConfig field value
// and a boolean to check if the value has been set.
func (o *TableAndSchemaConfig) GetTableConfigOk() (*TableConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableConfig, true
}

// SetTableConfig sets field value
func (o *TableAndSchemaConfig) SetTableConfig(v TableConfig) {
	o.TableConfig = v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *TableAndSchemaConfig) GetSchema() Schema {
	if o == nil || IsNil(o.Schema) {
		var ret Schema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableAndSchemaConfig) GetSchemaOk() (*Schema, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *TableAndSchemaConfig) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given Schema and assigns it to the Schema field.
func (o *TableAndSchemaConfig) SetSchema(v Schema) {
	o.Schema = &v
}

func (o TableAndSchemaConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TableAndSchemaConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tableConfig"] = o.TableConfig
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	return toSerialize, nil
}

type NullableTableAndSchemaConfig struct {
	value *TableAndSchemaConfig
	isSet bool
}

func (v NullableTableAndSchemaConfig) Get() *TableAndSchemaConfig {
	return v.value
}

func (v *NullableTableAndSchemaConfig) Set(val *TableAndSchemaConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTableAndSchemaConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTableAndSchemaConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableAndSchemaConfig(val *TableAndSchemaConfig) *NullableTableAndSchemaConfig {
	return &NullableTableAndSchemaConfig{value: val, isSet: true}
}

func (v NullableTableAndSchemaConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableAndSchemaConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

