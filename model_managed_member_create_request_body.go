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

// ManagedMemberCreateRequestBody struct for ManagedMemberCreateRequestBody
type ManagedMemberCreateRequestBody struct {
	Member *ManagedMemberCreateRequest `json:"member,omitempty"`
}

// NewManagedMemberCreateRequestBody instantiates a new ManagedMemberCreateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedMemberCreateRequestBody() *ManagedMemberCreateRequestBody {
	this := ManagedMemberCreateRequestBody{}
	return &this
}

// NewManagedMemberCreateRequestBodyWithDefaults instantiates a new ManagedMemberCreateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedMemberCreateRequestBodyWithDefaults() *ManagedMemberCreateRequestBody {
	this := ManagedMemberCreateRequestBody{}
	return &this
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *ManagedMemberCreateRequestBody) GetMember() ManagedMemberCreateRequest {
	if o == nil || o.Member == nil {
		var ret ManagedMemberCreateRequest
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMemberCreateRequestBody) GetMemberOk() (*ManagedMemberCreateRequest, bool) {
	if o == nil || o.Member == nil {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *ManagedMemberCreateRequestBody) HasMember() bool {
	if o != nil && o.Member != nil {
		return true
	}

	return false
}

// SetMember gets a reference to the given ManagedMemberCreateRequest and assigns it to the Member field.
func (o *ManagedMemberCreateRequestBody) SetMember(v ManagedMemberCreateRequest) {
	o.Member = &v
}

func (o ManagedMemberCreateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Member != nil {
		toSerialize["member"] = o.Member
	}
	return json.Marshal(toSerialize)
}

type NullableManagedMemberCreateRequestBody struct {
	value *ManagedMemberCreateRequestBody
	isSet bool
}

func (v NullableManagedMemberCreateRequestBody) Get() *ManagedMemberCreateRequestBody {
	return v.value
}

func (v *NullableManagedMemberCreateRequestBody) Set(val *ManagedMemberCreateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedMemberCreateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedMemberCreateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedMemberCreateRequestBody(val *ManagedMemberCreateRequestBody) *NullableManagedMemberCreateRequestBody {
	return &NullableManagedMemberCreateRequestBody{value: val, isSet: true}
}

func (v NullableManagedMemberCreateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedMemberCreateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

