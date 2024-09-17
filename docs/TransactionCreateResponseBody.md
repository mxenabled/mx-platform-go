# TransactionCreateResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
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
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserGuid** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTransactionCreateResponseBody

`func NewTransactionCreateResponseBody() *TransactionCreateResponseBody`

NewTransactionCreateResponseBody instantiates a new TransactionCreateResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCreateResponseBodyWithDefaults

`func NewTransactionCreateResponseBodyWithDefaults() *TransactionCreateResponseBody`

NewTransactionCreateResponseBodyWithDefaults instantiates a new TransactionCreateResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *TransactionCreateResponseBody) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *TransactionCreateResponseBody) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *TransactionCreateResponseBody) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.

### HasAccountGuid

`func (o *TransactionCreateResponseBody) HasAccountGuid() bool`

HasAccountGuid returns a boolean if a field has been set.

### SetAccountGuidNil

`func (o *TransactionCreateResponseBody) SetAccountGuidNil(b bool)`

 SetAccountGuidNil sets the value for AccountGuid to be an explicit nil

### UnsetAccountGuid
`func (o *TransactionCreateResponseBody) UnsetAccountGuid()`

UnsetAccountGuid ensures that no value is present for AccountGuid, not even an explicit nil
### GetAccountId

`func (o *TransactionCreateResponseBody) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransactionCreateResponseBody) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransactionCreateResponseBody) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TransactionCreateResponseBody) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *TransactionCreateResponseBody) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *TransactionCreateResponseBody) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAmount

`func (o *TransactionCreateResponseBody) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionCreateResponseBody) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionCreateResponseBody) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionCreateResponseBody) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCategory

`func (o *TransactionCreateResponseBody) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TransactionCreateResponseBody) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TransactionCreateResponseBody) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TransactionCreateResponseBody) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *TransactionCreateResponseBody) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *TransactionCreateResponseBody) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCategoryGuid

`func (o *TransactionCreateResponseBody) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *TransactionCreateResponseBody) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *TransactionCreateResponseBody) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *TransactionCreateResponseBody) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### SetCategoryGuidNil

`func (o *TransactionCreateResponseBody) SetCategoryGuidNil(b bool)`

 SetCategoryGuidNil sets the value for CategoryGuid to be an explicit nil

### UnsetCategoryGuid
`func (o *TransactionCreateResponseBody) UnsetCategoryGuid()`

UnsetCategoryGuid ensures that no value is present for CategoryGuid, not even an explicit nil
### GetCheckNumberString

`func (o *TransactionCreateResponseBody) GetCheckNumberString() string`

GetCheckNumberString returns the CheckNumberString field if non-nil, zero value otherwise.

### GetCheckNumberStringOk

`func (o *TransactionCreateResponseBody) GetCheckNumberStringOk() (*string, bool)`

GetCheckNumberStringOk returns a tuple with the CheckNumberString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumberString

`func (o *TransactionCreateResponseBody) SetCheckNumberString(v string)`

SetCheckNumberString sets CheckNumberString field to given value.

### HasCheckNumberString

`func (o *TransactionCreateResponseBody) HasCheckNumberString() bool`

HasCheckNumberString returns a boolean if a field has been set.

### SetCheckNumberStringNil

`func (o *TransactionCreateResponseBody) SetCheckNumberStringNil(b bool)`

 SetCheckNumberStringNil sets the value for CheckNumberString to be an explicit nil

### UnsetCheckNumberString
`func (o *TransactionCreateResponseBody) UnsetCheckNumberString()`

UnsetCheckNumberString ensures that no value is present for CheckNumberString, not even an explicit nil
### GetCreatedAt

`func (o *TransactionCreateResponseBody) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionCreateResponseBody) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionCreateResponseBody) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransactionCreateResponseBody) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TransactionCreateResponseBody) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TransactionCreateResponseBody) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCurrencyCode

`func (o *TransactionCreateResponseBody) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *TransactionCreateResponseBody) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *TransactionCreateResponseBody) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *TransactionCreateResponseBody) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *TransactionCreateResponseBody) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *TransactionCreateResponseBody) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDate

