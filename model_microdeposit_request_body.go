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

// checks if the MicrodepositRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MicrodepositRequestBody{}

// MicrodepositRequestBody struct for MicrodepositRequestBody
type MicrodepositRequestBody struct {
	MicroDeposit *MicrodepositRequest `json:"micro_deposit,omitempty"`
}

// NewMicrodepositRequestBody instantiates a new MicrodepositRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrodepositRequestBody() *MicrodepositRequestBody {
	this := MicrodepositRequestBody{}
	return &this
}

// NewMicrodepositRequestBodyWithDefaults instantiates a new MicrodepositRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrodepositRequestBodyWithDefaults() *MicrodepositRequestBody {
	this := MicrodepositRequestBody{}
	return &this
}

// GetMicroDeposit returns the MicroDeposit field value if set, zero value otherwise.
func (o *MicrodepositRequestBody) GetMicroDeposit() MicrodepositRequest {
	if o == nil || IsNil(o.MicroDeposit) {
		var ret MicrodepositRequest
		return ret
	}
	return *o.MicroDeposit
}

// GetMicroDepositOk returns a tuple with the MicroDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositRequestBody) GetMicroDepositOk() (*MicrodepositRequest, bool) {
	if o == nil || IsNil(o.MicroDeposit) {
		return nil, false
	}
	return o.MicroDeposit, true
}

// HasMicroDeposit returns a boolean if a field has been set.
func (o *MicrodepositRequestBody) HasMicroDeposit() bool {
	if o != nil && !IsNil(o.MicroDeposit) {
		return true
	}

	return false
}

// SetMicroDeposit gets a reference to the given MicrodepositRequest and assigns it to the MicroDeposit field.
func (o *MicrodepositRequestBody) SetMicroDeposit(v MicrodepositRequest) {
	o.MicroDeposit = &v
}

func (o MicrodepositRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MicrodepositRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MicroDeposit) {
		toSerialize["micro_deposit"] = o.MicroDeposit
	}
	return toSerialize, nil
}

type NullableMicrodepositRequestBody struct {
	value *MicrodepositRequestBody
	isSet bool
}

func (v NullableMicrodepositRequestBody) Get() *MicrodepositRequestBody {
	return v.value
}

func (v *NullableMicrodepositRequestBody) Set(val *MicrodepositRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrodepositRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrodepositRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrodepositRequestBody(val *MicrodepositRequestBody) *NullableMicrodepositRequestBody {
	return &NullableMicrodepositRequestBody{value: val, isSet: true}
}

func (v NullableMicrodepositRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrodepositRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


