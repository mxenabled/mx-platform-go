# MemberCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundAggregationIsDisabled** | Pointer to **bool** |  | [optional] 
**Credentials** | [**[]CredentialRequest**](CredentialRequest.md) |  | 
**Id** | Pointer to **string** |  | [optional] 
**InstitutionCode** | **string** |  | 
**IsOauth** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**SkipAggregation** | Pointer to **bool** |  | [optional] 

## Methods

### NewMemberCreateRequest

`func NewMemberCreateRequest(credentials []CredentialRequest, institutionCode string, ) *MemberCreateRequest`

NewMemberCreateRequest instantiates a new MemberCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberCreateRequestWithDefaults

`func NewMemberCreateRequestWithDefaults() *MemberCreateRequest`

NewMemberCreateRequestWithDefaults instantiates a new MemberCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundAggregationIsDisabled

`func (o *MemberCreateRequest) GetBackgroundAggregationIsDisabled() bool`

GetBackgroundAggregationIsDisabled returns the BackgroundAggregationIsDisabled field if non-nil, zero value otherwise.

### GetBackgroundAggregationIsDisabledOk

`func (o *MemberCreateRequest) GetBackgroundAggregationIsDisabledOk() (*bool, bool)`

GetBackgroundAggregationIsDisabledOk returns a tuple with the BackgroundAggregationIsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundAggregationIsDisabled

`func (o *MemberCreateRequest) SetBackgroundAggregationIsDisabled(v bool)`

SetBackgroundAggregationIsDisabled sets BackgroundAggregationIsDisabled field to given value.

### HasBackgroundAggregationIsDisabled

`func (o *MemberCreateRequest) HasBackgroundAggregationIsDisabled() bool`

HasBackgroundAggregationIsDisabled returns a boolean if a field has been set.

### GetCredentials

`func (o *MemberCreateRequest) GetCredentials() []CredentialRequest`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *MemberCreateRequest) GetCredentialsOk() (*[]CredentialRequest, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *MemberCreateRequest) SetCredentials(v []CredentialRequest)`

SetCredentials sets Credentials field to given value.


### GetId

`func (o *MemberCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberCreateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MemberCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstitutionCode

`func (o *MemberCreateRequest) GetInstitutionCode() string`

GetInstitutionCode returns the InstitutionCode field if non-nil, zero value otherwise.

### GetInstitutionCodeOk

`func (o *MemberCreateRequest) GetInstitutionCodeOk() (*string, bool)`

GetInstitutionCodeOk returns a tuple with the InstitutionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionCode

`func (o *MemberCreateRequest) SetInstitutionCode(v string)`

SetInstitutionCode sets InstitutionCode field to given value.


### GetIsOauth

`func (o *MemberCreateRequest) GetIsOauth() bool`

GetIsOauth returns the IsOauth field if non-nil, zero value otherwise.

### GetIsOauthOk

`func (o *MemberCreateRequest) GetIsOauthOk() (*bool, bool)`

GetIsOauthOk returns a tuple with the IsOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOauth

`func (o *MemberCreateRequest) SetIsOauth(v bool)`

SetIsOauth sets IsOauth field to given value.

### HasIsOauth

`func (o *MemberCreateRequest) HasIsOauth() bool`

HasIsOauth returns a boolean if a field has been set.

### GetMetadata

`func (o *MemberCreateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MemberCreateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MemberCreateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MemberCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSkipAggregation

`func (o *MemberCreateRequest) GetSkipAggregation() bool`

GetSkipAggregation returns the SkipAggregation field if non-nil, zero value otherwise.

### GetSkipAggregationOk

`func (o *MemberCreateRequest) GetSkipAggregationOk() (*bool, bool)`

GetSkipAggregationOk returns a tuple with the SkipAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipAggregation

`func (o *MemberCreateRequest) SetSkipAggregation(v bool)`

SetSkipAggregation sets SkipAggregation field to given value.

### HasSkipAggregation

`func (o *MemberCreateRequest) HasSkipAggregation() bool`

HasSkipAggregation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


