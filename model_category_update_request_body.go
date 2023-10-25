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

// checks if the CategoryUpdateRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryUpdateRequestBody{}

// CategoryUpdateRequestBody struct for CategoryUpdateRequestBody
type CategoryUpdateRequestBody struct {
	Category *CategoryUpdateRequest `json:"category,omitempty"`
}

// NewCategoryUpdateRequestBody instantiates a new CategoryUpdateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryUpdateRequestBody() *CategoryUpdateRequestBody {
	this := CategoryUpdateRequestBody{}
	return &this
}

// NewCategoryUpdateRequestBodyWithDefaults instantiates a new CategoryUpdateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryUpdateRequestBodyWithDefaults() *CategoryUpdateRequestBody {
	this := CategoryUpdateRequestBody{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CategoryUpdateRequestBody) GetCategory() CategoryUpdateRequest {
	if o == nil || IsNil(o.Category) {
		var ret CategoryUpdateRequest
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryUpdateRequestBody) GetCategoryOk() (*CategoryUpdateRequest, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CategoryUpdateRequestBody) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given CategoryUpdateRequest and assigns it to the Category field.
func (o *CategoryUpdateRequestBody) SetCategory(v CategoryUpdateRequest) {
	o.Category = &v
}

func (o CategoryUpdateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryUpdateRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	return toSerialize, nil
}

type NullableCategoryUpdateRequestBody struct {
	value *CategoryUpdateRequestBody
	isSet bool
}

func (v NullableCategoryUpdateRequestBody) Get() *CategoryUpdateRequestBody {
	return v.value
}

func (v *NullableCategoryUpdateRequestBody) Set(val *CategoryUpdateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryUpdateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryUpdateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryUpdateRequestBody(val *CategoryUpdateRequestBody) *NullableCategoryUpdateRequestBody {
	return &NullableCategoryUpdateRequestBody{value: val, isSet: true}
}

func (v NullableCategoryUpdateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryUpdateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


