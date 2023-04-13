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

// checks if the ServerReloadControllerJobStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerReloadControllerJobStatusResponse{}

// ServerReloadControllerJobStatusResponse struct for ServerReloadControllerJobStatusResponse
type ServerReloadControllerJobStatusResponse struct {
	Metadata *map[string]string `json:"metadata,omitempty"`
	SuccessCount *int32 `json:"successCount,omitempty"`
	TotalSegmentCount *int32 `json:"totalSegmentCount,omitempty"`
	TotalServersQueried *int32 `json:"totalServersQueried,omitempty"`
	TotalServerCallsFailed *int32 `json:"totalServerCallsFailed,omitempty"`
	TimeElapsedInMinutes *float64 `json:"timeElapsedInMinutes,omitempty"`
	EstimatedTimeRemainingInMinutes *float64 `json:"estimatedTimeRemainingInMinutes,omitempty"`
}

// NewServerReloadControllerJobStatusResponse instantiates a new ServerReloadControllerJobStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerReloadControllerJobStatusResponse() *ServerReloadControllerJobStatusResponse {
	this := ServerReloadControllerJobStatusResponse{}
	return &this
}

// NewServerReloadControllerJobStatusResponseWithDefaults instantiates a new ServerReloadControllerJobStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerReloadControllerJobStatusResponseWithDefaults() *ServerReloadControllerJobStatusResponse {
	this := ServerReloadControllerJobStatusResponse{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ServerReloadControllerJobStatusResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReloadControllerJobStatusResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServerReloadControllerJobStatusResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *ServerReloadControllerJobStatusResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetSuccessCount returns the SuccessCount field value if set, zero value otherwise.
func (o *ServerReloadControllerJobStatusResponse) GetSuccessCount() int32 {
	if o == nil || IsNil(o.SuccessCount) {
		var ret int32
		return ret
	}
	return *o.SuccessCount
}

// GetSuccessCountOk returns a tuple with the SuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReloadControllerJobStatusResponse) GetSuccessCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SuccessCount) {
		return nil, false
	}
	return o.SuccessCount, true
}

// HasSuccessCount returns a boolean if a field has been set.
func (o *ServerReloadControllerJobStatusResponse) HasSuccessCount() bool {
	if o != nil && !IsNil(o.SuccessCount) {
		return true
	}

	return false
}

// SetSuccessCount gets a reference to the given int32 and assigns it to the SuccessCount field.
func (o *ServerReloadControllerJobStatusResponse) SetSuccessCount(v int32) {
	o.SuccessCount = &v
}

// GetTotalSegmentCount returns the TotalSegmentCount field value if set, zero value otherwise.
func (o *ServerReloadControllerJobStatusResponse) GetTotalSegmentCount() int32 {
	if o == nil || IsNil(o.TotalSegmentCount) {
		var ret int32
		return ret
	}
	return *o.TotalSegmentCount
}

// GetTotalSegmentCountOk returns a tuple with the TotalSegmentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReloadControllerJobStatusResponse) GetTotalSegmentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSegmentCount) {
		return nil, false
	}
	return o.TotalSegmentCount, true
}

// HasTotalSegmentCount returns a boolean if a field has been set.
func (o *ServerReloadControllerJobStatusResponse) HasTotalSegmentCount() bool {
	if o != nil && !IsNil(o.TotalSegmentCount) {
		return true
	}

	return false
}

// SetTotalSegmentCount gets a reference to the given int32 and assigns it to the TotalSegmentCount field.
func (o *ServerReloadControllerJobStatusResponse) SetTotalSegmentCount(v int32) {
	o.TotalSegmentCount = &v
}

// GetTotalServersQueried returns the TotalServersQueried field value if set, zero value otherwise.
func (o *ServerReloadControllerJobStatusResponse) GetTotalServersQueried() int32 {
	if o == nil || IsNil(o.TotalServersQueried) {
		var ret int32
		return ret
	}
	return *o.TotalServersQueried
}

// GetTotalServersQueriedOk returns a tuple with the TotalServersQueried field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReloadControllerJobStatusResponse) GetTotalServersQueriedOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalServersQueried) {
		return nil, false
	}
	return o.TotalServersQueried, true
}

// HasTotalServersQueried returns a boolean if a field has been set.
func (o *ServerReloadControllerJobStatusResponse) HasTotalServersQueried() bool {
	if o != nil && !IsNil(o.TotalServersQueried) {
		return true
	}

	return false
}

// SetTotalServersQueried gets a reference to the given int32 and assigns it to the TotalServersQueried field.
func (o *ServerReloadControllerJobStatusResponse) SetTotalServersQueried(v int32) {
	o.TotalServersQueried = &v
}

// GetTotalServerCallsFailed returns the TotalServerCallsFailed field value if set, zero value otherwise.
func (o *ServerReloadControllerJobStatusResponse) GetTotalServerCallsFailed() int32 {
	if o == nil || IsNil(o.TotalServerCallsFailed) {
		var ret int32
		return ret
	}
	return *o.TotalServerCallsFailed
}

// GetTotalServerCallsFailedOk returns a tuple with the TotalServerCallsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReloadControllerJobStatusResponse) GetTotalServerCallsFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalServerCallsFailed) {
		return nil, false
	}
	return o.TotalServerCallsFailed, true
}

