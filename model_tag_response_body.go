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

// checks if the TagResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagResponseBody{}

// TagResponseBody struct for TagResponseBody
type TagResponseBody struct {
	Tag *TagResponse `json:"tag,omitempty"`
}

// NewTagResponseBody instantiates a new TagResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagResponseBody() *TagResponseBody {
	this := TagResponseBody{}
	return &this
}

// NewTagResponseBodyWithDefaults instantiates a new TagResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagResponseBodyWithDefaults() *TagResponseBody {
	this := TagResponseBody{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *TagResponseBody) GetTag() TagResponse {
	if o == nil || IsNil(o.Tag) {
		var ret TagResponse
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagResponseBody) GetTagOk() (*TagResponse, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *TagResponseBody) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given TagResponse and assigns it to the Tag field.
func (o *TagResponseBody) SetTag(v TagResponse) {
	o.Tag = &v
}

func (o TagResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullableTagResponseBody struct {
	value *TagResponseBody
	isSet bool
}

func (v NullableTagResponseBody) Get() *TagResponseBody {
	return v.value
}

func (v *NullableTagResponseBody) Set(val *TagResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTagResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTagResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagResponseBody(val *TagResponseBody) *NullableTagResponseBody {
	return &NullableTagResponseBody{value: val, isSet: true}
}

func (v NullableTagResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


