# TierConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [readonly] 
**SegmentSelectorType** | **string** |  | [readonly] 
**SegmentAge** | Pointer to **string** |  | [optional] [readonly] 
**SegmentList** | Pointer to **[]string** |  | [optional] [readonly] 
**StorageType** | **string** |  | [readonly] 
**ServerTag** | Pointer to **string** |  | [optional] [readonly] 
**TierBackend** | Pointer to **string** |  | [optional] [readonly] 
**TierBackendProperties** | Pointer to **map[string]string** |  | [optional] [readonly] 

## Methods

### NewTierConfig

`func NewTierConfig(name string, segmentSelectorType string, storageType string, ) *TierConfig`

NewTierConfig instantiates a new TierConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierConfigWithDefaults

`func NewTierConfigWithDefaults() *TierConfig`

NewTierConfigWithDefaults instantiates a new TierConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TierConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TierConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TierConfig) SetName(v string)`

SetName sets Name field to given value.


### GetSegmentSelectorType

`func (o *TierConfig) GetSegmentSelectorType() string`

GetSegmentSelectorType returns the SegmentSelectorType field if non-nil, zero value otherwise.

### GetSegmentSelectorTypeOk

`func (o *TierConfig) GetSegmentSelectorTypeOk() (*string, bool)`

GetSegmentSelectorTypeOk returns a tuple with the SegmentSelectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentSelectorType

`func (o *TierConfig) SetSegmentSelectorType(v string)`

SetSegmentSelectorType sets SegmentSelectorType field to given value.


### GetSegmentAge

`func (o *TierConfig) GetSegmentAge() string`

GetSegmentAge returns the SegmentAge field if non-nil, zero value otherwise.

### GetSegmentAgeOk

`func (o *TierConfig) GetSegmentAgeOk() (*string, bool)`

GetSegmentAgeOk returns a tuple with the SegmentAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentAge

`func (o *TierConfig) SetSegmentAge(v string)`

SetSegmentAge sets SegmentAge field to given value.

### HasSegmentAge

`func (o *TierConfig) HasSegmentAge() bool`

HasSegmentAge returns a boolean if a field has been set.

### GetSegmentList

`func (o *TierConfig) GetSegmentList() []string`

GetSegmentList returns the SegmentList field if non-nil, zero value otherwise.

### GetSegmentListOk

`func (o *TierConfig) GetSegmentListOk() (*[]string, bool)`

GetSegmentListOk returns a tuple with the SegmentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentList

`func (o *TierConfig) SetSegmentList(v []string)`

SetSegmentList sets SegmentList field to given value.

### HasSegmentList

`func (o *TierConfig) HasSegmentList() bool`

HasSegmentList returns a boolean if a field has been set.

### GetStorageType

`func (o *TierConfig) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *TierConfig) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *TierConfig) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetServerTag

`func (o *TierConfig) GetServerTag() string`

GetServerTag returns the ServerTag field if non-nil, zero value otherwise.

### GetServerTagOk

`func (o *TierConfig) GetServerTagOk() (*string, bool)`

GetServerTagOk returns a tuple with the ServerTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTag

`func (o *TierConfig) SetServerTag(v string)`

SetServerTag sets ServerTag field to given value.

### HasServerTag

`func (o *TierConfig) HasServerTag() bool`

HasServerTag returns a boolean if a field has been set.

### GetTierBackend

`func (o *TierConfig) GetTierBackend() string`

GetTierBackend returns the TierBackend field if non-nil, zero value otherwise.

### GetTierBackendOk

`func (o *TierConfig) GetTierBackendOk() (*string, bool)`

GetTierBackendOk returns a tuple with the TierBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierBackend

`func (o *TierConfig) SetTierBackend(v string)`

SetTierBackend sets TierBackend field to given value.

### HasTierBackend

`func (o *TierConfig) HasTierBackend() bool`

HasTierBackend returns a boolean if a field has been set.

### GetTierBackendProperties

`func (o *TierConfig) GetTierBackendProperties() map[string]string`

GetTierBackendProperties returns the TierBackendProperties field if non-nil, zero value otherwise.

### GetTierBackendPropertiesOk

`func (o *TierConfig) GetTierBackendPropertiesOk() (*map[string]string, bool)`

GetTierBackendPropertiesOk returns a tuple with the TierBackendProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierBackendProperties

`func (o *TierConfig) SetTierBackendProperties(v map[string]string)`

SetTierBackendProperties sets TierBackendProperties field to given value.

### HasTierBackendProperties

`func (o *TierConfig) HasTierBackendProperties() bool`

HasTierBackendProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


