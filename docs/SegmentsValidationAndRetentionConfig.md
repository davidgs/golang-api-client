# SegmentsValidationAndRetentionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeType** | Pointer to **string** |  | [optional] 
**SchemaName** | Pointer to **string** |  | [optional] 
**Replication** | Pointer to **string** |  | [optional] 
**RetentionTimeUnit** | Pointer to **string** |  | [optional] 
**RetentionTimeValue** | Pointer to **string** |  | [optional] 
**ReplicasPerPartition** | Pointer to **string** |  | [optional] 
**TimeColumnName** | Pointer to **string** |  | [optional] 
**CrypterClassName** | Pointer to **string** |  | [optional] 
**DeletedSegmentsRetentionPeriod** | Pointer to **string** |  | [optional] 
**SegmentAssignmentStrategy** | Pointer to **string** |  | [optional] 
**SegmentPushFrequency** | Pointer to **string** |  | [optional] 
**SegmentPushType** | Pointer to **string** |  | [optional] 
**ReplicaGroupStrategyConfig** | Pointer to [**ReplicaGroupStrategyConfig**](ReplicaGroupStrategyConfig.md) |  | [optional] 
**CompletionConfig** | Pointer to [**CompletionConfig**](CompletionConfig.md) |  | [optional] 
**PeerSegmentDownloadScheme** | Pointer to **string** |  | [optional] 
**MinimizeDataMovement** | Pointer to **bool** |  | [optional] 

## Methods

### NewSegmentsValidationAndRetentionConfig

`func NewSegmentsValidationAndRetentionConfig() *SegmentsValidationAndRetentionConfig`

NewSegmentsValidationAndRetentionConfig instantiates a new SegmentsValidationAndRetentionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentsValidationAndRetentionConfigWithDefaults

`func NewSegmentsValidationAndRetentionConfigWithDefaults() *SegmentsValidationAndRetentionConfig`

NewSegmentsValidationAndRetentionConfigWithDefaults instantiates a new SegmentsValidationAndRetentionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeType

`func (o *SegmentsValidationAndRetentionConfig) GetTimeType() string`

GetTimeType returns the TimeType field if non-nil, zero value otherwise.

### GetTimeTypeOk

`func (o *SegmentsValidationAndRetentionConfig) GetTimeTypeOk() (*string, bool)`

GetTimeTypeOk returns a tuple with the TimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeType

`func (o *SegmentsValidationAndRetentionConfig) SetTimeType(v string)`

SetTimeType sets TimeType field to given value.

### HasTimeType

`func (o *SegmentsValidationAndRetentionConfig) HasTimeType() bool`

HasTimeType returns a boolean if a field has been set.

### GetSchemaName

