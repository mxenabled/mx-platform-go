# AccountCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSubtypeName** | Pointer to **string** |  | [optional] 
**AccountType** | **int32** |  | 
**Apr** | Pointer to **float32** |  | [optional] 
**Apy** | Pointer to **float32** |  | [optional] 
**AvailableBalance** | Pointer to **float32** |  | [optional] 
**Balance** | Pointer to **float32** |  | [optional] 
**CashSurrenderValue** | Pointer to **float32** |  | [optional] 
**CreditLimit** | Pointer to **float32** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**DeathBenefit** | Pointer to **int32** |  | [optional] 
**InterestRate** | Pointer to **float32** |  | [optional] 
**IsBusiness** | Pointer to **bool** |  | [optional] 
**IsClosed** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**LoanAmount** | Pointer to **float32** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Nickname** | Pointer to **string** |  | [optional] 
**OriginalBalance** | Pointer to **float32** |  | [optional] 
**PropertyType** | Pointer to **int32** |  | [optional] 
**PropertyTypeName** | Pointer to **string** |  | [optional] 
**SkipWebhook** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccountCreateRequest

`func NewAccountCreateRequest(accountType int32, name string, ) *AccountCreateRequest`

NewAccountCreateRequest instantiates a new AccountCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreateRequestWithDefaults

`func NewAccountCreateRequestWithDefaults() *AccountCreateRequest`

NewAccountCreateRequestWithDefaults instantiates a new AccountCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSubtypeName

`func (o *AccountCreateRequest) GetAccountSubtypeName() string`

GetAccountSubtypeName returns the AccountSubtypeName field if non-nil, zero value otherwise.

### GetAccountSubtypeNameOk

`func (o *AccountCreateRequest) GetAccountSubtypeNameOk() (*string, bool)`

GetAccountSubtypeNameOk returns a tuple with the AccountSubtypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSubtypeName

`func (o *AccountCreateRequest) SetAccountSubtypeName(v string)`

SetAccountSubtypeName sets AccountSubtypeName field to given value.

### HasAccountSubtypeName

`func (o *AccountCreateRequest) HasAccountSubtypeName() bool`

HasAccountSubtypeName returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountCreateRequest) GetAccountType() int32`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountCreateRequest) GetAccountTypeOk() (*int32, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountCreateRequest) SetAccountType(v int32)`

SetAccountType sets AccountType field to given value.


### GetApr

`func (o *AccountCreateRequest) GetApr() float32`

GetApr returns the Apr field if non-nil, zero value otherwise.

### GetAprOk

`func (o *AccountCreateRequest) GetAprOk() (*float32, bool)`

GetAprOk returns a tuple with the Apr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApr

`func (o *AccountCreateRequest) SetApr(v float32)`

SetApr sets Apr field to given value.

### HasApr

`func (o *AccountCreateRequest) HasApr() bool`

HasApr returns a boolean if a field has been set.

### GetApy

`func (o *AccountCreateRequest) GetApy() float32`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *AccountCreateRequest) GetApyOk() (*float32, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *AccountCreateRequest) SetApy(v float32)`

SetApy sets Apy field to given value.

### HasApy

`func (o *AccountCreateRequest) HasApy() bool`

HasApy returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *AccountCreateRequest) GetAvailableBalance() float32`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *AccountCreateRequest) GetAvailableBalanceOk() (*float32, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *AccountCreateRequest) SetAvailableBalance(v float32)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *AccountCreateRequest) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetBalance

`func (o *AccountCreateRequest) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountCreateRequest) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountCreateRequest) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AccountCreateRequest) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCashSurrenderValue

`func (o *AccountCreateRequest) GetCashSurrenderValue() float32`

GetCashSurrenderValue returns the CashSurrenderValue field if non-nil, zero value otherwise.

### GetCashSurrenderValueOk

`func (o *AccountCreateRequest) GetCashSurrenderValueOk() (*float32, bool)`

GetCashSurrenderValueOk returns a tuple with the CashSurrenderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashSurrenderValue

`func (o *AccountCreateRequest) SetCashSurrenderValue(v float32)`

SetCashSurrenderValue sets CashSurrenderValue field to given value.

### HasCashSurrenderValue

`func (o *AccountCreateRequest) HasCashSurrenderValue() bool`

HasCashSurrenderValue returns a boolean if a field has been set.

### GetCreditLimit

`func (o *AccountCreateRequest) GetCreditLimit() float32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *AccountCreateRequest) GetCreditLimitOk() (*float32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *AccountCreateRequest) SetCreditLimit(v float32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *AccountCreateRequest) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AccountCreateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountCreateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountCreateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AccountCreateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDeathBenefit

`func (o *AccountCreateRequest) GetDeathBenefit() int32`

GetDeathBenefit returns the DeathBenefit field if non-nil, zero value otherwise.

### GetDeathBenefitOk