`func (o *TransactionCreateResponseBody) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionCreateResponseBody) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionCreateResponseBody) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *TransactionCreateResponseBody) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *TransactionCreateResponseBody) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *TransactionCreateResponseBody) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *TransactionCreateResponseBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionCreateResponseBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionCreateResponseBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionCreateResponseBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TransactionCreateResponseBody) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TransactionCreateResponseBody) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtendedTransactionType

`func (o *TransactionCreateResponseBody) GetExtendedTransactionType() string`

GetExtendedTransactionType returns the ExtendedTransactionType field if non-nil, zero value otherwise.

### GetExtendedTransactionTypeOk

`func (o *TransactionCreateResponseBody) GetExtendedTransactionTypeOk() (*string, bool)`

GetExtendedTransactionTypeOk returns a tuple with the ExtendedTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedTransactionType

`func (o *TransactionCreateResponseBody) SetExtendedTransactionType(v string)`

SetExtendedTransactionType sets ExtendedTransactionType field to given value.

### HasExtendedTransactionType

`func (o *TransactionCreateResponseBody) HasExtendedTransactionType() bool`

HasExtendedTransactionType returns a boolean if a field has been set.

### SetExtendedTransactionTypeNil

`func (o *TransactionCreateResponseBody) SetExtendedTransactionTypeNil(b bool)`

 SetExtendedTransactionTypeNil sets the value for ExtendedTransactionType to be an explicit nil

### UnsetExtendedTransactionType
`func (o *TransactionCreateResponseBody) UnsetExtendedTransactionType()`

UnsetExtendedTransactionType ensures that no value is present for ExtendedTransactionType, not even an explicit nil
### GetGuid

`func (o *TransactionCreateResponseBody) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *TransactionCreateResponseBody) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *TransactionCreateResponseBody) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *TransactionCreateResponseBody) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *TransactionCreateResponseBody) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *TransactionCreateResponseBody) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetId

`func (o *TransactionCreateResponseBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionCreateResponseBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionCreateResponseBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionCreateResponseBody) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TransactionCreateResponseBody) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TransactionCreateResponseBody) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsBillPay

`func (o *TransactionCreateResponseBody) GetIsBillPay() bool`

GetIsBillPay returns the IsBillPay field if non-nil, zero value otherwise.

### GetIsBillPayOk

`func (o *TransactionCreateResponseBody) GetIsBillPayOk() (*bool, bool)`

GetIsBillPayOk returns a tuple with the IsBillPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillPay

`func (o *TransactionCreateResponseBody) SetIsBillPay(v bool)`

SetIsBillPay sets IsBillPay field to given value.

### HasIsBillPay

`func (o *TransactionCreateResponseBody) HasIsBillPay() bool`

HasIsBillPay returns a boolean if a field has been set.

### SetIsBillPayNil

`func (o *TransactionCreateResponseBody) SetIsBillPayNil(b bool)`

 SetIsBillPayNil sets the value for IsBillPay to be an explicit nil

### UnsetIsBillPay
`func (o *TransactionCreateResponseBody) UnsetIsBillPay()`

UnsetIsBillPay ensures that no value is present for IsBillPay, not even an explicit nil
### GetIsDirectDeposit

`func (o *TransactionCreateResponseBody) GetIsDirectDeposit() bool`

GetIsDirectDeposit returns the IsDirectDeposit field if non-nil, zero value otherwise.

### GetIsDirectDepositOk

`func (o *TransactionCreateResponseBody) GetIsDirectDepositOk() (*bool, bool)`

GetIsDirectDepositOk returns a tuple with the IsDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectDeposit

`func (o *TransactionCreateResponseBody) SetIsDirectDeposit(v bool)`

SetIsDirectDeposit sets IsDirectDeposit field to given value.

### HasIsDirectDeposit

`func (o *TransactionCreateResponseBody) HasIsDirectDeposit() bool`

HasIsDirectDeposit returns a boolean if a field has been set.

### SetIsDirectDepositNil

`func (o *TransactionCreateResponseBody) SetIsDirectDepositNil(b bool)`

 SetIsDirectDepositNil sets the value for IsDirectDeposit to be an explicit nil

### UnsetIsDirectDeposit
`func (o *TransactionCreateResponseBody) UnsetIsDirectDeposit()`

UnsetIsDirectDeposit ensures that no value is present for IsDirectDeposit, not even an explicit nil
### GetIsExpense

`func (o *TransactionCreateResponseBody) GetIsExpense() bool`

GetIsExpense returns the IsExpense field if non-nil, zero value otherwise.

### GetIsExpenseOk

`func (o *TransactionCreateResponseBody) GetIsExpenseOk() (*bool, bool)`

GetIsExpenseOk returns a tuple with the IsExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpense

`func (o *TransactionCreateResponseBody) SetIsExpense(v bool)`

SetIsExpense sets IsExpense field to given value.

### HasIsExpense

`func (o *TransactionCreateResponseBody) HasIsExpense() bool`

HasIsExpense returns a boolean if a field has been set.

### SetIsExpenseNil

`func (o *TransactionCreateResponseBody) SetIsExpenseNil(b bool)`

 SetIsExpenseNil sets the value for IsExpense to be an explicit nil

### UnsetIsExpense
`func (o *TransactionCreateResponseBody) UnsetIsExpense()`

UnsetIsExpense ensures that no value is present for IsExpense, not even an explicit nil
### GetIsFee

`func (o *TransactionCreateResponseBody) GetIsFee() bool`

GetIsFee returns the IsFee field if non-nil, zero value otherwise.

### GetIsFeeOk

`func (o *TransactionCreateResponseBody) GetIsFeeOk() (*bool, bool)`

GetIsFeeOk returns a tuple with the IsFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFee

`func (o *TransactionCreateResponseBody) SetIsFee(v bool)`

SetIsFee sets IsFee field to given value.

### HasIsFee

`func (o *TransactionCreateResponseBody) HasIsFee() bool`

HasIsFee returns a boolean if a field has been set.

### SetIsFeeNil

`func (o *TransactionCreateResponseBody) SetIsFeeNil(b bool)`

 SetIsFeeNil sets the value for IsFee to be an explicit nil

### UnsetIsFee
`func (o *TransactionCreateResponseBody) UnsetIsFee()`

UnsetIsFee ensures that no value is present for IsFee, not even an explicit nil
### GetIsIncome

`func (o *TransactionCreateResponseBody) GetIsIncome() bool`

GetIsIncome returns the IsIncome field if non-nil, zero value otherwise.

### GetIsIncomeOk

`func (o *TransactionCreateResponseBody) GetIsIncomeOk() (*bool, bool)`

GetIsIncomeOk returns a tuple with the IsIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncome

`func (o *TransactionCreateResponseBody) SetIsIncome(v bool)`

SetIsIncome sets IsIncome field to given value.

### HasIsIncome

`func (o *TransactionCreateResponseBody) HasIsIncome() bool`

HasIsIncome returns a boolean if a field has been set.

### SetIsIncomeNil

`func (o *TransactionCreateResponseBody) SetIsIncomeNil(b bool)`

 SetIsIncomeNil sets the value for IsIncome to be an explicit nil

### UnsetIsIncome
`func (o *TransactionCreateResponseBody) UnsetIsIncome()`

UnsetIsIncome ensures that no value is present for IsIncome, not even an explicit nil
### GetIsInternational

`func (o *TransactionCreateResponseBody) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *TransactionCreateResponseBody) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *TransactionCreateResponseBody) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *TransactionCreateResponseBody) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### SetIsInternationalNil

