# MemberResumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Challenges** | Pointer to [**[]CredentialRequest**](CredentialRequest.md) |  | [optional] 

## Methods

### NewMemberResumeRequest

`func NewMemberResumeRequest() *MemberResumeRequest`

NewMemberResumeRequest instantiates a new MemberResumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberResumeRequestWithDefaults

`func NewMemberResumeRequestWithDefaults() *MemberResumeRequest`

NewMemberResumeRequestWithDefaults instantiates a new MemberResumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenges

`func (o *MemberResumeRequest) GetChallenges() []CredentialRequest`

GetChallenges returns the Challenges field if non-nil, zero value otherwise.

### GetChallengesOk

`func (o *MemberResumeRequest) GetChallengesOk() (*[]CredentialRequest, bool)`

GetChallengesOk returns a tuple with the Challenges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenges

`func (o *MemberResumeRequest) SetChallenges(v []CredentialRequest)`

SetChallenges sets Challenges field to given value.

### HasChallenges

`func (o *MemberResumeRequest) HasChallenges() bool`

HasChallenges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


