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

// checks if the TenantConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantConfig{}

// TenantConfig struct for TenantConfig
type TenantConfig struct {
	Broker *string `json:"broker,omitempty"`
	Server *string `json:"server,omitempty"`
	TagOverrideConfig *TagOverrideConfig `json:"tagOverrideConfig,omitempty"`
}

// NewTenantConfig instantiates a new TenantConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantConfig() *TenantConfig {
	this := TenantConfig{}
	return &this
}

// NewTenantConfigWithDefaults instantiates a new TenantConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantConfigWithDefaults() *TenantConfig {
	this := TenantConfig{}
	return &this
}

// GetBroker returns the Broker field value if set, zero value otherwise.
func (o *TenantConfig) GetBroker() string {
	if o == nil || IsNil(o.Broker) {
		var ret string
		return ret
	}
	return *o.Broker
}

// GetBrokerOk returns a tuple with the Broker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantConfig) GetBrokerOk() (*string, bool) {
	if o == nil || IsNil(o.Broker) {
		return nil, false
	}
	return o.Broker, true
}

// HasBroker returns a boolean if a field has been set.
func (o *TenantConfig) HasBroker() bool {
	if o != nil && !IsNil(o.Broker) {
		return true
	}

	return false
}

// SetBroker gets a reference to the given string and assigns it to the Broker field.
func (o *TenantConfig) SetBroker(v string) {
	o.Broker = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *TenantConfig) GetServer() string {
	if o == nil || IsNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantConfig) GetServerOk() (*string, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *TenantConfig) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *TenantConfig) SetServer(v string) {
	o.Server = &v
}

// GetTagOverrideConfig returns the TagOverrideConfig field value if set, zero value otherwise.
func (o *TenantConfig) GetTagOverrideConfig() TagOverrideConfig {
	if o == nil || IsNil(o.TagOverrideConfig) {
		var ret TagOverrideConfig
		return ret
	}
	return *o.TagOverrideConfig
}

// GetTagOverrideConfigOk returns a tuple with the TagOverrideConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantConfig) GetTagOverrideConfigOk() (*TagOverrideConfig, bool) {
	if o == nil || IsNil(o.TagOverrideConfig) {
		return nil, false
	}
	return o.TagOverrideConfig, true
}

// HasTagOverrideConfig returns a boolean if a field has been set.
func (o *TenantConfig) HasTagOverrideConfig() bool {
	if o != nil && !IsNil(o.TagOverrideConfig) {
		return true
	}

	return false
}

// SetTagOverrideConfig gets a reference to the given TagOverrideConfig and assigns it to the TagOverrideConfig field.
func (o *TenantConfig) SetTagOverrideConfig(v TagOverrideConfig) {
	o.TagOverrideConfig = &v
}

func (o TenantConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: broker is readOnly
	// skip: server is readOnly
	if !IsNil(o.TagOverrideConfig) {
		toSerialize["tagOverrideConfig"] = o.TagOverrideConfig
	}
	return toSerialize, nil
}

type NullableTenantConfig struct {
	value *TenantConfig
	isSet bool
}

func (v NullableTenantConfig) Get() *TenantConfig {
	return v.value
}

func (v *NullableTenantConfig) Set(val *TenantConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantConfig(val *TenantConfig) *NullableTenantConfig {
	return &NullableTenantConfig{value: val, isSet: true}
}

func (v NullableTenantConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


