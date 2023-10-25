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

// checks if the SpendingPlanIterationItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpendingPlanIterationItemResponse{}

// SpendingPlanIterationItemResponse struct for SpendingPlanIterationItemResponse
type SpendingPlanIterationItemResponse struct {
	ActualAmount NullableFloat32 `json:"actual_amount,omitempty"`
	CategoryGuid NullableString `json:"category_guid,omitempty"`
	CreatedAt NullableString `json:"created_at,omitempty"`
	Guid NullableString `json:"guid,omitempty"`
	ItemType NullableString `json:"item_type,omitempty"`
	PlannedAmount NullableFloat32 `json:"planned_amount,omitempty"`
	ScheduledPaymentGuid NullableString `json:"scheduled_payment_guid,omitempty"`
	SpendingPlanIterationGuid NullableString `json:"spending_plan_iteration_guid,omitempty"`
	TopLevelCategoryGuid NullableString `json:"top_level_category_guid,omitempty"`
	TransactionGuids []*string `json:"transaction_guids,omitempty"`
	UpdatedAt NullableString `json:"updated_at,omitempty"`
	UserGuid NullableString `json:"user_guid,omitempty"`
}

// NewSpendingPlanIterationItemResponse instantiates a new SpendingPlanIterationItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendingPlanIterationItemResponse() *SpendingPlanIterationItemResponse {
	this := SpendingPlanIterationItemResponse{}
	return &this
}

// NewSpendingPlanIterationItemResponseWithDefaults instantiates a new SpendingPlanIterationItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendingPlanIterationItemResponseWithDefaults() *SpendingPlanIterationItemResponse {
	this := SpendingPlanIterationItemResponse{}
	return &this
}

// GetActualAmount returns the ActualAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationItemResponse) GetActualAmount() float32 {
	if o == nil || IsNil(o.ActualAmount.Get()) {
		var ret float32
		return ret
	}
	return *o.ActualAmount.Get()
}

// GetActualAmountOk returns a tuple with the ActualAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationItemResponse) GetActualAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActualAmount.Get(), o.ActualAmount.IsSet()
}

// HasActualAmount returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasActualAmount() bool {
	if o != nil && o.ActualAmount.IsSet() {
		return true
	}

	return false
}

// SetActualAmount gets a reference to the given NullableFloat32 and assigns it to the ActualAmount field.
func (o *SpendingPlanIterationItemResponse) SetActualAmount(v float32) {
	o.ActualAmount.Set(&v)
}
// SetActualAmountNil sets the value for ActualAmount to be an explicit nil
func (o *SpendingPlanIterationItemResponse) SetActualAmountNil() {
	o.ActualAmount.Set(nil)
}

// UnsetActualAmount ensures that no value is present for ActualAmount, not even an explicit nil
func (o *SpendingPlanIterationItemResponse) UnsetActualAmount() {
	o.ActualAmount.Unset()
}

// GetCategoryGuid returns the CategoryGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationItemResponse) GetCategoryGuid() string {
	if o == nil || IsNil(o.CategoryGuid.Get()) {
		var ret string
		return ret
	}
	return *o.CategoryGuid.Get()
}

// GetCategoryGuidOk returns a tuple with the CategoryGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationItemResponse) GetCategoryGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategoryGuid.Get(), o.CategoryGuid.IsSet()
}

// HasCategoryGuid returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasCategoryGuid() bool {
	if o != nil && o.CategoryGuid.IsSet() {
		return true
	}

	return false
}

// SetCategoryGuid gets a reference to the given NullableString and assigns it to the CategoryGuid field.
func (o *SpendingPlanIterationItemResponse) SetCategoryGuid(v string) {
	o.CategoryGuid.Set(&v)
}
// SetCategoryGuidNil sets the value for CategoryGuid to be an explicit nil
func (o *SpendingPlanIterationItemResponse) SetCategoryGuidNil() {
	o.CategoryGuid.Set(nil)
}

