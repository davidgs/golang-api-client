# SegmentState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdealState** | Pointer to **string** |  | [optional] [readonly] 
**ExternalView** | Pointer to **string** |  | [optional] [readonly] 
**SegmentSize** | Pointer to **string** |  | [optional] [readonly] 
**ConsumerInfo** | Pointer to [**SegmentConsumerInfo**](SegmentConsumerInfo.md) |  | [optional] 
**ErrorInfo** | Pointer to [**SegmentErrorInfo**](SegmentErrorInfo.md) |  | [optional] 

## Methods

### NewSegmentState

`func NewSegmentState() *SegmentState`

NewSegmentState instantiates a new SegmentState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentStateWithDefaults

`func NewSegmentStateWithDefaults() *SegmentState`

NewSegmentStateWithDefaults instantiates a new SegmentState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdealState

`func (o *SegmentState) GetIdealState() string`

GetIdealState returns the IdealState field if non-nil, zero value otherwise.

### GetIdealStateOk

`func (o *SegmentState) GetIdealStateOk() (*string, bool)`

GetIdealStateOk returns a tuple with the IdealState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdealState

`func (o *SegmentState) SetIdealState(v string)`

SetIdealState sets IdealState field to given value.

### HasIdealState

`func (o *SegmentState) HasIdealState() bool`

HasIdealState returns a boolean if a field has been set.

### GetExternalView

`func (o *SegmentState) GetExternalView() string`

GetExternalView returns the ExternalView field if non-nil, zero value otherwise.

### GetExternalViewOk

`func (o *SegmentState) GetExternalViewOk() (*string, bool)`

GetExternalViewOk returns a tuple with the ExternalView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalView

`func (o *SegmentState) SetExternalView(v string)`

SetExternalView sets ExternalView field to given value.

### HasExternalView

`func (o *SegmentState) HasExternalView() bool`

HasExternalView returns a boolean if a field has been set.

### GetSegmentSize

`func (o *SegmentState) GetSegmentSize() string`

GetSegmentSize returns the SegmentSize field if non-nil, zero value otherwise.

### GetSegmentSizeOk

`func (o *SegmentState) GetSegmentSizeOk() (*string, bool)`

GetSegmentSizeOk returns a tuple with the SegmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentSize

`func (o *SegmentState) SetSegmentSize(v string)`

SetSegmentSize sets SegmentSize field to given value.

### HasSegmentSize

`func (o *SegmentState) HasSegmentSize() bool`

HasSegmentSize returns a boolean if a field has been set.

### GetConsumerInfo

`func (o *SegmentState) GetConsumerInfo() SegmentConsumerInfo`

GetConsumerInfo returns the ConsumerInfo field if non-nil, zero value otherwise.

### GetConsumerInfoOk

`func (o *SegmentState) GetConsumerInfoOk() (*SegmentConsumerInfo, bool)`

GetConsumerInfoOk returns a tuple with the ConsumerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerInfo

`func (o *SegmentState) SetConsumerInfo(v SegmentConsumerInfo)`

SetConsumerInfo sets ConsumerInfo field to given value.

### HasConsumerInfo

`func (o *SegmentState) HasConsumerInfo() bool`

HasConsumerInfo returns a boolean if a field has been set.

### GetErrorInfo

`func (o *SegmentState) GetErrorInfo() SegmentErrorInfo`

GetErrorInfo returns the ErrorInfo field if non-nil, zero value otherwise.

### GetErrorInfoOk

`func (o *SegmentState) GetErrorInfoOk() (*SegmentErrorInfo, bool)`

GetErrorInfoOk returns a tuple with the ErrorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorInfo

`func (o *SegmentState) SetErrorInfo(v SegmentErrorInfo)`

SetErrorInfo sets ErrorInfo field to given value.

### HasErrorInfo

`func (o *SegmentState) HasErrorInfo() bool`

HasErrorInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


