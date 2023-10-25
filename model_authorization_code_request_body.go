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

// checks if the AuthorizationCodeRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationCodeRequestBody{}

// AuthorizationCodeRequestBody struct for AuthorizationCodeRequestBody
type AuthorizationCodeRequestBody struct {
	AuthorizationCode *AuthorizationCodeRequest `json:"authorization_code,omitempty"`
}

// NewAuthorizationCodeRequestBody instantiates a new AuthorizationCodeRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationCodeRequestBody() *AuthorizationCodeRequestBody {
	this := AuthorizationCodeRequestBody{}
	return &this
}

// NewAuthorizationCodeRequestBodyWithDefaults instantiates a new AuthorizationCodeRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationCodeRequestBodyWithDefaults() *AuthorizationCodeRequestBody {
	this := AuthorizationCodeRequestBody{}
	return &this
}

// GetAuthorizationCode returns the AuthorizationCode field value if set, zero value otherwise.
func (o *AuthorizationCodeRequestBody) GetAuthorizationCode() AuthorizationCodeRequest {
	if o == nil || IsNil(o.AuthorizationCode) {
		var ret AuthorizationCodeRequest
		return ret
	}
	return *o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationCodeRequestBody) GetAuthorizationCodeOk() (*AuthorizationCodeRequest, bool) {
	if o == nil || IsNil(o.AuthorizationCode) {
		return nil, false
	}
	return o.AuthorizationCode, true
}

// HasAuthorizationCode returns a boolean if a field has been set.
func (o *AuthorizationCodeRequestBody) HasAuthorizationCode() bool {
	if o != nil && !IsNil(o.AuthorizationCode) {
		return true
	}

	return false
}

// SetAuthorizationCode gets a reference to the given AuthorizationCodeRequest and assigns it to the AuthorizationCode field.
func (o *AuthorizationCodeRequestBody) SetAuthorizationCode(v AuthorizationCodeRequest) {
	o.AuthorizationCode = &v
}

func (o AuthorizationCodeRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationCodeRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationCode) {
		toSerialize["authorization_code"] = o.AuthorizationCode
	}
	return toSerialize, nil
}

type NullableAuthorizationCodeRequestBody struct {
	value *AuthorizationCodeRequestBody
	isSet bool
}

func (v NullableAuthorizationCodeRequestBody) Get() *AuthorizationCodeRequestBody {
	return v.value
}

func (v *NullableAuthorizationCodeRequestBody) Set(val *AuthorizationCodeRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationCodeRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationCodeRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationCodeRequestBody(val *AuthorizationCodeRequestBody) *NullableAuthorizationCodeRequestBody {
	return &NullableAuthorizationCodeRequestBody{value: val, isSet: true}
}

func (v NullableAuthorizationCodeRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationCodeRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


