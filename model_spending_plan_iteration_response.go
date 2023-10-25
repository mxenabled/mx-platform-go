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

// checks if the SpendingPlanIterationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpendingPlanIterationResponse{}

// SpendingPlanIterationResponse struct for SpendingPlanIterationResponse
type SpendingPlanIterationResponse struct {
	CreatedAt NullableString `json:"created_at,omitempty"`
	EndOn NullableString `json:"end_on,omitempty"`
	Guid NullableString `json:"guid,omitempty"`
	IterationNumber NullableInt32 `json:"iteration_number,omitempty"`
	SpendingPlanGuid NullableString `json:"spending_plan_guid,omitempty"`
	StartOn NullableString `json:"start_on,omitempty"`
	UpdatedAt NullableString `json:"updated_at,omitempty"`
	UserGuid NullableString `json:"user_guid,omitempty"`
}

// NewSpendingPlanIterationResponse instantiates a new SpendingPlanIterationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendingPlanIterationResponse() *SpendingPlanIterationResponse {
	this := SpendingPlanIterationResponse{}
	return &this
}

// NewSpendingPlanIterationResponseWithDefaults instantiates a new SpendingPlanIterationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendingPlanIterationResponseWithDefaults() *SpendingPlanIterationResponse {
	this := SpendingPlanIterationResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SpendingPlanIterationResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *SpendingPlanIterationResponse) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *SpendingPlanIterationResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *SpendingPlanIterationResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetEndOn returns the EndOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationResponse) GetEndOn() string {
	if o == nil || IsNil(o.EndOn.Get()) {
		var ret string
		return ret
	}
	return *o.EndOn.Get()
}

// GetEndOnOk returns a tuple with the EndOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationResponse) GetEndOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndOn.Get(), o.EndOn.IsSet()
}

// HasEndOn returns a boolean if a field has been set.
func (o *SpendingPlanIterationResponse) HasEndOn() bool {
	if o != nil && o.EndOn.IsSet() {
		return true
	}

	return false
}

// SetEndOn gets a reference to the given NullableString and assigns it to the EndOn field.
func (o *SpendingPlanIterationResponse) SetEndOn(v string) {
	o.EndOn.Set(&v)
}
// SetEndOnNil sets the value for EndOn to be an explicit nil
func (o *SpendingPlanIterationResponse) SetEndOnNil() {
	o.EndOn.Set(nil)
}

// UnsetEndOn ensures that no value is present for EndOn, not even an explicit nil
func (o *SpendingPlanIterationResponse) UnsetEndOn() {
	o.EndOn.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationResponse) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *SpendingPlanIterationResponse) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *SpendingPlanIterationResponse) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *SpendingPlanIterationResponse) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *SpendingPlanIterationResponse) UnsetGuid() {
	o.Guid.Unset()
}

// GetIterationNumber returns the IterationNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationResponse) GetIterationNumber() int32 {
	if o == nil || IsNil(o.IterationNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.IterationNumber.Get()
}

// GetIterationNumberOk returns a tuple with the IterationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationResponse) GetIterationNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IterationNumber.Get(), o.IterationNumber.IsSet()
}

// HasIterationNumber returns a boolean if a field has been set.
func (o *SpendingPlanIterationResponse) HasIterationNumber() bool {
	if o != nil && o.IterationNumber.IsSet() {
		return true
	}

	return false
}

// SetIterationNumber gets a reference to the given NullableInt32 and assigns it to the IterationNumber field.
func (o *SpendingPlanIterationResponse) SetIterationNumber(v int32) {
	o.IterationNumber.Set(&v)
}
// SetIterationNumberNil sets the value for IterationNumber to be an explicit nil
func (o *SpendingPlanIterationResponse) SetIterationNumberNil() {
	o.IterationNumber.Set(nil)
}

// UnsetIterationNumber ensures that no value is present for IterationNumber, not even an explicit nil
func (o *SpendingPlanIterationResponse) UnsetIterationNumber() {
	o.IterationNumber.Unset()
}

// GetSpendingPlanGuid returns the SpendingPlanGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationResponse) GetSpendingPlanGuid() string {
	if o == nil || IsNil(o.SpendingPlanGuid.Get()) {
		var ret string
		return ret
	}
	return *o.SpendingPlanGuid.Get()
}

// GetSpendingPlanGuidOk returns a tuple with the SpendingPlanGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationResponse) GetSpendingPlanGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpendingPlanGuid.Get(), o.SpendingPlanGuid.IsSet()
}

