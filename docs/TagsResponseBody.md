# TagsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]TagResponse**](TagResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewTagsResponseBody

`func NewTagsResponseBody() *TagsResponseBody`

NewTagsResponseBody instantiates a new TagsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsResponseBodyWithDefaults

`func NewTagsResponseBodyWithDefaults() *TagsResponseBody`

NewTagsResponseBodyWithDefaults instantiates a new TagsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *TagsResponseBody) GetTags() []TagResponse`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagsResponseBody) GetTagsOk() (*[]TagResponse, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagsResponseBody) SetTags(v []TagResponse)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TagsResponseBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPagination

`func (o *TagsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TagsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TagsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TagsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


