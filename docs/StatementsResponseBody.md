# StatementsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 
**Statements** | Pointer to [**[]StatementResponse**](StatementResponse.md) |  | [optional] 

## Methods

### NewStatementsResponseBody

`func NewStatementsResponseBody() *StatementsResponseBody`

NewStatementsResponseBody instantiates a new StatementsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementsResponseBodyWithDefaults

`func NewStatementsResponseBodyWithDefaults() *StatementsResponseBody`

NewStatementsResponseBodyWithDefaults instantiates a new StatementsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *StatementsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *StatementsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *StatementsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *StatementsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetStatements

`func (o *StatementsResponseBody) GetStatements() []StatementResponse`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *StatementsResponseBody) GetStatementsOk() (*[]StatementResponse, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *StatementsResponseBody) SetStatements(v []StatementResponse)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *StatementsResponseBody) HasStatements() bool`

HasStatements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


