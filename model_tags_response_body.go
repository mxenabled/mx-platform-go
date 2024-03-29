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

// checks if the TagsResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagsResponseBody{}

// TagsResponseBody struct for TagsResponseBody
type TagsResponseBody struct {
	Pagination *PaginationResponse `json:"pagination,omitempty"`
	Tags []TagResponse `json:"tags,omitempty"`
}

// NewTagsResponseBody instantiates a new TagsResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagsResponseBody() *TagsResponseBody {
	this := TagsResponseBody{}
	return &this
}

// NewTagsResponseBodyWithDefaults instantiates a new TagsResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagsResponseBodyWithDefaults() *TagsResponseBody {
	this := TagsResponseBody{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *TagsResponseBody) GetPagination() PaginationResponse {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsResponseBody) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *TagsResponseBody) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *TagsResponseBody) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TagsResponseBody) GetTags() []TagResponse {
	if o == nil || IsNil(o.Tags) {
		var ret []TagResponse
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsResponseBody) GetTagsOk() ([]TagResponse, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TagsResponseBody) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagResponse and assigns it to the Tags field.
func (o *TagsResponseBody) SetTags(v []TagResponse) {
	o.Tags = v
}

func (o TagsResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagsResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableTagsResponseBody struct {
	value *TagsResponseBody
	isSet bool
}

func (v NullableTagsResponseBody) Get() *TagsResponseBody {
	return v.value
}

func (v *NullableTagsResponseBody) Set(val *TagsResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTagsResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTagsResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagsResponseBody(val *TagsResponseBody) *NullableTagsResponseBody {
	return &NullableTagsResponseBody{value: val, isSet: true}
}

func (v NullableTagsResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagsResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