`func (o *SegmentsValidationAndRetentionConfig) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *SegmentsValidationAndRetentionConfig) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *SegmentsValidationAndRetentionConfig) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *SegmentsValidationAndRetentionConfig) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### GetReplication

`func (o *SegmentsValidationAndRetentionConfig) GetReplication() string`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *SegmentsValidationAndRetentionConfig) GetReplicationOk() (*string, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *SegmentsValidationAndRetentionConfig) SetReplication(v string)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *SegmentsValidationAndRetentionConfig) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetRetentionTimeUnit

`func (o *SegmentsValidationAndRetentionConfig) GetRetentionTimeUnit() string`

GetRetentionTimeUnit returns the RetentionTimeUnit field if non-nil, zero value otherwise.

### GetRetentionTimeUnitOk

`func (o *SegmentsValidationAndRetentionConfig) GetRetentionTimeUnitOk() (*string, bool)`

GetRetentionTimeUnitOk returns a tuple with the RetentionTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTimeUnit

`func (o *SegmentsValidationAndRetentionConfig) SetRetentionTimeUnit(v string)`

SetRetentionTimeUnit sets RetentionTimeUnit field to given value.

### HasRetentionTimeUnit

`func (o *SegmentsValidationAndRetentionConfig) HasRetentionTimeUnit() bool`

HasRetentionTimeUnit returns a boolean if a field has been set.

### GetRetentionTimeValue

`func (o *SegmentsValidationAndRetentionConfig) GetRetentionTimeValue() string`

GetRetentionTimeValue returns the RetentionTimeValue field if non-nil, zero value otherwise.

### GetRetentionTimeValueOk

`func (o *SegmentsValidationAndRetentionConfig) GetRetentionTimeValueOk() (*string, bool)`

GetRetentionTimeValueOk returns a tuple with the RetentionTimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTimeValue

`func (o *SegmentsValidationAndRetentionConfig) SetRetentionTimeValue(v string)`

SetRetentionTimeValue sets RetentionTimeValue field to given value.

### HasRetentionTimeValue

`func (o *SegmentsValidationAndRetentionConfig) HasRetentionTimeValue() bool`

HasRetentionTimeValue returns a boolean if a field has been set.

### GetReplicasPerPartition

`func (o *SegmentsValidationAndRetentionConfig) GetReplicasPerPartition() string`

GetReplicasPerPartition returns the ReplicasPerPartition field if non-nil, zero value otherwise.

### GetReplicasPerPartitionOk

`func (o *SegmentsValidationAndRetentionConfig) GetReplicasPerPartitionOk() (*string, bool)`

GetReplicasPerPartitionOk returns a tuple with the ReplicasPerPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicasPerPartition

`func (o *SegmentsValidationAndRetentionConfig) SetReplicasPerPartition(v string)`

SetReplicasPerPartition sets ReplicasPerPartition field to given value.

### HasReplicasPerPartition

`func (o *SegmentsValidationAndRetentionConfig) HasReplicasPerPartition() bool`

HasReplicasPerPartition returns a boolean if a field has been set.

### GetTimeColumnName

`func (o *SegmentsValidationAndRetentionConfig) GetTimeColumnName() string`

GetTimeColumnName returns the TimeColumnName field if non-nil, zero value otherwise.

### GetTimeColumnNameOk

`func (o *SegmentsValidationAndRetentionConfig) GetTimeColumnNameOk() (*string, bool)`

GetTimeColumnNameOk returns a tuple with the TimeColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeColumnName

`func (o *SegmentsValidationAndRetentionConfig) SetTimeColumnName(v string)`

SetTimeColumnName sets TimeColumnName field to given value.

### HasTimeColumnName

`func (o *SegmentsValidationAndRetentionConfig) HasTimeColumnName() bool`

HasTimeColumnName returns a boolean if a field has been set.

### GetCrypterClassName

`func (o *SegmentsValidationAndRetentionConfig) GetCrypterClassName() string`

GetCrypterClassName returns the CrypterClassName field if non-nil, zero value otherwise.

### GetCrypterClassNameOk

`func (o *SegmentsValidationAndRetentionConfig) GetCrypterClassNameOk() (*string, bool)`

GetCrypterClassNameOk returns a tuple with the CrypterClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrypterClassName

`func (o *SegmentsValidationAndRetentionConfig) SetCrypterClassName(v string)`

SetCrypterClassName sets CrypterClassName field to given value.

### HasCrypterClassName

`func (o *SegmentsValidationAndRetentionConfig) HasCrypterClassName() bool`

HasCrypterClassName returns a boolean if a field has been set.

### GetDeletedSegmentsRetentionPeriod

`func (o *SegmentsValidationAndRetentionConfig) GetDeletedSegmentsRetentionPeriod() string`

GetDeletedSegmentsRetentionPeriod returns the DeletedSegmentsRetentionPeriod field if non-nil, zero value otherwise.

### GetDeletedSegmentsRetentionPeriodOk

`func (o *SegmentsValidationAndRetentionConfig) GetDeletedSegmentsRetentionPeriodOk() (*string, bool)`

GetDeletedSegmentsRetentionPeriodOk returns a tuple with the DeletedSegmentsRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedSegmentsRetentionPeriod

`func (o *SegmentsValidationAndRetentionConfig) SetDeletedSegmentsRetentionPeriod(v string)`

SetDeletedSegmentsRetentionPeriod sets DeletedSegmentsRetentionPeriod field to given value.

### HasDeletedSegmentsRetentionPeriod

`func (o *SegmentsValidationAndRetentionConfig) HasDeletedSegmentsRetentionPeriod() bool`

HasDeletedSegmentsRetentionPeriod returns a boolean if a field has been set.

### GetSegmentAssignmentStrategy

`func (o *SegmentsValidationAndRetentionConfig) GetSegmentAssignmentStrategy() string`

GetSegmentAssignmentStrategy returns the SegmentAssignmentStrategy field if non-nil, zero value otherwise.

### GetSegmentAssignmentStrategyOk

`func (o *SegmentsValidationAndRetentionConfig) GetSegmentAssignmentStrategyOk() (*string, bool)`

GetSegmentAssignmentStrategyOk returns a tuple with the SegmentAssignmentStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentAssignmentStrategy

`func (o *SegmentsValidationAndRetentionConfig) SetSegmentAssignmentStrategy(v string)`

SetSegmentAssignmentStrategy sets SegmentAssignmentStrategy field to given value.

### HasSegmentAssignmentStrategy

`func (o *SegmentsValidationAndRetentionConfig) HasSegmentAssignmentStrategy() bool`

HasSegmentAssignmentStrategy returns a boolean if a field has been set.

### GetSegmentPushFrequency

`func (o *SegmentsValidationAndRetentionConfig) GetSegmentPushFrequency() string`

GetSegmentPushFrequency returns the SegmentPushFrequency field if non-nil, zero value otherwise.

### GetSegmentPushFrequencyOk

`func (o *SegmentsValidationAndRetentionConfig) GetSegmentPushFrequencyOk() (*string, bool)`

GetSegmentPushFrequencyOk returns a tuple with the SegmentPushFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentPushFrequency

`func (o *SegmentsValidationAndRetentionConfig) SetSegmentPushFrequency(v string)`

SetSegmentPushFrequency sets SegmentPushFrequency field to given value.

### HasSegmentPushFrequency

`func (o *SegmentsValidationAndRetentionConfig) HasSegmentPushFrequency() bool`

HasSegmentPushFrequency returns a boolean if a field has been set.

### GetSegmentPushType

`func (o *SegmentsValidationAndRetentionConfig) GetSegmentPushType() string`

GetSegmentPushType returns the SegmentPushType field if non-nil, zero value otherwise.

### GetSegmentPushTypeOk

`func (o *SegmentsValidationAndRetentionConfig) GetSegmentPushTypeOk() (*string, bool)`

GetSegmentPushTypeOk returns a tuple with the SegmentPushType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentPushType

`func (o *SegmentsValidationAndRetentionConfig) SetSegmentPushType(v string)`

SetSegmentPushType sets SegmentPushType field to given value.

### HasSegmentPushType

`func (o *SegmentsValidationAndRetentionConfig) HasSegmentPushType() bool`

HasSegmentPushType returns a boolean if a field has been set.

### GetReplicaGroupStrategyConfig

`func (o *SegmentsValidationAndRetentionConfig) GetReplicaGroupStrategyConfig() ReplicaGroupStrategyConfig`

GetReplicaGroupStrategyConfig returns the ReplicaGroupStrategyConfig field if non-nil, zero value otherwise.

### GetReplicaGroupStrategyConfigOk

`func (o *SegmentsValidationAndRetentionConfig) GetReplicaGroupStrategyConfigOk() (*ReplicaGroupStrategyConfig, bool)`

GetReplicaGroupStrategyConfigOk returns a tuple with the ReplicaGroupStrategyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaGroupStrategyConfig

`func (o *SegmentsValidationAndRetentionConfig) SetReplicaGroupStrategyConfig(v ReplicaGroupStrategyConfig)`

SetReplicaGroupStrategyConfig sets ReplicaGroupStrategyConfig field to given value.

### HasReplicaGroupStrategyConfig

`func (o *SegmentsValidationAndRetentionConfig) HasReplicaGroupStrategyConfig() bool`

HasReplicaGroupStrategyConfig returns a boolean if a field has been set.

### GetCompletionConfig

`func (o *SegmentsValidationAndRetentionConfig) GetCompletionConfig() CompletionConfig`

GetCompletionConfig returns the CompletionConfig field if non-nil, zero value otherwise.

### GetCompletionConfigOk

`func (o *SegmentsValidationAndRetentionConfig) GetCompletionConfigOk() (*CompletionConfig, bool)`

GetCompletionConfigOk returns a tuple with the CompletionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionConfig

`func (o *SegmentsValidationAndRetentionConfig) SetCompletionConfig(v CompletionConfig)`

SetCompletionConfig sets CompletionConfig field to given value.

### HasCompletionConfig

`func (o *SegmentsValidationAndRetentionConfig) HasCompletionConfig() bool`

HasCompletionConfig returns a boolean if a field has been set.

### GetPeerSegmentDownloadScheme

`func (o *SegmentsValidationAndRetentionConfig) GetPeerSegmentDownloadScheme() string`

GetPeerSegmentDownloadScheme returns the PeerSegmentDownloadScheme field if non-nil, zero value otherwise.

### GetPeerSegmentDownloadSchemeOk

`func (o *SegmentsValidationAndRetentionConfig) GetPeerSegmentDownloadSchemeOk() (*string, bool)`

GetPeerSegmentDownloadSchemeOk returns a tuple with the PeerSegmentDownloadScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSegmentDownloadScheme

`func (o *SegmentsValidationAndRetentionConfig) SetPeerSegmentDownloadScheme(v string)`

SetPeerSegmentDownloadScheme sets PeerSegmentDownloadScheme field to given value.

### HasPeerSegmentDownloadScheme

`func (o *SegmentsValidationAndRetentionConfig) HasPeerSegmentDownloadScheme() bool`

HasPeerSegmentDownloadScheme returns a boolean if a field has been set.

### GetMinimizeDataMovement

`func (o *SegmentsValidationAndRetentionConfig) GetMinimizeDataMovement() bool`

GetMinimizeDataMovement returns the MinimizeDataMovement field if non-nil, zero value otherwise.

### GetMinimizeDataMovementOk

`func (o *SegmentsValidationAndRetentionConfig) GetMinimizeDataMovementOk() (*bool, bool)`

GetMinimizeDataMovementOk returns a tuple with the MinimizeDataMovement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimizeDataMovement

`func (o *SegmentsValidationAndRetentionConfig) SetMinimizeDataMovement(v bool)`

SetMinimizeDataMovement sets MinimizeDataMovement field to given value.

### HasMinimizeDataMovement

`func (o *SegmentsValidationAndRetentionConfig) HasMinimizeDataMovement() bool`

HasMinimizeDataMovement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


