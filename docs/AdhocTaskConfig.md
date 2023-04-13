# AdhocTaskConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskType** | **string** |  | [readonly] 
**TableName** | **string** |  | [readonly] 
**TaskName** | Pointer to **string** |  | [optional] [readonly] 
**TaskConfigs** | Pointer to **map[string]string** |  | [optional] [readonly] 

## Methods

### NewAdhocTaskConfig

`func NewAdhocTaskConfig(taskType string, tableName string, ) *AdhocTaskConfig`

NewAdhocTaskConfig instantiates a new AdhocTaskConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdhocTaskConfigWithDefaults

`func NewAdhocTaskConfigWithDefaults() *AdhocTaskConfig`

NewAdhocTaskConfigWithDefaults instantiates a new AdhocTaskConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskType

`func (o *AdhocTaskConfig) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *AdhocTaskConfig) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *AdhocTaskConfig) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.


### GetTableName

`func (o *AdhocTaskConfig) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *AdhocTaskConfig) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *AdhocTaskConfig) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetTaskName

`func (o *AdhocTaskConfig) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *AdhocTaskConfig) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *AdhocTaskConfig) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *AdhocTaskConfig) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTaskConfigs

`func (o *AdhocTaskConfig) GetTaskConfigs() map[string]string`

GetTaskConfigs returns the TaskConfigs field if non-nil, zero value otherwise.

### GetTaskConfigsOk

`func (o *AdhocTaskConfig) GetTaskConfigsOk() (*map[string]string, bool)`

GetTaskConfigsOk returns a tuple with the TaskConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskConfigs

`func (o *AdhocTaskConfig) SetTaskConfigs(v map[string]string)`

SetTaskConfigs sets TaskConfigs field to given value.

### HasTaskConfigs

`func (o *AdhocTaskConfig) HasTaskConfigs() bool`

HasTaskConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


