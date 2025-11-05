# PaymentProcessorAuthorizationCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | **string** |  | 
**MemberGuid** | **string** |  | 
**UserGuid** | **string** |  | 

## Methods

### NewPaymentProcessorAuthorizationCodeRequest

`func NewPaymentProcessorAuthorizationCodeRequest(accountGuid string, memberGuid string, userGuid string, ) *PaymentProcessorAuthorizationCodeRequest`

NewPaymentProcessorAuthorizationCodeRequest instantiates a new PaymentProcessorAuthorizationCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentProcessorAuthorizationCodeRequestWithDefaults

`func NewPaymentProcessorAuthorizationCodeRequestWithDefaults() *PaymentProcessorAuthorizationCodeRequest`

NewPaymentProcessorAuthorizationCodeRequestWithDefaults instantiates a new PaymentProcessorAuthorizationCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *PaymentProcessorAuthorizationCodeRequest) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *PaymentProcessorAuthorizationCodeRequest) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *PaymentProcessorAuthorizationCodeRequest) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.


### GetMemberGuid

`func (o *PaymentProcessorAuthorizationCodeRequest) GetMemberGuid() string`

GetMemberGuid returns the MemberGuid field if non-nil, zero value otherwise.

### GetMemberGuidOk

`func (o *PaymentProcessorAuthorizationCodeRequest) GetMemberGuidOk() (*string, bool)`

GetMemberGuidOk returns a tuple with the MemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGuid

`func (o *PaymentProcessorAuthorizationCodeRequest) SetMemberGuid(v string)`

SetMemberGuid sets MemberGuid field to given value.


### GetUserGuid

`func (o *PaymentProcessorAuthorizationCodeRequest) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *PaymentProcessorAuthorizationCodeRequest) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *PaymentProcessorAuthorizationCodeRequest) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


