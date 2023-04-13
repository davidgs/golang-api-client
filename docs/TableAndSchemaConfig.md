# TableAndSchemaConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableConfig** | [**TableConfig**](TableConfig.md) |  | 
**Schema** | Pointer to [**Schema**](Schema.md) |  | [optional] 

## Methods

### NewTableAndSchemaConfig

`func NewTableAndSchemaConfig(tableConfig TableConfig, ) *TableAndSchemaConfig`

NewTableAndSchemaConfig instantiates a new TableAndSchemaConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableAndSchemaConfigWithDefaults

`func NewTableAndSchemaConfigWithDefaults() *TableAndSchemaConfig`

NewTableAndSchemaConfigWithDefaults instantiates a new TableAndSchemaConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableConfig

`func (o *TableAndSchemaConfig) GetTableConfig() TableConfig`

GetTableConfig returns the TableConfig field if non-nil, zero value otherwise.

### GetTableConfigOk

`func (o *TableAndSchemaConfig) GetTableConfigOk() (*TableConfig, bool)`

GetTableConfigOk returns a tuple with the TableConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableConfig

`func (o *TableAndSchemaConfig) SetTableConfig(v TableConfig)`

SetTableConfig sets TableConfig field to given value.


### GetSchema

`func (o *TableAndSchemaConfig) GetSchema() Schema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *TableAndSchemaConfig) GetSchemaOk() (*Schema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *TableAndSchemaConfig) SetSchema(v Schema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *TableAndSchemaConfig) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


