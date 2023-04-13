# TableSizeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | Pointer to **string** |  | [optional] 
**ReportedSizeInBytes** | Pointer to **int64** |  | [optional] 
**EstimatedSizeInBytes** | Pointer to **int64** |  | [optional] 
**OfflineSegments** | Pointer to [**TableSubTypeSizeDetails**](TableSubTypeSizeDetails.md) |  | [optional] 
**RealtimeSegments** | Pointer to [**TableSubTypeSizeDetails**](TableSubTypeSizeDetails.md) |  | [optional] 

## Methods

### NewTableSizeDetails

`func NewTableSizeDetails() *TableSizeDetails`

NewTableSizeDetails instantiates a new TableSizeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableSizeDetailsWithDefaults

`func NewTableSizeDetailsWithDefaults() *TableSizeDetails`

NewTableSizeDetailsWithDefaults instantiates a new TableSizeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *TableSizeDetails) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableSizeDetails) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableSizeDetails) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *TableSizeDetails) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetReportedSizeInBytes

`func (o *TableSizeDetails) GetReportedSizeInBytes() int64`

GetReportedSizeInBytes returns the ReportedSizeInBytes field if non-nil, zero value otherwise.

### GetReportedSizeInBytesOk

`func (o *TableSizeDetails) GetReportedSizeInBytesOk() (*int64, bool)`

GetReportedSizeInBytesOk returns a tuple with the ReportedSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedSizeInBytes

`func (o *TableSizeDetails) SetReportedSizeInBytes(v int64)`

SetReportedSizeInBytes sets ReportedSizeInBytes field to given value.

### HasReportedSizeInBytes

`func (o *TableSizeDetails) HasReportedSizeInBytes() bool`

HasReportedSizeInBytes returns a boolean if a field has been set.

### GetEstimatedSizeInBytes

`func (o *TableSizeDetails) GetEstimatedSizeInBytes() int64`

GetEstimatedSizeInBytes returns the EstimatedSizeInBytes field if non-nil, zero value otherwise.

### GetEstimatedSizeInBytesOk

`func (o *TableSizeDetails) GetEstimatedSizeInBytesOk() (*int64, bool)`

GetEstimatedSizeInBytesOk returns a tuple with the EstimatedSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSizeInBytes

`func (o *TableSizeDetails) SetEstimatedSizeInBytes(v int64)`

SetEstimatedSizeInBytes sets EstimatedSizeInBytes field to given value.

### HasEstimatedSizeInBytes

`func (o *TableSizeDetails) HasEstimatedSizeInBytes() bool`

HasEstimatedSizeInBytes returns a boolean if a field has been set.

### GetOfflineSegments

`func (o *TableSizeDetails) GetOfflineSegments() TableSubTypeSizeDetails`

GetOfflineSegments returns the OfflineSegments field if non-nil, zero value otherwise.

### GetOfflineSegmentsOk

`func (o *TableSizeDetails) GetOfflineSegmentsOk() (*TableSubTypeSizeDetails, bool)`

GetOfflineSegmentsOk returns a tuple with the OfflineSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineSegments

`func (o *TableSizeDetails) SetOfflineSegments(v TableSubTypeSizeDetails)`

SetOfflineSegments sets OfflineSegments field to given value.

### HasOfflineSegments

`func (o *TableSizeDetails) HasOfflineSegments() bool`

HasOfflineSegments returns a boolean if a field has been set.

### GetRealtimeSegments

`func (o *TableSizeDetails) GetRealtimeSegments() TableSubTypeSizeDetails`

GetRealtimeSegments returns the RealtimeSegments field if non-nil, zero value otherwise.

### GetRealtimeSegmentsOk

`func (o *TableSizeDetails) GetRealtimeSegmentsOk() (*TableSubTypeSizeDetails, bool)`

GetRealtimeSegmentsOk returns a tuple with the RealtimeSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtimeSegments

`func (o *TableSizeDetails) SetRealtimeSegments(v TableSubTypeSizeDetails)`

SetRealtimeSegments sets RealtimeSegments field to given value.

### HasRealtimeSegments

`func (o *TableSizeDetails) HasRealtimeSegments() bool`

HasRealtimeSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


