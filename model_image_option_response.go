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

// ImageOptionResponse struct for ImageOptionResponse
type ImageOptionResponse struct {
	DataUri NullableString `json:"data_uri,omitempty"`
	Label NullableString `json:"label,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewImageOptionResponse instantiates a new ImageOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageOptionResponse() *ImageOptionResponse {
	this := ImageOptionResponse{}
	return &this
}

// NewImageOptionResponseWithDefaults instantiates a new ImageOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageOptionResponseWithDefaults() *ImageOptionResponse {
	this := ImageOptionResponse{}
	return &this
}

// GetDataUri returns the DataUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageOptionResponse) GetDataUri() string {
	if o == nil || o.DataUri.Get() == nil {
		var ret string
		return ret
	}
	return *o.DataUri.Get()
}

// GetDataUriOk returns a tuple with the DataUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageOptionResponse) GetDataUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataUri.Get(), o.DataUri.IsSet()
}

// HasDataUri returns a boolean if a field has been set.
func (o *ImageOptionResponse) HasDataUri() bool {
	if o != nil && o.DataUri.IsSet() {
		return true
	}

	return false
}

// SetDataUri gets a reference to the given NullableString and assigns it to the DataUri field.
func (o *ImageOptionResponse) SetDataUri(v string) {
	o.DataUri.Set(&v)
}
// SetDataUriNil sets the value for DataUri to be an explicit nil
func (o *ImageOptionResponse) SetDataUriNil() {
	o.DataUri.Set(nil)
}

// UnsetDataUri ensures that no value is present for DataUri, not even an explicit nil
func (o *ImageOptionResponse) UnsetDataUri() {
	o.DataUri.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageOptionResponse) GetLabel() string {
	if o == nil || o.Label.Get() == nil {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageOptionResponse) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *ImageOptionResponse) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *ImageOptionResponse) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *ImageOptionResponse) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *ImageOptionResponse) UnsetLabel() {
	o.Label.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageOptionResponse) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageOptionResponse) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ImageOptionResponse) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ImageOptionResponse) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ImageOptionResponse) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ImageOptionResponse) UnsetValue() {
	o.Value.Unset()
}

func (o ImageOptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataUri.IsSet() {
		toSerialize["data_uri"] = o.DataUri.Get()
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableImageOptionResponse struct {
	value *ImageOptionResponse
	isSet bool
}

func (v NullableImageOptionResponse) Get() *ImageOptionResponse {
	return v.value
}

func (v *NullableImageOptionResponse) Set(val *ImageOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableImageOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableImageOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageOptionResponse(val *ImageOptionResponse) *NullableImageOptionResponse {
	return &NullableImageOptionResponse{value: val, isSet: true}
}

func (v NullableImageOptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

