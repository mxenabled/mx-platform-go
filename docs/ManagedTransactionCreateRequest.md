# ManagedTransactionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** |  | 
**Category** | Pointer to **string** |  | [optional] 
**CheckNumberString** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 
**IsInternational** | Pointer to **bool** |  | [optional] 
**Latitude** | Pointer to **float32** |  | [optional] 
**LocalizedDescription** | Pointer to **string** |  | [optional] 
**LocalizedMemo** | Pointer to **string** |  | [optional] 
**Longitude** | Pointer to **float32** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**MerchantCategoryCode** | Pointer to **int32** |  | [optional] 
**MerchantGuid** | Pointer to **string** |  | [optional] 
**MerchantLocationGuid** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**PostedAt** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**TransactedAt** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewManagedTransactionCreateRequest

`func NewManagedTransactionCreateRequest(amount string, description string, status string, transactedAt string, type_ string, ) *ManagedTransactionCreateRequest`

NewManagedTransactionCreateRequest instantiates a new ManagedTransactionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedTransactionCreateRequestWithDefaults

`func NewManagedTransactionCreateRequestWithDefaults() *ManagedTransactionCreateRequest`

NewManagedTransactionCreateRequestWithDefaults instantiates a new ManagedTransactionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ManagedTransactionCreateRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ManagedTransactionCreateRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ManagedTransactionCreateRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCategory

`func (o *ManagedTransactionCreateRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ManagedTransactionCreateRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ManagedTransactionCreateRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ManagedTransactionCreateRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCheckNumberString

`func (o *ManagedTransactionCreateRequest) GetCheckNumberString() string`

GetCheckNumberString returns the CheckNumberString field if non-nil, zero value otherwise.

### GetCheckNumberStringOk

`func (o *ManagedTransactionCreateRequest) GetCheckNumberStringOk() (*string, bool)`

GetCheckNumberStringOk returns a tuple with the CheckNumberString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumberString

`func (o *ManagedTransactionCreateRequest) SetCheckNumberString(v string)`

SetCheckNumberString sets CheckNumberString field to given value.

### HasCheckNumberString

`func (o *ManagedTransactionCreateRequest) HasCheckNumberString() bool`

HasCheckNumberString returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ManagedTransactionCreateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ManagedTransactionCreateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ManagedTransactionCreateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ManagedTransactionCreateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDescription

`func (o *ManagedTransactionCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagedTransactionCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagedTransactionCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *ManagedTransactionCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedTransactionCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedTransactionCreateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedTransactionCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsInternational

`func (o *ManagedTransactionCreateRequest) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *ManagedTransactionCreateRequest) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *ManagedTransactionCreateRequest) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *ManagedTransactionCreateRequest) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### GetLatitude

`func (o *ManagedTransactionCreateRequest) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ManagedTransactionCreateRequest) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ManagedTransactionCreateRequest) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *ManagedTransactionCreateRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLocalizedDescription

`func (o *ManagedTransactionCreateRequest) GetLocalizedDescription() string`

GetLocalizedDescription returns the LocalizedDescription field if non-nil, zero value otherwise.

### GetLocalizedDescriptionOk

`func (o *ManagedTransactionCreateRequest) GetLocalizedDescriptionOk() (*string, bool)`

GetLocalizedDescriptionOk returns a tuple with the LocalizedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedDescription

`func (o *ManagedTransactionCreateRequest) SetLocalizedDescription(v string)`

SetLocalizedDescription sets LocalizedDescription field to given value.

### HasLocalizedDescription

`func (o *ManagedTransactionCreateRequest) HasLocalizedDescription() bool`

HasLocalizedDescription returns a boolean if a field has been set.

### GetLocalizedMemo

`func (o *ManagedTransactionCreateRequest) GetLocalizedMemo() string`

GetLocalizedMemo returns the LocalizedMemo field if non-nil, zero value otherwise.

### GetLocalizedMemoOk

`func (o *ManagedTransactionCreateRequest) GetLocalizedMemoOk() (*string, bool)`

GetLocalizedMemoOk returns a tuple with the LocalizedMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedMemo

`func (o *ManagedTransactionCreateRequest) SetLocalizedMemo(v string)`

SetLocalizedMemo sets LocalizedMemo field to given value.

### HasLocalizedMemo

`func (o *ManagedTransactionCreateRequest) HasLocalizedMemo() bool`

HasLocalizedMemo returns a boolean if a field has been set.

### GetLongitude

`func (o *ManagedTransactionCreateRequest) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ManagedTransactionCreateRequest) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ManagedTransactionCreateRequest) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *ManagedTransactionCreateRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetMemo

`func (o *ManagedTransactionCreateRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ManagedTransactionCreateRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ManagedTransactionCreateRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ManagedTransactionCreateRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMerchantCategoryCode

`func (o *ManagedTransactionCreateRequest) GetMerchantCategoryCode() int32`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *ManagedTransactionCreateRequest) GetMerchantCategoryCodeOk() (*int32, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *ManagedTransactionCreateRequest) SetMerchantCategoryCode(v int32)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *ManagedTransactionCreateRequest) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### GetMerchantGuid

`func (o *ManagedTransactionCreateRequest) GetMerchantGuid() string`

GetMerchantGuid returns the MerchantGuid field if non-nil, zero value otherwise.

### GetMerchantGuidOk

`func (o *ManagedTransactionCreateRequest) GetMerchantGuidOk() (*string, bool)`

GetMerchantGuidOk returns a tuple with the MerchantGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGuid

`func (o *ManagedTransactionCreateRequest) SetMerchantGuid(v string)`

SetMerchantGuid sets MerchantGuid field to given value.

### HasMerchantGuid

`func (o *ManagedTransactionCreateRequest) HasMerchantGuid() bool`

HasMerchantGuid returns a boolean if a field has been set.

### GetMerchantLocationGuid

`func (o *ManagedTransactionCreateRequest) GetMerchantLocationGuid() string`

GetMerchantLocationGuid returns the MerchantLocationGuid field if non-nil, zero value otherwise.

### GetMerchantLocationGuidOk

`func (o *ManagedTransactionCreateRequest) GetMerchantLocationGuidOk() (*string, bool)`

GetMerchantLocationGuidOk returns a tuple with the MerchantLocationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantLocationGuid

`func (o *ManagedTransactionCreateRequest) SetMerchantLocationGuid(v string)`

SetMerchantLocationGuid sets MerchantLocationGuid field to given value.

### HasMerchantLocationGuid

`func (o *ManagedTransactionCreateRequest) HasMerchantLocationGuid() bool`

HasMerchantLocationGuid returns a boolean if a field has been set.

### GetMetadata

`func (o *ManagedTransactionCreateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ManagedTransactionCreateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ManagedTransactionCreateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ManagedTransactionCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPostedAt

`func (o *ManagedTransactionCreateRequest) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *ManagedTransactionCreateRequest) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *ManagedTransactionCreateRequest) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *ManagedTransactionCreateRequest) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ManagedTransactionCreateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedTransactionCreateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedTransactionCreateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransactedAt

`func (o *ManagedTransactionCreateRequest) GetTransactedAt() string`

GetTransactedAt returns the TransactedAt field if non-nil, zero value otherwise.

### GetTransactedAtOk

`func (o *ManagedTransactionCreateRequest) GetTransactedAtOk() (*string, bool)`

GetTransactedAtOk returns a tuple with the TransactedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactedAt

`func (o *ManagedTransactionCreateRequest) SetTransactedAt(v string)`

SetTransactedAt sets TransactedAt field to given value.


### GetType

`func (o *ManagedTransactionCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagedTransactionCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagedTransactionCreateRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


