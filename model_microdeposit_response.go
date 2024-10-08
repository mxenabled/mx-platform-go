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

// checks if the MicrodepositResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MicrodepositResponse{}

// MicrodepositResponse struct for MicrodepositResponse
type MicrodepositResponse struct {
	AccountName *string `json:"account_name,omitempty"`
	AccountNumber *string `json:"account_number,omitempty"`
	AccountType *string `json:"account_type,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	RoutingNumber *string `json:"routing_number,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	Guid *string `json:"guid,omitempty"`
	InstitutionCode *string `json:"institution_code,omitempty"`
	InstitutionName *string `json:"institution_name,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	VerifiedAt *string `json:"verified_at,omitempty"`
}

// NewMicrodepositResponse instantiates a new MicrodepositResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrodepositResponse() *MicrodepositResponse {
	this := MicrodepositResponse{}
	return &this
}

// NewMicrodepositResponseWithDefaults instantiates a new MicrodepositResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrodepositResponseWithDefaults() *MicrodepositResponse {
	this := MicrodepositResponse{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *MicrodepositResponse) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *MicrodepositResponse) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *MicrodepositResponse) SetAccountType(v string) {
	o.AccountType = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MicrodepositResponse) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *MicrodepositResponse) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *MicrodepositResponse) SetLastName(v string) {
	o.LastName = &v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetRoutingNumber() string {
	if o == nil || IsNil(o.RoutingNumber) {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetRoutingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingNumber) {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasRoutingNumber() bool {
	if o != nil && !IsNil(o.RoutingNumber) {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *MicrodepositResponse) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *MicrodepositResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *MicrodepositResponse) SetGuid(v string) {
	o.Guid = &v
}

// GetInstitutionCode returns the InstitutionCode field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetInstitutionCode() string {
	if o == nil || IsNil(o.InstitutionCode) {
		var ret string
		return ret
	}
	return *o.InstitutionCode
}

// GetInstitutionCodeOk returns a tuple with the InstitutionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetInstitutionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InstitutionCode) {
		return nil, false
	}
	return o.InstitutionCode, true
}

// HasInstitutionCode returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasInstitutionCode() bool {
	if o != nil && !IsNil(o.InstitutionCode) {
		return true
	}

	return false
}

// SetInstitutionCode gets a reference to the given string and assigns it to the InstitutionCode field.
func (o *MicrodepositResponse) SetInstitutionCode(v string) {
	o.InstitutionCode = &v
}

// GetInstitutionName returns the InstitutionName field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetInstitutionName() string {
	if o == nil || IsNil(o.InstitutionName) {
		var ret string
		return ret
	}
	return *o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetInstitutionNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstitutionName) {
		return nil, false
	}
	return o.InstitutionName, true
}

// HasInstitutionName returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasInstitutionName() bool {
	if o != nil && !IsNil(o.InstitutionName) {
		return true
	}

	return false
}

// SetInstitutionName gets a reference to the given string and assigns it to the InstitutionName field.
func (o *MicrodepositResponse) SetInstitutionName(v string) {
	o.InstitutionName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MicrodepositResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *MicrodepositResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetVerifiedAt returns the VerifiedAt field value if set, zero value otherwise.
func (o *MicrodepositResponse) GetVerifiedAt() string {
	if o == nil || IsNil(o.VerifiedAt) {
		var ret string
		return ret
	}
	return *o.VerifiedAt
}

// GetVerifiedAtOk returns a tuple with the VerifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrodepositResponse) GetVerifiedAtOk() (*string, bool) {
	if o == nil || IsNil(o.VerifiedAt) {
		return nil, false
	}
	return o.VerifiedAt, true
}

// HasVerifiedAt returns a boolean if a field has been set.
func (o *MicrodepositResponse) HasVerifiedAt() bool {
	if o != nil && !IsNil(o.VerifiedAt) {
		return true
	}

	return false
}

// SetVerifiedAt gets a reference to the given string and assigns it to the VerifiedAt field.
func (o *MicrodepositResponse) SetVerifiedAt(v string) {
	o.VerifiedAt = &v
}

func (o MicrodepositResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MicrodepositResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.RoutingNumber) {
		toSerialize["routing_number"] = o.RoutingNumber
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.InstitutionCode) {
		toSerialize["institution_code"] = o.InstitutionCode
	}
	if !IsNil(o.InstitutionName) {
		toSerialize["institution_name"] = o.InstitutionName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.VerifiedAt) {
		toSerialize["verified_at"] = o.VerifiedAt
	}
	return toSerialize, nil
}

type NullableMicrodepositResponse struct {
	value *MicrodepositResponse
	isSet bool
}

func (v NullableMicrodepositResponse) Get() *MicrodepositResponse {
	return v.value
}

func (v *NullableMicrodepositResponse) Set(val *MicrodepositResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrodepositResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrodepositResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrodepositResponse(val *MicrodepositResponse) *NullableMicrodepositResponse {
	return &NullableMicrodepositResponse{value: val, isSet: true}
}

func (v NullableMicrodepositResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrodepositResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


