# StarTreeIndexConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DimensionsSplitOrder** | **[]string** |  | [readonly] 
**SkipStarNodeCreationForDimensions** | Pointer to **[]string** |  | [optional] [readonly] 
**FunctionColumnPairs** | **[]string** |  | [readonly] 
**MaxLeafRecords** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewStarTreeIndexConfig

`func NewStarTreeIndexConfig(dimensionsSplitOrder []string, functionColumnPairs []string, ) *StarTreeIndexConfig`

NewStarTreeIndexConfig instantiates a new StarTreeIndexConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStarTreeIndexConfigWithDefaults

`func NewStarTreeIndexConfigWithDefaults() *StarTreeIndexConfig`

NewStarTreeIndexConfigWithDefaults instantiates a new StarTreeIndexConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensionsSplitOrder

`func (o *StarTreeIndexConfig) GetDimensionsSplitOrder() []string`

GetDimensionsSplitOrder returns the DimensionsSplitOrder field if non-nil, zero value otherwise.

### GetDimensionsSplitOrderOk

`func (o *StarTreeIndexConfig) GetDimensionsSplitOrderOk() (*[]string, bool)`

GetDimensionsSplitOrderOk returns a tuple with the DimensionsSplitOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsSplitOrder

`func (o *StarTreeIndexConfig) SetDimensionsSplitOrder(v []string)`

SetDimensionsSplitOrder sets DimensionsSplitOrder field to given value.


### GetSkipStarNodeCreationForDimensions

`func (o *StarTreeIndexConfig) GetSkipStarNodeCreationForDimensions() []string`

GetSkipStarNodeCreationForDimensions returns the SkipStarNodeCreationForDimensions field if non-nil, zero value otherwise.

### GetSkipStarNodeCreationForDimensionsOk

`func (o *StarTreeIndexConfig) GetSkipStarNodeCreationForDimensionsOk() (*[]string, bool)`

GetSkipStarNodeCreationForDimensionsOk returns a tuple with the SkipStarNodeCreationForDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipStarNodeCreationForDimensions

`func (o *StarTreeIndexConfig) SetSkipStarNodeCreationForDimensions(v []string)`

SetSkipStarNodeCreationForDimensions sets SkipStarNodeCreationForDimensions field to given value.

### HasSkipStarNodeCreationForDimensions

`func (o *StarTreeIndexConfig) HasSkipStarNodeCreationForDimensions() bool`

HasSkipStarNodeCreationForDimensions returns a boolean if a field has been set.

### GetFunctionColumnPairs

`func (o *StarTreeIndexConfig) GetFunctionColumnPairs() []string`

GetFunctionColumnPairs returns the FunctionColumnPairs field if non-nil, zero value otherwise.

### GetFunctionColumnPairsOk

`func (o *StarTreeIndexConfig) GetFunctionColumnPairsOk() (*[]string, bool)`

GetFunctionColumnPairsOk returns a tuple with the FunctionColumnPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionColumnPairs

`func (o *StarTreeIndexConfig) SetFunctionColumnPairs(v []string)`

SetFunctionColumnPairs sets FunctionColumnPairs field to given value.


### GetMaxLeafRecords

`func (o *StarTreeIndexConfig) GetMaxLeafRecords() int32`

GetMaxLeafRecords returns the MaxLeafRecords field if non-nil, zero value otherwise.

### GetMaxLeafRecordsOk

`func (o *StarTreeIndexConfig) GetMaxLeafRecordsOk() (*int32, bool)`

GetMaxLeafRecordsOk returns a tuple with the MaxLeafRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLeafRecords

`func (o *StarTreeIndexConfig) SetMaxLeafRecords(v int32)`

SetMaxLeafRecords sets MaxLeafRecords field to given value.

### HasMaxLeafRecords

`func (o *StarTreeIndexConfig) HasMaxLeafRecords() bool`

HasMaxLeafRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