// HasTotalServerCallsFailed returns a boolean if a field has been set.
func (o *ServerReloadControllerJobStatusResponse) HasTotalServerCallsFailed() bool {
	if o != nil && !IsNil(o.TotalServerCallsFailed) {
		return true
	}

	return false
}

// SetTotalServerCallsFailed gets a reference to the given int32 and assigns it to the TotalServerCallsFailed field.
func (o *ServerReloadControllerJobStatusResponse) SetTotalServerCallsFailed(v int32) {
	o.TotalServerCallsFailed = &v
}

// GetTimeElapsedInMinutes returns the TimeElapsedInMinutes field value if set, zero value otherwise.
func (o *ServerReloadControllerJobStatusResponse) GetTimeElapsedInMinutes() float64 {
	if o == nil || IsNil(o.TimeElapsedInMinutes) {
		var ret float64
		return ret
	}
	return *o.TimeElapsedInMinutes
}

// GetTimeElapsedInMinutesOk returns a tuple with the TimeElapsedInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReloadControllerJobStatusResponse) GetTimeElapsedInMinutesOk() (*float64, bool) {
	if o == nil || IsNil(o.TimeElapsedInMinutes) {
		return nil, false
	}
	return o.TimeElapsedInMinutes, true
}

// HasTimeElapsedInMinutes returns a boolean if a field has been set.
func (o *ServerReloadControllerJobStatusResponse) HasTimeElapsedInMinutes() bool {
	if o != nil && !IsNil(o.TimeElapsedInMinutes) {
		return true
	}

	return false
}

// SetTimeElapsedInMinutes gets a reference to the given float64 and assigns it to the TimeElapsedInMinutes field.
func (o *ServerReloadControllerJobStatusResponse) SetTimeElapsedInMinutes(v float64) {
	o.TimeElapsedInMinutes = &v
}

// GetEstimatedTimeRemainingInMinutes returns the EstimatedTimeRemainingInMinutes field value if set, zero value otherwise.
func (o *ServerReloadControllerJobStatusResponse) GetEstimatedTimeRemainingInMinutes() float64 {
	if o == nil || IsNil(o.EstimatedTimeRemainingInMinutes) {
		var ret float64
		return ret
	}
	return *o.EstimatedTimeRemainingInMinutes
}

// GetEstimatedTimeRemainingInMinutesOk returns a tuple with the EstimatedTimeRemainingInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReloadControllerJobStatusResponse) GetEstimatedTimeRemainingInMinutesOk() (*float64, bool) {
	if o == nil || IsNil(o.EstimatedTimeRemainingInMinutes) {
		return nil, false
	}
	return o.EstimatedTimeRemainingInMinutes, true
}

// HasEstimatedTimeRemainingInMinutes returns a boolean if a field has been set.
func (o *ServerReloadControllerJobStatusResponse) HasEstimatedTimeRemainingInMinutes() bool {
	if o != nil && !IsNil(o.EstimatedTimeRemainingInMinutes) {
		return true
	}

	return false
}

// SetEstimatedTimeRemainingInMinutes gets a reference to the given float64 and assigns it to the EstimatedTimeRemainingInMinutes field.
func (o *ServerReloadControllerJobStatusResponse) SetEstimatedTimeRemainingInMinutes(v float64) {
	o.EstimatedTimeRemainingInMinutes = &v
}

func (o ServerReloadControllerJobStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerReloadControllerJobStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.SuccessCount) {
		toSerialize["successCount"] = o.SuccessCount
	}
	if !IsNil(o.TotalSegmentCount) {
		toSerialize["totalSegmentCount"] = o.TotalSegmentCount
	}
	if !IsNil(o.TotalServersQueried) {
		toSerialize["totalServersQueried"] = o.TotalServersQueried
	}
	if !IsNil(o.TotalServerCallsFailed) {
		toSerialize["totalServerCallsFailed"] = o.TotalServerCallsFailed
	}
	if !IsNil(o.TimeElapsedInMinutes) {
		toSerialize["timeElapsedInMinutes"] = o.TimeElapsedInMinutes
	}
	if !IsNil(o.EstimatedTimeRemainingInMinutes) {
		toSerialize["estimatedTimeRemainingInMinutes"] = o.EstimatedTimeRemainingInMinutes
	}
	return toSerialize, nil
}

type NullableServerReloadControllerJobStatusResponse struct {
	value *ServerReloadControllerJobStatusResponse
	isSet bool
}

func (v NullableServerReloadControllerJobStatusResponse) Get() *ServerReloadControllerJobStatusResponse {
	return v.value
}

func (v *NullableServerReloadControllerJobStatusResponse) Set(val *ServerReloadControllerJobStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServerReloadControllerJobStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServerReloadControllerJobStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerReloadControllerJobStatusResponse(val *ServerReloadControllerJobStatusResponse) *NullableServerReloadControllerJobStatusResponse {
	return &NullableServerReloadControllerJobStatusResponse{value: val, isSet: true}
}

func (v NullableServerReloadControllerJobStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerReloadControllerJobStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


