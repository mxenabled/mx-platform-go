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

// checks if the MicrodepositVerifyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MicrodepositVerifyRequest{}

// MicrodepositVerifyRequest struct for MicrodepositVerifyRequest
type MicrodepositVerifyRequest struct {
	DepositAmount1 *int32 `json:"deposit_amount_1,omitempty"`
	DepositAmount2 *int32 `json:"deposit_amount_2,omitempty"`
}

// NewMicrodepositVerifyRequest instantiates a new MicrodepositVerifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrodepositVerifyRequest() *MicrodepositVerifyRequest {
	this := MicrodepositVerifyRequest{}
	return &this
}

// NewMicrodepositVerifyRequestWithDefaults instantiates a new MicrodepositVerifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrodepositVerifyRequestWithDefaults() *MicrodepositVerifyRequest {
	this := MicrodepositVerifyRequest{}
	return &this
}

// GetDepositAmount1 returns the DepositAmount1 field value if set, zero value otherwise.
func (o *MicrodepositVerifyRequest) GetDepositAmount1() int32 {
	if o == nil || IsNil(o.DepositAmount1) {
		var ret int32
		return ret
	}
	return *o.DepositAmount1
}

// GetDepositAmount1Ok returns a tuple with the DepositAmount1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositVerifyRequest) GetDepositAmount1Ok() (*int32, bool) {
	if o == nil || IsNil(o.DepositAmount1) {
		return nil, false
	}
	return o.DepositAmount1, true
}

// HasDepositAmount1 returns a boolean if a field has been set.
func (o *MicrodepositVerifyRequest) HasDepositAmount1() bool {
	if o != nil && !IsNil(o.DepositAmount1) {
		return true
	}

	return false
}

// SetDepositAmount1 gets a reference to the given int32 and assigns it to the DepositAmount1 field.
func (o *MicrodepositVerifyRequest) SetDepositAmount1(v int32) {
	o.DepositAmount1 = &v
}

// GetDepositAmount2 returns the DepositAmount2 field value if set, zero value otherwise.
func (o *MicrodepositVerifyRequest) GetDepositAmount2() int32 {
	if o == nil || IsNil(o.DepositAmount2) {
		var ret int32
		return ret
	}
	return *o.DepositAmount2
}

// GetDepositAmount2Ok returns a tuple with the DepositAmount2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositVerifyRequest) GetDepositAmount2Ok() (*int32, bool) {
	if o == nil || IsNil(o.DepositAmount2) {
		return nil, false
	}
	return o.DepositAmount2, true
}

// HasDepositAmount2 returns a boolean if a field has been set.
func (o *MicrodepositVerifyRequest) HasDepositAmount2() bool {
	if o != nil && !IsNil(o.DepositAmount2) {
		return true
	}

	return false
}

// SetDepositAmount2 gets a reference to the given int32 and assigns it to the DepositAmount2 field.
func (o *MicrodepositVerifyRequest) SetDepositAmount2(v int32) {
	o.DepositAmount2 = &v
}

func (o MicrodepositVerifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MicrodepositVerifyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DepositAmount1) {
		toSerialize["deposit_amount_1"] = o.DepositAmount1
	}
	if !IsNil(o.DepositAmount2) {
		toSerialize["deposit_amount_2"] = o.DepositAmount2
	}
	return toSerialize, nil
}

type NullableMicrodepositVerifyRequest struct {
	value *MicrodepositVerifyRequest
	isSet bool
}

func (v NullableMicrodepositVerifyRequest) Get() *MicrodepositVerifyRequest {
	return v.value
}

func (v *NullableMicrodepositVerifyRequest) Set(val *MicrodepositVerifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrodepositVerifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrodepositVerifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrodepositVerifyRequest(val *MicrodepositVerifyRequest) *NullableMicrodepositVerifyRequest {
	return &NullableMicrodepositVerifyRequest{value: val, isSet: true}
}

func (v NullableMicrodepositVerifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrodepositVerifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

