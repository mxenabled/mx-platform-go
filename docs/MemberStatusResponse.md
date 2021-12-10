# MemberStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregatedAt** | Pointer to **NullableString** |  | [optional] 
**Challenges** | Pointer to [**[]ChallengeResponse**](ChallengeResponse.md) |  | [optional] 
**ConnectionStatus** | Pointer to **NullableString** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**HasProcessedAccounts** | Pointer to **NullableBool** |  | [optional] 
**HasProcessedTransactions** | Pointer to **NullableBool** |  | [optional] 
**IsAuthenticated** | Pointer to **NullableBool** |  | [optional] 
**IsBeingAggregated** | Pointer to **NullableBool** |  | [optional] 
**SuccessfullyAggregatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMemberStatusResponse

`func NewMemberStatusResponse() *MemberStatusResponse`

NewMemberStatusResponse instantiates a new MemberStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberStatusResponseWithDefaults

`func NewMemberStatusResponseWithDefaults() *MemberStatusResponse`

NewMemberStatusResponseWithDefaults instantiates a new MemberStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatedAt

`func (o *MemberStatusResponse) GetAggregatedAt() string`

GetAggregatedAt returns the AggregatedAt field if non-nil, zero value otherwise.

### GetAggregatedAtOk

`func (o *MemberStatusResponse) GetAggregatedAtOk() (*string, bool)`

GetAggregatedAtOk returns a tuple with the AggregatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedAt

`func (o *MemberStatusResponse) SetAggregatedAt(v string)`

SetAggregatedAt sets AggregatedAt field to given value.

### HasAggregatedAt

`func (o *MemberStatusResponse) HasAggregatedAt() bool`

HasAggregatedAt returns a boolean if a field has been set.

### SetAggregatedAtNil

`func (o *MemberStatusResponse) SetAggregatedAtNil(b bool)`

 SetAggregatedAtNil sets the value for AggregatedAt to be an explicit nil

### UnsetAggregatedAt
`func (o *MemberStatusResponse) UnsetAggregatedAt()`

UnsetAggregatedAt ensures that no value is present for AggregatedAt, not even an explicit nil
### GetChallenges

`func (o *MemberStatusResponse) GetChallenges() []ChallengeResponse`

GetChallenges returns the Challenges field if non-nil, zero value otherwise.

### GetChallengesOk

`func (o *MemberStatusResponse) GetChallengesOk() (*[]ChallengeResponse, bool)`

GetChallengesOk returns a tuple with the Challenges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenges

`func (o *MemberStatusResponse) SetChallenges(v []ChallengeResponse)`

SetChallenges sets Challenges field to given value.

### HasChallenges

`func (o *MemberStatusResponse) HasChallenges() bool`

HasChallenges returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *MemberStatusResponse) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *MemberStatusResponse) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *MemberStatusResponse) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *MemberStatusResponse) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### SetConnectionStatusNil

`func (o *MemberStatusResponse) SetConnectionStatusNil(b bool)`

 SetConnectionStatusNil sets the value for ConnectionStatus to be an explicit nil

### UnsetConnectionStatus
`func (o *MemberStatusResponse) UnsetConnectionStatus()`

UnsetConnectionStatus ensures that no value is present for ConnectionStatus, not even an explicit nil
### GetGuid