`func (o *TransactionCreateResponseBody) SetIsInternationalNil(b bool)`

 SetIsInternationalNil sets the value for IsInternational to be an explicit nil

### UnsetIsInternational
`func (o *TransactionCreateResponseBody) UnsetIsInternational()`

UnsetIsInternational ensures that no value is present for IsInternational, not even an explicit nil
### GetIsManual

`func (o *TransactionCreateResponseBody) GetIsManual() bool`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *TransactionCreateResponseBody) GetIsManualOk() (*bool, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *TransactionCreateResponseBody) SetIsManual(v bool)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *TransactionCreateResponseBody) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### SetIsManualNil

`func (o *TransactionCreateResponseBody) SetIsManualNil(b bool)`

 SetIsManualNil sets the value for IsManual to be an explicit nil

### UnsetIsManual
`func (o *TransactionCreateResponseBody) UnsetIsManual()`

UnsetIsManual ensures that no value is present for IsManual, not even an explicit nil
### GetIsOverdraftFee

`func (o *TransactionCreateResponseBody) GetIsOverdraftFee() bool`

GetIsOverdraftFee returns the IsOverdraftFee field if non-nil, zero value otherwise.

### GetIsOverdraftFeeOk

`func (o *TransactionCreateResponseBody) GetIsOverdraftFeeOk() (*bool, bool)`

GetIsOverdraftFeeOk returns a tuple with the IsOverdraftFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdraftFee

`func (o *TransactionCreateResponseBody) SetIsOverdraftFee(v bool)`

SetIsOverdraftFee sets IsOverdraftFee field to given value.

### HasIsOverdraftFee

`func (o *TransactionCreateResponseBody) HasIsOverdraftFee() bool`

HasIsOverdraftFee returns a boolean if a field has been set.

### SetIsOverdraftFeeNil

`func (o *TransactionCreateResponseBody) SetIsOverdraftFeeNil(b bool)`

 SetIsOverdraftFeeNil sets the value for IsOverdraftFee to be an explicit nil

### UnsetIsOverdraftFee
`func (o *TransactionCreateResponseBody) UnsetIsOverdraftFee()`

UnsetIsOverdraftFee ensures that no value is present for IsOverdraftFee, not even an explicit nil
### GetIsPayrollAdvance

`func (o *TransactionCreateResponseBody) GetIsPayrollAdvance() bool`

GetIsPayrollAdvance returns the IsPayrollAdvance field if non-nil, zero value otherwise.

### GetIsPayrollAdvanceOk

`func (o *TransactionCreateResponseBody) GetIsPayrollAdvanceOk() (*bool, bool)`

GetIsPayrollAdvanceOk returns a tuple with the IsPayrollAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPayrollAdvance

`func (o *TransactionCreateResponseBody) SetIsPayrollAdvance(v bool)`

SetIsPayrollAdvance sets IsPayrollAdvance field to given value.

### HasIsPayrollAdvance

`func (o *TransactionCreateResponseBody) HasIsPayrollAdvance() bool`

HasIsPayrollAdvance returns a boolean if a field has been set.

### SetIsPayrollAdvanceNil

`func (o *TransactionCreateResponseBody) SetIsPayrollAdvanceNil(b bool)`

 SetIsPayrollAdvanceNil sets the value for IsPayrollAdvance to be an explicit nil

### UnsetIsPayrollAdvance
`func (o *TransactionCreateResponseBody) UnsetIsPayrollAdvance()`

UnsetIsPayrollAdvance ensures that no value is present for IsPayrollAdvance, not even an explicit nil
### GetIsRecurring

