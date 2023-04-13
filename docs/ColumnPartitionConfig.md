# ColumnPartitionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionName** | **string** |  | [readonly] 
**NumPartitions** | **int32** |  | [readonly] 
**FunctionConfig** | Pointer to **map[string]string** |  | [optional] [readonly] 

## Methods

### NewColumnPartitionConfig

`func NewColumnPartitionConfig(functionName string, numPartitions int32, ) *ColumnPartitionConfig`

NewColumnPartitionConfig instantiates a new ColumnPartitionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnPartitionConfigWithDefaults

`func NewColumnPartitionConfigWithDefaults() *ColumnPartitionConfig`

NewColumnPartitionConfigWithDefaults instantiates a new ColumnPartitionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionName

`func (o *ColumnPartitionConfig) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *ColumnPartitionConfig) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *ColumnPartitionConfig) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetNumPartitions

`func (o *ColumnPartitionConfig) GetNumPartitions() int32`

GetNumPartitions returns the NumPartitions field if non-nil, zero value otherwise.

### GetNumPartitionsOk

`func (o *ColumnPartitionConfig) GetNumPartitionsOk() (*int32, bool)`

GetNumPartitionsOk returns a tuple with the NumPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitions

`func (o *ColumnPartitionConfig) SetNumPartitions(v int32)`

SetNumPartitions sets NumPartitions field to given value.


### GetFunctionConfig

`func (o *ColumnPartitionConfig) GetFunctionConfig() map[string]string`

GetFunctionConfig returns the FunctionConfig field if non-nil, zero value otherwise.

### GetFunctionConfigOk

`func (o *ColumnPartitionConfig) GetFunctionConfigOk() (*map[string]string, bool)`

GetFunctionConfigOk returns a tuple with the FunctionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionConfig

`func (o *ColumnPartitionConfig) SetFunctionConfig(v map[string]string)`

SetFunctionConfig sets FunctionConfig field to given value.

### HasFunctionConfig

`func (o *ColumnPartitionConfig) HasFunctionConfig() bool`

HasFunctionConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


