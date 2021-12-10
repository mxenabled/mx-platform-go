# EnhanceTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Description** | **string** |  | 
**ExtendedTransactionType** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Memo** | Pointer to **string** |  | [optional] 
**MerchantCategoryCode** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewEnhanceTransactionsRequest

`func NewEnhanceTransactionsRequest(description string, id string, ) *EnhanceTransactionsRequest`

NewEnhanceTransactionsRequest instantiates a new EnhanceTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhanceTransactionsRequestWithDefaults

`func NewEnhanceTransactionsRequestWithDefaults() *EnhanceTransactionsRequest`

NewEnhanceTransactionsRequestWithDefaults instantiates a new EnhanceTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EnhanceTransactionsRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnhanceTransactionsRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnhanceTransactionsRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EnhanceTransactionsRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *EnhanceTransactionsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnhanceTransactionsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnhanceTransactionsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExtendedTransactionType

`func (o *EnhanceTransactionsRequest) GetExtendedTransactionType() string`

GetExtendedTransactionType returns the ExtendedTransactionType field if non-nil, zero value otherwise.

### GetExtendedTransactionTypeOk

`func (o *EnhanceTransactionsRequest) GetExtendedTransactionTypeOk() (*string, bool)`

GetExtendedTransactionTypeOk returns a tuple with the ExtendedTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedTransactionType

`func (o *EnhanceTransactionsRequest) SetExtendedTransactionType(v string)`

SetExtendedTransactionType sets ExtendedTransactionType field to given value.

### HasExtendedTransactionType

`func (o *EnhanceTransactionsRequest) HasExtendedTransactionType() bool`

HasExtendedTransactionType returns a boolean if a field has been set.

### GetId

`func (o *EnhanceTransactionsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnhanceTransactionsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnhanceTransactionsRequest) SetId(v string)`

SetId sets Id field to given value.


### GetMemo

`func (o *EnhanceTransactionsRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *EnhanceTransactionsRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *EnhanceTransactionsRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *EnhanceTransactionsRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMerchantCategoryCode

`func (o *EnhanceTransactionsRequest) GetMerchantCategoryCode() int32`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *EnhanceTransactionsRequest) GetMerchantCategoryCodeOk() (*int32, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *EnhanceTransactionsRequest) SetMerchantCategoryCode(v int32)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *EnhanceTransactionsRequest) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### GetType

`func (o *EnhanceTransactionsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnhanceTransactionsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnhanceTransactionsRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnhanceTransactionsRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


