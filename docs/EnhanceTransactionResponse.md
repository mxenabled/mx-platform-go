# EnhanceTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat32** |  | [optional] 
**CategorizedBy** | Pointer to **NullableInt32** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**CategoryGuid** | Pointer to **NullableString** |  | [optional] 
**DescribedBy** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExtendedTransactionType** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**IsBillPay** | Pointer to **NullableBool** |  | [optional] 
**IsDirectDeposit** | Pointer to **NullableBool** |  | [optional] 
**IsExpense** | Pointer to **NullableBool** |  | [optional] 
**IsFee** | Pointer to **NullableBool** |  | [optional] 
**IsIncome** | Pointer to **NullableBool** |  | [optional] 
**IsInternational** | Pointer to **NullableBool** |  | [optional] 
**IsOverdraftFee** | Pointer to **NullableBool** |  | [optional] 
**IsPayrollAdvance** | Pointer to **NullableBool** |  | [optional] 
**IsSubscription** | Pointer to **NullableBool** |  | [optional] 
**Memo** | Pointer to **NullableString** |  | [optional] 
**MerchantCategoryCode** | Pointer to **NullableInt32** |  | [optional] 
**MerchantGuid** | Pointer to **NullableString** |  | [optional] 
**MerchantLocationGuid** | Pointer to **NullableString** |  | [optional] 
**OriginalDescription** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEnhanceTransactionResponse

`func NewEnhanceTransactionResponse() *EnhanceTransactionResponse`

NewEnhanceTransactionResponse instantiates a new EnhanceTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhanceTransactionResponseWithDefaults

`func NewEnhanceTransactionResponseWithDefaults() *EnhanceTransactionResponse`

NewEnhanceTransactionResponseWithDefaults instantiates a new EnhanceTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EnhanceTransactionResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnhanceTransactionResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnhanceTransactionResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EnhanceTransactionResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *EnhanceTransactionResponse) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *EnhanceTransactionResponse) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCategorizedBy

`func (o *EnhanceTransactionResponse) GetCategorizedBy() int32`

GetCategorizedBy returns the CategorizedBy field if non-nil, zero value otherwise.

### GetCategorizedByOk

`func (o *EnhanceTransactionResponse) GetCategorizedByOk() (*int32, bool)`

GetCategorizedByOk returns a tuple with the CategorizedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorizedBy

`func (o *EnhanceTransactionResponse) SetCategorizedBy(v int32)`

SetCategorizedBy sets CategorizedBy field to given value.

### HasCategorizedBy

`func (o *EnhanceTransactionResponse) HasCategorizedBy() bool`

HasCategorizedBy returns a boolean if a field has been set.

### SetCategorizedByNil

`func (o *EnhanceTransactionResponse) SetCategorizedByNil(b bool)`

 SetCategorizedByNil sets the value for CategorizedBy to be an explicit nil

### UnsetCategorizedBy
`func (o *EnhanceTransactionResponse) UnsetCategorizedBy()`

UnsetCategorizedBy ensures that no value is present for CategorizedBy, not even an explicit nil
### GetCategory

`func (o *EnhanceTransactionResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EnhanceTransactionResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EnhanceTransactionResponse) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EnhanceTransactionResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *EnhanceTransactionResponse) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *EnhanceTransactionResponse) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCategoryGuid

`func (o *EnhanceTransactionResponse) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *EnhanceTransactionResponse) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *EnhanceTransactionResponse) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *EnhanceTransactionResponse) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### SetCategoryGuidNil

`func (o *EnhanceTransactionResponse) SetCategoryGuidNil(b bool)`

 SetCategoryGuidNil sets the value for CategoryGuid to be an explicit nil

### UnsetCategoryGuid
`func (o *EnhanceTransactionResponse) UnsetCategoryGuid()`

UnsetCategoryGuid ensures that no value is present for CategoryGuid, not even an explicit nil
### GetDescribedBy

`func (o *EnhanceTransactionResponse) GetDescribedBy() int32`

GetDescribedBy returns the DescribedBy field if non-nil, zero value otherwise.

### GetDescribedByOk

`func (o *EnhanceTransactionResponse) GetDescribedByOk() (*int32, bool)`

