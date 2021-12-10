# TransactionRuleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryGuid** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**MatchDescription** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTransactionRuleUpdateRequest

`func NewTransactionRuleUpdateRequest() *TransactionRuleUpdateRequest`

NewTransactionRuleUpdateRequest instantiates a new TransactionRuleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRuleUpdateRequestWithDefaults

`func NewTransactionRuleUpdateRequestWithDefaults() *TransactionRuleUpdateRequest`

NewTransactionRuleUpdateRequestWithDefaults instantiates a new TransactionRuleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryGuid

`func (o *TransactionRuleUpdateRequest) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *TransactionRuleUpdateRequest) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *TransactionRuleUpdateRequest) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *TransactionRuleUpdateRequest) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionRuleUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionRuleUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionRuleUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionRuleUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TransactionRuleUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TransactionRuleUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMatchDescription

`func (o *TransactionRuleUpdateRequest) GetMatchDescription() string`

GetMatchDescription returns the MatchDescription field if non-nil, zero value otherwise.

### GetMatchDescriptionOk

`func (o *TransactionRuleUpdateRequest) GetMatchDescriptionOk() (*string, bool)`

GetMatchDescriptionOk returns a tuple with the MatchDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchDescription

`func (o *TransactionRuleUpdateRequest) SetMatchDescription(v string)`

SetMatchDescription sets MatchDescription field to given value.

### HasMatchDescription

`func (o *TransactionRuleUpdateRequest) HasMatchDescription() bool`

HasMatchDescription returns a boolean if a field has been set.

### SetMatchDescriptionNil

`func (o *TransactionRuleUpdateRequest) SetMatchDescriptionNil(b bool)`

 SetMatchDescriptionNil sets the value for MatchDescription to be an explicit nil

### UnsetMatchDescription
`func (o *TransactionRuleUpdateRequest) UnsetMatchDescription()`

UnsetMatchDescription ensures that no value is present for MatchDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