// UnsetCategoryGuid ensures that no value is present for CategoryGuid, not even an explicit nil
func (o *SpendingPlanIterationItemResponse) UnsetCategoryGuid() {
	o.CategoryGuid.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationItemResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationItemResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *SpendingPlanIterationItemResponse) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *SpendingPlanIterationItemResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *SpendingPlanIterationItemResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationItemResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationItemResponse) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *SpendingPlanIterationItemResponse) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *SpendingPlanIterationItemResponse) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *SpendingPlanIterationItemResponse) UnsetGuid() {
	o.Guid.Unset()
}

// GetItemType returns the ItemType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationItemResponse) GetItemType() string {
	if o == nil || IsNil(o.ItemType.Get()) {
		var ret string
		return ret
	}
	return *o.ItemType.Get()
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationItemResponse) GetItemTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemType.Get(), o.ItemType.IsSet()
}

// HasItemType returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasItemType() bool {
	if o != nil && o.ItemType.IsSet() {
		return true
	}

	return false
}

// SetItemType gets a reference to the given NullableString and assigns it to the ItemType field.
func (o *SpendingPlanIterationItemResponse) SetItemType(v string) {
	o.ItemType.Set(&v)
}
// SetItemTypeNil sets the value for ItemType to be an explicit nil
func (o *SpendingPlanIterationItemResponse) SetItemTypeNil() {
	o.ItemType.Set(nil)
}

// UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
func (o *SpendingPlanIterationItemResponse) UnsetItemType() {
	o.ItemType.Unset()
}

// GetPlannedAmount returns the PlannedAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationItemResponse) GetPlannedAmount() float32 {
	if o == nil || IsNil(o.PlannedAmount.Get()) {
		var ret float32
		return ret
	}
	return *o.PlannedAmount.Get()
}

// GetPlannedAmountOk returns a tuple with the PlannedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationItemResponse) GetPlannedAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlannedAmount.Get(), o.PlannedAmount.IsSet()
}

// HasPlannedAmount returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasPlannedAmount() bool {
	if o != nil && o.PlannedAmount.IsSet() {
		return true
	}

	return false
}

// SetPlannedAmount gets a reference to the given NullableFloat32 and assigns it to the PlannedAmount field.
func (o *SpendingPlanIterationItemResponse) SetPlannedAmount(v float32) {
	o.PlannedAmount.Set(&v)
}
// SetPlannedAmountNil sets the value for PlannedAmount to be an explicit nil
func (o *SpendingPlanIterationItemResponse) SetPlannedAmountNil() {
	o.PlannedAmount.Set(nil)
}

// UnsetPlannedAmount ensures that no value is present for PlannedAmount, not even an explicit nil
func (o *SpendingPlanIterationItemResponse) UnsetPlannedAmount() {
	o.PlannedAmount.Unset()
}

// GetScheduledPaymentGuid returns the ScheduledPaymentGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationItemResponse) GetScheduledPaymentGuid() string {
	if o == nil || IsNil(o.ScheduledPaymentGuid.Get()) {
		var ret string
		return ret
	}
	return *o.ScheduledPaymentGuid.Get()
}

// GetScheduledPaymentGuidOk returns a tuple with the ScheduledPaymentGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationItemResponse) GetScheduledPaymentGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledPaymentGuid.Get(), o.ScheduledPaymentGuid.IsSet()
}

// HasScheduledPaymentGuid returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasScheduledPaymentGuid() bool {
	if o != nil && o.ScheduledPaymentGuid.IsSet() {
		return true
	}

	return false
}

// SetScheduledPaymentGuid gets a reference to the given NullableString and assigns it to the ScheduledPaymentGuid field.
func (o *SpendingPlanIterationItemResponse) SetScheduledPaymentGuid(v string) {
	o.ScheduledPaymentGuid.Set(&v)
}
// SetScheduledPaymentGuidNil sets the value for ScheduledPaymentGuid to be an explicit nil
func (o *SpendingPlanIterationItemResponse) SetScheduledPaymentGuidNil() {
	o.ScheduledPaymentGuid.Set(nil)
}

