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

// checks if the AggregationConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregationConfig{}

// AggregationConfig struct for AggregationConfig
type AggregationConfig struct {
	ColumnName *string `json:"columnName,omitempty"`
	AggregationFunction *string `json:"aggregationFunction,omitempty"`
}

// NewAggregationConfig instantiates a new AggregationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregationConfig() *AggregationConfig {
	this := AggregationConfig{}
	return &this
}

// NewAggregationConfigWithDefaults instantiates a new AggregationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationConfigWithDefaults() *AggregationConfig {
	this := AggregationConfig{}
	return &this
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *AggregationConfig) GetColumnName() string {
	if o == nil || IsNil(o.ColumnName) {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationConfig) GetColumnNameOk() (*string, bool) {
	if o == nil || IsNil(o.ColumnName) {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *AggregationConfig) HasColumnName() bool {
	if o != nil && !IsNil(o.ColumnName) {
		return true
	}

	return false
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *AggregationConfig) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetAggregationFunction returns the AggregationFunction field value if set, zero value otherwise.
func (o *AggregationConfig) GetAggregationFunction() string {
	if o == nil || IsNil(o.AggregationFunction) {
		var ret string
		return ret
	}
	return *o.AggregationFunction
}

// GetAggregationFunctionOk returns a tuple with the AggregationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationConfig) GetAggregationFunctionOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationFunction) {
		return nil, false
	}
	return o.AggregationFunction, true
}

// HasAggregationFunction returns a boolean if a field has been set.
func (o *AggregationConfig) HasAggregationFunction() bool {
	if o != nil && !IsNil(o.AggregationFunction) {
		return true
	}

	return false
}

// SetAggregationFunction gets a reference to the given string and assigns it to the AggregationFunction field.
func (o *AggregationConfig) SetAggregationFunction(v string) {
	o.AggregationFunction = &v
}

func (o AggregationConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregationConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: columnName is readOnly
	// skip: aggregationFunction is readOnly
	return toSerialize, nil
}

type NullableAggregationConfig struct {
	value *AggregationConfig
	isSet bool
}

func (v NullableAggregationConfig) Get() *AggregationConfig {
	return v.value
}

func (v *NullableAggregationConfig) Set(val *AggregationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationConfig(val *AggregationConfig) *NullableAggregationConfig {
	return &NullableAggregationConfig{value: val, isSet: true}
}

func (v NullableAggregationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


