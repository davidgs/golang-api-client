# SegmentSizeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentName** | Pointer to **string** |  | [optional] [readonly] 
**DiskSizeInBytes** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewSegmentSizeInfo

`func NewSegmentSizeInfo() *SegmentSizeInfo`

NewSegmentSizeInfo instantiates a new SegmentSizeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentSizeInfoWithDefaults

`func NewSegmentSizeInfoWithDefaults() *SegmentSizeInfo`

NewSegmentSizeInfoWithDefaults instantiates a new SegmentSizeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentName

`func (o *SegmentSizeInfo) GetSegmentName() string`

GetSegmentName returns the SegmentName field if non-nil, zero value otherwise.

### GetSegmentNameOk

`func (o *SegmentSizeInfo) GetSegmentNameOk() (*string, bool)`

GetSegmentNameOk returns a tuple with the SegmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentName

`func (o *SegmentSizeInfo) SetSegmentName(v string)`

SetSegmentName sets SegmentName field to given value.

### HasSegmentName

`func (o *SegmentSizeInfo) HasSegmentName() bool`

HasSegmentName returns a boolean if a field has been set.

### GetDiskSizeInBytes

`func (o *SegmentSizeInfo) GetDiskSizeInBytes() int64`

GetDiskSizeInBytes returns the DiskSizeInBytes field if non-nil, zero value otherwise.

### GetDiskSizeInBytesOk

`func (o *SegmentSizeInfo) GetDiskSizeInBytesOk() (*int64, bool)`

GetDiskSizeInBytesOk returns a tuple with the DiskSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeInBytes

`func (o *SegmentSizeInfo) SetDiskSizeInBytes(v int64)`

SetDiskSizeInBytes sets DiskSizeInBytes field to given value.

### HasDiskSizeInBytes

`func (o *SegmentSizeInfo) HasDiskSizeInBytes() bool`

HasDiskSizeInBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


