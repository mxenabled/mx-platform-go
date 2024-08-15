# GoalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | **string** | Unique identifier of the account for the goal. | 
**Amount** | **float32** | Amount of the goal. | 
**GoalTypeName** | **string** | The goal type. | 
**MetaTypeName** | **string** | The category of the goal. | 
**Name** | **string** | The name of the goal. | 
**CompletedAt** | Pointer to **string** | Date and time the goal was completed. | [optional] 
**HasBeenSpent** | Pointer to **bool** | Determines if the goal has been spent. | [optional] 
**IsComplete** | Pointer to **bool** | Determines if the goal is complete. | [optional] 
**Metadata** | Pointer to **string** | Additional information a partner can store on the goal. | [optional] 
**Position** | Pointer to **int32** | The priority of the goal in relation to multiple goals. | [optional] 
**TargetedToCompleteAt** | Pointer to **string** | Date and time the goal is to complete. Intended for users to set their own goal completion dates. | [optional] 

## Methods

### NewGoalRequest

`func NewGoalRequest(accountGuid string, amount float32, goalTypeName string, metaTypeName string, name string, ) *GoalRequest`

NewGoalRequest instantiates a new GoalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalRequestWithDefaults

`func NewGoalRequestWithDefaults() *GoalRequest`

NewGoalRequestWithDefaults instantiates a new GoalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *GoalRequest) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *GoalRequest) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *GoalRequest) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.


### GetAmount

`func (o *GoalRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GoalRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GoalRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetGoalTypeName

`func (o *GoalRequest) GetGoalTypeName() string`

GetGoalTypeName returns the GoalTypeName field if non-nil, zero value otherwise.

### GetGoalTypeNameOk

`func (o *GoalRequest) GetGoalTypeNameOk() (*string, bool)`

GetGoalTypeNameOk returns a tuple with the GoalTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalTypeName

`func (o *GoalRequest) SetGoalTypeName(v string)`

SetGoalTypeName sets GoalTypeName field to given value.


### GetMetaTypeName

`func (o *GoalRequest) GetMetaTypeName() string`

GetMetaTypeName returns the MetaTypeName field if non-nil, zero value otherwise.

### GetMetaTypeNameOk

`func (o *GoalRequest) GetMetaTypeNameOk() (*string, bool)`

GetMetaTypeNameOk returns a tuple with the MetaTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTypeName

`func (o *GoalRequest) SetMetaTypeName(v string)`

SetMetaTypeName sets MetaTypeName field to given value.


### GetName

`func (o *GoalRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoalRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoalRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCompletedAt

`func (o *GoalRequest) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GoalRequest) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GoalRequest) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GoalRequest) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetHasBeenSpent

`func (o *GoalRequest) GetHasBeenSpent() bool`

GetHasBeenSpent returns the HasBeenSpent field if non-nil, zero value otherwise.

### GetHasBeenSpentOk

`func (o *GoalRequest) GetHasBeenSpentOk() (*bool, bool)`

GetHasBeenSpentOk returns a tuple with the HasBeenSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenSpent

`func (o *GoalRequest) SetHasBeenSpent(v bool)`

SetHasBeenSpent sets HasBeenSpent field to given value.

### HasHasBeenSpent

`func (o *GoalRequest) HasHasBeenSpent() bool`

HasHasBeenSpent returns a boolean if a field has been set.

### GetIsComplete

`func (o *GoalRequest) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *GoalRequest) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *GoalRequest) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *GoalRequest) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.

### GetMetadata

`func (o *GoalRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GoalRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GoalRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GoalRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPosition

`func (o *GoalRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GoalRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GoalRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GoalRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTargetedToCompleteAt

`func (o *GoalRequest) GetTargetedToCompleteAt() string`

GetTargetedToCompleteAt returns the TargetedToCompleteAt field if non-nil, zero value otherwise.

### GetTargetedToCompleteAtOk

`func (o *GoalRequest) GetTargetedToCompleteAtOk() (*string, bool)`

GetTargetedToCompleteAtOk returns a tuple with the TargetedToCompleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedToCompleteAt

`func (o *GoalRequest) SetTargetedToCompleteAt(v string)`

SetTargetedToCompleteAt sets TargetedToCompleteAt field to given value.

### HasTargetedToCompleteAt

`func (o *GoalRequest) HasTargetedToCompleteAt() bool`

HasTargetedToCompleteAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


