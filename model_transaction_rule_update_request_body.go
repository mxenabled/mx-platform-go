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

// checks if the TransactionRuleUpdateRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionRuleUpdateRequestBody{}

// TransactionRuleUpdateRequestBody struct for TransactionRuleUpdateRequestBody
type TransactionRuleUpdateRequestBody struct {
	TransactionRule *TransactionRuleUpdateRequest `json:"transaction_rule,omitempty"`
}

// NewTransactionRuleUpdateRequestBody instantiates a new TransactionRuleUpdateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRuleUpdateRequestBody() *TransactionRuleUpdateRequestBody {
	this := TransactionRuleUpdateRequestBody{}
	return &this
}

// NewTransactionRuleUpdateRequestBodyWithDefaults instantiates a new TransactionRuleUpdateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRuleUpdateRequestBodyWithDefaults() *TransactionRuleUpdateRequestBody {
	this := TransactionRuleUpdateRequestBody{}
	return &this
}

// GetTransactionRule returns the TransactionRule field value if set, zero value otherwise.
func (o *TransactionRuleUpdateRequestBody) GetTransactionRule() TransactionRuleUpdateRequest {
	if o == nil || IsNil(o.TransactionRule) {
		var ret TransactionRuleUpdateRequest
		return ret
	}
	return *o.TransactionRule
}

// GetTransactionRuleOk returns a tuple with the TransactionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleUpdateRequestBody) GetTransactionRuleOk() (*TransactionRuleUpdateRequest, bool) {
	if o == nil || IsNil(o.TransactionRule) {
		return nil, false
	}
	return o.TransactionRule, true
}

// HasTransactionRule returns a boolean if a field has been set.
func (o *TransactionRuleUpdateRequestBody) HasTransactionRule() bool {
	if o != nil && !IsNil(o.TransactionRule) {
		return true
	}

	return false
}

// SetTransactionRule gets a reference to the given TransactionRuleUpdateRequest and assigns it to the TransactionRule field.
func (o *TransactionRuleUpdateRequestBody) SetTransactionRule(v TransactionRuleUpdateRequest) {
	o.TransactionRule = &v
}

func (o TransactionRuleUpdateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRuleUpdateRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionRule) {
		toSerialize["transaction_rule"] = o.TransactionRule
	}
	return toSerialize, nil
}

type NullableTransactionRuleUpdateRequestBody struct {
	value *TransactionRuleUpdateRequestBody
	isSet bool
}

func (v NullableTransactionRuleUpdateRequestBody) Get() *TransactionRuleUpdateRequestBody {
	return v.value
}

func (v *NullableTransactionRuleUpdateRequestBody) Set(val *TransactionRuleUpdateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRuleUpdateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRuleUpdateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRuleUpdateRequestBody(val *TransactionRuleUpdateRequestBody) *NullableTransactionRuleUpdateRequestBody {
	return &NullableTransactionRuleUpdateRequestBody{value: val, isSet: true}
}

func (v NullableTransactionRuleUpdateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRuleUpdateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


