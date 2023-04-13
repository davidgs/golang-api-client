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

// checks if the SegmentSizeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentSizeInfo{}

// SegmentSizeInfo struct for SegmentSizeInfo
type SegmentSizeInfo struct {
	SegmentName *string `json:"segmentName,omitempty"`
	DiskSizeInBytes *int64 `json:"diskSizeInBytes,omitempty"`
}

// NewSegmentSizeInfo instantiates a new SegmentSizeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentSizeInfo() *SegmentSizeInfo {
	this := SegmentSizeInfo{}
	return &this
}

// NewSegmentSizeInfoWithDefaults instantiates a new SegmentSizeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentSizeInfoWithDefaults() *SegmentSizeInfo {
	this := SegmentSizeInfo{}
	return &this
}

// GetSegmentName returns the SegmentName field value if set, zero value otherwise.
func (o *SegmentSizeInfo) GetSegmentName() string {
	if o == nil || IsNil(o.SegmentName) {
		var ret string
		return ret
	}
	return *o.SegmentName
}

// GetSegmentNameOk returns a tuple with the SegmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentSizeInfo) GetSegmentNameOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentName) {
		return nil, false
	}
	return o.SegmentName, true
}

// HasSegmentName returns a boolean if a field has been set.
func (o *SegmentSizeInfo) HasSegmentName() bool {
	if o != nil && !IsNil(o.SegmentName) {
		return true
	}

	return false
}

// SetSegmentName gets a reference to the given string and assigns it to the SegmentName field.
func (o *SegmentSizeInfo) SetSegmentName(v string) {
	o.SegmentName = &v
}

// GetDiskSizeInBytes returns the DiskSizeInBytes field value if set, zero value otherwise.
func (o *SegmentSizeInfo) GetDiskSizeInBytes() int64 {
	if o == nil || IsNil(o.DiskSizeInBytes) {
		var ret int64
		return ret
	}
	return *o.DiskSizeInBytes
}

// GetDiskSizeInBytesOk returns a tuple with the DiskSizeInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentSizeInfo) GetDiskSizeInBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DiskSizeInBytes) {
		return nil, false
	}
	return o.DiskSizeInBytes, true
}

// HasDiskSizeInBytes returns a boolean if a field has been set.
func (o *SegmentSizeInfo) HasDiskSizeInBytes() bool {
	if o != nil && !IsNil(o.DiskSizeInBytes) {
		return true
	}

	return false
}

// SetDiskSizeInBytes gets a reference to the given int64 and assigns it to the DiskSizeInBytes field.
func (o *SegmentSizeInfo) SetDiskSizeInBytes(v int64) {
	o.DiskSizeInBytes = &v
}

func (o SegmentSizeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentSizeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: segmentName is readOnly
	// skip: diskSizeInBytes is readOnly
	return toSerialize, nil
}

type NullableSegmentSizeInfo struct {
	value *SegmentSizeInfo
	isSet bool
}

func (v NullableSegmentSizeInfo) Get() *SegmentSizeInfo {
	return v.value
}

func (v *NullableSegmentSizeInfo) Set(val *SegmentSizeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentSizeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentSizeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentSizeInfo(val *SegmentSizeInfo) *NullableSegmentSizeInfo {
	return &NullableSegmentSizeInfo{value: val, isSet: true}
}

func (v NullableSegmentSizeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentSizeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

