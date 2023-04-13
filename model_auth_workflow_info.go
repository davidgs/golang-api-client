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

// checks if the AuthWorkflowInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthWorkflowInfo{}

// AuthWorkflowInfo struct for AuthWorkflowInfo
type AuthWorkflowInfo struct {
	Workflow *string `json:"workflow,omitempty"`
}

// NewAuthWorkflowInfo instantiates a new AuthWorkflowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthWorkflowInfo() *AuthWorkflowInfo {
	this := AuthWorkflowInfo{}
	return &this
}

// NewAuthWorkflowInfoWithDefaults instantiates a new AuthWorkflowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthWorkflowInfoWithDefaults() *AuthWorkflowInfo {
	this := AuthWorkflowInfo{}
	return &this
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *AuthWorkflowInfo) GetWorkflow() string {
	if o == nil || IsNil(o.Workflow) {
		var ret string
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthWorkflowInfo) GetWorkflowOk() (*string, bool) {
	if o == nil || IsNil(o.Workflow) {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *AuthWorkflowInfo) HasWorkflow() bool {
	if o != nil && !IsNil(o.Workflow) {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given string and assigns it to the Workflow field.
func (o *AuthWorkflowInfo) SetWorkflow(v string) {
	o.Workflow = &v
}

func (o AuthWorkflowInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthWorkflowInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Workflow) {
		toSerialize["workflow"] = o.Workflow
	}
	return toSerialize, nil
}

type NullableAuthWorkflowInfo struct {
	value *AuthWorkflowInfo
	isSet bool
}

func (v NullableAuthWorkflowInfo) Get() *AuthWorkflowInfo {
	return v.value
}

func (v *NullableAuthWorkflowInfo) Set(val *AuthWorkflowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthWorkflowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthWorkflowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthWorkflowInfo(val *AuthWorkflowInfo) *NullableAuthWorkflowInfo {
	return &NullableAuthWorkflowInfo{value: val, isSet: true}
}

func (v NullableAuthWorkflowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthWorkflowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

