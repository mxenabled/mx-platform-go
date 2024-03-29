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

// checks if the AccountCreateRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountCreateRequestBody{}

// AccountCreateRequestBody struct for AccountCreateRequestBody
type AccountCreateRequestBody struct {
	Account *AccountCreateRequest `json:"account,omitempty"`
}

// NewAccountCreateRequestBody instantiates a new AccountCreateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreateRequestBody() *AccountCreateRequestBody {
	this := AccountCreateRequestBody{}
	return &this
}

// NewAccountCreateRequestBodyWithDefaults instantiates a new AccountCreateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreateRequestBodyWithDefaults() *AccountCreateRequestBody {
	this := AccountCreateRequestBody{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AccountCreateRequestBody) GetAccount() AccountCreateRequest {
	if o == nil || IsNil(o.Account) {
		var ret AccountCreateRequest
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequestBody) GetAccountOk() (*AccountCreateRequest, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AccountCreateRequestBody) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AccountCreateRequest and assigns it to the Account field.
func (o *AccountCreateRequestBody) SetAccount(v AccountCreateRequest) {
	o.Account = &v
}

func (o AccountCreateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCreateRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return toSerialize, nil
}

type NullableAccountCreateRequestBody struct {
	value *AccountCreateRequestBody
	isSet bool
}

func (v NullableAccountCreateRequestBody) Get() *AccountCreateRequestBody {
	return v.value
}

func (v *NullableAccountCreateRequestBody) Set(val *AccountCreateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreateRequestBody(val *AccountCreateRequestBody) *NullableAccountCreateRequestBody {
	return &NullableAccountCreateRequestBody{value: val, isSet: true}
}

func (v NullableAccountCreateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


