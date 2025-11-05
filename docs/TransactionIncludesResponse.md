# TransactionIncludesResponse

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
**Classification** | Pointer to [**NullableTransactionIncludesResponseAllOfClassification**](TransactionIncludesResponseAllOfClassification.md) |  | [optional] 
**Geolocation** | Pointer to [**NullableTransactionIncludesResponseAllOfGeolocation**](TransactionIncludesResponseAllOfGeolocation.md) |  | [optional] 
**Merchant** | Pointer to [**NullableTransactionIncludesResponseAllOfMerchant**](TransactionIncludesResponseAllOfMerchant.md) |  | [optional] 
**RepeatingTransaction** | Pointer to [**NullableTransactionIncludesResponseAllOfRepeatingTransaction**](TransactionIncludesResponseAllOfRepeatingTransaction.md) |  | [optional] 

## Methods

### NewTransactionIncludesResponse

`func NewTransactionIncludesResponse() *TransactionIncludesResponse`

NewTransactionIncludesResponse instantiates a new TransactionIncludesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionIncludesResponseWithDefaults

`func NewTransactionIncludesResponseWithDefaults() *TransactionIncludesResponse`

NewTransactionIncludesResponseWithDefaults instantiates a new TransactionIncludesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *TransactionIncludesResponse) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *TransactionIncludesResponse) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *TransactionIncludesResponse) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.

### HasAccountGuid

`func (o *TransactionIncludesResponse) HasAccountGuid() bool`

HasAccountGuid returns a boolean if a field has been set.

### SetAccountGuidNil

`func (o *TransactionIncludesResponse) SetAccountGuidNil(b bool)`

 SetAccountGuidNil sets the value for AccountGuid to be an explicit nil

### UnsetAccountGuid
`func (o *TransactionIncludesResponse) UnsetAccountGuid()`

UnsetAccountGuid ensures that no value is present for AccountGuid, not even an explicit nil
### GetAccountId

`func (o *TransactionIncludesResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransactionIncludesResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransactionIncludesResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TransactionIncludesResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *TransactionIncludesResponse) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *TransactionIncludesResponse) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAmount

`func (o *TransactionIncludesResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionIncludesResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionIncludesResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionIncludesResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *TransactionIncludesResponse) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *TransactionIncludesResponse) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCategory

`func (o *TransactionIncludesResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TransactionIncludesResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TransactionIncludesResponse) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TransactionIncludesResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *TransactionIncludesResponse) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *TransactionIncludesResponse) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCategoryGuid

`func (o *TransactionIncludesResponse) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *TransactionIncludesResponse) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *TransactionIncludesResponse) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *TransactionIncludesResponse) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### SetCategoryGuidNil

`func (o *TransactionIncludesResponse) SetCategoryGuidNil(b bool)`

 SetCategoryGuidNil sets the value for CategoryGuid to be an explicit nil

### UnsetCategoryGuid
`func (o *TransactionIncludesResponse) UnsetCategoryGuid()`

UnsetCategoryGuid ensures that no value is present for CategoryGuid, not even an explicit nil
### GetCheckNumberString

`func (o *TransactionIncludesResponse) GetCheckNumberString() string`

GetCheckNumberString returns the CheckNumberString field if non-nil, zero value otherwise.

### GetCheckNumberStringOk

`func (o *TransactionIncludesResponse) GetCheckNumberStringOk() (*string, bool)`

GetCheckNumberStringOk returns a tuple with the CheckNumberString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumberString

`func (o *TransactionIncludesResponse) SetCheckNumberString(v string)`

SetCheckNumberString sets CheckNumberString field to given value.

### HasCheckNumberString

`func (o *TransactionIncludesResponse) HasCheckNumberString() bool`

HasCheckNumberString returns a boolean if a field has been set.

### SetCheckNumberStringNil

`func (o *TransactionIncludesResponse) SetCheckNumberStringNil(b bool)`

 SetCheckNumberStringNil sets the value for CheckNumberString to be an explicit nil

### UnsetCheckNumberString
`func (o *TransactionIncludesResponse) UnsetCheckNumberString()`

UnsetCheckNumberString ensures that no value is present for CheckNumberString, not even an explicit nil
### GetCreatedAt

