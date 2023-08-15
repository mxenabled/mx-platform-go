# AccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **NullableString** |  | [optional] 
**AccountOwnership** | Pointer to **NullableString** |  | [optional] 
**AnnuityPolicyToDate** | Pointer to **NullableString** |  | [optional] 
**AnnuityProvider** | Pointer to **NullableString** |  | [optional] 
**AnnuityTermYear** | Pointer to **NullableFloat32** |  | [optional] 
**Apr** | Pointer to **NullableFloat32** |  | [optional] 
**Apy** | Pointer to **NullableFloat32** |  | [optional] 
**AvailableBalance** | Pointer to **NullableFloat32** |  | [optional] 
**AvailableCredit** | Pointer to **NullableFloat32** |  | [optional] 
**Balance** | Pointer to **NullableFloat32** |  | [optional] 
**CashBalance** | Pointer to **NullableFloat32** |  | [optional] 
**CashSurrenderValue** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreditLimit** | Pointer to **NullableFloat32** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**DayPaymentIsDue** | Pointer to **NullableInt32** |  | [optional] 
**DeathBenefit** | Pointer to **NullableInt32** |  | [optional] 
**Guid** | Pointer to **NullableString** |  | [optional] 
**HoldingsValue** | Pointer to **NullableFloat32** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**ImportedAt** | Pointer to **NullableString** |  | [optional] 
**InterestRate** | Pointer to **NullableFloat32** |  | [optional] 
**InstitutionCode** | Pointer to **NullableString** |  | [optional] 
**InsuredName** | Pointer to **NullableString** |  | [optional] 
**IsClosed** | Pointer to **NullableBool** |  | [optional] 
**IsHidden** | Pointer to **NullableBool** |  | [optional] 
**IsManual** | Pointer to **NullableBool** |  | [optional] 
**LastPayment** | Pointer to **NullableFloat32** |  | [optional] 
**LastPaymentAt** | Pointer to **NullableString** |  | [optional] 
**LoanAmount** | Pointer to **NullableFloat32** |  | [optional] 
**MarginBalance** | Pointer to **NullableFloat32** |  | [optional] 
**MaturesOn** | Pointer to **NullableString** |  | [optional] 
**MemberGuid** | Pointer to **NullableString** |  | [optional] 
**MemberId** | Pointer to **NullableString** |  | [optional] 
**MemberIsManagedByUser** | Pointer to **NullableBool** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**MinimumBalance** | Pointer to **NullableFloat32** |  | [optional] 
**MinimumPayment** | Pointer to **NullableFloat32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Nickname** | Pointer to **NullableString** |  | [optional] 
**OriginalBalance** | Pointer to **NullableFloat32** |  | [optional] 
**PayOutAmount** | Pointer to **NullableFloat32** |  | [optional] 
**PaymentDueAt** | Pointer to **NullableString** |  | [optional] 
**PayoffBalance** | Pointer to **NullableFloat32** |  | [optional] 
**PremiumAmount** | Pointer to **NullableFloat32** |  | [optional] 
**PropertyType** | Pointer to **NullableString** |  | [optional] 
**RoutingNumber** | Pointer to **NullableString** |  | [optional] 
**StartedOn** | Pointer to **NullableString** |  | [optional] 
**Subtype** | Pointer to **NullableString** |  | [optional] 
**TodayUglAmount** | Pointer to **NullableFloat32** |  | [optional] 
**TodayUglPercentage** | Pointer to **NullableFloat32** |  | [optional] 
**TotalAccountValue** | Pointer to **NullableFloat32** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**UserGuid** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAccountResponse

`func NewAccountResponse() *AccountResponse`

NewAccountResponse instantiates a new AccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseWithDefaults

`func NewAccountResponseWithDefaults() *AccountResponse`

NewAccountResponseWithDefaults instantiates a new AccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *AccountResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountResponse) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *AccountResponse) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *AccountResponse) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetAccountOwnership

`func (o *AccountResponse) GetAccountOwnership() string`

GetAccountOwnership returns the AccountOwnership field if non-nil, zero value otherwise.

### GetAccountOwnershipOk

`func (o *AccountResponse) GetAccountOwnershipOk() (*string, bool)`

GetAccountOwnershipOk returns a tuple with the AccountOwnership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwnership

`func (o *AccountResponse) SetAccountOwnership(v string)`

SetAccountOwnership sets AccountOwnership field to given value.

### HasAccountOwnership

`func (o *AccountResponse) HasAccountOwnership() bool`

HasAccountOwnership returns a boolean if a field has been set.

### SetAccountOwnershipNil

`func (o *AccountResponse) SetAccountOwnershipNil(b bool)`

 SetAccountOwnershipNil sets the value for AccountOwnership to be an explicit nil

### UnsetAccountOwnership
`func (o *AccountResponse) UnsetAccountOwnership()`

UnsetAccountOwnership ensures that no value is present for AccountOwnership, not even an explicit nil
### GetAnnuityPolicyToDate

`func (o *AccountResponse) GetAnnuityPolicyToDate() string`

GetAnnuityPolicyToDate returns the AnnuityPolicyToDate field if non-nil, zero value otherwise.

### GetAnnuityPolicyToDateOk

`func (o *AccountResponse) GetAnnuityPolicyToDateOk() (*string, bool)`

GetAnnuityPolicyToDateOk returns a tuple with the AnnuityPolicyToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnuityPolicyToDate

`func (o *AccountResponse) SetAnnuityPolicyToDate(v string)`

