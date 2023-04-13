/*
Pinot Controller API

APIs for accessing Pinot Controller information

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the IndexingConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexingConfig{}

// IndexingConfig struct for IndexingConfig
type IndexingConfig struct {
	SegmentNameGeneratorType *string `json:"segmentNameGeneratorType,omitempty"`
	StreamConfigs *map[string]string `json:"streamConfigs,omitempty"`
	InvertedIndexColumns []string `json:"invertedIndexColumns,omitempty"`
	NoDictionaryColumns []string `json:"noDictionaryColumns,omitempty"`
	SegmentPartitionConfig *SegmentPartitionConfig `json:"segmentPartitionConfig,omitempty"`
	RangeIndexColumns []string `json:"rangeIndexColumns,omitempty"`
	RangeIndexVersion *int32 `json:"rangeIndexVersion,omitempty"`
	FstindexType *string `json:"fstindexType,omitempty"`
	JsonIndexColumns []string `json:"jsonIndexColumns,omitempty"`
	JsonIndexConfigs *map[string]JsonIndexConfig `json:"jsonIndexConfigs,omitempty"`
	AutoGeneratedInvertedIndex *bool `json:"autoGeneratedInvertedIndex,omitempty"`
	CreateInvertedIndexDuringSegmentGeneration *bool `json:"createInvertedIndexDuringSegmentGeneration,omitempty"`
	SortedColumn []string `json:"sortedColumn,omitempty"`
	BloomFilterColumns []string `json:"bloomFilterColumns,omitempty"`
	BloomFilterConfigs *map[string]BloomFilterConfig `json:"bloomFilterConfigs,omitempty"`
	LoadMode *string `json:"loadMode,omitempty"`
	SegmentFormatVersion *string `json:"segmentFormatVersion,omitempty"`
	ColumnMinMaxValueGeneratorMode *string `json:"columnMinMaxValueGeneratorMode,omitempty"`
	NoDictionaryConfig *map[string]string `json:"noDictionaryConfig,omitempty"`
	OnHeapDictionaryColumns []string `json:"onHeapDictionaryColumns,omitempty"`
	VarLengthDictionaryColumns []string `json:"varLengthDictionaryColumns,omitempty"`
	EnableDefaultStarTree *bool `json:"enableDefaultStarTree,omitempty"`
	StarTreeIndexConfigs []StarTreeIndexConfig `json:"starTreeIndexConfigs,omitempty"`
	EnableDynamicStarTreeCreation *bool `json:"enableDynamicStarTreeCreation,omitempty"`
	AggregateMetrics *bool `json:"aggregateMetrics,omitempty"`
	NullHandlingEnabled *bool `json:"nullHandlingEnabled,omitempty"`
	OptimizeDictionary *bool `json:"optimizeDictionary,omitempty"`
	OptimizeDictionaryForMetrics *bool `json:"optimizeDictionaryForMetrics,omitempty"`
	NoDictionarySizeRatioThreshold *float64 `json:"noDictionarySizeRatioThreshold,omitempty"`
}

// NewIndexingConfig instantiates a new IndexingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexingConfig() *IndexingConfig {
	this := IndexingConfig{}
	return &this
}

// NewIndexingConfigWithDefaults instantiates a new IndexingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexingConfigWithDefaults() *IndexingConfig {
	this := IndexingConfig{}
	return &this
}

// GetSegmentNameGeneratorType returns the SegmentNameGeneratorType field value if set, zero value otherwise.
func (o *IndexingConfig) GetSegmentNameGeneratorType() string {
	if o == nil || IsNil(o.SegmentNameGeneratorType) {
		var ret string
		return ret
	}
	return *o.SegmentNameGeneratorType
}

// GetSegmentNameGeneratorTypeOk returns a tuple with the SegmentNameGeneratorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetSegmentNameGeneratorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentNameGeneratorType) {
		return nil, false
	}
	return o.SegmentNameGeneratorType, true
}

// HasSegmentNameGeneratorType returns a boolean if a field has been set.
func (o *IndexingConfig) HasSegmentNameGeneratorType() bool {
	if o != nil && !IsNil(o.SegmentNameGeneratorType) {
		return true
	}

	return false
}

// SetSegmentNameGeneratorType gets a reference to the given string and assigns it to the SegmentNameGeneratorType field.
func (o *IndexingConfig) SetSegmentNameGeneratorType(v string) {
	o.SegmentNameGeneratorType = &v
}

// GetStreamConfigs returns the StreamConfigs field value if set, zero value otherwise.
func (o *IndexingConfig) GetStreamConfigs() map[string]string {
	if o == nil || IsNil(o.StreamConfigs) {
		var ret map[string]string
		return ret
	}
	return *o.StreamConfigs
}

// GetStreamConfigsOk returns a tuple with the StreamConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetStreamConfigsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.StreamConfigs) {
		return nil, false
	}
	return o.StreamConfigs, true
}

// HasStreamConfigs returns a boolean if a field has been set.
func (o *IndexingConfig) HasStreamConfigs() bool {
	if o != nil && !IsNil(o.StreamConfigs) {
		return true
	}

	return false
}

// SetStreamConfigs gets a reference to the given map[string]string and assigns it to the StreamConfigs field.
func (o *IndexingConfig) SetStreamConfigs(v map[string]string) {
	o.StreamConfigs = &v
}

// GetInvertedIndexColumns returns the InvertedIndexColumns field value if set, zero value otherwise.
func (o *IndexingConfig) GetInvertedIndexColumns() []string {
	if o == nil || IsNil(o.InvertedIndexColumns) {
		var ret []string
		return ret
	}
	return o.InvertedIndexColumns
}

// GetInvertedIndexColumnsOk returns a tuple with the InvertedIndexColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetInvertedIndexColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.InvertedIndexColumns) {
		return nil, false
	}
	return o.InvertedIndexColumns, true
}

// HasInvertedIndexColumns returns a boolean if a field has been set.
func (o *IndexingConfig) HasInvertedIndexColumns() bool {
	if o != nil && !IsNil(o.InvertedIndexColumns) {
		return true
	}

	return false
}

// SetInvertedIndexColumns gets a reference to the given []string and assigns it to the InvertedIndexColumns field.
func (o *IndexingConfig) SetInvertedIndexColumns(v []string) {
	o.InvertedIndexColumns = v
}

// GetNoDictionaryColumns returns the NoDictionaryColumns field value if set, zero value otherwise.
func (o *IndexingConfig) GetNoDictionaryColumns() []string {
	if o == nil || IsNil(o.NoDictionaryColumns) {
		var ret []string
		return ret
	}
	return o.NoDictionaryColumns
}

// GetNoDictionaryColumnsOk returns a tuple with the NoDictionaryColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetNoDictionaryColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.NoDictionaryColumns) {
		return nil, false
	}
	return o.NoDictionaryColumns, true
}

// HasNoDictionaryColumns returns a boolean if a field has been set.
func (o *IndexingConfig) HasNoDictionaryColumns() bool {
	if o != nil && !IsNil(o.NoDictionaryColumns) {
		return true
	}

	return false
}

// SetNoDictionaryColumns gets a reference to the given []string and assigns it to the NoDictionaryColumns field.
func (o *IndexingConfig) SetNoDictionaryColumns(v []string) {
	o.NoDictionaryColumns = v
}

// GetSegmentPartitionConfig returns the SegmentPartitionConfig field value if set, zero value otherwise.
func (o *IndexingConfig) GetSegmentPartitionConfig() SegmentPartitionConfig {
	if o == nil || IsNil(o.SegmentPartitionConfig) {
		var ret SegmentPartitionConfig
		return ret
	}
	return *o.SegmentPartitionConfig
}

// GetSegmentPartitionConfigOk returns a tuple with the SegmentPartitionConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetSegmentPartitionConfigOk() (*SegmentPartitionConfig, bool) {
	if o == nil || IsNil(o.SegmentPartitionConfig) {
		return nil, false
	}
	return o.SegmentPartitionConfig, true
}

// HasSegmentPartitionConfig returns a boolean if a field has been set.
func (o *IndexingConfig) HasSegmentPartitionConfig() bool {
	if o != nil && !IsNil(o.SegmentPartitionConfig) {
		return true
	}

	return false
}

// SetSegmentPartitionConfig gets a reference to the given SegmentPartitionConfig and assigns it to the SegmentPartitionConfig field.
func (o *IndexingConfig) SetSegmentPartitionConfig(v SegmentPartitionConfig) {
	o.SegmentPartitionConfig = &v
}

// GetRangeIndexColumns returns the RangeIndexColumns field value if set, zero value otherwise.
func (o *IndexingConfig) GetRangeIndexColumns() []string {
	if o == nil || IsNil(o.RangeIndexColumns) {
		var ret []string
		return ret
	}
	return o.RangeIndexColumns
}

// GetRangeIndexColumnsOk returns a tuple with the RangeIndexColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetRangeIndexColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.RangeIndexColumns) {
		return nil, false
	}
	return o.RangeIndexColumns, true
}

// HasRangeIndexColumns returns a boolean if a field has been set.
func (o *IndexingConfig) HasRangeIndexColumns() bool {
	if o != nil && !IsNil(o.RangeIndexColumns) {
		return true
	}

	return false
}

// SetRangeIndexColumns gets a reference to the given []string and assigns it to the RangeIndexColumns field.
func (o *IndexingConfig) SetRangeIndexColumns(v []string) {
	o.RangeIndexColumns = v
}

// GetRangeIndexVersion returns the RangeIndexVersion field value if set, zero value otherwise.
func (o *IndexingConfig) GetRangeIndexVersion() int32 {
	if o == nil || IsNil(o.RangeIndexVersion) {
		var ret int32
		return ret
	}
	return *o.RangeIndexVersion
}

// GetRangeIndexVersionOk returns a tuple with the RangeIndexVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetRangeIndexVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.RangeIndexVersion) {
		return nil, false
	}
	return o.RangeIndexVersion, true
}

// HasRangeIndexVersion returns a boolean if a field has been set.
func (o *IndexingConfig) HasRangeIndexVersion() bool {
	if o != nil && !IsNil(o.RangeIndexVersion) {
		return true
	}

	return false
}

// SetRangeIndexVersion gets a reference to the given int32 and assigns it to the RangeIndexVersion field.
func (o *IndexingConfig) SetRangeIndexVersion(v int32) {
	o.RangeIndexVersion = &v
}

// GetFstindexType returns the FstindexType field value if set, zero value otherwise.
func (o *IndexingConfig) GetFstindexType() string {
	if o == nil || IsNil(o.FstindexType) {
		var ret string
		return ret
	}
	return *o.FstindexType
}

// GetFstindexTypeOk returns a tuple with the FstindexType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetFstindexTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FstindexType) {
		return nil, false
	}
	return o.FstindexType, true
}

// HasFstindexType returns a boolean if a field has been set.
func (o *IndexingConfig) HasFstindexType() bool {
	if o != nil && !IsNil(o.FstindexType) {
		return true
	}

	return false
}

// SetFstindexType gets a reference to the given string and assigns it to the FstindexType field.
func (o *IndexingConfig) SetFstindexType(v string) {
	o.FstindexType = &v
}

// GetJsonIndexColumns returns the JsonIndexColumns field value if set, zero value otherwise.
func (o *IndexingConfig) GetJsonIndexColumns() []string {
	if o == nil || IsNil(o.JsonIndexColumns) {
		var ret []string
		return ret
	}
	return o.JsonIndexColumns
}

// GetJsonIndexColumnsOk returns a tuple with the JsonIndexColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetJsonIndexColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.JsonIndexColumns) {
		return nil, false
	}
	return o.JsonIndexColumns, true
}

// HasJsonIndexColumns returns a boolean if a field has been set.
func (o *IndexingConfig) HasJsonIndexColumns() bool {
	if o != nil && !IsNil(o.JsonIndexColumns) {
		return true
	}

	return false
}

// SetJsonIndexColumns gets a reference to the given []string and assigns it to the JsonIndexColumns field.
func (o *IndexingConfig) SetJsonIndexColumns(v []string) {
	o.JsonIndexColumns = v
}

// GetJsonIndexConfigs returns the JsonIndexConfigs field value if set, zero value otherwise.
func (o *IndexingConfig) GetJsonIndexConfigs() map[string]JsonIndexConfig {
	if o == nil || IsNil(o.JsonIndexConfigs) {
		var ret map[string]JsonIndexConfig
		return ret
	}
	return *o.JsonIndexConfigs
}

// GetJsonIndexConfigsOk returns a tuple with the JsonIndexConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetJsonIndexConfigsOk() (*map[string]JsonIndexConfig, bool) {
	if o == nil || IsNil(o.JsonIndexConfigs) {
		return nil, false
	}
	return o.JsonIndexConfigs, true
}

// HasJsonIndexConfigs returns a boolean if a field has been set.
func (o *IndexingConfig) HasJsonIndexConfigs() bool {
	if o != nil && !IsNil(o.JsonIndexConfigs) {
		return true
	}

	return false
}

// SetJsonIndexConfigs gets a reference to the given map[string]JsonIndexConfig and assigns it to the JsonIndexConfigs field.
func (o *IndexingConfig) SetJsonIndexConfigs(v map[string]JsonIndexConfig) {
	o.JsonIndexConfigs = &v
}

// GetAutoGeneratedInvertedIndex returns the AutoGeneratedInvertedIndex field value if set, zero value otherwise.
func (o *IndexingConfig) GetAutoGeneratedInvertedIndex() bool {
	if o == nil || IsNil(o.AutoGeneratedInvertedIndex) {
		var ret bool
		return ret
	}
	return *o.AutoGeneratedInvertedIndex
}

// GetAutoGeneratedInvertedIndexOk returns a tuple with the AutoGeneratedInvertedIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetAutoGeneratedInvertedIndexOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoGeneratedInvertedIndex) {
		return nil, false
	}
	return o.AutoGeneratedInvertedIndex, true
}

// HasAutoGeneratedInvertedIndex returns a boolean if a field has been set.
func (o *IndexingConfig) HasAutoGeneratedInvertedIndex() bool {
	if o != nil && !IsNil(o.AutoGeneratedInvertedIndex) {
		return true
	}

	return false
}

// SetAutoGeneratedInvertedIndex gets a reference to the given bool and assigns it to the AutoGeneratedInvertedIndex field.
func (o *IndexingConfig) SetAutoGeneratedInvertedIndex(v bool) {
	o.AutoGeneratedInvertedIndex = &v
}

// GetCreateInvertedIndexDuringSegmentGeneration returns the CreateInvertedIndexDuringSegmentGeneration field value if set, zero value otherwise.
func (o *IndexingConfig) GetCreateInvertedIndexDuringSegmentGeneration() bool {
	if o == nil || IsNil(o.CreateInvertedIndexDuringSegmentGeneration) {
		var ret bool
		return ret
	}
	return *o.CreateInvertedIndexDuringSegmentGeneration
}

// GetCreateInvertedIndexDuringSegmentGenerationOk returns a tuple with the CreateInvertedIndexDuringSegmentGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetCreateInvertedIndexDuringSegmentGenerationOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateInvertedIndexDuringSegmentGeneration) {
		return nil, false
	}
	return o.CreateInvertedIndexDuringSegmentGeneration, true
}

// HasCreateInvertedIndexDuringSegmentGeneration returns a boolean if a field has been set.
func (o *IndexingConfig) HasCreateInvertedIndexDuringSegmentGeneration() bool {
	if o != nil && !IsNil(o.CreateInvertedIndexDuringSegmentGeneration) {
		return true
	}

	return false
}

// SetCreateInvertedIndexDuringSegmentGeneration gets a reference to the given bool and assigns it to the CreateInvertedIndexDuringSegmentGeneration field.
func (o *IndexingConfig) SetCreateInvertedIndexDuringSegmentGeneration(v bool) {
	o.CreateInvertedIndexDuringSegmentGeneration = &v
}

// GetSortedColumn returns the SortedColumn field value if set, zero value otherwise.
func (o *IndexingConfig) GetSortedColumn() []string {
	if o == nil || IsNil(o.SortedColumn) {
		var ret []string
		return ret
	}
	return o.SortedColumn
}

// GetSortedColumnOk returns a tuple with the SortedColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetSortedColumnOk() ([]string, bool) {
	if o == nil || IsNil(o.SortedColumn) {
		return nil, false
	}
	return o.SortedColumn, true
}

// HasSortedColumn returns a boolean if a field has been set.
func (o *IndexingConfig) HasSortedColumn() bool {
	if o != nil && !IsNil(o.SortedColumn) {
		return true
	}

	return false
}

// SetSortedColumn gets a reference to the given []string and assigns it to the SortedColumn field.
func (o *IndexingConfig) SetSortedColumn(v []string) {
	o.SortedColumn = v
}

// GetBloomFilterColumns returns the BloomFilterColumns field value if set, zero value otherwise.
func (o *IndexingConfig) GetBloomFilterColumns() []string {
	if o == nil || IsNil(o.BloomFilterColumns) {
		var ret []string
		return ret
	}
	return o.BloomFilterColumns
}

// GetBloomFilterColumnsOk returns a tuple with the BloomFilterColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetBloomFilterColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.BloomFilterColumns) {
		return nil, false
	}
	return o.BloomFilterColumns, true
}

// HasBloomFilterColumns returns a boolean if a field has been set.
func (o *IndexingConfig) HasBloomFilterColumns() bool {
	if o != nil && !IsNil(o.BloomFilterColumns) {
		return true
	}

	return false
}

// SetBloomFilterColumns gets a reference to the given []string and assigns it to the BloomFilterColumns field.
func (o *IndexingConfig) SetBloomFilterColumns(v []string) {
	o.BloomFilterColumns = v
}

// GetBloomFilterConfigs returns the BloomFilterConfigs field value if set, zero value otherwise.
func (o *IndexingConfig) GetBloomFilterConfigs() map[string]BloomFilterConfig {
	if o == nil || IsNil(o.BloomFilterConfigs) {
		var ret map[string]BloomFilterConfig
		return ret
	}
	return *o.BloomFilterConfigs
}

// GetBloomFilterConfigsOk returns a tuple with the BloomFilterConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetBloomFilterConfigsOk() (*map[string]BloomFilterConfig, bool) {
	if o == nil || IsNil(o.BloomFilterConfigs) {
		return nil, false
	}
	return o.BloomFilterConfigs, true
}

// HasBloomFilterConfigs returns a boolean if a field has been set.
func (o *IndexingConfig) HasBloomFilterConfigs() bool {
	if o != nil && !IsNil(o.BloomFilterConfigs) {
		return true
	}

	return false
}

// SetBloomFilterConfigs gets a reference to the given map[string]BloomFilterConfig and assigns it to the BloomFilterConfigs field.
func (o *IndexingConfig) SetBloomFilterConfigs(v map[string]BloomFilterConfig) {
	o.BloomFilterConfigs = &v
}

// GetLoadMode returns the LoadMode field value if set, zero value otherwise.
func (o *IndexingConfig) GetLoadMode() string {
	if o == nil || IsNil(o.LoadMode) {
		var ret string
		return ret
	}
	return *o.LoadMode
}

// GetLoadModeOk returns a tuple with the LoadMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetLoadModeOk() (*string, bool) {
	if o == nil || IsNil(o.LoadMode) {
		return nil, false
	}
	return o.LoadMode, true
}

// HasLoadMode returns a boolean if a field has been set.
func (o *IndexingConfig) HasLoadMode() bool {
	if o != nil && !IsNil(o.LoadMode) {
		return true
	}

	return false
}

// SetLoadMode gets a reference to the given string and assigns it to the LoadMode field.
func (o *IndexingConfig) SetLoadMode(v string) {
	o.LoadMode = &v
}

// GetSegmentFormatVersion returns the SegmentFormatVersion field value if set, zero value otherwise.
func (o *IndexingConfig) GetSegmentFormatVersion() string {
	if o == nil || IsNil(o.SegmentFormatVersion) {
		var ret string
		return ret
	}
	return *o.SegmentFormatVersion
}

// GetSegmentFormatVersionOk returns a tuple with the SegmentFormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetSegmentFormatVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentFormatVersion) {
		return nil, false
	}
	return o.SegmentFormatVersion, true
}

// HasSegmentFormatVersion returns a boolean if a field has been set.
func (o *IndexingConfig) HasSegmentFormatVersion() bool {
	if o != nil && !IsNil(o.SegmentFormatVersion) {
		return true
	}

	return false
}

// SetSegmentFormatVersion gets a reference to the given string and assigns it to the SegmentFormatVersion field.
func (o *IndexingConfig) SetSegmentFormatVersion(v string) {
	o.SegmentFormatVersion = &v
}

// GetColumnMinMaxValueGeneratorMode returns the ColumnMinMaxValueGeneratorMode field value if set, zero value otherwise.
func (o *IndexingConfig) GetColumnMinMaxValueGeneratorMode() string {
	if o == nil || IsNil(o.ColumnMinMaxValueGeneratorMode) {
		var ret string
		return ret
	}
	return *o.ColumnMinMaxValueGeneratorMode
}

// GetColumnMinMaxValueGeneratorModeOk returns a tuple with the ColumnMinMaxValueGeneratorMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetColumnMinMaxValueGeneratorModeOk() (*string, bool) {
	if o == nil || IsNil(o.ColumnMinMaxValueGeneratorMode) {
		return nil, false
	}
	return o.ColumnMinMaxValueGeneratorMode, true
}

// HasColumnMinMaxValueGeneratorMode returns a boolean if a field has been set.
func (o *IndexingConfig) HasColumnMinMaxValueGeneratorMode() bool {
	if o != nil && !IsNil(o.ColumnMinMaxValueGeneratorMode) {
		return true
	}

	return false
}

// SetColumnMinMaxValueGeneratorMode gets a reference to the given string and assigns it to the ColumnMinMaxValueGeneratorMode field.
func (o *IndexingConfig) SetColumnMinMaxValueGeneratorMode(v string) {
	o.ColumnMinMaxValueGeneratorMode = &v
}

// GetNoDictionaryConfig returns the NoDictionaryConfig field value if set, zero value otherwise.
func (o *IndexingConfig) GetNoDictionaryConfig() map[string]string {
	if o == nil || IsNil(o.NoDictionaryConfig) {
		var ret map[string]string
		return ret
	}
	return *o.NoDictionaryConfig
}

// GetNoDictionaryConfigOk returns a tuple with the NoDictionaryConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetNoDictionaryConfigOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.NoDictionaryConfig) {
		return nil, false
	}
	return o.NoDictionaryConfig, true
}

// HasNoDictionaryConfig returns a boolean if a field has been set.
func (o *IndexingConfig) HasNoDictionaryConfig() bool {
	if o != nil && !IsNil(o.NoDictionaryConfig) {
		return true
	}

	return false
}

// SetNoDictionaryConfig gets a reference to the given map[string]string and assigns it to the NoDictionaryConfig field.
func (o *IndexingConfig) SetNoDictionaryConfig(v map[string]string) {
	o.NoDictionaryConfig = &v
}

// GetOnHeapDictionaryColumns returns the OnHeapDictionaryColumns field value if set, zero value otherwise.
func (o *IndexingConfig) GetOnHeapDictionaryColumns() []string {
	if o == nil || IsNil(o.OnHeapDictionaryColumns) {
		var ret []string
		return ret
	}
	return o.OnHeapDictionaryColumns
}

// GetOnHeapDictionaryColumnsOk returns a tuple with the OnHeapDictionaryColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetOnHeapDictionaryColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.OnHeapDictionaryColumns) {
		return nil, false
	}
	return o.OnHeapDictionaryColumns, true
}

// HasOnHeapDictionaryColumns returns a boolean if a field has been set.
func (o *IndexingConfig) HasOnHeapDictionaryColumns() bool {
	if o != nil && !IsNil(o.OnHeapDictionaryColumns) {
		return true
	}

	return false
}

// SetOnHeapDictionaryColumns gets a reference to the given []string and assigns it to the OnHeapDictionaryColumns field.
func (o *IndexingConfig) SetOnHeapDictionaryColumns(v []string) {
	o.OnHeapDictionaryColumns = v
}

// GetVarLengthDictionaryColumns returns the VarLengthDictionaryColumns field value if set, zero value otherwise.
func (o *IndexingConfig) GetVarLengthDictionaryColumns() []string {
	if o == nil || IsNil(o.VarLengthDictionaryColumns) {
		var ret []string
		return ret
	}
	return o.VarLengthDictionaryColumns
}

// GetVarLengthDictionaryColumnsOk returns a tuple with the VarLengthDictionaryColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetVarLengthDictionaryColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.VarLengthDictionaryColumns) {
		return nil, false
	}
	return o.VarLengthDictionaryColumns, true
}

// HasVarLengthDictionaryColumns returns a boolean if a field has been set.
func (o *IndexingConfig) HasVarLengthDictionaryColumns() bool {
	if o != nil && !IsNil(o.VarLengthDictionaryColumns) {
		return true
	}

	return false
}

// SetVarLengthDictionaryColumns gets a reference to the given []string and assigns it to the VarLengthDictionaryColumns field.
func (o *IndexingConfig) SetVarLengthDictionaryColumns(v []string) {
	o.VarLengthDictionaryColumns = v
}

// GetEnableDefaultStarTree returns the EnableDefaultStarTree field value if set, zero value otherwise.
func (o *IndexingConfig) GetEnableDefaultStarTree() bool {
	if o == nil || IsNil(o.EnableDefaultStarTree) {
		var ret bool
		return ret
	}
	return *o.EnableDefaultStarTree
}

// GetEnableDefaultStarTreeOk returns a tuple with the EnableDefaultStarTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetEnableDefaultStarTreeOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDefaultStarTree) {
		return nil, false
	}
	return o.EnableDefaultStarTree, true
}

// HasEnableDefaultStarTree returns a boolean if a field has been set.
func (o *IndexingConfig) HasEnableDefaultStarTree() bool {
	if o != nil && !IsNil(o.EnableDefaultStarTree) {
		return true
	}

	return false
}

// SetEnableDefaultStarTree gets a reference to the given bool and assigns it to the EnableDefaultStarTree field.
func (o *IndexingConfig) SetEnableDefaultStarTree(v bool) {
	o.EnableDefaultStarTree = &v
}

// GetStarTreeIndexConfigs returns the StarTreeIndexConfigs field value if set, zero value otherwise.
func (o *IndexingConfig) GetStarTreeIndexConfigs() []StarTreeIndexConfig {
	if o == nil || IsNil(o.StarTreeIndexConfigs) {
		var ret []StarTreeIndexConfig
		return ret
	}
	return o.StarTreeIndexConfigs
}

// GetStarTreeIndexConfigsOk returns a tuple with the StarTreeIndexConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetStarTreeIndexConfigsOk() ([]StarTreeIndexConfig, bool) {
	if o == nil || IsNil(o.StarTreeIndexConfigs) {
		return nil, false
	}
	return o.StarTreeIndexConfigs, true
}

// HasStarTreeIndexConfigs returns a boolean if a field has been set.
func (o *IndexingConfig) HasStarTreeIndexConfigs() bool {
	if o != nil && !IsNil(o.StarTreeIndexConfigs) {
		return true
	}

	return false
}

// SetStarTreeIndexConfigs gets a reference to the given []StarTreeIndexConfig and assigns it to the StarTreeIndexConfigs field.
func (o *IndexingConfig) SetStarTreeIndexConfigs(v []StarTreeIndexConfig) {
	o.StarTreeIndexConfigs = v
}

// GetEnableDynamicStarTreeCreation returns the EnableDynamicStarTreeCreation field value if set, zero value otherwise.
func (o *IndexingConfig) GetEnableDynamicStarTreeCreation() bool {
	if o == nil || IsNil(o.EnableDynamicStarTreeCreation) {
		var ret bool
		return ret
	}
	return *o.EnableDynamicStarTreeCreation
}

// GetEnableDynamicStarTreeCreationOk returns a tuple with the EnableDynamicStarTreeCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetEnableDynamicStarTreeCreationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDynamicStarTreeCreation) {
		return nil, false
	}
	return o.EnableDynamicStarTreeCreation, true
}

// HasEnableDynamicStarTreeCreation returns a boolean if a field has been set.
func (o *IndexingConfig) HasEnableDynamicStarTreeCreation() bool {
	if o != nil && !IsNil(o.EnableDynamicStarTreeCreation) {
		return true
	}

	return false
}

// SetEnableDynamicStarTreeCreation gets a reference to the given bool and assigns it to the EnableDynamicStarTreeCreation field.
func (o *IndexingConfig) SetEnableDynamicStarTreeCreation(v bool) {
	o.EnableDynamicStarTreeCreation = &v
}

// GetAggregateMetrics returns the AggregateMetrics field value if set, zero value otherwise.
func (o *IndexingConfig) GetAggregateMetrics() bool {
	if o == nil || IsNil(o.AggregateMetrics) {
		var ret bool
		return ret
	}
	return *o.AggregateMetrics
}

// GetAggregateMetricsOk returns a tuple with the AggregateMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetAggregateMetricsOk() (*bool, bool) {
	if o == nil || IsNil(o.AggregateMetrics) {
		return nil, false
	}
	return o.AggregateMetrics, true
}

// HasAggregateMetrics returns a boolean if a field has been set.
func (o *IndexingConfig) HasAggregateMetrics() bool {
	if o != nil && !IsNil(o.AggregateMetrics) {
		return true
	}

	return false
}

// SetAggregateMetrics gets a reference to the given bool and assigns it to the AggregateMetrics field.
func (o *IndexingConfig) SetAggregateMetrics(v bool) {
	o.AggregateMetrics = &v
}

// GetNullHandlingEnabled returns the NullHandlingEnabled field value if set, zero value otherwise.
func (o *IndexingConfig) GetNullHandlingEnabled() bool {
	if o == nil || IsNil(o.NullHandlingEnabled) {
		var ret bool
		return ret
	}
	return *o.NullHandlingEnabled
}

// GetNullHandlingEnabledOk returns a tuple with the NullHandlingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetNullHandlingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NullHandlingEnabled) {
		return nil, false
	}
	return o.NullHandlingEnabled, true
}

// HasNullHandlingEnabled returns a boolean if a field has been set.
func (o *IndexingConfig) HasNullHandlingEnabled() bool {
	if o != nil && !IsNil(o.NullHandlingEnabled) {
		return true
	}

	return false
}

// SetNullHandlingEnabled gets a reference to the given bool and assigns it to the NullHandlingEnabled field.
func (o *IndexingConfig) SetNullHandlingEnabled(v bool) {
	o.NullHandlingEnabled = &v
}

// GetOptimizeDictionary returns the OptimizeDictionary field value if set, zero value otherwise.
func (o *IndexingConfig) GetOptimizeDictionary() bool {
	if o == nil || IsNil(o.OptimizeDictionary) {
		var ret bool
		return ret
	}
	return *o.OptimizeDictionary
}

// GetOptimizeDictionaryOk returns a tuple with the OptimizeDictionary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetOptimizeDictionaryOk() (*bool, bool) {
	if o == nil || IsNil(o.OptimizeDictionary) {
		return nil, false
	}
	return o.OptimizeDictionary, true
}

// HasOptimizeDictionary returns a boolean if a field has been set.
func (o *IndexingConfig) HasOptimizeDictionary() bool {
	if o != nil && !IsNil(o.OptimizeDictionary) {
		return true
	}

	return false
}

// SetOptimizeDictionary gets a reference to the given bool and assigns it to the OptimizeDictionary field.
func (o *IndexingConfig) SetOptimizeDictionary(v bool) {
	o.OptimizeDictionary = &v
}

// GetOptimizeDictionaryForMetrics returns the OptimizeDictionaryForMetrics field value if set, zero value otherwise.
func (o *IndexingConfig) GetOptimizeDictionaryForMetrics() bool {
	if o == nil || IsNil(o.OptimizeDictionaryForMetrics) {
		var ret bool
		return ret
	}
	return *o.OptimizeDictionaryForMetrics
}

// GetOptimizeDictionaryForMetricsOk returns a tuple with the OptimizeDictionaryForMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetOptimizeDictionaryForMetricsOk() (*bool, bool) {
	if o == nil || IsNil(o.OptimizeDictionaryForMetrics) {
		return nil, false
	}
	return o.OptimizeDictionaryForMetrics, true
}

// HasOptimizeDictionaryForMetrics returns a boolean if a field has been set.
func (o *IndexingConfig) HasOptimizeDictionaryForMetrics() bool {
	if o != nil && !IsNil(o.OptimizeDictionaryForMetrics) {
		return true
	}

	return false
}

// SetOptimizeDictionaryForMetrics gets a reference to the given bool and assigns it to the OptimizeDictionaryForMetrics field.
func (o *IndexingConfig) SetOptimizeDictionaryForMetrics(v bool) {
	o.OptimizeDictionaryForMetrics = &v
}

// GetNoDictionarySizeRatioThreshold returns the NoDictionarySizeRatioThreshold field value if set, zero value otherwise.
func (o *IndexingConfig) GetNoDictionarySizeRatioThreshold() float64 {
	if o == nil || IsNil(o.NoDictionarySizeRatioThreshold) {
		var ret float64
		return ret
	}
	return *o.NoDictionarySizeRatioThreshold
}

// GetNoDictionarySizeRatioThresholdOk returns a tuple with the NoDictionarySizeRatioThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingConfig) GetNoDictionarySizeRatioThresholdOk() (*float64, bool) {
	if o == nil || IsNil(o.NoDictionarySizeRatioThreshold) {
		return nil, false
	}
	return o.NoDictionarySizeRatioThreshold, true
}

// HasNoDictionarySizeRatioThreshold returns a boolean if a field has been set.
func (o *IndexingConfig) HasNoDictionarySizeRatioThreshold() bool {
	if o != nil && !IsNil(o.NoDictionarySizeRatioThreshold) {
		return true
	}

	return false
}

// SetNoDictionarySizeRatioThreshold gets a reference to the given float64 and assigns it to the NoDictionarySizeRatioThreshold field.
func (o *IndexingConfig) SetNoDictionarySizeRatioThreshold(v float64) {
	o.NoDictionarySizeRatioThreshold = &v
}

func (o IndexingConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexingConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SegmentNameGeneratorType) {
		toSerialize["segmentNameGeneratorType"] = o.SegmentNameGeneratorType
	}
	if !IsNil(o.StreamConfigs) {
		toSerialize["streamConfigs"] = o.StreamConfigs
	}
	if !IsNil(o.InvertedIndexColumns) {
		toSerialize["invertedIndexColumns"] = o.InvertedIndexColumns
	}
	if !IsNil(o.NoDictionaryColumns) {
		toSerialize["noDictionaryColumns"] = o.NoDictionaryColumns
	}
	if !IsNil(o.SegmentPartitionConfig) {
		toSerialize["segmentPartitionConfig"] = o.SegmentPartitionConfig
	}
	if !IsNil(o.RangeIndexColumns) {
		toSerialize["rangeIndexColumns"] = o.RangeIndexColumns
	}
	if !IsNil(o.RangeIndexVersion) {
		toSerialize["rangeIndexVersion"] = o.RangeIndexVersion
	}
	if !IsNil(o.FstindexType) {
		toSerialize["fstindexType"] = o.FstindexType
	}
	if !IsNil(o.JsonIndexColumns) {
		toSerialize["jsonIndexColumns"] = o.JsonIndexColumns
	}
	if !IsNil(o.JsonIndexConfigs) {
		toSerialize["jsonIndexConfigs"] = o.JsonIndexConfigs
	}
	if !IsNil(o.AutoGeneratedInvertedIndex) {
		toSerialize["autoGeneratedInvertedIndex"] = o.AutoGeneratedInvertedIndex
	}
	if !IsNil(o.CreateInvertedIndexDuringSegmentGeneration) {
		toSerialize["createInvertedIndexDuringSegmentGeneration"] = o.CreateInvertedIndexDuringSegmentGeneration
	}
	if !IsNil(o.SortedColumn) {
		toSerialize["sortedColumn"] = o.SortedColumn
	}
	if !IsNil(o.BloomFilterColumns) {
		toSerialize["bloomFilterColumns"] = o.BloomFilterColumns
	}
	if !IsNil(o.BloomFilterConfigs) {
		toSerialize["bloomFilterConfigs"] = o.BloomFilterConfigs
	}
	if !IsNil(o.LoadMode) {
		toSerialize["loadMode"] = o.LoadMode
	}
	if !IsNil(o.SegmentFormatVersion) {
		toSerialize["segmentFormatVersion"] = o.SegmentFormatVersion
	}
	if !IsNil(o.ColumnMinMaxValueGeneratorMode) {
		toSerialize["columnMinMaxValueGeneratorMode"] = o.ColumnMinMaxValueGeneratorMode
	}
	if !IsNil(o.NoDictionaryConfig) {
		toSerialize["noDictionaryConfig"] = o.NoDictionaryConfig
	}
	if !IsNil(o.OnHeapDictionaryColumns) {
		toSerialize["onHeapDictionaryColumns"] = o.OnHeapDictionaryColumns
	}
	if !IsNil(o.VarLengthDictionaryColumns) {
		toSerialize["varLengthDictionaryColumns"] = o.VarLengthDictionaryColumns
	}
	if !IsNil(o.EnableDefaultStarTree) {
		toSerialize["enableDefaultStarTree"] = o.EnableDefaultStarTree
	}
	if !IsNil(o.StarTreeIndexConfigs) {
		toSerialize["starTreeIndexConfigs"] = o.StarTreeIndexConfigs
	}
	if !IsNil(o.EnableDynamicStarTreeCreation) {
		toSerialize["enableDynamicStarTreeCreation"] = o.EnableDynamicStarTreeCreation
	}
	if !IsNil(o.AggregateMetrics) {
		toSerialize["aggregateMetrics"] = o.AggregateMetrics
	}
	if !IsNil(o.NullHandlingEnabled) {
		toSerialize["nullHandlingEnabled"] = o.NullHandlingEnabled
	}
	if !IsNil(o.OptimizeDictionary) {
		toSerialize["optimizeDictionary"] = o.OptimizeDictionary
	}
	if !IsNil(o.OptimizeDictionaryForMetrics) {
		toSerialize["optimizeDictionaryForMetrics"] = o.OptimizeDictionaryForMetrics
	}
	if !IsNil(o.NoDictionarySizeRatioThreshold) {
		toSerialize["noDictionarySizeRatioThreshold"] = o.NoDictionarySizeRatioThreshold
	}
	return toSerialize, nil
}

type NullableIndexingConfig struct {
	value *IndexingConfig
	isSet bool
}

func (v NullableIndexingConfig) Get() *IndexingConfig {
	return v.value
}

func (v *NullableIndexingConfig) Set(val *IndexingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexingConfig(val *IndexingConfig) *NullableIndexingConfig {
	return &NullableIndexingConfig{value: val, isSet: true}
}

func (v NullableIndexingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


