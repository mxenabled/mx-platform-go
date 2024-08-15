/*
MX Platform API

The MX Platform API is a powerful, fully-featured API designed to make aggregating and enhancing financial data easy and reliable. It can seamlessly connect your app or website to tens of thousands of financial institutions.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mxplatformgo

import (
	"encoding/json"
)

// checks if the GoalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoalResponse{}

// GoalResponse struct for GoalResponse
type GoalResponse struct {
	// Unique identifier of the account for the goal.
	AccountGuid *string `json:"account_guid,omitempty"`
	// Amount of the goal.
	Amount *float32 `json:"amount,omitempty"`
	// Date and time the goal was completed.
	CompletedAt *string `json:"completed_at,omitempty"`
	// The current amount of the goal.
	CurrentAmount *float32 `json:"current_amount,omitempty"`
	// The goal type.
	GoalTypeName *string `json:"goal_type_name,omitempty"`
	// Unique identifier for the goal. Defined by MX.
	Guid *string `json:"guid,omitempty"`
	// Determines if the goal has been spent.
	HasBeenSpent *bool `json:"has_been_spent,omitempty"`
	// Determines if the goal is complete.
	IsComplete *bool `json:"is_complete,omitempty"`
	// Additional information a partner can store on the goal.
	Metadata *string `json:"metadata,omitempty"`
	// The category of the goal.
	MetaTypeName *string `json:"meta_type_name,omitempty"`
	// The name of the goal.
	Name *string `json:"name,omitempty"`
	// The priority of the goal in relation to multiple goals.
	Position *int32 `json:"position,omitempty"`
	// Date and time the goal is projected to be completed.
	ProjectedToCompleteAt *string `json:"projected_to_complete_at,omitempty"`
	// Date and time the goal is to complete. Intended for users to set their own goal completion dates.
	TargetedToCompleteAt *string `json:"targeted_to_complete_at,omitempty"`
	TrackTypeName *string `json:"track_type_name,omitempty"`
	// The unique identifier for the the user. Defined by MX.
	UserGuid *string `json:"user_guid,omitempty"`
}

// NewGoalResponse instantiates a new GoalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoalResponse() *GoalResponse {
	this := GoalResponse{}
	return &this
}

// NewGoalResponseWithDefaults instantiates a new GoalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoalResponseWithDefaults() *GoalResponse {
	this := GoalResponse{}
	return &this
}

// GetAccountGuid returns the AccountGuid field value if set, zero value otherwise.
func (o *GoalResponse) GetAccountGuid() string {
	if o == nil || IsNil(o.AccountGuid) {
		var ret string
		return ret
	}
	return *o.AccountGuid
}

// GetAccountGuidOk returns a tuple with the AccountGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetAccountGuidOk() (*string, bool) {
	if o == nil || IsNil(o.AccountGuid) {
		return nil, false
	}
	return o.AccountGuid, true
}

// HasAccountGuid returns a boolean if a field has been set.
func (o *GoalResponse) HasAccountGuid() bool {
	if o != nil && !IsNil(o.AccountGuid) {
		return true
	}

	return false
}

// SetAccountGuid gets a reference to the given string and assigns it to the AccountGuid field.
func (o *GoalResponse) SetAccountGuid(v string) {
	o.AccountGuid = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GoalResponse) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GoalResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *GoalResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *GoalResponse) GetCompletedAt() string {
	if o == nil || IsNil(o.CompletedAt) {
		var ret string
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetCompletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CompletedAt) {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *GoalResponse) HasCompletedAt() bool {
	if o != nil && !IsNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given string and assigns it to the CompletedAt field.
func (o *GoalResponse) SetCompletedAt(v string) {
	o.CompletedAt = &v
}

// GetCurrentAmount returns the CurrentAmount field value if set, zero value otherwise.
func (o *GoalResponse) GetCurrentAmount() float32 {
	if o == nil || IsNil(o.CurrentAmount) {
		var ret float32
		return ret
	}
	return *o.CurrentAmount
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetCurrentAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CurrentAmount) {
		return nil, false
	}
	return o.CurrentAmount, true
}

// HasCurrentAmount returns a boolean if a field has been set.
func (o *GoalResponse) HasCurrentAmount() bool {
	if o != nil && !IsNil(o.CurrentAmount) {
		return true
	}

	return false
}

// SetCurrentAmount gets a reference to the given float32 and assigns it to the CurrentAmount field.
func (o *GoalResponse) SetCurrentAmount(v float32) {
	o.CurrentAmount = &v
}

// GetGoalTypeName returns the GoalTypeName field value if set, zero value otherwise.
func (o *GoalResponse) GetGoalTypeName() string {
	if o == nil || IsNil(o.GoalTypeName) {
		var ret string
		return ret
	}
	return *o.GoalTypeName
}

// GetGoalTypeNameOk returns a tuple with the GoalTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetGoalTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.GoalTypeName) {
		return nil, false
	}
	return o.GoalTypeName, true
}

// HasGoalTypeName returns a boolean if a field has been set.
func (o *GoalResponse) HasGoalTypeName() bool {
	if o != nil && !IsNil(o.GoalTypeName) {
		return true
	}

	return false
}

// SetGoalTypeName gets a reference to the given string and assigns it to the GoalTypeName field.
func (o *GoalResponse) SetGoalTypeName(v string) {
	o.GoalTypeName = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *GoalResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *GoalResponse) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *GoalResponse) SetGuid(v string) {
	o.Guid = &v
}

// GetHasBeenSpent returns the HasBeenSpent field value if set, zero value otherwise.
func (o *GoalResponse) GetHasBeenSpent() bool {
	if o == nil || IsNil(o.HasBeenSpent) {
		var ret bool
		return ret
	}
	return *o.HasBeenSpent
}

// GetHasBeenSpentOk returns a tuple with the HasBeenSpent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetHasBeenSpentOk() (*bool, bool) {
	if o == nil || IsNil(o.HasBeenSpent) {
		return nil, false
	}
	return o.HasBeenSpent, true
}

// HasHasBeenSpent returns a boolean if a field has been set.
func (o *GoalResponse) HasHasBeenSpent() bool {
	if o != nil && !IsNil(o.HasBeenSpent) {
		return true
	}

	return false
}

// SetHasBeenSpent gets a reference to the given bool and assigns it to the HasBeenSpent field.
func (o *GoalResponse) SetHasBeenSpent(v bool) {
	o.HasBeenSpent = &v
}

// GetIsComplete returns the IsComplete field value if set, zero value otherwise.
func (o *GoalResponse) GetIsComplete() bool {
	if o == nil || IsNil(o.IsComplete) {
		var ret bool
		return ret
	}
	return *o.IsComplete
}

// GetIsCompleteOk returns a tuple with the IsComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetIsCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsComplete) {
		return nil, false
	}
	return o.IsComplete, true
}

// HasIsComplete returns a boolean if a field has been set.
func (o *GoalResponse) HasIsComplete() bool {
	if o != nil && !IsNil(o.IsComplete) {
		return true
	}

	return false
}

// SetIsComplete gets a reference to the given bool and assigns it to the IsComplete field.
func (o *GoalResponse) SetIsComplete(v bool) {
	o.IsComplete = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GoalResponse) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GoalResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *GoalResponse) SetMetadata(v string) {
	o.Metadata = &v
}

// GetMetaTypeName returns the MetaTypeName field value if set, zero value otherwise.
func (o *GoalResponse) GetMetaTypeName() string {
	if o == nil || IsNil(o.MetaTypeName) {
		var ret string
		return ret
	}
	return *o.MetaTypeName
}

// GetMetaTypeNameOk returns a tuple with the MetaTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetMetaTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetaTypeName) {
		return nil, false
	}
	return o.MetaTypeName, true
}

// HasMetaTypeName returns a boolean if a field has been set.
func (o *GoalResponse) HasMetaTypeName() bool {
	if o != nil && !IsNil(o.MetaTypeName) {
		return true
	}

	return false
}

// SetMetaTypeName gets a reference to the given string and assigns it to the MetaTypeName field.
func (o *GoalResponse) SetMetaTypeName(v string) {
	o.MetaTypeName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GoalResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GoalResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GoalResponse) SetName(v string) {
	o.Name = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *GoalResponse) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *GoalResponse) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *GoalResponse) SetPosition(v int32) {
	o.Position = &v
}

// GetProjectedToCompleteAt returns the ProjectedToCompleteAt field value if set, zero value otherwise.
func (o *GoalResponse) GetProjectedToCompleteAt() string {
	if o == nil || IsNil(o.ProjectedToCompleteAt) {
		var ret string
		return ret
	}
	return *o.ProjectedToCompleteAt
}

// GetProjectedToCompleteAtOk returns a tuple with the ProjectedToCompleteAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetProjectedToCompleteAtOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectedToCompleteAt) {
		return nil, false
	}
	return o.ProjectedToCompleteAt, true
}

// HasProjectedToCompleteAt returns a boolean if a field has been set.
func (o *GoalResponse) HasProjectedToCompleteAt() bool {
	if o != nil && !IsNil(o.ProjectedToCompleteAt) {
		return true
	}

	return false
}

// SetProjectedToCompleteAt gets a reference to the given string and assigns it to the ProjectedToCompleteAt field.
func (o *GoalResponse) SetProjectedToCompleteAt(v string) {
	o.ProjectedToCompleteAt = &v
}

// GetTargetedToCompleteAt returns the TargetedToCompleteAt field value if set, zero value otherwise.
func (o *GoalResponse) GetTargetedToCompleteAt() string {
	if o == nil || IsNil(o.TargetedToCompleteAt) {
		var ret string
		return ret
	}
	return *o.TargetedToCompleteAt
}

// GetTargetedToCompleteAtOk returns a tuple with the TargetedToCompleteAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetTargetedToCompleteAtOk() (*string, bool) {
	if o == nil || IsNil(o.TargetedToCompleteAt) {
		return nil, false
	}
	return o.TargetedToCompleteAt, true
}

// HasTargetedToCompleteAt returns a boolean if a field has been set.
func (o *GoalResponse) HasTargetedToCompleteAt() bool {
	if o != nil && !IsNil(o.TargetedToCompleteAt) {
		return true
	}

	return false
}

// SetTargetedToCompleteAt gets a reference to the given string and assigns it to the TargetedToCompleteAt field.
func (o *GoalResponse) SetTargetedToCompleteAt(v string) {
	o.TargetedToCompleteAt = &v
}

// GetTrackTypeName returns the TrackTypeName field value if set, zero value otherwise.
func (o *GoalResponse) GetTrackTypeName() string {
	if o == nil || IsNil(o.TrackTypeName) {
		var ret string
		return ret
	}
	return *o.TrackTypeName
}

// GetTrackTypeNameOk returns a tuple with the TrackTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetTrackTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.TrackTypeName) {
		return nil, false
	}
	return o.TrackTypeName, true
}

// HasTrackTypeName returns a boolean if a field has been set.
func (o *GoalResponse) HasTrackTypeName() bool {
	if o != nil && !IsNil(o.TrackTypeName) {
		return true
	}

	return false
}

// SetTrackTypeName gets a reference to the given string and assigns it to the TrackTypeName field.
func (o *GoalResponse) SetTrackTypeName(v string) {
	o.TrackTypeName = &v
}

// GetUserGuid returns the UserGuid field value if set, zero value otherwise.
func (o *GoalResponse) GetUserGuid() string {
	if o == nil || IsNil(o.UserGuid) {
		var ret string
		return ret
	}
	return *o.UserGuid
}

// GetUserGuidOk returns a tuple with the UserGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalResponse) GetUserGuidOk() (*string, bool) {
	if o == nil || IsNil(o.UserGuid) {
		return nil, false
	}
	return o.UserGuid, true
}

// HasUserGuid returns a boolean if a field has been set.
func (o *GoalResponse) HasUserGuid() bool {
	if o != nil && !IsNil(o.UserGuid) {
		return true
	}

	return false
}

// SetUserGuid gets a reference to the given string and assigns it to the UserGuid field.
func (o *GoalResponse) SetUserGuid(v string) {
	o.UserGuid = &v
}

func (o GoalResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountGuid) {
		toSerialize["account_guid"] = o.AccountGuid
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CompletedAt) {
		toSerialize["completed_at"] = o.CompletedAt
	}
	if !IsNil(o.CurrentAmount) {
		toSerialize["current_amount"] = o.CurrentAmount
	}
	if !IsNil(o.GoalTypeName) {
		toSerialize["goal_type_name"] = o.GoalTypeName
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.HasBeenSpent) {
		toSerialize["has_been_spent"] = o.HasBeenSpent
	}
	if !IsNil(o.IsComplete) {
		toSerialize["is_complete"] = o.IsComplete
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MetaTypeName) {
		toSerialize["meta_type_name"] = o.MetaTypeName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.ProjectedToCompleteAt) {
		toSerialize["projected_to_complete_at"] = o.ProjectedToCompleteAt
	}
	if !IsNil(o.TargetedToCompleteAt) {
		toSerialize["targeted_to_complete_at"] = o.TargetedToCompleteAt
	}
	if !IsNil(o.TrackTypeName) {
		toSerialize["track_type_name"] = o.TrackTypeName
	}
	if !IsNil(o.UserGuid) {
		toSerialize["user_guid"] = o.UserGuid
	}
	return toSerialize, nil
}

type NullableGoalResponse struct {
	value *GoalResponse
	isSet bool
}

func (v NullableGoalResponse) Get() *GoalResponse {
	return v.value
}

func (v *NullableGoalResponse) Set(val *GoalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGoalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGoalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoalResponse(val *GoalResponse) *NullableGoalResponse {
	return &NullableGoalResponse{value: val, isSet: true}
}

func (v NullableGoalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


