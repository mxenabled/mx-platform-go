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

// TaggingUpdateRequestBody struct for TaggingUpdateRequestBody
type TaggingUpdateRequestBody struct {
	Tagging *TaggingUpdateRequest `json:"tagging,omitempty"`
}

// NewTaggingUpdateRequestBody instantiates a new TaggingUpdateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaggingUpdateRequestBody() *TaggingUpdateRequestBody {
	this := TaggingUpdateRequestBody{}
	return &this
}

// NewTaggingUpdateRequestBodyWithDefaults instantiates a new TaggingUpdateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaggingUpdateRequestBodyWithDefaults() *TaggingUpdateRequestBody {
	this := TaggingUpdateRequestBody{}
	return &this
}

// GetTagging returns the Tagging field value if set, zero value otherwise.
func (o *TaggingUpdateRequestBody) GetTagging() TaggingUpdateRequest {
	if o == nil || o.Tagging == nil {
		var ret TaggingUpdateRequest
		return ret
	}
	return *o.Tagging
}

// GetTaggingOk returns a tuple with the Tagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaggingUpdateRequestBody) GetTaggingOk() (*TaggingUpdateRequest, bool) {
	if o == nil || o.Tagging == nil {
		return nil, false
	}
	return o.Tagging, true
}

// HasTagging returns a boolean if a field has been set.
func (o *TaggingUpdateRequestBody) HasTagging() bool {
	if o != nil && o.Tagging != nil {
		return true
	}

	return false
}

// SetTagging gets a reference to the given TaggingUpdateRequest and assigns it to the Tagging field.
func (o *TaggingUpdateRequestBody) SetTagging(v TaggingUpdateRequest) {
	o.Tagging = &v
}

func (o TaggingUpdateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tagging != nil {
		toSerialize["tagging"] = o.Tagging
	}
	return json.Marshal(toSerialize)
}

type NullableTaggingUpdateRequestBody struct {
	value *TaggingUpdateRequestBody
	isSet bool
}

func (v NullableTaggingUpdateRequestBody) Get() *TaggingUpdateRequestBody {
	return v.value
}

func (v *NullableTaggingUpdateRequestBody) Set(val *TaggingUpdateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTaggingUpdateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTaggingUpdateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaggingUpdateRequestBody(val *TaggingUpdateRequestBody) *NullableTaggingUpdateRequestBody {
	return &NullableTaggingUpdateRequestBody{value: val, isSet: true}
}

func (v NullableTaggingUpdateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaggingUpdateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


