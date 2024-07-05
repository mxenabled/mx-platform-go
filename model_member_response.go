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

// checks if the MemberResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberResponse{}

// MemberResponse struct for MemberResponse
type MemberResponse struct {
	AggregatedAt NullableString `json:"aggregated_at,omitempty"`
	BackgroundAggregationIsDisabled *bool `json:"background_aggregation_is_disabled,omitempty"`
	ConnectionStatus NullableString `json:"connection_status,omitempty"`
	Guid NullableString `json:"guid,omitempty"`
	Id NullableString `json:"id,omitempty"`
	InstitutionCode NullableString `json:"institution_code,omitempty"`
	IsBeingAggregated NullableBool `json:"is_being_aggregated,omitempty"`
	IsManagedByUser NullableBool `json:"is_managed_by_user,omitempty"`
	IsManual NullableBool `json:"is_manual,omitempty"`
	IsOauth NullableBool `json:"is_oauth,omitempty"`
	Metadata NullableString `json:"metadata,omitempty"`
	MostRecentJobDetailCode NullableString `json:"most_recent_job_detail_code,omitempty"`
	MostRecentJobDetailText NullableString `json:"most_recent_job_detail_text,omitempty"`
	Name NullableString `json:"name,omitempty"`
	OauthWindowUri NullableString `json:"oauth_window_uri,omitempty"`
	SuccessfullyAggregatedAt NullableString `json:"successfully_aggregated_at,omitempty"`
	UserGuid NullableString `json:"user_guid,omitempty"`
	UserId NullableString `json:"user_id,omitempty"`
}

// NewMemberResponse instantiates a new MemberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberResponse() *MemberResponse {
	this := MemberResponse{}
	return &this
}

// NewMemberResponseWithDefaults instantiates a new MemberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberResponseWithDefaults() *MemberResponse {
	this := MemberResponse{}
	return &this
}

// GetAggregatedAt returns the AggregatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetAggregatedAt() string {
	if o == nil || IsNil(o.AggregatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.AggregatedAt.Get()
}

// GetAggregatedAtOk returns a tuple with the AggregatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetAggregatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AggregatedAt.Get(), o.AggregatedAt.IsSet()
}

// HasAggregatedAt returns a boolean if a field has been set.
func (o *MemberResponse) HasAggregatedAt() bool {
	if o != nil && o.AggregatedAt.IsSet() {
		return true
	}

	return false
}

// SetAggregatedAt gets a reference to the given NullableString and assigns it to the AggregatedAt field.
func (o *MemberResponse) SetAggregatedAt(v string) {
	o.AggregatedAt.Set(&v)
}
// SetAggregatedAtNil sets the value for AggregatedAt to be an explicit nil
func (o *MemberResponse) SetAggregatedAtNil() {
	o.AggregatedAt.Set(nil)
}

// UnsetAggregatedAt ensures that no value is present for AggregatedAt, not even an explicit nil
func (o *MemberResponse) UnsetAggregatedAt() {
	o.AggregatedAt.Unset()
}

// GetBackgroundAggregationIsDisabled returns the BackgroundAggregationIsDisabled field value if set, zero value otherwise.
func (o *MemberResponse) GetBackgroundAggregationIsDisabled() bool {
	if o == nil || IsNil(o.BackgroundAggregationIsDisabled) {
		var ret bool
		return ret
	}
	return *o.BackgroundAggregationIsDisabled
}

// GetBackgroundAggregationIsDisabledOk returns a tuple with the BackgroundAggregationIsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberResponse) GetBackgroundAggregationIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BackgroundAggregationIsDisabled) {
		return nil, false
	}
	return o.BackgroundAggregationIsDisabled, true
}

// HasBackgroundAggregationIsDisabled returns a boolean if a field has been set.
func (o *MemberResponse) HasBackgroundAggregationIsDisabled() bool {
	if o != nil && !IsNil(o.BackgroundAggregationIsDisabled) {
		return true
	}

	return false
}

