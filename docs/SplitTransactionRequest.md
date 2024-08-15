# SplitTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of money you want to re-categorize. | 
**Description** | Pointer to **string** | Description for the split transaction. | [optional] 
**CategoryGuid** | Pointer to **string** | Unique identifier of the category. | [optional] 
**Memo** | Pointer to **string** | Memo for the split transaction | [optional] 

## Methods

### NewSplitTransactionRequest

`func NewSplitTransactionRequest(amount float32, ) *SplitTransactionRequest`

NewSplitTransactionRequest instantiates a new SplitTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitTransactionRequestWithDefaults

`func NewSplitTransactionRequestWithDefaults() *SplitTransactionRequest`

NewSplitTransactionRequestWithDefaults instantiates a new SplitTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SplitTransactionRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SplitTransactionRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SplitTransactionRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *SplitTransactionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SplitTransactionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SplitTransactionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SplitTransactionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategoryGuid

`func (o *SplitTransactionRequest) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *SplitTransactionRequest) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *SplitTransactionRequest) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *SplitTransactionRequest) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### GetMemo

`func (o *SplitTransactionRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *SplitTransactionRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *SplitTransactionRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *SplitTransactionRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


