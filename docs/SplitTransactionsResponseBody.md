# SplitTransactionsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**[]TransactionResponse**](TransactionResponse.md) |  | [optional] 

## Methods

### NewSplitTransactionsResponseBody

`func NewSplitTransactionsResponseBody() *SplitTransactionsResponseBody`

NewSplitTransactionsResponseBody instantiates a new SplitTransactionsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitTransactionsResponseBodyWithDefaults

`func NewSplitTransactionsResponseBodyWithDefaults() *SplitTransactionsResponseBody`

NewSplitTransactionsResponseBodyWithDefaults instantiates a new SplitTransactionsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *SplitTransactionsResponseBody) GetTransactions() []TransactionResponse`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *SplitTransactionsResponseBody) GetTransactionsOk() (*[]TransactionResponse, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *SplitTransactionsResponseBody) SetTransactions(v []TransactionResponse)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *SplitTransactionsResponseBody) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