// SetBackgroundAggregationIsDisabled gets a reference to the given bool and assigns it to the BackgroundAggregationIsDisabled field.
func (o *MemberResponse) SetBackgroundAggregationIsDisabled(v bool) {
	o.BackgroundAggregationIsDisabled = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetConnectionStatus() string {
	if o == nil || IsNil(o.ConnectionStatus.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectionStatus.Get()
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetConnectionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionStatus.Get(), o.ConnectionStatus.IsSet()
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *MemberResponse) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus.IsSet() {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given NullableString and assigns it to the ConnectionStatus field.
func (o *MemberResponse) SetConnectionStatus(v string) {
	o.ConnectionStatus.Set(&v)
}
// SetConnectionStatusNil sets the value for ConnectionStatus to be an explicit nil
func (o *MemberResponse) SetConnectionStatusNil() {
	o.ConnectionStatus.Set(nil)
}

// UnsetConnectionStatus ensures that no value is present for ConnectionStatus, not even an explicit nil
func (o *MemberResponse) UnsetConnectionStatus() {
	o.ConnectionStatus.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *MemberResponse) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *MemberResponse) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *MemberResponse) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *MemberResponse) UnsetGuid() {
	o.Guid.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MemberResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MemberResponse) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MemberResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MemberResponse) UnsetId() {
	o.Id.Unset()
}

// GetInstitutionCode returns the InstitutionCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetInstitutionCode() string {
	if o == nil || IsNil(o.InstitutionCode.Get()) {
		var ret string
		return ret
	}
	return *o.InstitutionCode.Get()
}

// GetInstitutionCodeOk returns a tuple with the InstitutionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetInstitutionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstitutionCode.Get(), o.InstitutionCode.IsSet()
}

// HasInstitutionCode returns a boolean if a field has been set.
func (o *MemberResponse) HasInstitutionCode() bool {
	if o != nil && o.InstitutionCode.IsSet() {
		return true
	}

	return false
}

// SetInstitutionCode gets a reference to the given NullableString and assigns it to the InstitutionCode field.
func (o *MemberResponse) SetInstitutionCode(v string) {
	o.InstitutionCode.Set(&v)
}
// SetInstitutionCodeNil sets the value for InstitutionCode to be an explicit nil
func (o *MemberResponse) SetInstitutionCodeNil() {
	o.InstitutionCode.Set(nil)
}

// UnsetInstitutionCode ensures that no value is present for InstitutionCode, not even an explicit nil
func (o *MemberResponse) UnsetInstitutionCode() {
	o.InstitutionCode.Unset()
}

// GetIsBeingAggregated returns the IsBeingAggregated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetIsBeingAggregated() bool {
	if o == nil || IsNil(o.IsBeingAggregated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsBeingAggregated.Get()
}

// GetIsBeingAggregatedOk returns a tuple with the IsBeingAggregated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetIsBeingAggregatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsBeingAggregated.Get(), o.IsBeingAggregated.IsSet()
}

// HasIsBeingAggregated returns a boolean if a field has been set.
func (o *MemberResponse) HasIsBeingAggregated() bool {
	if o != nil && o.IsBeingAggregated.IsSet() {
		return true
	}

	return false
}

// SetIsBeingAggregated gets a reference to the given NullableBool and assigns it to the IsBeingAggregated field.
func (o *MemberResponse) SetIsBeingAggregated(v bool) {
	o.IsBeingAggregated.Set(&v)
}
// SetIsBeingAggregatedNil sets the value for IsBeingAggregated to be an explicit nil
func (o *MemberResponse) SetIsBeingAggregatedNil() {
	o.IsBeingAggregated.Set(nil)
}

// UnsetIsBeingAggregated ensures that no value is present for IsBeingAggregated, not even an explicit nil
func (o *MemberResponse) UnsetIsBeingAggregated() {
	o.IsBeingAggregated.Unset()
}

// GetIsManagedByUser returns the IsManagedByUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetIsManagedByUser() bool {
	if o == nil || IsNil(o.IsManagedByUser.Get()) {
		var ret bool
		return ret
	}
	return *o.IsManagedByUser.Get()
}

// GetIsManagedByUserOk returns a tuple with the IsManagedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetIsManagedByUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsManagedByUser.Get(), o.IsManagedByUser.IsSet()
}

// HasIsManagedByUser returns a boolean if a field has been set.
func (o *MemberResponse) HasIsManagedByUser() bool {
	if o != nil && o.IsManagedByUser.IsSet() {
		return true
	}

	return false
}

