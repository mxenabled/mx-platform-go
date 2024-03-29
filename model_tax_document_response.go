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

// checks if the TaxDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxDocumentResponse{}

// TaxDocumentResponse struct for TaxDocumentResponse
type TaxDocumentResponse struct {
	ContentHash NullableString `json:"content_hash,omitempty"`
	CreatedAt NullableString `json:"created_at,omitempty"`
	DocumentType NullableString `json:"document_type,omitempty"`
	Guid NullableString `json:"guid,omitempty"`
	IssuedOn NullableString `json:"issued_on,omitempty"`
	MemberGuid NullableString `json:"member_guid,omitempty"`
	TaxYear NullableString `json:"tax_year,omitempty"`
	UpdatedAt NullableString `json:"updated_at,omitempty"`
	Uri NullableString `json:"uri,omitempty"`
	UserGuid NullableString `json:"user_guid,omitempty"`
}

// NewTaxDocumentResponse instantiates a new TaxDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxDocumentResponse() *TaxDocumentResponse {
	this := TaxDocumentResponse{}
	return &this
}

// NewTaxDocumentResponseWithDefaults instantiates a new TaxDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxDocumentResponseWithDefaults() *TaxDocumentResponse {
	this := TaxDocumentResponse{}
	return &this
}

// GetContentHash returns the ContentHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxDocumentResponse) GetContentHash() string {
	if o == nil || IsNil(o.ContentHash.Get()) {
		var ret string
		return ret
	}
	return *o.ContentHash.Get()
}

// GetContentHashOk returns a tuple with the ContentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxDocumentResponse) GetContentHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentHash.Get(), o.ContentHash.IsSet()
}

// HasContentHash returns a boolean if a field has been set.
func (o *TaxDocumentResponse) HasContentHash() bool {
	if o != nil && o.ContentHash.IsSet() {
		return true
	}

	return false
}

// SetContentHash gets a reference to the given NullableString and assigns it to the ContentHash field.
func (o *TaxDocumentResponse) SetContentHash(v string) {
	o.ContentHash.Set(&v)
}
// SetContentHashNil sets the value for ContentHash to be an explicit nil
func (o *TaxDocumentResponse) SetContentHashNil() {
	o.ContentHash.Set(nil)
}

// UnsetContentHash ensures that no value is present for ContentHash, not even an explicit nil
func (o *TaxDocumentResponse) UnsetContentHash() {
	o.ContentHash.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxDocumentResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxDocumentResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TaxDocumentResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *TaxDocumentResponse) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *TaxDocumentResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *TaxDocumentResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxDocumentResponse) GetDocumentType() string {
	if o == nil || IsNil(o.DocumentType.Get()) {
		var ret string
		return ret
	}
	return *o.DocumentType.Get()
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxDocumentResponse) GetDocumentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentType.Get(), o.DocumentType.IsSet()
}

// HasDocumentType returns a boolean if a field has been set.
func (o *TaxDocumentResponse) HasDocumentType() bool {
	if o != nil && o.DocumentType.IsSet() {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given NullableString and assigns it to the DocumentType field.
func (o *TaxDocumentResponse) SetDocumentType(v string) {
	o.DocumentType.Set(&v)
}
// SetDocumentTypeNil sets the value for DocumentType to be an explicit nil
func (o *TaxDocumentResponse) SetDocumentTypeNil() {
	o.DocumentType.Set(nil)
}

// UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
func (o *TaxDocumentResponse) UnsetDocumentType() {
	o.DocumentType.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxDocumentResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxDocumentResponse) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *TaxDocumentResponse) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *TaxDocumentResponse) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *TaxDocumentResponse) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *TaxDocumentResponse) UnsetGuid() {
	o.Guid.Unset()
}

// GetIssuedOn returns the IssuedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxDocumentResponse) GetIssuedOn() string {
	if o == nil || IsNil(o.IssuedOn.Get()) {
		var ret string
		return ret
	}
	return *o.IssuedOn.Get()
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxDocumentResponse) GetIssuedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuedOn.Get(), o.IssuedOn.IsSet()
}

// HasIssuedOn returns a boolean if a field has been set.
func (o *TaxDocumentResponse) HasIssuedOn() bool {
	if o != nil && o.IssuedOn.IsSet() {
		return true
	}

	return false
}

// SetIssuedOn gets a reference to the given NullableString and assigns it to the IssuedOn field.
func (o *TaxDocumentResponse) SetIssuedOn(v string) {
	o.IssuedOn.Set(&v)
}
// SetIssuedOnNil sets the value for IssuedOn to be an explicit nil
func (o *TaxDocumentResponse) SetIssuedOnNil() {
	o.IssuedOn.Set(nil)
}

// UnsetIssuedOn ensures that no value is present for IssuedOn, not even an explicit nil
func (o *TaxDocumentResponse) UnsetIssuedOn() {
	o.IssuedOn.Unset()
}

