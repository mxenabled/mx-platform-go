# GoalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | Pointer to **string** | Unique identifier of the account for the goal. | [optional] 
**Amount** | Pointer to **float32** | Amount of the goal. | [optional] 
**CurrentAmount** | Pointer to **float32** | The current amount of the goal. | [optional] 
**Guid** | Pointer to **string** | The unique identifier for the goal. Defined by MX. | [optional] 
**GoalTypeName** | Pointer to **string** | The goal type. | [optional] 
**MetaTypeName** | Pointer to **string** | The category of the goal. | [optional] 
**Name** | Pointer to **string** | The name of the goal. | [optional] 
**CompletedAt** | Pointer to **string** | Date and time the goal was completed. | [optional] 
**HasBeenSpent** | Pointer to **bool** | Determines if the goal has been spent. | [optional] 
**IsComplete** | Pointer to **bool** | Determines if the goal is complete. | [optional] 
**Metadata** | Pointer to **string** | Additional information a partner can store on the goal. | [optional] 
**Position** | Pointer to **int32** | The priority of the goal in relation to multiple goals. | [optional] 
**ProjectedToCompleteAt** | Pointer to **string** | The date on which the project was completed. | [optional] 
**TargetedToCompleteAt** | Pointer to **string** |  | [optional] 
**TrackTypeName** | Pointer to **string** |  | [optional] 
**UserGuid** | Pointer to **string** | The unique identifier for the the user. Defined by MX. | [optional] 

## Methods

### NewGoalsResponse

`func NewGoalsResponse() *GoalsResponse`

NewGoalsResponse instantiates a new GoalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalsResponseWithDefaults

`func NewGoalsResponseWithDefaults() *GoalsResponse`

NewGoalsResponseWithDefaults instantiates a new GoalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *GoalsResponse) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *GoalsResponse) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *GoalsResponse) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.

### HasAccountGuid

`func (o *GoalsResponse) HasAccountGuid() bool`

HasAccountGuid returns a boolean if a field has been set.

### GetAmount

`func (o *GoalsResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GoalsResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GoalsResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GoalsResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrentAmount

`func (o *GoalsResponse) GetCurrentAmount() float32`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *GoalsResponse) GetCurrentAmountOk() (*float32, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *GoalsResponse) SetCurrentAmount(v float32)`

SetCurrentAmount sets CurrentAmount field to given value.

### HasCurrentAmount

`func (o *GoalsResponse) HasCurrentAmount() bool`

HasCurrentAmount returns a boolean if a field has been set.

### GetGuid

`func (o *GoalsResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *GoalsResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *GoalsResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *GoalsResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetGoalTypeName

`func (o *GoalsResponse) GetGoalTypeName() string`

GetGoalTypeName returns the GoalTypeName field if non-nil, zero value otherwise.

### GetGoalTypeNameOk

`func (o *GoalsResponse) GetGoalTypeNameOk() (*string, bool)`

GetGoalTypeNameOk returns a tuple with the GoalTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalTypeName

`func (o *GoalsResponse) SetGoalTypeName(v string)`

SetGoalTypeName sets GoalTypeName field to given value.

### HasGoalTypeName

`func (o *GoalsResponse) HasGoalTypeName() bool`

HasGoalTypeName returns a boolean if a field has been set.

### GetMetaTypeName

`func (o *GoalsResponse) GetMetaTypeName() string`

GetMetaTypeName returns the MetaTypeName field if non-nil, zero value otherwise.

### GetMetaTypeNameOk

`func (o *GoalsResponse) GetMetaTypeNameOk() (*string, bool)`

GetMetaTypeNameOk returns a tuple with the MetaTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTypeName

`func (o *GoalsResponse) SetMetaTypeName(v string)`

SetMetaTypeName sets MetaTypeName field to given value.

### HasMetaTypeName

`func (o *GoalsResponse) HasMetaTypeName() bool`

HasMetaTypeName returns a boolean if a field has been set.

### GetName

`func (o *GoalsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoalsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoalsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GoalsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GoalsResponse) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GoalsResponse) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GoalsResponse) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GoalsResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetHasBeenSpent

`func (o *GoalsResponse) GetHasBeenSpent() bool`

GetHasBeenSpent returns the HasBeenSpent field if non-nil, zero value otherwise.

### GetHasBeenSpentOk

`func (o *GoalsResponse) GetHasBeenSpentOk() (*bool, bool)`

GetHasBeenSpentOk returns a tuple with the HasBeenSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenSpent

`func (o *GoalsResponse) SetHasBeenSpent(v bool)`

SetHasBeenSpent sets HasBeenSpent field to given value.

### HasHasBeenSpent

`func (o *GoalsResponse) HasHasBeenSpent() bool`

HasHasBeenSpent returns a boolean if a field has been set.

### GetIsComplete

`func (o *GoalsResponse) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *GoalsResponse) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *GoalsResponse) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *GoalsResponse) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.

### GetMetadata

`func (o *GoalsResponse) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GoalsResponse) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GoalsResponse) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GoalsResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPosition

`func (o *GoalsResponse) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GoalsResponse) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GoalsResponse) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GoalsResponse) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProjectedToCompleteAt

`func (o *GoalsResponse) GetProjectedToCompleteAt() string`

GetProjectedToCompleteAt returns the ProjectedToCompleteAt field if non-nil, zero value otherwise.

### GetProjectedToCompleteAtOk

`func (o *GoalsResponse) GetProjectedToCompleteAtOk() (*string, bool)`

GetProjectedToCompleteAtOk returns a tuple with the ProjectedToCompleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedToCompleteAt

`func (o *GoalsResponse) SetProjectedToCompleteAt(v string)`

SetProjectedToCompleteAt sets ProjectedToCompleteAt field to given value.

### HasProjectedToCompleteAt

`func (o *GoalsResponse) HasProjectedToCompleteAt() bool`

HasProjectedToCompleteAt returns a boolean if a field has been set.

### GetTargetedToCompleteAt

`func (o *GoalsResponse) GetTargetedToCompleteAt() string`

GetTargetedToCompleteAt returns the TargetedToCompleteAt field if non-nil, zero value otherwise.

### GetTargetedToCompleteAtOk

`func (o *GoalsResponse) GetTargetedToCompleteAtOk() (*string, bool)`

GetTargetedToCompleteAtOk returns a tuple with the TargetedToCompleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedToCompleteAt

`func (o *GoalsResponse) SetTargetedToCompleteAt(v string)`

SetTargetedToCompleteAt sets TargetedToCompleteAt field to given value.

### HasTargetedToCompleteAt

`func (o *GoalsResponse) HasTargetedToCompleteAt() bool`

HasTargetedToCompleteAt returns a boolean if a field has been set.

### GetTrackTypeName

`func (o *GoalsResponse) GetTrackTypeName() string`

GetTrackTypeName returns the TrackTypeName field if non-nil, zero value otherwise.

### GetTrackTypeNameOk

`func (o *GoalsResponse) GetTrackTypeNameOk() (*string, bool)`

GetTrackTypeNameOk returns a tuple with the TrackTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTypeName

`func (o *GoalsResponse) SetTrackTypeName(v string)`

SetTrackTypeName sets TrackTypeName field to given value.

### HasTrackTypeName

`func (o *GoalsResponse) HasTrackTypeName() bool`

HasTrackTypeName returns a boolean if a field has been set.

### GetUserGuid

`func (o *GoalsResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *GoalsResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *GoalsResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *GoalsResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


