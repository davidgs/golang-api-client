# ConfigSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnrecognizedProperties** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigSuccessResponse

`func NewConfigSuccessResponse() *ConfigSuccessResponse`

NewConfigSuccessResponse instantiates a new ConfigSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSuccessResponseWithDefaults

`func NewConfigSuccessResponseWithDefaults() *ConfigSuccessResponse`

NewConfigSuccessResponseWithDefaults instantiates a new ConfigSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnrecognizedProperties

`func (o *ConfigSuccessResponse) GetUnrecognizedProperties() map[string]map[string]interface{}`

GetUnrecognizedProperties returns the UnrecognizedProperties field if non-nil, zero value otherwise.

### GetUnrecognizedPropertiesOk

`func (o *ConfigSuccessResponse) GetUnrecognizedPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetUnrecognizedPropertiesOk returns a tuple with the UnrecognizedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrecognizedProperties

`func (o *ConfigSuccessResponse) SetUnrecognizedProperties(v map[string]map[string]interface{})`

SetUnrecognizedProperties sets UnrecognizedProperties field to given value.

### HasUnrecognizedProperties

`func (o *ConfigSuccessResponse) HasUnrecognizedProperties() bool`

HasUnrecognizedProperties returns a boolean if a field has been set.

### GetStatus

`func (o *ConfigSuccessResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConfigSuccessResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConfigSuccessResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConfigSuccessResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


