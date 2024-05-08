/*
MX Platform API

The MX Platform API is a powerful, fully-featured API designed to make aggregating and enhancing financial data easy and reliable. It can seamlessly connect your app or website to tens of thousands of financial institutions.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mxplatformgo

import (
	"encoding/json"
)

// checks if the InsightUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsightUpdateRequest{}

// InsightUpdateRequest struct for InsightUpdateRequest
type InsightUpdateRequest struct {
	HasBeenDisplayed *bool `json:"has_been_displayed,omitempty"`
	IsDismissed *bool `json:"is_dismissed,omitempty"`
}

// NewInsightUpdateRequest instantiates a new InsightUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsightUpdateRequest() *InsightUpdateRequest {
	this := InsightUpdateRequest{}
	return &this
}

// NewInsightUpdateRequestWithDefaults instantiates a new InsightUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightUpdateRequestWithDefaults() *InsightUpdateRequest {
	this := InsightUpdateRequest{}
	return &this
}

// GetHasBeenDisplayed returns the HasBeenDisplayed field value if set, zero value otherwise.
func (o *InsightUpdateRequest) GetHasBeenDisplayed() bool {
	if o == nil || IsNil(o.HasBeenDisplayed) {
		var ret bool
		return ret
	}
	return *o.HasBeenDisplayed
}

// GetHasBeenDisplayedOk returns a tuple with the HasBeenDisplayed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightUpdateRequest) GetHasBeenDisplayedOk() (*bool, bool) {
	if o == nil || IsNil(o.HasBeenDisplayed) {
		return nil, false
	}
	return o.HasBeenDisplayed, true
}

// HasHasBeenDisplayed returns a boolean if a field has been set.
func (o *InsightUpdateRequest) HasHasBeenDisplayed() bool {
	if o != nil && !IsNil(o.HasBeenDisplayed) {
		return true
	}

	return false
}

// SetHasBeenDisplayed gets a reference to the given bool and assigns it to the HasBeenDisplayed field.
func (o *InsightUpdateRequest) SetHasBeenDisplayed(v bool) {
	o.HasBeenDisplayed = &v
}

// GetIsDismissed returns the IsDismissed field value if set, zero value otherwise.
func (o *InsightUpdateRequest) GetIsDismissed() bool {
	if o == nil || IsNil(o.IsDismissed) {
		var ret bool
		return ret
	}
	return *o.IsDismissed
}

// GetIsDismissedOk returns a tuple with the IsDismissed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightUpdateRequest) GetIsDismissedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDismissed) {
		return nil, false
	}
	return o.IsDismissed, true
}

// HasIsDismissed returns a boolean if a field has been set.
func (o *InsightUpdateRequest) HasIsDismissed() bool {
	if o != nil && !IsNil(o.IsDismissed) {
		return true
	}

	return false
}

// SetIsDismissed gets a reference to the given bool and assigns it to the IsDismissed field.
func (o *InsightUpdateRequest) SetIsDismissed(v bool) {
	o.IsDismissed = &v
}

func (o InsightUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsightUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HasBeenDisplayed) {
		toSerialize["has_been_displayed"] = o.HasBeenDisplayed
	}
	if !IsNil(o.IsDismissed) {
		toSerialize["is_dismissed"] = o.IsDismissed
	}
	return toSerialize, nil
}

type NullableInsightUpdateRequest struct {
	value *InsightUpdateRequest
	isSet bool
}

func (v NullableInsightUpdateRequest) Get() *InsightUpdateRequest {
	return v.value
}

func (v *NullableInsightUpdateRequest) Set(val *InsightUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightUpdateRequest(val *InsightUpdateRequest) *NullableInsightUpdateRequest {
	return &NullableInsightUpdateRequest{value: val, isSet: true}
}

func (v NullableInsightUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsightUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


