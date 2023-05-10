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

// PaymentProcessorAuthorizationCodeRequest struct for PaymentProcessorAuthorizationCodeRequest
type PaymentProcessorAuthorizationCodeRequest struct {
	AccountGuid NullableString `json:"account_guid,omitempty"`
	MemberGuid NullableString `json:"member_guid,omitempty"`
	UserGuid NullableString `json:"user_guid,omitempty"`
}

// NewPaymentProcessorAuthorizationCodeRequest instantiates a new PaymentProcessorAuthorizationCodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentProcessorAuthorizationCodeRequest() *PaymentProcessorAuthorizationCodeRequest {
	this := PaymentProcessorAuthorizationCodeRequest{}
	return &this
}

// NewPaymentProcessorAuthorizationCodeRequestWithDefaults instantiates a new PaymentProcessorAuthorizationCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentProcessorAuthorizationCodeRequestWithDefaults() *PaymentProcessorAuthorizationCodeRequest {
	this := PaymentProcessorAuthorizationCodeRequest{}
	return &this
}

// GetAccountGuid returns the AccountGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentProcessorAuthorizationCodeRequest) GetAccountGuid() string {
	if o == nil || o.AccountGuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountGuid.Get()
}

// GetAccountGuidOk returns a tuple with the AccountGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentProcessorAuthorizationCodeRequest) GetAccountGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountGuid.Get(), o.AccountGuid.IsSet()
}

// HasAccountGuid returns a boolean if a field has been set.
func (o *PaymentProcessorAuthorizationCodeRequest) HasAccountGuid() bool {
	if o != nil && o.AccountGuid.IsSet() {
		return true
	}

	return false
}

// SetAccountGuid gets a reference to the given NullableString and assigns it to the AccountGuid field.
func (o *PaymentProcessorAuthorizationCodeRequest) SetAccountGuid(v string) {
	o.AccountGuid.Set(&v)
}
// SetAccountGuidNil sets the value for AccountGuid to be an explicit nil
func (o *PaymentProcessorAuthorizationCodeRequest) SetAccountGuidNil() {
	o.AccountGuid.Set(nil)
}

// UnsetAccountGuid ensures that no value is present for AccountGuid, not even an explicit nil
func (o *PaymentProcessorAuthorizationCodeRequest) UnsetAccountGuid() {
	o.AccountGuid.Unset()
}

// GetMemberGuid returns the MemberGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentProcessorAuthorizationCodeRequest) GetMemberGuid() string {
	if o == nil || o.MemberGuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.MemberGuid.Get()
}

// GetMemberGuidOk returns a tuple with the MemberGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentProcessorAuthorizationCodeRequest) GetMemberGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MemberGuid.Get(), o.MemberGuid.IsSet()
}

// HasMemberGuid returns a boolean if a field has been set.
func (o *PaymentProcessorAuthorizationCodeRequest) HasMemberGuid() bool {
	if o != nil && o.MemberGuid.IsSet() {
		return true
	}

	return false
}

// SetMemberGuid gets a reference to the given NullableString and assigns it to the MemberGuid field.
func (o *PaymentProcessorAuthorizationCodeRequest) SetMemberGuid(v string) {
	o.MemberGuid.Set(&v)
}
// SetMemberGuidNil sets the value for MemberGuid to be an explicit nil
func (o *PaymentProcessorAuthorizationCodeRequest) SetMemberGuidNil() {
	o.MemberGuid.Set(nil)
}

// UnsetMemberGuid ensures that no value is present for MemberGuid, not even an explicit nil
func (o *PaymentProcessorAuthorizationCodeRequest) UnsetMemberGuid() {
	o.MemberGuid.Unset()
}

// GetUserGuid returns the UserGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentProcessorAuthorizationCodeRequest) GetUserGuid() string {
	if o == nil || o.UserGuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserGuid.Get()
}

// GetUserGuidOk returns a tuple with the UserGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentProcessorAuthorizationCodeRequest) GetUserGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserGuid.Get(), o.UserGuid.IsSet()
}

// HasUserGuid returns a boolean if a field has been set.
func (o *PaymentProcessorAuthorizationCodeRequest) HasUserGuid() bool {
	if o != nil && o.UserGuid.IsSet() {
		return true
	}

	return false
}

// SetUserGuid gets a reference to the given NullableString and assigns it to the UserGuid field.
func (o *PaymentProcessorAuthorizationCodeRequest) SetUserGuid(v string) {
	o.UserGuid.Set(&v)
}
// SetUserGuidNil sets the value for UserGuid to be an explicit nil
func (o *PaymentProcessorAuthorizationCodeRequest) SetUserGuidNil() {
	o.UserGuid.Set(nil)
}

// UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
func (o *PaymentProcessorAuthorizationCodeRequest) UnsetUserGuid() {
	o.UserGuid.Unset()
}

func (o PaymentProcessorAuthorizationCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountGuid.IsSet() {
		toSerialize["account_guid"] = o.AccountGuid.Get()
	}
	if o.MemberGuid.IsSet() {
		toSerialize["member_guid"] = o.MemberGuid.Get()
	}
	if o.UserGuid.IsSet() {
		toSerialize["user_guid"] = o.UserGuid.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentProcessorAuthorizationCodeRequest struct {
	value *PaymentProcessorAuthorizationCodeRequest
	isSet bool
}

func (v NullablePaymentProcessorAuthorizationCodeRequest) Get() *PaymentProcessorAuthorizationCodeRequest {
	return v.value
}

func (v *NullablePaymentProcessorAuthorizationCodeRequest) Set(val *PaymentProcessorAuthorizationCodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentProcessorAuthorizationCodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentProcessorAuthorizationCodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentProcessorAuthorizationCodeRequest(val *PaymentProcessorAuthorizationCodeRequest) *NullablePaymentProcessorAuthorizationCodeRequest {
	return &NullablePaymentProcessorAuthorizationCodeRequest{value: val, isSet: true}
}

func (v NullablePaymentProcessorAuthorizationCodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentProcessorAuthorizationCodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

