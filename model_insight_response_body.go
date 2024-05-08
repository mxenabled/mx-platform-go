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

// checks if the InsightResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsightResponseBody{}

// InsightResponseBody struct for InsightResponseBody
type InsightResponseBody struct {
	Insight []InsightResponse `json:"insight,omitempty"`
}

// NewInsightResponseBody instantiates a new InsightResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsightResponseBody() *InsightResponseBody {
	this := InsightResponseBody{}
	return &this
}

// NewInsightResponseBodyWithDefaults instantiates a new InsightResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightResponseBodyWithDefaults() *InsightResponseBody {
	this := InsightResponseBody{}
	return &this
}

// GetInsight returns the Insight field value if set, zero value otherwise.
func (o *InsightResponseBody) GetInsight() []InsightResponse {
	if o == nil || IsNil(o.Insight) {
		var ret []InsightResponse
		return ret
	}
	return o.Insight
}

// GetInsightOk returns a tuple with the Insight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightResponseBody) GetInsightOk() ([]InsightResponse, bool) {
	if o == nil || IsNil(o.Insight) {
		return nil, false
	}
	return o.Insight, true
}

// HasInsight returns a boolean if a field has been set.
func (o *InsightResponseBody) HasInsight() bool {
	if o != nil && !IsNil(o.Insight) {
		return true
	}

	return false
}

// SetInsight gets a reference to the given []InsightResponse and assigns it to the Insight field.
func (o *InsightResponseBody) SetInsight(v []InsightResponse) {
	o.Insight = v
}

func (o InsightResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsightResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Insight) {
		toSerialize["insight"] = o.Insight
	}
	return toSerialize, nil
}

type NullableInsightResponseBody struct {
	value *InsightResponseBody
	isSet bool
}

func (v NullableInsightResponseBody) Get() *InsightResponseBody {
	return v.value
}

func (v *NullableInsightResponseBody) Set(val *InsightResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightResponseBody(val *InsightResponseBody) *NullableInsightResponseBody {
	return &NullableInsightResponseBody{value: val, isSet: true}
}

func (v NullableInsightResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsightResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


