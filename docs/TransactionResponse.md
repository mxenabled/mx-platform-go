# TransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **NullableFloat32** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**CategoryGuid** | Pointer to **NullableString** |  | [optional] 
**CheckNumberString** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**Date** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExtendedTransactionType** | Pointer to **NullableString** |  | [optional] 
**Guid** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**IsBillPay** | Pointer to **NullableBool** |  | [optional] 
**IsDirectDeposit** | Pointer to **NullableBool** |  | [optional] 
**IsExpense** | Pointer to **NullableBool** |  | [optional] 
**IsFee** | Pointer to **NullableBool** |  | [optional] 
**IsIncome** | Pointer to **NullableBool** |  | [optional] 
**IsInternational** | Pointer to **NullableBool** |  | [optional] 
**IsManual** | Pointer to **NullableBool** |  | [optional] 
**IsOverdraftFee** | Pointer to **NullableBool** |  | [optional] 
**IsPayrollAdvance** | Pointer to **NullableBool** |  | [optional] 
**IsRecurring** | Pointer to **NullableBool** |  | [optional] 
**IsSubscription** | Pointer to **NullableBool** |  | [optional] 
**Latitude** | Pointer to **NullableFloat32** |  | [optional] 
**LocalizedDescription** | Pointer to **NullableString** |  | [optional] 
**LocalizedMemo** | Pointer to **NullableString** |  | [optional] 
**Longitude** | Pointer to **NullableFloat32** |  | [optional] 
**MemberGuid** | Pointer to **NullableString** |  | [optional] 
**MemberIsManagedByUser** | Pointer to **NullableBool** |  | [optional] 
**Memo** | Pointer to **NullableString** |  | [optional] 
**MerchantCategoryCode** | Pointer to **NullableInt32** |  | [optional] 
**MerchantGuid** | Pointer to **NullableString** |  | [optional] 
**MerchantLocationGuid** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**OriginalDescription** | Pointer to **NullableString** |  | [optional] 
**PostedAt** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**TopLevelCategory** | Pointer to **NullableString** |  | [optional] 
**TransactedAt** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**UserGuid** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTransactionResponse

`func NewTransactionResponse() *TransactionResponse`

NewTransactionResponse instantiates a new TransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResponseWithDefaults

`func NewTransactionResponseWithDefaults() *TransactionResponse`

NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *TransactionResponse) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *TransactionResponse) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *TransactionResponse) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.

### HasAccountGuid

`func (o *TransactionResponse) HasAccountGuid() bool`

HasAccountGuid returns a boolean if a field has been set.

### SetAccountGuidNil

`func (o *TransactionResponse) SetAccountGuidNil(b bool)`

 SetAccountGuidNil sets the value for AccountGuid to be an explicit nil

### UnsetAccountGuid
`func (o *TransactionResponse) UnsetAccountGuid()`

UnsetAccountGuid ensures that no value is present for AccountGuid, not even an explicit nil
### GetAccountId

`func (o *TransactionResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransactionResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransactionResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TransactionResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *TransactionResponse) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *TransactionResponse) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAmount

`func (o *TransactionResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *TransactionResponse) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *TransactionResponse) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCategory

`func (o *TransactionResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TransactionResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TransactionResponse) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TransactionResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *TransactionResponse) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *TransactionResponse) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCategoryGuid

`func (o *TransactionResponse) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *TransactionResponse) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *TransactionResponse) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *TransactionResponse) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### SetCategoryGuidNil

`func (o *TransactionResponse) SetCategoryGuidNil(b bool)`

 SetCategoryGuidNil sets the value for CategoryGuid to be an explicit nil

### UnsetCategoryGuid
`func (o *TransactionResponse) UnsetCategoryGuid()`

UnsetCategoryGuid ensures that no value is present for CategoryGuid, not even an explicit nil
### GetCheckNumberString

`func (o *TransactionResponse) GetCheckNumberString() string`

GetCheckNumberString returns the CheckNumberString field if non-nil, zero value otherwise.

### GetCheckNumberStringOk

`func (o *TransactionResponse) GetCheckNumberStringOk() (*string, bool)`

GetCheckNumberStringOk returns a tuple with the CheckNumberString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumberString

`func (o *TransactionResponse) SetCheckNumberString(v string)`

SetCheckNumberString sets CheckNumberString field to given value.

### HasCheckNumberString

`func (o *TransactionResponse) HasCheckNumberString() bool`

HasCheckNumberString returns a boolean if a field has been set.

### SetCheckNumberStringNil

`func (o *TransactionResponse) SetCheckNumberStringNil(b bool)`

 SetCheckNumberStringNil sets the value for CheckNumberString to be an explicit nil

### UnsetCheckNumberString
`func (o *TransactionResponse) UnsetCheckNumberString()`

UnsetCheckNumberString ensures that no value is present for CheckNumberString, not even an explicit nil
### GetCreatedAt

`func (o *TransactionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransactionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TransactionResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TransactionResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCurrencyCode

`func (o *TransactionResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *TransactionResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *TransactionResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *TransactionResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *TransactionResponse) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *TransactionResponse) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDate

`func (o *TransactionResponse) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionResponse) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionResponse) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *TransactionResponse) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *TransactionResponse) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *TransactionResponse) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *TransactionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TransactionResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TransactionResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtendedTransactionType

`func (o *TransactionResponse) GetExtendedTransactionType() string`

GetExtendedTransactionType returns the ExtendedTransactionType field if non-nil, zero value otherwise.

### GetExtendedTransactionTypeOk

`func (o *TransactionResponse) GetExtendedTransactionTypeOk() (*string, bool)`

GetExtendedTransactionTypeOk returns a tuple with the ExtendedTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedTransactionType

`func (o *TransactionResponse) SetExtendedTransactionType(v string)`

SetExtendedTransactionType sets ExtendedTransactionType field to given value.

### HasExtendedTransactionType

`func (o *TransactionResponse) HasExtendedTransactionType() bool`

HasExtendedTransactionType returns a boolean if a field has been set.

### SetExtendedTransactionTypeNil

`func (o *TransactionResponse) SetExtendedTransactionTypeNil(b bool)`

 SetExtendedTransactionTypeNil sets the value for ExtendedTransactionType to be an explicit nil

### UnsetExtendedTransactionType
`func (o *TransactionResponse) UnsetExtendedTransactionType()`

UnsetExtendedTransactionType ensures that no value is present for ExtendedTransactionType, not even an explicit nil
### GetGuid

`func (o *TransactionResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *TransactionResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *TransactionResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *TransactionResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *TransactionResponse) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *TransactionResponse) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetId

`func (o *TransactionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TransactionResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TransactionResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsBillPay

`func (o *TransactionResponse) GetIsBillPay() bool`

GetIsBillPay returns the IsBillPay field if non-nil, zero value otherwise.

### GetIsBillPayOk

`func (o *TransactionResponse) GetIsBillPayOk() (*bool, bool)`

GetIsBillPayOk returns a tuple with the IsBillPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillPay

`func (o *TransactionResponse) SetIsBillPay(v bool)`

SetIsBillPay sets IsBillPay field to given value.

### HasIsBillPay

`func (o *TransactionResponse) HasIsBillPay() bool`

HasIsBillPay returns a boolean if a field has been set.

### SetIsBillPayNil

`func (o *TransactionResponse) SetIsBillPayNil(b bool)`

 SetIsBillPayNil sets the value for IsBillPay to be an explicit nil

### UnsetIsBillPay
`func (o *TransactionResponse) UnsetIsBillPay()`

UnsetIsBillPay ensures that no value is present for IsBillPay, not even an explicit nil
### GetIsDirectDeposit

`func (o *TransactionResponse) GetIsDirectDeposit() bool`

GetIsDirectDeposit returns the IsDirectDeposit field if non-nil, zero value otherwise.

### GetIsDirectDepositOk

`func (o *TransactionResponse) GetIsDirectDepositOk() (*bool, bool)`

GetIsDirectDepositOk returns a tuple with the IsDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectDeposit

`func (o *TransactionResponse) SetIsDirectDeposit(v bool)`

SetIsDirectDeposit sets IsDirectDeposit field to given value.

### HasIsDirectDeposit

`func (o *TransactionResponse) HasIsDirectDeposit() bool`

HasIsDirectDeposit returns a boolean if a field has been set.

### SetIsDirectDepositNil

`func (o *TransactionResponse) SetIsDirectDepositNil(b bool)`

 SetIsDirectDepositNil sets the value for IsDirectDeposit to be an explicit nil

### UnsetIsDirectDeposit
`func (o *TransactionResponse) UnsetIsDirectDeposit()`

UnsetIsDirectDeposit ensures that no value is present for IsDirectDeposit, not even an explicit nil
### GetIsExpense

`func (o *TransactionResponse) GetIsExpense() bool`

GetIsExpense returns the IsExpense field if non-nil, zero value otherwise.

### GetIsExpenseOk

`func (o *TransactionResponse) GetIsExpenseOk() (*bool, bool)`

GetIsExpenseOk returns a tuple with the IsExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpense

`func (o *TransactionResponse) SetIsExpense(v bool)`

SetIsExpense sets IsExpense field to given value.

### HasIsExpense

`func (o *TransactionResponse) HasIsExpense() bool`

HasIsExpense returns a boolean if a field has been set.

### SetIsExpenseNil

`func (o *TransactionResponse) SetIsExpenseNil(b bool)`

 SetIsExpenseNil sets the value for IsExpense to be an explicit nil

### UnsetIsExpense
`func (o *TransactionResponse) UnsetIsExpense()`

UnsetIsExpense ensures that no value is present for IsExpense, not even an explicit nil
### GetIsFee

`func (o *TransactionResponse) GetIsFee() bool`

GetIsFee returns the IsFee field if non-nil, zero value otherwise.

### GetIsFeeOk

`func (o *TransactionResponse) GetIsFeeOk() (*bool, bool)`

GetIsFeeOk returns a tuple with the IsFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFee

`func (o *TransactionResponse) SetIsFee(v bool)`

SetIsFee sets IsFee field to given value.

### HasIsFee

`func (o *TransactionResponse) HasIsFee() bool`

HasIsFee returns a boolean if a field has been set.

### SetIsFeeNil

`func (o *TransactionResponse) SetIsFeeNil(b bool)`

 SetIsFeeNil sets the value for IsFee to be an explicit nil

### UnsetIsFee
`func (o *TransactionResponse) UnsetIsFee()`

UnsetIsFee ensures that no value is present for IsFee, not even an explicit nil
### GetIsIncome

`func (o *TransactionResponse) GetIsIncome() bool`

GetIsIncome returns the IsIncome field if non-nil, zero value otherwise.

### GetIsIncomeOk

`func (o *TransactionResponse) GetIsIncomeOk() (*bool, bool)`

GetIsIncomeOk returns a tuple with the IsIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncome

`func (o *TransactionResponse) SetIsIncome(v bool)`

SetIsIncome sets IsIncome field to given value.

### HasIsIncome

`func (o *TransactionResponse) HasIsIncome() bool`

HasIsIncome returns a boolean if a field has been set.

### SetIsIncomeNil

`func (o *TransactionResponse) SetIsIncomeNil(b bool)`

 SetIsIncomeNil sets the value for IsIncome to be an explicit nil

### UnsetIsIncome
`func (o *TransactionResponse) UnsetIsIncome()`

UnsetIsIncome ensures that no value is present for IsIncome, not even an explicit nil
### GetIsInternational

`func (o *TransactionResponse) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *TransactionResponse) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *TransactionResponse) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *TransactionResponse) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### SetIsInternationalNil

`func (o *TransactionResponse) SetIsInternationalNil(b bool)`

 SetIsInternationalNil sets the value for IsInternational to be an explicit nil

### UnsetIsInternational
`func (o *TransactionResponse) UnsetIsInternational()`

UnsetIsInternational ensures that no value is present for IsInternational, not even an explicit nil
### GetIsManual

`func (o *TransactionResponse) GetIsManual() bool`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *TransactionResponse) GetIsManualOk() (*bool, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *TransactionResponse) SetIsManual(v bool)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *TransactionResponse) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### SetIsManualNil

`func (o *TransactionResponse) SetIsManualNil(b bool)`

 SetIsManualNil sets the value for IsManual to be an explicit nil

