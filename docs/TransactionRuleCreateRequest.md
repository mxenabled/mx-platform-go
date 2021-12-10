# TransactionRuleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryGuid** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**MatchDescription** | **string** |  | 

## Methods

### NewTransactionRuleCreateRequest

`func NewTransactionRuleCreateRequest(categoryGuid string, matchDescription string, ) *TransactionRuleCreateRequest`

NewTransactionRuleCreateRequest instantiates a new TransactionRuleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRuleCreateRequestWithDefaults

`func NewTransactionRuleCreateRequestWithDefaults() *TransactionRuleCreateRequest`

NewTransactionRuleCreateRequestWithDefaults instantiates a new TransactionRuleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryGuid

`func (o *TransactionRuleCreateRequest) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *TransactionRuleCreateRequest) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *TransactionRuleCreateRequest) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.


### GetDescription

`func (o *TransactionRuleCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionRuleCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionRuleCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionRuleCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMatchDescription

`func (o *TransactionRuleCreateRequest) GetMatchDescription() string`

GetMatchDescription returns the MatchDescription field if non-nil, zero value otherwise.

### GetMatchDescriptionOk

`func (o *TransactionRuleCreateRequest) GetMatchDescriptionOk() (*string, bool)`

GetMatchDescriptionOk returns a tuple with the MatchDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchDescription

`func (o *TransactionRuleCreateRequest) SetMatchDescription(v string)`

SetMatchDescription sets MatchDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


