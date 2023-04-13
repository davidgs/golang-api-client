# ServerRebalanceJobStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableRebalanceProgressStats** | Pointer to [**TableRebalanceProgressStats**](TableRebalanceProgressStats.md) |  | [optional] 
**TimeElapsedSinceStartInSeconds** | Pointer to **int64** |  | [optional] 

## Methods

### NewServerRebalanceJobStatusResponse

`func NewServerRebalanceJobStatusResponse() *ServerRebalanceJobStatusResponse`

NewServerRebalanceJobStatusResponse instantiates a new ServerRebalanceJobStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerRebalanceJobStatusResponseWithDefaults

`func NewServerRebalanceJobStatusResponseWithDefaults() *ServerRebalanceJobStatusResponse`

NewServerRebalanceJobStatusResponseWithDefaults instantiates a new ServerRebalanceJobStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableRebalanceProgressStats

`func (o *ServerRebalanceJobStatusResponse) GetTableRebalanceProgressStats() TableRebalanceProgressStats`

GetTableRebalanceProgressStats returns the TableRebalanceProgressStats field if non-nil, zero value otherwise.

### GetTableRebalanceProgressStatsOk

`func (o *ServerRebalanceJobStatusResponse) GetTableRebalanceProgressStatsOk() (*TableRebalanceProgressStats, bool)`

GetTableRebalanceProgressStatsOk returns a tuple with the TableRebalanceProgressStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableRebalanceProgressStats

`func (o *ServerRebalanceJobStatusResponse) SetTableRebalanceProgressStats(v TableRebalanceProgressStats)`

SetTableRebalanceProgressStats sets TableRebalanceProgressStats field to given value.

### HasTableRebalanceProgressStats

`func (o *ServerRebalanceJobStatusResponse) HasTableRebalanceProgressStats() bool`

HasTableRebalanceProgressStats returns a boolean if a field has been set.

### GetTimeElapsedSinceStartInSeconds

`func (o *ServerRebalanceJobStatusResponse) GetTimeElapsedSinceStartInSeconds() int64`

GetTimeElapsedSinceStartInSeconds returns the TimeElapsedSinceStartInSeconds field if non-nil, zero value otherwise.

### GetTimeElapsedSinceStartInSecondsOk

`func (o *ServerRebalanceJobStatusResponse) GetTimeElapsedSinceStartInSecondsOk() (*int64, bool)`

GetTimeElapsedSinceStartInSecondsOk returns a tuple with the TimeElapsedSinceStartInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeElapsedSinceStartInSeconds

`func (o *ServerRebalanceJobStatusResponse) SetTimeElapsedSinceStartInSeconds(v int64)`

SetTimeElapsedSinceStartInSeconds sets TimeElapsedSinceStartInSeconds field to given value.

### HasTimeElapsedSinceStartInSeconds

`func (o *ServerRebalanceJobStatusResponse) HasTimeElapsedSinceStartInSeconds() bool`

HasTimeElapsedSinceStartInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


