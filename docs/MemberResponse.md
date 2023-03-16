# MemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregatedAt** | Pointer to **NullableString** |  | [optional] 
**BackgroundAggregationIsDisabled** | Pointer to **bool** |  | [optional] 
**ConnectionStatus** | Pointer to **NullableString** |  | [optional] 
**Guid** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**InstitutionCode** | Pointer to **NullableString** |  | [optional] 
**IsBeingAggregated** | Pointer to **NullableBool** |  | [optional] 
**IsManagedByUser** | Pointer to **NullableBool** |  | [optional] 
**IsOauth** | Pointer to **NullableBool** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OauthWindowUri** | Pointer to **NullableString** |  | [optional] 
**SuccessfullyAggregatedAt** | Pointer to **NullableString** |  | [optional] 
**UserGuid** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMemberResponse

`func NewMemberResponse() *MemberResponse`

NewMemberResponse instantiates a new MemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberResponseWithDefaults

`func NewMemberResponseWithDefaults() *MemberResponse`

NewMemberResponseWithDefaults instantiates a new MemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatedAt

`func (o *MemberResponse) GetAggregatedAt() string`

GetAggregatedAt returns the AggregatedAt field if non-nil, zero value otherwise.

### GetAggregatedAtOk

`func (o *MemberResponse) GetAggregatedAtOk() (*string, bool)`

GetAggregatedAtOk returns a tuple with the AggregatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedAt

`func (o *MemberResponse) SetAggregatedAt(v string)`

SetAggregatedAt sets AggregatedAt field to given value.

### HasAggregatedAt

`func (o *MemberResponse) HasAggregatedAt() bool`

HasAggregatedAt returns a boolean if a field has been set.

### SetAggregatedAtNil

`func (o *MemberResponse) SetAggregatedAtNil(b bool)`

 SetAggregatedAtNil sets the value for AggregatedAt to be an explicit nil

### UnsetAggregatedAt
`func (o *MemberResponse) UnsetAggregatedAt()`

UnsetAggregatedAt ensures that no value is present for AggregatedAt, not even an explicit nil
### GetBackgroundAggregationIsDisabled

`func (o *MemberResponse) GetBackgroundAggregationIsDisabled() bool`

GetBackgroundAggregationIsDisabled returns the BackgroundAggregationIsDisabled field if non-nil, zero value otherwise.

### GetBackgroundAggregationIsDisabledOk

`func (o *MemberResponse) GetBackgroundAggregationIsDisabledOk() (*bool, bool)`

GetBackgroundAggregationIsDisabledOk returns a tuple with the BackgroundAggregationIsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundAggregationIsDisabled

`func (o *MemberResponse) SetBackgroundAggregationIsDisabled(v bool)`

SetBackgroundAggregationIsDisabled sets BackgroundAggregationIsDisabled field to given value.

### HasBackgroundAggregationIsDisabled

`func (o *MemberResponse) HasBackgroundAggregationIsDisabled() bool`

HasBackgroundAggregationIsDisabled returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *MemberResponse) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *MemberResponse) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *MemberResponse) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *MemberResponse) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### SetConnectionStatusNil

`func (o *MemberResponse) SetConnectionStatusNil(b bool)`

 SetConnectionStatusNil sets the value for ConnectionStatus to be an explicit nil

### UnsetConnectionStatus
`func (o *MemberResponse) UnsetConnectionStatus()`

UnsetConnectionStatus ensures that no value is present for ConnectionStatus, not even an explicit nil
### GetGuid

