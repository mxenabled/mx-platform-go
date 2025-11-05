# TransactionRulesResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionRules** | Pointer to [**[]TransactionRuleResponse**](TransactionRuleResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewTransactionRulesResponseBody

`func NewTransactionRulesResponseBody() *TransactionRulesResponseBody`

NewTransactionRulesResponseBody instantiates a new TransactionRulesResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRulesResponseBodyWithDefaults

`func NewTransactionRulesResponseBodyWithDefaults() *TransactionRulesResponseBody`

NewTransactionRulesResponseBodyWithDefaults instantiates a new TransactionRulesResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionRules

`func (o *TransactionRulesResponseBody) GetTransactionRules() []TransactionRuleResponse`

GetTransactionRules returns the TransactionRules field if non-nil, zero value otherwise.

### GetTransactionRulesOk

`func (o *TransactionRulesResponseBody) GetTransactionRulesOk() (*[]TransactionRuleResponse, bool)`

GetTransactionRulesOk returns a tuple with the TransactionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRules

`func (o *TransactionRulesResponseBody) SetTransactionRules(v []TransactionRuleResponse)`

SetTransactionRules sets TransactionRules field to given value.

### HasTransactionRules

`func (o *TransactionRulesResponseBody) HasTransactionRules() bool`

HasTransactionRules returns a boolean if a field has been set.

### GetPagination

`func (o *TransactionRulesResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TransactionRulesResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TransactionRulesResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TransactionRulesResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


