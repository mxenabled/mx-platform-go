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

// WidgetRequest struct for WidgetRequest
type WidgetRequest struct {
	ClientRedirectUrl *string `json:"client_redirect_url,omitempty"`
	ColorScheme *string `json:"color_scheme,omitempty"`
	CurrentInstitutionCode *string `json:"current_institution_code,omitempty"`
	CurrentInstitutionGuid *string `json:"current_institution_guid,omitempty"`
	CurrentMemberGuid *string `json:"current_member_guid,omitempty"`
	DisableInstitutionSearch *bool `json:"disable_institution_search,omitempty"`
	IncludeTransactions *bool `json:"include_transactions,omitempty"`
	IsMobileWebview *bool `json:"is_mobile_webview,omitempty"`
	Mode *string `json:"mode,omitempty"`
	UiMessageVersion *int32 `json:"ui_message_version,omitempty"`
	UiMessageWebviewUrlScheme *string `json:"ui_message_webview_url_scheme,omitempty"`
	UpdateCredentials *bool `json:"update_credentials,omitempty"`
	WidgetType string `json:"widget_type"`
}

// NewWidgetRequest instantiates a new WidgetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetRequest(widgetType string) *WidgetRequest {
	this := WidgetRequest{}
	this.WidgetType = widgetType
	return &this
}

// NewWidgetRequestWithDefaults instantiates a new WidgetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetRequestWithDefaults() *WidgetRequest {
	this := WidgetRequest{}
	return &this
}

// GetClientRedirectUrl returns the ClientRedirectUrl field value if set, zero value otherwise.
func (o *WidgetRequest) GetClientRedirectUrl() string {
	if o == nil || o.ClientRedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.ClientRedirectUrl
}

// GetClientRedirectUrlOk returns a tuple with the ClientRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetClientRedirectUrlOk() (*string, bool) {
	if o == nil || o.ClientRedirectUrl == nil {
		return nil, false
	}
	return o.ClientRedirectUrl, true
}

// HasClientRedirectUrl returns a boolean if a field has been set.
func (o *WidgetRequest) HasClientRedirectUrl() bool {
	if o != nil && o.ClientRedirectUrl != nil {
		return true
	}

	return false
}

// SetClientRedirectUrl gets a reference to the given string and assigns it to the ClientRedirectUrl field.
func (o *WidgetRequest) SetClientRedirectUrl(v string) {
	o.ClientRedirectUrl = &v
}

// GetColorScheme returns the ColorScheme field value if set, zero value otherwise.
func (o *WidgetRequest) GetColorScheme() string {
	if o == nil || o.ColorScheme == nil {
		var ret string
		return ret
	}
	return *o.ColorScheme
}

// GetColorSchemeOk returns a tuple with the ColorScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetColorSchemeOk() (*string, bool) {
	if o == nil || o.ColorScheme == nil {
		return nil, false
	}
	return o.ColorScheme, true
}

// HasColorScheme returns a boolean if a field has been set.
func (o *WidgetRequest) HasColorScheme() bool {
	if o != nil && o.ColorScheme != nil {
		return true
	}

	return false
}

// SetColorScheme gets a reference to the given string and assigns it to the ColorScheme field.
func (o *WidgetRequest) SetColorScheme(v string) {
	o.ColorScheme = &v
}

// GetCurrentInstitutionCode returns the CurrentInstitutionCode field value if set, zero value otherwise.
func (o *WidgetRequest) GetCurrentInstitutionCode() string {
	if o == nil || o.CurrentInstitutionCode == nil {
		var ret string
		return ret
	}
	return *o.CurrentInstitutionCode
}

// GetCurrentInstitutionCodeOk returns a tuple with the CurrentInstitutionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetCurrentInstitutionCodeOk() (*string, bool) {
	if o == nil || o.CurrentInstitutionCode == nil {
		return nil, false
	}
	return o.CurrentInstitutionCode, true
}

// HasCurrentInstitutionCode returns a boolean if a field has been set.
func (o *WidgetRequest) HasCurrentInstitutionCode() bool {
	if o != nil && o.CurrentInstitutionCode != nil {
		return true
	}

	return false
}

// SetCurrentInstitutionCode gets a reference to the given string and assigns it to the CurrentInstitutionCode field.
func (o *WidgetRequest) SetCurrentInstitutionCode(v string) {
	o.CurrentInstitutionCode = &v
}

// GetCurrentInstitutionGuid returns the CurrentInstitutionGuid field value if set, zero value otherwise.
func (o *WidgetRequest) GetCurrentInstitutionGuid() string {
	if o == nil || o.CurrentInstitutionGuid == nil {
		var ret string
		return ret
	}
	return *o.CurrentInstitutionGuid
}

// GetCurrentInstitutionGuidOk returns a tuple with the CurrentInstitutionGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetCurrentInstitutionGuidOk() (*string, bool) {
	if o == nil || o.CurrentInstitutionGuid == nil {
		return nil, false
	}
	return o.CurrentInstitutionGuid, true
}

