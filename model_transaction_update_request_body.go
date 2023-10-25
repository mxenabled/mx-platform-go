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

// checks if the TransactionUpdateRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionUpdateRequestBody{}

// TransactionUpdateRequestBody struct for TransactionUpdateRequestBody
type TransactionUpdateRequestBody struct {
	Transaction *TransactionUpdateRequest `json:"transaction,omitempty"`
}

// NewTransactionUpdateRequestBody instantiates a new TransactionUpdateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionUpdateRequestBody() *TransactionUpdateRequestBody {
	this := TransactionUpdateRequestBody{}
	return &this
}

// NewTransactionUpdateRequestBodyWithDefaults instantiates a new TransactionUpdateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionUpdateRequestBodyWithDefaults() *TransactionUpdateRequestBody {
	this := TransactionUpdateRequestBody{}
	return &this
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *TransactionUpdateRequestBody) GetTransaction() TransactionUpdateRequest {
	if o == nil || IsNil(o.Transaction) {
		var ret TransactionUpdateRequest
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionUpdateRequestBody) GetTransactionOk() (*TransactionUpdateRequest, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *TransactionUpdateRequestBody) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given TransactionUpdateRequest and assigns it to the Transaction field.
func (o *TransactionUpdateRequestBody) SetTransaction(v TransactionUpdateRequest) {
	o.Transaction = &v
}

func (o TransactionUpdateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionUpdateRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	return toSerialize, nil
}

type NullableTransactionUpdateRequestBody struct {
	value *TransactionUpdateRequestBody
	isSet bool
}

func (v NullableTransactionUpdateRequestBody) Get() *TransactionUpdateRequestBody {
	return v.value
}

func (v *NullableTransactionUpdateRequestBody) Set(val *TransactionUpdateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionUpdateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionUpdateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionUpdateRequestBody(val *TransactionUpdateRequestBody) *NullableTransactionUpdateRequestBody {
	return &NullableTransactionUpdateRequestBody{value: val, isSet: true}
}

func (v NullableTransactionUpdateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionUpdateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


