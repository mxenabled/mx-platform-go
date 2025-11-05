# ACHReturnCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | **string** | The unique identifier for the account associated with the transaction. Defined by MX. | 
**AccountNumberLastFour** | Pointer to **string** | The last 4 digits of the account number used for the transaction by the Originating Depository Financial Institution (ODFI). | [optional] 
**AchInitiatedAt** | Pointer to **string** | The date and time when the transaction was initiated by the Originating Depository Financial Institution (ODFI) in ISO 8601 format without timestamp. | [optional] 
**CorrectedAccountNumber** | Pointer to **string** | The account number correction reported by the RDFI. Populate only if the &#x60;resolution_code&#x60; is &#x60;NOTICE_OF_CHANGE&#x60;. | [optional] 
**CorrectedRoutingNumber** | Pointer to **string** | The routing number correction reported by the RDFI. Populate only if the &#x60;resolution_code&#x60; is &#x60;NOTICE_OF_CHANGE&#x60;. Must be a valid 9-digit routing number format. | [optional] 
**Id** | **string** | Client-defined identifier for this specific return submission. Allows you to track and reference you requests. | 
**MemberGuid** | **string** | The unique identifier for the member associated with the transaction. Defined by MX. | 
**ReturnAccountNumber** | Pointer to **string** | Incorrect account number used in the ACH transaction. | [optional] 
**ReturnCode** | **string** | The associated ACH return code and notice of change code (for example, R02, R03, R04, R05, R20, NOC). See [Return Codes](/api-reference/platform-api/reference/ach-return-fields#return-codes) for a complete list. | 
**ReturnNotes** | Pointer to **string** | Notes that you set to inform MX on internal ACH processing. | [optional] 
**ReturnRoutingNumber** | Pointer to **string** | Incorrect routing number used in the ACH transaction. | [optional] 
**ReturnedAt** | Pointer to **string** | The date and time when the return was reported by the Receiving Financial Depository Institution (RDFI) in ISO 8601 format without timestamp. | [optional] 
**SecCode** | Pointer to **string** | The SEC code (Standard Entry Class Code)–a three-letter code describing how a payment was authorized (for example, &#x60;WEB&#x60;). See [SEC Codes](/api-reference/platform-api/reference/ach-return-fields#sec-codes) for a complete list. | [optional] 
**TransactionAmount** | Pointer to **float32** | The amount of the transaction. | [optional] 
**TransactionAmountRange** | Pointer to **float32** | The transaction amount range, used for impact assessment. | [optional] 
**UserGuid** | **string** | MX-defined identifier for the user associated with the ACH return. | 

## Methods

### NewACHReturnCreateRequest

`func NewACHReturnCreateRequest(accountGuid string, id string, memberGuid string, returnCode string, userGuid string, ) *ACHReturnCreateRequest`

NewACHReturnCreateRequest instantiates a new ACHReturnCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACHReturnCreateRequestWithDefaults

`func NewACHReturnCreateRequestWithDefaults() *ACHReturnCreateRequest`

NewACHReturnCreateRequestWithDefaults instantiates a new ACHReturnCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *ACHReturnCreateRequest) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *ACHReturnCreateRequest) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *ACHReturnCreateRequest) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.


### GetAccountNumberLastFour

`func (o *ACHReturnCreateRequest) GetAccountNumberLastFour() string`

GetAccountNumberLastFour returns the AccountNumberLastFour field if non-nil, zero value otherwise.

### GetAccountNumberLastFourOk

`func (o *ACHReturnCreateRequest) GetAccountNumberLastFourOk() (*string, bool)`

GetAccountNumberLastFourOk returns a tuple with the AccountNumberLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberLastFour

`func (o *ACHReturnCreateRequest) SetAccountNumberLastFour(v string)`

SetAccountNumberLastFour sets AccountNumberLastFour field to given value.

### HasAccountNumberLastFour

`func (o *ACHReturnCreateRequest) HasAccountNumberLastFour() bool`

HasAccountNumberLastFour returns a boolean if a field has been set.

