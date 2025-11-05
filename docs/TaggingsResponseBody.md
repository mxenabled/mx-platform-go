# TaggingsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taggings** | Pointer to [**[]TaggingResponse**](TaggingResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewTaggingsResponseBody

`func NewTaggingsResponseBody() *TaggingsResponseBody`

NewTaggingsResponseBody instantiates a new TaggingsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaggingsResponseBodyWithDefaults

`func NewTaggingsResponseBodyWithDefaults() *TaggingsResponseBody`

NewTaggingsResponseBodyWithDefaults instantiates a new TaggingsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaggings

`func (o *TaggingsResponseBody) GetTaggings() []TaggingResponse`

GetTaggings returns the Taggings field if non-nil, zero value otherwise.

### GetTaggingsOk

`func (o *TaggingsResponseBody) GetTaggingsOk() (*[]TaggingResponse, bool)`

GetTaggingsOk returns a tuple with the Taggings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggings

`func (o *TaggingsResponseBody) SetTaggings(v []TaggingResponse)`

SetTaggings sets Taggings field to given value.

### HasTaggings

`func (o *TaggingsResponseBody) HasTaggings() bool`

HasTaggings returns a boolean if a field has been set.

### GetPagination

`func (o *TaggingsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TaggingsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TaggingsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TaggingsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


