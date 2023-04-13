# WritePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to **[]map[string]map[string]interface{}** |  | [optional] 
**ColumnNames** | Pointer to **[]string** |  | [optional] 
**Rows** | Pointer to **[][]map[string]interface{}** |  | [optional] 

## Methods

### NewWritePayload

`func NewWritePayload() *WritePayload`

NewWritePayload instantiates a new WritePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritePayloadWithDefaults

`func NewWritePayloadWithDefaults() *WritePayload`

NewWritePayloadWithDefaults instantiates a new WritePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *WritePayload) GetValues() []map[string]map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WritePayload) GetValuesOk() (*[]map[string]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *WritePayload) SetValues(v []map[string]map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *WritePayload) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetColumnNames

`func (o *WritePayload) GetColumnNames() []string`

GetColumnNames returns the ColumnNames field if non-nil, zero value otherwise.

### GetColumnNamesOk

`func (o *WritePayload) GetColumnNamesOk() (*[]string, bool)`

GetColumnNamesOk returns a tuple with the ColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNames

`func (o *WritePayload) SetColumnNames(v []string)`

SetColumnNames sets ColumnNames field to given value.

### HasColumnNames

`func (o *WritePayload) HasColumnNames() bool`

HasColumnNames returns a boolean if a field has been set.

### GetRows

`func (o *WritePayload) GetRows() [][]map[string]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *WritePayload) GetRowsOk() (*[][]map[string]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *WritePayload) SetRows(v [][]map[string]interface{})`

SetRows sets Rows field to given value.

### HasRows

`func (o *WritePayload) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


