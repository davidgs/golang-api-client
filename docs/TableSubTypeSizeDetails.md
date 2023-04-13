# TableSubTypeSizeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportedSizeInBytes** | Pointer to **int64** |  | [optional] 
**EstimatedSizeInBytes** | Pointer to **int64** |  | [optional] 
**MissingSegments** | Pointer to **int32** |  | [optional] 
**Segments** | Pointer to [**map[string]SegmentSizeDetails**](SegmentSizeDetails.md) |  | [optional] 

## Methods

### NewTableSubTypeSizeDetails

`func NewTableSubTypeSizeDetails() *TableSubTypeSizeDetails`

NewTableSubTypeSizeDetails instantiates a new TableSubTypeSizeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableSubTypeSizeDetailsWithDefaults

`func NewTableSubTypeSizeDetailsWithDefaults() *TableSubTypeSizeDetails`

NewTableSubTypeSizeDetailsWithDefaults instantiates a new TableSubTypeSizeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportedSizeInBytes

`func (o *TableSubTypeSizeDetails) GetReportedSizeInBytes() int64`

GetReportedSizeInBytes returns the ReportedSizeInBytes field if non-nil, zero value otherwise.

### GetReportedSizeInBytesOk

`func (o *TableSubTypeSizeDetails) GetReportedSizeInBytesOk() (*int64, bool)`

GetReportedSizeInBytesOk returns a tuple with the ReportedSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedSizeInBytes

`func (o *TableSubTypeSizeDetails) SetReportedSizeInBytes(v int64)`

SetReportedSizeInBytes sets ReportedSizeInBytes field to given value.

### HasReportedSizeInBytes

`func (o *TableSubTypeSizeDetails) HasReportedSizeInBytes() bool`

HasReportedSizeInBytes returns a boolean if a field has been set.

### GetEstimatedSizeInBytes

`func (o *TableSubTypeSizeDetails) GetEstimatedSizeInBytes() int64`

GetEstimatedSizeInBytes returns the EstimatedSizeInBytes field if non-nil, zero value otherwise.

### GetEstimatedSizeInBytesOk

`func (o *TableSubTypeSizeDetails) GetEstimatedSizeInBytesOk() (*int64, bool)`

GetEstimatedSizeInBytesOk returns a tuple with the EstimatedSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSizeInBytes

`func (o *TableSubTypeSizeDetails) SetEstimatedSizeInBytes(v int64)`

SetEstimatedSizeInBytes sets EstimatedSizeInBytes field to given value.

### HasEstimatedSizeInBytes

`func (o *TableSubTypeSizeDetails) HasEstimatedSizeInBytes() bool`

HasEstimatedSizeInBytes returns a boolean if a field has been set.

### GetMissingSegments

`func (o *TableSubTypeSizeDetails) GetMissingSegments() int32`

GetMissingSegments returns the MissingSegments field if non-nil, zero value otherwise.

### GetMissingSegmentsOk

`func (o *TableSubTypeSizeDetails) GetMissingSegmentsOk() (*int32, bool)`

GetMissingSegmentsOk returns a tuple with the MissingSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingSegments

`func (o *TableSubTypeSizeDetails) SetMissingSegments(v int32)`

SetMissingSegments sets MissingSegments field to given value.

### HasMissingSegments

`func (o *TableSubTypeSizeDetails) HasMissingSegments() bool`

HasMissingSegments returns a boolean if a field has been set.

### GetSegments

`func (o *TableSubTypeSizeDetails) GetSegments() map[string]SegmentSizeDetails`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *TableSubTypeSizeDetails) GetSegmentsOk() (*map[string]SegmentSizeDetails, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *TableSubTypeSizeDetails) SetSegments(v map[string]SegmentSizeDetails)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *TableSubTypeSizeDetails) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


