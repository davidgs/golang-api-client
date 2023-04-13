# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaName** | Pointer to **string** |  | [optional] 
**PrimaryKeyColumns** | Pointer to **[]string** |  | [optional] 
**DimensionFieldSpecs** | Pointer to [**[]DimensionFieldSpec**](DimensionFieldSpec.md) |  | [optional] 
**MetricFieldSpecs** | Pointer to [**[]MetricFieldSpec**](MetricFieldSpec.md) |  | [optional] 
**DateTimeFieldSpecs** | Pointer to [**[]DateTimeFieldSpec**](DateTimeFieldSpec.md) |  | [optional] 
**TimeFieldSpec** | Pointer to [**TimeFieldSpec**](TimeFieldSpec.md) |  | [optional] 

## Methods

### NewSchema

`func NewSchema() *Schema`

NewSchema instantiates a new Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithDefaults

`func NewSchemaWithDefaults() *Schema`

NewSchemaWithDefaults instantiates a new Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaName

`func (o *Schema) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *Schema) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *Schema) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *Schema) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### GetPrimaryKeyColumns

`func (o *Schema) GetPrimaryKeyColumns() []string`

GetPrimaryKeyColumns returns the PrimaryKeyColumns field if non-nil, zero value otherwise.

### GetPrimaryKeyColumnsOk

`func (o *Schema) GetPrimaryKeyColumnsOk() (*[]string, bool)`

GetPrimaryKeyColumnsOk returns a tuple with the PrimaryKeyColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeyColumns

`func (o *Schema) SetPrimaryKeyColumns(v []string)`

SetPrimaryKeyColumns sets PrimaryKeyColumns field to given value.

### HasPrimaryKeyColumns

`func (o *Schema) HasPrimaryKeyColumns() bool`

HasPrimaryKeyColumns returns a boolean if a field has been set.

### GetDimensionFieldSpecs

`func (o *Schema) GetDimensionFieldSpecs() []DimensionFieldSpec`

GetDimensionFieldSpecs returns the DimensionFieldSpecs field if non-nil, zero value otherwise.

### GetDimensionFieldSpecsOk

`func (o *Schema) GetDimensionFieldSpecsOk() (*[]DimensionFieldSpec, bool)`

GetDimensionFieldSpecsOk returns a tuple with the DimensionFieldSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionFieldSpecs

`func (o *Schema) SetDimensionFieldSpecs(v []DimensionFieldSpec)`

SetDimensionFieldSpecs sets DimensionFieldSpecs field to given value.

### HasDimensionFieldSpecs

`func (o *Schema) HasDimensionFieldSpecs() bool`

HasDimensionFieldSpecs returns a boolean if a field has been set.

### GetMetricFieldSpecs

`func (o *Schema) GetMetricFieldSpecs() []MetricFieldSpec`

GetMetricFieldSpecs returns the MetricFieldSpecs field if non-nil, zero value otherwise.

### GetMetricFieldSpecsOk

`func (o *Schema) GetMetricFieldSpecsOk() (*[]MetricFieldSpec, bool)`

GetMetricFieldSpecsOk returns a tuple with the MetricFieldSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricFieldSpecs

`func (o *Schema) SetMetricFieldSpecs(v []MetricFieldSpec)`

SetMetricFieldSpecs sets MetricFieldSpecs field to given value.

### HasMetricFieldSpecs

`func (o *Schema) HasMetricFieldSpecs() bool`

HasMetricFieldSpecs returns a boolean if a field has been set.

### GetDateTimeFieldSpecs

`func (o *Schema) GetDateTimeFieldSpecs() []DateTimeFieldSpec`

GetDateTimeFieldSpecs returns the DateTimeFieldSpecs field if non-nil, zero value otherwise.

### GetDateTimeFieldSpecsOk

`func (o *Schema) GetDateTimeFieldSpecsOk() (*[]DateTimeFieldSpec, bool)`

GetDateTimeFieldSpecsOk returns a tuple with the DateTimeFieldSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFieldSpecs

`func (o *Schema) SetDateTimeFieldSpecs(v []DateTimeFieldSpec)`

SetDateTimeFieldSpecs sets DateTimeFieldSpecs field to given value.

### HasDateTimeFieldSpecs

`func (o *Schema) HasDateTimeFieldSpecs() bool`

HasDateTimeFieldSpecs returns a boolean if a field has been set.

### GetTimeFieldSpec

`func (o *Schema) GetTimeFieldSpec() TimeFieldSpec`

GetTimeFieldSpec returns the TimeFieldSpec field if non-nil, zero value otherwise.

### GetTimeFieldSpecOk

`func (o *Schema) GetTimeFieldSpecOk() (*TimeFieldSpec, bool)`

GetTimeFieldSpecOk returns a tuple with the TimeFieldSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFieldSpec

`func (o *Schema) SetTimeFieldSpec(v TimeFieldSpec)`

SetTimeFieldSpec sets TimeFieldSpec field to given value.

### HasTimeFieldSpec

`func (o *Schema) HasTimeFieldSpec() bool`

HasTimeFieldSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


