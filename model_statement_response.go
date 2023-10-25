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

// checks if the StatementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatementResponse{}

// StatementResponse struct for StatementResponse
type StatementResponse struct {
	AccountGuid NullableString `json:"account_guid,omitempty"`
	ContentHash NullableString `json:"content_hash,omitempty"`
	CreatedAt NullableString `json:"created_at,omitempty"`
	Guid NullableString `json:"guid,omitempty"`
	MemberGuid NullableString `json:"member_guid,omitempty"`
	UpdatedAt NullableString `json:"updated_at,omitempty"`
	Uri NullableString `json:"uri,omitempty"`
	UserGuid NullableString `json:"user_guid,omitempty"`
}

// NewStatementResponse instantiates a new StatementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementResponse() *StatementResponse {
	this := StatementResponse{}
	return &this
}

// NewStatementResponseWithDefaults instantiates a new StatementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementResponseWithDefaults() *StatementResponse {
	this := StatementResponse{}
	return &this
}

// GetAccountGuid returns the AccountGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatementResponse) GetAccountGuid() string {
	if o == nil || IsNil(o.AccountGuid.Get()) {
		var ret string
		return ret
	}
	return *o.AccountGuid.Get()
}

// GetAccountGuidOk returns a tuple with the AccountGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatementResponse) GetAccountGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountGuid.Get(), o.AccountGuid.IsSet()
}

// HasAccountGuid returns a boolean if a field has been set.
func (o *StatementResponse) HasAccountGuid() bool {
	if o != nil && o.AccountGuid.IsSet() {
		return true
	}

	return false
}

// SetAccountGuid gets a reference to the given NullableString and assigns it to the AccountGuid field.
func (o *StatementResponse) SetAccountGuid(v string) {
	o.AccountGuid.Set(&v)
}
// SetAccountGuidNil sets the value for AccountGuid to be an explicit nil
func (o *StatementResponse) SetAccountGuidNil() {
	o.AccountGuid.Set(nil)
}

// UnsetAccountGuid ensures that no value is present for AccountGuid, not even an explicit nil
func (o *StatementResponse) UnsetAccountGuid() {
	o.AccountGuid.Unset()
}

// GetContentHash returns the ContentHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatementResponse) GetContentHash() string {
	if o == nil || IsNil(o.ContentHash.Get()) {
		var ret string
		return ret
	}
	return *o.ContentHash.Get()
}

// GetContentHashOk returns a tuple with the ContentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatementResponse) GetContentHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentHash.Get(), o.ContentHash.IsSet()
}

// HasContentHash returns a boolean if a field has been set.
func (o *StatementResponse) HasContentHash() bool {
	if o != nil && o.ContentHash.IsSet() {
		return true
	}

	return false
}

// SetContentHash gets a reference to the given NullableString and assigns it to the ContentHash field.
func (o *StatementResponse) SetContentHash(v string) {
	o.ContentHash.Set(&v)
}
// SetContentHashNil sets the value for ContentHash to be an explicit nil
func (o *StatementResponse) SetContentHashNil() {
	o.ContentHash.Set(nil)
}

// UnsetContentHash ensures that no value is present for ContentHash, not even an explicit nil
func (o *StatementResponse) UnsetContentHash() {
	o.ContentHash.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatementResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatementResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StatementResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *StatementResponse) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *StatementResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *StatementResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatementResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatementResponse) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *StatementResponse) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *StatementResponse) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *StatementResponse) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *StatementResponse) UnsetGuid() {
	o.Guid.Unset()
}

// GetMemberGuid returns the MemberGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatementResponse) GetMemberGuid() string {
	if o == nil || IsNil(o.MemberGuid.Get()) {
		var ret string
		return ret
	}
	return *o.MemberGuid.Get()
}

// GetMemberGuidOk returns a tuple with the MemberGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatementResponse) GetMemberGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemberGuid.Get(), o.MemberGuid.IsSet()
}

// HasMemberGuid returns a boolean if a field has been set.
func (o *StatementResponse) HasMemberGuid() bool {
	if o != nil && o.MemberGuid.IsSet() {
		return true
	}

	return false
}

// SetMemberGuid gets a reference to the given NullableString and assigns it to the MemberGuid field.
func (o *StatementResponse) SetMemberGuid(v string) {
	o.MemberGuid.Set(&v)
}
// SetMemberGuidNil sets the value for MemberGuid to be an explicit nil
func (o *StatementResponse) SetMemberGuidNil() {
	o.MemberGuid.Set(nil)
}

// UnsetMemberGuid ensures that no value is present for MemberGuid, not even an explicit nil
func (o *StatementResponse) UnsetMemberGuid() {
	o.MemberGuid.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatementResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatementResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *StatementResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *StatementResponse) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *StatementResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *StatementResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatementResponse) GetUri() string {
	if o == nil || IsNil(o.Uri.Get()) {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatementResponse) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *StatementResponse) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *StatementResponse) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *StatementResponse) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *StatementResponse) UnsetUri() {
	o.Uri.Unset()
}

// GetUserGuid returns the UserGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatementResponse) GetUserGuid() string {
	if o == nil || IsNil(o.UserGuid.Get()) {
		var ret string
		return ret
	}
	return *o.UserGuid.Get()
}

// GetUserGuidOk returns a tuple with the UserGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatementResponse) GetUserGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserGuid.Get(), o.UserGuid.IsSet()
}

// HasUserGuid returns a boolean if a field has been set.
func (o *StatementResponse) HasUserGuid() bool {
	if o != nil && o.UserGuid.IsSet() {
		return true
	}

	return false
}

// SetUserGuid gets a reference to the given NullableString and assigns it to the UserGuid field.
func (o *StatementResponse) SetUserGuid(v string) {
	o.UserGuid.Set(&v)
}
// SetUserGuidNil sets the value for UserGuid to be an explicit nil
func (o *StatementResponse) SetUserGuidNil() {
	o.UserGuid.Set(nil)
}

// UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
func (o *StatementResponse) UnsetUserGuid() {
	o.UserGuid.Unset()
}

func (o StatementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountGuid.IsSet() {
		toSerialize["account_guid"] = o.AccountGuid.Get()
	}
	if o.ContentHash.IsSet() {
		toSerialize["content_hash"] = o.ContentHash.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.MemberGuid.IsSet() {
		toSerialize["member_guid"] = o.MemberGuid.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	if o.UserGuid.IsSet() {
		toSerialize["user_guid"] = o.UserGuid.Get()
	}
	return toSerialize, nil
}

type NullableStatementResponse struct {
	value *StatementResponse
	isSet bool
}

func (v NullableStatementResponse) Get() *StatementResponse {
	return v.value
}

func (v *NullableStatementResponse) Set(val *StatementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementResponse(val *StatementResponse) *NullableStatementResponse {
	return &NullableStatementResponse{value: val, isSet: true}
}

func (v NullableStatementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


