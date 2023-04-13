# QueryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeoutMs** | Pointer to **int64** |  | [optional] [readonly] 
**DisableGroovy** | Pointer to **bool** |  | [optional] [readonly] 
**UseApproximateFunction** | Pointer to **bool** |  | [optional] [readonly] 
**ExpressionOverrideMap** | Pointer to **map[string]string** |  | [optional] [readonly] 

## Methods

### NewQueryConfig

`func NewQueryConfig() *QueryConfig`

NewQueryConfig instantiates a new QueryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryConfigWithDefaults

`func NewQueryConfigWithDefaults() *QueryConfig`

NewQueryConfigWithDefaults instantiates a new QueryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeoutMs

`func (o *QueryConfig) GetTimeoutMs() int64`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *QueryConfig) GetTimeoutMsOk() (*int64, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *QueryConfig) SetTimeoutMs(v int64)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *QueryConfig) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.

### GetDisableGroovy

`func (o *QueryConfig) GetDisableGroovy() bool`

GetDisableGroovy returns the DisableGroovy field if non-nil, zero value otherwise.

### GetDisableGroovyOk

`func (o *QueryConfig) GetDisableGroovyOk() (*bool, bool)`

GetDisableGroovyOk returns a tuple with the DisableGroovy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableGroovy

`func (o *QueryConfig) SetDisableGroovy(v bool)`

SetDisableGroovy sets DisableGroovy field to given value.

### HasDisableGroovy

`func (o *QueryConfig) HasDisableGroovy() bool`

HasDisableGroovy returns a boolean if a field has been set.

### GetUseApproximateFunction

`func (o *QueryConfig) GetUseApproximateFunction() bool`

GetUseApproximateFunction returns the UseApproximateFunction field if non-nil, zero value otherwise.

### GetUseApproximateFunctionOk

`func (o *QueryConfig) GetUseApproximateFunctionOk() (*bool, bool)`

GetUseApproximateFunctionOk returns a tuple with the UseApproximateFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseApproximateFunction

`func (o *QueryConfig) SetUseApproximateFunction(v bool)`

SetUseApproximateFunction sets UseApproximateFunction field to given value.

### HasUseApproximateFunction

`func (o *QueryConfig) HasUseApproximateFunction() bool`

HasUseApproximateFunction returns a boolean if a field has been set.

### GetExpressionOverrideMap

`func (o *QueryConfig) GetExpressionOverrideMap() map[string]string`

GetExpressionOverrideMap returns the ExpressionOverrideMap field if non-nil, zero value otherwise.

### GetExpressionOverrideMapOk

`func (o *QueryConfig) GetExpressionOverrideMapOk() (*map[string]string, bool)`

GetExpressionOverrideMapOk returns a tuple with the ExpressionOverrideMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionOverrideMap

`func (o *QueryConfig) SetExpressionOverrideMap(v map[string]string)`

SetExpressionOverrideMap sets ExpressionOverrideMap field to given value.

### HasExpressionOverrideMap

`func (o *QueryConfig) HasExpressionOverrideMap() bool`

HasExpressionOverrideMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


