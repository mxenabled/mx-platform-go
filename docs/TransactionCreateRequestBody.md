# TransactionCreateRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | Pointer to [**TransactionCreateRequest**](TransactionCreateRequest.md) |  | [optional] 

## Methods

### NewTransactionCreateRequestBody

`func NewTransactionCreateRequestBody() *TransactionCreateRequestBody`

NewTransactionCreateRequestBody instantiates a new TransactionCreateRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCreateRequestBodyWithDefaults

`func NewTransactionCreateRequestBodyWithDefaults() *TransactionCreateRequestBody`

NewTransactionCreateRequestBodyWithDefaults instantiates a new TransactionCreateRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *TransactionCreateRequestBody) GetTransaction() TransactionCreateRequest`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *TransactionCreateRequestBody) GetTransactionOk() (*TransactionCreateRequest, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *TransactionCreateRequestBody) SetTransaction(v TransactionCreateRequest)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *TransactionCreateRequestBody) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


