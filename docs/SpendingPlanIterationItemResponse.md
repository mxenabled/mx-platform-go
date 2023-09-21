# SpendingPlanIterationItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualAmount** | Pointer to **NullableFloat32** |  | [optional] 
**CategoryGuid** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**Guid** | Pointer to **NullableString** |  | [optional] 
**ItemType** | Pointer to **NullableString** |  | [optional] 
**PlannedAmount** | Pointer to **NullableFloat32** |  | [optional] 
**ScheduledPaymentGuid** | Pointer to **NullableString** |  | [optional] 
**SpendingPlanIterationGuid** | Pointer to **NullableString** |  | [optional] 
**TopLevelCategoryGuid** | Pointer to **NullableString** |  | [optional] 
**TransactionGuids** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**UserGuid** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSpendingPlanIterationItemResponse

`func NewSpendingPlanIterationItemResponse() *SpendingPlanIterationItemResponse`

NewSpendingPlanIterationItemResponse instantiates a new SpendingPlanIterationItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendingPlanIterationItemResponseWithDefaults

`func NewSpendingPlanIterationItemResponseWithDefaults() *SpendingPlanIterationItemResponse`

NewSpendingPlanIterationItemResponseWithDefaults instantiates a new SpendingPlanIterationItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualAmount

`func (o *SpendingPlanIterationItemResponse) GetActualAmount() float32`

GetActualAmount returns the ActualAmount field if non-nil, zero value otherwise.

### GetActualAmountOk

`func (o *SpendingPlanIterationItemResponse) GetActualAmountOk() (*float32, bool)`

GetActualAmountOk returns a tuple with the ActualAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualAmount

`func (o *SpendingPlanIterationItemResponse) SetActualAmount(v float32)`

SetActualAmount sets ActualAmount field to given value.

### HasActualAmount

`func (o *SpendingPlanIterationItemResponse) HasActualAmount() bool`

HasActualAmount returns a boolean if a field has been set.

### SetActualAmountNil

`func (o *SpendingPlanIterationItemResponse) SetActualAmountNil(b bool)`

 SetActualAmountNil sets the value for ActualAmount to be an explicit nil

### UnsetActualAmount
`func (o *SpendingPlanIterationItemResponse) UnsetActualAmount()`

UnsetActualAmount ensures that no value is present for ActualAmount, not even an explicit nil
### GetCategoryGuid

`func (o *SpendingPlanIterationItemResponse) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *SpendingPlanIterationItemResponse) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *SpendingPlanIterationItemResponse) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *SpendingPlanIterationItemResponse) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### SetCategoryGuidNil

`func (o *SpendingPlanIterationItemResponse) SetCategoryGuidNil(b bool)`

 SetCategoryGuidNil sets the value for CategoryGuid to be an explicit nil

### UnsetCategoryGuid
`func (o *SpendingPlanIterationItemResponse) UnsetCategoryGuid()`

UnsetCategoryGuid ensures that no value is present for CategoryGuid, not even an explicit nil
### GetCreatedAt

`func (o *SpendingPlanIterationItemResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpendingPlanIterationItemResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpendingPlanIterationItemResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SpendingPlanIterationItemResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *SpendingPlanIterationItemResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SpendingPlanIterationItemResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetGuid

`func (o *SpendingPlanIterationItemResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *SpendingPlanIterationItemResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *SpendingPlanIterationItemResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *SpendingPlanIterationItemResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *SpendingPlanIterationItemResponse) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *SpendingPlanIterationItemResponse) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetItemType

`func (o *SpendingPlanIterationItemResponse) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *SpendingPlanIterationItemResponse) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *SpendingPlanIterationItemResponse) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *SpendingPlanIterationItemResponse) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *SpendingPlanIterationItemResponse) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *SpendingPlanIterationItemResponse) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetPlannedAmount

`func (o *SpendingPlanIterationItemResponse) GetPlannedAmount() float32`

GetPlannedAmount returns the PlannedAmount field if non-nil, zero value otherwise.

### GetPlannedAmountOk

`func (o *SpendingPlanIterationItemResponse) GetPlannedAmountOk() (*float32, bool)`

GetPlannedAmountOk returns a tuple with the PlannedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedAmount

`func (o *SpendingPlanIterationItemResponse) SetPlannedAmount(v float32)`

SetPlannedAmount sets PlannedAmount field to given value.

### HasPlannedAmount

`func (o *SpendingPlanIterationItemResponse) HasPlannedAmount() bool`

HasPlannedAmount returns a boolean if a field has been set.

### SetPlannedAmountNil

`func (o *SpendingPlanIterationItemResponse) SetPlannedAmountNil(b bool)`

 SetPlannedAmountNil sets the value for PlannedAmount to be an explicit nil

### UnsetPlannedAmount
`func (o *SpendingPlanIterationItemResponse) UnsetPlannedAmount()`

UnsetPlannedAmount ensures that no value is present for PlannedAmount, not even an explicit nil
### GetScheduledPaymentGuid

`func (o *SpendingPlanIterationItemResponse) GetScheduledPaymentGuid() string`

GetScheduledPaymentGuid returns the ScheduledPaymentGuid field if non-nil, zero value otherwise.

### GetScheduledPaymentGuidOk

`func (o *SpendingPlanIterationItemResponse) GetScheduledPaymentGuidOk() (*string, bool)`

