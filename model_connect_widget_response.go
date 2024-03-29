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

// checks if the ConnectWidgetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectWidgetResponse{}

// ConnectWidgetResponse struct for ConnectWidgetResponse
type ConnectWidgetResponse struct {
	ConnectWidgetUrl NullableString `json:"connect_widget_url,omitempty"`
	Guid NullableString `json:"guid,omitempty"`
}

// NewConnectWidgetResponse instantiates a new ConnectWidgetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectWidgetResponse() *ConnectWidgetResponse {
	this := ConnectWidgetResponse{}
	return &this
}

// NewConnectWidgetResponseWithDefaults instantiates a new ConnectWidgetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectWidgetResponseWithDefaults() *ConnectWidgetResponse {
	this := ConnectWidgetResponse{}
	return &this
}

// GetConnectWidgetUrl returns the ConnectWidgetUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectWidgetResponse) GetConnectWidgetUrl() string {
	if o == nil || IsNil(o.ConnectWidgetUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectWidgetUrl.Get()
}

// GetConnectWidgetUrlOk returns a tuple with the ConnectWidgetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectWidgetResponse) GetConnectWidgetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectWidgetUrl.Get(), o.ConnectWidgetUrl.IsSet()
}

// HasConnectWidgetUrl returns a boolean if a field has been set.
func (o *ConnectWidgetResponse) HasConnectWidgetUrl() bool {
	if o != nil && o.ConnectWidgetUrl.IsSet() {
		return true
	}

	return false
}

// SetConnectWidgetUrl gets a reference to the given NullableString and assigns it to the ConnectWidgetUrl field.
func (o *ConnectWidgetResponse) SetConnectWidgetUrl(v string) {
	o.ConnectWidgetUrl.Set(&v)
}
// SetConnectWidgetUrlNil sets the value for ConnectWidgetUrl to be an explicit nil
func (o *ConnectWidgetResponse) SetConnectWidgetUrlNil() {
	o.ConnectWidgetUrl.Set(nil)
}

// UnsetConnectWidgetUrl ensures that no value is present for ConnectWidgetUrl, not even an explicit nil
func (o *ConnectWidgetResponse) UnsetConnectWidgetUrl() {
	o.ConnectWidgetUrl.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectWidgetResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectWidgetResponse) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *ConnectWidgetResponse) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *ConnectWidgetResponse) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *ConnectWidgetResponse) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *ConnectWidgetResponse) UnsetGuid() {
	o.Guid.Unset()
}

func (o ConnectWidgetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectWidgetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectWidgetUrl.IsSet() {
		toSerialize["connect_widget_url"] = o.ConnectWidgetUrl.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	return toSerialize, nil
}

type NullableConnectWidgetResponse struct {
	value *ConnectWidgetResponse
	isSet bool
}

func (v NullableConnectWidgetResponse) Get() *ConnectWidgetResponse {
	return v.value
}

func (v *NullableConnectWidgetResponse) Set(val *ConnectWidgetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectWidgetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectWidgetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectWidgetResponse(val *ConnectWidgetResponse) *NullableConnectWidgetResponse {
	return &NullableConnectWidgetResponse{value: val, isSet: true}
}

func (v NullableConnectWidgetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectWidgetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


