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

// checks if the AuthorizationCodeResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationCodeResponseBody{}

// AuthorizationCodeResponseBody struct for AuthorizationCodeResponseBody
type AuthorizationCodeResponseBody struct {
	AuthorizationCode []AuthorizationCodeResponse `json:"authorization_code,omitempty"`
}

// NewAuthorizationCodeResponseBody instantiates a new AuthorizationCodeResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationCodeResponseBody() *AuthorizationCodeResponseBody {
	this := AuthorizationCodeResponseBody{}
	return &this
}

// NewAuthorizationCodeResponseBodyWithDefaults instantiates a new AuthorizationCodeResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationCodeResponseBodyWithDefaults() *AuthorizationCodeResponseBody {
	this := AuthorizationCodeResponseBody{}
	return &this
}

// GetAuthorizationCode returns the AuthorizationCode field value if set, zero value otherwise.
func (o *AuthorizationCodeResponseBody) GetAuthorizationCode() []AuthorizationCodeResponse {
	if o == nil || IsNil(o.AuthorizationCode) {
		var ret []AuthorizationCodeResponse
		return ret
	}
	return o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationCodeResponseBody) GetAuthorizationCodeOk() ([]AuthorizationCodeResponse, bool) {
	if o == nil || IsNil(o.AuthorizationCode) {
		return nil, false
	}
	return o.AuthorizationCode, true
}

// HasAuthorizationCode returns a boolean if a field has been set.
func (o *AuthorizationCodeResponseBody) HasAuthorizationCode() bool {
	if o != nil && !IsNil(o.AuthorizationCode) {
		return true
	}

	return false
}

// SetAuthorizationCode gets a reference to the given []AuthorizationCodeResponse and assigns it to the AuthorizationCode field.
func (o *AuthorizationCodeResponseBody) SetAuthorizationCode(v []AuthorizationCodeResponse) {
	o.AuthorizationCode = v
}

func (o AuthorizationCodeResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationCodeResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationCode) {
		toSerialize["authorization_code"] = o.AuthorizationCode
	}
	return toSerialize, nil
}

type NullableAuthorizationCodeResponseBody struct {
	value *AuthorizationCodeResponseBody
	isSet bool
}

func (v NullableAuthorizationCodeResponseBody) Get() *AuthorizationCodeResponseBody {
	return v.value
}

func (v *NullableAuthorizationCodeResponseBody) Set(val *AuthorizationCodeResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationCodeResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationCodeResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationCodeResponseBody(val *AuthorizationCodeResponseBody) *NullableAuthorizationCodeResponseBody {
	return &NullableAuthorizationCodeResponseBody{value: val, isSet: true}
}

func (v NullableAuthorizationCodeResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationCodeResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


