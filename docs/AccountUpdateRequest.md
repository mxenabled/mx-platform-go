# AccountUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSubtype** | Pointer to **string** | Can only be updated for manual accounts. | [optional] 
**AccountType** | Pointer to **string** | Can only be updated for manual accounts. | [optional] 
**Apr** | Pointer to **float32** | Can only be updated for manual accounts. | [optional] 
**Apy** | Pointer to **float32** | Can only be updated for manual accounts. | [optional] 
**AvailableBalance** | Pointer to **float32** | Can only be updated for manual accounts. | [optional] 
**Balance** | Pointer to **float32** | Can only be updated for manual accounts. | [optional] 
**CashSurrenderValue** | Pointer to **float32** | Can only be updated for manual accounts. | [optional] 
**CreditLimit** | Pointer to **float32** | Can only be updated for manual accounts. | [optional] 
**CurrencyCode** | Pointer to **string** | Can only be updated for manual accounts. | [optional] 
**DeathBenefit** | Pointer to **int32** | Can only be updated for manual accounts. | [optional] 
**InterestRate** | Pointer to **float32** | Can only be updated for manual accounts. | [optional] 
**IsBusiness** | Pointer to **bool** | Can be updated for manual accounts and aggregated accounts. | [optional] 
**IsClosed** | Pointer to **bool** | Can only be updated for manual accounts. | [optional] 
**IsHidden** | Pointer to **bool** | Can be updated for manual accounts and aggregated accounts. | [optional] 
**LoanAmount** | Pointer to **float32** | Can only be updated for manual accounts. | [optional] 
**Metadata** | Pointer to **string** | Can only be updated for manual accounts. | [optional] 
**Name** | Pointer to **string** | Can only be updated for manual accounts. | [optional] 
**Nickname** | Pointer to **string** | Can only be updated for manual accounts. | [optional] 
**OriginalBalance** | Pointer to **float32** | Can only be updated for manual accounts. | [optional] 
**PropertyType** | Pointer to **string** | Can only be updated for manual accounts. | [optional] 
**SkipWebhook** | Pointer to **bool** | If set to true, prevents sending an account webhook for the update if that webhook type is enabled for you. | [optional] 

## Methods

### NewAccountUpdateRequest

`func NewAccountUpdateRequest() *AccountUpdateRequest`

NewAccountUpdateRequest instantiates a new AccountUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateRequestWithDefaults

`func NewAccountUpdateRequestWithDefaults() *AccountUpdateRequest`

NewAccountUpdateRequestWithDefaults instantiates a new AccountUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSubtype

`func (o *AccountUpdateRequest) GetAccountSubtype() string`

GetAccountSubtype returns the AccountSubtype field if non-nil, zero value otherwise.

### GetAccountSubtypeOk

`func (o *AccountUpdateRequest) GetAccountSubtypeOk() (*string, bool)`

GetAccountSubtypeOk returns a tuple with the AccountSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSubtype

`func (o *AccountUpdateRequest) SetAccountSubtype(v string)`

SetAccountSubtype sets AccountSubtype field to given value.

### HasAccountSubtype

`func (o *AccountUpdateRequest) HasAccountSubtype() bool`

HasAccountSubtype returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountUpdateRequest) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountUpdateRequest) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountUpdateRequest) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountUpdateRequest) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetApr

`func (o *AccountUpdateRequest) GetApr() float32`

GetApr returns the Apr field if non-nil, zero value otherwise.

### GetAprOk

`func (o *AccountUpdateRequest) GetAprOk() (*float32, bool)`

GetAprOk returns a tuple with the Apr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApr

`func (o *AccountUpdateRequest) SetApr(v float32)`

SetApr sets Apr field to given value.

### HasApr

`func (o *AccountUpdateRequest) HasApr() bool`

HasApr returns a boolean if a field has been set.

### GetApy

`func (o *AccountUpdateRequest) GetApy() float32`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *AccountUpdateRequest) GetApyOk() (*float32, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *AccountUpdateRequest) SetApy(v float32)`

SetApy sets Apy field to given value.

### HasApy

`func (o *AccountUpdateRequest) HasApy() bool`

HasApy returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *AccountUpdateRequest) GetAvailableBalance() float32`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *AccountUpdateRequest) GetAvailableBalanceOk() (*float32, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *AccountUpdateRequest) SetAvailableBalance(v float32)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *AccountUpdateRequest) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetBalance

`func (o *AccountUpdateRequest) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountUpdateRequest) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountUpdateRequest) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AccountUpdateRequest) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCashSurrenderValue

`func (o *AccountUpdateRequest) GetCashSurrenderValue() float32`

GetCashSurrenderValue returns the CashSurrenderValue field if non-nil, zero value otherwise.

### GetCashSurrenderValueOk

`func (o *AccountUpdateRequest) GetCashSurrenderValueOk() (*float32, bool)`

GetCashSurrenderValueOk returns a tuple with the CashSurrenderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashSurrenderValue

`func (o *AccountUpdateRequest) SetCashSurrenderValue(v float32)`

SetCashSurrenderValue sets CashSurrenderValue field to given value.

### HasCashSurrenderValue

`func (o *AccountUpdateRequest) HasCashSurrenderValue() bool`

HasCashSurrenderValue returns a boolean if a field has been set.

### GetCreditLimit

`func (o *AccountUpdateRequest) GetCreditLimit() float32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *AccountUpdateRequest) GetCreditLimitOk() (*float32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *AccountUpdateRequest) SetCreditLimit(v float32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *AccountUpdateRequest) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AccountUpdateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountUpdateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountUpdateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AccountUpdateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDeathBenefit

`func (o *AccountUpdateRequest) GetDeathBenefit() int32`

GetDeathBenefit returns the DeathBenefit field if non-nil, zero value otherwise.

### GetDeathBenefitOk

`func (o *AccountUpdateRequest) GetDeathBenefitOk() (*int32, bool)`

GetDeathBenefitOk returns a tuple with the DeathBenefit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathBenefit

`func (o *AccountUpdateRequest) SetDeathBenefit(v int32)`

SetDeathBenefit sets DeathBenefit field to given value.

### HasDeathBenefit

`func (o *AccountUpdateRequest) HasDeathBenefit() bool`

HasDeathBenefit returns a boolean if a field has been set.

### GetInterestRate

`func (o *AccountUpdateRequest) GetInterestRate() float32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *AccountUpdateRequest) GetInterestRateOk() (*float32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *AccountUpdateRequest) SetInterestRate(v float32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *AccountUpdateRequest) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetIsBusiness

`func (o *AccountUpdateRequest) GetIsBusiness() bool`

GetIsBusiness returns the IsBusiness field if non-nil, zero value otherwise.

### GetIsBusinessOk

`func (o *AccountUpdateRequest) GetIsBusinessOk() (*bool, bool)`

GetIsBusinessOk returns a tuple with the IsBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBusiness

`func (o *AccountUpdateRequest) SetIsBusiness(v bool)`

SetIsBusiness sets IsBusiness field to given value.

### HasIsBusiness

`func (o *AccountUpdateRequest) HasIsBusiness() bool`

HasIsBusiness returns a boolean if a field has been set.

### GetIsClosed

`func (o *AccountUpdateRequest) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *AccountUpdateRequest) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *AccountUpdateRequest) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *AccountUpdateRequest) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetIsHidden

`func (o *AccountUpdateRequest) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *AccountUpdateRequest) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *AccountUpdateRequest) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *AccountUpdateRequest) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetLoanAmount

`func (o *AccountUpdateRequest) GetLoanAmount() float32`

GetLoanAmount returns the LoanAmount field if non-nil, zero value otherwise.

### GetLoanAmountOk

`func (o *AccountUpdateRequest) GetLoanAmountOk() (*float32, bool)`

GetLoanAmountOk returns a tuple with the LoanAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanAmount

`func (o *AccountUpdateRequest) SetLoanAmount(v float32)`

SetLoanAmount sets LoanAmount field to given value.

### HasLoanAmount

`func (o *AccountUpdateRequest) HasLoanAmount() bool`

HasLoanAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountUpdateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountUpdateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountUpdateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *AccountUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *AccountUpdateRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountUpdateRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountUpdateRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountUpdateRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetOriginalBalance

`func (o *AccountUpdateRequest) GetOriginalBalance() float32`

GetOriginalBalance returns the OriginalBalance field if non-nil, zero value otherwise.

### GetOriginalBalanceOk

`func (o *AccountUpdateRequest) GetOriginalBalanceOk() (*float32, bool)`

GetOriginalBalanceOk returns a tuple with the OriginalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalBalance

`func (o *AccountUpdateRequest) SetOriginalBalance(v float32)`

SetOriginalBalance sets OriginalBalance field to given value.

### HasOriginalBalance

`func (o *AccountUpdateRequest) HasOriginalBalance() bool`

HasOriginalBalance returns a boolean if a field has been set.

### GetPropertyType

`func (o *AccountUpdateRequest) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *AccountUpdateRequest) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *AccountUpdateRequest) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.

### HasPropertyType

`func (o *AccountUpdateRequest) HasPropertyType() bool`

HasPropertyType returns a boolean if a field has been set.

### GetSkipWebhook

`func (o *AccountUpdateRequest) GetSkipWebhook() bool`

GetSkipWebhook returns the SkipWebhook field if non-nil, zero value otherwise.

### GetSkipWebhookOk

`func (o *AccountUpdateRequest) GetSkipWebhookOk() (*bool, bool)`

GetSkipWebhookOk returns a tuple with the SkipWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhook

`func (o *AccountUpdateRequest) SetSkipWebhook(v bool)`

SetSkipWebhook sets SkipWebhook field to given value.

### HasSkipWebhook

`func (o *AccountUpdateRequest) HasSkipWebhook() bool`

HasSkipWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


