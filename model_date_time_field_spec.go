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

// checks if the DateTimeFieldSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateTimeFieldSpec{}

// DateTimeFieldSpec struct for DateTimeFieldSpec
type DateTimeFieldSpec struct {
	Format *string `json:"format,omitempty"`
	SampleValue map[string]interface{} `json:"sampleValue,omitempty"`
	Granularity *string `json:"granularity,omitempty"`
	SingleValueField *bool `json:"singleValueField,omitempty"`
	Name *string `json:"name,omitempty"`
	MaxLength *int32 `json:"maxLength,omitempty"`
	DataType *string `json:"dataType,omitempty"`
	TransformFunction *string `json:"transformFunction,omitempty"`
	DefaultNullValue map[string]interface{} `json:"defaultNullValue,omitempty"`
	VirtualColumnProvider *string `json:"virtualColumnProvider,omitempty"`
	DefaultNullValueString *string `json:"defaultNullValueString,omitempty"`
}

// NewDateTimeFieldSpec instantiates a new DateTimeFieldSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateTimeFieldSpec() *DateTimeFieldSpec {
	this := DateTimeFieldSpec{}
	return &this
}

// NewDateTimeFieldSpecWithDefaults instantiates a new DateTimeFieldSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateTimeFieldSpecWithDefaults() *DateTimeFieldSpec {
	this := DateTimeFieldSpec{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *DateTimeFieldSpec) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldSpec) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *DateTimeFieldSpec) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *DateTimeFieldSpec) SetFormat(v string) {
	o.Format = &v
}

// GetSampleValue returns the SampleValue field value if set, zero value otherwise.
func (o *DateTimeFieldSpec) GetSampleValue() map[string]interface{} {
	if o == nil || IsNil(o.SampleValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.SampleValue
}

// GetSampleValueOk returns a tuple with the SampleValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldSpec) GetSampleValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SampleValue) {
		return map[string]interface{}{}, false
	}
	return o.SampleValue, true
}

// HasSampleValue returns a boolean if a field has been set.
func (o *DateTimeFieldSpec) HasSampleValue() bool {
	if o != nil && !IsNil(o.SampleValue) {
		return true
	}

	return false
}

// SetSampleValue gets a reference to the given map[string]interface{} and assigns it to the SampleValue field.
func (o *DateTimeFieldSpec) SetSampleValue(v map[string]interface{}) {
	o.SampleValue = v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *DateTimeFieldSpec) GetGranularity() string {
	if o == nil || IsNil(o.Granularity) {
		var ret string
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldSpec) GetGranularityOk() (*string, bool) {
	if o == nil || IsNil(o.Granularity) {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *DateTimeFieldSpec) HasGranularity() bool {
	if o != nil && !IsNil(o.Granularity) {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given string and assigns it to the Granularity field.
func (o *DateTimeFieldSpec) SetGranularity(v string) {
	o.Granularity = &v
}

// GetSingleValueField returns the SingleValueField field value if set, zero value otherwise.
func (o *DateTimeFieldSpec) GetSingleValueField() bool {
	if o == nil || IsNil(o.SingleValueField) {
		var ret bool
		return ret
	}
	return *o.SingleValueField
}

// GetSingleValueFieldOk returns a tuple with the SingleValueField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldSpec) GetSingleValueFieldOk() (*bool, bool) {
	if o == nil || IsNil(o.SingleValueField) {
		return nil, false
	}
	return o.SingleValueField, true
}

// HasSingleValueField returns a boolean if a field has been set.
func (o *DateTimeFieldSpec) HasSingleValueField() bool {
	if o != nil && !IsNil(o.SingleValueField) {
		return true
	}

	return false
}

// SetSingleValueField gets a reference to the given bool and assigns it to the SingleValueField field.
func (o *DateTimeFieldSpec) SetSingleValueField(v bool) {
	o.SingleValueField = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DateTimeFieldSpec) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldSpec) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DateTimeFieldSpec) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DateTimeFieldSpec) SetName(v string) {
	o.Name = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *DateTimeFieldSpec) GetMaxLength() int32 {
	if o == nil || IsNil(o.MaxLength) {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldSpec) GetMaxLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLength) {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *DateTimeFieldSpec) HasMaxLength() bool {
	if o != nil && !IsNil(o.MaxLength) {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given int32 and assigns it to the MaxLength field.
func (o *DateTimeFieldSpec) SetMaxLength(v int32) {
	o.MaxLength = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *DateTimeFieldSpec) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldSpec) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *DateTimeFieldSpec) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *DateTimeFieldSpec) SetDataType(v string) {
	o.DataType = &v
}

// GetTransformFunction returns the TransformFunction field value if set, zero value otherwise.
func (o *DateTimeFieldSpec) GetTransformFunction() string {
	if o == nil || IsNil(o.TransformFunction) {
		var ret string
		return ret
	}
	return *o.TransformFunction
}

// GetTransformFunctionOk returns a tuple with the TransformFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldSpec) GetTransformFunctionOk() (*string, bool) {
	if o == nil || IsNil(o.TransformFunction) {
		return nil, false
	}
	return o.TransformFunction, true
}

