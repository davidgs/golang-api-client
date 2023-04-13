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

// checks if the UpsertConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertConfig{}

// UpsertConfig struct for UpsertConfig
type UpsertConfig struct {
	Mode *string `json:"mode,omitempty"`
	ComparisonColumns []string `json:"comparisonColumns,omitempty"`
	HashFunction *string `json:"hashFunction,omitempty"`
	PartialUpsertStrategies *map[string]string `json:"partialUpsertStrategies,omitempty"`
	DefaultPartialUpsertStrategy *string `json:"defaultPartialUpsertStrategy,omitempty"`
	EnableSnapshot *bool `json:"enableSnapshot,omitempty"`
	MetadataManagerClass *string `json:"metadataManagerClass,omitempty"`
	MetadataManagerConfigs *map[string]string `json:"metadataManagerConfigs,omitempty"`
}

// NewUpsertConfig instantiates a new UpsertConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertConfig() *UpsertConfig {
	this := UpsertConfig{}
	return &this
}

// NewUpsertConfigWithDefaults instantiates a new UpsertConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertConfigWithDefaults() *UpsertConfig {
	this := UpsertConfig{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *UpsertConfig) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertConfig) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *UpsertConfig) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *UpsertConfig) SetMode(v string) {
	o.Mode = &v
}

// GetComparisonColumns returns the ComparisonColumns field value if set, zero value otherwise.
func (o *UpsertConfig) GetComparisonColumns() []string {
	if o == nil || IsNil(o.ComparisonColumns) {
		var ret []string
		return ret
	}
	return o.ComparisonColumns
}

// GetComparisonColumnsOk returns a tuple with the ComparisonColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertConfig) GetComparisonColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.ComparisonColumns) {
		return nil, false
	}
	return o.ComparisonColumns, true
}

// HasComparisonColumns returns a boolean if a field has been set.
func (o *UpsertConfig) HasComparisonColumns() bool {
	if o != nil && !IsNil(o.ComparisonColumns) {
		return true
	}

	return false
}

// SetComparisonColumns gets a reference to the given []string and assigns it to the ComparisonColumns field.
func (o *UpsertConfig) SetComparisonColumns(v []string) {
	o.ComparisonColumns = v
}

// GetHashFunction returns the HashFunction field value if set, zero value otherwise.
func (o *UpsertConfig) GetHashFunction() string {
	if o == nil || IsNil(o.HashFunction) {
		var ret string
		return ret
	}
	return *o.HashFunction
}

// GetHashFunctionOk returns a tuple with the HashFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertConfig) GetHashFunctionOk() (*string, bool) {
	if o == nil || IsNil(o.HashFunction) {
		return nil, false
	}
	return o.HashFunction, true
}

// HasHashFunction returns a boolean if a field has been set.
func (o *UpsertConfig) HasHashFunction() bool {
	if o != nil && !IsNil(o.HashFunction) {
		return true
	}

	return false
}

// SetHashFunction gets a reference to the given string and assigns it to the HashFunction field.
func (o *UpsertConfig) SetHashFunction(v string) {
	o.HashFunction = &v
}

// GetPartialUpsertStrategies returns the PartialUpsertStrategies field value if set, zero value otherwise.
func (o *UpsertConfig) GetPartialUpsertStrategies() map[string]string {
	if o == nil || IsNil(o.PartialUpsertStrategies) {
		var ret map[string]string
		return ret
	}
	return *o.PartialUpsertStrategies
}

// GetPartialUpsertStrategiesOk returns a tuple with the PartialUpsertStrategies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertConfig) GetPartialUpsertStrategiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PartialUpsertStrategies) {
		return nil, false
	}
	return o.PartialUpsertStrategies, true
}

// HasPartialUpsertStrategies returns a boolean if a field has been set.
func (o *UpsertConfig) HasPartialUpsertStrategies() bool {
	if o != nil && !IsNil(o.PartialUpsertStrategies) {
		return true
	}

	return false
}

// SetPartialUpsertStrategies gets a reference to the given map[string]string and assigns it to the PartialUpsertStrategies field.
func (o *UpsertConfig) SetPartialUpsertStrategies(v map[string]string) {
	o.PartialUpsertStrategies = &v
}

// GetDefaultPartialUpsertStrategy returns the DefaultPartialUpsertStrategy field value if set, zero value otherwise.
func (o *UpsertConfig) GetDefaultPartialUpsertStrategy() string {
	if o == nil || IsNil(o.DefaultPartialUpsertStrategy) {
		var ret string
		return ret
	}
	return *o.DefaultPartialUpsertStrategy
}

// GetDefaultPartialUpsertStrategyOk returns a tuple with the DefaultPartialUpsertStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertConfig) GetDefaultPartialUpsertStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPartialUpsertStrategy) {
		return nil, false
	}
	return o.DefaultPartialUpsertStrategy, true
}

