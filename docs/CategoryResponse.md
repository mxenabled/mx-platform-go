# CategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**IsIncome** | Pointer to **NullableBool** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ParentGuid** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCategoryResponse

`func NewCategoryResponse() *CategoryResponse`

NewCategoryResponse instantiates a new CategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryResponseWithDefaults

`func NewCategoryResponseWithDefaults() *CategoryResponse`

NewCategoryResponseWithDefaults instantiates a new CategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CategoryResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CategoryResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CategoryResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CategoryResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CategoryResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CategoryResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetGuid

`func (o *CategoryResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *CategoryResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *CategoryResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *CategoryResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetIsDefault

`func (o *CategoryResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CategoryResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CategoryResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CategoryResponse) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *CategoryResponse) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *CategoryResponse) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetIsIncome

`func (o *CategoryResponse) GetIsIncome() bool`

GetIsIncome returns the IsIncome field if non-nil, zero value otherwise.

### GetIsIncomeOk

`func (o *CategoryResponse) GetIsIncomeOk() (*bool, bool)`

GetIsIncomeOk returns a tuple with the IsIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncome

`func (o *CategoryResponse) SetIsIncome(v bool)`

SetIsIncome sets IsIncome field to given value.

### HasIsIncome

`func (o *CategoryResponse) HasIsIncome() bool`

HasIsIncome returns a boolean if a field has been set.

### SetIsIncomeNil

`func (o *CategoryResponse) SetIsIncomeNil(b bool)`

 SetIsIncomeNil sets the value for IsIncome to be an explicit nil

### UnsetIsIncome
`func (o *CategoryResponse) UnsetIsIncome()`

UnsetIsIncome ensures that no value is present for IsIncome, not even an explicit nil
### GetMetadata

`func (o *CategoryResponse) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CategoryResponse) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CategoryResponse) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CategoryResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CategoryResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CategoryResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *CategoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CategoryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CategoryResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CategoryResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentGuid

`func (o *CategoryResponse) GetParentGuid() string`

GetParentGuid returns the ParentGuid field if non-nil, zero value otherwise.

### GetParentGuidOk

`func (o *CategoryResponse) GetParentGuidOk() (*string, bool)`

GetParentGuidOk returns a tuple with the ParentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGuid

`func (o *CategoryResponse) SetParentGuid(v string)`

SetParentGuid sets ParentGuid field to given value.

### HasParentGuid

`func (o *CategoryResponse) HasParentGuid() bool`

HasParentGuid returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CategoryResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CategoryResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CategoryResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CategoryResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CategoryResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CategoryResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


