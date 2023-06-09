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

// checks if the SuccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuccessResponse{}

// SuccessResponse struct for SuccessResponse
type SuccessResponse struct {
	Status *string `json:"status,omitempty"`
}

// NewSuccessResponse instantiates a new SuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessResponse() *SuccessResponse {
	this := SuccessResponse{}
	return &this
}

// NewSuccessResponseWithDefaults instantiates a new SuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessResponseWithDefaults() *SuccessResponse {
	this := SuccessResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SuccessResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SuccessResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SuccessResponse) SetStatus(v string) {
	o.Status = &v
}

func (o SuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableSuccessResponse struct {
	value *SuccessResponse
	isSet bool
}

func (v NullableSuccessResponse) Get() *SuccessResponse {
	return v.value
}

func (v *NullableSuccessResponse) Set(val *SuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessResponse(val *SuccessResponse) *NullableSuccessResponse {
	return &NullableSuccessResponse{value: val, isSet: true}
}

func (v NullableSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


