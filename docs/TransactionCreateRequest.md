# TransactionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** |  | 
**Date** | **string** |  | 
**Description** | **string** |  | 
**Type** | **string** | The type of transaction, which must be CREDIT or DEBIT. See Transaction Fields for more information. | 
**CategoryGuid** | Pointer to **string** | Unique identifier of the category. | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**HasBeenViewed** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**IsInternational** | Pointer to **bool** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**SkipWebhook** | Pointer to **bool** | When set to true, this parameter will prevent a webhook from being triggered by the request. | [optional] 

## Methods

### NewTransactionCreateRequest

`func NewTransactionCreateRequest(amount float32, date string, description string, type_ string, ) *TransactionCreateRequest`

NewTransactionCreateRequest instantiates a new TransactionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCreateRequestWithDefaults

`func NewTransactionCreateRequestWithDefaults() *TransactionCreateRequest`

NewTransactionCreateRequestWithDefaults instantiates a new TransactionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionCreateRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionCreateRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionCreateRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDate

`func (o *TransactionCreateRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionCreateRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionCreateRequest) SetDate(v string)`

SetDate sets Date field to given value.


### GetDescription

`func (o *TransactionCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *TransactionCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCategoryGuid

`func (o *TransactionCreateRequest) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *TransactionCreateRequest) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *TransactionCreateRequest) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *TransactionCreateRequest) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *TransactionCreateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *TransactionCreateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *TransactionCreateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *TransactionCreateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetHasBeenViewed

`func (o *TransactionCreateRequest) GetHasBeenViewed() bool`

GetHasBeenViewed returns the HasBeenViewed field if non-nil, zero value otherwise.

### GetHasBeenViewedOk

`func (o *TransactionCreateRequest) GetHasBeenViewedOk() (*bool, bool)`

GetHasBeenViewedOk returns a tuple with the HasBeenViewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenViewed

`func (o *TransactionCreateRequest) SetHasBeenViewed(v bool)`

SetHasBeenViewed sets HasBeenViewed field to given value.

### HasHasBeenViewed

`func (o *TransactionCreateRequest) HasHasBeenViewed() bool`

HasHasBeenViewed returns a boolean if a field has been set.

### GetIsHidden

`func (o *TransactionCreateRequest) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *TransactionCreateRequest) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *TransactionCreateRequest) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *TransactionCreateRequest) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsInternational

`func (o *TransactionCreateRequest) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *TransactionCreateRequest) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *TransactionCreateRequest) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *TransactionCreateRequest) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### GetMemo

`func (o *TransactionCreateRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionCreateRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionCreateRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionCreateRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMetadata

`func (o *TransactionCreateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransactionCreateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransactionCreateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransactionCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSkipWebhook

`func (o *TransactionCreateRequest) GetSkipWebhook() bool`

GetSkipWebhook returns the SkipWebhook field if non-nil, zero value otherwise.

### GetSkipWebhookOk

`func (o *TransactionCreateRequest) GetSkipWebhookOk() (*bool, bool)`

GetSkipWebhookOk returns a tuple with the SkipWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhook

`func (o *TransactionCreateRequest) SetSkipWebhook(v bool)`

SetSkipWebhook sets SkipWebhook field to given value.

### HasSkipWebhook

`func (o *TransactionCreateRequest) HasSkipWebhook() bool`

HasSkipWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


