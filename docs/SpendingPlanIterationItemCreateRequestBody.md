# SpendingPlanIterationItemCreateRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryGuid** | Pointer to **string** |  | [optional] 
**ItemType** | Pointer to **float32** |  | [optional] 
**PlannedAmount** | **float32** |  | 
**ScheduledPaymentGuid** | Pointer to **string** |  | [optional] 
**TopLevelCategoryGuid** | Pointer to **string** |  | [optional] 

## Methods

### NewSpendingPlanIterationItemCreateRequestBody

`func NewSpendingPlanIterationItemCreateRequestBody(plannedAmount float32, ) *SpendingPlanIterationItemCreateRequestBody`

NewSpendingPlanIterationItemCreateRequestBody instantiates a new SpendingPlanIterationItemCreateRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendingPlanIterationItemCreateRequestBodyWithDefaults

`func NewSpendingPlanIterationItemCreateRequestBodyWithDefaults() *SpendingPlanIterationItemCreateRequestBody`

NewSpendingPlanIterationItemCreateRequestBodyWithDefaults instantiates a new SpendingPlanIterationItemCreateRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryGuid

`func (o *SpendingPlanIterationItemCreateRequestBody) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *SpendingPlanIterationItemCreateRequestBody) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *SpendingPlanIterationItemCreateRequestBody) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *SpendingPlanIterationItemCreateRequestBody) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### GetItemType

`func (o *SpendingPlanIterationItemCreateRequestBody) GetItemType() float32`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *SpendingPlanIterationItemCreateRequestBody) GetItemTypeOk() (*float32, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *SpendingPlanIterationItemCreateRequestBody) SetItemType(v float32)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *SpendingPlanIterationItemCreateRequestBody) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetPlannedAmount

`func (o *SpendingPlanIterationItemCreateRequestBody) GetPlannedAmount() float32`

GetPlannedAmount returns the PlannedAmount field if non-nil, zero value otherwise.

### GetPlannedAmountOk

`func (o *SpendingPlanIterationItemCreateRequestBody) GetPlannedAmountOk() (*float32, bool)`

GetPlannedAmountOk returns a tuple with the PlannedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedAmount

`func (o *SpendingPlanIterationItemCreateRequestBody) SetPlannedAmount(v float32)`

SetPlannedAmount sets PlannedAmount field to given value.


### GetScheduledPaymentGuid

`func (o *SpendingPlanIterationItemCreateRequestBody) GetScheduledPaymentGuid() string`

GetScheduledPaymentGuid returns the ScheduledPaymentGuid field if non-nil, zero value otherwise.

### GetScheduledPaymentGuidOk

`func (o *SpendingPlanIterationItemCreateRequestBody) GetScheduledPaymentGuidOk() (*string, bool)`

GetScheduledPaymentGuidOk returns a tuple with the ScheduledPaymentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPaymentGuid

`func (o *SpendingPlanIterationItemCreateRequestBody) SetScheduledPaymentGuid(v string)`

SetScheduledPaymentGuid sets ScheduledPaymentGuid field to given value.

### HasScheduledPaymentGuid

`func (o *SpendingPlanIterationItemCreateRequestBody) HasScheduledPaymentGuid() bool`

HasScheduledPaymentGuid returns a boolean if a field has been set.

### GetTopLevelCategoryGuid

`func (o *SpendingPlanIterationItemCreateRequestBody) GetTopLevelCategoryGuid() string`

GetTopLevelCategoryGuid returns the TopLevelCategoryGuid field if non-nil, zero value otherwise.

### GetTopLevelCategoryGuidOk

`func (o *SpendingPlanIterationItemCreateRequestBody) GetTopLevelCategoryGuidOk() (*string, bool)`

GetTopLevelCategoryGuidOk returns a tuple with the TopLevelCategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevelCategoryGuid

`func (o *SpendingPlanIterationItemCreateRequestBody) SetTopLevelCategoryGuid(v string)`

SetTopLevelCategoryGuid sets TopLevelCategoryGuid field to given value.

### HasTopLevelCategoryGuid

`func (o *SpendingPlanIterationItemCreateRequestBody) HasTopLevelCategoryGuid() bool`

HasTopLevelCategoryGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


