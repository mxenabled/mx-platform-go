# ManagedTransactionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CheckNumberString** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
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
**Status** | Pointer to **string** |  | [optional] 
**TransactedAt** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedTransactionUpdateRequest

`func NewManagedTransactionUpdateRequest() *ManagedTransactionUpdateRequest`

NewManagedTransactionUpdateRequest instantiates a new ManagedTransactionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedTransactionUpdateRequestWithDefaults

`func NewManagedTransactionUpdateRequestWithDefaults() *ManagedTransactionUpdateRequest`

NewManagedTransactionUpdateRequestWithDefaults instantiates a new ManagedTransactionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ManagedTransactionUpdateRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ManagedTransactionUpdateRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ManagedTransactionUpdateRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ManagedTransactionUpdateRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCategory

`func (o *ManagedTransactionUpdateRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ManagedTransactionUpdateRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ManagedTransactionUpdateRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ManagedTransactionUpdateRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCheckNumberString

`func (o *ManagedTransactionUpdateRequest) GetCheckNumberString() string`

GetCheckNumberString returns the CheckNumberString field if non-nil, zero value otherwise.

### GetCheckNumberStringOk

`func (o *ManagedTransactionUpdateRequest) GetCheckNumberStringOk() (*string, bool)`

GetCheckNumberStringOk returns a tuple with the CheckNumberString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumberString

`func (o *ManagedTransactionUpdateRequest) SetCheckNumberString(v string)`

SetCheckNumberString sets CheckNumberString field to given value.

### HasCheckNumberString

`func (o *ManagedTransactionUpdateRequest) HasCheckNumberString() bool`

HasCheckNumberString returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ManagedTransactionUpdateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ManagedTransactionUpdateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ManagedTransactionUpdateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ManagedTransactionUpdateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDescription

`func (o *ManagedTransactionUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagedTransactionUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagedTransactionUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManagedTransactionUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ManagedTransactionUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedTransactionUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedTransactionUpdateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedTransactionUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsInternational

`func (o *ManagedTransactionUpdateRequest) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *ManagedTransactionUpdateRequest) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *ManagedTransactionUpdateRequest) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *ManagedTransactionUpdateRequest) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### GetLatitude

`func (o *ManagedTransactionUpdateRequest) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ManagedTransactionUpdateRequest) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ManagedTransactionUpdateRequest) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *ManagedTransactionUpdateRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLocalizedDescription

`func (o *ManagedTransactionUpdateRequest) GetLocalizedDescription() string`

GetLocalizedDescription returns the LocalizedDescription field if non-nil, zero value otherwise.

### GetLocalizedDescriptionOk

`func (o *ManagedTransactionUpdateRequest) GetLocalizedDescriptionOk() (*string, bool)`

GetLocalizedDescriptionOk returns a tuple with the LocalizedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedDescription

`func (o *ManagedTransactionUpdateRequest) SetLocalizedDescription(v string)`

SetLocalizedDescription sets LocalizedDescription field to given value.

### HasLocalizedDescription

`func (o *ManagedTransactionUpdateRequest) HasLocalizedDescription() bool`

HasLocalizedDescription returns a boolean if a field has been set.

### GetLocalizedMemo

`func (o *ManagedTransactionUpdateRequest) GetLocalizedMemo() string`

GetLocalizedMemo returns the LocalizedMemo field if non-nil, zero value otherwise.

### GetLocalizedMemoOk

`func (o *ManagedTransactionUpdateRequest) GetLocalizedMemoOk() (*string, bool)`

GetLocalizedMemoOk returns a tuple with the LocalizedMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedMemo

`func (o *ManagedTransactionUpdateRequest) SetLocalizedMemo(v string)`

SetLocalizedMemo sets LocalizedMemo field to given value.

### HasLocalizedMemo

`func (o *ManagedTransactionUpdateRequest) HasLocalizedMemo() bool`

HasLocalizedMemo returns a boolean if a field has been set.

### GetLongitude

`func (o *ManagedTransactionUpdateRequest) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ManagedTransactionUpdateRequest) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ManagedTransactionUpdateRequest) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *ManagedTransactionUpdateRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetMemo

`func (o *ManagedTransactionUpdateRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ManagedTransactionUpdateRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ManagedTransactionUpdateRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ManagedTransactionUpdateRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMerchantCategoryCode

`func (o *ManagedTransactionUpdateRequest) GetMerchantCategoryCode() int32`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *ManagedTransactionUpdateRequest) GetMerchantCategoryCodeOk() (*int32, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *ManagedTransactionUpdateRequest) SetMerchantCategoryCode(v int32)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *ManagedTransactionUpdateRequest) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### GetMerchantGuid

`func (o *ManagedTransactionUpdateRequest) GetMerchantGuid() string`

GetMerchantGuid returns the MerchantGuid field if non-nil, zero value otherwise.

### GetMerchantGuidOk

`func (o *ManagedTransactionUpdateRequest) GetMerchantGuidOk() (*string, bool)`

GetMerchantGuidOk returns a tuple with the MerchantGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGuid

`func (o *ManagedTransactionUpdateRequest) SetMerchantGuid(v string)`

SetMerchantGuid sets MerchantGuid field to given value.

### HasMerchantGuid

`func (o *ManagedTransactionUpdateRequest) HasMerchantGuid() bool`

HasMerchantGuid returns a boolean if a field has been set.

### GetMerchantLocationGuid

`func (o *ManagedTransactionUpdateRequest) GetMerchantLocationGuid() string`

GetMerchantLocationGuid returns the MerchantLocationGuid field if non-nil, zero value otherwise.

### GetMerchantLocationGuidOk

`func (o *ManagedTransactionUpdateRequest) GetMerchantLocationGuidOk() (*string, bool)`

GetMerchantLocationGuidOk returns a tuple with the MerchantLocationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantLocationGuid

`func (o *ManagedTransactionUpdateRequest) SetMerchantLocationGuid(v string)`

SetMerchantLocationGuid sets MerchantLocationGuid field to given value.

### HasMerchantLocationGuid

`func (o *ManagedTransactionUpdateRequest) HasMerchantLocationGuid() bool`

HasMerchantLocationGuid returns a boolean if a field has been set.

### GetMetadata

`func (o *ManagedTransactionUpdateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ManagedTransactionUpdateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ManagedTransactionUpdateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ManagedTransactionUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPostedAt

`func (o *ManagedTransactionUpdateRequest) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *ManagedTransactionUpdateRequest) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *ManagedTransactionUpdateRequest) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *ManagedTransactionUpdateRequest) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ManagedTransactionUpdateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedTransactionUpdateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedTransactionUpdateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManagedTransactionUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactedAt

`func (o *ManagedTransactionUpdateRequest) GetTransactedAt() string`

GetTransactedAt returns the TransactedAt field if non-nil, zero value otherwise.

### GetTransactedAtOk

`func (o *ManagedTransactionUpdateRequest) GetTransactedAtOk() (*string, bool)`

GetTransactedAtOk returns a tuple with the TransactedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactedAt

`func (o *ManagedTransactionUpdateRequest) SetTransactedAt(v string)`

SetTransactedAt sets TransactedAt field to given value.

### HasTransactedAt

`func (o *ManagedTransactionUpdateRequest) HasTransactedAt() bool`

HasTransactedAt returns a boolean if a field has been set.

### GetType

`func (o *ManagedTransactionUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagedTransactionUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagedTransactionUpdateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManagedTransactionUpdateRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


