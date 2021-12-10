# AccountOwnersResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountOwners** | Pointer to [**[]AccountOwnerResponse**](AccountOwnerResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewAccountOwnersResponseBody

`func NewAccountOwnersResponseBody() *AccountOwnersResponseBody`

NewAccountOwnersResponseBody instantiates a new AccountOwnersResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOwnersResponseBodyWithDefaults

`func NewAccountOwnersResponseBodyWithDefaults() *AccountOwnersResponseBody`

NewAccountOwnersResponseBodyWithDefaults instantiates a new AccountOwnersResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountOwners

`func (o *AccountOwnersResponseBody) GetAccountOwners() []AccountOwnerResponse`

GetAccountOwners returns the AccountOwners field if non-nil, zero value otherwise.

### GetAccountOwnersOk

`func (o *AccountOwnersResponseBody) GetAccountOwnersOk() (*[]AccountOwnerResponse, bool)`

GetAccountOwnersOk returns a tuple with the AccountOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwners

`func (o *AccountOwnersResponseBody) SetAccountOwners(v []AccountOwnerResponse)`

SetAccountOwners sets AccountOwners field to given value.

### HasAccountOwners

`func (o *AccountOwnersResponseBody) HasAccountOwners() bool`

HasAccountOwners returns a boolean if a field has been set.

### GetPagination

`func (o *AccountOwnersResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AccountOwnersResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AccountOwnersResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AccountOwnersResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


