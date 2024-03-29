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

// checks if the MerchantsResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantsResponseBody{}

// MerchantsResponseBody struct for MerchantsResponseBody
type MerchantsResponseBody struct {
	Merchants []MerchantResponse `json:"merchants,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewMerchantsResponseBody instantiates a new MerchantsResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantsResponseBody() *MerchantsResponseBody {
	this := MerchantsResponseBody{}
	return &this
}

// NewMerchantsResponseBodyWithDefaults instantiates a new MerchantsResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantsResponseBodyWithDefaults() *MerchantsResponseBody {
	this := MerchantsResponseBody{}
	return &this
}

// GetMerchants returns the Merchants field value if set, zero value otherwise.
func (o *MerchantsResponseBody) GetMerchants() []MerchantResponse {
	if o == nil || IsNil(o.Merchants) {
		var ret []MerchantResponse
		return ret
	}
	return o.Merchants
}

// GetMerchantsOk returns a tuple with the Merchants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantsResponseBody) GetMerchantsOk() ([]MerchantResponse, bool) {
	if o == nil || IsNil(o.Merchants) {
		return nil, false
	}
	return o.Merchants, true
}

// HasMerchants returns a boolean if a field has been set.
func (o *MerchantsResponseBody) HasMerchants() bool {
	if o != nil && !IsNil(o.Merchants) {
		return true
	}

	return false
}

// SetMerchants gets a reference to the given []MerchantResponse and assigns it to the Merchants field.
func (o *MerchantsResponseBody) SetMerchants(v []MerchantResponse) {
	o.Merchants = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *MerchantsResponseBody) GetPagination() PaginationResponse {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantsResponseBody) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *MerchantsResponseBody) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *MerchantsResponseBody) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o MerchantsResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantsResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Merchants) {
		toSerialize["merchants"] = o.Merchants
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableMerchantsResponseBody struct {
	value *MerchantsResponseBody
	isSet bool
}

func (v NullableMerchantsResponseBody) Get() *MerchantsResponseBody {
	return v.value
}

func (v *NullableMerchantsResponseBody) Set(val *MerchantsResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantsResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantsResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantsResponseBody(val *MerchantsResponseBody) *NullableMerchantsResponseBody {
	return &NullableMerchantsResponseBody{value: val, isSet: true}
}

func (v NullableMerchantsResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantsResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


