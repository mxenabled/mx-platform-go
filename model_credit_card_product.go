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

// checks if the CreditCardProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditCardProduct{}

// CreditCardProduct struct for CreditCardProduct
type CreditCardProduct struct {
	AnnualFee interface{} `json:"annual_fee,omitempty"`
	DurationOfIntroductoryRateOnBalanceTransfer interface{} `json:"duration_of_introductory_rate_on_balance_transfer,omitempty"`
	DurationOfIntroductoryRateOnPurchases interface{} `json:"duration_of_introductory_rate_on_purchases,omitempty"`
	Guid interface{} `json:"guid,omitempty"`
	HasCashbackRewards *bool `json:"has_cashback_rewards,omitempty"`
	HasOtherRewards *bool `json:"has_other_rewards,omitempty"`
	HasTravelRewards *bool `json:"has_travel_rewards,omitempty"`
	HasZeroIntroductoryAnnualFee *bool `json:"has_zero_introductory_annual_fee,omitempty"`
	HasZeroPercentIntroductoryRate *bool `json:"has_zero_percent_introductory_rate,omitempty"`
	HasZeroPercentIntroductoryRateOnBalanceTransfer *bool `json:"has_zero_percent_introductory_rate_on_balance_transfer,omitempty"`
	FinancialInstitution *bool `json:"financial_institution,omitempty"`
	IsAcceptingApplications *bool `json:"is_accepting_applications,omitempty"`
	IsSmallBusinessCard *bool `json:"is_small_business_card,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewCreditCardProduct instantiates a new CreditCardProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCardProduct() *CreditCardProduct {
	this := CreditCardProduct{}
	return &this
}

// NewCreditCardProductWithDefaults instantiates a new CreditCardProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCardProductWithDefaults() *CreditCardProduct {
	this := CreditCardProduct{}
	return &this
}

// GetAnnualFee returns the AnnualFee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditCardProduct) GetAnnualFee() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AnnualFee
}

// GetAnnualFeeOk returns a tuple with the AnnualFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditCardProduct) GetAnnualFeeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AnnualFee) {
		return nil, false
	}
	return &o.AnnualFee, true
}

// HasAnnualFee returns a boolean if a field has been set.
func (o *CreditCardProduct) HasAnnualFee() bool {
	if o != nil && IsNil(o.AnnualFee) {
		return true
	}

	return false
}

// SetAnnualFee gets a reference to the given interface{} and assigns it to the AnnualFee field.
func (o *CreditCardProduct) SetAnnualFee(v interface{}) {
	o.AnnualFee = v
}

// GetDurationOfIntroductoryRateOnBalanceTransfer returns the DurationOfIntroductoryRateOnBalanceTransfer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditCardProduct) GetDurationOfIntroductoryRateOnBalanceTransfer() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DurationOfIntroductoryRateOnBalanceTransfer
}

// GetDurationOfIntroductoryRateOnBalanceTransferOk returns a tuple with the DurationOfIntroductoryRateOnBalanceTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditCardProduct) GetDurationOfIntroductoryRateOnBalanceTransferOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DurationOfIntroductoryRateOnBalanceTransfer) {
		return nil, false
	}
	return &o.DurationOfIntroductoryRateOnBalanceTransfer, true
}

// HasDurationOfIntroductoryRateOnBalanceTransfer returns a boolean if a field has been set.
func (o *CreditCardProduct) HasDurationOfIntroductoryRateOnBalanceTransfer() bool {
	if o != nil && IsNil(o.DurationOfIntroductoryRateOnBalanceTransfer) {
		return true
	}

	return false
}

// SetDurationOfIntroductoryRateOnBalanceTransfer gets a reference to the given interface{} and assigns it to the DurationOfIntroductoryRateOnBalanceTransfer field.
func (o *CreditCardProduct) SetDurationOfIntroductoryRateOnBalanceTransfer(v interface{}) {
	o.DurationOfIntroductoryRateOnBalanceTransfer = v
}

// GetDurationOfIntroductoryRateOnPurchases returns the DurationOfIntroductoryRateOnPurchases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditCardProduct) GetDurationOfIntroductoryRateOnPurchases() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DurationOfIntroductoryRateOnPurchases
}

// GetDurationOfIntroductoryRateOnPurchasesOk returns a tuple with the DurationOfIntroductoryRateOnPurchases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditCardProduct) GetDurationOfIntroductoryRateOnPurchasesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DurationOfIntroductoryRateOnPurchases) {
		return nil, false
	}
	return &o.DurationOfIntroductoryRateOnPurchases, true
}

// HasDurationOfIntroductoryRateOnPurchases returns a boolean if a field has been set.
func (o *CreditCardProduct) HasDurationOfIntroductoryRateOnPurchases() bool {
	if o != nil && IsNil(o.DurationOfIntroductoryRateOnPurchases) {
		return true
	}

	return false
}

// SetDurationOfIntroductoryRateOnPurchases gets a reference to the given interface{} and assigns it to the DurationOfIntroductoryRateOnPurchases field.
func (o *CreditCardProduct) SetDurationOfIntroductoryRateOnPurchases(v interface{}) {
	o.DurationOfIntroductoryRateOnPurchases = v
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditCardProduct) GetGuid() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditCardProduct) GetGuidOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return &o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *CreditCardProduct) HasGuid() bool {
	if o != nil && IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given interface{} and assigns it to the Guid field.
func (o *CreditCardProduct) SetGuid(v interface{}) {
	o.Guid = v
}

// GetHasCashbackRewards returns the HasCashbackRewards field value if set, zero value otherwise.
func (o *CreditCardProduct) GetHasCashbackRewards() bool {
	if o == nil || IsNil(o.HasCashbackRewards) {
		var ret bool
		return ret
	}
	return *o.HasCashbackRewards
}

// GetHasCashbackRewardsOk returns a tuple with the HasCashbackRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardProduct) GetHasCashbackRewardsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasCashbackRewards) {
		return nil, false
	}
	return o.HasCashbackRewards, true
}

// HasHasCashbackRewards returns a boolean if a field has been set.
func (o *CreditCardProduct) HasHasCashbackRewards() bool {
	if o != nil && !IsNil(o.HasCashbackRewards) {
		return true
	}

	return false
}

// SetHasCashbackRewards gets a reference to the given bool and assigns it to the HasCashbackRewards field.
func (o *CreditCardProduct) SetHasCashbackRewards(v bool) {
	o.HasCashbackRewards = &v
}

// GetHasOtherRewards returns the HasOtherRewards field value if set, zero value otherwise.
func (o *CreditCardProduct) GetHasOtherRewards() bool {
	if o == nil || IsNil(o.HasOtherRewards) {
		var ret bool
		return ret
	}
	return *o.HasOtherRewards
}

// GetHasOtherRewardsOk returns a tuple with the HasOtherRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardProduct) GetHasOtherRewardsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasOtherRewards) {
		return nil, false
	}
	return o.HasOtherRewards, true
}

// HasHasOtherRewards returns a boolean if a field has been set.
func (o *CreditCardProduct) HasHasOtherRewards() bool {
	if o != nil && !IsNil(o.HasOtherRewards) {
		return true
	}

	return false
}

// SetHasOtherRewards gets a reference to the given bool and assigns it to the HasOtherRewards field.
func (o *CreditCardProduct) SetHasOtherRewards(v bool) {
	o.HasOtherRewards = &v
}

// GetHasTravelRewards returns the HasTravelRewards field value if set, zero value otherwise.
func (o *CreditCardProduct) GetHasTravelRewards() bool {
	if o == nil || IsNil(o.HasTravelRewards) {
		var ret bool
		return ret
	}
	return *o.HasTravelRewards
}

// GetHasTravelRewardsOk returns a tuple with the HasTravelRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardProduct) GetHasTravelRewardsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasTravelRewards) {
		return nil, false
	}
	return o.HasTravelRewards, true
}

// HasHasTravelRewards returns a boolean if a field has been set.
func (o *CreditCardProduct) HasHasTravelRewards() bool {
	if o != nil && !IsNil(o.HasTravelRewards) {
		return true
	}

	return false
}

// SetHasTravelRewards gets a reference to the given bool and assigns it to the HasTravelRewards field.
func (o *CreditCardProduct) SetHasTravelRewards(v bool) {
	o.HasTravelRewards = &v
}

// GetHasZeroIntroductoryAnnualFee returns the HasZeroIntroductoryAnnualFee field value if set, zero value otherwise.
func (o *CreditCardProduct) GetHasZeroIntroductoryAnnualFee() bool {
	if o == nil || IsNil(o.HasZeroIntroductoryAnnualFee) {
		var ret bool
		return ret
	}
	return *o.HasZeroIntroductoryAnnualFee
}

// GetHasZeroIntroductoryAnnualFeeOk returns a tuple with the HasZeroIntroductoryAnnualFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardProduct) GetHasZeroIntroductoryAnnualFeeOk() (*bool, bool) {
	if o == nil || IsNil(o.HasZeroIntroductoryAnnualFee) {
		return nil, false
	}
	return o.HasZeroIntroductoryAnnualFee, true
}

// HasHasZeroIntroductoryAnnualFee returns a boolean if a field has been set.
func (o *CreditCardProduct) HasHasZeroIntroductoryAnnualFee() bool {
	if o != nil && !IsNil(o.HasZeroIntroductoryAnnualFee) {
		return true
	}

	return false
}

// SetHasZeroIntroductoryAnnualFee gets a reference to the given bool and assigns it to the HasZeroIntroductoryAnnualFee field.
func (o *CreditCardProduct) SetHasZeroIntroductoryAnnualFee(v bool) {
	o.HasZeroIntroductoryAnnualFee = &v
}

// GetHasZeroPercentIntroductoryRate returns the HasZeroPercentIntroductoryRate field value if set, zero value otherwise.
func (o *CreditCardProduct) GetHasZeroPercentIntroductoryRate() bool {
	if o == nil || IsNil(o.HasZeroPercentIntroductoryRate) {
		var ret bool
		return ret
	}
	return *o.HasZeroPercentIntroductoryRate
}

// GetHasZeroPercentIntroductoryRateOk returns a tuple with the HasZeroPercentIntroductoryRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardProduct) GetHasZeroPercentIntroductoryRateOk() (*bool, bool) {
	if o == nil || IsNil(o.HasZeroPercentIntroductoryRate) {
		return nil, false
	}
	return o.HasZeroPercentIntroductoryRate, true
}

// HasHasZeroPercentIntroductoryRate returns a boolean if a field has been set.
func (o *CreditCardProduct) HasHasZeroPercentIntroductoryRate() bool {
	if o != nil && !IsNil(o.HasZeroPercentIntroductoryRate) {
		return true
	}

	return false
}

// SetHasZeroPercentIntroductoryRate gets a reference to the given bool and assigns it to the HasZeroPercentIntroductoryRate field.
func (o *CreditCardProduct) SetHasZeroPercentIntroductoryRate(v bool) {
	o.HasZeroPercentIntroductoryRate = &v
}

// GetHasZeroPercentIntroductoryRateOnBalanceTransfer returns the HasZeroPercentIntroductoryRateOnBalanceTransfer field value if set, zero value otherwise.
func (o *CreditCardProduct) GetHasZeroPercentIntroductoryRateOnBalanceTransfer() bool {
	if o == nil || IsNil(o.HasZeroPercentIntroductoryRateOnBalanceTransfer) {
		var ret bool
		return ret
	}
	return *o.HasZeroPercentIntroductoryRateOnBalanceTransfer
}

// GetHasZeroPercentIntroductoryRateOnBalanceTransferOk returns a tuple with the HasZeroPercentIntroductoryRateOnBalanceTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardProduct) GetHasZeroPercentIntroductoryRateOnBalanceTransferOk() (*bool, bool) {
	if o == nil || IsNil(o.HasZeroPercentIntroductoryRateOnBalanceTransfer) {
		return nil, false
	}
	return o.HasZeroPercentIntroductoryRateOnBalanceTransfer, true
}

// HasHasZeroPercentIntroductoryRateOnBalanceTransfer returns a boolean if a field has been set.
func (o *CreditCardProduct) HasHasZeroPercentIntroductoryRateOnBalanceTransfer() bool {
	if o != nil && !IsNil(o.HasZeroPercentIntroductoryRateOnBalanceTransfer) {
		return true
	}

	return false
}

// SetHasZeroPercentIntroductoryRateOnBalanceTransfer gets a reference to the given bool and assigns it to the HasZeroPercentIntroductoryRateOnBalanceTransfer field.
func (o *CreditCardProduct) SetHasZeroPercentIntroductoryRateOnBalanceTransfer(v bool) {
	o.HasZeroPercentIntroductoryRateOnBalanceTransfer = &v
}

// GetFinancialInstitution returns the FinancialInstitution field value if set, zero value otherwise.
func (o *CreditCardProduct) GetFinancialInstitution() bool {
	if o == nil || IsNil(o.FinancialInstitution) {
		var ret bool
		return ret
	}
	return *o.FinancialInstitution
}

// GetFinancialInstitutionOk returns a tuple with the FinancialInstitution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardProduct) GetFinancialInstitutionOk() (*bool, bool) {
	if o == nil || IsNil(o.FinancialInstitution) {
		return nil, false
	}
	return o.FinancialInstitution, true
}

// HasFinancialInstitution returns a boolean if a field has been set.
func (o *CreditCardProduct) HasFinancialInstitution() bool {
	if o != nil && !IsNil(o.FinancialInstitution) {
		return true
	}

	return false
}

// SetFinancialInstitution gets a reference to the given bool and assigns it to the FinancialInstitution field.
func (o *CreditCardProduct) SetFinancialInstitution(v bool) {
	o.FinancialInstitution = &v
}

// GetIsAcceptingApplications returns the IsAcceptingApplications field value if set, zero value otherwise.
func (o *CreditCardProduct) GetIsAcceptingApplications() bool {
	if o == nil || IsNil(o.IsAcceptingApplications) {
		var ret bool
		return ret
	}
	return *o.IsAcceptingApplications
}

// GetIsAcceptingApplicationsOk returns a tuple with the IsAcceptingApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardProduct) GetIsAcceptingApplicationsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAcceptingApplications) {
		return nil, false
	}
	return o.IsAcceptingApplications, true
}

// HasIsAcceptingApplications returns a boolean if a field has been set.
func (o *CreditCardProduct) HasIsAcceptingApplications() bool {
	if o != nil && !IsNil(o.IsAcceptingApplications) {
		return true
	}

	return false
}

// SetIsAcceptingApplications gets a reference to the given bool and assigns it to the IsAcceptingApplications field.
func (o *CreditCardProduct) SetIsAcceptingApplications(v bool) {
	o.IsAcceptingApplications = &v
}

// GetIsSmallBusinessCard returns the IsSmallBusinessCard field value if set, zero value otherwise.
func (o *CreditCardProduct) GetIsSmallBusinessCard() bool {
	if o == nil || IsNil(o.IsSmallBusinessCard) {
		var ret bool
		return ret
	}
	return *o.IsSmallBusinessCard
}

// GetIsSmallBusinessCardOk returns a tuple with the IsSmallBusinessCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardProduct) GetIsSmallBusinessCardOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSmallBusinessCard) {
		return nil, false
	}
	return o.IsSmallBusinessCard, true
}

// HasIsSmallBusinessCard returns a boolean if a field has been set.
func (o *CreditCardProduct) HasIsSmallBusinessCard() bool {
	if o != nil && !IsNil(o.IsSmallBusinessCard) {
		return true
	}

	return false
}

// SetIsSmallBusinessCard gets a reference to the given bool and assigns it to the IsSmallBusinessCard field.
func (o *CreditCardProduct) SetIsSmallBusinessCard(v bool) {
	o.IsSmallBusinessCard = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreditCardProduct) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardProduct) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreditCardProduct) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreditCardProduct) SetName(v string) {
	o.Name = &v
}

func (o CreditCardProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditCardProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AnnualFee != nil {
		toSerialize["annual_fee"] = o.AnnualFee
	}
	if o.DurationOfIntroductoryRateOnBalanceTransfer != nil {
		toSerialize["duration_of_introductory_rate_on_balance_transfer"] = o.DurationOfIntroductoryRateOnBalanceTransfer
	}
	if o.DurationOfIntroductoryRateOnPurchases != nil {
		toSerialize["duration_of_introductory_rate_on_purchases"] = o.DurationOfIntroductoryRateOnPurchases
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.HasCashbackRewards) {
		toSerialize["has_cashback_rewards"] = o.HasCashbackRewards
	}
	if !IsNil(o.HasOtherRewards) {
		toSerialize["has_other_rewards"] = o.HasOtherRewards
	}
	if !IsNil(o.HasTravelRewards) {
		toSerialize["has_travel_rewards"] = o.HasTravelRewards
	}
	if !IsNil(o.HasZeroIntroductoryAnnualFee) {
		toSerialize["has_zero_introductory_annual_fee"] = o.HasZeroIntroductoryAnnualFee
	}
	if !IsNil(o.HasZeroPercentIntroductoryRate) {
		toSerialize["has_zero_percent_introductory_rate"] = o.HasZeroPercentIntroductoryRate
	}
	if !IsNil(o.HasZeroPercentIntroductoryRateOnBalanceTransfer) {
		toSerialize["has_zero_percent_introductory_rate_on_balance_transfer"] = o.HasZeroPercentIntroductoryRateOnBalanceTransfer
	}
	if !IsNil(o.FinancialInstitution) {
		toSerialize["financial_institution"] = o.FinancialInstitution
	}
	if !IsNil(o.IsAcceptingApplications) {
		toSerialize["is_accepting_applications"] = o.IsAcceptingApplications
	}
	if !IsNil(o.IsSmallBusinessCard) {
		toSerialize["is_small_business_card"] = o.IsSmallBusinessCard
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCreditCardProduct struct {
	value *CreditCardProduct
	isSet bool
}

func (v NullableCreditCardProduct) Get() *CreditCardProduct {
	return v.value
}

func (v *NullableCreditCardProduct) Set(val *CreditCardProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCardProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCardProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCardProduct(val *CreditCardProduct) *NullableCreditCardProduct {
	return &NullableCreditCardProduct{value: val, isSet: true}
}

func (v NullableCreditCardProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCardProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


