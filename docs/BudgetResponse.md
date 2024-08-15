# BudgetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | A goal amount set by the user for a category&#39;s transaction total during a month. | [optional] 
**CategoryGuid** | Pointer to **string** | Unique identifier for the budget category. Defined by MX. | [optional] 
**CreatedAt** | Pointer to **string** | Date and time the budget was created, represented in ISO 8601 format with timestamp. | [optional] 
**Guid** | Pointer to **string** | Unique identifier for the budget. Defined by MX. | [optional] 
**IsExceeded** | Pointer to **bool** | If the budget has been exceeded, this field will be true. Otherwise, this field will be false. | [optional] 
**IsOffTrack** | Pointer to **bool** | If the budget is off track, this field will be true. Otherwise, this field will be false. | [optional] 
**Metadata** | Pointer to **NullableString** | Additional information a partner can store on the budget. | [optional] 
**Name** | Pointer to **NullableString** | The name of the budget that is visible to the user (ie \&quot;Food\&quot;, \&quot;Bills\&quot;, \&quot;Entertainment\&quot;, etc). | [optional] 
**OffTrackPercentage** | Pointer to **NullableFloat32** | The percentage amount of off track spending. (Deprecated). | [optional] 
**ParentGuid** | Pointer to **NullableString** | Unique identifier for the parent budget. Defined by MX. | [optional] 
**PercentSpent** | Pointer to **NullableFloat32** | The percentage of a budget that has been spent during the current calendar month Calculated as the transaction total divided by the amount and then multiplied by 100.A value of zero will be returned when &#x60;amount&#x60; is zero. | [optional] 
**ProjectedSpending** | Pointer to **float32** | The projected amount of spending for the budget. | [optional] 
**Revision** | Pointer to **int32** | The revision number of this budget record. | [optional] 
**TransactionTotal** | Pointer to **interface{}** | The cumulative amount of all transactions under the budget. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Date and time the budget was updated, represented in ISO 8601 format with timestamp. | [optional] 
**UserGuid** | Pointer to **interface{}** | Unique identifier for the user. Defined by MX. | [optional] 

## Methods

### NewBudgetResponse

`func NewBudgetResponse() *BudgetResponse`

NewBudgetResponse instantiates a new BudgetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetResponseWithDefaults

`func NewBudgetResponseWithDefaults() *BudgetResponse`

NewBudgetResponseWithDefaults instantiates a new BudgetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BudgetResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BudgetResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BudgetResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BudgetResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCategoryGuid

`func (o *BudgetResponse) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *BudgetResponse) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *BudgetResponse) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *BudgetResponse) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BudgetResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BudgetResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BudgetResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BudgetResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *BudgetResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *BudgetResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *BudgetResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *BudgetResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetIsExceeded

`func (o *BudgetResponse) GetIsExceeded() bool`

GetIsExceeded returns the IsExceeded field if non-nil, zero value otherwise.

### GetIsExceededOk

`func (o *BudgetResponse) GetIsExceededOk() (*bool, bool)`

GetIsExceededOk returns a tuple with the IsExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExceeded

`func (o *BudgetResponse) SetIsExceeded(v bool)`

SetIsExceeded sets IsExceeded field to given value.

### HasIsExceeded

`func (o *BudgetResponse) HasIsExceeded() bool`

HasIsExceeded returns a boolean if a field has been set.

### GetIsOffTrack

`func (o *BudgetResponse) GetIsOffTrack() bool`

GetIsOffTrack returns the IsOffTrack field if non-nil, zero value otherwise.

### GetIsOffTrackOk

`func (o *BudgetResponse) GetIsOffTrackOk() (*bool, bool)`

GetIsOffTrackOk returns a tuple with the IsOffTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOffTrack

`func (o *BudgetResponse) SetIsOffTrack(v bool)`

SetIsOffTrack sets IsOffTrack field to given value.

### HasIsOffTrack

`func (o *BudgetResponse) HasIsOffTrack() bool`

HasIsOffTrack returns a boolean if a field has been set.

### GetMetadata

`func (o *BudgetResponse) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BudgetResponse) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BudgetResponse) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BudgetResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BudgetResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BudgetResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *BudgetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BudgetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BudgetResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BudgetResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BudgetResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BudgetResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOffTrackPercentage

`func (o *BudgetResponse) GetOffTrackPercentage() float32`

GetOffTrackPercentage returns the OffTrackPercentage field if non-nil, zero value otherwise.

### GetOffTrackPercentageOk

`func (o *BudgetResponse) GetOffTrackPercentageOk() (*float32, bool)`

GetOffTrackPercentageOk returns a tuple with the OffTrackPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffTrackPercentage

`func (o *BudgetResponse) SetOffTrackPercentage(v float32)`

SetOffTrackPercentage sets OffTrackPercentage field to given value.

### HasOffTrackPercentage

`func (o *BudgetResponse) HasOffTrackPercentage() bool`