`func (o *TransactionIncludesResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionIncludesResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionIncludesResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransactionIncludesResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TransactionIncludesResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TransactionIncludesResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCurrencyCode

`func (o *TransactionIncludesResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *TransactionIncludesResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *TransactionIncludesResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *TransactionIncludesResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *TransactionIncludesResponse) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *TransactionIncludesResponse) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDate

`func (o *TransactionIncludesResponse) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionIncludesResponse) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionIncludesResponse) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *TransactionIncludesResponse) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *TransactionIncludesResponse) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *TransactionIncludesResponse) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *TransactionIncludesResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionIncludesResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionIncludesResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionIncludesResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TransactionIncludesResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TransactionIncludesResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtendedTransactionType

`func (o *TransactionIncludesResponse) GetExtendedTransactionType() string`

GetExtendedTransactionType returns the ExtendedTransactionType field if non-nil, zero value otherwise.

### GetExtendedTransactionTypeOk

`func (o *TransactionIncludesResponse) GetExtendedTransactionTypeOk() (*string, bool)`

GetExtendedTransactionTypeOk returns a tuple with the ExtendedTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedTransactionType

`func (o *TransactionIncludesResponse) SetExtendedTransactionType(v string)`

SetExtendedTransactionType sets ExtendedTransactionType field to given value.

### HasExtendedTransactionType

`func (o *TransactionIncludesResponse) HasExtendedTransactionType() bool`

HasExtendedTransactionType returns a boolean if a field has been set.

### SetExtendedTransactionTypeNil

`func (o *TransactionIncludesResponse) SetExtendedTransactionTypeNil(b bool)`

 SetExtendedTransactionTypeNil sets the value for ExtendedTransactionType to be an explicit nil

### UnsetExtendedTransactionType
`func (o *TransactionIncludesResponse) UnsetExtendedTransactionType()`

UnsetExtendedTransactionType ensures that no value is present for ExtendedTransactionType, not even an explicit nil
### GetGuid

`func (o *TransactionIncludesResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *TransactionIncludesResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *TransactionIncludesResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *TransactionIncludesResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *TransactionIncludesResponse) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *TransactionIncludesResponse) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetId

