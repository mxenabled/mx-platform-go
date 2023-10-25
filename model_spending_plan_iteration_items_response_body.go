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

// checks if the SpendingPlanIterationItemsResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpendingPlanIterationItemsResponseBody{}

// SpendingPlanIterationItemsResponseBody struct for SpendingPlanIterationItemsResponseBody
type SpendingPlanIterationItemsResponseBody struct {
	IterationItems []SpendingPlanIterationItemResponse `json:"iteration_items,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewSpendingPlanIterationItemsResponseBody instantiates a new SpendingPlanIterationItemsResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendingPlanIterationItemsResponseBody() *SpendingPlanIterationItemsResponseBody {
	this := SpendingPlanIterationItemsResponseBody{}
	return &this
}

// NewSpendingPlanIterationItemsResponseBodyWithDefaults instantiates a new SpendingPlanIterationItemsResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendingPlanIterationItemsResponseBodyWithDefaults() *SpendingPlanIterationItemsResponseBody {
	this := SpendingPlanIterationItemsResponseBody{}
	return &this
}

// GetIterationItems returns the IterationItems field value if set, zero value otherwise.
func (o *SpendingPlanIterationItemsResponseBody) GetIterationItems() []SpendingPlanIterationItemResponse {
	if o == nil || IsNil(o.IterationItems) {
		var ret []SpendingPlanIterationItemResponse
		return ret
	}
	return o.IterationItems
}

// GetIterationItemsOk returns a tuple with the IterationItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingPlanIterationItemsResponseBody) GetIterationItemsOk() ([]SpendingPlanIterationItemResponse, bool) {
	if o == nil || IsNil(o.IterationItems) {
		return nil, false
	}
	return o.IterationItems, true
}

// HasIterationItems returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemsResponseBody) HasIterationItems() bool {
	if o != nil && !IsNil(o.IterationItems) {
		return true
	}

	return false
}

// SetIterationItems gets a reference to the given []SpendingPlanIterationItemResponse and assigns it to the IterationItems field.
func (o *SpendingPlanIterationItemsResponseBody) SetIterationItems(v []SpendingPlanIterationItemResponse) {
	o.IterationItems = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *SpendingPlanIterationItemsResponseBody) GetPagination() PaginationResponse {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingPlanIterationItemsResponseBody) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemsResponseBody) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *SpendingPlanIterationItemsResponseBody) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o SpendingPlanIterationItemsResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpendingPlanIterationItemsResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IterationItems) {
		toSerialize["iteration_items"] = o.IterationItems
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableSpendingPlanIterationItemsResponseBody struct {
	value *SpendingPlanIterationItemsResponseBody
	isSet bool
}

func (v NullableSpendingPlanIterationItemsResponseBody) Get() *SpendingPlanIterationItemsResponseBody {
	return v.value
}

func (v *NullableSpendingPlanIterationItemsResponseBody) Set(val *SpendingPlanIterationItemsResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendingPlanIterationItemsResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendingPlanIterationItemsResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendingPlanIterationItemsResponseBody(val *SpendingPlanIterationItemsResponseBody) *NullableSpendingPlanIterationItemsResponseBody {
	return &NullableSpendingPlanIterationItemsResponseBody{value: val, isSet: true}
}

func (v NullableSpendingPlanIterationItemsResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendingPlanIterationItemsResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


