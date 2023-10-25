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

// checks if the InstitutionsResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstitutionsResponseBody{}

// InstitutionsResponseBody struct for InstitutionsResponseBody
type InstitutionsResponseBody struct {
	Institutions []InstitutionResponse `json:"institutions,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewInstitutionsResponseBody instantiates a new InstitutionsResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsResponseBody() *InstitutionsResponseBody {
	this := InstitutionsResponseBody{}
	return &this
}

// NewInstitutionsResponseBodyWithDefaults instantiates a new InstitutionsResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsResponseBodyWithDefaults() *InstitutionsResponseBody {
	this := InstitutionsResponseBody{}
	return &this
}

// GetInstitutions returns the Institutions field value if set, zero value otherwise.
func (o *InstitutionsResponseBody) GetInstitutions() []InstitutionResponse {
	if o == nil || IsNil(o.Institutions) {
		var ret []InstitutionResponse
		return ret
	}
	return o.Institutions
}

// GetInstitutionsOk returns a tuple with the Institutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsResponseBody) GetInstitutionsOk() ([]InstitutionResponse, bool) {
	if o == nil || IsNil(o.Institutions) {
		return nil, false
	}
	return o.Institutions, true
}

// HasInstitutions returns a boolean if a field has been set.
func (o *InstitutionsResponseBody) HasInstitutions() bool {
	if o != nil && !IsNil(o.Institutions) {
		return true
	}

	return false
}

// SetInstitutions gets a reference to the given []InstitutionResponse and assigns it to the Institutions field.
func (o *InstitutionsResponseBody) SetInstitutions(v []InstitutionResponse) {
	o.Institutions = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *InstitutionsResponseBody) GetPagination() PaginationResponse {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsResponseBody) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *InstitutionsResponseBody) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *InstitutionsResponseBody) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o InstitutionsResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstitutionsResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Institutions) {
		toSerialize["institutions"] = o.Institutions
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableInstitutionsResponseBody struct {
	value *InstitutionsResponseBody
	isSet bool
}

func (v NullableInstitutionsResponseBody) Get() *InstitutionsResponseBody {
	return v.value
}

func (v *NullableInstitutionsResponseBody) Set(val *InstitutionsResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsResponseBody(val *InstitutionsResponseBody) *NullableInstitutionsResponseBody {
	return &NullableInstitutionsResponseBody{value: val, isSet: true}
}

func (v NullableInstitutionsResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


