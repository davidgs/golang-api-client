# DimensionFieldSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**TransformFunction** | Pointer to **string** |  | [optional] 
**DefaultNullValue** | Pointer to **map[string]interface{}** |  | [optional] 
**SingleValueField** | Pointer to **bool** |  | [optional] 
**VirtualColumnProvider** | Pointer to **string** |  | [optional] 
**DefaultNullValueString** | Pointer to **string** |  | [optional] 

## Methods

### NewDimensionFieldSpec

`func NewDimensionFieldSpec() *DimensionFieldSpec`

NewDimensionFieldSpec instantiates a new DimensionFieldSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionFieldSpecWithDefaults

`func NewDimensionFieldSpecWithDefaults() *DimensionFieldSpec`

NewDimensionFieldSpecWithDefaults instantiates a new DimensionFieldSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DimensionFieldSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DimensionFieldSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DimensionFieldSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DimensionFieldSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMaxLength

`func (o *DimensionFieldSpec) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *DimensionFieldSpec) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *DimensionFieldSpec) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *DimensionFieldSpec) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetDataType

`func (o *DimensionFieldSpec) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *DimensionFieldSpec) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *DimensionFieldSpec) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *DimensionFieldSpec) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetTransformFunction

`func (o *DimensionFieldSpec) GetTransformFunction() string`

GetTransformFunction returns the TransformFunction field if non-nil, zero value otherwise.

### GetTransformFunctionOk

`func (o *DimensionFieldSpec) GetTransformFunctionOk() (*string, bool)`

GetTransformFunctionOk returns a tuple with the TransformFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformFunction

`func (o *DimensionFieldSpec) SetTransformFunction(v string)`

SetTransformFunction sets TransformFunction field to given value.

### HasTransformFunction

`func (o *DimensionFieldSpec) HasTransformFunction() bool`

HasTransformFunction returns a boolean if a field has been set.

### GetDefaultNullValue

`func (o *DimensionFieldSpec) GetDefaultNullValue() map[string]interface{}`

GetDefaultNullValue returns the DefaultNullValue field if non-nil, zero value otherwise.

### GetDefaultNullValueOk

`func (o *DimensionFieldSpec) GetDefaultNullValueOk() (*map[string]interface{}, bool)`

GetDefaultNullValueOk returns a tuple with the DefaultNullValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNullValue

`func (o *DimensionFieldSpec) SetDefaultNullValue(v map[string]interface{})`

SetDefaultNullValue sets DefaultNullValue field to given value.

### HasDefaultNullValue

`func (o *DimensionFieldSpec) HasDefaultNullValue() bool`

HasDefaultNullValue returns a boolean if a field has been set.

### GetSingleValueField

`func (o *DimensionFieldSpec) GetSingleValueField() bool`

GetSingleValueField returns the SingleValueField field if non-nil, zero value otherwise.

### GetSingleValueFieldOk

`func (o *DimensionFieldSpec) GetSingleValueFieldOk() (*bool, bool)`

GetSingleValueFieldOk returns a tuple with the SingleValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueField

`func (o *DimensionFieldSpec) SetSingleValueField(v bool)`

SetSingleValueField sets SingleValueField field to given value.

### HasSingleValueField

`func (o *DimensionFieldSpec) HasSingleValueField() bool`

HasSingleValueField returns a boolean if a field has been set.

### GetVirtualColumnProvider

`func (o *DimensionFieldSpec) GetVirtualColumnProvider() string`

GetVirtualColumnProvider returns the VirtualColumnProvider field if non-nil, zero value otherwise.

### GetVirtualColumnProviderOk

`func (o *DimensionFieldSpec) GetVirtualColumnProviderOk() (*string, bool)`

GetVirtualColumnProviderOk returns a tuple with the VirtualColumnProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualColumnProvider

`func (o *DimensionFieldSpec) SetVirtualColumnProvider(v string)`

SetVirtualColumnProvider sets VirtualColumnProvider field to given value.

### HasVirtualColumnProvider

`func (o *DimensionFieldSpec) HasVirtualColumnProvider() bool`

HasVirtualColumnProvider returns a boolean if a field has been set.

### GetDefaultNullValueString

`func (o *DimensionFieldSpec) GetDefaultNullValueString() string`

GetDefaultNullValueString returns the DefaultNullValueString field if non-nil, zero value otherwise.

### GetDefaultNullValueStringOk

`func (o *DimensionFieldSpec) GetDefaultNullValueStringOk() (*string, bool)`

GetDefaultNullValueStringOk returns a tuple with the DefaultNullValueString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNullValueString

`func (o *DimensionFieldSpec) SetDefaultNullValueString(v string)`

SetDefaultNullValueString sets DefaultNullValueString field to given value.

### HasDefaultNullValueString

`func (o *DimensionFieldSpec) HasDefaultNullValueString() bool`

HasDefaultNullValueString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


