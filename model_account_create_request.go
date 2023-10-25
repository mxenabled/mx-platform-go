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

// checks if the AccountCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountCreateRequest{}

// AccountCreateRequest struct for AccountCreateRequest
type AccountCreateRequest struct {
	AccountSubtype *string `json:"account_subtype,omitempty"`
	AccountType string `json:"account_type"`
	Apr *float32 `json:"apr,omitempty"`
	Apy *float32 `json:"apy,omitempty"`
	AvailableBalance *float32 `json:"available_balance,omitempty"`
	Balance *float32 `json:"balance,omitempty"`
	CashSurrenderValue *float32 `json:"cash_surrender_value,omitempty"`
	CreditLimit *float32 `json:"credit_limit,omitempty"`
	CurrencyCode *string `json:"currency_code,omitempty"`
	DeathBenefit *int32 `json:"death_benefit,omitempty"`
	InterestRate *float32 `json:"interest_rate,omitempty"`
	IsBusiness *bool `json:"is_business,omitempty"`
	IsClosed *bool `json:"is_closed,omitempty"`
	IsHidden *bool `json:"is_hidden,omitempty"`
	LoanAmount *float32 `json:"loan_amount,omitempty"`
	Metadata *string `json:"metadata,omitempty"`
	Name string `json:"name"`
	Nickname *string `json:"nickname,omitempty"`
	OriginalBalance *float32 `json:"original_balance,omitempty"`
	PropertyType *string `json:"property_type,omitempty"`
	SkipWebhook *bool `json:"skip_webhook,omitempty"`
}

// NewAccountCreateRequest instantiates a new AccountCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreateRequest(accountType string, name string) *AccountCreateRequest {
	this := AccountCreateRequest{}
	this.AccountType = accountType
	this.Name = name
	return &this
}

// NewAccountCreateRequestWithDefaults instantiates a new AccountCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreateRequestWithDefaults() *AccountCreateRequest {
	this := AccountCreateRequest{}
	return &this
}

// GetAccountSubtype returns the AccountSubtype field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetAccountSubtype() string {
	if o == nil || IsNil(o.AccountSubtype) {
		var ret string
		return ret
	}
	return *o.AccountSubtype
}

// GetAccountSubtypeOk returns a tuple with the AccountSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetAccountSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountSubtype) {
		return nil, false
	}
	return o.AccountSubtype, true
}

// HasAccountSubtype returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasAccountSubtype() bool {
	if o != nil && !IsNil(o.AccountSubtype) {
		return true
	}

	return false
}

// SetAccountSubtype gets a reference to the given string and assigns it to the AccountSubtype field.
func (o *AccountCreateRequest) SetAccountSubtype(v string) {
	o.AccountSubtype = &v
}

// GetAccountType returns the AccountType field value
func (o *AccountCreateRequest) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *AccountCreateRequest) SetAccountType(v string) {
	o.AccountType = v
}

// GetApr returns the Apr field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetApr() float32 {
	if o == nil || IsNil(o.Apr) {
		var ret float32
		return ret
	}
	return *o.Apr
}

// GetAprOk returns a tuple with the Apr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetAprOk() (*float32, bool) {
	if o == nil || IsNil(o.Apr) {
		return nil, false
	}
	return o.Apr, true
}

// HasApr returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasApr() bool {
	if o != nil && !IsNil(o.Apr) {
		return true
	}

	return false
}

// SetApr gets a reference to the given float32 and assigns it to the Apr field.
func (o *AccountCreateRequest) SetApr(v float32) {
	o.Apr = &v
}

// GetApy returns the Apy field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetApy() float32 {
	if o == nil || IsNil(o.Apy) {
		var ret float32
		return ret
	}
	return *o.Apy
}

// GetApyOk returns a tuple with the Apy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetApyOk() (*float32, bool) {
	if o == nil || IsNil(o.Apy) {
		return nil, false
	}
	return o.Apy, true
}

// HasApy returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasApy() bool {
	if o != nil && !IsNil(o.Apy) {
		return true
	}

	return false
}

