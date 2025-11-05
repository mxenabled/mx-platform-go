# CategoryCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**ParentGuid** | **string** |  | 

## Methods

### NewCategoryCreateRequest

`func NewCategoryCreateRequest(name string, parentGuid string, ) *CategoryCreateRequest`

NewCategoryCreateRequest instantiates a new CategoryCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryCreateRequestWithDefaults

`func NewCategoryCreateRequestWithDefaults() *CategoryCreateRequest`

NewCategoryCreateRequestWithDefaults instantiates a new CategoryCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CategoryCreateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CategoryCreateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CategoryCreateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CategoryCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CategoryCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParentGuid

`func (o *CategoryCreateRequest) GetParentGuid() string`

GetParentGuid returns the ParentGuid field if non-nil, zero value otherwise.

### GetParentGuidOk

`func (o *CategoryCreateRequest) GetParentGuidOk() (*string, bool)`

GetParentGuidOk returns a tuple with the ParentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGuid

`func (o *CategoryCreateRequest) SetParentGuid(v string)`

SetParentGuid sets ParentGuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


