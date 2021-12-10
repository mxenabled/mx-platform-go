# AccountUpdateRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AccountUpdateRequest**](AccountUpdateRequest.md) |  | [optional] 

## Methods

### NewAccountUpdateRequestBody

`func NewAccountUpdateRequestBody() *AccountUpdateRequestBody`

NewAccountUpdateRequestBody instantiates a new AccountUpdateRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateRequestBodyWithDefaults

`func NewAccountUpdateRequestBodyWithDefaults() *AccountUpdateRequestBody`

NewAccountUpdateRequestBodyWithDefaults instantiates a new AccountUpdateRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AccountUpdateRequestBody) GetAccount() AccountUpdateRequest`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountUpdateRequestBody) GetAccountOk() (*AccountUpdateRequest, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountUpdateRequestBody) SetAccount(v AccountUpdateRequest)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AccountUpdateRequestBody) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


