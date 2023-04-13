# UpsertConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |  | [optional] 
**ComparisonColumns** | Pointer to **[]string** |  | [optional] 
**HashFunction** | Pointer to **string** |  | [optional] 
**PartialUpsertStrategies** | Pointer to **map[string]string** |  | [optional] 
**DefaultPartialUpsertStrategy** | Pointer to **string** |  | [optional] 
**EnableSnapshot** | Pointer to **bool** |  | [optional] 
**MetadataManagerClass** | Pointer to **string** |  | [optional] 
**MetadataManagerConfigs** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUpsertConfig

`func NewUpsertConfig() *UpsertConfig`

NewUpsertConfig instantiates a new UpsertConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertConfigWithDefaults

`func NewUpsertConfigWithDefaults() *UpsertConfig`

NewUpsertConfigWithDefaults instantiates a new UpsertConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *UpsertConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpsertConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpsertConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpsertConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetComparisonColumns

`func (o *UpsertConfig) GetComparisonColumns() []string`

GetComparisonColumns returns the ComparisonColumns field if non-nil, zero value otherwise.

### GetComparisonColumnsOk

`func (o *UpsertConfig) GetComparisonColumnsOk() (*[]string, bool)`

GetComparisonColumnsOk returns a tuple with the ComparisonColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonColumns

`func (o *UpsertConfig) SetComparisonColumns(v []string)`

SetComparisonColumns sets ComparisonColumns field to given value.

### HasComparisonColumns

`func (o *UpsertConfig) HasComparisonColumns() bool`

HasComparisonColumns returns a boolean if a field has been set.

### GetHashFunction

`func (o *UpsertConfig) GetHashFunction() string`

GetHashFunction returns the HashFunction field if non-nil, zero value otherwise.

### GetHashFunctionOk

`func (o *UpsertConfig) GetHashFunctionOk() (*string, bool)`

GetHashFunctionOk returns a tuple with the HashFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashFunction

`func (o *UpsertConfig) SetHashFunction(v string)`

SetHashFunction sets HashFunction field to given value.

### HasHashFunction

`func (o *UpsertConfig) HasHashFunction() bool`

HasHashFunction returns a boolean if a field has been set.

### GetPartialUpsertStrategies

`func (o *UpsertConfig) GetPartialUpsertStrategies() map[string]string`

GetPartialUpsertStrategies returns the PartialUpsertStrategies field if non-nil, zero value otherwise.

### GetPartialUpsertStrategiesOk

`func (o *UpsertConfig) GetPartialUpsertStrategiesOk() (*map[string]string, bool)`

GetPartialUpsertStrategiesOk returns a tuple with the PartialUpsertStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialUpsertStrategies

`func (o *UpsertConfig) SetPartialUpsertStrategies(v map[string]string)`

SetPartialUpsertStrategies sets PartialUpsertStrategies field to given value.

### HasPartialUpsertStrategies

`func (o *UpsertConfig) HasPartialUpsertStrategies() bool`

HasPartialUpsertStrategies returns a boolean if a field has been set.

### GetDefaultPartialUpsertStrategy

`func (o *UpsertConfig) GetDefaultPartialUpsertStrategy() string`

GetDefaultPartialUpsertStrategy returns the DefaultPartialUpsertStrategy field if non-nil, zero value otherwise.

### GetDefaultPartialUpsertStrategyOk

`func (o *UpsertConfig) GetDefaultPartialUpsertStrategyOk() (*string, bool)`

GetDefaultPartialUpsertStrategyOk returns a tuple with the DefaultPartialUpsertStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPartialUpsertStrategy

`func (o *UpsertConfig) SetDefaultPartialUpsertStrategy(v string)`

SetDefaultPartialUpsertStrategy sets DefaultPartialUpsertStrategy field to given value.

### HasDefaultPartialUpsertStrategy

`func (o *UpsertConfig) HasDefaultPartialUpsertStrategy() bool`

HasDefaultPartialUpsertStrategy returns a boolean if a field has been set.

### GetEnableSnapshot

`func (o *UpsertConfig) GetEnableSnapshot() bool`

GetEnableSnapshot returns the EnableSnapshot field if non-nil, zero value otherwise.

### GetEnableSnapshotOk

`func (o *UpsertConfig) GetEnableSnapshotOk() (*bool, bool)`

GetEnableSnapshotOk returns a tuple with the EnableSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSnapshot

`func (o *UpsertConfig) SetEnableSnapshot(v bool)`

SetEnableSnapshot sets EnableSnapshot field to given value.

### HasEnableSnapshot

`func (o *UpsertConfig) HasEnableSnapshot() bool`

HasEnableSnapshot returns a boolean if a field has been set.

### GetMetadataManagerClass

`func (o *UpsertConfig) GetMetadataManagerClass() string`

GetMetadataManagerClass returns the MetadataManagerClass field if non-nil, zero value otherwise.

### GetMetadataManagerClassOk

`func (o *UpsertConfig) GetMetadataManagerClassOk() (*string, bool)`

GetMetadataManagerClassOk returns a tuple with the MetadataManagerClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataManagerClass

`func (o *UpsertConfig) SetMetadataManagerClass(v string)`

SetMetadataManagerClass sets MetadataManagerClass field to given value.

### HasMetadataManagerClass

`func (o *UpsertConfig) HasMetadataManagerClass() bool`

HasMetadataManagerClass returns a boolean if a field has been set.

### GetMetadataManagerConfigs

`func (o *UpsertConfig) GetMetadataManagerConfigs() map[string]string`

GetMetadataManagerConfigs returns the MetadataManagerConfigs field if non-nil, zero value otherwise.

### GetMetadataManagerConfigsOk

`func (o *UpsertConfig) GetMetadataManagerConfigsOk() (*map[string]string, bool)`

GetMetadataManagerConfigsOk returns a tuple with the MetadataManagerConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataManagerConfigs

`func (o *UpsertConfig) SetMetadataManagerConfigs(v map[string]string)`

SetMetadataManagerConfigs sets MetadataManagerConfigs field to given value.

### HasMetadataManagerConfigs

`func (o *UpsertConfig) HasMetadataManagerConfigs() bool`

HasMetadataManagerConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


