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

// checks if the CreditCardProductResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditCardProductResponse{}

// CreditCardProductResponse struct for CreditCardProductResponse
type CreditCardProductResponse struct {
	Reward *CreditCardProduct `json:"reward,omitempty"`
}

// NewCreditCardProductResponse instantiates a new CreditCardProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCardProductResponse() *CreditCardProductResponse {
	this := CreditCardProductResponse{}
	return &this
}

// NewCreditCardProductResponseWithDefaults instantiates a new CreditCardProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCardProductResponseWithDefaults() *CreditCardProductResponse {
	this := CreditCardProductResponse{}
	return &this
}

// GetReward returns the Reward field value if set, zero value otherwise.
func (o *CreditCardProductResponse) GetReward() CreditCardProduct {
	if o == nil || IsNil(o.Reward) {
		var ret CreditCardProduct
		return ret
	}
	return *o.Reward
}

// GetRewardOk returns a tuple with the Reward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardProductResponse) GetRewardOk() (*CreditCardProduct, bool) {
	if o == nil || IsNil(o.Reward) {
		return nil, false
	}
	return o.Reward, true
}

// HasReward returns a boolean if a field has been set.
func (o *CreditCardProductResponse) HasReward() bool {
	if o != nil && !IsNil(o.Reward) {
		return true
	}

	return false
}

// SetReward gets a reference to the given CreditCardProduct and assigns it to the Reward field.
func (o *CreditCardProductResponse) SetReward(v CreditCardProduct) {
	o.Reward = &v
}

func (o CreditCardProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditCardProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reward) {
		toSerialize["reward"] = o.Reward
	}
	return toSerialize, nil
}

type NullableCreditCardProductResponse struct {
	value *CreditCardProductResponse
	isSet bool
}

func (v NullableCreditCardProductResponse) Get() *CreditCardProductResponse {
	return v.value
}

func (v *NullableCreditCardProductResponse) Set(val *CreditCardProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCardProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCardProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCardProductResponse(val *CreditCardProductResponse) *NullableCreditCardProductResponse {
	return &NullableCreditCardProductResponse{value: val, isSet: true}
}

func (v NullableCreditCardProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCardProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


