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

// checks if the InstitutionResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstitutionResponseBody{}

// InstitutionResponseBody struct for InstitutionResponseBody
type InstitutionResponseBody struct {
	Institution *InstitutionResponse `json:"institution,omitempty"`
}

// NewInstitutionResponseBody instantiates a new InstitutionResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionResponseBody() *InstitutionResponseBody {
	this := InstitutionResponseBody{}
	return &this
}

// NewInstitutionResponseBodyWithDefaults instantiates a new InstitutionResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionResponseBodyWithDefaults() *InstitutionResponseBody {
	this := InstitutionResponseBody{}
	return &this
}

// GetInstitution returns the Institution field value if set, zero value otherwise.
func (o *InstitutionResponseBody) GetInstitution() InstitutionResponse {
	if o == nil || IsNil(o.Institution) {
		var ret InstitutionResponse
		return ret
	}
	return *o.Institution
}

// GetInstitutionOk returns a tuple with the Institution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionResponseBody) GetInstitutionOk() (*InstitutionResponse, bool) {
	if o == nil || IsNil(o.Institution) {
		return nil, false
	}
	return o.Institution, true
}

// HasInstitution returns a boolean if a field has been set.
func (o *InstitutionResponseBody) HasInstitution() bool {
	if o != nil && !IsNil(o.Institution) {
		return true
	}

	return false
}

// SetInstitution gets a reference to the given InstitutionResponse and assigns it to the Institution field.
func (o *InstitutionResponseBody) SetInstitution(v InstitutionResponse) {
	o.Institution = &v
}

func (o InstitutionResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstitutionResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Institution) {
		toSerialize["institution"] = o.Institution
	}
	return toSerialize, nil
}

type NullableInstitutionResponseBody struct {
	value *InstitutionResponseBody
	isSet bool
}

func (v NullableInstitutionResponseBody) Get() *InstitutionResponseBody {
	return v.value
}

func (v *NullableInstitutionResponseBody) Set(val *InstitutionResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionResponseBody(val *InstitutionResponseBody) *NullableInstitutionResponseBody {
	return &NullableInstitutionResponseBody{value: val, isSet: true}
}

func (v NullableInstitutionResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