`func (o *MemberResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *MemberResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *MemberResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *MemberResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *MemberResponse) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *MemberResponse) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetId

`func (o *MemberResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MemberResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MemberResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MemberResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInstitutionCode

`func (o *MemberResponse) GetInstitutionCode() string`

GetInstitutionCode returns the InstitutionCode field if non-nil, zero value otherwise.

### GetInstitutionCodeOk

`func (o *MemberResponse) GetInstitutionCodeOk() (*string, bool)`

GetInstitutionCodeOk returns a tuple with the InstitutionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionCode

`func (o *MemberResponse) SetInstitutionCode(v string)`

SetInstitutionCode sets InstitutionCode field to given value.

### HasInstitutionCode

`func (o *MemberResponse) HasInstitutionCode() bool`

HasInstitutionCode returns a boolean if a field has been set.

### SetInstitutionCodeNil

`func (o *MemberResponse) SetInstitutionCodeNil(b bool)`

 SetInstitutionCodeNil sets the value for InstitutionCode to be an explicit nil

### UnsetInstitutionCode
`func (o *MemberResponse) UnsetInstitutionCode()`

UnsetInstitutionCode ensures that no value is present for InstitutionCode, not even an explicit nil
### GetIsBeingAggregated

`func (o *MemberResponse) GetIsBeingAggregated() bool`

GetIsBeingAggregated returns the IsBeingAggregated field if non-nil, zero value otherwise.

### GetIsBeingAggregatedOk

`func (o *MemberResponse) GetIsBeingAggregatedOk() (*bool, bool)`

GetIsBeingAggregatedOk returns a tuple with the IsBeingAggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBeingAggregated

`func (o *MemberResponse) SetIsBeingAggregated(v bool)`

SetIsBeingAggregated sets IsBeingAggregated field to given value.

### HasIsBeingAggregated

`func (o *MemberResponse) HasIsBeingAggregated() bool`

HasIsBeingAggregated returns a boolean if a field has been set.

### SetIsBeingAggregatedNil

`func (o *MemberResponse) SetIsBeingAggregatedNil(b bool)`

 SetIsBeingAggregatedNil sets the value for IsBeingAggregated to be an explicit nil

### UnsetIsBeingAggregated
`func (o *MemberResponse) UnsetIsBeingAggregated()`

UnsetIsBeingAggregated ensures that no value is present for IsBeingAggregated, not even an explicit nil
### GetIsManagedByUser

`func (o *MemberResponse) GetIsManagedByUser() bool`

GetIsManagedByUser returns the IsManagedByUser field if non-nil, zero value otherwise.

### GetIsManagedByUserOk

`func (o *MemberResponse) GetIsManagedByUserOk() (*bool, bool)`

GetIsManagedByUserOk returns a tuple with the IsManagedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManagedByUser

`func (o *MemberResponse) SetIsManagedByUser(v bool)`

SetIsManagedByUser sets IsManagedByUser field to given value.

### HasIsManagedByUser

`func (o *MemberResponse) HasIsManagedByUser() bool`

HasIsManagedByUser returns a boolean if a field has been set.

### SetIsManagedByUserNil

`func (o *MemberResponse) SetIsManagedByUserNil(b bool)`

 SetIsManagedByUserNil sets the value for IsManagedByUser to be an explicit nil

### UnsetIsManagedByUser
`func (o *MemberResponse) UnsetIsManagedByUser()`

UnsetIsManagedByUser ensures that no value is present for IsManagedByUser, not even an explicit nil
### GetIsOauth

`func (o *MemberResponse) GetIsOauth() bool`

GetIsOauth returns the IsOauth field if non-nil, zero value otherwise.

### GetIsOauthOk

`func (o *MemberResponse) GetIsOauthOk() (*bool, bool)`

GetIsOauthOk returns a tuple with the IsOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOauth

`func (o *MemberResponse) SetIsOauth(v bool)`

SetIsOauth sets IsOauth field to given value.

### HasIsOauth

`func (o *MemberResponse) HasIsOauth() bool`

HasIsOauth returns a boolean if a field has been set.

### SetIsOauthNil

`func (o *MemberResponse) SetIsOauthNil(b bool)`

 SetIsOauthNil sets the value for IsOauth to be an explicit nil

### UnsetIsOauth
`func (o *MemberResponse) UnsetIsOauth()`

UnsetIsOauth ensures that no value is present for IsOauth, not even an explicit nil
### GetMetadata

`func (o *MemberResponse) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MemberResponse) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MemberResponse) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MemberResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *MemberResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *MemberResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *MemberResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemberResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MemberResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MemberResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOauthWindowUri

`func (o *MemberResponse) GetOauthWindowUri() string`

GetOauthWindowUri returns the OauthWindowUri field if non-nil, zero value otherwise.

### GetOauthWindowUriOk

`func (o *MemberResponse) GetOauthWindowUriOk() (*string, bool)`

GetOauthWindowUriOk returns a tuple with the OauthWindowUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthWindowUri

`func (o *MemberResponse) SetOauthWindowUri(v string)`

SetOauthWindowUri sets OauthWindowUri field to given value.

### HasOauthWindowUri

`func (o *MemberResponse) HasOauthWindowUri() bool`

HasOauthWindowUri returns a boolean if a field has been set.

### SetOauthWindowUriNil

`func (o *MemberResponse) SetOauthWindowUriNil(b bool)`

 SetOauthWindowUriNil sets the value for OauthWindowUri to be an explicit nil

### UnsetOauthWindowUri
`func (o *MemberResponse) UnsetOauthWindowUri()`

UnsetOauthWindowUri ensures that no value is present for OauthWindowUri, not even an explicit nil
### GetSuccessfullyAggregatedAt

`func (o *MemberResponse) GetSuccessfullyAggregatedAt() string`

GetSuccessfullyAggregatedAt returns the SuccessfullyAggregatedAt field if non-nil, zero value otherwise.

### GetSuccessfullyAggregatedAtOk

`func (o *MemberResponse) GetSuccessfullyAggregatedAtOk() (*string, bool)`

GetSuccessfullyAggregatedAtOk returns a tuple with the SuccessfullyAggregatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfullyAggregatedAt

`func (o *MemberResponse) SetSuccessfullyAggregatedAt(v string)`

SetSuccessfullyAggregatedAt sets SuccessfullyAggregatedAt field to given value.

### HasSuccessfullyAggregatedAt

`func (o *MemberResponse) HasSuccessfullyAggregatedAt() bool`

HasSuccessfullyAggregatedAt returns a boolean if a field has been set.

### SetSuccessfullyAggregatedAtNil

`func (o *MemberResponse) SetSuccessfullyAggregatedAtNil(b bool)`

 SetSuccessfullyAggregatedAtNil sets the value for SuccessfullyAggregatedAt to be an explicit nil

### UnsetSuccessfullyAggregatedAt
`func (o *MemberResponse) UnsetSuccessfullyAggregatedAt()`

UnsetSuccessfullyAggregatedAt ensures that no value is present for SuccessfullyAggregatedAt, not even an explicit nil
### GetUserGuid

`func (o *MemberResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *MemberResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *MemberResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *MemberResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### SetUserGuidNil

`func (o *MemberResponse) SetUserGuidNil(b bool)`

 SetUserGuidNil sets the value for UserGuid to be an explicit nil

### UnsetUserGuid
`func (o *MemberResponse) UnsetUserGuid()`

UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
### GetUserId

`func (o *MemberResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MemberResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MemberResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MemberResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *MemberResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MemberResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


