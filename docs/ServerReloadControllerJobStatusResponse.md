# ServerReloadControllerJobStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**SuccessCount** | Pointer to **int32** |  | [optional] 
**TotalSegmentCount** | Pointer to **int32** |  | [optional] 
**TotalServersQueried** | Pointer to **int32** |  | [optional] 
**TotalServerCallsFailed** | Pointer to **int32** |  | [optional] 
**TimeElapsedInMinutes** | Pointer to **float64** |  | [optional] 
**EstimatedTimeRemainingInMinutes** | Pointer to **float64** |  | [optional] 

## Methods

### NewServerReloadControllerJobStatusResponse

`func NewServerReloadControllerJobStatusResponse() *ServerReloadControllerJobStatusResponse`

NewServerReloadControllerJobStatusResponse instantiates a new ServerReloadControllerJobStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerReloadControllerJobStatusResponseWithDefaults

`func NewServerReloadControllerJobStatusResponseWithDefaults() *ServerReloadControllerJobStatusResponse`

NewServerReloadControllerJobStatusResponseWithDefaults instantiates a new ServerReloadControllerJobStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ServerReloadControllerJobStatusResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServerReloadControllerJobStatusResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServerReloadControllerJobStatusResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServerReloadControllerJobStatusResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSuccessCount

`func (o *ServerReloadControllerJobStatusResponse) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *ServerReloadControllerJobStatusResponse) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *ServerReloadControllerJobStatusResponse) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *ServerReloadControllerJobStatusResponse) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetTotalSegmentCount

`func (o *ServerReloadControllerJobStatusResponse) GetTotalSegmentCount() int32`

GetTotalSegmentCount returns the TotalSegmentCount field if non-nil, zero value otherwise.

### GetTotalSegmentCountOk

`func (o *ServerReloadControllerJobStatusResponse) GetTotalSegmentCountOk() (*int32, bool)`

GetTotalSegmentCountOk returns a tuple with the TotalSegmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSegmentCount

`func (o *ServerReloadControllerJobStatusResponse) SetTotalSegmentCount(v int32)`

SetTotalSegmentCount sets TotalSegmentCount field to given value.

### HasTotalSegmentCount

`func (o *ServerReloadControllerJobStatusResponse) HasTotalSegmentCount() bool`

HasTotalSegmentCount returns a boolean if a field has been set.

### GetTotalServersQueried

`func (o *ServerReloadControllerJobStatusResponse) GetTotalServersQueried() int32`

GetTotalServersQueried returns the TotalServersQueried field if non-nil, zero value otherwise.

### GetTotalServersQueriedOk

`func (o *ServerReloadControllerJobStatusResponse) GetTotalServersQueriedOk() (*int32, bool)`

GetTotalServersQueriedOk returns a tuple with the TotalServersQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServersQueried

`func (o *ServerReloadControllerJobStatusResponse) SetTotalServersQueried(v int32)`

SetTotalServersQueried sets TotalServersQueried field to given value.

### HasTotalServersQueried

`func (o *ServerReloadControllerJobStatusResponse) HasTotalServersQueried() bool`

HasTotalServersQueried returns a boolean if a field has been set.

### GetTotalServerCallsFailed

`func (o *ServerReloadControllerJobStatusResponse) GetTotalServerCallsFailed() int32`

GetTotalServerCallsFailed returns the TotalServerCallsFailed field if non-nil, zero value otherwise.

### GetTotalServerCallsFailedOk

`func (o *ServerReloadControllerJobStatusResponse) GetTotalServerCallsFailedOk() (*int32, bool)`

GetTotalServerCallsFailedOk returns a tuple with the TotalServerCallsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServerCallsFailed

`func (o *ServerReloadControllerJobStatusResponse) SetTotalServerCallsFailed(v int32)`

SetTotalServerCallsFailed sets TotalServerCallsFailed field to given value.

### HasTotalServerCallsFailed

`func (o *ServerReloadControllerJobStatusResponse) HasTotalServerCallsFailed() bool`

HasTotalServerCallsFailed returns a boolean if a field has been set.

### GetTimeElapsedInMinutes

`func (o *ServerReloadControllerJobStatusResponse) GetTimeElapsedInMinutes() float64`

GetTimeElapsedInMinutes returns the TimeElapsedInMinutes field if non-nil, zero value otherwise.

### GetTimeElapsedInMinutesOk

`func (o *ServerReloadControllerJobStatusResponse) GetTimeElapsedInMinutesOk() (*float64, bool)`

GetTimeElapsedInMinutesOk returns a tuple with the TimeElapsedInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeElapsedInMinutes

`func (o *ServerReloadControllerJobStatusResponse) SetTimeElapsedInMinutes(v float64)`

SetTimeElapsedInMinutes sets TimeElapsedInMinutes field to given value.

### HasTimeElapsedInMinutes

`func (o *ServerReloadControllerJobStatusResponse) HasTimeElapsedInMinutes() bool`

HasTimeElapsedInMinutes returns a boolean if a field has been set.

### GetEstimatedTimeRemainingInMinutes

`func (o *ServerReloadControllerJobStatusResponse) GetEstimatedTimeRemainingInMinutes() float64`

GetEstimatedTimeRemainingInMinutes returns the EstimatedTimeRemainingInMinutes field if non-nil, zero value otherwise.

### GetEstimatedTimeRemainingInMinutesOk

`func (o *ServerReloadControllerJobStatusResponse) GetEstimatedTimeRemainingInMinutesOk() (*float64, bool)`

GetEstimatedTimeRemainingInMinutesOk returns a tuple with the EstimatedTimeRemainingInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeRemainingInMinutes

`func (o *ServerReloadControllerJobStatusResponse) SetEstimatedTimeRemainingInMinutes(v float64)`

SetEstimatedTimeRemainingInMinutes sets EstimatedTimeRemainingInMinutes field to given value.

### HasEstimatedTimeRemainingInMinutes

`func (o *ServerReloadControllerJobStatusResponse) HasEstimatedTimeRemainingInMinutes() bool`

HasEstimatedTimeRemainingInMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


