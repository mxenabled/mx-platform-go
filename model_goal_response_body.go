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

// checks if the GoalResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoalResponseBody{}

// GoalResponseBody struct for GoalResponseBody
type GoalResponseBody struct {
	Goal *GoalResponse `json:"goal,omitempty"`
}

// NewGoalResponseBody instantiates a new GoalResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoalResponseBody() *GoalResponseBody {
	this := GoalResponseBody{}
	return &this
}

// NewGoalResponseBodyWithDefaults instantiates a new GoalResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoalResponseBodyWithDefaults() *GoalResponseBody {
	this := GoalResponseBody{}
	return &this
}

// GetGoal returns the Goal field value if set, zero value otherwise.
func (o *GoalResponseBody) GetGoal() GoalResponse {
	if o == nil || IsNil(o.Goal) {
		var ret GoalResponse
		return ret
	}
	return *o.Goal
}

// GetGoalOk returns a tuple with the Goal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponseBody) GetGoalOk() (*GoalResponse, bool) {
	if o == nil || IsNil(o.Goal) {
		return nil, false
	}
	return o.Goal, true
}

// HasGoal returns a boolean if a field has been set.
func (o *GoalResponseBody) HasGoal() bool {
	if o != nil && !IsNil(o.Goal) {
		return true
	}

	return false
}

// SetGoal gets a reference to the given GoalResponse and assigns it to the Goal field.
func (o *GoalResponseBody) SetGoal(v GoalResponse) {
	o.Goal = &v
}

func (o GoalResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoalResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Goal) {
		toSerialize["goal"] = o.Goal
	}
	return toSerialize, nil
}

type NullableGoalResponseBody struct {
	value *GoalResponseBody
	isSet bool
}

func (v NullableGoalResponseBody) Get() *GoalResponseBody {
	return v.value
}

func (v *NullableGoalResponseBody) Set(val *GoalResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableGoalResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableGoalResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoalResponseBody(val *GoalResponseBody) *NullableGoalResponseBody {
	return &NullableGoalResponseBody{value: val, isSet: true}
}

func (v NullableGoalResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoalResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