// UnsetScheduledPaymentGuid ensures that no value is present for ScheduledPaymentGuid, not even an explicit nil
func (o *SpendingPlanIterationItemResponse) UnsetScheduledPaymentGuid() {
	o.ScheduledPaymentGuid.Unset()
}

// GetSpendingPlanIterationGuid returns the SpendingPlanIterationGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationItemResponse) GetSpendingPlanIterationGuid() string {
	if o == nil || IsNil(o.SpendingPlanIterationGuid.Get()) {
		var ret string
		return ret
	}
	return *o.SpendingPlanIterationGuid.Get()
}

// GetSpendingPlanIterationGuidOk returns a tuple with the SpendingPlanIterationGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationItemResponse) GetSpendingPlanIterationGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpendingPlanIterationGuid.Get(), o.SpendingPlanIterationGuid.IsSet()
}

// HasSpendingPlanIterationGuid returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasSpendingPlanIterationGuid() bool {
	if o != nil && o.SpendingPlanIterationGuid.IsSet() {
		return true
	}

	return false
}

// SetSpendingPlanIterationGuid gets a reference to the given NullableString and assigns it to the SpendingPlanIterationGuid field.
func (o *SpendingPlanIterationItemResponse) SetSpendingPlanIterationGuid(v string) {
	o.SpendingPlanIterationGuid.Set(&v)
}
// SetSpendingPlanIterationGuidNil sets the value for SpendingPlanIterationGuid to be an explicit nil
func (o *SpendingPlanIterationItemResponse) SetSpendingPlanIterationGuidNil() {
	o.SpendingPlanIterationGuid.Set(nil)
}

// UnsetSpendingPlanIterationGuid ensures that no value is present for SpendingPlanIterationGuid, not even an explicit nil
func (o *SpendingPlanIterationItemResponse) UnsetSpendingPlanIterationGuid() {
	o.SpendingPlanIterationGuid.Unset()
}

// GetTopLevelCategoryGuid returns the TopLevelCategoryGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationItemResponse) GetTopLevelCategoryGuid() string {
	if o == nil || IsNil(o.TopLevelCategoryGuid.Get()) {
		var ret string
		return ret
	}
	return *o.TopLevelCategoryGuid.Get()
}

// GetTopLevelCategoryGuidOk returns a tuple with the TopLevelCategoryGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationItemResponse) GetTopLevelCategoryGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopLevelCategoryGuid.Get(), o.TopLevelCategoryGuid.IsSet()
}

// HasTopLevelCategoryGuid returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasTopLevelCategoryGuid() bool {
	if o != nil && o.TopLevelCategoryGuid.IsSet() {
		return true
	}

	return false
}

// SetTopLevelCategoryGuid gets a reference to the given NullableString and assigns it to the TopLevelCategoryGuid field.
func (o *SpendingPlanIterationItemResponse) SetTopLevelCategoryGuid(v string) {
	o.TopLevelCategoryGuid.Set(&v)
}
// SetTopLevelCategoryGuidNil sets the value for TopLevelCategoryGuid to be an explicit nil
func (o *SpendingPlanIterationItemResponse) SetTopLevelCategoryGuidNil() {
	o.TopLevelCategoryGuid.Set(nil)
}

// UnsetTopLevelCategoryGuid ensures that no value is present for TopLevelCategoryGuid, not even an explicit nil
func (o *SpendingPlanIterationItemResponse) UnsetTopLevelCategoryGuid() {
	o.TopLevelCategoryGuid.Unset()
}

// GetTransactionGuids returns the TransactionGuids field value if set, zero value otherwise.
func (o *SpendingPlanIterationItemResponse) GetTransactionGuids() []*string {
	if o == nil || IsNil(o.TransactionGuids) {
		var ret []*string
		return ret
	}
	return o.TransactionGuids
}

