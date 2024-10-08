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

// checks if the SplitTransactionRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SplitTransactionRequestBody{}

// SplitTransactionRequestBody struct for SplitTransactionRequestBody
type SplitTransactionRequestBody struct {
	Transactions SplitTransactionRequest `json:"transactions"`
}

// NewSplitTransactionRequestBody instantiates a new SplitTransactionRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitTransactionRequestBody(transactions SplitTransactionRequest) *SplitTransactionRequestBody {
	this := SplitTransactionRequestBody{}
	this.Transactions = transactions
	return &this
}

// NewSplitTransactionRequestBodyWithDefaults instantiates a new SplitTransactionRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitTransactionRequestBodyWithDefaults() *SplitTransactionRequestBody {
	this := SplitTransactionRequestBody{}
	return &this
}

// GetTransactions returns the Transactions field value
func (o *SplitTransactionRequestBody) GetTransactions() SplitTransactionRequest {
	if o == nil {
		var ret SplitTransactionRequest
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *SplitTransactionRequestBody) GetTransactionsOk() (*SplitTransactionRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *SplitTransactionRequestBody) SetTransactions(v SplitTransactionRequest) {
	o.Transactions = v
}

func (o SplitTransactionRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitTransactionRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactions"] = o.Transactions
	return toSerialize, nil
}

type NullableSplitTransactionRequestBody struct {
	value *SplitTransactionRequestBody
	isSet bool
}

func (v NullableSplitTransactionRequestBody) Get() *SplitTransactionRequestBody {
	return v.value
}

func (v *NullableSplitTransactionRequestBody) Set(val *SplitTransactionRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitTransactionRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitTransactionRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitTransactionRequestBody(val *SplitTransactionRequestBody) *NullableSplitTransactionRequestBody {
	return &NullableSplitTransactionRequestBody{value: val, isSet: true}
}

func (v NullableSplitTransactionRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitTransactionRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


