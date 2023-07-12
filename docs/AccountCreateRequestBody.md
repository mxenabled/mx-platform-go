# AccountCreateRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipWebhook** | Pointer to **NullableBool** |  | [optional] 
**Account** | Pointer to [**AccountCreateRequest**](AccountCreateRequest.md) |  | [optional] 

## Methods

### NewAccountCreateRequestBody

`func NewAccountCreateRequestBody() *AccountCreateRequestBody`

NewAccountCreateRequestBody instantiates a new AccountCreateRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreateRequestBodyWithDefaults

`func NewAccountCreateRequestBodyWithDefaults() *AccountCreateRequestBody`

NewAccountCreateRequestBodyWithDefaults instantiates a new AccountCreateRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipWebhook

`func (o *AccountCreateRequestBody) GetSkipWebhook() bool`

GetSkipWebhook returns the SkipWebhook field if non-nil, zero value otherwise.

### GetSkipWebhookOk

`func (o *AccountCreateRequestBody) GetSkipWebhookOk() (*bool, bool)`

GetSkipWebhookOk returns a tuple with the SkipWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhook

`func (o *AccountCreateRequestBody) SetSkipWebhook(v bool)`

SetSkipWebhook sets SkipWebhook field to given value.

### HasSkipWebhook

`func (o *AccountCreateRequestBody) HasSkipWebhook() bool`

HasSkipWebhook returns a boolean if a field has been set.

### SetSkipWebhookNil

`func (o *AccountCreateRequestBody) SetSkipWebhookNil(b bool)`

 SetSkipWebhookNil sets the value for SkipWebhook to be an explicit nil

### UnsetSkipWebhook
`func (o *AccountCreateRequestBody) UnsetSkipWebhook()`

UnsetSkipWebhook ensures that no value is present for SkipWebhook, not even an explicit nil
### GetAccount

`func (o *AccountCreateRequestBody) GetAccount() AccountCreateRequest`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountCreateRequestBody) GetAccountOk() (*AccountCreateRequest, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountCreateRequestBody) SetAccount(v AccountCreateRequest)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AccountCreateRequestBody) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


