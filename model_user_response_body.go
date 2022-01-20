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

// UserResponseBody struct for UserResponseBody
type UserResponseBody struct {
	User *UserResponse `json:"user,omitempty"`
}

// NewUserResponseBody instantiates a new UserResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponseBody() *UserResponseBody {
	this := UserResponseBody{}
	return &this
}

// NewUserResponseBodyWithDefaults instantiates a new UserResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseBodyWithDefaults() *UserResponseBody {
	this := UserResponseBody{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserResponseBody) GetUser() UserResponse {
	if o == nil || o.User == nil {
		var ret UserResponse
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseBody) GetUserOk() (*UserResponse, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserResponseBody) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserResponse and assigns it to the User field.
func (o *UserResponseBody) SetUser(v UserResponse) {
	o.User = &v
}

func (o UserResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableUserResponseBody struct {
	value *UserResponseBody
	isSet bool
}

func (v NullableUserResponseBody) Get() *UserResponseBody {
	return v.value
}

func (v *NullableUserResponseBody) Set(val *UserResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponseBody(val *UserResponseBody) *NullableUserResponseBody {
	return &NullableUserResponseBody{value: val, isSet: true}
}

func (v NullableUserResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