SetAnnuityPolicyToDate sets AnnuityPolicyToDate field to given value.

### HasAnnuityPolicyToDate

`func (o *AccountResponse) HasAnnuityPolicyToDate() bool`

HasAnnuityPolicyToDate returns a boolean if a field has been set.

### SetAnnuityPolicyToDateNil

`func (o *AccountResponse) SetAnnuityPolicyToDateNil(b bool)`

 SetAnnuityPolicyToDateNil sets the value for AnnuityPolicyToDate to be an explicit nil

### UnsetAnnuityPolicyToDate
`func (o *AccountResponse) UnsetAnnuityPolicyToDate()`

UnsetAnnuityPolicyToDate ensures that no value is present for AnnuityPolicyToDate, not even an explicit nil
### GetAnnuityProvider

`func (o *AccountResponse) GetAnnuityProvider() string`

GetAnnuityProvider returns the AnnuityProvider field if non-nil, zero value otherwise.

### GetAnnuityProviderOk

`func (o *AccountResponse) GetAnnuityProviderOk() (*string, bool)`

GetAnnuityProviderOk returns a tuple with the AnnuityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnuityProvider

`func (o *AccountResponse) SetAnnuityProvider(v string)`

SetAnnuityProvider sets AnnuityProvider field to given value.

### HasAnnuityProvider

`func (o *AccountResponse) HasAnnuityProvider() bool`

HasAnnuityProvider returns a boolean if a field has been set.

### SetAnnuityProviderNil

`func (o *AccountResponse) SetAnnuityProviderNil(b bool)`

 SetAnnuityProviderNil sets the value for AnnuityProvider to be an explicit nil

### UnsetAnnuityProvider
`func (o *AccountResponse) UnsetAnnuityProvider()`

UnsetAnnuityProvider ensures that no value is present for AnnuityProvider, not even an explicit nil
### GetAnnuityTermYear

`func (o *AccountResponse) GetAnnuityTermYear() float32`

GetAnnuityTermYear returns the AnnuityTermYear field if non-nil, zero value otherwise.

### GetAnnuityTermYearOk

`func (o *AccountResponse) GetAnnuityTermYearOk() (*float32, bool)`

GetAnnuityTermYearOk returns a tuple with the AnnuityTermYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnuityTermYear

`func (o *AccountResponse) SetAnnuityTermYear(v float32)`

SetAnnuityTermYear sets AnnuityTermYear field to given value.

### HasAnnuityTermYear

`func (o *AccountResponse) HasAnnuityTermYear() bool`

HasAnnuityTermYear returns a boolean if a field has been set.

### SetAnnuityTermYearNil

`func (o *AccountResponse) SetAnnuityTermYearNil(b bool)`

 SetAnnuityTermYearNil sets the value for AnnuityTermYear to be an explicit nil

### UnsetAnnuityTermYear
`func (o *AccountResponse) UnsetAnnuityTermYear()`

UnsetAnnuityTermYear ensures that no value is present for AnnuityTermYear, not even an explicit nil
### GetApr

`func (o *AccountResponse) GetApr() float32`

GetApr returns the Apr field if non-nil, zero value otherwise.

### GetAprOk

`func (o *AccountResponse) GetAprOk() (*float32, bool)`

GetAprOk returns a tuple with the Apr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApr

`func (o *AccountResponse) SetApr(v float32)`

SetApr sets Apr field to given value.

### HasApr

`func (o *AccountResponse) HasApr() bool`

HasApr returns a boolean if a field has been set.

### SetAprNil

`func (o *AccountResponse) SetAprNil(b bool)`

 SetAprNil sets the value for Apr to be an explicit nil

### UnsetApr
`func (o *AccountResponse) UnsetApr()`

UnsetApr ensures that no value is present for Apr, not even an explicit nil
### GetApy

`func (o *AccountResponse) GetApy() float32`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *AccountResponse) GetApyOk() (*float32, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *AccountResponse) SetApy(v float32)`

SetApy sets Apy field to given value.

### HasApy

`func (o *AccountResponse) HasApy() bool`

HasApy returns a boolean if a field has been set.

### SetApyNil

`func (o *AccountResponse) SetApyNil(b bool)`

 SetApyNil sets the value for Apy to be an explicit nil

### UnsetApy
`func (o *AccountResponse) UnsetApy()`

UnsetApy ensures that no value is present for Apy, not even an explicit nil
### GetAvailableBalance

`func (o *AccountResponse) GetAvailableBalance() float32`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *AccountResponse) GetAvailableBalanceOk() (*float32, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *AccountResponse) SetAvailableBalance(v float32)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *AccountResponse) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### SetAvailableBalanceNil

`func (o *AccountResponse) SetAvailableBalanceNil(b bool)`

 SetAvailableBalanceNil sets the value for AvailableBalance to be an explicit nil

### UnsetAvailableBalance
`func (o *AccountResponse) UnsetAvailableBalance()`

UnsetAvailableBalance ensures that no value is present for AvailableBalance, not even an explicit nil
### GetAvailableCredit

`func (o *AccountResponse) GetAvailableCredit() float32`

GetAvailableCredit returns the AvailableCredit field if non-nil, zero value otherwise.

### GetAvailableCreditOk

`func (o *AccountResponse) GetAvailableCreditOk() (*float32, bool)`

GetAvailableCreditOk returns a tuple with the AvailableCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCredit

`func (o *AccountResponse) SetAvailableCredit(v float32)`

SetAvailableCredit sets AvailableCredit field to given value.

### HasAvailableCredit

`func (o *AccountResponse) HasAvailableCredit() bool`

HasAvailableCredit returns a boolean if a field has been set.

### SetAvailableCreditNil

`func (o *AccountResponse) SetAvailableCreditNil(b bool)`

 SetAvailableCreditNil sets the value for AvailableCredit to be an explicit nil

### UnsetAvailableCredit
`func (o *AccountResponse) UnsetAvailableCredit()`

UnsetAvailableCredit ensures that no value is present for AvailableCredit, not even an explicit nil
### GetBalance

`func (o *AccountResponse) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountResponse) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountResponse) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AccountResponse) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *AccountResponse) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *AccountResponse) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetCashBalance

