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

// PaymentProcessorTokenResponseBody struct for PaymentProcessorTokenResponseBody
type PaymentProcessorTokenResponseBody struct {
	AccessToken NullableString `json:"access_token,omitempty"`
	Scope NullableString `json:"scope,omitempty"`
	TokenType NullableString `json:"token_type,omitempty"`
}

// NewPaymentProcessorTokenResponseBody instantiates a new PaymentProcessorTokenResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentProcessorTokenResponseBody() *PaymentProcessorTokenResponseBody {
	this := PaymentProcessorTokenResponseBody{}
	return &this
}

// NewPaymentProcessorTokenResponseBodyWithDefaults instantiates a new PaymentProcessorTokenResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentProcessorTokenResponseBodyWithDefaults() *PaymentProcessorTokenResponseBody {
	this := PaymentProcessorTokenResponseBody{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentProcessorTokenResponseBody) GetAccessToken() string {
	if o == nil || o.AccessToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentProcessorTokenResponseBody) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *PaymentProcessorTokenResponseBody) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *PaymentProcessorTokenResponseBody) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *PaymentProcessorTokenResponseBody) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *PaymentProcessorTokenResponseBody) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentProcessorTokenResponseBody) GetScope() string {
	if o == nil || o.Scope.Get() == nil {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentProcessorTokenResponseBody) GetScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *PaymentProcessorTokenResponseBody) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *PaymentProcessorTokenResponseBody) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *PaymentProcessorTokenResponseBody) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *PaymentProcessorTokenResponseBody) UnsetScope() {
	o.Scope.Unset()
}

// GetTokenType returns the TokenType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentProcessorTokenResponseBody) GetTokenType() string {
	if o == nil || o.TokenType.Get() == nil {
		var ret string
		return ret
	}
	return *o.TokenType.Get()
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentProcessorTokenResponseBody) GetTokenTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TokenType.Get(), o.TokenType.IsSet()
}

// HasTokenType returns a boolean if a field has been set.
func (o *PaymentProcessorTokenResponseBody) HasTokenType() bool {
	if o != nil && o.TokenType.IsSet() {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given NullableString and assigns it to the TokenType field.
func (o *PaymentProcessorTokenResponseBody) SetTokenType(v string) {
	o.TokenType.Set(&v)
}
// SetTokenTypeNil sets the value for TokenType to be an explicit nil
func (o *PaymentProcessorTokenResponseBody) SetTokenTypeNil() {
	o.TokenType.Set(nil)
}

// UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil
func (o *PaymentProcessorTokenResponseBody) UnsetTokenType() {
	o.TokenType.Unset()
}

func (o PaymentProcessorTokenResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	if o.TokenType.IsSet() {
		toSerialize["token_type"] = o.TokenType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentProcessorTokenResponseBody struct {
	value *PaymentProcessorTokenResponseBody
	isSet bool
}

func (v NullablePaymentProcessorTokenResponseBody) Get() *PaymentProcessorTokenResponseBody {
	return v.value
}

func (v *NullablePaymentProcessorTokenResponseBody) Set(val *PaymentProcessorTokenResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentProcessorTokenResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentProcessorTokenResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentProcessorTokenResponseBody(val *PaymentProcessorTokenResponseBody) *NullablePaymentProcessorTokenResponseBody {
	return &NullablePaymentProcessorTokenResponseBody{value: val, isSet: true}
}

func (v NullablePaymentProcessorTokenResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentProcessorTokenResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