// SetIsManagedByUser gets a reference to the given NullableBool and assigns it to the IsManagedByUser field.
func (o *MemberResponse) SetIsManagedByUser(v bool) {
	o.IsManagedByUser.Set(&v)
}
// SetIsManagedByUserNil sets the value for IsManagedByUser to be an explicit nil
func (o *MemberResponse) SetIsManagedByUserNil() {
	o.IsManagedByUser.Set(nil)
}

// UnsetIsManagedByUser ensures that no value is present for IsManagedByUser, not even an explicit nil
func (o *MemberResponse) UnsetIsManagedByUser() {
	o.IsManagedByUser.Unset()
}

// GetIsManual returns the IsManual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetIsManual() bool {
	if o == nil || IsNil(o.IsManual.Get()) {
		var ret bool
		return ret
	}
	return *o.IsManual.Get()
}

// GetIsManualOk returns a tuple with the IsManual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetIsManualOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsManual.Get(), o.IsManual.IsSet()
}

// HasIsManual returns a boolean if a field has been set.
func (o *MemberResponse) HasIsManual() bool {
	if o != nil && o.IsManual.IsSet() {
		return true
	}

	return false
}

// SetIsManual gets a reference to the given NullableBool and assigns it to the IsManual field.
func (o *MemberResponse) SetIsManual(v bool) {
	o.IsManual.Set(&v)
}
// SetIsManualNil sets the value for IsManual to be an explicit nil
func (o *MemberResponse) SetIsManualNil() {
	o.IsManual.Set(nil)
}

// UnsetIsManual ensures that no value is present for IsManual, not even an explicit nil
func (o *MemberResponse) UnsetIsManual() {
	o.IsManual.Unset()
}

// GetIsOauth returns the IsOauth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetIsOauth() bool {
	if o == nil || IsNil(o.IsOauth.Get()) {
		var ret bool
		return ret
	}
	return *o.IsOauth.Get()
}

// GetIsOauthOk returns a tuple with the IsOauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetIsOauthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsOauth.Get(), o.IsOauth.IsSet()
}

// HasIsOauth returns a boolean if a field has been set.
func (o *MemberResponse) HasIsOauth() bool {
	if o != nil && o.IsOauth.IsSet() {
		return true
	}

	return false
}

// SetIsOauth gets a reference to the given NullableBool and assigns it to the IsOauth field.
func (o *MemberResponse) SetIsOauth(v bool) {
	o.IsOauth.Set(&v)
}
// SetIsOauthNil sets the value for IsOauth to be an explicit nil
func (o *MemberResponse) SetIsOauthNil() {
	o.IsOauth.Set(nil)
}

// UnsetIsOauth ensures that no value is present for IsOauth, not even an explicit nil
func (o *MemberResponse) UnsetIsOauth() {
	o.IsOauth.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetMetadata() string {
	if o == nil || IsNil(o.Metadata.Get()) {
		var ret string
		return ret
	}
	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// HasMetadata returns a boolean if a field has been set.
func (o *MemberResponse) HasMetadata() bool {
	if o != nil && o.Metadata.IsSet() {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given NullableString and assigns it to the Metadata field.
func (o *MemberResponse) SetMetadata(v string) {
	o.Metadata.Set(&v)
}
// SetMetadataNil sets the value for Metadata to be an explicit nil
func (o *MemberResponse) SetMetadataNil() {
	o.Metadata.Set(nil)
}

// UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
func (o *MemberResponse) UnsetMetadata() {
	o.Metadata.Unset()
}

// GetMostRecentJobDetailCode returns the MostRecentJobDetailCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetMostRecentJobDetailCode() string {
	if o == nil || IsNil(o.MostRecentJobDetailCode.Get()) {
		var ret string
		return ret
	}
	return *o.MostRecentJobDetailCode.Get()
}

// GetMostRecentJobDetailCodeOk returns a tuple with the MostRecentJobDetailCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetMostRecentJobDetailCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MostRecentJobDetailCode.Get(), o.MostRecentJobDetailCode.IsSet()
}

// HasMostRecentJobDetailCode returns a boolean if a field has been set.
func (o *MemberResponse) HasMostRecentJobDetailCode() bool {
	if o != nil && o.MostRecentJobDetailCode.IsSet() {
		return true
	}

	return false
}

// SetMostRecentJobDetailCode gets a reference to the given NullableString and assigns it to the MostRecentJobDetailCode field.
func (o *MemberResponse) SetMostRecentJobDetailCode(v string) {
	o.MostRecentJobDetailCode.Set(&v)
}
// SetMostRecentJobDetailCodeNil sets the value for MostRecentJobDetailCode to be an explicit nil
func (o *MemberResponse) SetMostRecentJobDetailCodeNil() {
	o.MostRecentJobDetailCode.Set(nil)
}

// UnsetMostRecentJobDetailCode ensures that no value is present for MostRecentJobDetailCode, not even an explicit nil
func (o *MemberResponse) UnsetMostRecentJobDetailCode() {
	o.MostRecentJobDetailCode.Unset()
}

// GetMostRecentJobDetailText returns the MostRecentJobDetailText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetMostRecentJobDetailText() string {
	if o == nil || IsNil(o.MostRecentJobDetailText.Get()) {
		var ret string
		return ret
	}
	return *o.MostRecentJobDetailText.Get()
}

// GetMostRecentJobDetailTextOk returns a tuple with the MostRecentJobDetailText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetMostRecentJobDetailTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MostRecentJobDetailText.Get(), o.MostRecentJobDetailText.IsSet()
}

