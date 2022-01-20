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

// CategoryResponseBody struct for CategoryResponseBody
type CategoryResponseBody struct {
	Category *CategoryResponse `json:"category,omitempty"`
}

// NewCategoryResponseBody instantiates a new CategoryResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryResponseBody() *CategoryResponseBody {
	this := CategoryResponseBody{}
	return &this
}

// NewCategoryResponseBodyWithDefaults instantiates a new CategoryResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryResponseBodyWithDefaults() *CategoryResponseBody {
	this := CategoryResponseBody{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CategoryResponseBody) GetCategory() CategoryResponse {
	if o == nil || o.Category == nil {
		var ret CategoryResponse
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryResponseBody) GetCategoryOk() (*CategoryResponse, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CategoryResponseBody) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given CategoryResponse and assigns it to the Category field.
func (o *CategoryResponseBody) SetCategory(v CategoryResponse) {
	o.Category = &v
}

func (o CategoryResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryResponseBody struct {
	value *CategoryResponseBody
	isSet bool
}

func (v NullableCategoryResponseBody) Get() *CategoryResponseBody {
	return v.value
}

func (v *NullableCategoryResponseBody) Set(val *CategoryResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryResponseBody(val *CategoryResponseBody) *NullableCategoryResponseBody {
	return &NullableCategoryResponseBody{value: val, isSet: true}
}

func (v NullableCategoryResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

