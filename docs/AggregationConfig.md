# AggregationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnName** | Pointer to **string** |  | [optional] [readonly] 
**AggregationFunction** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAggregationConfig

`func NewAggregationConfig() *AggregationConfig`

NewAggregationConfig instantiates a new AggregationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationConfigWithDefaults

`func NewAggregationConfigWithDefaults() *AggregationConfig`

NewAggregationConfigWithDefaults instantiates a new AggregationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnName

`func (o *AggregationConfig) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *AggregationConfig) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *AggregationConfig) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *AggregationConfig) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### GetAggregationFunction

`func (o *AggregationConfig) GetAggregationFunction() string`

GetAggregationFunction returns the AggregationFunction field if non-nil, zero value otherwise.

### GetAggregationFunctionOk

`func (o *AggregationConfig) GetAggregationFunctionOk() (*string, bool)`

GetAggregationFunctionOk returns a tuple with the AggregationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFunction

`func (o *AggregationConfig) SetAggregationFunction(v string)`

SetAggregationFunction sets AggregationFunction field to given value.

### HasAggregationFunction

`func (o *AggregationConfig) HasAggregationFunction() bool`

HasAggregationFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


