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

// checks if the ServerRebalanceJobStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerRebalanceJobStatusResponse{}

// ServerRebalanceJobStatusResponse struct for ServerRebalanceJobStatusResponse
type ServerRebalanceJobStatusResponse struct {
	TableRebalanceProgressStats *TableRebalanceProgressStats `json:"tableRebalanceProgressStats,omitempty"`
	TimeElapsedSinceStartInSeconds *int64 `json:"timeElapsedSinceStartInSeconds,omitempty"`
}

// NewServerRebalanceJobStatusResponse instantiates a new ServerRebalanceJobStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerRebalanceJobStatusResponse() *ServerRebalanceJobStatusResponse {
	this := ServerRebalanceJobStatusResponse{}
	return &this
}

// NewServerRebalanceJobStatusResponseWithDefaults instantiates a new ServerRebalanceJobStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerRebalanceJobStatusResponseWithDefaults() *ServerRebalanceJobStatusResponse {
	this := ServerRebalanceJobStatusResponse{}
	return &this
}

// GetTableRebalanceProgressStats returns the TableRebalanceProgressStats field value if set, zero value otherwise.
func (o *ServerRebalanceJobStatusResponse) GetTableRebalanceProgressStats() TableRebalanceProgressStats {
	if o == nil || IsNil(o.TableRebalanceProgressStats) {
		var ret TableRebalanceProgressStats
		return ret
	}
	return *o.TableRebalanceProgressStats
}

// GetTableRebalanceProgressStatsOk returns a tuple with the TableRebalanceProgressStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerRebalanceJobStatusResponse) GetTableRebalanceProgressStatsOk() (*TableRebalanceProgressStats, bool) {
	if o == nil || IsNil(o.TableRebalanceProgressStats) {
		return nil, false
	}
	return o.TableRebalanceProgressStats, true
}

// HasTableRebalanceProgressStats returns a boolean if a field has been set.
func (o *ServerRebalanceJobStatusResponse) HasTableRebalanceProgressStats() bool {
	if o != nil && !IsNil(o.TableRebalanceProgressStats) {
		return true
	}

	return false
}

// SetTableRebalanceProgressStats gets a reference to the given TableRebalanceProgressStats and assigns it to the TableRebalanceProgressStats field.
func (o *ServerRebalanceJobStatusResponse) SetTableRebalanceProgressStats(v TableRebalanceProgressStats) {
	o.TableRebalanceProgressStats = &v
}

// GetTimeElapsedSinceStartInSeconds returns the TimeElapsedSinceStartInSeconds field value if set, zero value otherwise.
func (o *ServerRebalanceJobStatusResponse) GetTimeElapsedSinceStartInSeconds() int64 {
	if o == nil || IsNil(o.TimeElapsedSinceStartInSeconds) {
		var ret int64
		return ret
	}
	return *o.TimeElapsedSinceStartInSeconds
}

// GetTimeElapsedSinceStartInSecondsOk returns a tuple with the TimeElapsedSinceStartInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerRebalanceJobStatusResponse) GetTimeElapsedSinceStartInSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeElapsedSinceStartInSeconds) {
		return nil, false
	}
	return o.TimeElapsedSinceStartInSeconds, true
}

// HasTimeElapsedSinceStartInSeconds returns a boolean if a field has been set.
func (o *ServerRebalanceJobStatusResponse) HasTimeElapsedSinceStartInSeconds() bool {
	if o != nil && !IsNil(o.TimeElapsedSinceStartInSeconds) {
		return true
	}

	return false
}

// SetTimeElapsedSinceStartInSeconds gets a reference to the given int64 and assigns it to the TimeElapsedSinceStartInSeconds field.
func (o *ServerRebalanceJobStatusResponse) SetTimeElapsedSinceStartInSeconds(v int64) {
	o.TimeElapsedSinceStartInSeconds = &v
}

func (o ServerRebalanceJobStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerRebalanceJobStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TableRebalanceProgressStats) {
		toSerialize["tableRebalanceProgressStats"] = o.TableRebalanceProgressStats
	}
	if !IsNil(o.TimeElapsedSinceStartInSeconds) {
		toSerialize["timeElapsedSinceStartInSeconds"] = o.TimeElapsedSinceStartInSeconds
	}
	return toSerialize, nil
}

type NullableServerRebalanceJobStatusResponse struct {
	value *ServerRebalanceJobStatusResponse
	isSet bool
}

func (v NullableServerRebalanceJobStatusResponse) Get() *ServerRebalanceJobStatusResponse {
	return v.value
}

func (v *NullableServerRebalanceJobStatusResponse) Set(val *ServerRebalanceJobStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServerRebalanceJobStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServerRebalanceJobStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerRebalanceJobStatusResponse(val *ServerRebalanceJobStatusResponse) *NullableServerRebalanceJobStatusResponse {
	return &NullableServerRebalanceJobStatusResponse{value: val, isSet: true}
}

func (v NullableServerRebalanceJobStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerRebalanceJobStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

