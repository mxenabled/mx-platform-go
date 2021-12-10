# MemberUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundAggregationIsDisabled** | Pointer to **bool** |  | [optional] 
**Credentials** | Pointer to [**[]CredentialRequest**](CredentialRequest.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**SkipAggregation** | Pointer to **bool** |  | [optional] 

## Methods

### NewMemberUpdateRequest

`func NewMemberUpdateRequest() *MemberUpdateRequest`

NewMemberUpdateRequest instantiates a new MemberUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberUpdateRequestWithDefaults

`func NewMemberUpdateRequestWithDefaults() *MemberUpdateRequest`

NewMemberUpdateRequestWithDefaults instantiates a new MemberUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundAggregationIsDisabled

`func (o *MemberUpdateRequest) GetBackgroundAggregationIsDisabled() bool`

GetBackgroundAggregationIsDisabled returns the BackgroundAggregationIsDisabled field if non-nil, zero value otherwise.

### GetBackgroundAggregationIsDisabledOk

`func (o *MemberUpdateRequest) GetBackgroundAggregationIsDisabledOk() (*bool, bool)`

GetBackgroundAggregationIsDisabledOk returns a tuple with the BackgroundAggregationIsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundAggregationIsDisabled

`func (o *MemberUpdateRequest) SetBackgroundAggregationIsDisabled(v bool)`

SetBackgroundAggregationIsDisabled sets BackgroundAggregationIsDisabled field to given value.

### HasBackgroundAggregationIsDisabled

`func (o *MemberUpdateRequest) HasBackgroundAggregationIsDisabled() bool`

HasBackgroundAggregationIsDisabled returns a boolean if a field has been set.

### GetCredentials

`func (o *MemberUpdateRequest) GetCredentials() []CredentialRequest`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *MemberUpdateRequest) GetCredentialsOk() (*[]CredentialRequest, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *MemberUpdateRequest) SetCredentials(v []CredentialRequest)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *MemberUpdateRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetId

`func (o *MemberUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberUpdateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MemberUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *MemberUpdateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MemberUpdateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MemberUpdateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MemberUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSkipAggregation

`func (o *MemberUpdateRequest) GetSkipAggregation() bool`

GetSkipAggregation returns the SkipAggregation field if non-nil, zero value otherwise.

### GetSkipAggregationOk

`func (o *MemberUpdateRequest) GetSkipAggregationOk() (*bool, bool)`

GetSkipAggregationOk returns a tuple with the SkipAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipAggregation

`func (o *MemberUpdateRequest) SetSkipAggregation(v bool)`

SetSkipAggregation sets SkipAggregation field to given value.

### HasSkipAggregation

`func (o *MemberUpdateRequest) HasSkipAggregation() bool`

HasSkipAggregation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