// GetTransactionGuidsOk returns a tuple with the TransactionGuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingPlanIterationItemResponse) GetTransactionGuidsOk() ([]*string, bool) {
	if o == nil || IsNil(o.TransactionGuids) {
		return nil, false
	}
	return o.TransactionGuids, true
}

// HasTransactionGuids returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasTransactionGuids() bool {
	if o != nil && !IsNil(o.TransactionGuids) {
		return true
	}

	return false
}

// SetTransactionGuids gets a reference to the given []*string and assigns it to the TransactionGuids field.
func (o *SpendingPlanIterationItemResponse) SetTransactionGuids(v []*string) {
	o.TransactionGuids = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationItemResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationItemResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *SpendingPlanIterationItemResponse) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *SpendingPlanIterationItemResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *SpendingPlanIterationItemResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUserGuid returns the UserGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpendingPlanIterationItemResponse) GetUserGuid() string {
	if o == nil || IsNil(o.UserGuid.Get()) {
		var ret string
		return ret
	}
	return *o.UserGuid.Get()
}

// GetUserGuidOk returns a tuple with the UserGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpendingPlanIterationItemResponse) GetUserGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserGuid.Get(), o.UserGuid.IsSet()
}

// HasUserGuid returns a boolean if a field has been set.
func (o *SpendingPlanIterationItemResponse) HasUserGuid() bool {
	if o != nil && o.UserGuid.IsSet() {
		return true
	}

	return false
}

// SetUserGuid gets a reference to the given NullableString and assigns it to the UserGuid field.
func (o *SpendingPlanIterationItemResponse) SetUserGuid(v string) {
	o.UserGuid.Set(&v)
}
// SetUserGuidNil sets the value for UserGuid to be an explicit nil
func (o *SpendingPlanIterationItemResponse) SetUserGuidNil() {
	o.UserGuid.Set(nil)
}

// UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
func (o *SpendingPlanIterationItemResponse) UnsetUserGuid() {
	o.UserGuid.Unset()
}

func (o SpendingPlanIterationItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpendingPlanIterationItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActualAmount.IsSet() {
		toSerialize["actual_amount"] = o.ActualAmount.Get()
	}
	if o.CategoryGuid.IsSet() {
		toSerialize["category_guid"] = o.CategoryGuid.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.ItemType.IsSet() {
		toSerialize["item_type"] = o.ItemType.Get()
	}
	if o.PlannedAmount.IsSet() {
		toSerialize["planned_amount"] = o.PlannedAmount.Get()
	}
	if o.ScheduledPaymentGuid.IsSet() {
		toSerialize["scheduled_payment_guid"] = o.ScheduledPaymentGuid.Get()
	}
	if o.SpendingPlanIterationGuid.IsSet() {
		toSerialize["spending_plan_iteration_guid"] = o.SpendingPlanIterationGuid.Get()
	}
	if o.TopLevelCategoryGuid.IsSet() {
		toSerialize["top_level_category_guid"] = o.TopLevelCategoryGuid.Get()
	}
	if !IsNil(o.TransactionGuids) {
		toSerialize["transaction_guids"] = o.TransactionGuids
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.UserGuid.IsSet() {
		toSerialize["user_guid"] = o.UserGuid.Get()
	}
	return toSerialize, nil
}

type NullableSpendingPlanIterationItemResponse struct {
	value *SpendingPlanIterationItemResponse
	isSet bool
}

func (v NullableSpendingPlanIterationItemResponse) Get() *SpendingPlanIterationItemResponse {
	return v.value
}

func (v *NullableSpendingPlanIterationItemResponse) Set(val *SpendingPlanIterationItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendingPlanIterationItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendingPlanIterationItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendingPlanIterationItemResponse(val *SpendingPlanIterationItemResponse) *NullableSpendingPlanIterationItemResponse {
	return &NullableSpendingPlanIterationItemResponse{value: val, isSet: true}
}

func (v NullableSpendingPlanIterationItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendingPlanIterationItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