`func (o *TransactionCreateResponseBody) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *TransactionCreateResponseBody) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *TransactionCreateResponseBody) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.

### HasIsRecurring

`func (o *TransactionCreateResponseBody) HasIsRecurring() bool`

HasIsRecurring returns a boolean if a field has been set.

### SetIsRecurringNil

`func (o *TransactionCreateResponseBody) SetIsRecurringNil(b bool)`

 SetIsRecurringNil sets the value for IsRecurring to be an explicit nil

### UnsetIsRecurring
`func (o *TransactionCreateResponseBody) UnsetIsRecurring()`

UnsetIsRecurring ensures that no value is present for IsRecurring, not even an explicit nil
### GetIsSubscription

`func (o *TransactionCreateResponseBody) GetIsSubscription() bool`

GetIsSubscription returns the IsSubscription field if non-nil, zero value otherwise.

### GetIsSubscriptionOk

`func (o *TransactionCreateResponseBody) GetIsSubscriptionOk() (*bool, bool)`

GetIsSubscriptionOk returns a tuple with the IsSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscription

`func (o *TransactionCreateResponseBody) SetIsSubscription(v bool)`

SetIsSubscription sets IsSubscription field to given value.

### HasIsSubscription

`func (o *TransactionCreateResponseBody) HasIsSubscription() bool`

HasIsSubscription returns a boolean if a field has been set.

### SetIsSubscriptionNil

`func (o *TransactionCreateResponseBody) SetIsSubscriptionNil(b bool)`

 SetIsSubscriptionNil sets the value for IsSubscription to be an explicit nil

### UnsetIsSubscription
`func (o *TransactionCreateResponseBody) UnsetIsSubscription()`

UnsetIsSubscription ensures that no value is present for IsSubscription, not even an explicit nil
### GetLatitude

`func (o *TransactionCreateResponseBody) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *TransactionCreateResponseBody) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *TransactionCreateResponseBody) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *TransactionCreateResponseBody) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *TransactionCreateResponseBody) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *TransactionCreateResponseBody) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLocalizedDescription

`func (o *TransactionCreateResponseBody) GetLocalizedDescription() string`

GetLocalizedDescription returns the LocalizedDescription field if non-nil, zero value otherwise.

### GetLocalizedDescriptionOk

`func (o *TransactionCreateResponseBody) GetLocalizedDescriptionOk() (*string, bool)`

GetLocalizedDescriptionOk returns a tuple with the LocalizedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedDescription

`func (o *TransactionCreateResponseBody) SetLocalizedDescription(v string)`

SetLocalizedDescription sets LocalizedDescription field to given value.

### HasLocalizedDescription

`func (o *TransactionCreateResponseBody) HasLocalizedDescription() bool`

HasLocalizedDescription returns a boolean if a field has been set.

### SetLocalizedDescriptionNil

`func (o *TransactionCreateResponseBody) SetLocalizedDescriptionNil(b bool)`

 SetLocalizedDescriptionNil sets the value for LocalizedDescription to be an explicit nil

### UnsetLocalizedDescription
`func (o *TransactionCreateResponseBody) UnsetLocalizedDescription()`

UnsetLocalizedDescription ensures that no value is present for LocalizedDescription, not even an explicit nil
### GetLocalizedMemo

`func (o *TransactionCreateResponseBody) GetLocalizedMemo() string`

GetLocalizedMemo returns the LocalizedMemo field if non-nil, zero value otherwise.

### GetLocalizedMemoOk

`func (o *TransactionCreateResponseBody) GetLocalizedMemoOk() (*string, bool)`

GetLocalizedMemoOk returns a tuple with the LocalizedMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedMemo

`func (o *TransactionCreateResponseBody) SetLocalizedMemo(v string)`

SetLocalizedMemo sets LocalizedMemo field to given value.

### HasLocalizedMemo

`func (o *TransactionCreateResponseBody) HasLocalizedMemo() bool`

HasLocalizedMemo returns a boolean if a field has been set.

### SetLocalizedMemoNil

`func (o *TransactionCreateResponseBody) SetLocalizedMemoNil(b bool)`

 SetLocalizedMemoNil sets the value for LocalizedMemo to be an explicit nil