### UnsetIsManual
`func (o *TransactionResponse) UnsetIsManual()`

UnsetIsManual ensures that no value is present for IsManual, not even an explicit nil
### GetIsOverdraftFee

`func (o *TransactionResponse) GetIsOverdraftFee() bool`

GetIsOverdraftFee returns the IsOverdraftFee field if non-nil, zero value otherwise.

### GetIsOverdraftFeeOk

`func (o *TransactionResponse) GetIsOverdraftFeeOk() (*bool, bool)`

GetIsOverdraftFeeOk returns a tuple with the IsOverdraftFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdraftFee

`func (o *TransactionResponse) SetIsOverdraftFee(v bool)`

SetIsOverdraftFee sets IsOverdraftFee field to given value.

### HasIsOverdraftFee

`func (o *TransactionResponse) HasIsOverdraftFee() bool`

HasIsOverdraftFee returns a boolean if a field has been set.

### SetIsOverdraftFeeNil

`func (o *TransactionResponse) SetIsOverdraftFeeNil(b bool)`

 SetIsOverdraftFeeNil sets the value for IsOverdraftFee to be an explicit nil

### UnsetIsOverdraftFee
`func (o *TransactionResponse) UnsetIsOverdraftFee()`

UnsetIsOverdraftFee ensures that no value is present for IsOverdraftFee, not even an explicit nil
### GetIsPayrollAdvance

`func (o *TransactionResponse) GetIsPayrollAdvance() bool`

GetIsPayrollAdvance returns the IsPayrollAdvance field if non-nil, zero value otherwise.

### GetIsPayrollAdvanceOk

`func (o *TransactionResponse) GetIsPayrollAdvanceOk() (*bool, bool)`

GetIsPayrollAdvanceOk returns a tuple with the IsPayrollAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPayrollAdvance

`func (o *TransactionResponse) SetIsPayrollAdvance(v bool)`

SetIsPayrollAdvance sets IsPayrollAdvance field to given value.

### HasIsPayrollAdvance

`func (o *TransactionResponse) HasIsPayrollAdvance() bool`

HasIsPayrollAdvance returns a boolean if a field has been set.

### SetIsPayrollAdvanceNil

`func (o *TransactionResponse) SetIsPayrollAdvanceNil(b bool)`

 SetIsPayrollAdvanceNil sets the value for IsPayrollAdvance to be an explicit nil

### UnsetIsPayrollAdvance
`func (o *TransactionResponse) UnsetIsPayrollAdvance()`

UnsetIsPayrollAdvance ensures that no value is present for IsPayrollAdvance, not even an explicit nil
### GetIsRecurring

`func (o *TransactionResponse) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *TransactionResponse) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *TransactionResponse) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.

### HasIsRecurring

`func (o *TransactionResponse) HasIsRecurring() bool`

HasIsRecurring returns a boolean if a field has been set.

### SetIsRecurringNil

`func (o *TransactionResponse) SetIsRecurringNil(b bool)`

 SetIsRecurringNil sets the value for IsRecurring to be an explicit nil

### UnsetIsRecurring
`func (o *TransactionResponse) UnsetIsRecurring()`

UnsetIsRecurring ensures that no value is present for IsRecurring, not even an explicit nil
### GetIsSubscription

`func (o *TransactionResponse) GetIsSubscription() bool`

GetIsSubscription returns the IsSubscription field if non-nil, zero value otherwise.

### GetIsSubscriptionOk

`func (o *TransactionResponse) GetIsSubscriptionOk() (*bool, bool)`

GetIsSubscriptionOk returns a tuple with the IsSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscription

`func (o *TransactionResponse) SetIsSubscription(v bool)`

SetIsSubscription sets IsSubscription field to given value.

### HasIsSubscription

`func (o *TransactionResponse) HasIsSubscription() bool`

HasIsSubscription returns a boolean if a field has been set.

### SetIsSubscriptionNil

`func (o *TransactionResponse) SetIsSubscriptionNil(b bool)`

 SetIsSubscriptionNil sets the value for IsSubscription to be an explicit nil

### UnsetIsSubscription
`func (o *TransactionResponse) UnsetIsSubscription()`

UnsetIsSubscription ensures that no value is present for IsSubscription, not even an explicit nil
### GetLatitude

`func (o *TransactionResponse) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *TransactionResponse) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *TransactionResponse) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *TransactionResponse) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *TransactionResponse) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *TransactionResponse) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLocalizedDescription

`func (o *TransactionResponse) GetLocalizedDescription() string`

GetLocalizedDescription returns the LocalizedDescription field if non-nil, zero value otherwise.

### GetLocalizedDescriptionOk

`func (o *TransactionResponse) GetLocalizedDescriptionOk() (*string, bool)`

GetLocalizedDescriptionOk returns a tuple with the LocalizedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedDescription

`func (o *TransactionResponse) SetLocalizedDescription(v string)`

SetLocalizedDescription sets LocalizedDescription field to given value.

### HasLocalizedDescription

`func (o *TransactionResponse) HasLocalizedDescription() bool`

HasLocalizedDescription returns a boolean if a field has been set.

### SetLocalizedDescriptionNil

`func (o *TransactionResponse) SetLocalizedDescriptionNil(b bool)`

 SetLocalizedDescriptionNil sets the value for LocalizedDescription to be an explicit nil

### UnsetLocalizedDescription
`func (o *TransactionResponse) UnsetLocalizedDescription()`

UnsetLocalizedDescription ensures that no value is present for LocalizedDescription, not even an explicit nil
### GetLocalizedMemo

`func (o *TransactionResponse) GetLocalizedMemo() string`

GetLocalizedMemo returns the LocalizedMemo field if non-nil, zero value otherwise.

### GetLocalizedMemoOk

`func (o *TransactionResponse) GetLocalizedMemoOk() (*string, bool)`

GetLocalizedMemoOk returns a tuple with the LocalizedMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedMemo

`func (o *TransactionResponse) SetLocalizedMemo(v string)`

SetLocalizedMemo sets LocalizedMemo field to given value.

### HasLocalizedMemo

`func (o *TransactionResponse) HasLocalizedMemo() bool`

HasLocalizedMemo returns a boolean if a field has been set.

### SetLocalizedMemoNil

`func (o *TransactionResponse) SetLocalizedMemoNil(b bool)`

 SetLocalizedMemoNil sets the value for LocalizedMemo to be an explicit nil

### UnsetLocalizedMemo
`func (o *TransactionResponse) UnsetLocalizedMemo()`

UnsetLocalizedMemo ensures that no value is present for LocalizedMemo, not even an explicit nil
### GetLongitude

`func (o *TransactionResponse) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *TransactionResponse) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *TransactionResponse) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *TransactionResponse) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *TransactionResponse) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *TransactionResponse) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetMemberGuid

`func (o *TransactionResponse) GetMemberGuid() string`

GetMemberGuid returns the MemberGuid field if non-nil, zero value otherwise.

### GetMemberGuidOk

`func (o *TransactionResponse) GetMemberGuidOk() (*string, bool)`

GetMemberGuidOk returns a tuple with the MemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGuid

`func (o *TransactionResponse) SetMemberGuid(v string)`

SetMemberGuid sets MemberGuid field to given value.

### HasMemberGuid

`func (o *TransactionResponse) HasMemberGuid() bool`

HasMemberGuid returns a boolean if a field has been set.

### SetMemberGuidNil

`func (o *TransactionResponse) SetMemberGuidNil(b bool)`

 SetMemberGuidNil sets the value for MemberGuid to be an explicit nil

### UnsetMemberGuid
`func (o *TransactionResponse) UnsetMemberGuid()`

UnsetMemberGuid ensures that no value is present for MemberGuid, not even an explicit nil
### GetMemberIsManagedByUser

`func (o *TransactionResponse) GetMemberIsManagedByUser() bool`

GetMemberIsManagedByUser returns the MemberIsManagedByUser field if non-nil, zero value otherwise.

### GetMemberIsManagedByUserOk

`func (o *TransactionResponse) GetMemberIsManagedByUserOk() (*bool, bool)`

GetMemberIsManagedByUserOk returns a tuple with the MemberIsManagedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIsManagedByUser

`func (o *TransactionResponse) SetMemberIsManagedByUser(v bool)`

SetMemberIsManagedByUser sets MemberIsManagedByUser field to given value.

### HasMemberIsManagedByUser

`func (o *TransactionResponse) HasMemberIsManagedByUser() bool`

HasMemberIsManagedByUser returns a boolean if a field has been set.

### SetMemberIsManagedByUserNil

`func (o *TransactionResponse) SetMemberIsManagedByUserNil(b bool)`

 SetMemberIsManagedByUserNil sets the value for MemberIsManagedByUser to be an explicit nil

### UnsetMemberIsManagedByUser
`func (o *TransactionResponse) UnsetMemberIsManagedByUser()`

UnsetMemberIsManagedByUser ensures that no value is present for MemberIsManagedByUser, not even an explicit nil
### GetMemo

`func (o *TransactionResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *TransactionResponse) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *TransactionResponse) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetMerchantCategoryCode

`func (o *TransactionResponse) GetMerchantCategoryCode() int32`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *TransactionResponse) GetMerchantCategoryCodeOk() (*int32, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *TransactionResponse) SetMerchantCategoryCode(v int32)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *TransactionResponse) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### SetMerchantCategoryCodeNil

`func (o *TransactionResponse) SetMerchantCategoryCodeNil(b bool)`

 SetMerchantCategoryCodeNil sets the value for MerchantCategoryCode to be an explicit nil

### UnsetMerchantCategoryCode
`func (o *TransactionResponse) UnsetMerchantCategoryCode()`

UnsetMerchantCategoryCode ensures that no value is present for MerchantCategoryCode, not even an explicit nil
### GetMerchantGuid

`func (o *TransactionResponse) GetMerchantGuid() string`

GetMerchantGuid returns the MerchantGuid field if non-nil, zero value otherwise.

### GetMerchantGuidOk

`func (o *TransactionResponse) GetMerchantGuidOk() (*string, bool)`

GetMerchantGuidOk returns a tuple with the MerchantGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGuid

`func (o *TransactionResponse) SetMerchantGuid(v string)`

SetMerchantGuid sets MerchantGuid field to given value.

### HasMerchantGuid

`func (o *TransactionResponse) HasMerchantGuid() bool`

HasMerchantGuid returns a boolean if a field has been set.

### SetMerchantGuidNil

`func (o *TransactionResponse) SetMerchantGuidNil(b bool)`

 SetMerchantGuidNil sets the value for MerchantGuid to be an explicit nil

### UnsetMerchantGuid
`func (o *TransactionResponse) UnsetMerchantGuid()`

UnsetMerchantGuid ensures that no value is present for MerchantGuid, not even an explicit nil
### GetMerchantLocationGuid

`func (o *TransactionResponse) GetMerchantLocationGuid() string`

GetMerchantLocationGuid returns the MerchantLocationGuid field if non-nil, zero value otherwise.

### GetMerchantLocationGuidOk

`func (o *TransactionResponse) GetMerchantLocationGuidOk() (*string, bool)`

GetMerchantLocationGuidOk returns a tuple with the MerchantLocationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantLocationGuid

`func (o *TransactionResponse) SetMerchantLocationGuid(v string)`

SetMerchantLocationGuid sets MerchantLocationGuid field to given value.

### HasMerchantLocationGuid

`func (o *TransactionResponse) HasMerchantLocationGuid() bool`

HasMerchantLocationGuid returns a boolean if a field has been set.

### SetMerchantLocationGuidNil

`func (o *TransactionResponse) SetMerchantLocationGuidNil(b bool)`

 SetMerchantLocationGuidNil sets the value for MerchantLocationGuid to be an explicit nil

### UnsetMerchantLocationGuid
`func (o *TransactionResponse) UnsetMerchantLocationGuid()`

UnsetMerchantLocationGuid ensures that no value is present for MerchantLocationGuid, not even an explicit nil
### GetMetadata

