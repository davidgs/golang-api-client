# TableConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | Pointer to **string** |  | [optional] [readonly] 
**TableType** | Pointer to **string** |  | [optional] [readonly] 
**SegmentsConfig** | Pointer to [**SegmentsValidationAndRetentionConfig**](SegmentsValidationAndRetentionConfig.md) |  | [optional] 
**Tenants** | Pointer to [**TenantConfig**](TenantConfig.md) |  | [optional] 
**TableIndexConfig** | Pointer to [**IndexingConfig**](IndexingConfig.md) |  | [optional] 
**Metadata** | Pointer to [**TableCustomConfig**](TableCustomConfig.md) |  | [optional] 
**Quota** | Pointer to [**QuotaConfig**](QuotaConfig.md) |  | [optional] 
**Task** | Pointer to [**TableTaskConfig**](TableTaskConfig.md) |  | [optional] 
**Routing** | Pointer to [**RoutingConfig**](RoutingConfig.md) |  | [optional] 
**Query** | Pointer to [**QueryConfig**](QueryConfig.md) |  | [optional] 
**InstanceAssignmentConfigMap** | Pointer to [**map[string]InstanceAssignmentConfig**](InstanceAssignmentConfig.md) |  | [optional] 
**FieldConfigList** | Pointer to [**[]FieldConfig**](FieldConfig.md) |  | [optional] 
**UpsertConfig** | Pointer to [**UpsertConfig**](UpsertConfig.md) |  | [optional] 
**DedupConfig** | Pointer to [**DedupConfig**](DedupConfig.md) |  | [optional] 
**DimensionTableConfig** | Pointer to [**DimensionTableConfig**](DimensionTableConfig.md) |  | [optional] 
**IngestionConfig** | Pointer to [**IngestionConfig**](IngestionConfig.md) |  | [optional] 
**TierConfigs** | Pointer to [**[]TierConfig**](TierConfig.md) |  | [optional] 
**IsDimTable** | Pointer to **bool** |  | [optional] [readonly] 
**TunerConfigs** | Pointer to [**[]TunerConfig**](TunerConfig.md) |  | [optional] 
**InstancePartitionsMap** | Pointer to **map[string]string** |  | [optional] 
**SegmentAssignmentConfigMap** | Pointer to [**map[string]SegmentAssignmentConfig**](SegmentAssignmentConfig.md) |  | [optional] 

## Methods

### NewTableConfig

`func NewTableConfig() *TableConfig`

NewTableConfig instantiates a new TableConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableConfigWithDefaults

`func NewTableConfigWithDefaults() *TableConfig`

NewTableConfigWithDefaults instantiates a new TableConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *TableConfig) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableConfig) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableConfig) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *TableConfig) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetTableType

`func (o *TableConfig) GetTableType() string`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *TableConfig) GetTableTypeOk() (*string, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *TableConfig) SetTableType(v string)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *TableConfig) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### GetSegmentsConfig

`func (o *TableConfig) GetSegmentsConfig() SegmentsValidationAndRetentionConfig`

GetSegmentsConfig returns the SegmentsConfig field if non-nil, zero value otherwise.

### GetSegmentsConfigOk

`func (o *TableConfig) GetSegmentsConfigOk() (*SegmentsValidationAndRetentionConfig, bool)`

GetSegmentsConfigOk returns a tuple with the SegmentsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentsConfig

`func (o *TableConfig) SetSegmentsConfig(v SegmentsValidationAndRetentionConfig)`

SetSegmentsConfig sets SegmentsConfig field to given value.

### HasSegmentsConfig

`func (o *TableConfig) HasSegmentsConfig() bool`

HasSegmentsConfig returns a boolean if a field has been set.

### GetTenants

`func (o *TableConfig) GetTenants() TenantConfig`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *TableConfig) GetTenantsOk() (*TenantConfig, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *TableConfig) SetTenants(v TenantConfig)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *TableConfig) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetTableIndexConfig

`func (o *TableConfig) GetTableIndexConfig() IndexingConfig`

GetTableIndexConfig returns the TableIndexConfig field if non-nil, zero value otherwise.

### GetTableIndexConfigOk

`func (o *TableConfig) GetTableIndexConfigOk() (*IndexingConfig, bool)`

GetTableIndexConfigOk returns a tuple with the TableIndexConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableIndexConfig

`func (o *TableConfig) SetTableIndexConfig(v IndexingConfig)`

SetTableIndexConfig sets TableIndexConfig field to given value.

### HasTableIndexConfig

`func (o *TableConfig) HasTableIndexConfig() bool`

HasTableIndexConfig returns a boolean if a field has been set.

### GetMetadata