### UnsetLocalizedMemo
`func (o *TransactionCreateResponseBody) UnsetLocalizedMemo()`

UnsetLocalizedMemo ensures that no value is present for LocalizedMemo, not even an explicit nil
### GetLongitude

`func (o *TransactionCreateResponseBody) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *TransactionCreateResponseBody) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *TransactionCreateResponseBody) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *TransactionCreateResponseBody) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *TransactionCreateResponseBody) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *TransactionCreateResponseBody) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetMemberGuid

`func (o *TransactionCreateResponseBody) GetMemberGuid() string`

GetMemberGuid returns the MemberGuid field if non-nil, zero value otherwise.

### GetMemberGuidOk

`func (o *TransactionCreateResponseBody) GetMemberGuidOk() (*string, bool)`

GetMemberGuidOk returns a tuple with the MemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGuid

`func (o *TransactionCreateResponseBody) SetMemberGuid(v string)`

SetMemberGuid sets MemberGuid field to given value.

### HasMemberGuid

`func (o *TransactionCreateResponseBody) HasMemberGuid() bool`

HasMemberGuid returns a boolean if a field has been set.

### SetMemberGuidNil

`func (o *TransactionCreateResponseBody) SetMemberGuidNil(b bool)`

 SetMemberGuidNil sets the value for MemberGuid to be an explicit nil

### UnsetMemberGuid
`func (o *TransactionCreateResponseBody) UnsetMemberGuid()`

UnsetMemberGuid ensures that no value is present for MemberGuid, not even an explicit nil
### GetMemberIsManagedByUser

`func (o *TransactionCreateResponseBody) GetMemberIsManagedByUser() bool`

GetMemberIsManagedByUser returns the MemberIsManagedByUser field if non-nil, zero value otherwise.

### GetMemberIsManagedByUserOk

`func (o *TransactionCreateResponseBody) GetMemberIsManagedByUserOk() (*bool, bool)`

GetMemberIsManagedByUserOk returns a tuple with the MemberIsManagedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIsManagedByUser

`func (o *TransactionCreateResponseBody) SetMemberIsManagedByUser(v bool)`

SetMemberIsManagedByUser sets MemberIsManagedByUser field to given value.

### HasMemberIsManagedByUser

`func (o *TransactionCreateResponseBody) HasMemberIsManagedByUser() bool`

HasMemberIsManagedByUser returns a boolean if a field has been set.

### SetMemberIsManagedByUserNil

`func (o *TransactionCreateResponseBody) SetMemberIsManagedByUserNil(b bool)`

 SetMemberIsManagedByUserNil sets the value for MemberIsManagedByUser to be an explicit nil

### UnsetMemberIsManagedByUser
`func (o *TransactionCreateResponseBody) UnsetMemberIsManagedByUser()`

UnsetMemberIsManagedByUser ensures that no value is present for MemberIsManagedByUser, not even an explicit nil
### GetMemo

`func (o *TransactionCreateResponseBody) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionCreateResponseBody) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionCreateResponseBody) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionCreateResponseBody) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *TransactionCreateResponseBody) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *TransactionCreateResponseBody) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetMerchantCategoryCode

`func (o *TransactionCreateResponseBody) GetMerchantCategoryCode() int32`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *TransactionCreateResponseBody) GetMerchantCategoryCodeOk() (*int32, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *TransactionCreateResponseBody) SetMerchantCategoryCode(v int32)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *TransactionCreateResponseBody) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### SetMerchantCategoryCodeNil

`func (o *TransactionCreateResponseBody) SetMerchantCategoryCodeNil(b bool)`

 SetMerchantCategoryCodeNil sets the value for MerchantCategoryCode to be an explicit nil

### UnsetMerchantCategoryCode
`func (o *TransactionCreateResponseBody) UnsetMerchantCategoryCode()`

UnsetMerchantCategoryCode ensures that no value is present for MerchantCategoryCode, not even an explicit nil
### GetMerchantGuid

