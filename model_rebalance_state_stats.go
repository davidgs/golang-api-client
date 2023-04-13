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

// checks if the RebalanceStateStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RebalanceStateStats{}

// RebalanceStateStats struct for RebalanceStateStats
type RebalanceStateStats struct {
	SegmentsMissing *int32 `json:"_segmentsMissing,omitempty"`
	SegmentsToRebalance *int32 `json:"_segmentsToRebalance,omitempty"`
	PercentSegmentsToRebalance *float64 `json:"_percentSegmentsToRebalance,omitempty"`
	ReplicasToRebalance *int32 `json:"_replicasToRebalance,omitempty"`
}

// NewRebalanceStateStats instantiates a new RebalanceStateStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRebalanceStateStats() *RebalanceStateStats {
	this := RebalanceStateStats{}
	return &this
}

// NewRebalanceStateStatsWithDefaults instantiates a new RebalanceStateStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRebalanceStateStatsWithDefaults() *RebalanceStateStats {
	this := RebalanceStateStats{}
	return &this
}

// GetSegmentsMissing returns the SegmentsMissing field value if set, zero value otherwise.
func (o *RebalanceStateStats) GetSegmentsMissing() int32 {
	if o == nil || IsNil(o.SegmentsMissing) {
		var ret int32
		return ret
	}
	return *o.SegmentsMissing
}

// GetSegmentsMissingOk returns a tuple with the SegmentsMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebalanceStateStats) GetSegmentsMissingOk() (*int32, bool) {
	if o == nil || IsNil(o.SegmentsMissing) {
		return nil, false
	}
	return o.SegmentsMissing, true
}

// HasSegmentsMissing returns a boolean if a field has been set.
func (o *RebalanceStateStats) HasSegmentsMissing() bool {
	if o != nil && !IsNil(o.SegmentsMissing) {
		return true
	}

	return false
}

// SetSegmentsMissing gets a reference to the given int32 and assigns it to the SegmentsMissing field.
func (o *RebalanceStateStats) SetSegmentsMissing(v int32) {
	o.SegmentsMissing = &v
}

// GetSegmentsToRebalance returns the SegmentsToRebalance field value if set, zero value otherwise.
func (o *RebalanceStateStats) GetSegmentsToRebalance() int32 {
	if o == nil || IsNil(o.SegmentsToRebalance) {
		var ret int32
		return ret
	}
	return *o.SegmentsToRebalance
}

// GetSegmentsToRebalanceOk returns a tuple with the SegmentsToRebalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebalanceStateStats) GetSegmentsToRebalanceOk() (*int32, bool) {
	if o == nil || IsNil(o.SegmentsToRebalance) {
		return nil, false
	}
	return o.SegmentsToRebalance, true
}

// HasSegmentsToRebalance returns a boolean if a field has been set.
func (o *RebalanceStateStats) HasSegmentsToRebalance() bool {
	if o != nil && !IsNil(o.SegmentsToRebalance) {
		return true
	}

	return false
}

// SetSegmentsToRebalance gets a reference to the given int32 and assigns it to the SegmentsToRebalance field.
func (o *RebalanceStateStats) SetSegmentsToRebalance(v int32) {
	o.SegmentsToRebalance = &v
}

// GetPercentSegmentsToRebalance returns the PercentSegmentsToRebalance field value if set, zero value otherwise.
func (o *RebalanceStateStats) GetPercentSegmentsToRebalance() float64 {
	if o == nil || IsNil(o.PercentSegmentsToRebalance) {
		var ret float64
		return ret
	}
	return *o.PercentSegmentsToRebalance
}

// GetPercentSegmentsToRebalanceOk returns a tuple with the PercentSegmentsToRebalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebalanceStateStats) GetPercentSegmentsToRebalanceOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentSegmentsToRebalance) {
		return nil, false
	}
	return o.PercentSegmentsToRebalance, true
}

// HasPercentSegmentsToRebalance returns a boolean if a field has been set.
func (o *RebalanceStateStats) HasPercentSegmentsToRebalance() bool {
	if o != nil && !IsNil(o.PercentSegmentsToRebalance) {
		return true
	}

	return false
}

// SetPercentSegmentsToRebalance gets a reference to the given float64 and assigns it to the PercentSegmentsToRebalance field.
func (o *RebalanceStateStats) SetPercentSegmentsToRebalance(v float64) {
	o.PercentSegmentsToRebalance = &v
}

// GetReplicasToRebalance returns the ReplicasToRebalance field value if set, zero value otherwise.
func (o *RebalanceStateStats) GetReplicasToRebalance() int32 {
	if o == nil || IsNil(o.ReplicasToRebalance) {
		var ret int32
		return ret
	}
	return *o.ReplicasToRebalance
}

// GetReplicasToRebalanceOk returns a tuple with the ReplicasToRebalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebalanceStateStats) GetReplicasToRebalanceOk() (*int32, bool) {
	if o == nil || IsNil(o.ReplicasToRebalance) {
		return nil, false
	}
	return o.ReplicasToRebalance, true
}

// HasReplicasToRebalance returns a boolean if a field has been set.
func (o *RebalanceStateStats) HasReplicasToRebalance() bool {
	if o != nil && !IsNil(o.ReplicasToRebalance) {
		return true
	}

	return false
}

// SetReplicasToRebalance gets a reference to the given int32 and assigns it to the ReplicasToRebalance field.
func (o *RebalanceStateStats) SetReplicasToRebalance(v int32) {
	o.ReplicasToRebalance = &v
}

func (o RebalanceStateStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RebalanceStateStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SegmentsMissing) {
		toSerialize["_segmentsMissing"] = o.SegmentsMissing
	}
	if !IsNil(o.SegmentsToRebalance) {
		toSerialize["_segmentsToRebalance"] = o.SegmentsToRebalance
	}
	if !IsNil(o.PercentSegmentsToRebalance) {
		toSerialize["_percentSegmentsToRebalance"] = o.PercentSegmentsToRebalance
	}
	if !IsNil(o.ReplicasToRebalance) {
		toSerialize["_replicasToRebalance"] = o.ReplicasToRebalance
	}
	return toSerialize, nil
}

type NullableRebalanceStateStats struct {
	value *RebalanceStateStats
	isSet bool
}

func (v NullableRebalanceStateStats) Get() *RebalanceStateStats {
	return v.value
}

func (v *NullableRebalanceStateStats) Set(val *RebalanceStateStats) {
	v.value = val
	v.isSet = true
}

func (v NullableRebalanceStateStats) IsSet() bool {
	return v.isSet
}

func (v *NullableRebalanceStateStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRebalanceStateStats(val *RebalanceStateStats) *NullableRebalanceStateStats {
	return &NullableRebalanceStateStats{value: val, isSet: true}
}

func (v NullableRebalanceStateStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRebalanceStateStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