`func (o *AccountResponse) GetCashBalance() float32`

GetCashBalance returns the CashBalance field if non-nil, zero value otherwise.

### GetCashBalanceOk

`func (o *AccountResponse) GetCashBalanceOk() (*float32, bool)`

GetCashBalanceOk returns a tuple with the CashBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBalance

`func (o *AccountResponse) SetCashBalance(v float32)`

SetCashBalance sets CashBalance field to given value.

### HasCashBalance

`func (o *AccountResponse) HasCashBalance() bool`

HasCashBalance returns a boolean if a field has been set.

### SetCashBalanceNil

`func (o *AccountResponse) SetCashBalanceNil(b bool)`

 SetCashBalanceNil sets the value for CashBalance to be an explicit nil

### UnsetCashBalance
`func (o *AccountResponse) UnsetCashBalance()`

UnsetCashBalance ensures that no value is present for CashBalance, not even an explicit nil
### GetCashSurrenderValue

`func (o *AccountResponse) GetCashSurrenderValue() float32`

GetCashSurrenderValue returns the CashSurrenderValue field if non-nil, zero value otherwise.

### GetCashSurrenderValueOk

`func (o *AccountResponse) GetCashSurrenderValueOk() (*float32, bool)`

GetCashSurrenderValueOk returns a tuple with the CashSurrenderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashSurrenderValue

`func (o *AccountResponse) SetCashSurrenderValue(v float32)`

SetCashSurrenderValue sets CashSurrenderValue field to given value.

### HasCashSurrenderValue

`func (o *AccountResponse) HasCashSurrenderValue() bool`

HasCashSurrenderValue returns a boolean if a field has been set.

### SetCashSurrenderValueNil

`func (o *AccountResponse) SetCashSurrenderValueNil(b bool)`

 SetCashSurrenderValueNil sets the value for CashSurrenderValue to be an explicit nil

### UnsetCashSurrenderValue
`func (o *AccountResponse) UnsetCashSurrenderValue()`

UnsetCashSurrenderValue ensures that no value is present for CashSurrenderValue, not even an explicit nil
### GetCreatedAt

`func (o *AccountResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreditLimit

`func (o *AccountResponse) GetCreditLimit() float32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *AccountResponse) GetCreditLimitOk() (*float32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *AccountResponse) SetCreditLimit(v float32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *AccountResponse) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### SetCreditLimitNil

`func (o *AccountResponse) SetCreditLimitNil(b bool)`

 SetCreditLimitNil sets the value for CreditLimit to be an explicit nil

### UnsetCreditLimit
`func (o *AccountResponse) UnsetCreditLimit()`

UnsetCreditLimit ensures that no value is present for CreditLimit, not even an explicit nil
### GetCurrencyCode

`func (o *AccountResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AccountResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *AccountResponse) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *AccountResponse) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDayPaymentIsDue

`func (o *AccountResponse) GetDayPaymentIsDue() int32`

GetDayPaymentIsDue returns the DayPaymentIsDue field if non-nil, zero value otherwise.

### GetDayPaymentIsDueOk

`func (o *AccountResponse) GetDayPaymentIsDueOk() (*int32, bool)`

GetDayPaymentIsDueOk returns a tuple with the DayPaymentIsDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayPaymentIsDue

`func (o *AccountResponse) SetDayPaymentIsDue(v int32)`

SetDayPaymentIsDue sets DayPaymentIsDue field to given value.

### HasDayPaymentIsDue

`func (o *AccountResponse) HasDayPaymentIsDue() bool`

HasDayPaymentIsDue returns a boolean if a field has been set.

### SetDayPaymentIsDueNil

`func (o *AccountResponse) SetDayPaymentIsDueNil(b bool)`

 SetDayPaymentIsDueNil sets the value for DayPaymentIsDue to be an explicit nil

### UnsetDayPaymentIsDue
`func (o *AccountResponse) UnsetDayPaymentIsDue()`

UnsetDayPaymentIsDue ensures that no value is present for DayPaymentIsDue, not even an explicit nil
### GetDeathBenefit

`func (o *AccountResponse) GetDeathBenefit() int32`

GetDeathBenefit returns the DeathBenefit field if non-nil, zero value otherwise.

### GetDeathBenefitOk

`func (o *AccountResponse) GetDeathBenefitOk() (*int32, bool)`

GetDeathBenefitOk returns a tuple with the DeathBenefit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathBenefit

`func (o *AccountResponse) SetDeathBenefit(v int32)`

SetDeathBenefit sets DeathBenefit field to given value.

### HasDeathBenefit

`func (o *AccountResponse) HasDeathBenefit() bool`

HasDeathBenefit returns a boolean if a field has been set.

### SetDeathBenefitNil

`func (o *AccountResponse) SetDeathBenefitNil(b bool)`

 SetDeathBenefitNil sets the value for DeathBenefit to be an explicit nil

### UnsetDeathBenefit
`func (o *AccountResponse) UnsetDeathBenefit()`

UnsetDeathBenefit ensures that no value is present for DeathBenefit, not even an explicit nil
### GetGuid

`func (o *AccountResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AccountResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AccountResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AccountResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *AccountResponse) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *AccountResponse) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetHoldingsValue

