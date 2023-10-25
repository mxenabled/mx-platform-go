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

// checks if the TagUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagUpdateRequest{}

// TagUpdateRequest struct for TagUpdateRequest
type TagUpdateRequest struct {
	Name string `json:"name"`
}

// NewTagUpdateRequest instantiates a new TagUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagUpdateRequest(name string) *TagUpdateRequest {
	this := TagUpdateRequest{}
	this.Name = name
	return &this
}

// NewTagUpdateRequestWithDefaults instantiates a new TagUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagUpdateRequestWithDefaults() *TagUpdateRequest {
	this := TagUpdateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *TagUpdateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagUpdateRequest) SetName(v string) {
	o.Name = v
}

func (o TagUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableTagUpdateRequest struct {
	value *TagUpdateRequest
	isSet bool
}

func (v NullableTagUpdateRequest) Get() *TagUpdateRequest {
	return v.value
}

func (v *NullableTagUpdateRequest) Set(val *TagUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTagUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTagUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagUpdateRequest(val *TagUpdateRequest) *NullableTagUpdateRequest {
	return &NullableTagUpdateRequest{value: val, isSet: true}
}

func (v NullableTagUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