`func (o *TableConfig) GetMetadata() TableCustomConfig`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TableConfig) GetMetadataOk() (*TableCustomConfig, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TableConfig) SetMetadata(v TableCustomConfig)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TableConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQuota

`func (o *TableConfig) GetQuota() QuotaConfig`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *TableConfig) GetQuotaOk() (*QuotaConfig, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *TableConfig) SetQuota(v QuotaConfig)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *TableConfig) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetTask

`func (o *TableConfig) GetTask() TableTaskConfig`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TableConfig) GetTaskOk() (*TableTaskConfig, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TableConfig) SetTask(v TableTaskConfig)`

SetTask sets Task field to given value.

### HasTask

`func (o *TableConfig) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetRouting

`func (o *TableConfig) GetRouting() RoutingConfig`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *TableConfig) GetRoutingOk() (*RoutingConfig, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *TableConfig) SetRouting(v RoutingConfig)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *TableConfig) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetQuery

`func (o *TableConfig) GetQuery() QueryConfig`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TableConfig) GetQueryOk() (*QueryConfig, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TableConfig) SetQuery(v QueryConfig)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *TableConfig) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetInstanceAssignmentConfigMap

`func (o *TableConfig) GetInstanceAssignmentConfigMap() map[string]InstanceAssignmentConfig`

GetInstanceAssignmentConfigMap returns the InstanceAssignmentConfigMap field if non-nil, zero value otherwise.

### GetInstanceAssignmentConfigMapOk

`func (o *TableConfig) GetInstanceAssignmentConfigMapOk() (*map[string]InstanceAssignmentConfig, bool)`

GetInstanceAssignmentConfigMapOk returns a tuple with the InstanceAssignmentConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAssignmentConfigMap

`func (o *TableConfig) SetInstanceAssignmentConfigMap(v map[string]InstanceAssignmentConfig)`

SetInstanceAssignmentConfigMap sets InstanceAssignmentConfigMap field to given value.

### HasInstanceAssignmentConfigMap

`func (o *TableConfig) HasInstanceAssignmentConfigMap() bool`

HasInstanceAssignmentConfigMap returns a boolean if a field has been set.

### GetFieldConfigList

`func (o *TableConfig) GetFieldConfigList() []FieldConfig`

GetFieldConfigList returns the FieldConfigList field if non-nil, zero value otherwise.

### GetFieldConfigListOk

`func (o *TableConfig) GetFieldConfigListOk() (*[]FieldConfig, bool)`

GetFieldConfigListOk returns a tuple with the FieldConfigList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldConfigList

`func (o *TableConfig) SetFieldConfigList(v []FieldConfig)`

SetFieldConfigList sets FieldConfigList field to given value.

### HasFieldConfigList

`func (o *TableConfig) HasFieldConfigList() bool`

HasFieldConfigList returns a boolean if a field has been set.

### GetUpsertConfig

`func (o *TableConfig) GetUpsertConfig() UpsertConfig`

GetUpsertConfig returns the UpsertConfig field if non-nil, zero value otherwise.

### GetUpsertConfigOk

`func (o *TableConfig) GetUpsertConfigOk() (*UpsertConfig, bool)`

GetUpsertConfigOk returns a tuple with the UpsertConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertConfig

`func (o *TableConfig) SetUpsertConfig(v UpsertConfig)`

SetUpsertConfig sets UpsertConfig field to given value.

### HasUpsertConfig

`func (o *TableConfig) HasUpsertConfig() bool`

HasUpsertConfig returns a boolean if a field has been set.

### GetDedupConfig

`func (o *TableConfig) GetDedupConfig() DedupConfig`

GetDedupConfig returns the DedupConfig field if non-nil, zero value otherwise.

### GetDedupConfigOk

`func (o *TableConfig) GetDedupConfigOk() (*DedupConfig, bool)`

GetDedupConfigOk returns a tuple with the DedupConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupConfig

`func (o *TableConfig) SetDedupConfig(v DedupConfig)`

SetDedupConfig sets DedupConfig field to given value.

### HasDedupConfig

`func (o *TableConfig) HasDedupConfig() bool`

HasDedupConfig returns a boolean if a field has been set.

### GetDimensionTableConfig

`func (o *TableConfig) GetDimensionTableConfig() DimensionTableConfig`

GetDimensionTableConfig returns the DimensionTableConfig field if non-nil, zero value otherwise.

### GetDimensionTableConfigOk

`func (o *TableConfig) GetDimensionTableConfigOk() (*DimensionTableConfig, bool)`

GetDimensionTableConfigOk returns a tuple with the DimensionTableConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionTableConfig

`func (o *TableConfig) SetDimensionTableConfig(v DimensionTableConfig)`

