# RepeatingTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**MemberGuid** | Pointer to **string** |  | [optional] 
**MerchantGuid** | Pointer to **string** |  | [optional] 
**LastPostedDate** | Pointer to **string** |  | [optional] 
**PredictedOccursOn** | Pointer to **string** |  | [optional] 
**RecurrenceType** | Pointer to **string** |  | [optional] 
**UserGuid** | Pointer to **string** |  | [optional] 
**RepeatingTransactionType** | Pointer to **string** |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 

## Methods

### NewRepeatingTransactionResponse

`func NewRepeatingTransactionResponse() *RepeatingTransactionResponse`

NewRepeatingTransactionResponse instantiates a new RepeatingTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepeatingTransactionResponseWithDefaults

`func NewRepeatingTransactionResponseWithDefaults() *RepeatingTransactionResponse`

NewRepeatingTransactionResponseWithDefaults instantiates a new RepeatingTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *RepeatingTransactionResponse) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *RepeatingTransactionResponse) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *RepeatingTransactionResponse) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.

### HasAccountGuid

`func (o *RepeatingTransactionResponse) HasAccountGuid() bool`

HasAccountGuid returns a boolean if a field has been set.

### SetAccountGuidNil

`func (o *RepeatingTransactionResponse) SetAccountGuidNil(b bool)`

 SetAccountGuidNil sets the value for AccountGuid to be an explicit nil

### UnsetAccountGuid
`func (o *RepeatingTransactionResponse) UnsetAccountGuid()`

UnsetAccountGuid ensures that no value is present for AccountGuid, not even an explicit nil
### GetAmount

`func (o *RepeatingTransactionResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RepeatingTransactionResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RepeatingTransactionResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RepeatingTransactionResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *RepeatingTransactionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepeatingTransactionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepeatingTransactionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RepeatingTransactionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGuid

`func (o *RepeatingTransactionResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *RepeatingTransactionResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *RepeatingTransactionResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *RepeatingTransactionResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetMemberGuid

`func (o *RepeatingTransactionResponse) GetMemberGuid() string`

GetMemberGuid returns the MemberGuid field if non-nil, zero value otherwise.

### GetMemberGuidOk

`func (o *RepeatingTransactionResponse) GetMemberGuidOk() (*string, bool)`

GetMemberGuidOk returns a tuple with the MemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGuid

`func (o *RepeatingTransactionResponse) SetMemberGuid(v string)`

SetMemberGuid sets MemberGuid field to given value.

### HasMemberGuid

`func (o *RepeatingTransactionResponse) HasMemberGuid() bool`

HasMemberGuid returns a boolean if a field has been set.

### GetMerchantGuid

`func (o *RepeatingTransactionResponse) GetMerchantGuid() string`

GetMerchantGuid returns the MerchantGuid field if non-nil, zero value otherwise.

### GetMerchantGuidOk

`func (o *RepeatingTransactionResponse) GetMerchantGuidOk() (*string, bool)`

GetMerchantGuidOk returns a tuple with the MerchantGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGuid

`func (o *RepeatingTransactionResponse) SetMerchantGuid(v string)`

SetMerchantGuid sets MerchantGuid field to given value.

### HasMerchantGuid

`func (o *RepeatingTransactionResponse) HasMerchantGuid() bool`

HasMerchantGuid returns a boolean if a field has been set.

### GetLastPostedDate

`func (o *RepeatingTransactionResponse) GetLastPostedDate() string`

GetLastPostedDate returns the LastPostedDate field if non-nil, zero value otherwise.

### GetLastPostedDateOk

`func (o *RepeatingTransactionResponse) GetLastPostedDateOk() (*string, bool)`

GetLastPostedDateOk returns a tuple with the LastPostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPostedDate

`func (o *RepeatingTransactionResponse) SetLastPostedDate(v string)`

SetLastPostedDate sets LastPostedDate field to given value.

### HasLastPostedDate

`func (o *RepeatingTransactionResponse) HasLastPostedDate() bool`

HasLastPostedDate returns a boolean if a field has been set.

### GetPredictedOccursOn

`func (o *RepeatingTransactionResponse) GetPredictedOccursOn() string`

GetPredictedOccursOn returns the PredictedOccursOn field if non-nil, zero value otherwise.

### GetPredictedOccursOnOk

`func (o *RepeatingTransactionResponse) GetPredictedOccursOnOk() (*string, bool)`

GetPredictedOccursOnOk returns a tuple with the PredictedOccursOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedOccursOn

`func (o *RepeatingTransactionResponse) SetPredictedOccursOn(v string)`

SetPredictedOccursOn sets PredictedOccursOn field to given value.

### HasPredictedOccursOn

`func (o *RepeatingTransactionResponse) HasPredictedOccursOn() bool`

HasPredictedOccursOn returns a boolean if a field has been set.

### GetRecurrenceType

`func (o *RepeatingTransactionResponse) GetRecurrenceType() string`

GetRecurrenceType returns the RecurrenceType field if non-nil, zero value otherwise.

### GetRecurrenceTypeOk

`func (o *RepeatingTransactionResponse) GetRecurrenceTypeOk() (*string, bool)`

GetRecurrenceTypeOk returns a tuple with the RecurrenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceType

`func (o *RepeatingTransactionResponse) SetRecurrenceType(v string)`

SetRecurrenceType sets RecurrenceType field to given value.

### HasRecurrenceType

`func (o *RepeatingTransactionResponse) HasRecurrenceType() bool`

HasRecurrenceType returns a boolean if a field has been set.

### GetUserGuid

`func (o *RepeatingTransactionResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *RepeatingTransactionResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *RepeatingTransactionResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *RepeatingTransactionResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### GetRepeatingTransactionType

`func (o *RepeatingTransactionResponse) GetRepeatingTransactionType() string`

GetRepeatingTransactionType returns the RepeatingTransactionType field if non-nil, zero value otherwise.

### GetRepeatingTransactionTypeOk

`func (o *RepeatingTransactionResponse) GetRepeatingTransactionTypeOk() (*string, bool)`

GetRepeatingTransactionTypeOk returns a tuple with the RepeatingTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatingTransactionType

`func (o *RepeatingTransactionResponse) SetRepeatingTransactionType(v string)`

SetRepeatingTransactionType sets RepeatingTransactionType field to given value.

### HasRepeatingTransactionType

`func (o *RepeatingTransactionResponse) HasRepeatingTransactionType() bool`

HasRepeatingTransactionType returns a boolean if a field has been set.

### GetTransactionType

`func (o *RepeatingTransactionResponse) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *RepeatingTransactionResponse) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *RepeatingTransactionResponse) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *RepeatingTransactionResponse) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


