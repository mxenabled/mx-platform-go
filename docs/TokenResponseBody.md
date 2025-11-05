# TokenResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tokens** | Pointer to [**[]TokenResponse**](TokenResponse.md) |  | [optional] 

## Methods

### NewTokenResponseBody

`func NewTokenResponseBody() *TokenResponseBody`

NewTokenResponseBody instantiates a new TokenResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseBodyWithDefaults

`func NewTokenResponseBodyWithDefaults() *TokenResponseBody`

NewTokenResponseBodyWithDefaults instantiates a new TokenResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokens

`func (o *TokenResponseBody) GetTokens() []TokenResponse`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TokenResponseBody) GetTokensOk() (*[]TokenResponse, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TokenResponseBody) SetTokens(v []TokenResponse)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *TokenResponseBody) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