SetDimensionTableConfig sets DimensionTableConfig field to given value.

### HasDimensionTableConfig

`func (o *TableConfig) HasDimensionTableConfig() bool`

HasDimensionTableConfig returns a boolean if a field has been set.

### GetIngestionConfig

`func (o *TableConfig) GetIngestionConfig() IngestionConfig`

GetIngestionConfig returns the IngestionConfig field if non-nil, zero value otherwise.

### GetIngestionConfigOk

`func (o *TableConfig) GetIngestionConfigOk() (*IngestionConfig, bool)`

GetIngestionConfigOk returns a tuple with the IngestionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionConfig

`func (o *TableConfig) SetIngestionConfig(v IngestionConfig)`

SetIngestionConfig sets IngestionConfig field to given value.

### HasIngestionConfig

`func (o *TableConfig) HasIngestionConfig() bool`

HasIngestionConfig returns a boolean if a field has been set.

### GetTierConfigs

`func (o *TableConfig) GetTierConfigs() []TierConfig`

GetTierConfigs returns the TierConfigs field if non-nil, zero value otherwise.

### GetTierConfigsOk

`func (o *TableConfig) GetTierConfigsOk() (*[]TierConfig, bool)`

GetTierConfigsOk returns a tuple with the TierConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierConfigs

`func (o *TableConfig) SetTierConfigs(v []TierConfig)`

SetTierConfigs sets TierConfigs field to given value.

### HasTierConfigs

`func (o *TableConfig) HasTierConfigs() bool`

HasTierConfigs returns a boolean if a field has been set.

### GetIsDimTable

`func (o *TableConfig) GetIsDimTable() bool`

GetIsDimTable returns the IsDimTable field if non-nil, zero value otherwise.

### GetIsDimTableOk

`func (o *TableConfig) GetIsDimTableOk() (*bool, bool)`

GetIsDimTableOk returns a tuple with the IsDimTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDimTable

`func (o *TableConfig) SetIsDimTable(v bool)`

SetIsDimTable sets IsDimTable field to given value.

### HasIsDimTable

`func (o *TableConfig) HasIsDimTable() bool`

HasIsDimTable returns a boolean if a field has been set.

### GetTunerConfigs

`func (o *TableConfig) GetTunerConfigs() []TunerConfig`

GetTunerConfigs returns the TunerConfigs field if non-nil, zero value otherwise.

### GetTunerConfigsOk

`func (o *TableConfig) GetTunerConfigsOk() (*[]TunerConfig, bool)`

GetTunerConfigsOk returns a tuple with the TunerConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunerConfigs

`func (o *TableConfig) SetTunerConfigs(v []TunerConfig)`

SetTunerConfigs sets TunerConfigs field to given value.

### HasTunerConfigs

`func (o *TableConfig) HasTunerConfigs() bool`

HasTunerConfigs returns a boolean if a field has been set.

### GetInstancePartitionsMap

`func (o *TableConfig) GetInstancePartitionsMap() map[string]string`

GetInstancePartitionsMap returns the InstancePartitionsMap field if non-nil, zero value otherwise.

### GetInstancePartitionsMapOk

`func (o *TableConfig) GetInstancePartitionsMapOk() (*map[string]string, bool)`

GetInstancePartitionsMapOk returns a tuple with the InstancePartitionsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePartitionsMap

`func (o *TableConfig) SetInstancePartitionsMap(v map[string]string)`

SetInstancePartitionsMap sets InstancePartitionsMap field to given value.

### HasInstancePartitionsMap

`func (o *TableConfig) HasInstancePartitionsMap() bool`

HasInstancePartitionsMap returns a boolean if a field has been set.

### GetSegmentAssignmentConfigMap

`func (o *TableConfig) GetSegmentAssignmentConfigMap() map[string]SegmentAssignmentConfig`

GetSegmentAssignmentConfigMap returns the SegmentAssignmentConfigMap field if non-nil, zero value otherwise.

### GetSegmentAssignmentConfigMapOk

`func (o *TableConfig) GetSegmentAssignmentConfigMapOk() (*map[string]SegmentAssignmentConfig, bool)`

GetSegmentAssignmentConfigMapOk returns a tuple with the SegmentAssignmentConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentAssignmentConfigMap

`func (o *TableConfig) SetSegmentAssignmentConfigMap(v map[string]SegmentAssignmentConfig)`

SetSegmentAssignmentConfigMap sets SegmentAssignmentConfigMap field to given value.

### HasSegmentAssignmentConfigMap

`func (o *TableConfig) HasSegmentAssignmentConfigMap() bool`

HasSegmentAssignmentConfigMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


