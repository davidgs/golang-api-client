# IngestionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamIngestionConfig** | Pointer to [**StreamIngestionConfig**](StreamIngestionConfig.md) |  | [optional] 
**BatchIngestionConfig** | Pointer to [**BatchIngestionConfig**](BatchIngestionConfig.md) |  | [optional] 
**SegmentTimeValueCheck** | Pointer to **bool** |  | [optional] 
**FilterConfig** | Pointer to [**FilterConfig**](FilterConfig.md) |  | [optional] 
**TransformConfigs** | Pointer to [**[]TransformConfig**](TransformConfig.md) |  | [optional] 
**ComplexTypeConfig** | Pointer to [**ComplexTypeConfig**](ComplexTypeConfig.md) |  | [optional] 
**AggregationConfigs** | Pointer to [**[]AggregationConfig**](AggregationConfig.md) |  | [optional] 
**ContinueOnError** | Pointer to **bool** |  | [optional] 
**RowTimeValueCheck** | Pointer to **bool** |  | [optional] 

## Methods

### NewIngestionConfig

`func NewIngestionConfig() *IngestionConfig`

NewIngestionConfig instantiates a new IngestionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionConfigWithDefaults

`func NewIngestionConfigWithDefaults() *IngestionConfig`

NewIngestionConfigWithDefaults instantiates a new IngestionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamIngestionConfig

`func (o *IngestionConfig) GetStreamIngestionConfig() StreamIngestionConfig`

GetStreamIngestionConfig returns the StreamIngestionConfig field if non-nil, zero value otherwise.

### GetStreamIngestionConfigOk

`func (o *IngestionConfig) GetStreamIngestionConfigOk() (*StreamIngestionConfig, bool)`

GetStreamIngestionConfigOk returns a tuple with the StreamIngestionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamIngestionConfig

`func (o *IngestionConfig) SetStreamIngestionConfig(v StreamIngestionConfig)`

SetStreamIngestionConfig sets StreamIngestionConfig field to given value.

### HasStreamIngestionConfig

`func (o *IngestionConfig) HasStreamIngestionConfig() bool`

HasStreamIngestionConfig returns a boolean if a field has been set.

### GetBatchIngestionConfig

`func (o *IngestionConfig) GetBatchIngestionConfig() BatchIngestionConfig`

GetBatchIngestionConfig returns the BatchIngestionConfig field if non-nil, zero value otherwise.

### GetBatchIngestionConfigOk

`func (o *IngestionConfig) GetBatchIngestionConfigOk() (*BatchIngestionConfig, bool)`

GetBatchIngestionConfigOk returns a tuple with the BatchIngestionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchIngestionConfig

`func (o *IngestionConfig) SetBatchIngestionConfig(v BatchIngestionConfig)`

SetBatchIngestionConfig sets BatchIngestionConfig field to given value.

### HasBatchIngestionConfig

`func (o *IngestionConfig) HasBatchIngestionConfig() bool`

HasBatchIngestionConfig returns a boolean if a field has been set.

### GetSegmentTimeValueCheck

`func (o *IngestionConfig) GetSegmentTimeValueCheck() bool`

GetSegmentTimeValueCheck returns the SegmentTimeValueCheck field if non-nil, zero value otherwise.

### GetSegmentTimeValueCheckOk

`func (o *IngestionConfig) GetSegmentTimeValueCheckOk() (*bool, bool)`

GetSegmentTimeValueCheckOk returns a tuple with the SegmentTimeValueCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentTimeValueCheck

`func (o *IngestionConfig) SetSegmentTimeValueCheck(v bool)`

SetSegmentTimeValueCheck sets SegmentTimeValueCheck field to given value.

### HasSegmentTimeValueCheck

`func (o *IngestionConfig) HasSegmentTimeValueCheck() bool`

HasSegmentTimeValueCheck returns a boolean if a field has been set.

### GetFilterConfig

`func (o *IngestionConfig) GetFilterConfig() FilterConfig`

GetFilterConfig returns the FilterConfig field if non-nil, zero value otherwise.

### GetFilterConfigOk

`func (o *IngestionConfig) GetFilterConfigOk() (*FilterConfig, bool)`

GetFilterConfigOk returns a tuple with the FilterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterConfig

`func (o *IngestionConfig) SetFilterConfig(v FilterConfig)`

SetFilterConfig sets FilterConfig field to given value.

### HasFilterConfig

`func (o *IngestionConfig) HasFilterConfig() bool`

HasFilterConfig returns a boolean if a field has been set.

### GetTransformConfigs

`func (o *IngestionConfig) GetTransformConfigs() []TransformConfig`

GetTransformConfigs returns the TransformConfigs field if non-nil, zero value otherwise.

### GetTransformConfigsOk

`func (o *IngestionConfig) GetTransformConfigsOk() (*[]TransformConfig, bool)`

GetTransformConfigsOk returns a tuple with the TransformConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformConfigs

`func (o *IngestionConfig) SetTransformConfigs(v []TransformConfig)`

SetTransformConfigs sets TransformConfigs field to given value.

### HasTransformConfigs

`func (o *IngestionConfig) HasTransformConfigs() bool`

HasTransformConfigs returns a boolean if a field has been set.

### GetComplexTypeConfig

`func (o *IngestionConfig) GetComplexTypeConfig() ComplexTypeConfig`

GetComplexTypeConfig returns the ComplexTypeConfig field if non-nil, zero value otherwise.

### GetComplexTypeConfigOk

`func (o *IngestionConfig) GetComplexTypeConfigOk() (*ComplexTypeConfig, bool)`

GetComplexTypeConfigOk returns a tuple with the ComplexTypeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexTypeConfig

`func (o *IngestionConfig) SetComplexTypeConfig(v ComplexTypeConfig)`

SetComplexTypeConfig sets ComplexTypeConfig field to given value.

### HasComplexTypeConfig

`func (o *IngestionConfig) HasComplexTypeConfig() bool`

HasComplexTypeConfig returns a boolean if a field has been set.

### GetAggregationConfigs

`func (o *IngestionConfig) GetAggregationConfigs() []AggregationConfig`

GetAggregationConfigs returns the AggregationConfigs field if non-nil, zero value otherwise.

### GetAggregationConfigsOk

`func (o *IngestionConfig) GetAggregationConfigsOk() (*[]AggregationConfig, bool)`

GetAggregationConfigsOk returns a tuple with the AggregationConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationConfigs

`func (o *IngestionConfig) SetAggregationConfigs(v []AggregationConfig)`

SetAggregationConfigs sets AggregationConfigs field to given value.

### HasAggregationConfigs

`func (o *IngestionConfig) HasAggregationConfigs() bool`

HasAggregationConfigs returns a boolean if a field has been set.

### GetContinueOnError

`func (o *IngestionConfig) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *IngestionConfig) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *IngestionConfig) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *IngestionConfig) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### GetRowTimeValueCheck

`func (o *IngestionConfig) GetRowTimeValueCheck() bool`

GetRowTimeValueCheck returns the RowTimeValueCheck field if non-nil, zero value otherwise.

### GetRowTimeValueCheckOk

`func (o *IngestionConfig) GetRowTimeValueCheckOk() (*bool, bool)`

GetRowTimeValueCheckOk returns a tuple with the RowTimeValueCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowTimeValueCheck

`func (o *IngestionConfig) SetRowTimeValueCheck(v bool)`

SetRowTimeValueCheck sets RowTimeValueCheck field to given value.

### HasRowTimeValueCheck

`func (o *IngestionConfig) HasRowTimeValueCheck() bool`

HasRowTimeValueCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