// HasMostRecentJobDetailText returns a boolean if a field has been set.
func (o *MemberResponse) HasMostRecentJobDetailText() bool {
	if o != nil && o.MostRecentJobDetailText.IsSet() {
		return true
	}

	return false
}

// SetMostRecentJobDetailText gets a reference to the given NullableString and assigns it to the MostRecentJobDetailText field.
func (o *MemberResponse) SetMostRecentJobDetailText(v string) {
	o.MostRecentJobDetailText.Set(&v)
}
// SetMostRecentJobDetailTextNil sets the value for MostRecentJobDetailText to be an explicit nil
func (o *MemberResponse) SetMostRecentJobDetailTextNil() {
	o.MostRecentJobDetailText.Set(nil)
}

// UnsetMostRecentJobDetailText ensures that no value is present for MostRecentJobDetailText, not even an explicit nil
func (o *MemberResponse) UnsetMostRecentJobDetailText() {
	o.MostRecentJobDetailText.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MemberResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MemberResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MemberResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MemberResponse) UnsetName() {
	o.Name.Unset()
}

// GetOauthWindowUri returns the OauthWindowUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetOauthWindowUri() string {
	if o == nil || IsNil(o.OauthWindowUri.Get()) {
		var ret string
		return ret
	}
	return *o.OauthWindowUri.Get()
}

// GetOauthWindowUriOk returns a tuple with the OauthWindowUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetOauthWindowUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OauthWindowUri.Get(), o.OauthWindowUri.IsSet()
}

// HasOauthWindowUri returns a boolean if a field has been set.
func (o *MemberResponse) HasOauthWindowUri() bool {
	if o != nil && o.OauthWindowUri.IsSet() {
		return true
	}

	return false
}

// SetOauthWindowUri gets a reference to the given NullableString and assigns it to the OauthWindowUri field.
func (o *MemberResponse) SetOauthWindowUri(v string) {
	o.OauthWindowUri.Set(&v)
}
// SetOauthWindowUriNil sets the value for OauthWindowUri to be an explicit nil
func (o *MemberResponse) SetOauthWindowUriNil() {
	o.OauthWindowUri.Set(nil)
}

// UnsetOauthWindowUri ensures that no value is present for OauthWindowUri, not even an explicit nil
func (o *MemberResponse) UnsetOauthWindowUri() {
	o.OauthWindowUri.Unset()
}

// GetSuccessfullyAggregatedAt returns the SuccessfullyAggregatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetSuccessfullyAggregatedAt() string {
	if o == nil || IsNil(o.SuccessfullyAggregatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.SuccessfullyAggregatedAt.Get()
}

// GetSuccessfullyAggregatedAtOk returns a tuple with the SuccessfullyAggregatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetSuccessfullyAggregatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuccessfullyAggregatedAt.Get(), o.SuccessfullyAggregatedAt.IsSet()
}

// HasSuccessfullyAggregatedAt returns a boolean if a field has been set.
func (o *MemberResponse) HasSuccessfullyAggregatedAt() bool {
	if o != nil && o.SuccessfullyAggregatedAt.IsSet() {
		return true
	}

	return false
}