### GetAchInitiatedAt

`func (o *ACHReturnCreateRequest) GetAchInitiatedAt() string`

GetAchInitiatedAt returns the AchInitiatedAt field if non-nil, zero value otherwise.

### GetAchInitiatedAtOk

`func (o *ACHReturnCreateRequest) GetAchInitiatedAtOk() (*string, bool)`

GetAchInitiatedAtOk returns a tuple with the AchInitiatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchInitiatedAt

`func (o *ACHReturnCreateRequest) SetAchInitiatedAt(v string)`

SetAchInitiatedAt sets AchInitiatedAt field to given value.

### HasAchInitiatedAt

`func (o *ACHReturnCreateRequest) HasAchInitiatedAt() bool`

HasAchInitiatedAt returns a boolean if a field has been set.

### GetCorrectedAccountNumber

`func (o *ACHReturnCreateRequest) GetCorrectedAccountNumber() string`

GetCorrectedAccountNumber returns the CorrectedAccountNumber field if non-nil, zero value otherwise.

### GetCorrectedAccountNumberOk

`func (o *ACHReturnCreateRequest) GetCorrectedAccountNumberOk() (*string, bool)`

GetCorrectedAccountNumberOk returns a tuple with the CorrectedAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectedAccountNumber

`func (o *ACHReturnCreateRequest) SetCorrectedAccountNumber(v string)`

SetCorrectedAccountNumber sets CorrectedAccountNumber field to given value.

### HasCorrectedAccountNumber

`func (o *ACHReturnCreateRequest) HasCorrectedAccountNumber() bool`

HasCorrectedAccountNumber returns a boolean if a field has been set.

### GetCorrectedRoutingNumber

`func (o *ACHReturnCreateRequest) GetCorrectedRoutingNumber() string`

GetCorrectedRoutingNumber returns the CorrectedRoutingNumber field if non-nil, zero value otherwise.

### GetCorrectedRoutingNumberOk

`func (o *ACHReturnCreateRequest) GetCorrectedRoutingNumberOk() (*string, bool)`

GetCorrectedRoutingNumberOk returns a tuple with the CorrectedRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectedRoutingNumber

`func (o *ACHReturnCreateRequest) SetCorrectedRoutingNumber(v string)`

SetCorrectedRoutingNumber sets CorrectedRoutingNumber field to given value.

### HasCorrectedRoutingNumber

`func (o *ACHReturnCreateRequest) HasCorrectedRoutingNumber() bool`

HasCorrectedRoutingNumber returns a boolean if a field has been set.

### GetId

`func (o *ACHReturnCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ACHReturnCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ACHReturnCreateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetMemberGuid

`func (o *ACHReturnCreateRequest) GetMemberGuid() string`

GetMemberGuid returns the MemberGuid field if non-nil, zero value otherwise.

### GetMemberGuidOk

`func (o *ACHReturnCreateRequest) GetMemberGuidOk() (*string, bool)`

GetMemberGuidOk returns a tuple with the MemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGuid

`func (o *ACHReturnCreateRequest) SetMemberGuid(v string)`

SetMemberGuid sets MemberGuid field to given value.


### GetReturnAccountNumber

`func (o *ACHReturnCreateRequest) GetReturnAccountNumber() string`

GetReturnAccountNumber returns the ReturnAccountNumber field if non-nil, zero value otherwise.

### GetReturnAccountNumberOk

`func (o *ACHReturnCreateRequest) GetReturnAccountNumberOk() (*string, bool)`

GetReturnAccountNumberOk returns a tuple with the ReturnAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAccountNumber

`func (o *ACHReturnCreateRequest) SetReturnAccountNumber(v string)`

SetReturnAccountNumber sets ReturnAccountNumber field to given value.

### HasReturnAccountNumber

`func (o *ACHReturnCreateRequest) HasReturnAccountNumber() bool`

HasReturnAccountNumber returns a boolean if a field has been set.

### GetReturnCode

`func (o *ACHReturnCreateRequest) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ACHReturnCreateRequest) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ACHReturnCreateRequest) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.


