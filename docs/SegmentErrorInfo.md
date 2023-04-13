# SegmentErrorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** |  | [optional] [readonly] 
**StackTrace** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSegmentErrorInfo

`func NewSegmentErrorInfo() *SegmentErrorInfo`

NewSegmentErrorInfo instantiates a new SegmentErrorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentErrorInfoWithDefaults

`func NewSegmentErrorInfoWithDefaults() *SegmentErrorInfo`

NewSegmentErrorInfoWithDefaults instantiates a new SegmentErrorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *SegmentErrorInfo) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SegmentErrorInfo) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SegmentErrorInfo) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SegmentErrorInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetErrorMessage

`func (o *SegmentErrorInfo) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SegmentErrorInfo) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SegmentErrorInfo) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SegmentErrorInfo) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStackTrace

`func (o *SegmentErrorInfo) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *SegmentErrorInfo) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *SegmentErrorInfo) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *SegmentErrorInfo) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


