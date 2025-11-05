# TransactionsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**[]TransactionResponse**](TransactionResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewTransactionsResponseBody

`func NewTransactionsResponseBody() *TransactionsResponseBody`

NewTransactionsResponseBody instantiates a new TransactionsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsResponseBodyWithDefaults

`func NewTransactionsResponseBodyWithDefaults() *TransactionsResponseBody`

NewTransactionsResponseBodyWithDefaults instantiates a new TransactionsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *TransactionsResponseBody) GetTransactions() []TransactionResponse`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *TransactionsResponseBody) GetTransactionsOk() (*[]TransactionResponse, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *TransactionsResponseBody) SetTransactions(v []TransactionResponse)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *TransactionsResponseBody) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetPagination

`func (o *TransactionsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TransactionsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TransactionsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TransactionsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


