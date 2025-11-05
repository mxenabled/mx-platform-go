# TransactionsResponseBodyIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**[]TransactionIncludesResponse**](TransactionIncludesResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewTransactionsResponseBodyIncludes

`func NewTransactionsResponseBodyIncludes() *TransactionsResponseBodyIncludes`

NewTransactionsResponseBodyIncludes instantiates a new TransactionsResponseBodyIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsResponseBodyIncludesWithDefaults

`func NewTransactionsResponseBodyIncludesWithDefaults() *TransactionsResponseBodyIncludes`

NewTransactionsResponseBodyIncludesWithDefaults instantiates a new TransactionsResponseBodyIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *TransactionsResponseBodyIncludes) GetTransactions() []TransactionIncludesResponse`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *TransactionsResponseBodyIncludes) GetTransactionsOk() (*[]TransactionIncludesResponse, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *TransactionsResponseBodyIncludes) SetTransactions(v []TransactionIncludesResponse)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *TransactionsResponseBodyIncludes) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetPagination

`func (o *TransactionsResponseBodyIncludes) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TransactionsResponseBodyIncludes) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TransactionsResponseBodyIncludes) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TransactionsResponseBodyIncludes) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


