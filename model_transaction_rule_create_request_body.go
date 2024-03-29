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

// checks if the TransactionRuleCreateRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionRuleCreateRequestBody{}

// TransactionRuleCreateRequestBody struct for TransactionRuleCreateRequestBody
type TransactionRuleCreateRequestBody struct {
	TransactionRule *TransactionRuleCreateRequest `json:"transaction_rule,omitempty"`
}

// NewTransactionRuleCreateRequestBody instantiates a new TransactionRuleCreateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRuleCreateRequestBody() *TransactionRuleCreateRequestBody {
	this := TransactionRuleCreateRequestBody{}
	return &this
}

// NewTransactionRuleCreateRequestBodyWithDefaults instantiates a new TransactionRuleCreateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRuleCreateRequestBodyWithDefaults() *TransactionRuleCreateRequestBody {
	this := TransactionRuleCreateRequestBody{}
	return &this
}

// GetTransactionRule returns the TransactionRule field value if set, zero value otherwise.
func (o *TransactionRuleCreateRequestBody) GetTransactionRule() TransactionRuleCreateRequest {
	if o == nil || IsNil(o.TransactionRule) {
		var ret TransactionRuleCreateRequest
		return ret
	}
	return *o.TransactionRule
}

// GetTransactionRuleOk returns a tuple with the TransactionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleCreateRequestBody) GetTransactionRuleOk() (*TransactionRuleCreateRequest, bool) {
	if o == nil || IsNil(o.TransactionRule) {
		return nil, false
	}
	return o.TransactionRule, true
}

// HasTransactionRule returns a boolean if a field has been set.
func (o *TransactionRuleCreateRequestBody) HasTransactionRule() bool {
	if o != nil && !IsNil(o.TransactionRule) {
		return true
	}

	return false
}

// SetTransactionRule gets a reference to the given TransactionRuleCreateRequest and assigns it to the TransactionRule field.
func (o *TransactionRuleCreateRequestBody) SetTransactionRule(v TransactionRuleCreateRequest) {
	o.TransactionRule = &v
}

func (o TransactionRuleCreateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRuleCreateRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionRule) {
		toSerialize["transaction_rule"] = o.TransactionRule
	}
	return toSerialize, nil
}

type NullableTransactionRuleCreateRequestBody struct {
	value *TransactionRuleCreateRequestBody
	isSet bool
}

func (v NullableTransactionRuleCreateRequestBody) Get() *TransactionRuleCreateRequestBody {
	return v.value
}

func (v *NullableTransactionRuleCreateRequestBody) Set(val *TransactionRuleCreateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRuleCreateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRuleCreateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRuleCreateRequestBody(val *TransactionRuleCreateRequestBody) *NullableTransactionRuleCreateRequestBody {
	return &NullableTransactionRuleCreateRequestBody{value: val, isSet: true}
}

func (v NullableTransactionRuleCreateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRuleCreateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


