# TransactionRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryGuid** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**MatchDescription** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**UserGuid** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionRuleResponse

`func NewTransactionRuleResponse() *TransactionRuleResponse`

NewTransactionRuleResponse instantiates a new TransactionRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRuleResponseWithDefaults

`func NewTransactionRuleResponseWithDefaults() *TransactionRuleResponse`

NewTransactionRuleResponseWithDefaults instantiates a new TransactionRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryGuid

`func (o *TransactionRuleResponse) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *TransactionRuleResponse) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *TransactionRuleResponse) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *TransactionRuleResponse) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TransactionRuleResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionRuleResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionRuleResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransactionRuleResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TransactionRuleResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TransactionRuleResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDescription

`func (o *TransactionRuleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionRuleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionRuleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionRuleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TransactionRuleResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TransactionRuleResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGuid

`func (o *TransactionRuleResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *TransactionRuleResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *TransactionRuleResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *TransactionRuleResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetMatchDescription

`func (o *TransactionRuleResponse) GetMatchDescription() string`

GetMatchDescription returns the MatchDescription field if non-nil, zero value otherwise.

### GetMatchDescriptionOk

`func (o *TransactionRuleResponse) GetMatchDescriptionOk() (*string, bool)`

GetMatchDescriptionOk returns a tuple with the MatchDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchDescription

`func (o *TransactionRuleResponse) SetMatchDescription(v string)`

SetMatchDescription sets MatchDescription field to given value.

### HasMatchDescription

`func (o *TransactionRuleResponse) HasMatchDescription() bool`

HasMatchDescription returns a boolean if a field has been set.

### SetMatchDescriptionNil

`func (o *TransactionRuleResponse) SetMatchDescriptionNil(b bool)`

 SetMatchDescriptionNil sets the value for MatchDescription to be an explicit nil

### UnsetMatchDescription
`func (o *TransactionRuleResponse) UnsetMatchDescription()`

UnsetMatchDescription ensures that no value is present for MatchDescription, not even an explicit nil
### GetUpdatedAt

`func (o *TransactionRuleResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransactionRuleResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransactionRuleResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TransactionRuleResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TransactionRuleResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TransactionRuleResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserGuid

`func (o *TransactionRuleResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *TransactionRuleResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *TransactionRuleResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *TransactionRuleResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