`func (o *TransactionIncludesResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionIncludesResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionIncludesResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionIncludesResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TransactionIncludesResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TransactionIncludesResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsBillPay

`func (o *TransactionIncludesResponse) GetIsBillPay() bool`

GetIsBillPay returns the IsBillPay field if non-nil, zero value otherwise.

### GetIsBillPayOk

`func (o *TransactionIncludesResponse) GetIsBillPayOk() (*bool, bool)`

GetIsBillPayOk returns a tuple with the IsBillPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillPay

`func (o *TransactionIncludesResponse) SetIsBillPay(v bool)`

SetIsBillPay sets IsBillPay field to given value.

### HasIsBillPay

`func (o *TransactionIncludesResponse) HasIsBillPay() bool`

HasIsBillPay returns a boolean if a field has been set.

### SetIsBillPayNil

`func (o *TransactionIncludesResponse) SetIsBillPayNil(b bool)`

 SetIsBillPayNil sets the value for IsBillPay to be an explicit nil

### UnsetIsBillPay
`func (o *TransactionIncludesResponse) UnsetIsBillPay()`

UnsetIsBillPay ensures that no value is present for IsBillPay, not even an explicit nil
### GetIsDirectDeposit

`func (o *TransactionIncludesResponse) GetIsDirectDeposit() bool`

GetIsDirectDeposit returns the IsDirectDeposit field if non-nil, zero value otherwise.

### GetIsDirectDepositOk

`func (o *TransactionIncludesResponse) GetIsDirectDepositOk() (*bool, bool)`

GetIsDirectDepositOk returns a tuple with the IsDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectDeposit

`func (o *TransactionIncludesResponse) SetIsDirectDeposit(v bool)`

SetIsDirectDeposit sets IsDirectDeposit field to given value.

### HasIsDirectDeposit

`func (o *TransactionIncludesResponse) HasIsDirectDeposit() bool`

HasIsDirectDeposit returns a boolean if a field has been set.

### SetIsDirectDepositNil

`func (o *TransactionIncludesResponse) SetIsDirectDepositNil(b bool)`

 SetIsDirectDepositNil sets the value for IsDirectDeposit to be an explicit nil

### UnsetIsDirectDeposit
`func (o *TransactionIncludesResponse) UnsetIsDirectDeposit()`

UnsetIsDirectDeposit ensures that no value is present for IsDirectDeposit, not even an explicit nil
### GetIsExpense

`func (o *TransactionIncludesResponse) GetIsExpense() bool`

GetIsExpense returns the IsExpense field if non-nil, zero value otherwise.

### GetIsExpenseOk

`func (o *TransactionIncludesResponse) GetIsExpenseOk() (*bool, bool)`

GetIsExpenseOk returns a tuple with the IsExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpense

`func (o *TransactionIncludesResponse) SetIsExpense(v bool)`

SetIsExpense sets IsExpense field to given value.

### HasIsExpense

`func (o *TransactionIncludesResponse) HasIsExpense() bool`

HasIsExpense returns a boolean if a field has been set.

### SetIsExpenseNil

`func (o *TransactionIncludesResponse) SetIsExpenseNil(b bool)`

 SetIsExpenseNil sets the value for IsExpense to be an explicit nil

### UnsetIsExpense
`func (o *TransactionIncludesResponse) UnsetIsExpense()`

UnsetIsExpense ensures that no value is present for IsExpense, not even an explicit nil
### GetIsFee

`func (o *TransactionIncludesResponse) GetIsFee() bool`

GetIsFee returns the IsFee field if non-nil, zero value otherwise.

### GetIsFeeOk

`func (o *TransactionIncludesResponse) GetIsFeeOk() (*bool, bool)`

GetIsFeeOk returns a tuple with the IsFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFee

`func (o *TransactionIncludesResponse) SetIsFee(v bool)`

SetIsFee sets IsFee field to given value.

### HasIsFee

`func (o *TransactionIncludesResponse) HasIsFee() bool`

HasIsFee returns a boolean if a field has been set.

### SetIsFeeNil

`func (o *TransactionIncludesResponse) SetIsFeeNil(b bool)`

 SetIsFeeNil sets the value for IsFee to be an explicit nil

### UnsetIsFee
`func (o *TransactionIncludesResponse) UnsetIsFee()`

UnsetIsFee ensures that no value is present for IsFee, not even an explicit nil
### GetIsIncome

`func (o *TransactionIncludesResponse) GetIsIncome() bool`

GetIsIncome returns the IsIncome field if non-nil, zero value otherwise.

### GetIsIncomeOk

`func (o *TransactionIncludesResponse) GetIsIncomeOk() (*bool, bool)`

GetIsIncomeOk returns a tuple with the IsIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncome

`func (o *TransactionIncludesResponse) SetIsIncome(v bool)`

SetIsIncome sets IsIncome field to given value.

### HasIsIncome

`func (o *TransactionIncludesResponse) HasIsIncome() bool`

HasIsIncome returns a boolean if a field has been set.

### SetIsIncomeNil

`func (o *TransactionIncludesResponse) SetIsIncomeNil(b bool)`

 SetIsIncomeNil sets the value for IsIncome to be an explicit nil

### UnsetIsIncome
`func (o *TransactionIncludesResponse) UnsetIsIncome()`

UnsetIsIncome ensures that no value is present for IsIncome, not even an explicit nil
### GetIsInternational

`func (o *TransactionIncludesResponse) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *TransactionIncludesResponse) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *TransactionIncludesResponse) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *TransactionIncludesResponse) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### SetIsInternationalNil

`func (o *TransactionIncludesResponse) SetIsInternationalNil(b bool)`

 SetIsInternationalNil sets the value for IsInternational to be an explicit nil

### UnsetIsInternational
`func (o *TransactionIncludesResponse) UnsetIsInternational()`

UnsetIsInternational ensures that no value is present for IsInternational, not even an explicit nil
### GetIsManual

`func (o *TransactionIncludesResponse) GetIsManual() bool`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *TransactionIncludesResponse) GetIsManualOk() (*bool, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *TransactionIncludesResponse) SetIsManual(v bool)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *TransactionIncludesResponse) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### SetIsManualNil

`func (o *TransactionIncludesResponse) SetIsManualNil(b bool)`

 SetIsManualNil sets the value for IsManual to be an explicit nil

### UnsetIsManual
`func (o *TransactionIncludesResponse) UnsetIsManual()`

UnsetIsManual ensures that no value is present for IsManual, not even an explicit nil
### GetIsOverdraftFee

