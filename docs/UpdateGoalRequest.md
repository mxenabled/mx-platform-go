# UpdateGoalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | Pointer to **string** | Unique identifier of the account for the goal. | [optional] 
**Amount** | Pointer to **float32** | Amount of the goal. | [optional] 
**GoalTypeName** | Pointer to **string** | The goal type. | [optional] 
**MetaTypeName** | Pointer to **string** | The category of the goal. | [optional] 
**Name** | Pointer to **string** | The name of the goal. | [optional] 
**CompletedAt** | Pointer to **string** | Date and time the goal was completed. | [optional] 
**HasBeenSpent** | Pointer to **bool** | Determines if the goal has been spent. | [optional] 
**IsComplete** | Pointer to **bool** | Determines if the goal is complete. | [optional] 
**Metadata** | Pointer to **string** | Additional information a partner can store on the goal. | [optional] 
**Position** | Pointer to **int32** | The priority of the goal in relation to multiple goals. | [optional] 
**TargetedToCompleteAt** | Pointer to **string** | Date and time the goal is to complete. Intended for users to set their own goal completion dates. | [optional] 

## Methods

### NewUpdateGoalRequest

`func NewUpdateGoalRequest() *UpdateGoalRequest`

NewUpdateGoalRequest instantiates a new UpdateGoalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGoalRequestWithDefaults

`func NewUpdateGoalRequestWithDefaults() *UpdateGoalRequest`

NewUpdateGoalRequestWithDefaults instantiates a new UpdateGoalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *UpdateGoalRequest) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *UpdateGoalRequest) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *UpdateGoalRequest) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.

### HasAccountGuid

`func (o *UpdateGoalRequest) HasAccountGuid() bool`

HasAccountGuid returns a boolean if a field has been set.

### GetAmount

`func (o *UpdateGoalRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UpdateGoalRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UpdateGoalRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UpdateGoalRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetGoalTypeName

`func (o *UpdateGoalRequest) GetGoalTypeName() string`

GetGoalTypeName returns the GoalTypeName field if non-nil, zero value otherwise.

### GetGoalTypeNameOk

`func (o *UpdateGoalRequest) GetGoalTypeNameOk() (*string, bool)`

GetGoalTypeNameOk returns a tuple with the GoalTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalTypeName

`func (o *UpdateGoalRequest) SetGoalTypeName(v string)`

SetGoalTypeName sets GoalTypeName field to given value.

### HasGoalTypeName

`func (o *UpdateGoalRequest) HasGoalTypeName() bool`

HasGoalTypeName returns a boolean if a field has been set.

### GetMetaTypeName

`func (o *UpdateGoalRequest) GetMetaTypeName() string`

GetMetaTypeName returns the MetaTypeName field if non-nil, zero value otherwise.

### GetMetaTypeNameOk

`func (o *UpdateGoalRequest) GetMetaTypeNameOk() (*string, bool)`

GetMetaTypeNameOk returns a tuple with the MetaTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTypeName

`func (o *UpdateGoalRequest) SetMetaTypeName(v string)`

SetMetaTypeName sets MetaTypeName field to given value.

### HasMetaTypeName

`func (o *UpdateGoalRequest) HasMetaTypeName() bool`

HasMetaTypeName returns a boolean if a field has been set.

### GetName

`func (o *UpdateGoalRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGoalRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGoalRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGoalRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCompletedAt

`func (o *UpdateGoalRequest) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *UpdateGoalRequest) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *UpdateGoalRequest) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *UpdateGoalRequest) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetHasBeenSpent

`func (o *UpdateGoalRequest) GetHasBeenSpent() bool`

GetHasBeenSpent returns the HasBeenSpent field if non-nil, zero value otherwise.

### GetHasBeenSpentOk

`func (o *UpdateGoalRequest) GetHasBeenSpentOk() (*bool, bool)`

GetHasBeenSpentOk returns a tuple with the HasBeenSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenSpent

`func (o *UpdateGoalRequest) SetHasBeenSpent(v bool)`

SetHasBeenSpent sets HasBeenSpent field to given value.

### HasHasBeenSpent

`func (o *UpdateGoalRequest) HasHasBeenSpent() bool`

HasHasBeenSpent returns a boolean if a field has been set.

### GetIsComplete

`func (o *UpdateGoalRequest) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *UpdateGoalRequest) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *UpdateGoalRequest) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *UpdateGoalRequest) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateGoalRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateGoalRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateGoalRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateGoalRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPosition

`func (o *UpdateGoalRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UpdateGoalRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UpdateGoalRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UpdateGoalRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTargetedToCompleteAt

`func (o *UpdateGoalRequest) GetTargetedToCompleteAt() string`

GetTargetedToCompleteAt returns the TargetedToCompleteAt field if non-nil, zero value otherwise.

### GetTargetedToCompleteAtOk

`func (o *UpdateGoalRequest) GetTargetedToCompleteAtOk() (*string, bool)`

GetTargetedToCompleteAtOk returns a tuple with the TargetedToCompleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedToCompleteAt

`func (o *UpdateGoalRequest) SetTargetedToCompleteAt(v string)`

SetTargetedToCompleteAt sets TargetedToCompleteAt field to given value.

### HasTargetedToCompleteAt

`func (o *UpdateGoalRequest) HasTargetedToCompleteAt() bool`

HasTargetedToCompleteAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


