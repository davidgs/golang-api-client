# TaskDebugInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskState** | Pointer to **string** |  | [optional] 
**SubtaskCount** | Pointer to [**TaskCount**](TaskCount.md) |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**ExecutionStartTime** | Pointer to **string** |  | [optional] 
**FinishTime** | Pointer to **string** |  | [optional] 
**SubtaskInfos** | Pointer to [**[]SubtaskDebugInfo**](SubtaskDebugInfo.md) |  | [optional] 

## Methods

### NewTaskDebugInfo

`func NewTaskDebugInfo() *TaskDebugInfo`

NewTaskDebugInfo instantiates a new TaskDebugInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskDebugInfoWithDefaults

`func NewTaskDebugInfoWithDefaults() *TaskDebugInfo`

NewTaskDebugInfoWithDefaults instantiates a new TaskDebugInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskState

`func (o *TaskDebugInfo) GetTaskState() string`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *TaskDebugInfo) GetTaskStateOk() (*string, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *TaskDebugInfo) SetTaskState(v string)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *TaskDebugInfo) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### GetSubtaskCount

`func (o *TaskDebugInfo) GetSubtaskCount() TaskCount`

GetSubtaskCount returns the SubtaskCount field if non-nil, zero value otherwise.

### GetSubtaskCountOk

`func (o *TaskDebugInfo) GetSubtaskCountOk() (*TaskCount, bool)`

GetSubtaskCountOk returns a tuple with the SubtaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtaskCount

`func (o *TaskDebugInfo) SetSubtaskCount(v TaskCount)`

SetSubtaskCount sets SubtaskCount field to given value.

### HasSubtaskCount

`func (o *TaskDebugInfo) HasSubtaskCount() bool`

HasSubtaskCount returns a boolean if a field has been set.

### GetStartTime

`func (o *TaskDebugInfo) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TaskDebugInfo) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TaskDebugInfo) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TaskDebugInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetExecutionStartTime

`func (o *TaskDebugInfo) GetExecutionStartTime() string`

GetExecutionStartTime returns the ExecutionStartTime field if non-nil, zero value otherwise.

### GetExecutionStartTimeOk

`func (o *TaskDebugInfo) GetExecutionStartTimeOk() (*string, bool)`

GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStartTime

`func (o *TaskDebugInfo) SetExecutionStartTime(v string)`

SetExecutionStartTime sets ExecutionStartTime field to given value.

### HasExecutionStartTime

`func (o *TaskDebugInfo) HasExecutionStartTime() bool`

HasExecutionStartTime returns a boolean if a field has been set.

### GetFinishTime

`func (o *TaskDebugInfo) GetFinishTime() string`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *TaskDebugInfo) GetFinishTimeOk() (*string, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *TaskDebugInfo) SetFinishTime(v string)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *TaskDebugInfo) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetSubtaskInfos

`func (o *TaskDebugInfo) GetSubtaskInfos() []SubtaskDebugInfo`

GetSubtaskInfos returns the SubtaskInfos field if non-nil, zero value otherwise.

### GetSubtaskInfosOk

`func (o *TaskDebugInfo) GetSubtaskInfosOk() (*[]SubtaskDebugInfo, bool)`

GetSubtaskInfosOk returns a tuple with the SubtaskInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtaskInfos

`func (o *TaskDebugInfo) SetSubtaskInfos(v []SubtaskDebugInfo)`

SetSubtaskInfos sets SubtaskInfos field to given value.

### HasSubtaskInfos

`func (o *TaskDebugInfo) HasSubtaskInfos() bool`

HasSubtaskInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