`func (o *TransactionIncludesResponse) GetIsOverdraftFee() bool`

GetIsOverdraftFee returns the IsOverdraftFee field if non-nil, zero value otherwise.

### GetIsOverdraftFeeOk

`func (o *TransactionIncludesResponse) GetIsOverdraftFeeOk() (*bool, bool)`

GetIsOverdraftFeeOk returns a tuple with the IsOverdraftFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdraftFee

`func (o *TransactionIncludesResponse) SetIsOverdraftFee(v bool)`

SetIsOverdraftFee sets IsOverdraftFee field to given value.

### HasIsOverdraftFee

`func (o *TransactionIncludesResponse) HasIsOverdraftFee() bool`

HasIsOverdraftFee returns a boolean if a field has been set.

### SetIsOverdraftFeeNil

`func (o *TransactionIncludesResponse) SetIsOverdraftFeeNil(b bool)`

 SetIsOverdraftFeeNil sets the value for IsOverdraftFee to be an explicit nil

### UnsetIsOverdraftFee
`func (o *TransactionIncludesResponse) UnsetIsOverdraftFee()`

UnsetIsOverdraftFee ensures that no value is present for IsOverdraftFee, not even an explicit nil
### GetIsPayrollAdvance

`func (o *TransactionIncludesResponse) GetIsPayrollAdvance() bool`

GetIsPayrollAdvance returns the IsPayrollAdvance field if non-nil, zero value otherwise.

### GetIsPayrollAdvanceOk

`func (o *TransactionIncludesResponse) GetIsPayrollAdvanceOk() (*bool, bool)`

GetIsPayrollAdvanceOk returns a tuple with the IsPayrollAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPayrollAdvance

`func (o *TransactionIncludesResponse) SetIsPayrollAdvance(v bool)`

SetIsPayrollAdvance sets IsPayrollAdvance field to given value.

### HasIsPayrollAdvance

`func (o *TransactionIncludesResponse) HasIsPayrollAdvance() bool`

HasIsPayrollAdvance returns a boolean if a field has been set.

### SetIsPayrollAdvanceNil

`func (o *TransactionIncludesResponse) SetIsPayrollAdvanceNil(b bool)`

 SetIsPayrollAdvanceNil sets the value for IsPayrollAdvance to be an explicit nil

### UnsetIsPayrollAdvance
`func (o *TransactionIncludesResponse) UnsetIsPayrollAdvance()`

UnsetIsPayrollAdvance ensures that no value is present for IsPayrollAdvance, not even an explicit nil
### GetIsRecurring

`func (o *TransactionIncludesResponse) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *TransactionIncludesResponse) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *TransactionIncludesResponse) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.

### HasIsRecurring

`func (o *TransactionIncludesResponse) HasIsRecurring() bool`

HasIsRecurring returns a boolean if a field has been set.

### SetIsRecurringNil

`func (o *TransactionIncludesResponse) SetIsRecurringNil(b bool)`

 SetIsRecurringNil sets the value for IsRecurring to be an explicit nil

### UnsetIsRecurring
`func (o *TransactionIncludesResponse) UnsetIsRecurring()`

UnsetIsRecurring ensures that no value is present for IsRecurring, not even an explicit nil
### GetIsSubscription

`func (o *TransactionIncludesResponse) GetIsSubscription() bool`

GetIsSubscription returns the IsSubscription field if non-nil, zero value otherwise.

### GetIsSubscriptionOk

`func (o *TransactionIncludesResponse) GetIsSubscriptionOk() (*bool, bool)`

GetIsSubscriptionOk returns a tuple with the IsSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscription

`func (o *TransactionIncludesResponse) SetIsSubscription(v bool)`

SetIsSubscription sets IsSubscription field to given value.

### HasIsSubscription

`func (o *TransactionIncludesResponse) HasIsSubscription() bool`

HasIsSubscription returns a boolean if a field has been set.

### SetIsSubscriptionNil

`func (o *TransactionIncludesResponse) SetIsSubscriptionNil(b bool)`

 SetIsSubscriptionNil sets the value for IsSubscription to be an explicit nil

### UnsetIsSubscription
`func (o *TransactionIncludesResponse) UnsetIsSubscription()`

UnsetIsSubscription ensures that no value is present for IsSubscription, not even an explicit nil
### GetLatitude

`func (o *TransactionIncludesResponse) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *TransactionIncludesResponse) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *TransactionIncludesResponse) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *TransactionIncludesResponse) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *TransactionIncludesResponse) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *TransactionIncludesResponse) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLocalizedDescription

`func (o *TransactionIncludesResponse) GetLocalizedDescription() string`

GetLocalizedDescription returns the LocalizedDescription field if non-nil, zero value otherwise.

### GetLocalizedDescriptionOk

`func (o *TransactionIncludesResponse) GetLocalizedDescriptionOk() (*string, bool)`

GetLocalizedDescriptionOk returns a tuple with the LocalizedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedDescription

`func (o *TransactionIncludesResponse) SetLocalizedDescription(v string)`

SetLocalizedDescription sets LocalizedDescription field to given value.

### HasLocalizedDescription

`func (o *TransactionIncludesResponse) HasLocalizedDescription() bool`

HasLocalizedDescription returns a boolean if a field has been set.

### SetLocalizedDescriptionNil

`func (o *TransactionIncludesResponse) SetLocalizedDescriptionNil(b bool)`

 SetLocalizedDescriptionNil sets the value for LocalizedDescription to be an explicit nil

### UnsetLocalizedDescription
`func (o *TransactionIncludesResponse) UnsetLocalizedDescription()`

UnsetLocalizedDescription ensures that no value is present for LocalizedDescription, not even an explicit nil
### GetLocalizedMemo

`func (o *TransactionIncludesResponse) GetLocalizedMemo() string`

GetLocalizedMemo returns the LocalizedMemo field if non-nil, zero value otherwise.

### GetLocalizedMemoOk

`func (o *TransactionIncludesResponse) GetLocalizedMemoOk() (*string, bool)`

GetLocalizedMemoOk returns a tuple with the LocalizedMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedMemo

`func (o *TransactionIncludesResponse) SetLocalizedMemo(v string)`

SetLocalizedMemo sets LocalizedMemo field to given value.

### HasLocalizedMemo

`func (o *TransactionIncludesResponse) HasLocalizedMemo() bool`

HasLocalizedMemo returns a boolean if a field has been set.

### SetLocalizedMemoNil

`func (o *TransactionIncludesResponse) SetLocalizedMemoNil(b bool)`

 SetLocalizedMemoNil sets the value for LocalizedMemo to be an explicit nil

### UnsetLocalizedMemo
`func (o *TransactionIncludesResponse) UnsetLocalizedMemo()`

UnsetLocalizedMemo ensures that no value is present for LocalizedMemo, not even an explicit nil
### GetLongitude

`func (o *TransactionIncludesResponse) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *TransactionIncludesResponse) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *TransactionIncludesResponse) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *TransactionIncludesResponse) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *TransactionIncludesResponse) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *TransactionIncludesResponse) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetMemberGuid

`func (o *TransactionIncludesResponse) GetMemberGuid() string`

GetMemberGuid returns the MemberGuid field if non-nil, zero value otherwise.

### GetMemberGuidOk

`func (o *TransactionIncludesResponse) GetMemberGuidOk() (*string, bool)`

GetMemberGuidOk returns a tuple with the MemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGuid

`func (o *TransactionIncludesResponse) SetMemberGuid(v string)`

SetMemberGuid sets MemberGuid field to given value.

### HasMemberGuid

`func (o *TransactionIncludesResponse) HasMemberGuid() bool`

HasMemberGuid returns a boolean if a field has been set.

### SetMemberGuidNil

`func (o *TransactionIncludesResponse) SetMemberGuidNil(b bool)`

 SetMemberGuidNil sets the value for MemberGuid to be an explicit nil

### UnsetMemberGuid
`func (o *TransactionIncludesResponse) UnsetMemberGuid()`

UnsetMemberGuid ensures that no value is present for MemberGuid, not even an explicit nil
### GetMemberIsManagedByUser

`func (o *TransactionIncludesResponse) GetMemberIsManagedByUser() bool`

GetMemberIsManagedByUser returns the MemberIsManagedByUser field if non-nil, zero value otherwise.

### GetMemberIsManagedByUserOk

`func (o *TransactionIncludesResponse) GetMemberIsManagedByUserOk() (*bool, bool)`

GetMemberIsManagedByUserOk returns a tuple with the MemberIsManagedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIsManagedByUser

`func (o *TransactionIncludesResponse) SetMemberIsManagedByUser(v bool)`

SetMemberIsManagedByUser sets MemberIsManagedByUser field to given value.

### HasMemberIsManagedByUser

`func (o *TransactionIncludesResponse) HasMemberIsManagedByUser() bool`

HasMemberIsManagedByUser returns a boolean if a field has been set.

### SetMemberIsManagedByUserNil

`func (o *TransactionIncludesResponse) SetMemberIsManagedByUserNil(b bool)`

 SetMemberIsManagedByUserNil sets the value for MemberIsManagedByUser to be an explicit nil

### UnsetMemberIsManagedByUser
`func (o *TransactionIncludesResponse) UnsetMemberIsManagedByUser()`

UnsetMemberIsManagedByUser ensures that no value is present for MemberIsManagedByUser, not even an explicit nil
### GetMemo

`func (o *TransactionIncludesResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionIncludesResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionIncludesResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionIncludesResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *TransactionIncludesResponse) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *TransactionIncludesResponse) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetMerchantCategoryCode

`func (o *TransactionIncludesResponse) GetMerchantCategoryCode() int32`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *TransactionIncludesResponse) GetMerchantCategoryCodeOk() (*int32, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *TransactionIncludesResponse) SetMerchantCategoryCode(v int32)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *TransactionIncludesResponse) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### SetMerchantCategoryCodeNil

`func (o *TransactionIncludesResponse) SetMerchantCategoryCodeNil(b bool)`

 SetMerchantCategoryCodeNil sets the value for MerchantCategoryCode to be an explicit nil

### UnsetMerchantCategoryCode
`func (o *TransactionIncludesResponse) UnsetMerchantCategoryCode()`

UnsetMerchantCategoryCode ensures that no value is present for MerchantCategoryCode, not even an explicit nil
### GetMerchantGuid

`func (o *TransactionIncludesResponse) GetMerchantGuid() string`

GetMerchantGuid returns the MerchantGuid field if non-nil, zero value otherwise.

### GetMerchantGuidOk

`func (o *TransactionIncludesResponse) GetMerchantGuidOk() (*string, bool)`

GetMerchantGuidOk returns a tuple with the MerchantGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGuid

`func (o *TransactionIncludesResponse) SetMerchantGuid(v string)`

SetMerchantGuid sets MerchantGuid field to given value.

### HasMerchantGuid

`func (o *TransactionIncludesResponse) HasMerchantGuid() bool`

HasMerchantGuid returns a boolean if a field has been set.

### SetMerchantGuidNil

`func (o *TransactionIncludesResponse) SetMerchantGuidNil(b bool)`

 SetMerchantGuidNil sets the value for MerchantGuid to be an explicit nil

### UnsetMerchantGuid
`func (o *TransactionIncludesResponse) UnsetMerchantGuid()`

UnsetMerchantGuid ensures that no value is present for MerchantGuid, not even an explicit nil
### GetMerchantLocationGuid

`func (o *TransactionIncludesResponse) GetMerchantLocationGuid() string`

GetMerchantLocationGuid returns the MerchantLocationGuid field if non-nil, zero value otherwise.

### GetMerchantLocationGuidOk

`func (o *TransactionIncludesResponse) GetMerchantLocationGuidOk() (*string, bool)`

GetMerchantLocationGuidOk returns a tuple with the MerchantLocationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantLocationGuid

`func (o *TransactionIncludesResponse) SetMerchantLocationGuid(v string)`

SetMerchantLocationGuid sets MerchantLocationGuid field to given value.

### HasMerchantLocationGuid

`func (o *TransactionIncludesResponse) HasMerchantLocationGuid() bool`

HasMerchantLocationGuid returns a boolean if a field has been set.

### SetMerchantLocationGuidNil

`func (o *TransactionIncludesResponse) SetMerchantLocationGuidNil(b bool)`

 SetMerchantLocationGuidNil sets the value for MerchantLocationGuid to be an explicit nil

### UnsetMerchantLocationGuid
`func (o *TransactionIncludesResponse) UnsetMerchantLocationGuid()`

UnsetMerchantLocationGuid ensures that no value is present for MerchantLocationGuid, not even an explicit nil
### GetMetadata