`func (o *AccountResponse) GetHoldingsValue() float32`

GetHoldingsValue returns the HoldingsValue field if non-nil, zero value otherwise.

### GetHoldingsValueOk

`func (o *AccountResponse) GetHoldingsValueOk() (*float32, bool)`

GetHoldingsValueOk returns a tuple with the HoldingsValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingsValue

`func (o *AccountResponse) SetHoldingsValue(v float32)`

SetHoldingsValue sets HoldingsValue field to given value.

### HasHoldingsValue

`func (o *AccountResponse) HasHoldingsValue() bool`

HasHoldingsValue returns a boolean if a field has been set.

### SetHoldingsValueNil

`func (o *AccountResponse) SetHoldingsValueNil(b bool)`

 SetHoldingsValueNil sets the value for HoldingsValue to be an explicit nil

### UnsetHoldingsValue
`func (o *AccountResponse) UnsetHoldingsValue()`

UnsetHoldingsValue ensures that no value is present for HoldingsValue, not even an explicit nil
### GetId

`func (o *AccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AccountResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AccountResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetImportedAt

`func (o *AccountResponse) GetImportedAt() string`

GetImportedAt returns the ImportedAt field if non-nil, zero value otherwise.

### GetImportedAtOk

`func (o *AccountResponse) GetImportedAtOk() (*string, bool)`

GetImportedAtOk returns a tuple with the ImportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedAt

`func (o *AccountResponse) SetImportedAt(v string)`

SetImportedAt sets ImportedAt field to given value.

### HasImportedAt

`func (o *AccountResponse) HasImportedAt() bool`

HasImportedAt returns a boolean if a field has been set.

### SetImportedAtNil

`func (o *AccountResponse) SetImportedAtNil(b bool)`

 SetImportedAtNil sets the value for ImportedAt to be an explicit nil

### UnsetImportedAt
`func (o *AccountResponse) UnsetImportedAt()`

UnsetImportedAt ensures that no value is present for ImportedAt, not even an explicit nil
### GetInterestRate

`func (o *AccountResponse) GetInterestRate() float32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *AccountResponse) GetInterestRateOk() (*float32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *AccountResponse) SetInterestRate(v float32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *AccountResponse) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### SetInterestRateNil

`func (o *AccountResponse) SetInterestRateNil(b bool)`

 SetInterestRateNil sets the value for InterestRate to be an explicit nil

### UnsetInterestRate
`func (o *AccountResponse) UnsetInterestRate()`

UnsetInterestRate ensures that no value is present for InterestRate, not even an explicit nil
### GetInstitutionCode

`func (o *AccountResponse) GetInstitutionCode() string`

GetInstitutionCode returns the InstitutionCode field if non-nil, zero value otherwise.

### GetInstitutionCodeOk

`func (o *AccountResponse) GetInstitutionCodeOk() (*string, bool)`

GetInstitutionCodeOk returns a tuple with the InstitutionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionCode

`func (o *AccountResponse) SetInstitutionCode(v string)`

SetInstitutionCode sets InstitutionCode field to given value.

### HasInstitutionCode

`func (o *AccountResponse) HasInstitutionCode() bool`

HasInstitutionCode returns a boolean if a field has been set.

### SetInstitutionCodeNil

`func (o *AccountResponse) SetInstitutionCodeNil(b bool)`

 SetInstitutionCodeNil sets the value for InstitutionCode to be an explicit nil

### UnsetInstitutionCode
`func (o *AccountResponse) UnsetInstitutionCode()`

UnsetInstitutionCode ensures that no value is present for InstitutionCode, not even an explicit nil
### GetInsuredName

`func (o *AccountResponse) GetInsuredName() string`

GetInsuredName returns the InsuredName field if non-nil, zero value otherwise.

### GetInsuredNameOk

`func (o *AccountResponse) GetInsuredNameOk() (*string, bool)`

GetInsuredNameOk returns a tuple with the InsuredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuredName

`func (o *AccountResponse) SetInsuredName(v string)`

SetInsuredName sets InsuredName field to given value.

### HasInsuredName

`func (o *AccountResponse) HasInsuredName() bool`

HasInsuredName returns a boolean if a field has been set.

### SetInsuredNameNil

`func (o *AccountResponse) SetInsuredNameNil(b bool)`

 SetInsuredNameNil sets the value for InsuredName to be an explicit nil

### UnsetInsuredName
`func (o *AccountResponse) UnsetInsuredName()`

UnsetInsuredName ensures that no value is present for InsuredName, not even an explicit nil
### GetIsClosed

`func (o *AccountResponse) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *AccountResponse) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *AccountResponse) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *AccountResponse) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### SetIsClosedNil

`func (o *AccountResponse) SetIsClosedNil(b bool)`

 SetIsClosedNil sets the value for IsClosed to be an explicit nil

### UnsetIsClosed
`func (o *AccountResponse) UnsetIsClosed()`

UnsetIsClosed ensures that no value is present for IsClosed, not even an explicit nil
### GetIsHidden

`func (o *AccountResponse) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *AccountResponse) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *AccountResponse) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *AccountResponse) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### SetIsHiddenNil

