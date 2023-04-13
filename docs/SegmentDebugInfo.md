# SegmentDebugInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentName** | Pointer to **string** |  | [optional] [readonly] 
**ServerState** | Pointer to [**map[string]SegmentState**](SegmentState.md) |  | [optional] [readonly] 

## Methods

### NewSegmentDebugInfo

`func NewSegmentDebugInfo() *SegmentDebugInfo`

NewSegmentDebugInfo instantiates a new SegmentDebugInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentDebugInfoWithDefaults

`func NewSegmentDebugInfoWithDefaults() *SegmentDebugInfo`

NewSegmentDebugInfoWithDefaults instantiates a new SegmentDebugInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentName

`func (o *SegmentDebugInfo) GetSegmentName() string`

GetSegmentName returns the SegmentName field if non-nil, zero value otherwise.

### GetSegmentNameOk

`func (o *SegmentDebugInfo) GetSegmentNameOk() (*string, bool)`

GetSegmentNameOk returns a tuple with the SegmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentName

`func (o *SegmentDebugInfo) SetSegmentName(v string)`

SetSegmentName sets SegmentName field to given value.

### HasSegmentName

`func (o *SegmentDebugInfo) HasSegmentName() bool`

HasSegmentName returns a boolean if a field has been set.

### GetServerState

`func (o *SegmentDebugInfo) GetServerState() map[string]SegmentState`

GetServerState returns the ServerState field if non-nil, zero value otherwise.

### GetServerStateOk

`func (o *SegmentDebugInfo) GetServerStateOk() (*map[string]SegmentState, bool)`

GetServerStateOk returns a tuple with the ServerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerState

`func (o *SegmentDebugInfo) SetServerState(v map[string]SegmentState)`

SetServerState sets ServerState field to given value.

### HasServerState

`func (o *SegmentDebugInfo) HasServerState() bool`

HasServerState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