`func (o *TransactionIncludesResponse) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransactionIncludesResponse) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransactionIncludesResponse) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransactionIncludesResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *TransactionIncludesResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TransactionIncludesResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOriginalDescription

`func (o *TransactionIncludesResponse) GetOriginalDescription() string`

GetOriginalDescription returns the OriginalDescription field if non-nil, zero value otherwise.

### GetOriginalDescriptionOk

`func (o *TransactionIncludesResponse) GetOriginalDescriptionOk() (*string, bool)`

GetOriginalDescriptionOk returns a tuple with the OriginalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDescription

`func (o *TransactionIncludesResponse) SetOriginalDescription(v string)`

SetOriginalDescription sets OriginalDescription field to given value.

### HasOriginalDescription

`func (o *TransactionIncludesResponse) HasOriginalDescription() bool`

HasOriginalDescription returns a boolean if a field has been set.

### SetOriginalDescriptionNil

`func (o *TransactionIncludesResponse) SetOriginalDescriptionNil(b bool)`

 SetOriginalDescriptionNil sets the value for OriginalDescription to be an explicit nil

### UnsetOriginalDescription
`func (o *TransactionIncludesResponse) UnsetOriginalDescription()`

UnsetOriginalDescription ensures that no value is present for OriginalDescription, not even an explicit nil
### GetPostedAt

`func (o *TransactionIncludesResponse) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *TransactionIncludesResponse) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *TransactionIncludesResponse) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *TransactionIncludesResponse) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### SetPostedAtNil

`func (o *TransactionIncludesResponse) SetPostedAtNil(b bool)`

 SetPostedAtNil sets the value for PostedAt to be an explicit nil

### UnsetPostedAt
`func (o *TransactionIncludesResponse) UnsetPostedAt()`

UnsetPostedAt ensures that no value is present for PostedAt, not even an explicit nil
### GetStatus

`func (o *TransactionIncludesResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionIncludesResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionIncludesResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionIncludesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TransactionIncludesResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TransactionIncludesResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTopLevelCategory

`func (o *TransactionIncludesResponse) GetTopLevelCategory() string`

GetTopLevelCategory returns the TopLevelCategory field if non-nil, zero value otherwise.

### GetTopLevelCategoryOk

`func (o *TransactionIncludesResponse) GetTopLevelCategoryOk() (*string, bool)`

GetTopLevelCategoryOk returns a tuple with the TopLevelCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevelCategory

`func (o *TransactionIncludesResponse) SetTopLevelCategory(v string)`

SetTopLevelCategory sets TopLevelCategory field to given value.

### HasTopLevelCategory

`func (o *TransactionIncludesResponse) HasTopLevelCategory() bool`

HasTopLevelCategory returns a boolean if a field has been set.

### SetTopLevelCategoryNil

`func (o *TransactionIncludesResponse) SetTopLevelCategoryNil(b bool)`

 SetTopLevelCategoryNil sets the value for TopLevelCategory to be an explicit nil

### UnsetTopLevelCategory
`func (o *TransactionIncludesResponse) UnsetTopLevelCategory()`

UnsetTopLevelCategory ensures that no value is present for TopLevelCategory, not even an explicit nil
### GetTransactedAt

`func (o *TransactionIncludesResponse) GetTransactedAt() string`

GetTransactedAt returns the TransactedAt field if non-nil, zero value otherwise.

### GetTransactedAtOk

`func (o *TransactionIncludesResponse) GetTransactedAtOk() (*string, bool)`

GetTransactedAtOk returns a tuple with the TransactedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactedAt

`func (o *TransactionIncludesResponse) SetTransactedAt(v string)`

SetTransactedAt sets TransactedAt field to given value.

### HasTransactedAt

`func (o *TransactionIncludesResponse) HasTransactedAt() bool`

HasTransactedAt returns a boolean if a field has been set.

### SetTransactedAtNil

`func (o *TransactionIncludesResponse) SetTransactedAtNil(b bool)`

 SetTransactedAtNil sets the value for TransactedAt to be an explicit nil

### UnsetTransactedAt
`func (o *TransactionIncludesResponse) UnsetTransactedAt()`

UnsetTransactedAt ensures that no value is present for TransactedAt, not even an explicit nil
### GetType

`func (o *TransactionIncludesResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionIncludesResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionIncludesResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionIncludesResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TransactionIncludesResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TransactionIncludesResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUpdatedAt

`func (o *TransactionIncludesResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransactionIncludesResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransactionIncludesResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TransactionIncludesResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TransactionIncludesResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TransactionIncludesResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserGuid

`func (o *TransactionIncludesResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *TransactionIncludesResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *TransactionIncludesResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *TransactionIncludesResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### SetUserGuidNil

`func (o *TransactionIncludesResponse) SetUserGuidNil(b bool)`

 SetUserGuidNil sets the value for UserGuid to be an explicit nil

### UnsetUserGuid
`func (o *TransactionIncludesResponse) UnsetUserGuid()`

UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
### GetUserId

`func (o *TransactionIncludesResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TransactionIncludesResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TransactionIncludesResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TransactionIncludesResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *TransactionIncludesResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *TransactionIncludesResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetClassification

`func (o *TransactionIncludesResponse) GetClassification() TransactionIncludesResponseAllOfClassification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *TransactionIncludesResponse) GetClassificationOk() (*TransactionIncludesResponseAllOfClassification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *TransactionIncludesResponse) SetClassification(v TransactionIncludesResponseAllOfClassification)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *TransactionIncludesResponse) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *TransactionIncludesResponse) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *TransactionIncludesResponse) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetGeolocation

