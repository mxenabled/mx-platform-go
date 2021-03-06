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

// ConnectWidgetRequestBody struct for ConnectWidgetRequestBody
type ConnectWidgetRequestBody struct {
	Config *ConnectWidgetRequest `json:"config,omitempty"`
}

// NewConnectWidgetRequestBody instantiates a new ConnectWidgetRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectWidgetRequestBody() *ConnectWidgetRequestBody {
	this := ConnectWidgetRequestBody{}
	return &this
}

// NewConnectWidgetRequestBodyWithDefaults instantiates a new ConnectWidgetRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectWidgetRequestBodyWithDefaults() *ConnectWidgetRequestBody {
	this := ConnectWidgetRequestBody{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ConnectWidgetRequestBody) GetConfig() ConnectWidgetRequest {
	if o == nil || o.Config == nil {
		var ret ConnectWidgetRequest
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectWidgetRequestBody) GetConfigOk() (*ConnectWidgetRequest, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ConnectWidgetRequestBody) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ConnectWidgetRequest and assigns it to the Config field.
func (o *ConnectWidgetRequestBody) SetConfig(v ConnectWidgetRequest) {
	o.Config = &v
}

func (o ConnectWidgetRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableConnectWidgetRequestBody struct {
	value *ConnectWidgetRequestBody
	isSet bool
}

func (v NullableConnectWidgetRequestBody) Get() *ConnectWidgetRequestBody {
	return v.value
}

func (v *NullableConnectWidgetRequestBody) Set(val *ConnectWidgetRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectWidgetRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectWidgetRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectWidgetRequestBody(val *ConnectWidgetRequestBody) *NullableConnectWidgetRequestBody {
	return &NullableConnectWidgetRequestBody{value: val, isSet: true}
}

func (v NullableConnectWidgetRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectWidgetRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


