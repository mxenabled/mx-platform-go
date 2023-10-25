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

// checks if the OAuthWindowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthWindowResponse{}

// OAuthWindowResponse struct for OAuthWindowResponse
type OAuthWindowResponse struct {
	Guid NullableString `json:"guid,omitempty"`
	OauthWindowUri NullableString `json:"oauth_window_uri,omitempty"`
}

// NewOAuthWindowResponse instantiates a new OAuthWindowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthWindowResponse() *OAuthWindowResponse {
	this := OAuthWindowResponse{}
	return &this
}

// NewOAuthWindowResponseWithDefaults instantiates a new OAuthWindowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthWindowResponseWithDefaults() *OAuthWindowResponse {
	this := OAuthWindowResponse{}
	return &this
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthWindowResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthWindowResponse) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *OAuthWindowResponse) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *OAuthWindowResponse) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *OAuthWindowResponse) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *OAuthWindowResponse) UnsetGuid() {
	o.Guid.Unset()
}

// GetOauthWindowUri returns the OauthWindowUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthWindowResponse) GetOauthWindowUri() string {
	if o == nil || IsNil(o.OauthWindowUri.Get()) {
		var ret string
		return ret
	}
	return *o.OauthWindowUri.Get()
}

// GetOauthWindowUriOk returns a tuple with the OauthWindowUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthWindowResponse) GetOauthWindowUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OauthWindowUri.Get(), o.OauthWindowUri.IsSet()
}

// HasOauthWindowUri returns a boolean if a field has been set.
func (o *OAuthWindowResponse) HasOauthWindowUri() bool {
	if o != nil && o.OauthWindowUri.IsSet() {
		return true
	}

	return false
}

// SetOauthWindowUri gets a reference to the given NullableString and assigns it to the OauthWindowUri field.
func (o *OAuthWindowResponse) SetOauthWindowUri(v string) {
	o.OauthWindowUri.Set(&v)
}
// SetOauthWindowUriNil sets the value for OauthWindowUri to be an explicit nil
func (o *OAuthWindowResponse) SetOauthWindowUriNil() {
	o.OauthWindowUri.Set(nil)
}

// UnsetOauthWindowUri ensures that no value is present for OauthWindowUri, not even an explicit nil
func (o *OAuthWindowResponse) UnsetOauthWindowUri() {
	o.OauthWindowUri.Unset()
}

func (o OAuthWindowResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthWindowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.OauthWindowUri.IsSet() {
		toSerialize["oauth_window_uri"] = o.OauthWindowUri.Get()
	}
	return toSerialize, nil
}

type NullableOAuthWindowResponse struct {
	value *OAuthWindowResponse
	isSet bool
}

func (v NullableOAuthWindowResponse) Get() *OAuthWindowResponse {
	return v.value
}

func (v *NullableOAuthWindowResponse) Set(val *OAuthWindowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthWindowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthWindowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthWindowResponse(val *OAuthWindowResponse) *NullableOAuthWindowResponse {
	return &NullableOAuthWindowResponse{value: val, isSet: true}
}

func (v NullableOAuthWindowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthWindowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


