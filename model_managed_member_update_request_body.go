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

// checks if the ManagedMemberUpdateRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedMemberUpdateRequestBody{}

// ManagedMemberUpdateRequestBody struct for ManagedMemberUpdateRequestBody
type ManagedMemberUpdateRequestBody struct {
	Member *ManagedMemberUpdateRequest `json:"member,omitempty"`
}

// NewManagedMemberUpdateRequestBody instantiates a new ManagedMemberUpdateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedMemberUpdateRequestBody() *ManagedMemberUpdateRequestBody {
	this := ManagedMemberUpdateRequestBody{}
	return &this
}

// NewManagedMemberUpdateRequestBodyWithDefaults instantiates a new ManagedMemberUpdateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedMemberUpdateRequestBodyWithDefaults() *ManagedMemberUpdateRequestBody {
	this := ManagedMemberUpdateRequestBody{}
	return &this
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *ManagedMemberUpdateRequestBody) GetMember() ManagedMemberUpdateRequest {
	if o == nil || IsNil(o.Member) {
		var ret ManagedMemberUpdateRequest
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMemberUpdateRequestBody) GetMemberOk() (*ManagedMemberUpdateRequest, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *ManagedMemberUpdateRequestBody) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given ManagedMemberUpdateRequest and assigns it to the Member field.
func (o *ManagedMemberUpdateRequestBody) SetMember(v ManagedMemberUpdateRequest) {
	o.Member = &v
}

func (o ManagedMemberUpdateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedMemberUpdateRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	return toSerialize, nil
}

type NullableManagedMemberUpdateRequestBody struct {
	value *ManagedMemberUpdateRequestBody
	isSet bool
}

func (v NullableManagedMemberUpdateRequestBody) Get() *ManagedMemberUpdateRequestBody {
	return v.value
}

func (v *NullableManagedMemberUpdateRequestBody) Set(val *ManagedMemberUpdateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedMemberUpdateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedMemberUpdateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedMemberUpdateRequestBody(val *ManagedMemberUpdateRequestBody) *NullableManagedMemberUpdateRequestBody {
	return &NullableManagedMemberUpdateRequestBody{value: val, isSet: true}
}

func (v NullableManagedMemberUpdateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedMemberUpdateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


