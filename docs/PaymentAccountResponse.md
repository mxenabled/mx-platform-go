# PaymentAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **NullableString** |  | [optional] 
**AccountNumber** | Pointer to **NullableString** |  | [optional] 
**AccountType** | Pointer to **NullableString** |  | [optional] 
**AvailableBalance** | Pointer to **NullableFloat32** |  | [optional] 
**Balance** | Pointer to **NullableFloat32** |  | [optional] 
**RoutingNumber** | Pointer to **NullableString** |  | [optional] 
**TransitNumber** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentAccountResponse

`func NewPaymentAccountResponse() *PaymentAccountResponse`

NewPaymentAccountResponse instantiates a new PaymentAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAccountResponseWithDefaults

`func NewPaymentAccountResponseWithDefaults() *PaymentAccountResponse`

NewPaymentAccountResponseWithDefaults instantiates a new PaymentAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *PaymentAccountResponse) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PaymentAccountResponse) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PaymentAccountResponse) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PaymentAccountResponse) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *PaymentAccountResponse) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *PaymentAccountResponse) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetAccountNumber

`func (o *PaymentAccountResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PaymentAccountResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PaymentAccountResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *PaymentAccountResponse) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *PaymentAccountResponse) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *PaymentAccountResponse) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetAccountType

`func (o *PaymentAccountResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PaymentAccountResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PaymentAccountResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *PaymentAccountResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountTypeNil

`func (o *PaymentAccountResponse) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *PaymentAccountResponse) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetAvailableBalance

`func (o *PaymentAccountResponse) GetAvailableBalance() float32`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *PaymentAccountResponse) GetAvailableBalanceOk() (*float32, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *PaymentAccountResponse) SetAvailableBalance(v float32)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *PaymentAccountResponse) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### SetAvailableBalanceNil

`func (o *PaymentAccountResponse) SetAvailableBalanceNil(b bool)`

 SetAvailableBalanceNil sets the value for AvailableBalance to be an explicit nil

### UnsetAvailableBalance
`func (o *PaymentAccountResponse) UnsetAvailableBalance()`

UnsetAvailableBalance ensures that no value is present for AvailableBalance, not even an explicit nil
### GetBalance

`func (o *PaymentAccountResponse) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *PaymentAccountResponse) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *PaymentAccountResponse) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *PaymentAccountResponse) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *PaymentAccountResponse) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *PaymentAccountResponse) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetRoutingNumber

`func (o *PaymentAccountResponse) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *PaymentAccountResponse) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *PaymentAccountResponse) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *PaymentAccountResponse) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### SetRoutingNumberNil

`func (o *PaymentAccountResponse) SetRoutingNumberNil(b bool)`

 SetRoutingNumberNil sets the value for RoutingNumber to be an explicit nil

### UnsetRoutingNumber
`func (o *PaymentAccountResponse) UnsetRoutingNumber()`

UnsetRoutingNumber ensures that no value is present for RoutingNumber, not even an explicit nil
### GetTransitNumber

`func (o *PaymentAccountResponse) GetTransitNumber() string`

GetTransitNumber returns the TransitNumber field if non-nil, zero value otherwise.

### GetTransitNumberOk

`func (o *PaymentAccountResponse) GetTransitNumberOk() (*string, bool)`

GetTransitNumberOk returns a tuple with the TransitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitNumber

`func (o *PaymentAccountResponse) SetTransitNumber(v string)`

SetTransitNumber sets TransitNumber field to given value.

### HasTransitNumber

`func (o *PaymentAccountResponse) HasTransitNumber() bool`

HasTransitNumber returns a boolean if a field has been set.

### SetTransitNumberNil

`func (o *PaymentAccountResponse) SetTransitNumberNil(b bool)`

 SetTransitNumberNil sets the value for TransitNumber to be an explicit nil

### UnsetTransitNumber
`func (o *PaymentAccountResponse) UnsetTransitNumber()`

UnsetTransitNumber ensures that no value is present for TransitNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