`func (o *TransactionResponse) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransactionResponse) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransactionResponse) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransactionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *TransactionResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TransactionResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOriginalDescription

`func (o *TransactionResponse) GetOriginalDescription() string`

GetOriginalDescription returns the OriginalDescription field if non-nil, zero value otherwise.

### GetOriginalDescriptionOk

`func (o *TransactionResponse) GetOriginalDescriptionOk() (*string, bool)`

GetOriginalDescriptionOk returns a tuple with the OriginalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDescription

`func (o *TransactionResponse) SetOriginalDescription(v string)`

SetOriginalDescription sets OriginalDescription field to given value.

### HasOriginalDescription

`func (o *TransactionResponse) HasOriginalDescription() bool`

HasOriginalDescription returns a boolean if a field has been set.

### SetOriginalDescriptionNil

`func (o *TransactionResponse) SetOriginalDescriptionNil(b bool)`

 SetOriginalDescriptionNil sets the value for OriginalDescription to be an explicit nil

### UnsetOriginalDescription
`func (o *TransactionResponse) UnsetOriginalDescription()`

UnsetOriginalDescription ensures that no value is present for OriginalDescription, not even an explicit nil
### GetPostedAt

`func (o *TransactionResponse) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *TransactionResponse) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *TransactionResponse) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *TransactionResponse) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### SetPostedAtNil

`func (o *TransactionResponse) SetPostedAtNil(b bool)`

 SetPostedAtNil sets the value for PostedAt to be an explicit nil

### UnsetPostedAt
`func (o *TransactionResponse) UnsetPostedAt()`

UnsetPostedAt ensures that no value is present for PostedAt, not even an explicit nil
### GetStatus

`func (o *TransactionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TransactionResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TransactionResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTopLevelCategory

`func (o *TransactionResponse) GetTopLevelCategory() string`

GetTopLevelCategory returns the TopLevelCategory field if non-nil, zero value otherwise.

### GetTopLevelCategoryOk

`func (o *TransactionResponse) GetTopLevelCategoryOk() (*string, bool)`

GetTopLevelCategoryOk returns a tuple with the TopLevelCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevelCategory

`func (o *TransactionResponse) SetTopLevelCategory(v string)`

SetTopLevelCategory sets TopLevelCategory field to given value.

### HasTopLevelCategory

`func (o *TransactionResponse) HasTopLevelCategory() bool`

HasTopLevelCategory returns a boolean if a field has been set.

### SetTopLevelCategoryNil

`func (o *TransactionResponse) SetTopLevelCategoryNil(b bool)`

 SetTopLevelCategoryNil sets the value for TopLevelCategory to be an explicit nil

### UnsetTopLevelCategory
`func (o *TransactionResponse) UnsetTopLevelCategory()`

UnsetTopLevelCategory ensures that no value is present for TopLevelCategory, not even an explicit nil
### GetTransactedAt

`func (o *TransactionResponse) GetTransactedAt() string`

GetTransactedAt returns the TransactedAt field if non-nil, zero value otherwise.

### GetTransactedAtOk

`func (o *TransactionResponse) GetTransactedAtOk() (*string, bool)`

GetTransactedAtOk returns a tuple with the TransactedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactedAt

`func (o *TransactionResponse) SetTransactedAt(v string)`

SetTransactedAt sets TransactedAt field to given value.

### HasTransactedAt

`func (o *TransactionResponse) HasTransactedAt() bool`

HasTransactedAt returns a boolean if a field has been set.

### SetTransactedAtNil

`func (o *TransactionResponse) SetTransactedAtNil(b bool)`

 SetTransactedAtNil sets the value for TransactedAt to be an explicit nil

### UnsetTransactedAt
`func (o *TransactionResponse) UnsetTransactedAt()`

UnsetTransactedAt ensures that no value is present for TransactedAt, not even an explicit nil
### GetType

`func (o *TransactionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TransactionResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TransactionResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUpdatedAt

`func (o *TransactionResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransactionResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransactionResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TransactionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TransactionResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TransactionResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserGuid

`func (o *TransactionResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *TransactionResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *TransactionResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *TransactionResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### SetUserGuidNil

`func (o *TransactionResponse) SetUserGuidNil(b bool)`

 SetUserGuidNil sets the value for UserGuid to be an explicit nil

### UnsetUserGuid
`func (o *TransactionResponse) UnsetUserGuid()`

UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
### GetUserId

`func (o *TransactionResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TransactionResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TransactionResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TransactionResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *TransactionResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *TransactionResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


