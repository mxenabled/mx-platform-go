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

// checks if the StatementResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatementResponseBody{}

// StatementResponseBody struct for StatementResponseBody
type StatementResponseBody struct {
	Statement *StatementResponse `json:"statement,omitempty"`
}

// NewStatementResponseBody instantiates a new StatementResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementResponseBody() *StatementResponseBody {
	this := StatementResponseBody{}
	return &this
}

// NewStatementResponseBodyWithDefaults instantiates a new StatementResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementResponseBodyWithDefaults() *StatementResponseBody {
	this := StatementResponseBody{}
	return &this
}

// GetStatement returns the Statement field value if set, zero value otherwise.
func (o *StatementResponseBody) GetStatement() StatementResponse {
	if o == nil || IsNil(o.Statement) {
		var ret StatementResponse
		return ret
	}
	return *o.Statement
}

// GetStatementOk returns a tuple with the Statement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementResponseBody) GetStatementOk() (*StatementResponse, bool) {
	if o == nil || IsNil(o.Statement) {
		return nil, false
	}
	return o.Statement, true
}

// HasStatement returns a boolean if a field has been set.
func (o *StatementResponseBody) HasStatement() bool {
	if o != nil && !IsNil(o.Statement) {
		return true
	}

	return false
}

// SetStatement gets a reference to the given StatementResponse and assigns it to the Statement field.
func (o *StatementResponseBody) SetStatement(v StatementResponse) {
	o.Statement = &v
}

func (o StatementResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatementResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Statement) {
		toSerialize["statement"] = o.Statement
	}
	return toSerialize, nil
}

type NullableStatementResponseBody struct {
	value *StatementResponseBody
	isSet bool
}

func (v NullableStatementResponseBody) Get() *StatementResponseBody {
	return v.value
}

func (v *NullableStatementResponseBody) Set(val *StatementResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementResponseBody(val *StatementResponseBody) *NullableStatementResponseBody {
	return &NullableStatementResponseBody{value: val, isSet: true}
}

func (v NullableStatementResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


