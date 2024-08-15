# GoalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | Pointer to **string** | Unique identifier of the account for the goal. | [optional] 
**Amount** | Pointer to **float32** | Amount of the goal. | [optional] 
**CompletedAt** | Pointer to **string** | Date and time the goal was completed. | [optional] 
**CurrentAmount** | Pointer to **float32** | The current amount of the goal. | [optional] 
**GoalTypeName** | Pointer to **string** | The goal type. | [optional] 
**Guid** | Pointer to **string** | Unique identifier for the goal. Defined by MX. | [optional] 
**HasBeenSpent** | Pointer to **bool** | Determines if the goal has been spent. | [optional] 
**IsComplete** | Pointer to **bool** | Determines if the goal is complete. | [optional] 
**Metadata** | Pointer to **string** | Additional information a partner can store on the goal. | [optional] 
**MetaTypeName** | Pointer to **string** | The category of the goal. | [optional] 
**Name** | Pointer to **string** | The name of the goal. | [optional] 
**Position** | Pointer to **int32** | The priority of the goal in relation to multiple goals. | [optional] 
**ProjectedToCompleteAt** | Pointer to **string** | Date and time the goal is projected to be completed. | [optional] 
**TargetedToCompleteAt** | Pointer to **string** | Date and time the goal is to complete. Intended for users to set their own goal completion dates. | [optional] 
**TrackTypeName** | Pointer to **string** |  | [optional] 
**UserGuid** | Pointer to **string** | The unique identifier for the the user. Defined by MX. | [optional] 

## Methods

### NewGoalResponse

`func NewGoalResponse() *GoalResponse`

NewGoalResponse instantiates a new GoalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalResponseWithDefaults

`func NewGoalResponseWithDefaults() *GoalResponse`

NewGoalResponseWithDefaults instantiates a new GoalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *GoalResponse) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *GoalResponse) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *GoalResponse) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.

### HasAccountGuid

`func (o *GoalResponse) HasAccountGuid() bool`

HasAccountGuid returns a boolean if a field has been set.

### GetAmount

`func (o *GoalResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GoalResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GoalResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GoalResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GoalResponse) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GoalResponse) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GoalResponse) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GoalResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCurrentAmount

`func (o *GoalResponse) GetCurrentAmount() float32`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *GoalResponse) GetCurrentAmountOk() (*float32, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *GoalResponse) SetCurrentAmount(v float32)`

SetCurrentAmount sets CurrentAmount field to given value.

### HasCurrentAmount

`func (o *GoalResponse) HasCurrentAmount() bool`

HasCurrentAmount returns a boolean if a field has been set.

### GetGoalTypeName

`func (o *GoalResponse) GetGoalTypeName() string`

GetGoalTypeName returns the GoalTypeName field if non-nil, zero value otherwise.

### GetGoalTypeNameOk

`func (o *GoalResponse) GetGoalTypeNameOk() (*string, bool)`

GetGoalTypeNameOk returns a tuple with the GoalTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalTypeName

`func (o *GoalResponse) SetGoalTypeName(v string)`

SetGoalTypeName sets GoalTypeName field to given value.

### HasGoalTypeName

`func (o *GoalResponse) HasGoalTypeName() bool`

HasGoalTypeName returns a boolean if a field has been set.

### GetGuid

`func (o *GoalResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *GoalResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *GoalResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *GoalResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetHasBeenSpent

`func (o *GoalResponse) GetHasBeenSpent() bool`

GetHasBeenSpent returns the HasBeenSpent field if non-nil, zero value otherwise.

### GetHasBeenSpentOk

`func (o *GoalResponse) GetHasBeenSpentOk() (*bool, bool)`

GetHasBeenSpentOk returns a tuple with the HasBeenSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenSpent

`func (o *GoalResponse) SetHasBeenSpent(v bool)`

SetHasBeenSpent sets HasBeenSpent field to given value.

### HasHasBeenSpent

`func (o *GoalResponse) HasHasBeenSpent() bool`

HasHasBeenSpent returns a boolean if a field has been set.

### GetIsComplete

`func (o *GoalResponse) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *GoalResponse) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *GoalResponse) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *GoalResponse) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.

### GetMetadata

`func (o *GoalResponse) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GoalResponse) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GoalResponse) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GoalResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetaTypeName

`func (o *GoalResponse) GetMetaTypeName() string`

GetMetaTypeName returns the MetaTypeName field if non-nil, zero value otherwise.

### GetMetaTypeNameOk

`func (o *GoalResponse) GetMetaTypeNameOk() (*string, bool)`

GetMetaTypeNameOk returns a tuple with the MetaTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTypeName

`func (o *GoalResponse) SetMetaTypeName(v string)`

SetMetaTypeName sets MetaTypeName field to given value.

### HasMetaTypeName

`func (o *GoalResponse) HasMetaTypeName() bool`

HasMetaTypeName returns a boolean if a field has been set.

### GetName

`func (o *GoalResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoalResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoalResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GoalResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *GoalResponse) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GoalResponse) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GoalResponse) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GoalResponse) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProjectedToCompleteAt

`func (o *GoalResponse) GetProjectedToCompleteAt() string`

GetProjectedToCompleteAt returns the ProjectedToCompleteAt field if non-nil, zero value otherwise.

### GetProjectedToCompleteAtOk

`func (o *GoalResponse) GetProjectedToCompleteAtOk() (*string, bool)`

GetProjectedToCompleteAtOk returns a tuple with the ProjectedToCompleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedToCompleteAt

`func (o *GoalResponse) SetProjectedToCompleteAt(v string)`

SetProjectedToCompleteAt sets ProjectedToCompleteAt field to given value.

### HasProjectedToCompleteAt

`func (o *GoalResponse) HasProjectedToCompleteAt() bool`

HasProjectedToCompleteAt returns a boolean if a field has been set.

### GetTargetedToCompleteAt

`func (o *GoalResponse) GetTargetedToCompleteAt() string`

GetTargetedToCompleteAt returns the TargetedToCompleteAt field if non-nil, zero value otherwise.

### GetTargetedToCompleteAtOk

`func (o *GoalResponse) GetTargetedToCompleteAtOk() (*string, bool)`

GetTargetedToCompleteAtOk returns a tuple with the TargetedToCompleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedToCompleteAt

`func (o *GoalResponse) SetTargetedToCompleteAt(v string)`

SetTargetedToCompleteAt sets TargetedToCompleteAt field to given value.

### HasTargetedToCompleteAt

`func (o *GoalResponse) HasTargetedToCompleteAt() bool`

HasTargetedToCompleteAt returns a boolean if a field has been set.

### GetTrackTypeName

`func (o *GoalResponse) GetTrackTypeName() string`

GetTrackTypeName returns the TrackTypeName field if non-nil, zero value otherwise.

### GetTrackTypeNameOk

`func (o *GoalResponse) GetTrackTypeNameOk() (*string, bool)`

GetTrackTypeNameOk returns a tuple with the TrackTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTypeName

`func (o *GoalResponse) SetTrackTypeName(v string)`

SetTrackTypeName sets TrackTypeName field to given value.

### HasTrackTypeName

`func (o *GoalResponse) HasTrackTypeName() bool`

HasTrackTypeName returns a boolean if a field has been set.

### GetUserGuid

`func (o *GoalResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *GoalResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *GoalResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *GoalResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


