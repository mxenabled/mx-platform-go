# AccountNumbersResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumbers** | Pointer to [**[]AccountNumberResponse**](AccountNumberResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewAccountNumbersResponseBody

`func NewAccountNumbersResponseBody() *AccountNumbersResponseBody`

NewAccountNumbersResponseBody instantiates a new AccountNumbersResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountNumbersResponseBodyWithDefaults

`func NewAccountNumbersResponseBodyWithDefaults() *AccountNumbersResponseBody`

NewAccountNumbersResponseBodyWithDefaults instantiates a new AccountNumbersResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumbers

`func (o *AccountNumbersResponseBody) GetAccountNumbers() []AccountNumberResponse`

GetAccountNumbers returns the AccountNumbers field if non-nil, zero value otherwise.

### GetAccountNumbersOk

`func (o *AccountNumbersResponseBody) GetAccountNumbersOk() (*[]AccountNumberResponse, bool)`

GetAccountNumbersOk returns a tuple with the AccountNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumbers

`func (o *AccountNumbersResponseBody) SetAccountNumbers(v []AccountNumberResponse)`

SetAccountNumbers sets AccountNumbers field to given value.

### HasAccountNumbers

`func (o *AccountNumbersResponseBody) HasAccountNumbers() bool`

HasAccountNumbers returns a boolean if a field has been set.

### GetPagination

`func (o *AccountNumbersResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AccountNumbersResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AccountNumbersResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AccountNumbersResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