// SetSuccessfullyAggregatedAt gets a reference to the given NullableString and assigns it to the SuccessfullyAggregatedAt field.
func (o *MemberResponse) SetSuccessfullyAggregatedAt(v string) {
	o.SuccessfullyAggregatedAt.Set(&v)
}
// SetSuccessfullyAggregatedAtNil sets the value for SuccessfullyAggregatedAt to be an explicit nil
func (o *MemberResponse) SetSuccessfullyAggregatedAtNil() {
	o.SuccessfullyAggregatedAt.Set(nil)
}

// UnsetSuccessfullyAggregatedAt ensures that no value is present for SuccessfullyAggregatedAt, not even an explicit nil
func (o *MemberResponse) UnsetSuccessfullyAggregatedAt() {
	o.SuccessfullyAggregatedAt.Unset()
}

// GetUserGuid returns the UserGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetUserGuid() string {
	if o == nil || IsNil(o.UserGuid.Get()) {
		var ret string
		return ret
	}
	return *o.UserGuid.Get()
}

// GetUserGuidOk returns a tuple with the UserGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetUserGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserGuid.Get(), o.UserGuid.IsSet()
}

// HasUserGuid returns a boolean if a field has been set.
func (o *MemberResponse) HasUserGuid() bool {
	if o != nil && o.UserGuid.IsSet() {
		return true
	}

	return false
}

// SetUserGuid gets a reference to the given NullableString and assigns it to the UserGuid field.
func (o *MemberResponse) SetUserGuid(v string) {
	o.UserGuid.Set(&v)
}
// SetUserGuidNil sets the value for UserGuid to be an explicit nil
func (o *MemberResponse) SetUserGuidNil() {
	o.UserGuid.Set(nil)
}

// UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
func (o *MemberResponse) UnsetUserGuid() {
	o.UserGuid.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberResponse) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *MemberResponse) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *MemberResponse) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *MemberResponse) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *MemberResponse) UnsetUserId() {
	o.UserId.Unset()
}

func (o MemberResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AggregatedAt.IsSet() {
		toSerialize["aggregated_at"] = o.AggregatedAt.Get()
	}
	if !IsNil(o.BackgroundAggregationIsDisabled) {
		toSerialize["background_aggregation_is_disabled"] = o.BackgroundAggregationIsDisabled
	}
	if o.ConnectionStatus.IsSet() {
		toSerialize["connection_status"] = o.ConnectionStatus.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.InstitutionCode.IsSet() {
		toSerialize["institution_code"] = o.InstitutionCode.Get()
	}
	if o.IsBeingAggregated.IsSet() {
		toSerialize["is_being_aggregated"] = o.IsBeingAggregated.Get()
	}
	if o.IsManagedByUser.IsSet() {
		toSerialize["is_managed_by_user"] = o.IsManagedByUser.Get()
	}
	if o.IsManual.IsSet() {
		toSerialize["is_manual"] = o.IsManual.Get()
	}
	if o.IsOauth.IsSet() {
		toSerialize["is_oauth"] = o.IsOauth.Get()
	}
	if o.Metadata.IsSet() {
		toSerialize["metadata"] = o.Metadata.Get()
	}
	if o.MostRecentJobDetailCode.IsSet() {
		toSerialize["most_recent_job_detail_code"] = o.MostRecentJobDetailCode.Get()
	}
	if o.MostRecentJobDetailText.IsSet() {
		toSerialize["most_recent_job_detail_text"] = o.MostRecentJobDetailText.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.OauthWindowUri.IsSet() {
		toSerialize["oauth_window_uri"] = o.OauthWindowUri.Get()
	}
	if o.SuccessfullyAggregatedAt.IsSet() {
		toSerialize["successfully_aggregated_at"] = o.SuccessfullyAggregatedAt.Get()
	}
	if o.UserGuid.IsSet() {
		toSerialize["user_guid"] = o.UserGuid.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}
	return toSerialize, nil
}

type NullableMemberResponse struct {
	value *MemberResponse
	isSet bool
}

func (v NullableMemberResponse) Get() *MemberResponse {
	return v.value
}

func (v *NullableMemberResponse) Set(val *MemberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberResponse(val *MemberResponse) *NullableMemberResponse {
	return &NullableMemberResponse{value: val, isSet: true}
}

func (v NullableMemberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