`func (o *MemberStatusResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *MemberStatusResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *MemberStatusResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *MemberStatusResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetHasProcessedAccounts

`func (o *MemberStatusResponse) GetHasProcessedAccounts() bool`

GetHasProcessedAccounts returns the HasProcessedAccounts field if non-nil, zero value otherwise.

### GetHasProcessedAccountsOk

`func (o *MemberStatusResponse) GetHasProcessedAccountsOk() (*bool, bool)`

GetHasProcessedAccountsOk returns a tuple with the HasProcessedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProcessedAccounts

`func (o *MemberStatusResponse) SetHasProcessedAccounts(v bool)`

SetHasProcessedAccounts sets HasProcessedAccounts field to given value.

### HasHasProcessedAccounts

`func (o *MemberStatusResponse) HasHasProcessedAccounts() bool`

HasHasProcessedAccounts returns a boolean if a field has been set.

### SetHasProcessedAccountsNil

`func (o *MemberStatusResponse) SetHasProcessedAccountsNil(b bool)`

 SetHasProcessedAccountsNil sets the value for HasProcessedAccounts to be an explicit nil

### UnsetHasProcessedAccounts
`func (o *MemberStatusResponse) UnsetHasProcessedAccounts()`

UnsetHasProcessedAccounts ensures that no value is present for HasProcessedAccounts, not even an explicit nil
### GetHasProcessedTransactions

`func (o *MemberStatusResponse) GetHasProcessedTransactions() bool`

GetHasProcessedTransactions returns the HasProcessedTransactions field if non-nil, zero value otherwise.

### GetHasProcessedTransactionsOk

`func (o *MemberStatusResponse) GetHasProcessedTransactionsOk() (*bool, bool)`

GetHasProcessedTransactionsOk returns a tuple with the HasProcessedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProcessedTransactions

`func (o *MemberStatusResponse) SetHasProcessedTransactions(v bool)`

SetHasProcessedTransactions sets HasProcessedTransactions field to given value.

### HasHasProcessedTransactions

`func (o *MemberStatusResponse) HasHasProcessedTransactions() bool`

HasHasProcessedTransactions returns a boolean if a field has been set.

### SetHasProcessedTransactionsNil

`func (o *MemberStatusResponse) SetHasProcessedTransactionsNil(b bool)`

 SetHasProcessedTransactionsNil sets the value for HasProcessedTransactions to be an explicit nil

### UnsetHasProcessedTransactions
`func (o *MemberStatusResponse) UnsetHasProcessedTransactions()`

UnsetHasProcessedTransactions ensures that no value is present for HasProcessedTransactions, not even an explicit nil
### GetIsAuthenticated

`func (o *MemberStatusResponse) GetIsAuthenticated() bool`

GetIsAuthenticated returns the IsAuthenticated field if non-nil, zero value otherwise.

### GetIsAuthenticatedOk

`func (o *MemberStatusResponse) GetIsAuthenticatedOk() (*bool, bool)`

GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticated

`func (o *MemberStatusResponse) SetIsAuthenticated(v bool)`

SetIsAuthenticated sets IsAuthenticated field to given value.

### HasIsAuthenticated

`func (o *MemberStatusResponse) HasIsAuthenticated() bool`

HasIsAuthenticated returns a boolean if a field has been set.

### SetIsAuthenticatedNil

`func (o *MemberStatusResponse) SetIsAuthenticatedNil(b bool)`

 SetIsAuthenticatedNil sets the value for IsAuthenticated to be an explicit nil

### UnsetIsAuthenticated
`func (o *MemberStatusResponse) UnsetIsAuthenticated()`

UnsetIsAuthenticated ensures that no value is present for IsAuthenticated, not even an explicit nil
### GetIsBeingAggregated

`func (o *MemberStatusResponse) GetIsBeingAggregated() bool`

GetIsBeingAggregated returns the IsBeingAggregated field if non-nil, zero value otherwise.

### GetIsBeingAggregatedOk

`func (o *MemberStatusResponse) GetIsBeingAggregatedOk() (*bool, bool)`

GetIsBeingAggregatedOk returns a tuple with the IsBeingAggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBeingAggregated

`func (o *MemberStatusResponse) SetIsBeingAggregated(v bool)`

SetIsBeingAggregated sets IsBeingAggregated field to given value.

### HasIsBeingAggregated

`func (o *MemberStatusResponse) HasIsBeingAggregated() bool`

HasIsBeingAggregated returns a boolean if a field has been set.

### SetIsBeingAggregatedNil

`func (o *MemberStatusResponse) SetIsBeingAggregatedNil(b bool)`

 SetIsBeingAggregatedNil sets the value for IsBeingAggregated to be an explicit nil

### UnsetIsBeingAggregated
`func (o *MemberStatusResponse) UnsetIsBeingAggregated()`

UnsetIsBeingAggregated ensures that no value is present for IsBeingAggregated, not even an explicit nil
### GetSuccessfullyAggregatedAt

`func (o *MemberStatusResponse) GetSuccessfullyAggregatedAt() string`

GetSuccessfullyAggregatedAt returns the SuccessfullyAggregatedAt field if non-nil, zero value otherwise.

### GetSuccessfullyAggregatedAtOk

`func (o *MemberStatusResponse) GetSuccessfullyAggregatedAtOk() (*string, bool)`

GetSuccessfullyAggregatedAtOk returns a tuple with the SuccessfullyAggregatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfullyAggregatedAt

`func (o *MemberStatusResponse) SetSuccessfullyAggregatedAt(v string)`

SetSuccessfullyAggregatedAt sets SuccessfullyAggregatedAt field to given value.

### HasSuccessfullyAggregatedAt

`func (o *MemberStatusResponse) HasSuccessfullyAggregatedAt() bool`

HasSuccessfullyAggregatedAt returns a boolean if a field has been set.

### SetSuccessfullyAggregatedAtNil

`func (o *MemberStatusResponse) SetSuccessfullyAggregatedAtNil(b bool)`

 SetSuccessfullyAggregatedAtNil sets the value for SuccessfullyAggregatedAt to be an explicit nil

### UnsetSuccessfullyAggregatedAt
`func (o *MemberStatusResponse) UnsetSuccessfullyAggregatedAt()`

UnsetSuccessfullyAggregatedAt ensures that no value is present for SuccessfullyAggregatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


