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

// checks if the InsightResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsightResponse{}

// InsightResponse struct for InsightResponse
type InsightResponse struct {
	ActiveAt NullableString `json:"active_at,omitempty"`
	ClientGuid NullableString `json:"client_guid,omitempty"`
	CreatedAt NullableString `json:"created_at,omitempty"`
	CtaClickedAt NullableString `json:"cta_clicked_at,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Guid NullableString `json:"guid,omitempty"`
	HasAssociatedAccounts NullableBool `json:"has_associated_accounts,omitempty"`
	HasAssociatedMerchants NullableBool `json:"has_associated_merchants,omitempty"`
	HasAssociatedScheduledPayments NullableBool `json:"has_associated_scheduled_payments,omitempty"`
	HasAssociatedTransactions NullableBool `json:"has_associated_transactions,omitempty"`
	HasBeenDisplayed NullableBool `json:"has_been_displayed,omitempty"`
	IsDismissed NullableBool `json:"is_dismissed,omitempty"`
	MicroCallToAction NullableString `json:"micro_call_to_action,omitempty"`
	MicroDescription NullableString `json:"micro_description,omitempty"`
	MicroTitle NullableString `json:"micro_title,omitempty"`
	Template NullableString `json:"template,omitempty"`
	Title NullableString `json:"title,omitempty"`
	UpdatedAt NullableString `json:"updated_at,omitempty"`
	UserGuid *string `json:"user_guid,omitempty"`
	UserId interface{} `json:"user_id,omitempty"`
}

// NewInsightResponse instantiates a new InsightResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsightResponse() *InsightResponse {
	this := InsightResponse{}
	return &this
}

// NewInsightResponseWithDefaults instantiates a new InsightResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightResponseWithDefaults() *InsightResponse {
	this := InsightResponse{}
	return &this
}

// GetActiveAt returns the ActiveAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetActiveAt() string {
	if o == nil || IsNil(o.ActiveAt.Get()) {
		var ret string
		return ret
	}
	return *o.ActiveAt.Get()
}

// GetActiveAtOk returns a tuple with the ActiveAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetActiveAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveAt.Get(), o.ActiveAt.IsSet()
}

// HasActiveAt returns a boolean if a field has been set.
func (o *InsightResponse) HasActiveAt() bool {
	if o != nil && o.ActiveAt.IsSet() {
		return true
	}

	return false
}

// SetActiveAt gets a reference to the given NullableString and assigns it to the ActiveAt field.
func (o *InsightResponse) SetActiveAt(v string) {
	o.ActiveAt.Set(&v)
}
// SetActiveAtNil sets the value for ActiveAt to be an explicit nil
func (o *InsightResponse) SetActiveAtNil() {
	o.ActiveAt.Set(nil)
}

// UnsetActiveAt ensures that no value is present for ActiveAt, not even an explicit nil
func (o *InsightResponse) UnsetActiveAt() {
	o.ActiveAt.Unset()
}

// GetClientGuid returns the ClientGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetClientGuid() string {
	if o == nil || IsNil(o.ClientGuid.Get()) {
		var ret string
		return ret
	}
	return *o.ClientGuid.Get()
}

// GetClientGuidOk returns a tuple with the ClientGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetClientGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientGuid.Get(), o.ClientGuid.IsSet()
}

// HasClientGuid returns a boolean if a field has been set.
func (o *InsightResponse) HasClientGuid() bool {
	if o != nil && o.ClientGuid.IsSet() {
		return true
	}

	return false
}

// SetClientGuid gets a reference to the given NullableString and assigns it to the ClientGuid field.
func (o *InsightResponse) SetClientGuid(v string) {
	o.ClientGuid.Set(&v)
}
// SetClientGuidNil sets the value for ClientGuid to be an explicit nil
func (o *InsightResponse) SetClientGuidNil() {
	o.ClientGuid.Set(nil)
}

// UnsetClientGuid ensures that no value is present for ClientGuid, not even an explicit nil
func (o *InsightResponse) UnsetClientGuid() {
	o.ClientGuid.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InsightResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *InsightResponse) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *InsightResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *InsightResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCtaClickedAt returns the CtaClickedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetCtaClickedAt() string {
	if o == nil || IsNil(o.CtaClickedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CtaClickedAt.Get()
}

// GetCtaClickedAtOk returns a tuple with the CtaClickedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetCtaClickedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CtaClickedAt.Get(), o.CtaClickedAt.IsSet()
}

// HasCtaClickedAt returns a boolean if a field has been set.
func (o *InsightResponse) HasCtaClickedAt() bool {
	if o != nil && o.CtaClickedAt.IsSet() {
		return true
	}

	return false
}

// SetCtaClickedAt gets a reference to the given NullableString and assigns it to the CtaClickedAt field.
func (o *InsightResponse) SetCtaClickedAt(v string) {
	o.CtaClickedAt.Set(&v)
}
// SetCtaClickedAtNil sets the value for CtaClickedAt to be an explicit nil
func (o *InsightResponse) SetCtaClickedAtNil() {
	o.CtaClickedAt.Set(nil)
}

// UnsetCtaClickedAt ensures that no value is present for CtaClickedAt, not even an explicit nil
func (o *InsightResponse) UnsetCtaClickedAt() {
	o.CtaClickedAt.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *InsightResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *InsightResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *InsightResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *InsightResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *InsightResponse) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *InsightResponse) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *InsightResponse) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *InsightResponse) UnsetGuid() {
	o.Guid.Unset()
}

// GetHasAssociatedAccounts returns the HasAssociatedAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetHasAssociatedAccounts() bool {
	if o == nil || IsNil(o.HasAssociatedAccounts.Get()) {
		var ret bool
		return ret
	}
	return *o.HasAssociatedAccounts.Get()
}

// GetHasAssociatedAccountsOk returns a tuple with the HasAssociatedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetHasAssociatedAccountsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasAssociatedAccounts.Get(), o.HasAssociatedAccounts.IsSet()
}

// HasHasAssociatedAccounts returns a boolean if a field has been set.
func (o *InsightResponse) HasHasAssociatedAccounts() bool {
	if o != nil && o.HasAssociatedAccounts.IsSet() {
		return true
	}

	return false
}

// SetHasAssociatedAccounts gets a reference to the given NullableBool and assigns it to the HasAssociatedAccounts field.
func (o *InsightResponse) SetHasAssociatedAccounts(v bool) {
	o.HasAssociatedAccounts.Set(&v)
}
// SetHasAssociatedAccountsNil sets the value for HasAssociatedAccounts to be an explicit nil
func (o *InsightResponse) SetHasAssociatedAccountsNil() {
	o.HasAssociatedAccounts.Set(nil)
}

// UnsetHasAssociatedAccounts ensures that no value is present for HasAssociatedAccounts, not even an explicit nil
func (o *InsightResponse) UnsetHasAssociatedAccounts() {
	o.HasAssociatedAccounts.Unset()
}

// GetHasAssociatedMerchants returns the HasAssociatedMerchants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetHasAssociatedMerchants() bool {
	if o == nil || IsNil(o.HasAssociatedMerchants.Get()) {
		var ret bool
		return ret
	}
	return *o.HasAssociatedMerchants.Get()
}

// GetHasAssociatedMerchantsOk returns a tuple with the HasAssociatedMerchants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetHasAssociatedMerchantsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasAssociatedMerchants.Get(), o.HasAssociatedMerchants.IsSet()
}

// HasHasAssociatedMerchants returns a boolean if a field has been set.
func (o *InsightResponse) HasHasAssociatedMerchants() bool {
	if o != nil && o.HasAssociatedMerchants.IsSet() {
		return true
	}

	return false
}

// SetHasAssociatedMerchants gets a reference to the given NullableBool and assigns it to the HasAssociatedMerchants field.
func (o *InsightResponse) SetHasAssociatedMerchants(v bool) {
	o.HasAssociatedMerchants.Set(&v)
}
// SetHasAssociatedMerchantsNil sets the value for HasAssociatedMerchants to be an explicit nil
func (o *InsightResponse) SetHasAssociatedMerchantsNil() {
	o.HasAssociatedMerchants.Set(nil)
}

// UnsetHasAssociatedMerchants ensures that no value is present for HasAssociatedMerchants, not even an explicit nil
func (o *InsightResponse) UnsetHasAssociatedMerchants() {
	o.HasAssociatedMerchants.Unset()
}

// GetHasAssociatedScheduledPayments returns the HasAssociatedScheduledPayments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetHasAssociatedScheduledPayments() bool {
	if o == nil || IsNil(o.HasAssociatedScheduledPayments.Get()) {
		var ret bool
		return ret
	}
	return *o.HasAssociatedScheduledPayments.Get()
}

// GetHasAssociatedScheduledPaymentsOk returns a tuple with the HasAssociatedScheduledPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetHasAssociatedScheduledPaymentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasAssociatedScheduledPayments.Get(), o.HasAssociatedScheduledPayments.IsSet()
}

// HasHasAssociatedScheduledPayments returns a boolean if a field has been set.
func (o *InsightResponse) HasHasAssociatedScheduledPayments() bool {
	if o != nil && o.HasAssociatedScheduledPayments.IsSet() {
		return true
	}

	return false
}

// SetHasAssociatedScheduledPayments gets a reference to the given NullableBool and assigns it to the HasAssociatedScheduledPayments field.
func (o *InsightResponse) SetHasAssociatedScheduledPayments(v bool) {
	o.HasAssociatedScheduledPayments.Set(&v)
}
// SetHasAssociatedScheduledPaymentsNil sets the value for HasAssociatedScheduledPayments to be an explicit nil
func (o *InsightResponse) SetHasAssociatedScheduledPaymentsNil() {
	o.HasAssociatedScheduledPayments.Set(nil)
}

// UnsetHasAssociatedScheduledPayments ensures that no value is present for HasAssociatedScheduledPayments, not even an explicit nil
func (o *InsightResponse) UnsetHasAssociatedScheduledPayments() {
	o.HasAssociatedScheduledPayments.Unset()
}

// GetHasAssociatedTransactions returns the HasAssociatedTransactions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetHasAssociatedTransactions() bool {
	if o == nil || IsNil(o.HasAssociatedTransactions.Get()) {
		var ret bool
		return ret
	}
	return *o.HasAssociatedTransactions.Get()
}

// GetHasAssociatedTransactionsOk returns a tuple with the HasAssociatedTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetHasAssociatedTransactionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasAssociatedTransactions.Get(), o.HasAssociatedTransactions.IsSet()
}

// HasHasAssociatedTransactions returns a boolean if a field has been set.
func (o *InsightResponse) HasHasAssociatedTransactions() bool {
	if o != nil && o.HasAssociatedTransactions.IsSet() {
		return true
	}

	return false
}

// SetHasAssociatedTransactions gets a reference to the given NullableBool and assigns it to the HasAssociatedTransactions field.
func (o *InsightResponse) SetHasAssociatedTransactions(v bool) {
	o.HasAssociatedTransactions.Set(&v)
}
// SetHasAssociatedTransactionsNil sets the value for HasAssociatedTransactions to be an explicit nil
func (o *InsightResponse) SetHasAssociatedTransactionsNil() {
	o.HasAssociatedTransactions.Set(nil)
}

// UnsetHasAssociatedTransactions ensures that no value is present for HasAssociatedTransactions, not even an explicit nil
func (o *InsightResponse) UnsetHasAssociatedTransactions() {
	o.HasAssociatedTransactions.Unset()
}

// GetHasBeenDisplayed returns the HasBeenDisplayed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetHasBeenDisplayed() bool {
	if o == nil || IsNil(o.HasBeenDisplayed.Get()) {
		var ret bool
		return ret
	}
	return *o.HasBeenDisplayed.Get()
}

// GetHasBeenDisplayedOk returns a tuple with the HasBeenDisplayed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetHasBeenDisplayedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasBeenDisplayed.Get(), o.HasBeenDisplayed.IsSet()
}

// HasHasBeenDisplayed returns a boolean if a field has been set.
func (o *InsightResponse) HasHasBeenDisplayed() bool {
	if o != nil && o.HasBeenDisplayed.IsSet() {
		return true
	}

	return false
}

// SetHasBeenDisplayed gets a reference to the given NullableBool and assigns it to the HasBeenDisplayed field.
func (o *InsightResponse) SetHasBeenDisplayed(v bool) {
	o.HasBeenDisplayed.Set(&v)
}
// SetHasBeenDisplayedNil sets the value for HasBeenDisplayed to be an explicit nil
func (o *InsightResponse) SetHasBeenDisplayedNil() {
	o.HasBeenDisplayed.Set(nil)
}

// UnsetHasBeenDisplayed ensures that no value is present for HasBeenDisplayed, not even an explicit nil
func (o *InsightResponse) UnsetHasBeenDisplayed() {
	o.HasBeenDisplayed.Unset()
}

// GetIsDismissed returns the IsDismissed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetIsDismissed() bool {
	if o == nil || IsNil(o.IsDismissed.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDismissed.Get()
}

// GetIsDismissedOk returns a tuple with the IsDismissed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetIsDismissedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDismissed.Get(), o.IsDismissed.IsSet()
}

// HasIsDismissed returns a boolean if a field has been set.
func (o *InsightResponse) HasIsDismissed() bool {
	if o != nil && o.IsDismissed.IsSet() {
		return true
	}

	return false
}

// SetIsDismissed gets a reference to the given NullableBool and assigns it to the IsDismissed field.
func (o *InsightResponse) SetIsDismissed(v bool) {
	o.IsDismissed.Set(&v)
}
// SetIsDismissedNil sets the value for IsDismissed to be an explicit nil
func (o *InsightResponse) SetIsDismissedNil() {
	o.IsDismissed.Set(nil)
}

// UnsetIsDismissed ensures that no value is present for IsDismissed, not even an explicit nil
func (o *InsightResponse) UnsetIsDismissed() {
	o.IsDismissed.Unset()
}

// GetMicroCallToAction returns the MicroCallToAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetMicroCallToAction() string {
	if o == nil || IsNil(o.MicroCallToAction.Get()) {
		var ret string
		return ret
	}
	return *o.MicroCallToAction.Get()
}

// GetMicroCallToActionOk returns a tuple with the MicroCallToAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetMicroCallToActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MicroCallToAction.Get(), o.MicroCallToAction.IsSet()
}

// HasMicroCallToAction returns a boolean if a field has been set.
func (o *InsightResponse) HasMicroCallToAction() bool {
	if o != nil && o.MicroCallToAction.IsSet() {
		return true
	}

	return false
}

// SetMicroCallToAction gets a reference to the given NullableString and assigns it to the MicroCallToAction field.
func (o *InsightResponse) SetMicroCallToAction(v string) {
	o.MicroCallToAction.Set(&v)
}
// SetMicroCallToActionNil sets the value for MicroCallToAction to be an explicit nil
func (o *InsightResponse) SetMicroCallToActionNil() {
	o.MicroCallToAction.Set(nil)
}

// UnsetMicroCallToAction ensures that no value is present for MicroCallToAction, not even an explicit nil
func (o *InsightResponse) UnsetMicroCallToAction() {
	o.MicroCallToAction.Unset()
}

// GetMicroDescription returns the MicroDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetMicroDescription() string {
	if o == nil || IsNil(o.MicroDescription.Get()) {
		var ret string
		return ret
	}
	return *o.MicroDescription.Get()
}

// GetMicroDescriptionOk returns a tuple with the MicroDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetMicroDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MicroDescription.Get(), o.MicroDescription.IsSet()
}

// HasMicroDescription returns a boolean if a field has been set.
func (o *InsightResponse) HasMicroDescription() bool {
	if o != nil && o.MicroDescription.IsSet() {
		return true
	}

	return false
}

// SetMicroDescription gets a reference to the given NullableString and assigns it to the MicroDescription field.
func (o *InsightResponse) SetMicroDescription(v string) {
	o.MicroDescription.Set(&v)
}
// SetMicroDescriptionNil sets the value for MicroDescription to be an explicit nil
func (o *InsightResponse) SetMicroDescriptionNil() {
	o.MicroDescription.Set(nil)
}

// UnsetMicroDescription ensures that no value is present for MicroDescription, not even an explicit nil
func (o *InsightResponse) UnsetMicroDescription() {
	o.MicroDescription.Unset()
}

// GetMicroTitle returns the MicroTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetMicroTitle() string {
	if o == nil || IsNil(o.MicroTitle.Get()) {
		var ret string
		return ret
	}
	return *o.MicroTitle.Get()
}

// GetMicroTitleOk returns a tuple with the MicroTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetMicroTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MicroTitle.Get(), o.MicroTitle.IsSet()
}

// HasMicroTitle returns a boolean if a field has been set.
func (o *InsightResponse) HasMicroTitle() bool {
	if o != nil && o.MicroTitle.IsSet() {
		return true
	}

	return false
}

// SetMicroTitle gets a reference to the given NullableString and assigns it to the MicroTitle field.
func (o *InsightResponse) SetMicroTitle(v string) {
	o.MicroTitle.Set(&v)
}
// SetMicroTitleNil sets the value for MicroTitle to be an explicit nil
func (o *InsightResponse) SetMicroTitleNil() {
	o.MicroTitle.Set(nil)
}

// UnsetMicroTitle ensures that no value is present for MicroTitle, not even an explicit nil
func (o *InsightResponse) UnsetMicroTitle() {
	o.MicroTitle.Unset()
}

// GetTemplate returns the Template field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetTemplate() string {
	if o == nil || IsNil(o.Template.Get()) {
		var ret string
		return ret
	}
	return *o.Template.Get()
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Template.Get(), o.Template.IsSet()
}

// HasTemplate returns a boolean if a field has been set.
func (o *InsightResponse) HasTemplate() bool {
	if o != nil && o.Template.IsSet() {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given NullableString and assigns it to the Template field.
func (o *InsightResponse) SetTemplate(v string) {
	o.Template.Set(&v)
}
// SetTemplateNil sets the value for Template to be an explicit nil
func (o *InsightResponse) SetTemplateNil() {
	o.Template.Set(nil)
}

// UnsetTemplate ensures that no value is present for Template, not even an explicit nil
func (o *InsightResponse) UnsetTemplate() {
	o.Template.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *InsightResponse) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *InsightResponse) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *InsightResponse) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *InsightResponse) UnsetTitle() {
	o.Title.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InsightResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *InsightResponse) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *InsightResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *InsightResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUserGuid returns the UserGuid field value if set, zero value otherwise.
func (o *InsightResponse) GetUserGuid() string {
	if o == nil || IsNil(o.UserGuid) {
		var ret string
		return ret
	}
	return *o.UserGuid
}

// GetUserGuidOk returns a tuple with the UserGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightResponse) GetUserGuidOk() (*string, bool) {
	if o == nil || IsNil(o.UserGuid) {
		return nil, false
	}
	return o.UserGuid, true
}

// HasUserGuid returns a boolean if a field has been set.
func (o *InsightResponse) HasUserGuid() bool {
	if o != nil && !IsNil(o.UserGuid) {
		return true
	}

	return false
}

// SetUserGuid gets a reference to the given string and assigns it to the UserGuid field.
func (o *InsightResponse) SetUserGuid(v string) {
	o.UserGuid = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InsightResponse) GetUserId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InsightResponse) GetUserIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return &o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *InsightResponse) HasUserId() bool {
	if o != nil && IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given interface{} and assigns it to the UserId field.
func (o *InsightResponse) SetUserId(v interface{}) {
	o.UserId = v
}

func (o InsightResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsightResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveAt.IsSet() {
		toSerialize["active_at"] = o.ActiveAt.Get()
	}
	if o.ClientGuid.IsSet() {
		toSerialize["client_guid"] = o.ClientGuid.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.CtaClickedAt.IsSet() {
		toSerialize["cta_clicked_at"] = o.CtaClickedAt.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.HasAssociatedAccounts.IsSet() {
		toSerialize["has_associated_accounts"] = o.HasAssociatedAccounts.Get()
	}
	if o.HasAssociatedMerchants.IsSet() {
		toSerialize["has_associated_merchants"] = o.HasAssociatedMerchants.Get()
	}
	if o.HasAssociatedScheduledPayments.IsSet() {
		toSerialize["has_associated_scheduled_payments"] = o.HasAssociatedScheduledPayments.Get()
	}
	if o.HasAssociatedTransactions.IsSet() {
		toSerialize["has_associated_transactions"] = o.HasAssociatedTransactions.Get()
	}
	if o.HasBeenDisplayed.IsSet() {
		toSerialize["has_been_displayed"] = o.HasBeenDisplayed.Get()
	}
	if o.IsDismissed.IsSet() {
		toSerialize["is_dismissed"] = o.IsDismissed.Get()
	}
	if o.MicroCallToAction.IsSet() {
		toSerialize["micro_call_to_action"] = o.MicroCallToAction.Get()
	}
	if o.MicroDescription.IsSet() {
		toSerialize["micro_description"] = o.MicroDescription.Get()
	}
	if o.MicroTitle.IsSet() {
		toSerialize["micro_title"] = o.MicroTitle.Get()
	}
	if o.Template.IsSet() {
		toSerialize["template"] = o.Template.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if !IsNil(o.UserGuid) {
		toSerialize["user_guid"] = o.UserGuid
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableInsightResponse struct {
	value *InsightResponse
	isSet bool
}

func (v NullableInsightResponse) Get() *InsightResponse {
	return v.value
}

func (v *NullableInsightResponse) Set(val *InsightResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightResponse(val *InsightResponse) *NullableInsightResponse {
	return &NullableInsightResponse{value: val, isSet: true}
}

func (v NullableInsightResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsightResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

