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

// checks if the RepositionRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositionRequestBody{}

// RepositionRequestBody struct for RepositionRequestBody
type RepositionRequestBody struct {
	Goals []RepositionRequest `json:"goals,omitempty"`
}

// NewRepositionRequestBody instantiates a new RepositionRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositionRequestBody() *RepositionRequestBody {
	this := RepositionRequestBody{}
	return &this
}

// NewRepositionRequestBodyWithDefaults instantiates a new RepositionRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositionRequestBodyWithDefaults() *RepositionRequestBody {
	this := RepositionRequestBody{}
	return &this
}

// GetGoals returns the Goals field value if set, zero value otherwise.
func (o *RepositionRequestBody) GetGoals() []RepositionRequest {
	if o == nil || IsNil(o.Goals) {
		var ret []RepositionRequest
		return ret
	}
	return o.Goals
}

// GetGoalsOk returns a tuple with the Goals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositionRequestBody) GetGoalsOk() ([]RepositionRequest, bool) {
	if o == nil || IsNil(o.Goals) {
		return nil, false
	}
	return o.Goals, true
}

// HasGoals returns a boolean if a field has been set.
func (o *RepositionRequestBody) HasGoals() bool {
	if o != nil && !IsNil(o.Goals) {
		return true
	}

	return false
}

// SetGoals gets a reference to the given []RepositionRequest and assigns it to the Goals field.
func (o *RepositionRequestBody) SetGoals(v []RepositionRequest) {
	o.Goals = v
}

func (o RepositionRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositionRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Goals) {
		toSerialize["goals"] = o.Goals
	}
	return toSerialize, nil
}

type NullableRepositionRequestBody struct {
	value *RepositionRequestBody
	isSet bool
}

func (v NullableRepositionRequestBody) Get() *RepositionRequestBody {
	return v.value
}

func (v *NullableRepositionRequestBody) Set(val *RepositionRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositionRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositionRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositionRequestBody(val *RepositionRequestBody) *NullableRepositionRequestBody {
	return &NullableRepositionRequestBody{value: val, isSet: true}
}

func (v NullableRepositionRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositionRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


