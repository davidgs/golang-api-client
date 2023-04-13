# DateTimeFieldSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] 
**SampleValue** | Pointer to **map[string]interface{}** |  | [optional] 
**Granularity** | Pointer to **string** |  | [optional] 
**SingleValueField** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**TransformFunction** | Pointer to **string** |  | [optional] 
**DefaultNullValue** | Pointer to **map[string]interface{}** |  | [optional] 
**VirtualColumnProvider** | Pointer to **string** |  | [optional] 
**DefaultNullValueString** | Pointer to **string** |  | [optional] 

## Methods

### NewDateTimeFieldSpec

`func NewDateTimeFieldSpec() *DateTimeFieldSpec`

NewDateTimeFieldSpec instantiates a new DateTimeFieldSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeFieldSpecWithDefaults

`func NewDateTimeFieldSpecWithDefaults() *DateTimeFieldSpec`

NewDateTimeFieldSpecWithDefaults instantiates a new DateTimeFieldSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *DateTimeFieldSpec) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DateTimeFieldSpec) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DateTimeFieldSpec) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DateTimeFieldSpec) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetSampleValue

`func (o *DateTimeFieldSpec) GetSampleValue() map[string]interface{}`

GetSampleValue returns the SampleValue field if non-nil, zero value otherwise.

### GetSampleValueOk

`func (o *DateTimeFieldSpec) GetSampleValueOk() (*map[string]interface{}, bool)`

GetSampleValueOk returns a tuple with the SampleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleValue

`func (o *DateTimeFieldSpec) SetSampleValue(v map[string]interface{})`

SetSampleValue sets SampleValue field to given value.

### HasSampleValue

`func (o *DateTimeFieldSpec) HasSampleValue() bool`

HasSampleValue returns a boolean if a field has been set.

### GetGranularity

`func (o *DateTimeFieldSpec) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *DateTimeFieldSpec) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *DateTimeFieldSpec) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *DateTimeFieldSpec) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetSingleValueField

`func (o *DateTimeFieldSpec) GetSingleValueField() bool`

GetSingleValueField returns the SingleValueField field if non-nil, zero value otherwise.

### GetSingleValueFieldOk

`func (o *DateTimeFieldSpec) GetSingleValueFieldOk() (*bool, bool)`

GetSingleValueFieldOk returns a tuple with the SingleValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueField

`func (o *DateTimeFieldSpec) SetSingleValueField(v bool)`

SetSingleValueField sets SingleValueField field to given value.

### HasSingleValueField

`func (o *DateTimeFieldSpec) HasSingleValueField() bool`

HasSingleValueField returns a boolean if a field has been set.

### GetName

`func (o *DateTimeFieldSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DateTimeFieldSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DateTimeFieldSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DateTimeFieldSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMaxLength

`func (o *DateTimeFieldSpec) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *DateTimeFieldSpec) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *DateTimeFieldSpec) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *DateTimeFieldSpec) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetDataType

`func (o *DateTimeFieldSpec) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *DateTimeFieldSpec) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *DateTimeFieldSpec) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *DateTimeFieldSpec) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetTransformFunction

`func (o *DateTimeFieldSpec) GetTransformFunction() string`

GetTransformFunction returns the TransformFunction field if non-nil, zero value otherwise.

### GetTransformFunctionOk

`func (o *DateTimeFieldSpec) GetTransformFunctionOk() (*string, bool)`

GetTransformFunctionOk returns a tuple with the TransformFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformFunction

`func (o *DateTimeFieldSpec) SetTransformFunction(v string)`

SetTransformFunction sets TransformFunction field to given value.

### HasTransformFunction

`func (o *DateTimeFieldSpec) HasTransformFunction() bool`

HasTransformFunction returns a boolean if a field has been set.

### GetDefaultNullValue

`func (o *DateTimeFieldSpec) GetDefaultNullValue() map[string]interface{}`

GetDefaultNullValue returns the DefaultNullValue field if non-nil, zero value otherwise.

### GetDefaultNullValueOk

`func (o *DateTimeFieldSpec) GetDefaultNullValueOk() (*map[string]interface{}, bool)`

GetDefaultNullValueOk returns a tuple with the DefaultNullValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNullValue

`func (o *DateTimeFieldSpec) SetDefaultNullValue(v map[string]interface{})`

SetDefaultNullValue sets DefaultNullValue field to given value.

### HasDefaultNullValue

`func (o *DateTimeFieldSpec) HasDefaultNullValue() bool`

HasDefaultNullValue returns a boolean if a field has been set.

### GetVirtualColumnProvider

`func (o *DateTimeFieldSpec) GetVirtualColumnProvider() string`

GetVirtualColumnProvider returns the VirtualColumnProvider field if non-nil, zero value otherwise.

### GetVirtualColumnProviderOk

`func (o *DateTimeFieldSpec) GetVirtualColumnProviderOk() (*string, bool)`

GetVirtualColumnProviderOk returns a tuple with the VirtualColumnProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualColumnProvider

`func (o *DateTimeFieldSpec) SetVirtualColumnProvider(v string)`

SetVirtualColumnProvider sets VirtualColumnProvider field to given value.

### HasVirtualColumnProvider

`func (o *DateTimeFieldSpec) HasVirtualColumnProvider() bool`

HasVirtualColumnProvider returns a boolean if a field has been set.

### GetDefaultNullValueString

`func (o *DateTimeFieldSpec) GetDefaultNullValueString() string`

GetDefaultNullValueString returns the DefaultNullValueString field if non-nil, zero value otherwise.

### GetDefaultNullValueStringOk

`func (o *DateTimeFieldSpec) GetDefaultNullValueStringOk() (*string, bool)`

GetDefaultNullValueStringOk returns a tuple with the DefaultNullValueString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNullValueString

`func (o *DateTimeFieldSpec) SetDefaultNullValueString(v string)`

SetDefaultNullValueString sets DefaultNullValueString field to given value.

### HasDefaultNullValueString

`func (o *DateTimeFieldSpec) HasDefaultNullValueString() bool`

HasDefaultNullValueString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