`func (o *TransactionCreateResponseBody) GetMerchantGuid() string`

GetMerchantGuid returns the MerchantGuid field if non-nil, zero value otherwise.

### GetMerchantGuidOk

`func (o *TransactionCreateResponseBody) GetMerchantGuidOk() (*string, bool)`

GetMerchantGuidOk returns a tuple with the MerchantGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGuid

`func (o *TransactionCreateResponseBody) SetMerchantGuid(v string)`

SetMerchantGuid sets MerchantGuid field to given value.

### HasMerchantGuid

`func (o *TransactionCreateResponseBody) HasMerchantGuid() bool`

HasMerchantGuid returns a boolean if a field has been set.

### SetMerchantGuidNil

`func (o *TransactionCreateResponseBody) SetMerchantGuidNil(b bool)`

 SetMerchantGuidNil sets the value for MerchantGuid to be an explicit nil

### UnsetMerchantGuid
`func (o *TransactionCreateResponseBody) UnsetMerchantGuid()`

UnsetMerchantGuid ensures that no value is present for MerchantGuid, not even an explicit nil
### GetMerchantLocationGuid

`func (o *TransactionCreateResponseBody) GetMerchantLocationGuid() string`

GetMerchantLocationGuid returns the MerchantLocationGuid field if non-nil, zero value otherwise.

### GetMerchantLocationGuidOk

`func (o *TransactionCreateResponseBody) GetMerchantLocationGuidOk() (*string, bool)`

GetMerchantLocationGuidOk returns a tuple with the MerchantLocationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantLocationGuid

`func (o *TransactionCreateResponseBody) SetMerchantLocationGuid(v string)`

SetMerchantLocationGuid sets MerchantLocationGuid field to given value.

### HasMerchantLocationGuid

`func (o *TransactionCreateResponseBody) HasMerchantLocationGuid() bool`

HasMerchantLocationGuid returns a boolean if a field has been set.

### SetMerchantLocationGuidNil

`func (o *TransactionCreateResponseBody) SetMerchantLocationGuidNil(b bool)`

 SetMerchantLocationGuidNil sets the value for MerchantLocationGuid to be an explicit nil

### UnsetMerchantLocationGuid
`func (o *TransactionCreateResponseBody) UnsetMerchantLocationGuid()`

UnsetMerchantLocationGuid ensures that no value is present for MerchantLocationGuid, not even an explicit nil
### GetMetadata

`func (o *TransactionCreateResponseBody) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransactionCreateResponseBody) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransactionCreateResponseBody) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransactionCreateResponseBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *TransactionCreateResponseBody) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TransactionCreateResponseBody) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOriginalDescription

`func (o *TransactionCreateResponseBody) GetOriginalDescription() string`

GetOriginalDescription returns the OriginalDescription field if non-nil, zero value otherwise.

### GetOriginalDescriptionOk

`func (o *TransactionCreateResponseBody) GetOriginalDescriptionOk() (*string, bool)`

GetOriginalDescriptionOk returns a tuple with the OriginalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDescription

`func (o *TransactionCreateResponseBody) SetOriginalDescription(v string)`

SetOriginalDescription sets OriginalDescription field to given value.

### HasOriginalDescription

`func (o *TransactionCreateResponseBody) HasOriginalDescription() bool`

HasOriginalDescription returns a boolean if a field has been set.

### SetOriginalDescriptionNil

`func (o *TransactionCreateResponseBody) SetOriginalDescriptionNil(b bool)`

 SetOriginalDescriptionNil sets the value for OriginalDescription to be an explicit nil

### UnsetOriginalDescription
`func (o *TransactionCreateResponseBody) UnsetOriginalDescription()`

UnsetOriginalDescription ensures that no value is present for OriginalDescription, not even an explicit nil
### GetPostedAt

`func (o *TransactionCreateResponseBody) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *TransactionCreateResponseBody) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *TransactionCreateResponseBody) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *TransactionCreateResponseBody) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### SetPostedAtNil