GetDescribedByOk returns a tuple with the DescribedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribedBy

`func (o *EnhanceTransactionResponse) SetDescribedBy(v int32)`

SetDescribedBy sets DescribedBy field to given value.

### HasDescribedBy

`func (o *EnhanceTransactionResponse) HasDescribedBy() bool`

HasDescribedBy returns a boolean if a field has been set.

### SetDescribedByNil

`func (o *EnhanceTransactionResponse) SetDescribedByNil(b bool)`

 SetDescribedByNil sets the value for DescribedBy to be an explicit nil

### UnsetDescribedBy
`func (o *EnhanceTransactionResponse) UnsetDescribedBy()`

UnsetDescribedBy ensures that no value is present for DescribedBy, not even an explicit nil
### GetDescription

`func (o *EnhanceTransactionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnhanceTransactionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnhanceTransactionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnhanceTransactionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EnhanceTransactionResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EnhanceTransactionResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtendedTransactionType

`func (o *EnhanceTransactionResponse) GetExtendedTransactionType() string`

GetExtendedTransactionType returns the ExtendedTransactionType field if non-nil, zero value otherwise.

### GetExtendedTransactionTypeOk

`func (o *EnhanceTransactionResponse) GetExtendedTransactionTypeOk() (*string, bool)`

GetExtendedTransactionTypeOk returns a tuple with the ExtendedTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedTransactionType

`func (o *EnhanceTransactionResponse) SetExtendedTransactionType(v string)`

SetExtendedTransactionType sets ExtendedTransactionType field to given value.

### HasExtendedTransactionType

`func (o *EnhanceTransactionResponse) HasExtendedTransactionType() bool`

HasExtendedTransactionType returns a boolean if a field has been set.

### SetExtendedTransactionTypeNil

`func (o *EnhanceTransactionResponse) SetExtendedTransactionTypeNil(b bool)`

 SetExtendedTransactionTypeNil sets the value for ExtendedTransactionType to be an explicit nil

### UnsetExtendedTransactionType
`func (o *EnhanceTransactionResponse) UnsetExtendedTransactionType()`

UnsetExtendedTransactionType ensures that no value is present for ExtendedTransactionType, not even an explicit nil
### GetId

`func (o *EnhanceTransactionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnhanceTransactionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnhanceTransactionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnhanceTransactionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EnhanceTransactionResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EnhanceTransactionResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsBillPay

`func (o *EnhanceTransactionResponse) GetIsBillPay() bool`

GetIsBillPay returns the IsBillPay field if non-nil, zero value otherwise.

### GetIsBillPayOk

`func (o *EnhanceTransactionResponse) GetIsBillPayOk() (*bool, bool)`

GetIsBillPayOk returns a tuple with the IsBillPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillPay

`func (o *EnhanceTransactionResponse) SetIsBillPay(v bool)`

SetIsBillPay sets IsBillPay field to given value.

### HasIsBillPay

`func (o *EnhanceTransactionResponse) HasIsBillPay() bool`

HasIsBillPay returns a boolean if a field has been set.

### SetIsBillPayNil

`func (o *EnhanceTransactionResponse) SetIsBillPayNil(b bool)`

 SetIsBillPayNil sets the value for IsBillPay to be an explicit nil

### UnsetIsBillPay
`func (o *EnhanceTransactionResponse) UnsetIsBillPay()`

UnsetIsBillPay ensures that no value is present for IsBillPay, not even an explicit nil
### GetIsDirectDeposit

`func (o *EnhanceTransactionResponse) GetIsDirectDeposit() bool`

GetIsDirectDeposit returns the IsDirectDeposit field if non-nil, zero value otherwise.

### GetIsDirectDepositOk

`func (o *EnhanceTransactionResponse) GetIsDirectDepositOk() (*bool, bool)`

GetIsDirectDepositOk returns a tuple with the IsDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectDeposit

`func (o *EnhanceTransactionResponse) SetIsDirectDeposit(v bool)`

SetIsDirectDeposit sets IsDirectDeposit field to given value.

### HasIsDirectDeposit

`func (o *EnhanceTransactionResponse) HasIsDirectDeposit() bool`

HasIsDirectDeposit returns a boolean if a field has been set.

### SetIsDirectDepositNil

`func (o *EnhanceTransactionResponse) SetIsDirectDepositNil(b bool)`

 SetIsDirectDepositNil sets the value for IsDirectDeposit to be an explicit nil

### UnsetIsDirectDeposit
`func (o *EnhanceTransactionResponse) UnsetIsDirectDeposit()`

UnsetIsDirectDeposit ensures that no value is present for IsDirectDeposit, not even an explicit nil
### GetIsExpense

`func (o *EnhanceTransactionResponse) GetIsExpense() bool`

GetIsExpense returns the IsExpense field if non-nil, zero value otherwise.

### GetIsExpenseOk

`func (o *EnhanceTransactionResponse) GetIsExpenseOk() (*bool, bool)`

GetIsExpenseOk returns a tuple with the IsExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpense

`func (o *EnhanceTransactionResponse) SetIsExpense(v bool)`

SetIsExpense sets IsExpense field to given value.

### HasIsExpense

`func (o *EnhanceTransactionResponse) HasIsExpense() bool`

HasIsExpense returns a boolean if a field has been set.

### SetIsExpenseNil

`func (o *EnhanceTransactionResponse) SetIsExpenseNil(b bool)`

 SetIsExpenseNil sets the value for IsExpense to be an explicit nil

### UnsetIsExpense
`func (o *EnhanceTransactionResponse) UnsetIsExpense()`

UnsetIsExpense ensures that no value is present for IsExpense, not even an explicit nil
### GetIsFee

`func (o *EnhanceTransactionResponse) GetIsFee() bool`

GetIsFee returns the IsFee field if non-nil, zero value otherwise.

### GetIsFeeOk

`func (o *EnhanceTransactionResponse) GetIsFeeOk() (*bool, bool)`

GetIsFeeOk returns a tuple with the IsFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFee

`func (o *EnhanceTransactionResponse) SetIsFee(v bool)`

SetIsFee sets IsFee field to given value.

### HasIsFee

`func (o *EnhanceTransactionResponse) HasIsFee() bool`

HasIsFee returns a boolean if a field has been set.

### SetIsFeeNil

`func (o *EnhanceTransactionResponse) SetIsFeeNil(b bool)`

 SetIsFeeNil sets the value for IsFee to be an explicit nil

### UnsetIsFee
`func (o *EnhanceTransactionResponse) UnsetIsFee()`

UnsetIsFee ensures that no value is present for IsFee, not even an explicit nil
### GetIsIncome

`func (o *EnhanceTransactionResponse) GetIsIncome() bool`

GetIsIncome returns the IsIncome field if non-nil, zero value otherwise.

### GetIsIncomeOk

`func (o *EnhanceTransactionResponse) GetIsIncomeOk() (*bool, bool)`

GetIsIncomeOk returns a tuple with the IsIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncome

`func (o *EnhanceTransactionResponse) SetIsIncome(v bool)`

SetIsIncome sets IsIncome field to given value.

### HasIsIncome

`func (o *EnhanceTransactionResponse) HasIsIncome() bool`

HasIsIncome returns a boolean if a field has been set.

### SetIsIncomeNil

`func (o *EnhanceTransactionResponse) SetIsIncomeNil(b bool)`

 SetIsIncomeNil sets the value for IsIncome to be an explicit nil

### UnsetIsIncome
`func (o *EnhanceTransactionResponse) UnsetIsIncome()`

UnsetIsIncome ensures that no value is present for IsIncome, not even an explicit nil
### GetIsInternational

`func (o *EnhanceTransactionResponse) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *EnhanceTransactionResponse) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *EnhanceTransactionResponse) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *EnhanceTransactionResponse) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### SetIsInternationalNil

`func (o *EnhanceTransactionResponse) SetIsInternationalNil(b bool)`

 SetIsInternationalNil sets the value for IsInternational to be an explicit nil

### UnsetIsInternational
`func (o *EnhanceTransactionResponse) UnsetIsInternational()`

UnsetIsInternational ensures that no value is present for IsInternational, not even an explicit nil
### GetIsOverdraftFee

`func (o *EnhanceTransactionResponse) GetIsOverdraftFee() bool`

GetIsOverdraftFee returns the IsOverdraftFee field if non-nil, zero value otherwise.

### GetIsOverdraftFeeOk

`func (o *EnhanceTransactionResponse) GetIsOverdraftFeeOk() (*bool, bool)`

GetIsOverdraftFeeOk returns a tuple with the IsOverdraftFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdraftFee

`func (o *EnhanceTransactionResponse) SetIsOverdraftFee(v bool)`

SetIsOverdraftFee sets IsOverdraftFee field to given value.

### HasIsOverdraftFee

`func (o *EnhanceTransactionResponse) HasIsOverdraftFee() bool`

HasIsOverdraftFee returns a boolean if a field has been set.

### SetIsOverdraftFeeNil

`func (o *EnhanceTransactionResponse) SetIsOverdraftFeeNil(b bool)`

 SetIsOverdraftFeeNil sets the value for IsOverdraftFee to be an explicit nil

### UnsetIsOverdraftFee
`func (o *EnhanceTransactionResponse) UnsetIsOverdraftFee()`

UnsetIsOverdraftFee ensures that no value is present for IsOverdraftFee, not even an explicit nil
### GetIsPayrollAdvance

`func (o *EnhanceTransactionResponse) GetIsPayrollAdvance() bool`

GetIsPayrollAdvance returns the IsPayrollAdvance field if non-nil, zero value otherwise.

### GetIsPayrollAdvanceOk

`func (o *EnhanceTransactionResponse) GetIsPayrollAdvanceOk() (*bool, bool)`

GetIsPayrollAdvanceOk returns a tuple with the IsPayrollAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPayrollAdvance

`func (o *EnhanceTransactionResponse) SetIsPayrollAdvance(v bool)`

SetIsPayrollAdvance sets IsPayrollAdvance field to given value.

### HasIsPayrollAdvance

`func (o *EnhanceTransactionResponse) HasIsPayrollAdvance() bool`

HasIsPayrollAdvance returns a boolean if a field has been set.

### SetIsPayrollAdvanceNil

`func (o *EnhanceTransactionResponse) SetIsPayrollAdvanceNil(b bool)`

 SetIsPayrollAdvanceNil sets the value for IsPayrollAdvance to be an explicit nil

### UnsetIsPayrollAdvance
`func (o *EnhanceTransactionResponse) UnsetIsPayrollAdvance()`

UnsetIsPayrollAdvance ensures that no value is present for IsPayrollAdvance, not even an explicit nil
### GetIsSubscription

`func (o *EnhanceTransactionResponse) GetIsSubscription() bool`

GetIsSubscription returns the IsSubscription field if non-nil, zero value otherwise.

### GetIsSubscriptionOk

`func (o *EnhanceTransactionResponse) GetIsSubscriptionOk() (*bool, bool)`

GetIsSubscriptionOk returns a tuple with the IsSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscription

`func (o *EnhanceTransactionResponse) SetIsSubscription(v bool)`

SetIsSubscription sets IsSubscription field to given value.

### HasIsSubscription

`func (o *EnhanceTransactionResponse) HasIsSubscription() bool`

HasIsSubscription returns a boolean if a field has been set.

### SetIsSubscriptionNil

`func (o *EnhanceTransactionResponse) SetIsSubscriptionNil(b bool)`

 SetIsSubscriptionNil sets the value for IsSubscription to be an explicit nil

### UnsetIsSubscription
`func (o *EnhanceTransactionResponse) UnsetIsSubscription()`

UnsetIsSubscription ensures that no value is present for IsSubscription, not even an explicit nil
### GetMemo

`func (o *EnhanceTransactionResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *EnhanceTransactionResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *EnhanceTransactionResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *EnhanceTransactionResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *EnhanceTransactionResponse) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *EnhanceTransactionResponse) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetMerchantCategoryCode

