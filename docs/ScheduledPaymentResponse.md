# ScheduledPaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**IsCompleted** | Pointer to **bool** |  | [optional] 
**IsRecurring** | Pointer to **bool** |  | [optional] 
**MerchantGuid** | Pointer to **string** |  | [optional] 
**OccursOn** | Pointer to **string** |  | [optional] 
**RecurrenceDay** | Pointer to **int32** |  | [optional] 
**RecurrenceType** | Pointer to **string** |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserGuid** | Pointer to **string** |  | [optional] 

## Methods

### NewScheduledPaymentResponse

`func NewScheduledPaymentResponse() *ScheduledPaymentResponse`

NewScheduledPaymentResponse instantiates a new ScheduledPaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledPaymentResponseWithDefaults

`func NewScheduledPaymentResponseWithDefaults() *ScheduledPaymentResponse`

NewScheduledPaymentResponseWithDefaults instantiates a new ScheduledPaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ScheduledPaymentResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ScheduledPaymentResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ScheduledPaymentResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ScheduledPaymentResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ScheduledPaymentResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScheduledPaymentResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScheduledPaymentResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ScheduledPaymentResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ScheduledPaymentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduledPaymentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduledPaymentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduledPaymentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGuid

`func (o *ScheduledPaymentResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ScheduledPaymentResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ScheduledPaymentResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ScheduledPaymentResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetIsCompleted

`func (o *ScheduledPaymentResponse) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *ScheduledPaymentResponse) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *ScheduledPaymentResponse) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.

### HasIsCompleted

`func (o *ScheduledPaymentResponse) HasIsCompleted() bool`

HasIsCompleted returns a boolean if a field has been set.

### GetIsRecurring

`func (o *ScheduledPaymentResponse) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *ScheduledPaymentResponse) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *ScheduledPaymentResponse) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.

### HasIsRecurring

`func (o *ScheduledPaymentResponse) HasIsRecurring() bool`

HasIsRecurring returns a boolean if a field has been set.

### GetMerchantGuid

`func (o *ScheduledPaymentResponse) GetMerchantGuid() string`

GetMerchantGuid returns the MerchantGuid field if non-nil, zero value otherwise.

### GetMerchantGuidOk

`func (o *ScheduledPaymentResponse) GetMerchantGuidOk() (*string, bool)`

GetMerchantGuidOk returns a tuple with the MerchantGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGuid

`func (o *ScheduledPaymentResponse) SetMerchantGuid(v string)`

SetMerchantGuid sets MerchantGuid field to given value.

### HasMerchantGuid

`func (o *ScheduledPaymentResponse) HasMerchantGuid() bool`

HasMerchantGuid returns a boolean if a field has been set.

### GetOccursOn

`func (o *ScheduledPaymentResponse) GetOccursOn() string`

GetOccursOn returns the OccursOn field if non-nil, zero value otherwise.

### GetOccursOnOk

`func (o *ScheduledPaymentResponse) GetOccursOnOk() (*string, bool)`

GetOccursOnOk returns a tuple with the OccursOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccursOn

`func (o *ScheduledPaymentResponse) SetOccursOn(v string)`

SetOccursOn sets OccursOn field to given value.

### HasOccursOn

`func (o *ScheduledPaymentResponse) HasOccursOn() bool`

HasOccursOn returns a boolean if a field has been set.

### GetRecurrenceDay

`func (o *ScheduledPaymentResponse) GetRecurrenceDay() int32`

GetRecurrenceDay returns the RecurrenceDay field if non-nil, zero value otherwise.

### GetRecurrenceDayOk

`func (o *ScheduledPaymentResponse) GetRecurrenceDayOk() (*int32, bool)`

GetRecurrenceDayOk returns a tuple with the RecurrenceDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceDay

`func (o *ScheduledPaymentResponse) SetRecurrenceDay(v int32)`

SetRecurrenceDay sets RecurrenceDay field to given value.

### HasRecurrenceDay

`func (o *ScheduledPaymentResponse) HasRecurrenceDay() bool`

HasRecurrenceDay returns a boolean if a field has been set.

### GetRecurrenceType

`func (o *ScheduledPaymentResponse) GetRecurrenceType() string`

GetRecurrenceType returns the RecurrenceType field if non-nil, zero value otherwise.

### GetRecurrenceTypeOk

`func (o *ScheduledPaymentResponse) GetRecurrenceTypeOk() (*string, bool)`

GetRecurrenceTypeOk returns a tuple with the RecurrenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceType

`func (o *ScheduledPaymentResponse) SetRecurrenceType(v string)`

SetRecurrenceType sets RecurrenceType field to given value.

### HasRecurrenceType

`func (o *ScheduledPaymentResponse) HasRecurrenceType() bool`

HasRecurrenceType returns a boolean if a field has been set.

### GetTransactionType

`func (o *ScheduledPaymentResponse) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ScheduledPaymentResponse) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ScheduledPaymentResponse) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ScheduledPaymentResponse) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ScheduledPaymentResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScheduledPaymentResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScheduledPaymentResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ScheduledPaymentResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserGuid

`func (o *ScheduledPaymentResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *ScheduledPaymentResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *ScheduledPaymentResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *ScheduledPaymentResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


