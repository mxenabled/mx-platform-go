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

// TagCreateRequestBody struct for TagCreateRequestBody
type TagCreateRequestBody struct {
	Tag *TagCreateRequest `json:"tag,omitempty"`
}

// NewTagCreateRequestBody instantiates a new TagCreateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagCreateRequestBody() *TagCreateRequestBody {
	this := TagCreateRequestBody{}
	return &this
}

// NewTagCreateRequestBodyWithDefaults instantiates a new TagCreateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagCreateRequestBodyWithDefaults() *TagCreateRequestBody {
	this := TagCreateRequestBody{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *TagCreateRequestBody) GetTag() TagCreateRequest {
	if o == nil || o.Tag == nil {
		var ret TagCreateRequest
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagCreateRequestBody) GetTagOk() (*TagCreateRequest, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *TagCreateRequestBody) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given TagCreateRequest and assigns it to the Tag field.
func (o *TagCreateRequestBody) SetTag(v TagCreateRequest) {
	o.Tag = &v
}

func (o TagCreateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableTagCreateRequestBody struct {
	value *TagCreateRequestBody
	isSet bool
}

func (v NullableTagCreateRequestBody) Get() *TagCreateRequestBody {
	return v.value
}

func (v *NullableTagCreateRequestBody) Set(val *TagCreateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTagCreateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTagCreateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagCreateRequestBody(val *TagCreateRequestBody) *NullableTagCreateRequestBody {
	return &NullableTagCreateRequestBody{value: val, isSet: true}
}

func (v NullableTagCreateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagCreateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


