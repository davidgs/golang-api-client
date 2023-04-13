# BatchIngestionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchConfigMaps** | Pointer to **[]map[string]string** |  | [optional] 
**SegmentIngestionType** | Pointer to **string** |  | [optional] 
**SegmentIngestionFrequency** | Pointer to **string** |  | [optional] 
**ConsistentDataPush** | Pointer to **bool** |  | [optional] 

## Methods

### NewBatchIngestionConfig

`func NewBatchIngestionConfig() *BatchIngestionConfig`

NewBatchIngestionConfig instantiates a new BatchIngestionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchIngestionConfigWithDefaults

`func NewBatchIngestionConfigWithDefaults() *BatchIngestionConfig`

NewBatchIngestionConfigWithDefaults instantiates a new BatchIngestionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchConfigMaps

`func (o *BatchIngestionConfig) GetBatchConfigMaps() []map[string]string`

GetBatchConfigMaps returns the BatchConfigMaps field if non-nil, zero value otherwise.

### GetBatchConfigMapsOk

`func (o *BatchIngestionConfig) GetBatchConfigMapsOk() (*[]map[string]string, bool)`

GetBatchConfigMapsOk returns a tuple with the BatchConfigMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchConfigMaps

`func (o *BatchIngestionConfig) SetBatchConfigMaps(v []map[string]string)`

SetBatchConfigMaps sets BatchConfigMaps field to given value.

### HasBatchConfigMaps

`func (o *BatchIngestionConfig) HasBatchConfigMaps() bool`

HasBatchConfigMaps returns a boolean if a field has been set.

### GetSegmentIngestionType

`func (o *BatchIngestionConfig) GetSegmentIngestionType() string`

GetSegmentIngestionType returns the SegmentIngestionType field if non-nil, zero value otherwise.

### GetSegmentIngestionTypeOk

`func (o *BatchIngestionConfig) GetSegmentIngestionTypeOk() (*string, bool)`

GetSegmentIngestionTypeOk returns a tuple with the SegmentIngestionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIngestionType

`func (o *BatchIngestionConfig) SetSegmentIngestionType(v string)`

SetSegmentIngestionType sets SegmentIngestionType field to given value.

### HasSegmentIngestionType

`func (o *BatchIngestionConfig) HasSegmentIngestionType() bool`

HasSegmentIngestionType returns a boolean if a field has been set.

### GetSegmentIngestionFrequency

`func (o *BatchIngestionConfig) GetSegmentIngestionFrequency() string`

GetSegmentIngestionFrequency returns the SegmentIngestionFrequency field if non-nil, zero value otherwise.

### GetSegmentIngestionFrequencyOk

`func (o *BatchIngestionConfig) GetSegmentIngestionFrequencyOk() (*string, bool)`

GetSegmentIngestionFrequencyOk returns a tuple with the SegmentIngestionFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIngestionFrequency

`func (o *BatchIngestionConfig) SetSegmentIngestionFrequency(v string)`

SetSegmentIngestionFrequency sets SegmentIngestionFrequency field to given value.

### HasSegmentIngestionFrequency

`func (o *BatchIngestionConfig) HasSegmentIngestionFrequency() bool`

HasSegmentIngestionFrequency returns a boolean if a field has been set.

### GetConsistentDataPush

`func (o *BatchIngestionConfig) GetConsistentDataPush() bool`

GetConsistentDataPush returns the ConsistentDataPush field if non-nil, zero value otherwise.

### GetConsistentDataPushOk

`func (o *BatchIngestionConfig) GetConsistentDataPushOk() (*bool, bool)`

GetConsistentDataPushOk returns a tuple with the ConsistentDataPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistentDataPush

`func (o *BatchIngestionConfig) SetConsistentDataPush(v bool)`

SetConsistentDataPush sets ConsistentDataPush field to given value.

### HasConsistentDataPush

`func (o *BatchIngestionConfig) HasConsistentDataPush() bool`

HasConsistentDataPush returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


