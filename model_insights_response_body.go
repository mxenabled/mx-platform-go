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

// checks if the InsightsResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsightsResponseBody{}

// InsightsResponseBody struct for InsightsResponseBody
type InsightsResponseBody struct {
	Insights []InsightResponse `json:"insights,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewInsightsResponseBody instantiates a new InsightsResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsightsResponseBody() *InsightsResponseBody {
	this := InsightsResponseBody{}
	return &this
}

// NewInsightsResponseBodyWithDefaults instantiates a new InsightsResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightsResponseBodyWithDefaults() *InsightsResponseBody {
	this := InsightsResponseBody{}
	return &this
}

// GetInsights returns the Insights field value if set, zero value otherwise.
func (o *InsightsResponseBody) GetInsights() []InsightResponse {
	if o == nil || IsNil(o.Insights) {
		var ret []InsightResponse
		return ret
	}
	return o.Insights
}

// GetInsightsOk returns a tuple with the Insights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponseBody) GetInsightsOk() ([]InsightResponse, bool) {
	if o == nil || IsNil(o.Insights) {
		return nil, false
	}
	return o.Insights, true
}

// HasInsights returns a boolean if a field has been set.
func (o *InsightsResponseBody) HasInsights() bool {
	if o != nil && !IsNil(o.Insights) {
		return true
	}

	return false
}

// SetInsights gets a reference to the given []InsightResponse and assigns it to the Insights field.
func (o *InsightsResponseBody) SetInsights(v []InsightResponse) {
	o.Insights = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *InsightsResponseBody) GetPagination() PaginationResponse {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponseBody) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *InsightsResponseBody) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *InsightsResponseBody) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o InsightsResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsightsResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Insights) {
		toSerialize["insights"] = o.Insights
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableInsightsResponseBody struct {
	value *InsightsResponseBody
	isSet bool
}

func (v NullableInsightsResponseBody) Get() *InsightsResponseBody {
	return v.value
}

func (v *NullableInsightsResponseBody) Set(val *InsightsResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightsResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightsResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightsResponseBody(val *InsightsResponseBody) *NullableInsightsResponseBody {
	return &NullableInsightsResponseBody{value: val, isSet: true}
}

func (v NullableInsightsResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsightsResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


