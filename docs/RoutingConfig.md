# RoutingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingTableBuilderName** | Pointer to **string** |  | [optional] [readonly] 
**SegmentPrunerTypes** | Pointer to **[]string** |  | [optional] [readonly] 
**InstanceSelectorType** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRoutingConfig

`func NewRoutingConfig() *RoutingConfig`

NewRoutingConfig instantiates a new RoutingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingConfigWithDefaults

`func NewRoutingConfigWithDefaults() *RoutingConfig`

NewRoutingConfigWithDefaults instantiates a new RoutingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingTableBuilderName

`func (o *RoutingConfig) GetRoutingTableBuilderName() string`

GetRoutingTableBuilderName returns the RoutingTableBuilderName field if non-nil, zero value otherwise.

### GetRoutingTableBuilderNameOk

`func (o *RoutingConfig) GetRoutingTableBuilderNameOk() (*string, bool)`

GetRoutingTableBuilderNameOk returns a tuple with the RoutingTableBuilderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingTableBuilderName

`func (o *RoutingConfig) SetRoutingTableBuilderName(v string)`

SetRoutingTableBuilderName sets RoutingTableBuilderName field to given value.

### HasRoutingTableBuilderName

`func (o *RoutingConfig) HasRoutingTableBuilderName() bool`

HasRoutingTableBuilderName returns a boolean if a field has been set.

### GetSegmentPrunerTypes

`func (o *RoutingConfig) GetSegmentPrunerTypes() []string`

GetSegmentPrunerTypes returns the SegmentPrunerTypes field if non-nil, zero value otherwise.

### GetSegmentPrunerTypesOk

`func (o *RoutingConfig) GetSegmentPrunerTypesOk() (*[]string, bool)`

GetSegmentPrunerTypesOk returns a tuple with the SegmentPrunerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentPrunerTypes

`func (o *RoutingConfig) SetSegmentPrunerTypes(v []string)`

SetSegmentPrunerTypes sets SegmentPrunerTypes field to given value.

### HasSegmentPrunerTypes

`func (o *RoutingConfig) HasSegmentPrunerTypes() bool`

HasSegmentPrunerTypes returns a boolean if a field has been set.

### GetInstanceSelectorType

`func (o *RoutingConfig) GetInstanceSelectorType() string`

GetInstanceSelectorType returns the InstanceSelectorType field if non-nil, zero value otherwise.

### GetInstanceSelectorTypeOk

`func (o *RoutingConfig) GetInstanceSelectorTypeOk() (*string, bool)`

GetInstanceSelectorTypeOk returns a tuple with the InstanceSelectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSelectorType

`func (o *RoutingConfig) SetInstanceSelectorType(v string)`

SetInstanceSelectorType sets InstanceSelectorType field to given value.

### HasInstanceSelectorType

`func (o *RoutingConfig) HasInstanceSelectorType() bool`

HasInstanceSelectorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