// HasCurrentInstitutionGuid returns a boolean if a field has been set.
func (o *WidgetRequest) HasCurrentInstitutionGuid() bool {
	if o != nil && o.CurrentInstitutionGuid != nil {
		return true
	}

	return false
}

// SetCurrentInstitutionGuid gets a reference to the given string and assigns it to the CurrentInstitutionGuid field.
func (o *WidgetRequest) SetCurrentInstitutionGuid(v string) {
	o.CurrentInstitutionGuid = &v
}

// GetCurrentMemberGuid returns the CurrentMemberGuid field value if set, zero value otherwise.
func (o *WidgetRequest) GetCurrentMemberGuid() string {
	if o == nil || o.CurrentMemberGuid == nil {
		var ret string
		return ret
	}
	return *o.CurrentMemberGuid
}

// GetCurrentMemberGuidOk returns a tuple with the CurrentMemberGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetCurrentMemberGuidOk() (*string, bool) {
	if o == nil || o.CurrentMemberGuid == nil {
		return nil, false
	}
	return o.CurrentMemberGuid, true
}

// HasCurrentMemberGuid returns a boolean if a field has been set.
func (o *WidgetRequest) HasCurrentMemberGuid() bool {
	if o != nil && o.CurrentMemberGuid != nil {
		return true
	}

	return false
}

// SetCurrentMemberGuid gets a reference to the given string and assigns it to the CurrentMemberGuid field.
func (o *WidgetRequest) SetCurrentMemberGuid(v string) {
	o.CurrentMemberGuid = &v
}

// GetDisableInstitutionSearch returns the DisableInstitutionSearch field value if set, zero value otherwise.
func (o *WidgetRequest) GetDisableInstitutionSearch() bool {
	if o == nil || o.DisableInstitutionSearch == nil {
		var ret bool
		return ret
	}
	return *o.DisableInstitutionSearch
}

// GetDisableInstitutionSearchOk returns a tuple with the DisableInstitutionSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetDisableInstitutionSearchOk() (*bool, bool) {
	if o == nil || o.DisableInstitutionSearch == nil {
		return nil, false
	}
	return o.DisableInstitutionSearch, true
}

// HasDisableInstitutionSearch returns a boolean if a field has been set.
func (o *WidgetRequest) HasDisableInstitutionSearch() bool {
	if o != nil && o.DisableInstitutionSearch != nil {
		return true
	}

	return false
}

// SetDisableInstitutionSearch gets a reference to the given bool and assigns it to the DisableInstitutionSearch field.
func (o *WidgetRequest) SetDisableInstitutionSearch(v bool) {
	o.DisableInstitutionSearch = &v
}

// GetIncludeTransactions returns the IncludeTransactions field value if set, zero value otherwise.
func (o *WidgetRequest) GetIncludeTransactions() bool {
	if o == nil || o.IncludeTransactions == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTransactions
}

// GetIncludeTransactionsOk returns a tuple with the IncludeTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetIncludeTransactionsOk() (*bool, bool) {
	if o == nil || o.IncludeTransactions == nil {
		return nil, false
	}
	return o.IncludeTransactions, true
}

// HasIncludeTransactions returns a boolean if a field has been set.
func (o *WidgetRequest) HasIncludeTransactions() bool {
	if o != nil && o.IncludeTransactions != nil {
		return true
	}

	return false
}

// SetIncludeTransactions gets a reference to the given bool and assigns it to the IncludeTransactions field.
func (o *WidgetRequest) SetIncludeTransactions(v bool) {
	o.IncludeTransactions = &v
}

// GetIsMobileWebview returns the IsMobileWebview field value if set, zero value otherwise.
func (o *WidgetRequest) GetIsMobileWebview() bool {
	if o == nil || o.IsMobileWebview == nil {
		var ret bool
		return ret
	}
	return *o.IsMobileWebview
}

// GetIsMobileWebviewOk returns a tuple with the IsMobileWebview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetIsMobileWebviewOk() (*bool, bool) {
	if o == nil || o.IsMobileWebview == nil {
		return nil, false
	}
	return o.IsMobileWebview, true
}

// HasIsMobileWebview returns a boolean if a field has been set.
func (o *WidgetRequest) HasIsMobileWebview() bool {
	if o != nil && o.IsMobileWebview != nil {
		return true
	}

	return false
}

// SetIsMobileWebview gets a reference to the given bool and assigns it to the IsMobileWebview field.
func (o *WidgetRequest) SetIsMobileWebview(v bool) {
	o.IsMobileWebview = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *WidgetRequest) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *WidgetRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *WidgetRequest) SetMode(v string) {
	o.Mode = &v
}

// GetUiMessageVersion returns the UiMessageVersion field value if set, zero value otherwise.
func (o *WidgetRequest) GetUiMessageVersion() int32 {
	if o == nil || o.UiMessageVersion == nil {
		var ret int32
		return ret
	}
	return *o.UiMessageVersion
}

// GetUiMessageVersionOk returns a tuple with the UiMessageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetUiMessageVersionOk() (*int32, bool) {
	if o == nil || o.UiMessageVersion == nil {
		return nil, false
	}
	return o.UiMessageVersion, true
}