`func (o *AccountCreateRequest) GetDeathBenefitOk() (*int32, bool)`

GetDeathBenefitOk returns a tuple with the DeathBenefit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathBenefit

`func (o *AccountCreateRequest) SetDeathBenefit(v int32)`

SetDeathBenefit sets DeathBenefit field to given value.

### HasDeathBenefit

`func (o *AccountCreateRequest) HasDeathBenefit() bool`

HasDeathBenefit returns a boolean if a field has been set.

### GetInterestRate

`func (o *AccountCreateRequest) GetInterestRate() float32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *AccountCreateRequest) GetInterestRateOk() (*float32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *AccountCreateRequest) SetInterestRate(v float32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *AccountCreateRequest) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetIsBusiness

`func (o *AccountCreateRequest) GetIsBusiness() bool`

GetIsBusiness returns the IsBusiness field if non-nil, zero value otherwise.

### GetIsBusinessOk

`func (o *AccountCreateRequest) GetIsBusinessOk() (*bool, bool)`

GetIsBusinessOk returns a tuple with the IsBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBusiness

`func (o *AccountCreateRequest) SetIsBusiness(v bool)`

SetIsBusiness sets IsBusiness field to given value.

### HasIsBusiness

`func (o *AccountCreateRequest) HasIsBusiness() bool`

HasIsBusiness returns a boolean if a field has been set.

### GetIsClosed

`func (o *AccountCreateRequest) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *AccountCreateRequest) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *AccountCreateRequest) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *AccountCreateRequest) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetIsHidden

`func (o *AccountCreateRequest) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *AccountCreateRequest) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *AccountCreateRequest) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *AccountCreateRequest) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetLoanAmount

`func (o *AccountCreateRequest) GetLoanAmount() float32`

GetLoanAmount returns the LoanAmount field if non-nil, zero value otherwise.

### GetLoanAmountOk

`func (o *AccountCreateRequest) GetLoanAmountOk() (*float32, bool)`

GetLoanAmountOk returns a tuple with the LoanAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanAmount

`func (o *AccountCreateRequest) SetLoanAmount(v float32)`

SetLoanAmount sets LoanAmount field to given value.

### HasLoanAmount

`func (o *AccountCreateRequest) HasLoanAmount() bool`

HasLoanAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountCreateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountCreateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountCreateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *AccountCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNickname

`func (o *AccountCreateRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountCreateRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountCreateRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountCreateRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetOriginalBalance

`func (o *AccountCreateRequest) GetOriginalBalance() float32`

GetOriginalBalance returns the OriginalBalance field if non-nil, zero value otherwise.

### GetOriginalBalanceOk

`func (o *AccountCreateRequest) GetOriginalBalanceOk() (*float32, bool)`

GetOriginalBalanceOk returns a tuple with the OriginalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalBalance

`func (o *AccountCreateRequest) SetOriginalBalance(v float32)`

SetOriginalBalance sets OriginalBalance field to given value.

### HasOriginalBalance

`func (o *AccountCreateRequest) HasOriginalBalance() bool`

HasOriginalBalance returns a boolean if a field has been set.

### GetPropertyType

`func (o *AccountCreateRequest) GetPropertyType() int32`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *AccountCreateRequest) GetPropertyTypeOk() (*int32, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *AccountCreateRequest) SetPropertyType(v int32)`

SetPropertyType sets PropertyType field to given value.

### HasPropertyType

`func (o *AccountCreateRequest) HasPropertyType() bool`

HasPropertyType returns a boolean if a field has been set.

### GetPropertyTypeName

`func (o *AccountCreateRequest) GetPropertyTypeName() string`

GetPropertyTypeName returns the PropertyTypeName field if non-nil, zero value otherwise.

### GetPropertyTypeNameOk

`func (o *AccountCreateRequest) GetPropertyTypeNameOk() (*string, bool)`

GetPropertyTypeNameOk returns a tuple with the PropertyTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyTypeName

`func (o *AccountCreateRequest) SetPropertyTypeName(v string)`

SetPropertyTypeName sets PropertyTypeName field to given value.

### HasPropertyTypeName

`func (o *AccountCreateRequest) HasPropertyTypeName() bool`

HasPropertyTypeName returns a boolean if a field has been set.

### GetSkipWebhook

`func (o *AccountCreateRequest) GetSkipWebhook() bool`

GetSkipWebhook returns the SkipWebhook field if non-nil, zero value otherwise.

### GetSkipWebhookOk

`func (o *AccountCreateRequest) GetSkipWebhookOk() (*bool, bool)`

GetSkipWebhookOk returns a tuple with the SkipWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhook

`func (o *AccountCreateRequest) SetSkipWebhook(v bool)`

SetSkipWebhook sets SkipWebhook field to given value.

### HasSkipWebhook

`func (o *AccountCreateRequest) HasSkipWebhook() bool`

HasSkipWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


