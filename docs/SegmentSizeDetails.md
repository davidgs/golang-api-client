# SegmentSizeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportedSizeInBytes** | Pointer to **int64** |  | [optional] 
**EstimatedSizeInBytes** | Pointer to **int64** |  | [optional] 
**ServerInfo** | Pointer to [**map[string]SegmentSizeInfo**](SegmentSizeInfo.md) |  | [optional] 

## Methods

### NewSegmentSizeDetails

`func NewSegmentSizeDetails() *SegmentSizeDetails`

NewSegmentSizeDetails instantiates a new SegmentSizeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentSizeDetailsWithDefaults

`func NewSegmentSizeDetailsWithDefaults() *SegmentSizeDetails`

NewSegmentSizeDetailsWithDefaults instantiates a new SegmentSizeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportedSizeInBytes

`func (o *SegmentSizeDetails) GetReportedSizeInBytes() int64`

GetReportedSizeInBytes returns the ReportedSizeInBytes field if non-nil, zero value otherwise.

### GetReportedSizeInBytesOk

`func (o *SegmentSizeDetails) GetReportedSizeInBytesOk() (*int64, bool)`

GetReportedSizeInBytesOk returns a tuple with the ReportedSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedSizeInBytes

`func (o *SegmentSizeDetails) SetReportedSizeInBytes(v int64)`

SetReportedSizeInBytes sets ReportedSizeInBytes field to given value.

### HasReportedSizeInBytes

`func (o *SegmentSizeDetails) HasReportedSizeInBytes() bool`

HasReportedSizeInBytes returns a boolean if a field has been set.

### GetEstimatedSizeInBytes

`func (o *SegmentSizeDetails) GetEstimatedSizeInBytes() int64`

GetEstimatedSizeInBytes returns the EstimatedSizeInBytes field if non-nil, zero value otherwise.

### GetEstimatedSizeInBytesOk

`func (o *SegmentSizeDetails) GetEstimatedSizeInBytesOk() (*int64, bool)`

GetEstimatedSizeInBytesOk returns a tuple with the EstimatedSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSizeInBytes

`func (o *SegmentSizeDetails) SetEstimatedSizeInBytes(v int64)`

SetEstimatedSizeInBytes sets EstimatedSizeInBytes field to given value.

### HasEstimatedSizeInBytes

`func (o *SegmentSizeDetails) HasEstimatedSizeInBytes() bool`

HasEstimatedSizeInBytes returns a boolean if a field has been set.

### GetServerInfo

`func (o *SegmentSizeDetails) GetServerInfo() map[string]SegmentSizeInfo`

GetServerInfo returns the ServerInfo field if non-nil, zero value otherwise.

### GetServerInfoOk

`func (o *SegmentSizeDetails) GetServerInfoOk() (*map[string]SegmentSizeInfo, bool)`

GetServerInfoOk returns a tuple with the ServerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInfo

`func (o *SegmentSizeDetails) SetServerInfo(v map[string]SegmentSizeInfo)`

SetServerInfo sets ServerInfo field to given value.

### HasServerInfo

`func (o *SegmentSizeDetails) HasServerInfo() bool`

HasServerInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


