# ComplexTypeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldsToUnnest** | Pointer to **[]string** |  | [optional] [readonly] 
**Delimiter** | Pointer to **string** |  | [optional] [readonly] 
**CollectionNotUnnestedToJson** | Pointer to **string** |  | [optional] [readonly] 
**PrefixesToRename** | Pointer to **map[string]string** |  | [optional] [readonly] 

## Methods

### NewComplexTypeConfig

`func NewComplexTypeConfig() *ComplexTypeConfig`

NewComplexTypeConfig instantiates a new ComplexTypeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplexTypeConfigWithDefaults

`func NewComplexTypeConfigWithDefaults() *ComplexTypeConfig`

NewComplexTypeConfigWithDefaults instantiates a new ComplexTypeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldsToUnnest

`func (o *ComplexTypeConfig) GetFieldsToUnnest() []string`

GetFieldsToUnnest returns the FieldsToUnnest field if non-nil, zero value otherwise.

### GetFieldsToUnnestOk

`func (o *ComplexTypeConfig) GetFieldsToUnnestOk() (*[]string, bool)`

GetFieldsToUnnestOk returns a tuple with the FieldsToUnnest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsToUnnest

`func (o *ComplexTypeConfig) SetFieldsToUnnest(v []string)`

SetFieldsToUnnest sets FieldsToUnnest field to given value.

### HasFieldsToUnnest

`func (o *ComplexTypeConfig) HasFieldsToUnnest() bool`

HasFieldsToUnnest returns a boolean if a field has been set.

### GetDelimiter

`func (o *ComplexTypeConfig) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *ComplexTypeConfig) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *ComplexTypeConfig) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *ComplexTypeConfig) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetCollectionNotUnnestedToJson

`func (o *ComplexTypeConfig) GetCollectionNotUnnestedToJson() string`

GetCollectionNotUnnestedToJson returns the CollectionNotUnnestedToJson field if non-nil, zero value otherwise.

### GetCollectionNotUnnestedToJsonOk

`func (o *ComplexTypeConfig) GetCollectionNotUnnestedToJsonOk() (*string, bool)`

GetCollectionNotUnnestedToJsonOk returns a tuple with the CollectionNotUnnestedToJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionNotUnnestedToJson

`func (o *ComplexTypeConfig) SetCollectionNotUnnestedToJson(v string)`

SetCollectionNotUnnestedToJson sets CollectionNotUnnestedToJson field to given value.

### HasCollectionNotUnnestedToJson

`func (o *ComplexTypeConfig) HasCollectionNotUnnestedToJson() bool`

HasCollectionNotUnnestedToJson returns a boolean if a field has been set.

### GetPrefixesToRename

`func (o *ComplexTypeConfig) GetPrefixesToRename() map[string]string`

GetPrefixesToRename returns the PrefixesToRename field if non-nil, zero value otherwise.

### GetPrefixesToRenameOk

`func (o *ComplexTypeConfig) GetPrefixesToRenameOk() (*map[string]string, bool)`

GetPrefixesToRenameOk returns a tuple with the PrefixesToRename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixesToRename

`func (o *ComplexTypeConfig) SetPrefixesToRename(v map[string]string)`

SetPrefixesToRename sets PrefixesToRename field to given value.

### HasPrefixesToRename

`func (o *ComplexTypeConfig) HasPrefixesToRename() bool`

HasPrefixesToRename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


