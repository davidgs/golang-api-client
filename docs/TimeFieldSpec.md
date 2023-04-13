# TimeFieldSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**IncomingGranularitySpec** | Pointer to [**TimeGranularitySpec**](TimeGranularitySpec.md) |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**OutgoingGranularitySpec** | Pointer to [**TimeGranularitySpec**](TimeGranularitySpec.md) |  | [optional] 
**SingleValueField** | Pointer to **bool** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 
**TransformFunction** | Pointer to **string** |  | [optional] 
**DefaultNullValue** | Pointer to **map[string]interface{}** |  | [optional] 
**VirtualColumnProvider** | Pointer to **string** |  | [optional] 
**DefaultNullValueString** | Pointer to **string** |  | [optional] 

## Methods

### NewTimeFieldSpec

`func NewTimeFieldSpec() *TimeFieldSpec`

NewTimeFieldSpec instantiates a new TimeFieldSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeFieldSpecWithDefaults

`func NewTimeFieldSpecWithDefaults() *TimeFieldSpec`

NewTimeFieldSpecWithDefaults instantiates a new TimeFieldSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TimeFieldSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimeFieldSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimeFieldSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimeFieldSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIncomingGranularitySpec

`func (o *TimeFieldSpec) GetIncomingGranularitySpec() TimeGranularitySpec`

GetIncomingGranularitySpec returns the IncomingGranularitySpec field if non-nil, zero value otherwise.

### GetIncomingGranularitySpecOk

`func (o *TimeFieldSpec) GetIncomingGranularitySpecOk() (*TimeGranularitySpec, bool)`

GetIncomingGranularitySpecOk returns a tuple with the IncomingGranularitySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingGranularitySpec

`func (o *TimeFieldSpec) SetIncomingGranularitySpec(v TimeGranularitySpec)`

SetIncomingGranularitySpec sets IncomingGranularitySpec field to given value.

### HasIncomingGranularitySpec

`func (o *TimeFieldSpec) HasIncomingGranularitySpec() bool`

HasIncomingGranularitySpec returns a boolean if a field has been set.

### GetDataType

`func (o *TimeFieldSpec) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *TimeFieldSpec) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *TimeFieldSpec) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *TimeFieldSpec) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetOutgoingGranularitySpec

`func (o *TimeFieldSpec) GetOutgoingGranularitySpec() TimeGranularitySpec`

GetOutgoingGranularitySpec returns the OutgoingGranularitySpec field if non-nil, zero value otherwise.

### GetOutgoingGranularitySpecOk

`func (o *TimeFieldSpec) GetOutgoingGranularitySpecOk() (*TimeGranularitySpec, bool)`

GetOutgoingGranularitySpecOk returns a tuple with the OutgoingGranularitySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingGranularitySpec

`func (o *TimeFieldSpec) SetOutgoingGranularitySpec(v TimeGranularitySpec)`

SetOutgoingGranularitySpec sets OutgoingGranularitySpec field to given value.

### HasOutgoingGranularitySpec

`func (o *TimeFieldSpec) HasOutgoingGranularitySpec() bool`

HasOutgoingGranularitySpec returns a boolean if a field has been set.

### GetSingleValueField

`func (o *TimeFieldSpec) GetSingleValueField() bool`

GetSingleValueField returns the SingleValueField field if non-nil, zero value otherwise.

### GetSingleValueFieldOk

`func (o *TimeFieldSpec) GetSingleValueFieldOk() (*bool, bool)`

GetSingleValueFieldOk returns a tuple with the SingleValueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueField

`func (o *TimeFieldSpec) SetSingleValueField(v bool)`

SetSingleValueField sets SingleValueField field to given value.

### HasSingleValueField

`func (o *TimeFieldSpec) HasSingleValueField() bool`

HasSingleValueField returns a boolean if a field has been set.

### GetMaxLength

`func (o *TimeFieldSpec) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *TimeFieldSpec) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *TimeFieldSpec) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *TimeFieldSpec) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetTransformFunction

`func (o *TimeFieldSpec) GetTransformFunction() string`

GetTransformFunction returns the TransformFunction field if non-nil, zero value otherwise.

### GetTransformFunctionOk

`func (o *TimeFieldSpec) GetTransformFunctionOk() (*string, bool)`

GetTransformFunctionOk returns a tuple with the TransformFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformFunction

`func (o *TimeFieldSpec) SetTransformFunction(v string)`

SetTransformFunction sets TransformFunction field to given value.

### HasTransformFunction

`func (o *TimeFieldSpec) HasTransformFunction() bool`

HasTransformFunction returns a boolean if a field has been set.

### GetDefaultNullValue

`func (o *TimeFieldSpec) GetDefaultNullValue() map[string]interface{}`

GetDefaultNullValue returns the DefaultNullValue field if non-nil, zero value otherwise.

### GetDefaultNullValueOk

`func (o *TimeFieldSpec) GetDefaultNullValueOk() (*map[string]interface{}, bool)`

GetDefaultNullValueOk returns a tuple with the DefaultNullValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNullValue

`func (o *TimeFieldSpec) SetDefaultNullValue(v map[string]interface{})`

SetDefaultNullValue sets DefaultNullValue field to given value.

### HasDefaultNullValue

`func (o *TimeFieldSpec) HasDefaultNullValue() bool`

HasDefaultNullValue returns a boolean if a field has been set.

### GetVirtualColumnProvider

`func (o *TimeFieldSpec) GetVirtualColumnProvider() string`

GetVirtualColumnProvider returns the VirtualColumnProvider field if non-nil, zero value otherwise.

### GetVirtualColumnProviderOk

`func (o *TimeFieldSpec) GetVirtualColumnProviderOk() (*string, bool)`

GetVirtualColumnProviderOk returns a tuple with the VirtualColumnProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualColumnProvider

`func (o *TimeFieldSpec) SetVirtualColumnProvider(v string)`

SetVirtualColumnProvider sets VirtualColumnProvider field to given value.

### HasVirtualColumnProvider

`func (o *TimeFieldSpec) HasVirtualColumnProvider() bool`

HasVirtualColumnProvider returns a boolean if a field has been set.

### GetDefaultNullValueString

`func (o *TimeFieldSpec) GetDefaultNullValueString() string`

GetDefaultNullValueString returns the DefaultNullValueString field if non-nil, zero value otherwise.

### GetDefaultNullValueStringOk

`func (o *TimeFieldSpec) GetDefaultNullValueStringOk() (*string, bool)`

GetDefaultNullValueStringOk returns a tuple with the DefaultNullValueString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNullValueString

`func (o *TimeFieldSpec) SetDefaultNullValueString(v string)`

SetDefaultNullValueString sets DefaultNullValueString field to given value.

### HasDefaultNullValueString

`func (o *TimeFieldSpec) HasDefaultNullValueString() bool`

HasDefaultNullValueString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