// HasUiMessageVersion returns a boolean if a field has been set.
func (o *WidgetRequest) HasUiMessageVersion() bool {
	if o != nil && o.UiMessageVersion != nil {
		return true
	}

	return false
}

// SetUiMessageVersion gets a reference to the given int32 and assigns it to the UiMessageVersion field.
func (o *WidgetRequest) SetUiMessageVersion(v int32) {
	o.UiMessageVersion = &v
}

// GetUiMessageWebviewUrlScheme returns the UiMessageWebviewUrlScheme field value if set, zero value otherwise.
func (o *WidgetRequest) GetUiMessageWebviewUrlScheme() string {
	if o == nil || o.UiMessageWebviewUrlScheme == nil {
		var ret string
		return ret
	}
	return *o.UiMessageWebviewUrlScheme
}

// GetUiMessageWebviewUrlSchemeOk returns a tuple with the UiMessageWebviewUrlScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetUiMessageWebviewUrlSchemeOk() (*string, bool) {
	if o == nil || o.UiMessageWebviewUrlScheme == nil {
		return nil, false
	}
	return o.UiMessageWebviewUrlScheme, true
}

// HasUiMessageWebviewUrlScheme returns a boolean if a field has been set.
func (o *WidgetRequest) HasUiMessageWebviewUrlScheme() bool {
	if o != nil && o.UiMessageWebviewUrlScheme != nil {
		return true
	}

	return false
}

// SetUiMessageWebviewUrlScheme gets a reference to the given string and assigns it to the UiMessageWebviewUrlScheme field.
func (o *WidgetRequest) SetUiMessageWebviewUrlScheme(v string) {
	o.UiMessageWebviewUrlScheme = &v
}

// GetUpdateCredentials returns the UpdateCredentials field value if set, zero value otherwise.
func (o *WidgetRequest) GetUpdateCredentials() bool {
	if o == nil || o.UpdateCredentials == nil {
		var ret bool
		return ret
	}
	return *o.UpdateCredentials
}

// GetUpdateCredentialsOk returns a tuple with the UpdateCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetUpdateCredentialsOk() (*bool, bool) {
	if o == nil || o.UpdateCredentials == nil {
		return nil, false
	}
	return o.UpdateCredentials, true
}

// HasUpdateCredentials returns a boolean if a field has been set.
func (o *WidgetRequest) HasUpdateCredentials() bool {
	if o != nil && o.UpdateCredentials != nil {
		return true
	}

	return false
}

// SetUpdateCredentials gets a reference to the given bool and assigns it to the UpdateCredentials field.
func (o *WidgetRequest) SetUpdateCredentials(v bool) {
	o.UpdateCredentials = &v
}

// GetWidgetType returns the WidgetType field value
func (o *WidgetRequest) GetWidgetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WidgetType
}

// GetWidgetTypeOk returns a tuple with the WidgetType field value
// and a boolean to check if the value has been set.
func (o *WidgetRequest) GetWidgetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WidgetType, true
}

// SetWidgetType sets field value
func (o *WidgetRequest) SetWidgetType(v string) {
	o.WidgetType = v
}

func (o WidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientRedirectUrl != nil {
		toSerialize["client_redirect_url"] = o.ClientRedirectUrl
	}
	if o.ColorScheme != nil {
		toSerialize["color_scheme"] = o.ColorScheme
	}
	if o.CurrentInstitutionCode != nil {
		toSerialize["current_institution_code"] = o.CurrentInstitutionCode
	}
	if o.CurrentInstitutionGuid != nil {
		toSerialize["current_institution_guid"] = o.CurrentInstitutionGuid
	}
	if o.CurrentMemberGuid != nil {
		toSerialize["current_member_guid"] = o.CurrentMemberGuid
	}
	if o.DisableInstitutionSearch != nil {
		toSerialize["disable_institution_search"] = o.DisableInstitutionSearch
	}
	if o.IncludeTransactions != nil {
		toSerialize["include_transactions"] = o.IncludeTransactions
	}
	if o.IsMobileWebview != nil {
		toSerialize["is_mobile_webview"] = o.IsMobileWebview
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.UiMessageVersion != nil {
		toSerialize["ui_message_version"] = o.UiMessageVersion
	}
	if o.UiMessageWebviewUrlScheme != nil {
		toSerialize["ui_message_webview_url_scheme"] = o.UiMessageWebviewUrlScheme
	}
	if o.UpdateCredentials != nil {
		toSerialize["update_credentials"] = o.UpdateCredentials
	}
	if true {
		toSerialize["widget_type"] = o.WidgetType
	}
	return json.Marshal(toSerialize)
}

type NullableWidgetRequest struct {
	value *WidgetRequest
	isSet bool
}

func (v NullableWidgetRequest) Get() *WidgetRequest {
	return v.value
}

func (v *NullableWidgetRequest) Set(val *WidgetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetRequest(val *WidgetRequest) *NullableWidgetRequest {
	return &NullableWidgetRequest{value: val, isSet: true}
}

func (v NullableWidgetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


