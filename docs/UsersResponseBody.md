# UsersResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]UserResponse**](UserResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewUsersResponseBody

`func NewUsersResponseBody() *UsersResponseBody`

NewUsersResponseBody instantiates a new UsersResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersResponseBodyWithDefaults

`func NewUsersResponseBodyWithDefaults() *UsersResponseBody`

NewUsersResponseBodyWithDefaults instantiates a new UsersResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UsersResponseBody) GetUsers() []UserResponse`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersResponseBody) GetUsersOk() (*[]UserResponse, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersResponseBody) SetUsers(v []UserResponse)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UsersResponseBody) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetPagination

`func (o *UsersResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *UsersResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *UsersResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *UsersResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


