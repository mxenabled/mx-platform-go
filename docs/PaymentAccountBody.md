# PaymentAccountBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentAccount** | Pointer to [**PaymentAccountBodyPaymentAccount**](PaymentAccountBodyPaymentAccount.md) |  | [optional] 

## Methods

### NewPaymentAccountBody

`func NewPaymentAccountBody() *PaymentAccountBody`

NewPaymentAccountBody instantiates a new PaymentAccountBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAccountBodyWithDefaults

`func NewPaymentAccountBodyWithDefaults() *PaymentAccountBody`

NewPaymentAccountBodyWithDefaults instantiates a new PaymentAccountBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentAccount

`func (o *PaymentAccountBody) GetPaymentAccount() PaymentAccountBodyPaymentAccount`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *PaymentAccountBody) GetPaymentAccountOk() (*PaymentAccountBodyPaymentAccount, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *PaymentAccountBody) SetPaymentAccount(v PaymentAccountBodyPaymentAccount)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *PaymentAccountBody) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


