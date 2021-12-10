# ManagedMemberCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**InstitutionCode** | **string** |  | 
**Metadata** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedMemberCreateRequest

`func NewManagedMemberCreateRequest(institutionCode string, ) *ManagedMemberCreateRequest`

NewManagedMemberCreateRequest instantiates a new ManagedMemberCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedMemberCreateRequestWithDefaults

`func NewManagedMemberCreateRequestWithDefaults() *ManagedMemberCreateRequest`

NewManagedMemberCreateRequestWithDefaults instantiates a new ManagedMemberCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedMemberCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedMemberCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedMemberCreateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedMemberCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstitutionCode

`func (o *ManagedMemberCreateRequest) GetInstitutionCode() string`

GetInstitutionCode returns the InstitutionCode field if non-nil, zero value otherwise.

### GetInstitutionCodeOk

`func (o *ManagedMemberCreateRequest) GetInstitutionCodeOk() (*string, bool)`

GetInstitutionCodeOk returns a tuple with the InstitutionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionCode

`func (o *ManagedMemberCreateRequest) SetInstitutionCode(v string)`

SetInstitutionCode sets InstitutionCode field to given value.


### GetMetadata

`func (o *ManagedMemberCreateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ManagedMemberCreateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ManagedMemberCreateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ManagedMemberCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ManagedMemberCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedMemberCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedMemberCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagedMemberCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


