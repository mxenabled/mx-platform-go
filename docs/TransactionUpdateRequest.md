# TransactionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**CategoryGuid** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionUpdateRequest

`func NewTransactionUpdateRequest() *TransactionUpdateRequest`

NewTransactionUpdateRequest instantiates a new TransactionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionUpdateRequestWithDefaults

`func NewTransactionUpdateRequestWithDefaults() *TransactionUpdateRequest`

NewTransactionUpdateRequestWithDefaults instantiates a new TransactionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *TransactionUpdateRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionUpdateRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionUpdateRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *TransactionUpdateRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMemo

`func (o *TransactionUpdateRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionUpdateRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionUpdateRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionUpdateRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetCategoryGuid

`func (o *TransactionUpdateRequest) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *TransactionUpdateRequest) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *TransactionUpdateRequest) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *TransactionUpdateRequest) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


