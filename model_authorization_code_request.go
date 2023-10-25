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

// checks if the AuthorizationCodeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationCodeRequest{}

// AuthorizationCodeRequest struct for AuthorizationCodeRequest
type AuthorizationCodeRequest struct {
	Scope NullableString `json:"scope,omitempty"`
}

// NewAuthorizationCodeRequest instantiates a new AuthorizationCodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationCodeRequest() *AuthorizationCodeRequest {
	this := AuthorizationCodeRequest{}
	return &this
}

// NewAuthorizationCodeRequestWithDefaults instantiates a new AuthorizationCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationCodeRequestWithDefaults() *AuthorizationCodeRequest {
	this := AuthorizationCodeRequest{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizationCodeRequest) GetScope() string {
	if o == nil || IsNil(o.Scope.Get()) {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizationCodeRequest) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *AuthorizationCodeRequest) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *AuthorizationCodeRequest) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *AuthorizationCodeRequest) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *AuthorizationCodeRequest) UnsetScope() {
	o.Scope.Unset()
}

func (o AuthorizationCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationCodeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	return toSerialize, nil
}

type NullableAuthorizationCodeRequest struct {
	value *AuthorizationCodeRequest
	isSet bool
}

func (v NullableAuthorizationCodeRequest) Get() *AuthorizationCodeRequest {
	return v.value
}

func (v *NullableAuthorizationCodeRequest) Set(val *AuthorizationCodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationCodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationCodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationCodeRequest(val *AuthorizationCodeRequest) *NullableAuthorizationCodeRequest {
	return &NullableAuthorizationCodeRequest{value: val, isSet: true}
}

func (v NullableAuthorizationCodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationCodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