`func (o *AccountResponse) SetIsHiddenNil(b bool)`

 SetIsHiddenNil sets the value for IsHidden to be an explicit nil

### UnsetIsHidden
`func (o *AccountResponse) UnsetIsHidden()`

UnsetIsHidden ensures that no value is present for IsHidden, not even an explicit nil
### GetIsManual

`func (o *AccountResponse) GetIsManual() bool`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *AccountResponse) GetIsManualOk() (*bool, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *AccountResponse) SetIsManual(v bool)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *AccountResponse) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### SetIsManualNil

`func (o *AccountResponse) SetIsManualNil(b bool)`

 SetIsManualNil sets the value for IsManual to be an explicit nil

### UnsetIsManual
`func (o *AccountResponse) UnsetIsManual()`

UnsetIsManual ensures that no value is present for IsManual, not even an explicit nil
### GetLastPayment

`func (o *AccountResponse) GetLastPayment() float32`

GetLastPayment returns the LastPayment field if non-nil, zero value otherwise.

### GetLastPaymentOk

`func (o *AccountResponse) GetLastPaymentOk() (*float32, bool)`

GetLastPaymentOk returns a tuple with the LastPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPayment

`func (o *AccountResponse) SetLastPayment(v float32)`

SetLastPayment sets LastPayment field to given value.

### HasLastPayment

`func (o *AccountResponse) HasLastPayment() bool`

HasLastPayment returns a boolean if a field has been set.

### SetLastPaymentNil

`func (o *AccountResponse) SetLastPaymentNil(b bool)`

 SetLastPaymentNil sets the value for LastPayment to be an explicit nil

### UnsetLastPayment
`func (o *AccountResponse) UnsetLastPayment()`

UnsetLastPayment ensures that no value is present for LastPayment, not even an explicit nil
### GetLastPaymentAt

`func (o *AccountResponse) GetLastPaymentAt() string`

GetLastPaymentAt returns the LastPaymentAt field if non-nil, zero value otherwise.

### GetLastPaymentAtOk

`func (o *AccountResponse) GetLastPaymentAtOk() (*string, bool)`

GetLastPaymentAtOk returns a tuple with the LastPaymentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAt

`func (o *AccountResponse) SetLastPaymentAt(v string)`

SetLastPaymentAt sets LastPaymentAt field to given value.

### HasLastPaymentAt

`func (o *AccountResponse) HasLastPaymentAt() bool`

HasLastPaymentAt returns a boolean if a field has been set.

### SetLastPaymentAtNil

`func (o *AccountResponse) SetLastPaymentAtNil(b bool)`

 SetLastPaymentAtNil sets the value for LastPaymentAt to be an explicit nil

### UnsetLastPaymentAt
`func (o *AccountResponse) UnsetLastPaymentAt()`

UnsetLastPaymentAt ensures that no value is present for LastPaymentAt, not even an explicit nil
### GetLoanAmount

`func (o *AccountResponse) GetLoanAmount() float32`

GetLoanAmount returns the LoanAmount field if non-nil, zero value otherwise.

### GetLoanAmountOk

`func (o *AccountResponse) GetLoanAmountOk() (*float32, bool)`

GetLoanAmountOk returns a tuple with the LoanAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanAmount

`func (o *AccountResponse) SetLoanAmount(v float32)`

SetLoanAmount sets LoanAmount field to given value.

### HasLoanAmount

`func (o *AccountResponse) HasLoanAmount() bool`

HasLoanAmount returns a boolean if a field has been set.

### SetLoanAmountNil

`func (o *AccountResponse) SetLoanAmountNil(b bool)`

 SetLoanAmountNil sets the value for LoanAmount to be an explicit nil

### UnsetLoanAmount
`func (o *AccountResponse) UnsetLoanAmount()`

UnsetLoanAmount ensures that no value is present for LoanAmount, not even an explicit nil
### GetMarginBalance

`func (o *AccountResponse) GetMarginBalance() float32`

GetMarginBalance returns the MarginBalance field if non-nil, zero value otherwise.

### GetMarginBalanceOk

`func (o *AccountResponse) GetMarginBalanceOk() (*float32, bool)`

GetMarginBalanceOk returns a tuple with the MarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBalance

`func (o *AccountResponse) SetMarginBalance(v float32)`

SetMarginBalance sets MarginBalance field to given value.

### HasMarginBalance

`func (o *AccountResponse) HasMarginBalance() bool`

HasMarginBalance returns a boolean if a field has been set.

### SetMarginBalanceNil

`func (o *AccountResponse) SetMarginBalanceNil(b bool)`

 SetMarginBalanceNil sets the value for MarginBalance to be an explicit nil

### UnsetMarginBalance
`func (o *AccountResponse) UnsetMarginBalance()`

UnsetMarginBalance ensures that no value is present for MarginBalance, not even an explicit nil
### GetMaturesOn

`func (o *AccountResponse) GetMaturesOn() string`

GetMaturesOn returns the MaturesOn field if non-nil, zero value otherwise.

### GetMaturesOnOk

`func (o *AccountResponse) GetMaturesOnOk() (*string, bool)`

GetMaturesOnOk returns a tuple with the MaturesOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturesOn

`func (o *AccountResponse) SetMaturesOn(v string)`