// HasSpendingPlanGuid returns a boolean if a field has been set.
func (o *SpendingPlanIterationResponse) HasSpendingPlanGuid() bool {
	if o != nil && o.SpendingPlanGuid.IsSet() {
		return true
	}

	return false
}

// SetSpendingPlanGuid gets a reference to the given NullableString and assigns it to the SpendingPlanGuid field.
func (o *SpendingPlanIterationResponse) SetSpendingPlanGuid(v string) {
	o.SpendingPlanGuid.Set(&v)
}
// SetSpendingPlanGuidNil sets the value for SpendingPlanGuid to be an explicit nil
func (o *SpendingPlanIterationResponse) SetSpendingPlanGuidNil() {
	o.SpendingPlanGuid.Set(nil)
}

// UnsetSpendingPlanGuid ensures that no value is present for SpendingPlanGuid, not even an explicit nil
func (o *SpendingPlanIterationResponse) UnsetSpendingPlanGuid() {
	o.SpendingPlanGuid.Unset()
}

// GetStartOn returns the StartOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationResponse) GetStartOn() string {
	if o == nil || IsNil(o.StartOn.Get()) {
		var ret string
		return ret
	}
	return *o.StartOn.Get()
}

// GetStartOnOk returns a tuple with the StartOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationResponse) GetStartOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartOn.Get(), o.StartOn.IsSet()
}

// HasStartOn returns a boolean if a field has been set.
func (o *SpendingPlanIterationResponse) HasStartOn() bool {
	if o != nil && o.StartOn.IsSet() {
		return true
	}

	return false
}

// SetStartOn gets a reference to the given NullableString and assigns it to the StartOn field.
func (o *SpendingPlanIterationResponse) SetStartOn(v string) {
	o.StartOn.Set(&v)
}
// SetStartOnNil sets the value for StartOn to be an explicit nil
func (o *SpendingPlanIterationResponse) SetStartOnNil() {
	o.StartOn.Set(nil)
}

// UnsetStartOn ensures that no value is present for StartOn, not even an explicit nil
func (o *SpendingPlanIterationResponse) UnsetStartOn() {
	o.StartOn.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SpendingPlanIterationResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *SpendingPlanIterationResponse) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *SpendingPlanIterationResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *SpendingPlanIterationResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUserGuid returns the UserGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationResponse) GetUserGuid() string {
	if o == nil || IsNil(o.UserGuid.Get()) {
		var ret string
		return ret
	}
	return *o.UserGuid.Get()
}

// GetUserGuidOk returns a tuple with the UserGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationResponse) GetUserGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserGuid.Get(), o.UserGuid.IsSet()
}

// HasUserGuid returns a boolean if a field has been set.
func (o *SpendingPlanIterationResponse) HasUserGuid() bool {
	if o != nil && o.UserGuid.IsSet() {
		return true
	}

	return false
}

// SetUserGuid gets a reference to the given NullableString and assigns it to the UserGuid field.
func (o *SpendingPlanIterationResponse) SetUserGuid(v string) {
	o.UserGuid.Set(&v)
}
// SetUserGuidNil sets the value for UserGuid to be an explicit nil
func (o *SpendingPlanIterationResponse) SetUserGuidNil() {
	o.UserGuid.Set(nil)
}

// UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
func (o *SpendingPlanIterationResponse) UnsetUserGuid() {
	o.UserGuid.Unset()
}

func (o SpendingPlanIterationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpendingPlanIterationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.EndOn.IsSet() {
		toSerialize["end_on"] = o.EndOn.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.IterationNumber.IsSet() {
		toSerialize["iteration_number"] = o.IterationNumber.Get()
	}
	if o.SpendingPlanGuid.IsSet() {
		toSerialize["spending_plan_guid"] = o.SpendingPlanGuid.Get()
	}
	if o.StartOn.IsSet() {
		toSerialize["start_on"] = o.StartOn.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.UserGuid.IsSet() {
		toSerialize["user_guid"] = o.UserGuid.Get()
	}
	return toSerialize, nil
}

type NullableSpendingPlanIterationResponse struct {
	value *SpendingPlanIterationResponse
	isSet bool
}

func (v NullableSpendingPlanIterationResponse) Get() *SpendingPlanIterationResponse {
	return v.value
}

func (v *NullableSpendingPlanIterationResponse) Set(val *SpendingPlanIterationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendingPlanIterationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendingPlanIterationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendingPlanIterationResponse(val *SpendingPlanIterationResponse) *NullableSpendingPlanIterationResponse {
	return &NullableSpendingPlanIterationResponse{value: val, isSet: true}
}

func (v NullableSpendingPlanIterationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendingPlanIterationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