// HasTransformFunction returns a boolean if a field has been set.
func (o *DateTimeFieldSpec) HasTransformFunction() bool {
	if o != nil && !IsNil(o.TransformFunction) {
		return true
	}

	return false
}

// SetTransformFunction gets a reference to the given string and assigns it to the TransformFunction field.
func (o *DateTimeFieldSpec) SetTransformFunction(v string) {
	o.TransformFunction = &v
}

// GetDefaultNullValue returns the DefaultNullValue field value if set, zero value otherwise.
func (o *DateTimeFieldSpec) GetDefaultNullValue() map[string]interface{} {
	if o == nil || IsNil(o.DefaultNullValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.DefaultNullValue
}

// GetDefaultNullValueOk returns a tuple with the DefaultNullValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldSpec) GetDefaultNullValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DefaultNullValue) {
		return map[string]interface{}{}, false
	}
	return o.DefaultNullValue, true
}

// HasDefaultNullValue returns a boolean if a field has been set.
func (o *DateTimeFieldSpec) HasDefaultNullValue() bool {
	if o != nil && !IsNil(o.DefaultNullValue) {
		return true
	}

	return false
}

// SetDefaultNullValue gets a reference to the given map[string]interface{} and assigns it to the DefaultNullValue field.
func (o *DateTimeFieldSpec) SetDefaultNullValue(v map[string]interface{}) {
	o.DefaultNullValue = v
}

// GetVirtualColumnProvider returns the VirtualColumnProvider field value if set, zero value otherwise.
func (o *DateTimeFieldSpec) GetVirtualColumnProvider() string {
	if o == nil || IsNil(o.VirtualColumnProvider) {
		var ret string
		return ret
	}
	return *o.VirtualColumnProvider
}

// GetVirtualColumnProviderOk returns a tuple with the VirtualColumnProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldSpec) GetVirtualColumnProviderOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualColumnProvider) {
		return nil, false
	}
	return o.VirtualColumnProvider, true
}

// HasVirtualColumnProvider returns a boolean if a field has been set.
func (o *DateTimeFieldSpec) HasVirtualColumnProvider() bool {
	if o != nil && !IsNil(o.VirtualColumnProvider) {
		return true
	}

	return false
}

// SetVirtualColumnProvider gets a reference to the given string and assigns it to the VirtualColumnProvider field.
func (o *DateTimeFieldSpec) SetVirtualColumnProvider(v string) {
	o.VirtualColumnProvider = &v
}

// GetDefaultNullValueString returns the DefaultNullValueString field value if set, zero value otherwise.
func (o *DateTimeFieldSpec) GetDefaultNullValueString() string {
	if o == nil || IsNil(o.DefaultNullValueString) {
		var ret string
		return ret
	}
	return *o.DefaultNullValueString
}

// GetDefaultNullValueStringOk returns a tuple with the DefaultNullValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldSpec) GetDefaultNullValueStringOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultNullValueString) {
		return nil, false
	}
	return o.DefaultNullValueString, true
}

// HasDefaultNullValueString returns a boolean if a field has been set.
func (o *DateTimeFieldSpec) HasDefaultNullValueString() bool {
	if o != nil && !IsNil(o.DefaultNullValueString) {
		return true
	}

	return false
}

// SetDefaultNullValueString gets a reference to the given string and assigns it to the DefaultNullValueString field.
func (o *DateTimeFieldSpec) SetDefaultNullValueString(v string) {
	o.DefaultNullValueString = &v
}

func (o DateTimeFieldSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DateTimeFieldSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.SampleValue) {
		toSerialize["sampleValue"] = o.SampleValue
	}
	if !IsNil(o.Granularity) {
		toSerialize["granularity"] = o.Granularity
	}
	if !IsNil(o.SingleValueField) {
		toSerialize["singleValueField"] = o.SingleValueField
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.MaxLength) {
		toSerialize["maxLength"] = o.MaxLength
	}
	if !IsNil(o.DataType) {
		toSerialize["dataType"] = o.DataType
	}
	if !IsNil(o.TransformFunction) {
		toSerialize["transformFunction"] = o.TransformFunction
	}
	if !IsNil(o.DefaultNullValue) {
		toSerialize["defaultNullValue"] = o.DefaultNullValue
	}
	if !IsNil(o.VirtualColumnProvider) {
		toSerialize["virtualColumnProvider"] = o.VirtualColumnProvider
	}
	if !IsNil(o.DefaultNullValueString) {
		toSerialize["defaultNullValueString"] = o.DefaultNullValueString
	}
	return toSerialize, nil
}

type NullableDateTimeFieldSpec struct {
	value *DateTimeFieldSpec
	isSet bool
}

func (v NullableDateTimeFieldSpec) Get() *DateTimeFieldSpec {
	return v.value
}

func (v *NullableDateTimeFieldSpec) Set(val *DateTimeFieldSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDateTimeFieldSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDateTimeFieldSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateTimeFieldSpec(val *DateTimeFieldSpec) *NullableDateTimeFieldSpec {
	return &NullableDateTimeFieldSpec{value: val, isSet: true}
}

func (v NullableDateTimeFieldSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateTimeFieldSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