GetScheduledPaymentGuidOk returns a tuple with the ScheduledPaymentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPaymentGuid

`func (o *SpendingPlanIterationItemResponse) SetScheduledPaymentGuid(v string)`

SetScheduledPaymentGuid sets ScheduledPaymentGuid field to given value.

### HasScheduledPaymentGuid

`func (o *SpendingPlanIterationItemResponse) HasScheduledPaymentGuid() bool`

HasScheduledPaymentGuid returns a boolean if a field has been set.

### SetScheduledPaymentGuidNil

`func (o *SpendingPlanIterationItemResponse) SetScheduledPaymentGuidNil(b bool)`

 SetScheduledPaymentGuidNil sets the value for ScheduledPaymentGuid to be an explicit nil

### UnsetScheduledPaymentGuid
`func (o *SpendingPlanIterationItemResponse) UnsetScheduledPaymentGuid()`

UnsetScheduledPaymentGuid ensures that no value is present for ScheduledPaymentGuid, not even an explicit nil
### GetSpendingPlanIterationGuid

`func (o *SpendingPlanIterationItemResponse) GetSpendingPlanIterationGuid() string`

GetSpendingPlanIterationGuid returns the SpendingPlanIterationGuid field if non-nil, zero value otherwise.

### GetSpendingPlanIterationGuidOk

`func (o *SpendingPlanIterationItemResponse) GetSpendingPlanIterationGuidOk() (*string, bool)`

GetSpendingPlanIterationGuidOk returns a tuple with the SpendingPlanIterationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendingPlanIterationGuid

`func (o *SpendingPlanIterationItemResponse) SetSpendingPlanIterationGuid(v string)`

SetSpendingPlanIterationGuid sets SpendingPlanIterationGuid field to given value.

### HasSpendingPlanIterationGuid

`func (o *SpendingPlanIterationItemResponse) HasSpendingPlanIterationGuid() bool`

HasSpendingPlanIterationGuid returns a boolean if a field has been set.

### SetSpendingPlanIterationGuidNil

`func (o *SpendingPlanIterationItemResponse) SetSpendingPlanIterationGuidNil(b bool)`

 SetSpendingPlanIterationGuidNil sets the value for SpendingPlanIterationGuid to be an explicit nil

### UnsetSpendingPlanIterationGuid
`func (o *SpendingPlanIterationItemResponse) UnsetSpendingPlanIterationGuid()`

UnsetSpendingPlanIterationGuid ensures that no value is present for SpendingPlanIterationGuid, not even an explicit nil
### GetTopLevelCategoryGuid

`func (o *SpendingPlanIterationItemResponse) GetTopLevelCategoryGuid() string`

GetTopLevelCategoryGuid returns the TopLevelCategoryGuid field if non-nil, zero value otherwise.

### GetTopLevelCategoryGuidOk

`func (o *SpendingPlanIterationItemResponse) GetTopLevelCategoryGuidOk() (*string, bool)`

GetTopLevelCategoryGuidOk returns a tuple with the TopLevelCategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevelCategoryGuid

`func (o *SpendingPlanIterationItemResponse) SetTopLevelCategoryGuid(v string)`

SetTopLevelCategoryGuid sets TopLevelCategoryGuid field to given value.

### HasTopLevelCategoryGuid

`func (o *SpendingPlanIterationItemResponse) HasTopLevelCategoryGuid() bool`

HasTopLevelCategoryGuid returns a boolean if a field has been set.

### SetTopLevelCategoryGuidNil

`func (o *SpendingPlanIterationItemResponse) SetTopLevelCategoryGuidNil(b bool)`

 SetTopLevelCategoryGuidNil sets the value for TopLevelCategoryGuid to be an explicit nil

### UnsetTopLevelCategoryGuid
`func (o *SpendingPlanIterationItemResponse) UnsetTopLevelCategoryGuid()`

UnsetTopLevelCategoryGuid ensures that no value is present for TopLevelCategoryGuid, not even an explicit nil
### GetTransactionGuids

`func (o *SpendingPlanIterationItemResponse) GetTransactionGuids() []*string`

GetTransactionGuids returns the TransactionGuids field if non-nil, zero value otherwise.

### GetTransactionGuidsOk

`func (o *SpendingPlanIterationItemResponse) GetTransactionGuidsOk() (*[]*string, bool)`

GetTransactionGuidsOk returns a tuple with the TransactionGuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionGuids

`func (o *SpendingPlanIterationItemResponse) SetTransactionGuids(v []*string)`

SetTransactionGuids sets TransactionGuids field to given value.

### HasTransactionGuids

`func (o *SpendingPlanIterationItemResponse) HasTransactionGuids() bool`

HasTransactionGuids returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SpendingPlanIterationItemResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SpendingPlanIterationItemResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SpendingPlanIterationItemResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SpendingPlanIterationItemResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *SpendingPlanIterationItemResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SpendingPlanIterationItemResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserGuid

`func (o *SpendingPlanIterationItemResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *SpendingPlanIterationItemResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *SpendingPlanIterationItemResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *SpendingPlanIterationItemResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### SetUserGuidNil

`func (o *SpendingPlanIterationItemResponse) SetUserGuidNil(b bool)`

 SetUserGuidNil sets the value for UserGuid to be an explicit nil

### UnsetUserGuid
`func (o *SpendingPlanIterationItemResponse) UnsetUserGuid()`

UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