`func (o *EnhanceTransactionResponse) GetMerchantCategoryCode() int32`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *EnhanceTransactionResponse) GetMerchantCategoryCodeOk() (*int32, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *EnhanceTransactionResponse) SetMerchantCategoryCode(v int32)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *EnhanceTransactionResponse) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### SetMerchantCategoryCodeNil

`func (o *EnhanceTransactionResponse) SetMerchantCategoryCodeNil(b bool)`

 SetMerchantCategoryCodeNil sets the value for MerchantCategoryCode to be an explicit nil

### UnsetMerchantCategoryCode
`func (o *EnhanceTransactionResponse) UnsetMerchantCategoryCode()`

UnsetMerchantCategoryCode ensures that no value is present for MerchantCategoryCode, not even an explicit nil
### GetMerchantGuid

`func (o *EnhanceTransactionResponse) GetMerchantGuid() string`

GetMerchantGuid returns the MerchantGuid field if non-nil, zero value otherwise.

### GetMerchantGuidOk

`func (o *EnhanceTransactionResponse) GetMerchantGuidOk() (*string, bool)`

GetMerchantGuidOk returns a tuple with the MerchantGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGuid

`func (o *EnhanceTransactionResponse) SetMerchantGuid(v string)`

SetMerchantGuid sets MerchantGuid field to given value.

