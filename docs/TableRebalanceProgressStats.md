# TableRebalanceProgressStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**StartTimeMs** | Pointer to **int64** |  | [optional] 
**InitialToTargetStateConvergence** | Pointer to [**RebalanceStateStats**](RebalanceStateStats.md) |  | [optional] 
**TimeToFinishInSeconds** | Pointer to **int64** |  | [optional] 
**ExternalViewToIdealStateConvergence** | Pointer to [**RebalanceStateStats**](RebalanceStateStats.md) |  | [optional] 
**CurrentToTargetConvergence** | Pointer to [**RebalanceStateStats**](RebalanceStateStats.md) |  | [optional] 
**CompletionStatusMsg** | Pointer to **string** |  | [optional] 

## Methods

### NewTableRebalanceProgressStats

`func NewTableRebalanceProgressStats() *TableRebalanceProgressStats`

NewTableRebalanceProgressStats instantiates a new TableRebalanceProgressStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableRebalanceProgressStatsWithDefaults

`func NewTableRebalanceProgressStatsWithDefaults() *TableRebalanceProgressStats`

NewTableRebalanceProgressStatsWithDefaults instantiates a new TableRebalanceProgressStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TableRebalanceProgressStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TableRebalanceProgressStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TableRebalanceProgressStats) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TableRebalanceProgressStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTimeMs

`func (o *TableRebalanceProgressStats) GetStartTimeMs() int64`

GetStartTimeMs returns the StartTimeMs field if non-nil, zero value otherwise.

### GetStartTimeMsOk

`func (o *TableRebalanceProgressStats) GetStartTimeMsOk() (*int64, bool)`

GetStartTimeMsOk returns a tuple with the StartTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeMs

`func (o *TableRebalanceProgressStats) SetStartTimeMs(v int64)`

SetStartTimeMs sets StartTimeMs field to given value.

### HasStartTimeMs

`func (o *TableRebalanceProgressStats) HasStartTimeMs() bool`

HasStartTimeMs returns a boolean if a field has been set.

### GetInitialToTargetStateConvergence

`func (o *TableRebalanceProgressStats) GetInitialToTargetStateConvergence() RebalanceStateStats`

GetInitialToTargetStateConvergence returns the InitialToTargetStateConvergence field if non-nil, zero value otherwise.

### GetInitialToTargetStateConvergenceOk

`func (o *TableRebalanceProgressStats) GetInitialToTargetStateConvergenceOk() (*RebalanceStateStats, bool)`

GetInitialToTargetStateConvergenceOk returns a tuple with the InitialToTargetStateConvergence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialToTargetStateConvergence

`func (o *TableRebalanceProgressStats) SetInitialToTargetStateConvergence(v RebalanceStateStats)`

SetInitialToTargetStateConvergence sets InitialToTargetStateConvergence field to given value.

### HasInitialToTargetStateConvergence

`func (o *TableRebalanceProgressStats) HasInitialToTargetStateConvergence() bool`

HasInitialToTargetStateConvergence returns a boolean if a field has been set.

### GetTimeToFinishInSeconds

`func (o *TableRebalanceProgressStats) GetTimeToFinishInSeconds() int64`

GetTimeToFinishInSeconds returns the TimeToFinishInSeconds field if non-nil, zero value otherwise.

### GetTimeToFinishInSecondsOk

`func (o *TableRebalanceProgressStats) GetTimeToFinishInSecondsOk() (*int64, bool)`

GetTimeToFinishInSecondsOk returns a tuple with the TimeToFinishInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToFinishInSeconds

`func (o *TableRebalanceProgressStats) SetTimeToFinishInSeconds(v int64)`

SetTimeToFinishInSeconds sets TimeToFinishInSeconds field to given value.

### HasTimeToFinishInSeconds

`func (o *TableRebalanceProgressStats) HasTimeToFinishInSeconds() bool`

HasTimeToFinishInSeconds returns a boolean if a field has been set.

### GetExternalViewToIdealStateConvergence

`func (o *TableRebalanceProgressStats) GetExternalViewToIdealStateConvergence() RebalanceStateStats`

GetExternalViewToIdealStateConvergence returns the ExternalViewToIdealStateConvergence field if non-nil, zero value otherwise.

### GetExternalViewToIdealStateConvergenceOk

`func (o *TableRebalanceProgressStats) GetExternalViewToIdealStateConvergenceOk() (*RebalanceStateStats, bool)`

GetExternalViewToIdealStateConvergenceOk returns a tuple with the ExternalViewToIdealStateConvergence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalViewToIdealStateConvergence

`func (o *TableRebalanceProgressStats) SetExternalViewToIdealStateConvergence(v RebalanceStateStats)`

SetExternalViewToIdealStateConvergence sets ExternalViewToIdealStateConvergence field to given value.

### HasExternalViewToIdealStateConvergence

`func (o *TableRebalanceProgressStats) HasExternalViewToIdealStateConvergence() bool`

HasExternalViewToIdealStateConvergence returns a boolean if a field has been set.

### GetCurrentToTargetConvergence

`func (o *TableRebalanceProgressStats) GetCurrentToTargetConvergence() RebalanceStateStats`

GetCurrentToTargetConvergence returns the CurrentToTargetConvergence field if non-nil, zero value otherwise.

### GetCurrentToTargetConvergenceOk

`func (o *TableRebalanceProgressStats) GetCurrentToTargetConvergenceOk() (*RebalanceStateStats, bool)`

GetCurrentToTargetConvergenceOk returns a tuple with the CurrentToTargetConvergence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentToTargetConvergence

`func (o *TableRebalanceProgressStats) SetCurrentToTargetConvergence(v RebalanceStateStats)`

SetCurrentToTargetConvergence sets CurrentToTargetConvergence field to given value.

### HasCurrentToTargetConvergence

`func (o *TableRebalanceProgressStats) HasCurrentToTargetConvergence() bool`

HasCurrentToTargetConvergence returns a boolean if a field has been set.

### GetCompletionStatusMsg

`func (o *TableRebalanceProgressStats) GetCompletionStatusMsg() string`

GetCompletionStatusMsg returns the CompletionStatusMsg field if non-nil, zero value otherwise.

### GetCompletionStatusMsgOk

`func (o *TableRebalanceProgressStats) GetCompletionStatusMsgOk() (*string, bool)`

GetCompletionStatusMsgOk returns a tuple with the CompletionStatusMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionStatusMsg

`func (o *TableRebalanceProgressStats) SetCompletionStatusMsg(v string)`

SetCompletionStatusMsg sets CompletionStatusMsg field to given value.

### HasCompletionStatusMsg

`func (o *TableRebalanceProgressStats) HasCompletionStatusMsg() bool`

HasCompletionStatusMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


