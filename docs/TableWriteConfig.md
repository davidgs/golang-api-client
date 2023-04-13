# TableWriteConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** |  | [optional] 
**ProducerConfig** | Pointer to **map[string]string** |  | [optional] 
**EncoderConfig** | Pointer to **map[string]string** |  | [optional] 
**EncoderClass** | Pointer to **string** |  | [optional] 
**ProducerType** | Pointer to **string** |  | [optional] 
**PartitionColumns** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTableWriteConfig

`func NewTableWriteConfig() *TableWriteConfig`

NewTableWriteConfig instantiates a new TableWriteConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableWriteConfigWithDefaults

`func NewTableWriteConfigWithDefaults() *TableWriteConfig`

NewTableWriteConfigWithDefaults instantiates a new TableWriteConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *TableWriteConfig) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TableWriteConfig) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TableWriteConfig) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *TableWriteConfig) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetProducerConfig

`func (o *TableWriteConfig) GetProducerConfig() map[string]string`

GetProducerConfig returns the ProducerConfig field if non-nil, zero value otherwise.

### GetProducerConfigOk

`func (o *TableWriteConfig) GetProducerConfigOk() (*map[string]string, bool)`

GetProducerConfigOk returns a tuple with the ProducerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerConfig

`func (o *TableWriteConfig) SetProducerConfig(v map[string]string)`

SetProducerConfig sets ProducerConfig field to given value.

### HasProducerConfig

`func (o *TableWriteConfig) HasProducerConfig() bool`

HasProducerConfig returns a boolean if a field has been set.

### GetEncoderConfig

`func (o *TableWriteConfig) GetEncoderConfig() map[string]string`

GetEncoderConfig returns the EncoderConfig field if non-nil, zero value otherwise.

### GetEncoderConfigOk

`func (o *TableWriteConfig) GetEncoderConfigOk() (*map[string]string, bool)`

GetEncoderConfigOk returns a tuple with the EncoderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderConfig

`func (o *TableWriteConfig) SetEncoderConfig(v map[string]string)`

SetEncoderConfig sets EncoderConfig field to given value.

### HasEncoderConfig

`func (o *TableWriteConfig) HasEncoderConfig() bool`

HasEncoderConfig returns a boolean if a field has been set.

### GetEncoderClass

`func (o *TableWriteConfig) GetEncoderClass() string`

GetEncoderClass returns the EncoderClass field if non-nil, zero value otherwise.

### GetEncoderClassOk

`func (o *TableWriteConfig) GetEncoderClassOk() (*string, bool)`

GetEncoderClassOk returns a tuple with the EncoderClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderClass

`func (o *TableWriteConfig) SetEncoderClass(v string)`

SetEncoderClass sets EncoderClass field to given value.

### HasEncoderClass

`func (o *TableWriteConfig) HasEncoderClass() bool`

HasEncoderClass returns a boolean if a field has been set.

### GetProducerType

`func (o *TableWriteConfig) GetProducerType() string`

GetProducerType returns the ProducerType field if non-nil, zero value otherwise.

### GetProducerTypeOk

`func (o *TableWriteConfig) GetProducerTypeOk() (*string, bool)`

GetProducerTypeOk returns a tuple with the ProducerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerType

`func (o *TableWriteConfig) SetProducerType(v string)`

SetProducerType sets ProducerType field to given value.

### HasProducerType

`func (o *TableWriteConfig) HasProducerType() bool`

HasProducerType returns a boolean if a field has been set.

### GetPartitionColumns

`func (o *TableWriteConfig) GetPartitionColumns() []string`

GetPartitionColumns returns the PartitionColumns field if non-nil, zero value otherwise.

### GetPartitionColumnsOk

`func (o *TableWriteConfig) GetPartitionColumnsOk() (*[]string, bool)`

GetPartitionColumnsOk returns a tuple with the PartitionColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionColumns

`func (o *TableWriteConfig) SetPartitionColumns(v []string)`

SetPartitionColumns sets PartitionColumns field to given value.

### HasPartitionColumns

`func (o *TableWriteConfig) HasPartitionColumns() bool`

HasPartitionColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