### GetReturnNotes

`func (o *ACHReturnCreateRequest) GetReturnNotes() string`

GetReturnNotes returns the ReturnNotes field if non-nil, zero value otherwise.

### GetReturnNotesOk

`func (o *ACHReturnCreateRequest) GetReturnNotesOk() (*string, bool)`

GetReturnNotesOk returns a tuple with the ReturnNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnNotes

`func (o *ACHReturnCreateRequest) SetReturnNotes(v string)`

SetReturnNotes sets ReturnNotes field to given value.

### HasReturnNotes

`func (o *ACHReturnCreateRequest) HasReturnNotes() bool`

HasReturnNotes returns a boolean if a field has been set.

### GetReturnRoutingNumber

`func (o *ACHReturnCreateRequest) GetReturnRoutingNumber() string`

GetReturnRoutingNumber returns the ReturnRoutingNumber field if non-nil, zero value otherwise.

### GetReturnRoutingNumberOk

`func (o *ACHReturnCreateRequest) GetReturnRoutingNumberOk() (*string, bool)`

GetReturnRoutingNumberOk returns a tuple with the ReturnRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnRoutingNumber

`func (o *ACHReturnCreateRequest) SetReturnRoutingNumber(v string)`

SetReturnRoutingNumber sets ReturnRoutingNumber field to given value.

### HasReturnRoutingNumber

`func (o *ACHReturnCreateRequest) HasReturnRoutingNumber() bool`

HasReturnRoutingNumber returns a boolean if a field has been set.

### GetReturnedAt

`func (o *ACHReturnCreateRequest) GetReturnedAt() string`

GetReturnedAt returns the ReturnedAt field if non-nil, zero value otherwise.

### GetReturnedAtOk

`func (o *ACHReturnCreateRequest) GetReturnedAtOk() (*string, bool)`

GetReturnedAtOk returns a tuple with the ReturnedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedAt

`func (o *ACHReturnCreateRequest) SetReturnedAt(v string)`

SetReturnedAt sets ReturnedAt field to given value.

### HasReturnedAt

`func (o *ACHReturnCreateRequest) HasReturnedAt() bool`

HasReturnedAt returns a boolean if a field has been set.

### GetSecCode

`func (o *ACHReturnCreateRequest) GetSecCode() string`

GetSecCode returns the SecCode field if non-nil, zero value otherwise.

### GetSecCodeOk

`func (o *ACHReturnCreateRequest) GetSecCodeOk() (*string, bool)`

GetSecCodeOk returns a tuple with the SecCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecCode

`func (o *ACHReturnCreateRequest) SetSecCode(v string)`

SetSecCode sets SecCode field to given value.

### HasSecCode

`func (o *ACHReturnCreateRequest) HasSecCode() bool`

HasSecCode returns a boolean if a field has been set.

### GetTransactionAmount

`func (o *ACHReturnCreateRequest) GetTransactionAmount() float32`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *ACHReturnCreateRequest) GetTransactionAmountOk() (*float32, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *ACHReturnCreateRequest) SetTransactionAmount(v float32)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *ACHReturnCreateRequest) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetTransactionAmountRange

`func (o *ACHReturnCreateRequest) GetTransactionAmountRange() float32`

GetTransactionAmountRange returns the TransactionAmountRange field if non-nil, zero value otherwise.

### GetTransactionAmountRangeOk

`func (o *ACHReturnCreateRequest) GetTransactionAmountRangeOk() (*float32, bool)`

GetTransactionAmountRangeOk returns a tuple with the TransactionAmountRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmountRange

`func (o *ACHReturnCreateRequest) SetTransactionAmountRange(v float32)`

SetTransactionAmountRange sets TransactionAmountRange field to given value.

### HasTransactionAmountRange

`func (o *ACHReturnCreateRequest) HasTransactionAmountRange() bool`

HasTransactionAmountRange returns a boolean if a field has been set.

### GetUserGuid

`func (o *ACHReturnCreateRequest) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *ACHReturnCreateRequest) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *ACHReturnCreateRequest) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