`func (o *TransactionCreateResponseBody) SetPostedAtNil(b bool)`

 SetPostedAtNil sets the value for PostedAt to be an explicit nil

### UnsetPostedAt
`func (o *TransactionCreateResponseBody) UnsetPostedAt()`

UnsetPostedAt ensures that no value is present for PostedAt, not even an explicit nil
### GetStatus

`func (o *TransactionCreateResponseBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionCreateResponseBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionCreateResponseBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionCreateResponseBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TransactionCreateResponseBody) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TransactionCreateResponseBody) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTopLevelCategory

`func (o *TransactionCreateResponseBody) GetTopLevelCategory() string`

GetTopLevelCategory returns the TopLevelCategory field if non-nil, zero value otherwise.

### GetTopLevelCategoryOk

`func (o *TransactionCreateResponseBody) GetTopLevelCategoryOk() (*string, bool)`

GetTopLevelCategoryOk returns a tuple with the TopLevelCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevelCategory

`func (o *TransactionCreateResponseBody) SetTopLevelCategory(v string)`

SetTopLevelCategory sets TopLevelCategory field to given value.

### HasTopLevelCategory

`func (o *TransactionCreateResponseBody) HasTopLevelCategory() bool`

HasTopLevelCategory returns a boolean if a field has been set.

### SetTopLevelCategoryNil

`func (o *TransactionCreateResponseBody) SetTopLevelCategoryNil(b bool)`

 SetTopLevelCategoryNil sets the value for TopLevelCategory to be an explicit nil

### UnsetTopLevelCategory
`func (o *TransactionCreateResponseBody) UnsetTopLevelCategory()`

UnsetTopLevelCategory ensures that no value is present for TopLevelCategory, not even an explicit nil
### GetTransactedAt

`func (o *TransactionCreateResponseBody) GetTransactedAt() string`

GetTransactedAt returns the TransactedAt field if non-nil, zero value otherwise.

### GetTransactedAtOk

`func (o *TransactionCreateResponseBody) GetTransactedAtOk() (*string, bool)`

GetTransactedAtOk returns a tuple with the TransactedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactedAt

`func (o *TransactionCreateResponseBody) SetTransactedAt(v string)`

SetTransactedAt sets TransactedAt field to given value.

### HasTransactedAt

`func (o *TransactionCreateResponseBody) HasTransactedAt() bool`

HasTransactedAt returns a boolean if a field has been set.

### SetTransactedAtNil

`func (o *TransactionCreateResponseBody) SetTransactedAtNil(b bool)`

 SetTransactedAtNil sets the value for TransactedAt to be an explicit nil

### UnsetTransactedAt
`func (o *TransactionCreateResponseBody) UnsetTransactedAt()`

UnsetTransactedAt ensures that no value is present for TransactedAt, not even an explicit nil
### GetType

`func (o *TransactionCreateResponseBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionCreateResponseBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionCreateResponseBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionCreateResponseBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TransactionCreateResponseBody) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransactionCreateResponseBody) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransactionCreateResponseBody) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TransactionCreateResponseBody) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserGuid

`func (o *TransactionCreateResponseBody) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *TransactionCreateResponseBody) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *TransactionCreateResponseBody) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *TransactionCreateResponseBody) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### SetUserGuidNil

`func (o *TransactionCreateResponseBody) SetUserGuidNil(b bool)`

 SetUserGuidNil sets the value for UserGuid to be an explicit nil

### UnsetUserGuid
`func (o *TransactionCreateResponseBody) UnsetUserGuid()`

UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
### GetUserId

`func (o *TransactionCreateResponseBody) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TransactionCreateResponseBody) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TransactionCreateResponseBody) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TransactionCreateResponseBody) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *TransactionCreateResponseBody) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *TransactionCreateResponseBody) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


