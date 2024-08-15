# BudgetUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Amount of the budget. | [optional] 
**Metadata** | Pointer to **string** | Additional information a partner can store on the budget. | [optional] 
**SkipWebhook** | Pointer to **bool** | When set to true, this parameter will prevent a webhook from being triggered by the request. | [optional] 

## Methods

### NewBudgetUpdateRequest

`func NewBudgetUpdateRequest() *BudgetUpdateRequest`

NewBudgetUpdateRequest instantiates a new BudgetUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetUpdateRequestWithDefaults

`func NewBudgetUpdateRequestWithDefaults() *BudgetUpdateRequest`

NewBudgetUpdateRequestWithDefaults instantiates a new BudgetUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BudgetUpdateRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BudgetUpdateRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BudgetUpdateRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BudgetUpdateRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *BudgetUpdateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BudgetUpdateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BudgetUpdateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BudgetUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSkipWebhook

`func (o *BudgetUpdateRequest) GetSkipWebhook() bool`

GetSkipWebhook returns the SkipWebhook field if non-nil, zero value otherwise.

### GetSkipWebhookOk

`func (o *BudgetUpdateRequest) GetSkipWebhookOk() (*bool, bool)`

GetSkipWebhookOk returns a tuple with the SkipWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhook

`func (o *BudgetUpdateRequest) SetSkipWebhook(v bool)`

SetSkipWebhook sets SkipWebhook field to given value.

### HasSkipWebhook

`func (o *BudgetUpdateRequest) HasSkipWebhook() bool`

HasSkipWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