// GetMemberGuid returns the MemberGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxDocumentResponse) GetMemberGuid() string {
	if o == nil || IsNil(o.MemberGuid.Get()) {
		var ret string
		return ret
	}
	return *o.MemberGuid.Get()
}

// GetMemberGuidOk returns a tuple with the MemberGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxDocumentResponse) GetMemberGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemberGuid.Get(), o.MemberGuid.IsSet()
}

// HasMemberGuid returns a boolean if a field has been set.
func (o *TaxDocumentResponse) HasMemberGuid() bool {
	if o != nil && o.MemberGuid.IsSet() {
		return true
	}

	return false
}

// SetMemberGuid gets a reference to the given NullableString and assigns it to the MemberGuid field.
func (o *TaxDocumentResponse) SetMemberGuid(v string) {
	o.MemberGuid.Set(&v)
}
// SetMemberGuidNil sets the value for MemberGuid to be an explicit nil
func (o *TaxDocumentResponse) SetMemberGuidNil() {
	o.MemberGuid.Set(nil)
}

// UnsetMemberGuid ensures that no value is present for MemberGuid, not even an explicit nil
func (o *TaxDocumentResponse) UnsetMemberGuid() {
	o.MemberGuid.Unset()
}

// GetTaxYear returns the TaxYear field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxDocumentResponse) GetTaxYear() string {
	if o == nil || IsNil(o.TaxYear.Get()) {
		var ret string
		return ret
	}
	return *o.TaxYear.Get()
}

// GetTaxYearOk returns a tuple with the TaxYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxDocumentResponse) GetTaxYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxYear.Get(), o.TaxYear.IsSet()
}

// HasTaxYear returns a boolean if a field has been set.
func (o *TaxDocumentResponse) HasTaxYear() bool {
	if o != nil && o.TaxYear.IsSet() {
		return true
	}

	return false
}

// SetTaxYear gets a reference to the given NullableString and assigns it to the TaxYear field.
func (o *TaxDocumentResponse) SetTaxYear(v string) {
	o.TaxYear.Set(&v)
}
// SetTaxYearNil sets the value for TaxYear to be an explicit nil
func (o *TaxDocumentResponse) SetTaxYearNil() {
	o.TaxYear.Set(nil)
}

// UnsetTaxYear ensures that no value is present for TaxYear, not even an explicit nil
func (o *TaxDocumentResponse) UnsetTaxYear() {
	o.TaxYear.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxDocumentResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxDocumentResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TaxDocumentResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *TaxDocumentResponse) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *TaxDocumentResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *TaxDocumentResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxDocumentResponse) GetUri() string {
	if o == nil || IsNil(o.Uri.Get()) {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxDocumentResponse) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *TaxDocumentResponse) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *TaxDocumentResponse) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *TaxDocumentResponse) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *TaxDocumentResponse) UnsetUri() {
	o.Uri.Unset()
}

// GetUserGuid returns the UserGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxDocumentResponse) GetUserGuid() string {
	if o == nil || IsNil(o.UserGuid.Get()) {
		var ret string
		return ret
	}
	return *o.UserGuid.Get()
}

// GetUserGuidOk returns a tuple with the UserGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxDocumentResponse) GetUserGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserGuid.Get(), o.UserGuid.IsSet()
}

// HasUserGuid returns a boolean if a field has been set.
func (o *TaxDocumentResponse) HasUserGuid() bool {
	if o != nil && o.UserGuid.IsSet() {
		return true
	}

	return false
}

// SetUserGuid gets a reference to the given NullableString and assigns it to the UserGuid field.
func (o *TaxDocumentResponse) SetUserGuid(v string) {
	o.UserGuid.Set(&v)
}
// SetUserGuidNil sets the value for UserGuid to be an explicit nil
func (o *TaxDocumentResponse) SetUserGuidNil() {
	o.UserGuid.Set(nil)
}

// UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
func (o *TaxDocumentResponse) UnsetUserGuid() {
	o.UserGuid.Unset()
}

func (o TaxDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentHash.IsSet() {
		toSerialize["content_hash"] = o.ContentHash.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DocumentType.IsSet() {
		toSerialize["document_type"] = o.DocumentType.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.IssuedOn.IsSet() {
		toSerialize["issued_on"] = o.IssuedOn.Get()
	}
	if o.MemberGuid.IsSet() {
		toSerialize["member_guid"] = o.MemberGuid.Get()
	}
	if o.TaxYear.IsSet() {
		toSerialize["tax_year"] = o.TaxYear.Get()
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

type NullableTaxDocumentResponse struct {
	value *TaxDocumentResponse
	isSet bool
}

func (v NullableTaxDocumentResponse) Get() *TaxDocumentResponse {
	return v.value
}

func (v *NullableTaxDocumentResponse) Set(val *TaxDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxDocumentResponse(val *TaxDocumentResponse) *NullableTaxDocumentResponse {
	return &NullableTaxDocumentResponse{value: val, isSet: true}
}

func (v NullableTaxDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