SetMaturesOn sets MaturesOn field to given value.

### HasMaturesOn

`func (o *AccountResponse) HasMaturesOn() bool`

HasMaturesOn returns a boolean if a field has been set.

### SetMaturesOnNil

`func (o *AccountResponse) SetMaturesOnNil(b bool)`

 SetMaturesOnNil sets the value for MaturesOn to be an explicit nil

### UnsetMaturesOn
`func (o *AccountResponse) UnsetMaturesOn()`

UnsetMaturesOn ensures that no value is present for MaturesOn, not even an explicit nil
### GetMemberGuid

`func (o *AccountResponse) GetMemberGuid() string`

GetMemberGuid returns the MemberGuid field if non-nil, zero value otherwise.

### GetMemberGuidOk

`func (o *AccountResponse) GetMemberGuidOk() (*string, bool)`

GetMemberGuidOk returns a tuple with the MemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGuid

`func (o *AccountResponse) SetMemberGuid(v string)`

SetMemberGuid sets MemberGuid field to given value.

### HasMemberGuid

`func (o *AccountResponse) HasMemberGuid() bool`

HasMemberGuid returns a boolean if a field has been set.

### SetMemberGuidNil

`func (o *AccountResponse) SetMemberGuidNil(b bool)`

 SetMemberGuidNil sets the value for MemberGuid to be an explicit nil

### UnsetMemberGuid
`func (o *AccountResponse) UnsetMemberGuid()`

UnsetMemberGuid ensures that no value is present for MemberGuid, not even an explicit nil
### GetMemberId

`func (o *AccountResponse) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *AccountResponse) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *AccountResponse) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *AccountResponse) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### SetMemberIdNil

`func (o *AccountResponse) SetMemberIdNil(b bool)`

 SetMemberIdNil sets the value for MemberId to be an explicit nil

### UnsetMemberId
`func (o *AccountResponse) UnsetMemberId()`

UnsetMemberId ensures that no value is present for MemberId, not even an explicit nil
### GetMemberIsManagedByUser

`func (o *AccountResponse) GetMemberIsManagedByUser() bool`

GetMemberIsManagedByUser returns the MemberIsManagedByUser field if non-nil, zero value otherwise.

### GetMemberIsManagedByUserOk

`func (o *AccountResponse) GetMemberIsManagedByUserOk() (*bool, bool)`

GetMemberIsManagedByUserOk returns a tuple with the MemberIsManagedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIsManagedByUser

`func (o *AccountResponse) SetMemberIsManagedByUser(v bool)`

SetMemberIsManagedByUser sets MemberIsManagedByUser field to given value.

### HasMemberIsManagedByUser

`func (o *AccountResponse) HasMemberIsManagedByUser() bool`

HasMemberIsManagedByUser returns a boolean if a field has been set.

### SetMemberIsManagedByUserNil

`func (o *AccountResponse) SetMemberIsManagedByUserNil(b bool)`

 SetMemberIsManagedByUserNil sets the value for MemberIsManagedByUser to be an explicit nil

### UnsetMemberIsManagedByUser
`func (o *AccountResponse) UnsetMemberIsManagedByUser()`

UnsetMemberIsManagedByUser ensures that no value is present for MemberIsManagedByUser, not even an explicit nil
### GetMetadata

`func (o *AccountResponse) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountResponse) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountResponse) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AccountResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AccountResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMinimumBalance

`func (o *AccountResponse) GetMinimumBalance() float32`

GetMinimumBalance returns the MinimumBalance field if non-nil, zero value otherwise.

### GetMinimumBalanceOk

`func (o *AccountResponse) GetMinimumBalanceOk() (*float32, bool)`

GetMinimumBalanceOk returns a tuple with the MinimumBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBalance

`func (o *AccountResponse) SetMinimumBalance(v float32)`

SetMinimumBalance sets MinimumBalance field to given value.

### HasMinimumBalance

`func (o *AccountResponse) HasMinimumBalance() bool`

HasMinimumBalance returns a boolean if a field has been set.

### SetMinimumBalanceNil

`func (o *AccountResponse) SetMinimumBalanceNil(b bool)`

 SetMinimumBalanceNil sets the value for MinimumBalance to be an explicit nil

### UnsetMinimumBalance
`func (o *AccountResponse) UnsetMinimumBalance()`

UnsetMinimumBalance ensures that no value is present for MinimumBalance, not even an explicit nil
### GetMinimumPayment