### HasMerchantGuid

`func (o *EnhanceTransactionResponse) HasMerchantGuid() bool`

HasMerchantGuid returns a boolean if a field has been set.

### SetMerchantGuidNil

`func (o *EnhanceTransactionResponse) SetMerchantGuidNil(b bool)`

 SetMerchantGuidNil sets the value for MerchantGuid to be an explicit nil

### UnsetMerchantGuid
`func (o *EnhanceTransactionResponse) UnsetMerchantGuid()`

UnsetMerchantGuid ensures that no value is present for MerchantGuid, not even an explicit nil
### GetMerchantLocationGuid

`func (o *EnhanceTransactionResponse) GetMerchantLocationGuid() string`

GetMerchantLocationGuid returns the MerchantLocationGuid field if non-nil, zero value otherwise.

### GetMerchantLocationGuidOk

`func (o *EnhanceTransactionResponse) GetMerchantLocationGuidOk() (*string, bool)`

GetMerchantLocationGuidOk returns a tuple with the MerchantLocationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantLocationGuid

`func (o *EnhanceTransactionResponse) SetMerchantLocationGuid(v string)`

SetMerchantLocationGuid sets MerchantLocationGuid field to given value.

### HasMerchantLocationGuid

`func (o *EnhanceTransactionResponse) HasMerchantLocationGuid() bool`

