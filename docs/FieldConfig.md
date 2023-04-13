# FieldConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [readonly] 
**EncodingType** | Pointer to **string** |  | [optional] [readonly] 
**IndexType** | Pointer to **string** |  | [optional] [readonly] 
**IndexTypes** | Pointer to **[]string** |  | [optional] [readonly] 
**CompressionCodec** | Pointer to **string** |  | [optional] [readonly] 
**TimestampConfig** | Pointer to [**TimestampConfig**](TimestampConfig.md) |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] [readonly] 

## Methods

### NewFieldConfig

`func NewFieldConfig(name string, ) *FieldConfig`

NewFieldConfig instantiates a new FieldConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldConfigWithDefaults

`func NewFieldConfigWithDefaults() *FieldConfig`

NewFieldConfigWithDefaults instantiates a new FieldConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FieldConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldConfig) SetName(v string)`

SetName sets Name field to given value.


### GetEncodingType

`func (o *FieldConfig) GetEncodingType() string`

GetEncodingType returns the EncodingType field if non-nil, zero value otherwise.

### GetEncodingTypeOk

`func (o *FieldConfig) GetEncodingTypeOk() (*string, bool)`

GetEncodingTypeOk returns a tuple with the EncodingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingType

`func (o *FieldConfig) SetEncodingType(v string)`

SetEncodingType sets EncodingType field to given value.

### HasEncodingType

`func (o *FieldConfig) HasEncodingType() bool`

HasEncodingType returns a boolean if a field has been set.

### GetIndexType

`func (o *FieldConfig) GetIndexType() string`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *FieldConfig) GetIndexTypeOk() (*string, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *FieldConfig) SetIndexType(v string)`

SetIndexType sets IndexType field to given value.

### HasIndexType

`func (o *FieldConfig) HasIndexType() bool`

HasIndexType returns a boolean if a field has been set.

### GetIndexTypes

`func (o *FieldConfig) GetIndexTypes() []string`

GetIndexTypes returns the IndexTypes field if non-nil, zero value otherwise.

### GetIndexTypesOk

`func (o *FieldConfig) GetIndexTypesOk() (*[]string, bool)`

GetIndexTypesOk returns a tuple with the IndexTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexTypes

`func (o *FieldConfig) SetIndexTypes(v []string)`

SetIndexTypes sets IndexTypes field to given value.

### HasIndexTypes

`func (o *FieldConfig) HasIndexTypes() bool`

HasIndexTypes returns a boolean if a field has been set.

### GetCompressionCodec

`func (o *FieldConfig) GetCompressionCodec() string`

GetCompressionCodec returns the CompressionCodec field if non-nil, zero value otherwise.

### GetCompressionCodecOk

`func (o *FieldConfig) GetCompressionCodecOk() (*string, bool)`

GetCompressionCodecOk returns a tuple with the CompressionCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionCodec

`func (o *FieldConfig) SetCompressionCodec(v string)`

SetCompressionCodec sets CompressionCodec field to given value.

### HasCompressionCodec

`func (o *FieldConfig) HasCompressionCodec() bool`

HasCompressionCodec returns a boolean if a field has been set.

### GetTimestampConfig

`func (o *FieldConfig) GetTimestampConfig() TimestampConfig`

GetTimestampConfig returns the TimestampConfig field if non-nil, zero value otherwise.

### GetTimestampConfigOk

`func (o *FieldConfig) GetTimestampConfigOk() (*TimestampConfig, bool)`

GetTimestampConfigOk returns a tuple with the TimestampConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampConfig

`func (o *FieldConfig) SetTimestampConfig(v TimestampConfig)`

SetTimestampConfig sets TimestampConfig field to given value.

### HasTimestampConfig

`func (o *FieldConfig) HasTimestampConfig() bool`

HasTimestampConfig returns a boolean if a field has been set.

### GetProperties

`func (o *FieldConfig) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FieldConfig) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FieldConfig) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *FieldConfig) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


