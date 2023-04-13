# TransformConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnName** | Pointer to **string** |  | [optional] [readonly] 
**TransformFunction** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTransformConfig

`func NewTransformConfig() *TransformConfig`

NewTransformConfig instantiates a new TransformConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformConfigWithDefaults

`func NewTransformConfigWithDefaults() *TransformConfig`

NewTransformConfigWithDefaults instantiates a new TransformConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnName

`func (o *TransformConfig) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *TransformConfig) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *TransformConfig) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *TransformConfig) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### GetTransformFunction

`func (o *TransformConfig) GetTransformFunction() string`

GetTransformFunction returns the TransformFunction field if non-nil, zero value otherwise.

### GetTransformFunctionOk

`func (o *TransformConfig) GetTransformFunctionOk() (*string, bool)`

GetTransformFunctionOk returns a tuple with the TransformFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformFunction

`func (o *TransformConfig) SetTransformFunction(v string)`

SetTransformFunction sets TransformFunction field to given value.

### HasTransformFunction

`func (o *TransformConfig) HasTransformFunction() bool`

HasTransformFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