// SetApy gets a reference to the given float32 and assigns it to the Apy field.
func (o *AccountCreateRequest) SetApy(v float32) {
	o.Apy = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetAvailableBalance() float32 {
	if o == nil || IsNil(o.AvailableBalance) {
		var ret float32
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetAvailableBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasAvailableBalance() bool {
	if o != nil && !IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given float32 and assigns it to the AvailableBalance field.
func (o *AccountCreateRequest) SetAvailableBalance(v float32) {
	o.AvailableBalance = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *AccountCreateRequest) SetBalance(v float32) {
	o.Balance = &v
}

// GetCashSurrenderValue returns the CashSurrenderValue field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetCashSurrenderValue() float32 {
	if o == nil || IsNil(o.CashSurrenderValue) {
		var ret float32
		return ret
	}
	return *o.CashSurrenderValue
}

// GetCashSurrenderValueOk returns a tuple with the CashSurrenderValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetCashSurrenderValueOk() (*float32, bool) {
	if o == nil || IsNil(o.CashSurrenderValue) {
		return nil, false
	}
	return o.CashSurrenderValue, true
}

// HasCashSurrenderValue returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasCashSurrenderValue() bool {
	if o != nil && !IsNil(o.CashSurrenderValue) {
		return true
	}

	return false
}

// SetCashSurrenderValue gets a reference to the given float32 and assigns it to the CashSurrenderValue field.
func (o *AccountCreateRequest) SetCashSurrenderValue(v float32) {
	o.CashSurrenderValue = &v
}

// GetCreditLimit returns the CreditLimit field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetCreditLimit() float32 {
	if o == nil || IsNil(o.CreditLimit) {
		var ret float32
		return ret
	}
	return *o.CreditLimit
}

// GetCreditLimitOk returns a tuple with the CreditLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetCreditLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.CreditLimit) {
		return nil, false
	}
	return o.CreditLimit, true
}

// HasCreditLimit returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasCreditLimit() bool {
	if o != nil && !IsNil(o.CreditLimit) {
		return true
	}

	return false
}

// SetCreditLimit gets a reference to the given float32 and assigns it to the CreditLimit field.
func (o *AccountCreateRequest) SetCreditLimit(v float32) {
	o.CreditLimit = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *AccountCreateRequest) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDeathBenefit returns the DeathBenefit field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetDeathBenefit() int32 {
	if o == nil || IsNil(o.DeathBenefit) {
		var ret int32
		return ret
	}
	return *o.DeathBenefit
}

// GetDeathBenefitOk returns a tuple with the DeathBenefit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetDeathBenefitOk() (*int32, bool) {
	if o == nil || IsNil(o.DeathBenefit) {
		return nil, false
	}
	return o.DeathBenefit, true
}

// HasDeathBenefit returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasDeathBenefit() bool {
	if o != nil && !IsNil(o.DeathBenefit) {
		return true
	}

	return false
}

// SetDeathBenefit gets a reference to the given int32 and assigns it to the DeathBenefit field.
func (o *AccountCreateRequest) SetDeathBenefit(v int32) {
	o.DeathBenefit = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetInterestRate() float32 {
	if o == nil || IsNil(o.InterestRate) {
		var ret float32
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetInterestRateOk() (*float32, bool) {
	if o == nil || IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasInterestRate() bool {
	if o != nil && !IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given float32 and assigns it to the InterestRate field.
func (o *AccountCreateRequest) SetInterestRate(v float32) {
	o.InterestRate = &v
}

// GetIsBusiness returns the IsBusiness field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetIsBusiness() bool {
	if o == nil || IsNil(o.IsBusiness) {
		var ret bool
		return ret
	}
	return *o.IsBusiness
}

// GetIsBusinessOk returns a tuple with the IsBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetIsBusinessOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBusiness) {
		return nil, false
	}
	return o.IsBusiness, true
}

// HasIsBusiness returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasIsBusiness() bool {
	if o != nil && !IsNil(o.IsBusiness) {
		return true
	}

	return false
}

// SetIsBusiness gets a reference to the given bool and assigns it to the IsBusiness field.
func (o *AccountCreateRequest) SetIsBusiness(v bool) {
	o.IsBusiness = &v
}

// GetIsClosed returns the IsClosed field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetIsClosed() bool {
	if o == nil || IsNil(o.IsClosed) {
		var ret bool
		return ret
	}
	return *o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetIsClosedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsClosed) {
		return nil, false
	}
	return o.IsClosed, true
}

// HasIsClosed returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasIsClosed() bool {
	if o != nil && !IsNil(o.IsClosed) {
		return true
	}

	return false
}

// SetIsClosed gets a reference to the given bool and assigns it to the IsClosed field.
func (o *AccountCreateRequest) SetIsClosed(v bool) {
	o.IsClosed = &v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetIsHidden() bool {
	if o == nil || IsNil(o.IsHidden) {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetIsHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHidden) {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasIsHidden() bool {
	if o != nil && !IsNil(o.IsHidden) {
		return true
	}

	return false
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *AccountCreateRequest) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetLoanAmount returns the LoanAmount field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetLoanAmount() float32 {
	if o == nil || IsNil(o.LoanAmount) {
		var ret float32
		return ret
	}
	return *o.LoanAmount
}

// GetLoanAmountOk returns a tuple with the LoanAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetLoanAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.LoanAmount) {
		return nil, false
	}
	return o.LoanAmount, true
}

// HasLoanAmount returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasLoanAmount() bool {
	if o != nil && !IsNil(o.LoanAmount) {
		return true
	}

	return false
}