`func (o *AccountResponse) GetMinimumPayment() float32`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *AccountResponse) GetMinimumPaymentOk() (*float32, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *AccountResponse) SetMinimumPayment(v float32)`

SetMinimumPayment sets MinimumPayment field to given value.

### HasMinimumPayment

`func (o *AccountResponse) HasMinimumPayment() bool`

HasMinimumPayment returns a boolean if a field has been set.

### SetMinimumPaymentNil

`func (o *AccountResponse) SetMinimumPaymentNil(b bool)`

 SetMinimumPaymentNil sets the value for MinimumPayment to be an explicit nil

### UnsetMinimumPayment
`func (o *AccountResponse) UnsetMinimumPayment()`

UnsetMinimumPayment ensures that no value is present for MinimumPayment, not even an explicit nil
### GetName

`func (o *AccountResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AccountResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AccountResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNickname

`func (o *AccountResponse) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountResponse) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountResponse) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountResponse) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### SetNicknameNil

`func (o *AccountResponse) SetNicknameNil(b bool)`

 SetNicknameNil sets the value for Nickname to be an explicit nil

### UnsetNickname
`func (o *AccountResponse) UnsetNickname()`

UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
### GetOriginalBalance

`func (o *AccountResponse) GetOriginalBalance() float32`

GetOriginalBalance returns the OriginalBalance field if non-nil, zero value otherwise.

### GetOriginalBalanceOk

`func (o *AccountResponse) GetOriginalBalanceOk() (*float32, bool)`

GetOriginalBalanceOk returns a tuple with the OriginalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalBalance

`func (o *AccountResponse) SetOriginalBalance(v float32)`

SetOriginalBalance sets OriginalBalance field to given value.

### HasOriginalBalance

`func (o *AccountResponse) HasOriginalBalance() bool`

HasOriginalBalance returns a boolean if a field has been set.

### SetOriginalBalanceNil

`func (o *AccountResponse) SetOriginalBalanceNil(b bool)`

 SetOriginalBalanceNil sets the value for OriginalBalance to be an explicit nil

### UnsetOriginalBalance
`func (o *AccountResponse) UnsetOriginalBalance()`

UnsetOriginalBalance ensures that no value is present for OriginalBalance, not even an explicit nil
### GetPayOutAmount

`func (o *AccountResponse) GetPayOutAmount() float32`

GetPayOutAmount returns the PayOutAmount field if non-nil, zero value otherwise.

### GetPayOutAmountOk

`func (o *AccountResponse) GetPayOutAmountOk() (*float32, bool)`

GetPayOutAmountOk returns a tuple with the PayOutAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOutAmount

`func (o *AccountResponse) SetPayOutAmount(v float32)`

SetPayOutAmount sets PayOutAmount field to given value.

### HasPayOutAmount

`func (o *AccountResponse) HasPayOutAmount() bool`

HasPayOutAmount returns a boolean if a field has been set.

### SetPayOutAmountNil

`func (o *AccountResponse) SetPayOutAmountNil(b bool)`

 SetPayOutAmountNil sets the value for PayOutAmount to be an explicit nil

### UnsetPayOutAmount
`func (o *AccountResponse) UnsetPayOutAmount()`

UnsetPayOutAmount ensures that no value is present for PayOutAmount, not even an explicit nil
### GetPaymentDueAt

`func (o *AccountResponse) GetPaymentDueAt() string`

GetPaymentDueAt returns the PaymentDueAt field if non-nil, zero value otherwise.

### GetPaymentDueAtOk

`func (o *AccountResponse) GetPaymentDueAtOk() (*string, bool)`

GetPaymentDueAtOk returns a tuple with the PaymentDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueAt

`func (o *AccountResponse) SetPaymentDueAt(v string)`

SetPaymentDueAt sets PaymentDueAt field to given value.

### HasPaymentDueAt

`func (o *AccountResponse) HasPaymentDueAt() bool`

HasPaymentDueAt returns a boolean if a field has been set.

### SetPaymentDueAtNil

`func (o *AccountResponse) SetPaymentDueAtNil(b bool)`

 SetPaymentDueAtNil sets the value for PaymentDueAt to be an explicit nil

### UnsetPaymentDueAt
`func (o *AccountResponse) UnsetPaymentDueAt()`

UnsetPaymentDueAt ensures that no value is present for PaymentDueAt, not even an explicit nil
### GetPayoffBalance

`func (o *AccountResponse) GetPayoffBalance() float32`

GetPayoffBalance returns the PayoffBalance field if non-nil, zero value otherwise.

### GetPayoffBalanceOk

`func (o *AccountResponse) GetPayoffBalanceOk() (*float32, bool)`

GetPayoffBalanceOk returns a tuple with the PayoffBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoffBalance

`func (o *AccountResponse) SetPayoffBalance(v float32)`

SetPayoffBalance sets PayoffBalance field to given value.

### HasPayoffBalance

`func (o *AccountResponse) HasPayoffBalance() bool`

HasPayoffBalance returns a boolean if a field has been set.

### SetPayoffBalanceNil

`func (o *AccountResponse) SetPayoffBalanceNil(b bool)`

 SetPayoffBalanceNil sets the value for PayoffBalance to be an explicit nil

### UnsetPayoffBalance
`func (o *AccountResponse) UnsetPayoffBalance()`

UnsetPayoffBalance ensures that no value is present for PayoffBalance, not even an explicit nil
### GetPremiumAmount

`func (o *AccountResponse) GetPremiumAmount() float32`

GetPremiumAmount returns the PremiumAmount field if non-nil, zero value otherwise.

### GetPremiumAmountOk

`func (o *AccountResponse) GetPremiumAmountOk() (*float32, bool)`

GetPremiumAmountOk returns a tuple with the PremiumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumAmount

`func (o *AccountResponse) SetPremiumAmount(v float32)`

SetPremiumAmount sets PremiumAmount field to given value.

### HasPremiumAmount

`func (o *AccountResponse) HasPremiumAmount() bool`

HasPremiumAmount returns a boolean if a field has been set.

### SetPremiumAmountNil

`func (o *AccountResponse) SetPremiumAmountNil(b bool)`

 SetPremiumAmountNil sets the value for PremiumAmount to be an explicit nil

### UnsetPremiumAmount
`func (o *AccountResponse) UnsetPremiumAmount()`

UnsetPremiumAmount ensures that no value is present for PremiumAmount, not even an explicit nil
### GetPropertyType

`func (o *AccountResponse) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *AccountResponse) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *AccountResponse) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.

