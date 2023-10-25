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

// checks if the SpendingPlanResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpendingPlanResponse{}

// SpendingPlanResponse struct for SpendingPlanResponse
type SpendingPlanResponse struct {
	CreatedAt NullableString `json:"created_at,omitempty"`
	CurrentIterationNumber NullableInt32 `json:"current_iteration_number,omitempty"`
	Guid NullableString `json:"guid,omitempty"`
	UpdatedAt NullableString `json:"updated_at,omitempty"`
	UserGuid NullableString `json:"user_guid,omitempty"`
}

// NewSpendingPlanResponse instantiates a new SpendingPlanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendingPlanResponse() *SpendingPlanResponse {
	this := SpendingPlanResponse{}
	return &this
}

// NewSpendingPlanResponseWithDefaults instantiates a new SpendingPlanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendingPlanResponseWithDefaults() *SpendingPlanResponse {
	this := SpendingPlanResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SpendingPlanResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *SpendingPlanResponse) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *SpendingPlanResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *SpendingPlanResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCurrentIterationNumber returns the CurrentIterationNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanResponse) GetCurrentIterationNumber() int32 {
	if o == nil || IsNil(o.CurrentIterationNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.CurrentIterationNumber.Get()
}

// GetCurrentIterationNumberOk returns a tuple with the CurrentIterationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanResponse) GetCurrentIterationNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentIterationNumber.Get(), o.CurrentIterationNumber.IsSet()
}

// HasCurrentIterationNumber returns a boolean if a field has been set.
func (o *SpendingPlanResponse) HasCurrentIterationNumber() bool {
	if o != nil && o.CurrentIterationNumber.IsSet() {
		return true
	}

	return false
}

// SetCurrentIterationNumber gets a reference to the given NullableInt32 and assigns it to the CurrentIterationNumber field.
func (o *SpendingPlanResponse) SetCurrentIterationNumber(v int32) {
	o.CurrentIterationNumber.Set(&v)
}
// SetCurrentIterationNumberNil sets the value for CurrentIterationNumber to be an explicit nil
func (o *SpendingPlanResponse) SetCurrentIterationNumberNil() {
	o.CurrentIterationNumber.Set(nil)
}

// UnsetCurrentIterationNumber ensures that no value is present for CurrentIterationNumber, not even an explicit nil
func (o *SpendingPlanResponse) UnsetCurrentIterationNumber() {
	o.CurrentIterationNumber.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanResponse) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *SpendingPlanResponse) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *SpendingPlanResponse) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *SpendingPlanResponse) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *SpendingPlanResponse) UnsetGuid() {
	o.Guid.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SpendingPlanResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *SpendingPlanResponse) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *SpendingPlanResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *SpendingPlanResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUserGuid returns the UserGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanResponse) GetUserGuid() string {
	if o == nil || IsNil(o.UserGuid.Get()) {
		var ret string
		return ret
	}
	return *o.UserGuid.Get()
}

// GetUserGuidOk returns a tuple with the UserGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanResponse) GetUserGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserGuid.Get(), o.UserGuid.IsSet()
}

// HasUserGuid returns a boolean if a field has been set.
func (o *SpendingPlanResponse) HasUserGuid() bool {
	if o != nil && o.UserGuid.IsSet() {
		return true
	}

	return false
}

// SetUserGuid gets a reference to the given NullableString and assigns it to the UserGuid field.
func (o *SpendingPlanResponse) SetUserGuid(v string) {
	o.UserGuid.Set(&v)
}
// SetUserGuidNil sets the value for UserGuid to be an explicit nil
func (o *SpendingPlanResponse) SetUserGuidNil() {
	o.UserGuid.Set(nil)
}

// UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
func (o *SpendingPlanResponse) UnsetUserGuid() {
	o.UserGuid.Unset()
}

func (o SpendingPlanResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpendingPlanResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.CurrentIterationNumber.IsSet() {
		toSerialize["current_iteration_number"] = o.CurrentIterationNumber.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.UserGuid.IsSet() {
		toSerialize["user_guid"] = o.UserGuid.Get()
	}
	return toSerialize, nil
}

type NullableSpendingPlanResponse struct {
	value *SpendingPlanResponse
	isSet bool
}

func (v NullableSpendingPlanResponse) Get() *SpendingPlanResponse {
	return v.value
}

func (v *NullableSpendingPlanResponse) Set(val *SpendingPlanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendingPlanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendingPlanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendingPlanResponse(val *SpendingPlanResponse) *NullableSpendingPlanResponse {
	return &NullableSpendingPlanResponse{value: val, isSet: true}
}

func (v NullableSpendingPlanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendingPlanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


