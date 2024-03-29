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

// checks if the CategoriesResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoriesResponseBody{}

// CategoriesResponseBody struct for CategoriesResponseBody
type CategoriesResponseBody struct {
	Categories []CategoryResponse `json:"categories,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewCategoriesResponseBody instantiates a new CategoriesResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoriesResponseBody() *CategoriesResponseBody {
	this := CategoriesResponseBody{}
	return &this
}

// NewCategoriesResponseBodyWithDefaults instantiates a new CategoriesResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoriesResponseBodyWithDefaults() *CategoriesResponseBody {
	this := CategoriesResponseBody{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *CategoriesResponseBody) GetCategories() []CategoryResponse {
	if o == nil || IsNil(o.Categories) {
		var ret []CategoryResponse
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoriesResponseBody) GetCategoriesOk() ([]CategoryResponse, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *CategoriesResponseBody) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []CategoryResponse and assigns it to the Categories field.
func (o *CategoriesResponseBody) SetCategories(v []CategoryResponse) {
	o.Categories = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *CategoriesResponseBody) GetPagination() PaginationResponse {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoriesResponseBody) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *CategoriesResponseBody) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *CategoriesResponseBody) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o CategoriesResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoriesResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableCategoriesResponseBody struct {
	value *CategoriesResponseBody
	isSet bool
}

func (v NullableCategoriesResponseBody) Get() *CategoriesResponseBody {
	return v.value
}

func (v *NullableCategoriesResponseBody) Set(val *CategoriesResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoriesResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoriesResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoriesResponseBody(val *CategoriesResponseBody) *NullableCategoriesResponseBody {
	return &NullableCategoriesResponseBody{value: val, isSet: true}
}

func (v NullableCategoriesResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoriesResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