### HasPropertyType

`func (o *AccountResponse) HasPropertyType() bool`

HasPropertyType returns a boolean if a field has been set.

### SetPropertyTypeNil

`func (o *AccountResponse) SetPropertyTypeNil(b bool)`

 SetPropertyTypeNil sets the value for PropertyType to be an explicit nil

### UnsetPropertyType
`func (o *AccountResponse) UnsetPropertyType()`

UnsetPropertyType ensures that no value is present for PropertyType, not even an explicit nil
### GetRoutingNumber

`func (o *AccountResponse) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *AccountResponse) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *AccountResponse) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *AccountResponse) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### SetRoutingNumberNil

`func (o *AccountResponse) SetRoutingNumberNil(b bool)`

 SetRoutingNumberNil sets the value for RoutingNumber to be an explicit nil

### UnsetRoutingNumber
`func (o *AccountResponse) UnsetRoutingNumber()`

UnsetRoutingNumber ensures that no value is present for RoutingNumber, not even an explicit nil
### GetStartedOn

`func (o *AccountResponse) GetStartedOn() string`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *AccountResponse) GetStartedOnOk() (*string, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *AccountResponse) SetStartedOn(v string)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *AccountResponse) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *AccountResponse) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *AccountResponse) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetSubtype

`func (o *AccountResponse) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AccountResponse) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AccountResponse) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *AccountResponse) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *AccountResponse) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *AccountResponse) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetTodayUglAmount

`func (o *AccountResponse) GetTodayUglAmount() float32`

GetTodayUglAmount returns the TodayUglAmount field if non-nil, zero value otherwise.

### GetTodayUglAmountOk

`func (o *AccountResponse) GetTodayUglAmountOk() (*float32, bool)`

GetTodayUglAmountOk returns a tuple with the TodayUglAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodayUglAmount

`func (o *AccountResponse) SetTodayUglAmount(v float32)`

SetTodayUglAmount sets TodayUglAmount field to given value.

### HasTodayUglAmount

`func (o *AccountResponse) HasTodayUglAmount() bool`

HasTodayUglAmount returns a boolean if a field has been set.

### SetTodayUglAmountNil

`func (o *AccountResponse) SetTodayUglAmountNil(b bool)`

 SetTodayUglAmountNil sets the value for TodayUglAmount to be an explicit nil

### UnsetTodayUglAmount
`func (o *AccountResponse) UnsetTodayUglAmount()`

UnsetTodayUglAmount ensures that no value is present for TodayUglAmount, not even an explicit nil
### GetTodayUglPercentage

`func (o *AccountResponse) GetTodayUglPercentage() float32`

GetTodayUglPercentage returns the TodayUglPercentage field if non-nil, zero value otherwise.

### GetTodayUglPercentageOk

`func (o *AccountResponse) GetTodayUglPercentageOk() (*float32, bool)`

GetTodayUglPercentageOk returns a tuple with the TodayUglPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodayUglPercentage

`func (o *AccountResponse) SetTodayUglPercentage(v float32)`

SetTodayUglPercentage sets TodayUglPercentage field to given value.

### HasTodayUglPercentage

`func (o *AccountResponse) HasTodayUglPercentage() bool`

HasTodayUglPercentage returns a boolean if a field has been set.

### SetTodayUglPercentageNil

`func (o *AccountResponse) SetTodayUglPercentageNil(b bool)`

 SetTodayUglPercentageNil sets the value for TodayUglPercentage to be an explicit nil

### UnsetTodayUglPercentage
`func (o *AccountResponse) UnsetTodayUglPercentage()`

UnsetTodayUglPercentage ensures that no value is present for TodayUglPercentage, not even an explicit nil
### GetTotalAccountValue

`func (o *AccountResponse) GetTotalAccountValue() float32`

GetTotalAccountValue returns the TotalAccountValue field if non-nil, zero value otherwise.

### GetTotalAccountValueOk

`func (o *AccountResponse) GetTotalAccountValueOk() (*float32, bool)`

GetTotalAccountValueOk returns a tuple with the TotalAccountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAccountValue

`func (o *AccountResponse) SetTotalAccountValue(v float32)`

SetTotalAccountValue sets TotalAccountValue field to given value.

### HasTotalAccountValue

`func (o *AccountResponse) HasTotalAccountValue() bool`

HasTotalAccountValue returns a boolean if a field has been set.

### SetTotalAccountValueNil

`func (o *AccountResponse) SetTotalAccountValueNil(b bool)`

 SetTotalAccountValueNil sets the value for TotalAccountValue to be an explicit nil

### UnsetTotalAccountValue
`func (o *AccountResponse) UnsetTotalAccountValue()`

UnsetTotalAccountValue ensures that no value is present for TotalAccountValue, not even an explicit nil
### GetType

`func (o *AccountResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AccountResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AccountResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUpdatedAt

`func (o *AccountResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *AccountResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *AccountResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserGuid

`func (o *AccountResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *AccountResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *AccountResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *AccountResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### SetUserGuidNil

`func (o *AccountResponse) SetUserGuidNil(b bool)`

 SetUserGuidNil sets the value for UserGuid to be an explicit nil

### UnsetUserGuid
`func (o *AccountResponse) UnsetUserGuid()`

UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
### GetUserId

`func (o *AccountResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AccountResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AccountResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AccountResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *AccountResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AccountResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


