# TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentProcessorGuid** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**AccessToken** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewTokenResponse

`func NewTokenResponse() *TokenResponse`

NewTokenResponse instantiates a new TokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseWithDefaults

`func NewTokenResponseWithDefaults() *TokenResponse`

NewTokenResponseWithDefaults instantiates a new TokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentProcessorGuid

`func (o *TokenResponse) GetPaymentProcessorGuid() string`

GetPaymentProcessorGuid returns the PaymentProcessorGuid field if non-nil, zero value otherwise.

### GetPaymentProcessorGuidOk

`func (o *TokenResponse) GetPaymentProcessorGuidOk() (*string, bool)`

GetPaymentProcessorGuidOk returns a tuple with the PaymentProcessorGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProcessorGuid

`func (o *TokenResponse) SetPaymentProcessorGuid(v string)`

SetPaymentProcessorGuid sets PaymentProcessorGuid field to given value.

### HasPaymentProcessorGuid

`func (o *TokenResponse) HasPaymentProcessorGuid() bool`

HasPaymentProcessorGuid returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TokenResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetAccessToken

`func (o *TokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TokenResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetActive

`func (o *TokenResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TokenResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TokenResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TokenResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