`func (o *TransactionIncludesResponse) GetGeolocation() TransactionIncludesResponseAllOfGeolocation`

GetGeolocation returns the Geolocation field if non-nil, zero value otherwise.

### GetGeolocationOk

`func (o *TransactionIncludesResponse) GetGeolocationOk() (*TransactionIncludesResponseAllOfGeolocation, bool)`

GetGeolocationOk returns a tuple with the Geolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocation

`func (o *TransactionIncludesResponse) SetGeolocation(v TransactionIncludesResponseAllOfGeolocation)`

SetGeolocation sets Geolocation field to given value.

### HasGeolocation

`func (o *TransactionIncludesResponse) HasGeolocation() bool`

HasGeolocation returns a boolean if a field has been set.

### SetGeolocationNil

`func (o *TransactionIncludesResponse) SetGeolocationNil(b bool)`

 SetGeolocationNil sets the value for Geolocation to be an explicit nil

### UnsetGeolocation
`func (o *TransactionIncludesResponse) UnsetGeolocation()`

UnsetGeolocation ensures that no value is present for Geolocation, not even an explicit nil
### GetMerchant

`func (o *TransactionIncludesResponse) GetMerchant() TransactionIncludesResponseAllOfMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransactionIncludesResponse) GetMerchantOk() (*TransactionIncludesResponseAllOfMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransactionIncludesResponse) SetMerchant(v TransactionIncludesResponseAllOfMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *TransactionIncludesResponse) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### SetMerchantNil

`func (o *TransactionIncludesResponse) SetMerchantNil(b bool)`

 SetMerchantNil sets the value for Merchant to be an explicit nil

### UnsetMerchant
`func (o *TransactionIncludesResponse) UnsetMerchant()`

UnsetMerchant ensures that no value is present for Merchant, not even an explicit nil
### GetRepeatingTransaction

`func (o *TransactionIncludesResponse) GetRepeatingTransaction() TransactionIncludesResponseAllOfRepeatingTransaction`

GetRepeatingTransaction returns the RepeatingTransaction field if non-nil, zero value otherwise.

### GetRepeatingTransactionOk

`func (o *TransactionIncludesResponse) GetRepeatingTransactionOk() (*TransactionIncludesResponseAllOfRepeatingTransaction, bool)`

GetRepeatingTransactionOk returns a tuple with the RepeatingTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatingTransaction

`func (o *TransactionIncludesResponse) SetRepeatingTransaction(v TransactionIncludesResponseAllOfRepeatingTransaction)`

SetRepeatingTransaction sets RepeatingTransaction field to given value.

### HasRepeatingTransaction

`func (o *TransactionIncludesResponse) HasRepeatingTransaction() bool`

HasRepeatingTransaction returns a boolean if a field has been set.

### SetRepeatingTransactionNil

`func (o *TransactionIncludesResponse) SetRepeatingTransactionNil(b bool)`

 SetRepeatingTransactionNil sets the value for RepeatingTransaction to be an explicit nil

### UnsetRepeatingTransaction
`func (o *TransactionIncludesResponse) UnsetRepeatingTransaction()`

UnsetRepeatingTransaction ensures that no value is present for RepeatingTransaction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


