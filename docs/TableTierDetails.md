# TableTierDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | Pointer to **string** | Name of table to look for segment storage tiers | [optional] [readonly] 
**SegmentTiers** | Pointer to **map[string]map[string]string** | Storage tiers of segments for the given table | [optional] [readonly] 

## Methods

### NewTableTierDetails

`func NewTableTierDetails() *TableTierDetails`

NewTableTierDetails instantiates a new TableTierDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableTierDetailsWithDefaults

`func NewTableTierDetailsWithDefaults() *TableTierDetails`

NewTableTierDetailsWithDefaults instantiates a new TableTierDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *TableTierDetails) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableTierDetails) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableTierDetails) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *TableTierDetails) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetSegmentTiers

`func (o *TableTierDetails) GetSegmentTiers() map[string]map[string]string`

GetSegmentTiers returns the SegmentTiers field if non-nil, zero value otherwise.

### GetSegmentTiersOk

`func (o *TableTierDetails) GetSegmentTiersOk() (*map[string]map[string]string, bool)`

GetSegmentTiersOk returns a tuple with the SegmentTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentTiers

`func (o *TableTierDetails) SetSegmentTiers(v map[string]map[string]string)`

SetSegmentTiers sets SegmentTiers field to given value.

### HasSegmentTiers

`func (o *TableTierDetails) HasSegmentTiers() bool`

HasSegmentTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