HasMerchantLocationGuid returns a boolean if a field has been set.

### SetMerchantLocationGuidNil

`func (o *EnhanceTransactionResponse) SetMerchantLocationGuidNil(b bool)`

 SetMerchantLocationGuidNil sets the value for MerchantLocationGuid to be an explicit nil

### UnsetMerchantLocationGuid
`func (o *EnhanceTransactionResponse) UnsetMerchantLocationGuid()`

UnsetMerchantLocationGuid ensures that no value is present for MerchantLocationGuid, not even an explicit nil
### GetOriginalDescription

`func (o *EnhanceTransactionResponse) GetOriginalDescription() string`

GetOriginalDescription returns the OriginalDescription field if non-nil, zero value otherwise.

### GetOriginalDescriptionOk

`func (o *EnhanceTransactionResponse) GetOriginalDescriptionOk() (*string, bool)`

GetOriginalDescriptionOk returns a tuple with the OriginalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDescription

`func (o *EnhanceTransactionResponse) SetOriginalDescription(v string)`

SetOriginalDescription sets OriginalDescription field to given value.

### HasOriginalDescription

`func (o *EnhanceTransactionResponse) HasOriginalDescription() bool`

HasOriginalDescription returns a boolean if a field has been set.

### SetOriginalDescriptionNil

`func (o *EnhanceTransactionResponse) SetOriginalDescriptionNil(b bool)`

 SetOriginalDescriptionNil sets the value for OriginalDescription to be an explicit nil

### UnsetOriginalDescription
`func (o *EnhanceTransactionResponse) UnsetOriginalDescription()`

UnsetOriginalDescription ensures that no value is present for OriginalDescription, not even an explicit nil
### GetType

`func (o *EnhanceTransactionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnhanceTransactionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnhanceTransactionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnhanceTransactionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *EnhanceTransactionResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *EnhanceTransactionResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