HasOffTrackPercentage returns a boolean if a field has been set.

### SetOffTrackPercentageNil

`func (o *BudgetResponse) SetOffTrackPercentageNil(b bool)`

 SetOffTrackPercentageNil sets the value for OffTrackPercentage to be an explicit nil

### UnsetOffTrackPercentage
`func (o *BudgetResponse) UnsetOffTrackPercentage()`

UnsetOffTrackPercentage ensures that no value is present for OffTrackPercentage, not even an explicit nil
### GetParentGuid

`func (o *BudgetResponse) GetParentGuid() string`

GetParentGuid returns the ParentGuid field if non-nil, zero value otherwise.

### GetParentGuidOk

`func (o *BudgetResponse) GetParentGuidOk() (*string, bool)`

GetParentGuidOk returns a tuple with the ParentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGuid

`func (o *BudgetResponse) SetParentGuid(v string)`

SetParentGuid sets ParentGuid field to given value.

### HasParentGuid

`func (o *BudgetResponse) HasParentGuid() bool`

HasParentGuid returns a boolean if a field has been set.

### SetParentGuidNil

`func (o *BudgetResponse) SetParentGuidNil(b bool)`

 SetParentGuidNil sets the value for ParentGuid to be an explicit nil

### UnsetParentGuid
`func (o *BudgetResponse) UnsetParentGuid()`

UnsetParentGuid ensures that no value is present for ParentGuid, not even an explicit nil
### GetPercentSpent

`func (o *BudgetResponse) GetPercentSpent() float32`

GetPercentSpent returns the PercentSpent field if non-nil, zero value otherwise.

### GetPercentSpentOk

`func (o *BudgetResponse) GetPercentSpentOk() (*float32, bool)`

GetPercentSpentOk returns a tuple with the PercentSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentSpent

`func (o *BudgetResponse) SetPercentSpent(v float32)`

SetPercentSpent sets PercentSpent field to given value.

### HasPercentSpent

`func (o *BudgetResponse) HasPercentSpent() bool`

HasPercentSpent returns a boolean if a field has been set.

### SetPercentSpentNil

`func (o *BudgetResponse) SetPercentSpentNil(b bool)`

 SetPercentSpentNil sets the value for PercentSpent to be an explicit nil

### UnsetPercentSpent
`func (o *BudgetResponse) UnsetPercentSpent()`

UnsetPercentSpent ensures that no value is present for PercentSpent, not even an explicit nil
### GetProjectedSpending

`func (o *BudgetResponse) GetProjectedSpending() float32`

GetProjectedSpending returns the ProjectedSpending field if non-nil, zero value otherwise.

### GetProjectedSpendingOk

`func (o *BudgetResponse) GetProjectedSpendingOk() (*float32, bool)`

GetProjectedSpendingOk returns a tuple with the ProjectedSpending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedSpending

`func (o *BudgetResponse) SetProjectedSpending(v float32)`

SetProjectedSpending sets ProjectedSpending field to given value.

### HasProjectedSpending

`func (o *BudgetResponse) HasProjectedSpending() bool`

HasProjectedSpending returns a boolean if a field has been set.

### GetRevision

`func (o *BudgetResponse) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BudgetResponse) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BudgetResponse) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BudgetResponse) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetTransactionTotal

`func (o *BudgetResponse) GetTransactionTotal() interface{}`

GetTransactionTotal returns the TransactionTotal field if non-nil, zero value otherwise.

### GetTransactionTotalOk

`func (o *BudgetResponse) GetTransactionTotalOk() (*interface{}, bool)`

GetTransactionTotalOk returns a tuple with the TransactionTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTotal

`func (o *BudgetResponse) SetTransactionTotal(v interface{})`

SetTransactionTotal sets TransactionTotal field to given value.

### HasTransactionTotal

`func (o *BudgetResponse) HasTransactionTotal() bool`

HasTransactionTotal returns a boolean if a field has been set.

### SetTransactionTotalNil

`func (o *BudgetResponse) SetTransactionTotalNil(b bool)`

 SetTransactionTotalNil sets the value for TransactionTotal to be an explicit nil

### UnsetTransactionTotal
`func (o *BudgetResponse) UnsetTransactionTotal()`

UnsetTransactionTotal ensures that no value is present for TransactionTotal, not even an explicit nil
### GetUpdatedAt

`func (o *BudgetResponse) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BudgetResponse) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BudgetResponse) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BudgetResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BudgetResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BudgetResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserGuid

`func (o *BudgetResponse) GetUserGuid() interface{}`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *BudgetResponse) GetUserGuidOk() (*interface{}, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *BudgetResponse) SetUserGuid(v interface{})`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *BudgetResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### SetUserGuidNil

`func (o *BudgetResponse) SetUserGuidNil(b bool)`

 SetUserGuidNil sets the value for UserGuid to be an explicit nil

### UnsetUserGuid
`func (o *BudgetResponse) UnsetUserGuid()`

UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


