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

// CredentialsResponseBody struct for CredentialsResponseBody
type CredentialsResponseBody struct {
	Credentials *[]CredentialResponse `json:"credentials,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewCredentialsResponseBody instantiates a new CredentialsResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsResponseBody() *CredentialsResponseBody {
	this := CredentialsResponseBody{}
	return &this
}

// NewCredentialsResponseBodyWithDefaults instantiates a new CredentialsResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsResponseBodyWithDefaults() *CredentialsResponseBody {
	this := CredentialsResponseBody{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *CredentialsResponseBody) GetCredentials() []CredentialResponse {
	if o == nil || o.Credentials == nil {
		var ret []CredentialResponse
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsResponseBody) GetCredentialsOk() (*[]CredentialResponse, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *CredentialsResponseBody) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []CredentialResponse and assigns it to the Credentials field.
func (o *CredentialsResponseBody) SetCredentials(v []CredentialResponse) {
	o.Credentials = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *CredentialsResponseBody) GetPagination() PaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsResponseBody) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *CredentialsResponseBody) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *CredentialsResponseBody) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o CredentialsResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialsResponseBody struct {
	value *CredentialsResponseBody
	isSet bool
}

func (v NullableCredentialsResponseBody) Get() *CredentialsResponseBody {
	return v.value
}

func (v *NullableCredentialsResponseBody) Set(val *CredentialsResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsResponseBody(val *CredentialsResponseBody) *NullableCredentialsResponseBody {
	return &NullableCredentialsResponseBody{value: val, isSet: true}
}

func (v NullableCredentialsResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

