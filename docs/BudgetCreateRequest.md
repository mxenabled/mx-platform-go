# BudgetCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryGuid** | **string** | Unique identifier of the category. | 
**ParentGuid** | **string** | Unique identifier of the parent budget. This is only required when creating a budget on a sub-category. | 
**Amount** | Pointer to **int32** | Amount of the budget. | [optional] 
**Metadata** | Pointer to **string** | Additional information a partner can store on the budget. | [optional] 
**SkipWebhook** | Pointer to **bool** | When set to true, this parameter will prevent a webhook from being triggered by the request. | [optional] 

## Methods

### NewBudgetCreateRequest

`func NewBudgetCreateRequest(categoryGuid string, parentGuid string, ) *BudgetCreateRequest`

NewBudgetCreateRequest instantiates a new BudgetCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetCreateRequestWithDefaults

`func NewBudgetCreateRequestWithDefaults() *BudgetCreateRequest`

NewBudgetCreateRequestWithDefaults instantiates a new BudgetCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryGuid

`func (o *BudgetCreateRequest) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *BudgetCreateRequest) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *BudgetCreateRequest) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.


### GetParentGuid

`func (o *BudgetCreateRequest) GetParentGuid() string`

GetParentGuid returns the ParentGuid field if non-nil, zero value otherwise.

### GetParentGuidOk

`func (o *BudgetCreateRequest) GetParentGuidOk() (*string, bool)`

GetParentGuidOk returns a tuple with the ParentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGuid

`func (o *BudgetCreateRequest) SetParentGuid(v string)`

SetParentGuid sets ParentGuid field to given value.


### GetAmount

`func (o *BudgetCreateRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BudgetCreateRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BudgetCreateRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BudgetCreateRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *BudgetCreateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BudgetCreateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BudgetCreateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BudgetCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSkipWebhook

`func (o *BudgetCreateRequest) GetSkipWebhook() bool`

GetSkipWebhook returns the SkipWebhook field if non-nil, zero value otherwise.

### GetSkipWebhookOk

`func (o *BudgetCreateRequest) GetSkipWebhookOk() (*bool, bool)`

GetSkipWebhookOk returns a tuple with the SkipWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhook

`func (o *BudgetCreateRequest) SetSkipWebhook(v bool)`

SetSkipWebhook sets SkipWebhook field to given value.

### HasSkipWebhook

`func (o *BudgetCreateRequest) HasSkipWebhook() bool`

HasSkipWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