// HasDefaultPartialUpsertStrategy returns a boolean if a field has been set.
func (o *UpsertConfig) HasDefaultPartialUpsertStrategy() bool {
	if o != nil && !IsNil(o.DefaultPartialUpsertStrategy) {
		return true
	}

	return false
}

// SetDefaultPartialUpsertStrategy gets a reference to the given string and assigns it to the DefaultPartialUpsertStrategy field.
func (o *UpsertConfig) SetDefaultPartialUpsertStrategy(v string) {
	o.DefaultPartialUpsertStrategy = &v
}

// GetEnableSnapshot returns the EnableSnapshot field value if set, zero value otherwise.
func (o *UpsertConfig) GetEnableSnapshot() bool {
	if o == nil || IsNil(o.EnableSnapshot) {
		var ret bool
		return ret
	}
	return *o.EnableSnapshot
}

// GetEnableSnapshotOk returns a tuple with the EnableSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertConfig) GetEnableSnapshotOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSnapshot) {
		return nil, false
	}
	return o.EnableSnapshot, true
}

// HasEnableSnapshot returns a boolean if a field has been set.
func (o *UpsertConfig) HasEnableSnapshot() bool {
	if o != nil && !IsNil(o.EnableSnapshot) {
		return true
	}

	return false
}

// SetEnableSnapshot gets a reference to the given bool and assigns it to the EnableSnapshot field.
func (o *UpsertConfig) SetEnableSnapshot(v bool) {
	o.EnableSnapshot = &v
}

// GetMetadataManagerClass returns the MetadataManagerClass field value if set, zero value otherwise.
func (o *UpsertConfig) GetMetadataManagerClass() string {
	if o == nil || IsNil(o.MetadataManagerClass) {
		var ret string
		return ret
	}
	return *o.MetadataManagerClass
}

// GetMetadataManagerClassOk returns a tuple with the MetadataManagerClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertConfig) GetMetadataManagerClassOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataManagerClass) {
		return nil, false
	}
	return o.MetadataManagerClass, true
}

// HasMetadataManagerClass returns a boolean if a field has been set.
func (o *UpsertConfig) HasMetadataManagerClass() bool {
	if o != nil && !IsNil(o.MetadataManagerClass) {
		return true
	}

	return false
}

// SetMetadataManagerClass gets a reference to the given string and assigns it to the MetadataManagerClass field.
func (o *UpsertConfig) SetMetadataManagerClass(v string) {
	o.MetadataManagerClass = &v
}

// GetMetadataManagerConfigs returns the MetadataManagerConfigs field value if set, zero value otherwise.
func (o *UpsertConfig) GetMetadataManagerConfigs() map[string]string {
	if o == nil || IsNil(o.MetadataManagerConfigs) {
		var ret map[string]string
		return ret
	}
	return *o.MetadataManagerConfigs
}

// GetMetadataManagerConfigsOk returns a tuple with the MetadataManagerConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertConfig) GetMetadataManagerConfigsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.MetadataManagerConfigs) {
		return nil, false
	}
	return o.MetadataManagerConfigs, true
}

// HasMetadataManagerConfigs returns a boolean if a field has been set.
func (o *UpsertConfig) HasMetadataManagerConfigs() bool {
	if o != nil && !IsNil(o.MetadataManagerConfigs) {
		return true
	}

	return false
}

// SetMetadataManagerConfigs gets a reference to the given map[string]string and assigns it to the MetadataManagerConfigs field.
func (o *UpsertConfig) SetMetadataManagerConfigs(v map[string]string) {
	o.MetadataManagerConfigs = &v
}

func (o UpsertConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.ComparisonColumns) {
		toSerialize["comparisonColumns"] = o.ComparisonColumns
	}
	if !IsNil(o.HashFunction) {
		toSerialize["hashFunction"] = o.HashFunction
	}
	if !IsNil(o.PartialUpsertStrategies) {
		toSerialize["partialUpsertStrategies"] = o.PartialUpsertStrategies
	}
	if !IsNil(o.DefaultPartialUpsertStrategy) {
		toSerialize["defaultPartialUpsertStrategy"] = o.DefaultPartialUpsertStrategy
	}
	if !IsNil(o.EnableSnapshot) {
		toSerialize["enableSnapshot"] = o.EnableSnapshot
	}
	if !IsNil(o.MetadataManagerClass) {
		toSerialize["metadataManagerClass"] = o.MetadataManagerClass
	}
	if !IsNil(o.MetadataManagerConfigs) {
		toSerialize["metadataManagerConfigs"] = o.MetadataManagerConfigs
	}
	return toSerialize, nil
}

type NullableUpsertConfig struct {
	value *UpsertConfig
	isSet bool
}

func (v NullableUpsertConfig) Get() *UpsertConfig {
	return v.value
}

func (v *NullableUpsertConfig) Set(val *UpsertConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertConfig(val *UpsertConfig) *NullableUpsertConfig {
	return &NullableUpsertConfig{value: val, isSet: true}
}

func (v NullableUpsertConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


