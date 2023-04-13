# PinotTaskConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | Pointer to **string** |  | [optional] 
**Configs** | Pointer to **map[string]string** |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**TaskType** | Pointer to **string** |  | [optional] 

## Methods

### NewPinotTaskConfig

`func NewPinotTaskConfig() *PinotTaskConfig`

NewPinotTaskConfig instantiates a new PinotTaskConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinotTaskConfigWithDefaults

`func NewPinotTaskConfigWithDefaults() *PinotTaskConfig`

NewPinotTaskConfigWithDefaults instantiates a new PinotTaskConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *PinotTaskConfig) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *PinotTaskConfig) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *PinotTaskConfig) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *PinotTaskConfig) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetConfigs

`func (o *PinotTaskConfig) GetConfigs() map[string]string`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *PinotTaskConfig) GetConfigsOk() (*map[string]string, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *PinotTaskConfig) SetConfigs(v map[string]string)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *PinotTaskConfig) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetTaskId

`func (o *PinotTaskConfig) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *PinotTaskConfig) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *PinotTaskConfig) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *PinotTaskConfig) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskType

`func (o *PinotTaskConfig) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *PinotTaskConfig) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *PinotTaskConfig) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *PinotTaskConfig) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