// SetLoanAmount gets a reference to the given float32 and assigns it to the LoanAmount field.
func (o *AccountCreateRequest) SetLoanAmount(v float32) {
	o.LoanAmount = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *AccountCreateRequest) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *AccountCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountCreateRequest) SetName(v string) {
	o.Name = v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetNickname() string {
	if o == nil || IsNil(o.Nickname) {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname) {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasNickname() bool {
	if o != nil && !IsNil(o.Nickname) {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *AccountCreateRequest) SetNickname(v string) {
	o.Nickname = &v
}

// GetOriginalBalance returns the OriginalBalance field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetOriginalBalance() float32 {
	if o == nil || IsNil(o.OriginalBalance) {
		var ret float32
		return ret
	}
	return *o.OriginalBalance
}

// GetOriginalBalanceOk returns a tuple with the OriginalBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetOriginalBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.OriginalBalance) {
		return nil, false
	}
	return o.OriginalBalance, true
}

// HasOriginalBalance returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasOriginalBalance() bool {
	if o != nil && !IsNil(o.OriginalBalance) {
		return true
	}

	return false
}

// SetOriginalBalance gets a reference to the given float32 and assigns it to the OriginalBalance field.
func (o *AccountCreateRequest) SetOriginalBalance(v float32) {
	o.OriginalBalance = &v
}

// GetPropertyType returns the PropertyType field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetPropertyType() string {
	if o == nil || IsNil(o.PropertyType) {
		var ret string
		return ret
	}
	return *o.PropertyType
}

// GetPropertyTypeOk returns a tuple with the PropertyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetPropertyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyType) {
		return nil, false
	}
	return o.PropertyType, true
}

// HasPropertyType returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasPropertyType() bool {
	if o != nil && !IsNil(o.PropertyType) {
		return true
	}

	return false
}

// SetPropertyType gets a reference to the given string and assigns it to the PropertyType field.
func (o *AccountCreateRequest) SetPropertyType(v string) {
	o.PropertyType = &v
}

// GetSkipWebhook returns the SkipWebhook field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetSkipWebhook() bool {
	if o == nil || IsNil(o.SkipWebhook) {
		var ret bool
		return ret
	}
	return *o.SkipWebhook
}

// GetSkipWebhookOk returns a tuple with the SkipWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetSkipWebhookOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipWebhook) {
		return nil, false
	}
	return o.SkipWebhook, true
}

// HasSkipWebhook returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasSkipWebhook() bool {
	if o != nil && !IsNil(o.SkipWebhook) {
		return true
	}

	return false
}

// SetSkipWebhook gets a reference to the given bool and assigns it to the SkipWebhook field.
func (o *AccountCreateRequest) SetSkipWebhook(v bool) {
	o.SkipWebhook = &v
}

func (o AccountCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountSubtype) {
		toSerialize["account_subtype"] = o.AccountSubtype
	}
	toSerialize["account_type"] = o.AccountType
	if !IsNil(o.Apr) {
		toSerialize["apr"] = o.Apr
	}
	if !IsNil(o.Apy) {
		toSerialize["apy"] = o.Apy
	}
	if !IsNil(o.AvailableBalance) {
		toSerialize["available_balance"] = o.AvailableBalance
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.CashSurrenderValue) {
		toSerialize["cash_surrender_value"] = o.CashSurrenderValue
	}
	if !IsNil(o.CreditLimit) {
		toSerialize["credit_limit"] = o.CreditLimit
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.DeathBenefit) {
		toSerialize["death_benefit"] = o.DeathBenefit
	}
	if !IsNil(o.InterestRate) {
		toSerialize["interest_rate"] = o.InterestRate
	}
	if !IsNil(o.IsBusiness) {
		toSerialize["is_business"] = o.IsBusiness
	}
	if !IsNil(o.IsClosed) {
		toSerialize["is_closed"] = o.IsClosed
	}
	if !IsNil(o.IsHidden) {
		toSerialize["is_hidden"] = o.IsHidden
	}
	if !IsNil(o.LoanAmount) {
		toSerialize["loan_amount"] = o.LoanAmount
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Nickname) {
		toSerialize["nickname"] = o.Nickname
	}
	if !IsNil(o.OriginalBalance) {
		toSerialize["original_balance"] = o.OriginalBalance
	}
	if !IsNil(o.PropertyType) {
		toSerialize["property_type"] = o.PropertyType
	}
	if !IsNil(o.SkipWebhook) {
		toSerialize["skip_webhook"] = o.SkipWebhook
	}
	return toSerialize, nil
}

type NullableAccountCreateRequest struct {
	value *AccountCreateRequest
	isSet bool
}

func (v NullableAccountCreateRequest) Get() *AccountCreateRequest {
	return v.value
}

func (v *NullableAccountCreateRequest) Set(val *AccountCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreateRequest(val *AccountCreateRequest) *NullableAccountCreateRequest {
	return &NullableAccountCreateRequest{value: val, isSet: true}
}

func (v NullableAccountCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


