# ManagedAccountUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** |  | [optional] 
**Apr** | Pointer to **float32** |  | [optional] 
**Apy** | Pointer to **float32** |  | [optional] 
**AvailableBalance** | Pointer to **float32** |  | [optional] 
**AvailableCredit** | Pointer to **float32** |  | [optional] 
**Balance** | Pointer to **float32** |  | [optional] 
**CashSurrenderValue** | Pointer to **float32** |  | [optional] 
**CreditLimit** | Pointer to **float32** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**DayPaymentIsDue** | Pointer to **int32** |  | [optional] 
**DeathBenefit** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InterestRate** | Pointer to **float32** |  | [optional] 
**IsClosed** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**LastPayment** | Pointer to **float32** |  | [optional] 
**LastPaymentAt** | Pointer to **string** |  | [optional] 
**LoanAmount** | Pointer to **float32** |  | [optional] 
**MaturesOn** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**MinimumBalance** | Pointer to **float32** |  | [optional] 
**MinimumPayment** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**OriginalBalance** | Pointer to **float32** |  | [optional] 
**PaymentDueAt** | Pointer to **string** |  | [optional] 
**PayoffBalance** | Pointer to **float32** |  | [optional] 
**RoutingNumber** | Pointer to **string** |  | [optional] 
**StartedOn** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedAccountUpdateRequest

`func NewManagedAccountUpdateRequest() *ManagedAccountUpdateRequest`

NewManagedAccountUpdateRequest instantiates a new ManagedAccountUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAccountUpdateRequestWithDefaults

`func NewManagedAccountUpdateRequestWithDefaults() *ManagedAccountUpdateRequest`

NewManagedAccountUpdateRequestWithDefaults instantiates a new ManagedAccountUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *ManagedAccountUpdateRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *ManagedAccountUpdateRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *ManagedAccountUpdateRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *ManagedAccountUpdateRequest) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetApr

`func (o *ManagedAccountUpdateRequest) GetApr() float32`

GetApr returns the Apr field if non-nil, zero value otherwise.

### GetAprOk

`func (o *ManagedAccountUpdateRequest) GetAprOk() (*float32, bool)`

GetAprOk returns a tuple with the Apr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApr

`func (o *ManagedAccountUpdateRequest) SetApr(v float32)`

SetApr sets Apr field to given value.

### HasApr

`func (o *ManagedAccountUpdateRequest) HasApr() bool`

HasApr returns a boolean if a field has been set.

### GetApy

`func (o *ManagedAccountUpdateRequest) GetApy() float32`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *ManagedAccountUpdateRequest) GetApyOk() (*float32, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *ManagedAccountUpdateRequest) SetApy(v float32)`

SetApy sets Apy field to given value.

### HasApy

`func (o *ManagedAccountUpdateRequest) HasApy() bool`

HasApy returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *ManagedAccountUpdateRequest) GetAvailableBalance() float32`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *ManagedAccountUpdateRequest) GetAvailableBalanceOk() (*float32, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *ManagedAccountUpdateRequest) SetAvailableBalance(v float32)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *ManagedAccountUpdateRequest) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetAvailableCredit

`func (o *ManagedAccountUpdateRequest) GetAvailableCredit() float32`

GetAvailableCredit returns the AvailableCredit field if non-nil, zero value otherwise.

### GetAvailableCreditOk

`func (o *ManagedAccountUpdateRequest) GetAvailableCreditOk() (*float32, bool)`

GetAvailableCreditOk returns a tuple with the AvailableCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCredit

`func (o *ManagedAccountUpdateRequest) SetAvailableCredit(v float32)`

SetAvailableCredit sets AvailableCredit field to given value.

### HasAvailableCredit

`func (o *ManagedAccountUpdateRequest) HasAvailableCredit() bool`

HasAvailableCredit returns a boolean if a field has been set.

### GetBalance

`func (o *ManagedAccountUpdateRequest) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ManagedAccountUpdateRequest) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ManagedAccountUpdateRequest) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *ManagedAccountUpdateRequest) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCashSurrenderValue

`func (o *ManagedAccountUpdateRequest) GetCashSurrenderValue() float32`

GetCashSurrenderValue returns the CashSurrenderValue field if non-nil, zero value otherwise.

### GetCashSurrenderValueOk

`func (o *ManagedAccountUpdateRequest) GetCashSurrenderValueOk() (*float32, bool)`

GetCashSurrenderValueOk returns a tuple with the CashSurrenderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashSurrenderValue

`func (o *ManagedAccountUpdateRequest) SetCashSurrenderValue(v float32)`

SetCashSurrenderValue sets CashSurrenderValue field to given value.

### HasCashSurrenderValue

`func (o *ManagedAccountUpdateRequest) HasCashSurrenderValue() bool`

HasCashSurrenderValue returns a boolean if a field has been set.

### GetCreditLimit

`func (o *ManagedAccountUpdateRequest) GetCreditLimit() float32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *ManagedAccountUpdateRequest) GetCreditLimitOk() (*float32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *ManagedAccountUpdateRequest) SetCreditLimit(v float32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *ManagedAccountUpdateRequest) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ManagedAccountUpdateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ManagedAccountUpdateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ManagedAccountUpdateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ManagedAccountUpdateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDayPaymentIsDue

`func (o *ManagedAccountUpdateRequest) GetDayPaymentIsDue() int32`

GetDayPaymentIsDue returns the DayPaymentIsDue field if non-nil, zero value otherwise.

### GetDayPaymentIsDueOk

`func (o *ManagedAccountUpdateRequest) GetDayPaymentIsDueOk() (*int32, bool)`

GetDayPaymentIsDueOk returns a tuple with the DayPaymentIsDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayPaymentIsDue

`func (o *ManagedAccountUpdateRequest) SetDayPaymentIsDue(v int32)`

SetDayPaymentIsDue sets DayPaymentIsDue field to given value.

### HasDayPaymentIsDue

`func (o *ManagedAccountUpdateRequest) HasDayPaymentIsDue() bool`

HasDayPaymentIsDue returns a boolean if a field has been set.

### GetDeathBenefit

`func (o *ManagedAccountUpdateRequest) GetDeathBenefit() int32`

GetDeathBenefit returns the DeathBenefit field if non-nil, zero value otherwise.

### GetDeathBenefitOk

`func (o *ManagedAccountUpdateRequest) GetDeathBenefitOk() (*int32, bool)`

GetDeathBenefitOk returns a tuple with the DeathBenefit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathBenefit

`func (o *ManagedAccountUpdateRequest) SetDeathBenefit(v int32)`

SetDeathBenefit sets DeathBenefit field to given value.

### HasDeathBenefit

`func (o *ManagedAccountUpdateRequest) HasDeathBenefit() bool`

HasDeathBenefit returns a boolean if a field has been set.

### GetId

`func (o *ManagedAccountUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedAccountUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedAccountUpdateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedAccountUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterestRate

`func (o *ManagedAccountUpdateRequest) GetInterestRate() float32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *ManagedAccountUpdateRequest) GetInterestRateOk() (*float32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *ManagedAccountUpdateRequest) SetInterestRate(v float32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *ManagedAccountUpdateRequest) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetIsClosed

`func (o *ManagedAccountUpdateRequest) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *ManagedAccountUpdateRequest) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *ManagedAccountUpdateRequest) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *ManagedAccountUpdateRequest) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetIsHidden

`func (o *ManagedAccountUpdateRequest) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *ManagedAccountUpdateRequest) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *ManagedAccountUpdateRequest) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *ManagedAccountUpdateRequest) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetLastPayment

`func (o *ManagedAccountUpdateRequest) GetLastPayment() float32`

GetLastPayment returns the LastPayment field if non-nil, zero value otherwise.

### GetLastPaymentOk

`func (o *ManagedAccountUpdateRequest) GetLastPaymentOk() (*float32, bool)`

GetLastPaymentOk returns a tuple with the LastPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPayment

`func (o *ManagedAccountUpdateRequest) SetLastPayment(v float32)`

SetLastPayment sets LastPayment field to given value.

### HasLastPayment

`func (o *ManagedAccountUpdateRequest) HasLastPayment() bool`

HasLastPayment returns a boolean if a field has been set.

### GetLastPaymentAt

`func (o *ManagedAccountUpdateRequest) GetLastPaymentAt() string`

GetLastPaymentAt returns the LastPaymentAt field if non-nil, zero value otherwise.

### GetLastPaymentAtOk

`func (o *ManagedAccountUpdateRequest) GetLastPaymentAtOk() (*string, bool)`

GetLastPaymentAtOk returns a tuple with the LastPaymentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAt

`func (o *ManagedAccountUpdateRequest) SetLastPaymentAt(v string)`

SetLastPaymentAt sets LastPaymentAt field to given value.

### HasLastPaymentAt

`func (o *ManagedAccountUpdateRequest) HasLastPaymentAt() bool`

HasLastPaymentAt returns a boolean if a field has been set.

### GetLoanAmount

`func (o *ManagedAccountUpdateRequest) GetLoanAmount() float32`

GetLoanAmount returns the LoanAmount field if non-nil, zero value otherwise.

### GetLoanAmountOk

`func (o *ManagedAccountUpdateRequest) GetLoanAmountOk() (*float32, bool)`

GetLoanAmountOk returns a tuple with the LoanAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanAmount

`func (o *ManagedAccountUpdateRequest) SetLoanAmount(v float32)`

SetLoanAmount sets LoanAmount field to given value.

### HasLoanAmount

`func (o *ManagedAccountUpdateRequest) HasLoanAmount() bool`

HasLoanAmount returns a boolean if a field has been set.

### GetMaturesOn

`func (o *ManagedAccountUpdateRequest) GetMaturesOn() string`

GetMaturesOn returns the MaturesOn field if non-nil, zero value otherwise.

### GetMaturesOnOk

`func (o *ManagedAccountUpdateRequest) GetMaturesOnOk() (*string, bool)`

GetMaturesOnOk returns a tuple with the MaturesOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturesOn

`func (o *ManagedAccountUpdateRequest) SetMaturesOn(v string)`

SetMaturesOn sets MaturesOn field to given value.

### HasMaturesOn

`func (o *ManagedAccountUpdateRequest) HasMaturesOn() bool`

HasMaturesOn returns a boolean if a field has been set.

### GetMetadata

`func (o *ManagedAccountUpdateRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ManagedAccountUpdateRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ManagedAccountUpdateRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ManagedAccountUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMinimumBalance

`func (o *ManagedAccountUpdateRequest) GetMinimumBalance() float32`

GetMinimumBalance returns the MinimumBalance field if non-nil, zero value otherwise.

### GetMinimumBalanceOk

`func (o *ManagedAccountUpdateRequest) GetMinimumBalanceOk() (*float32, bool)`

GetMinimumBalanceOk returns a tuple with the MinimumBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBalance

`func (o *ManagedAccountUpdateRequest) SetMinimumBalance(v float32)`

SetMinimumBalance sets MinimumBalance field to given value.

### HasMinimumBalance

`func (o *ManagedAccountUpdateRequest) HasMinimumBalance() bool`

HasMinimumBalance returns a boolean if a field has been set.

### GetMinimumPayment

`func (o *ManagedAccountUpdateRequest) GetMinimumPayment() float32`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *ManagedAccountUpdateRequest) GetMinimumPaymentOk() (*float32, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *ManagedAccountUpdateRequest) SetMinimumPayment(v float32)`

SetMinimumPayment sets MinimumPayment field to given value.

### HasMinimumPayment

`func (o *ManagedAccountUpdateRequest) HasMinimumPayment() bool`

HasMinimumPayment returns a boolean if a field has been set.

### GetName

`func (o *ManagedAccountUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedAccountUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedAccountUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagedAccountUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *ManagedAccountUpdateRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *ManagedAccountUpdateRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *ManagedAccountUpdateRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *ManagedAccountUpdateRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetOriginalBalance

`func (o *ManagedAccountUpdateRequest) GetOriginalBalance() float32`

GetOriginalBalance returns the OriginalBalance field if non-nil, zero value otherwise.

### GetOriginalBalanceOk

`func (o *ManagedAccountUpdateRequest) GetOriginalBalanceOk() (*float32, bool)`

GetOriginalBalanceOk returns a tuple with the OriginalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalBalance

`func (o *ManagedAccountUpdateRequest) SetOriginalBalance(v float32)`

SetOriginalBalance sets OriginalBalance field to given value.

### HasOriginalBalance

`func (o *ManagedAccountUpdateRequest) HasOriginalBalance() bool`

HasOriginalBalance returns a boolean if a field has been set.

### GetPaymentDueAt

`func (o *ManagedAccountUpdateRequest) GetPaymentDueAt() string`

GetPaymentDueAt returns the PaymentDueAt field if non-nil, zero value otherwise.

### GetPaymentDueAtOk

`func (o *ManagedAccountUpdateRequest) GetPaymentDueAtOk() (*string, bool)`

GetPaymentDueAtOk returns a tuple with the PaymentDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueAt

`func (o *ManagedAccountUpdateRequest) SetPaymentDueAt(v string)`

SetPaymentDueAt sets PaymentDueAt field to given value.

### HasPaymentDueAt

`func (o *ManagedAccountUpdateRequest) HasPaymentDueAt() bool`

HasPaymentDueAt returns a boolean if a field has been set.

### GetPayoffBalance

`func (o *ManagedAccountUpdateRequest) GetPayoffBalance() float32`

GetPayoffBalance returns the PayoffBalance field if non-nil, zero value otherwise.

### GetPayoffBalanceOk

`func (o *ManagedAccountUpdateRequest) GetPayoffBalanceOk() (*float32, bool)`

GetPayoffBalanceOk returns a tuple with the PayoffBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoffBalance

`func (o *ManagedAccountUpdateRequest) SetPayoffBalance(v float32)`

SetPayoffBalance sets PayoffBalance field to given value.

### HasPayoffBalance

`func (o *ManagedAccountUpdateRequest) HasPayoffBalance() bool`

HasPayoffBalance returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *ManagedAccountUpdateRequest) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *ManagedAccountUpdateRequest) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *ManagedAccountUpdateRequest) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *ManagedAccountUpdateRequest) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetStartedOn

`func (o *ManagedAccountUpdateRequest) GetStartedOn() string`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *ManagedAccountUpdateRequest) GetStartedOnOk() (*string, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *ManagedAccountUpdateRequest) SetStartedOn(v string)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *ManagedAccountUpdateRequest) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### GetSubtype

`func (o *ManagedAccountUpdateRequest) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ManagedAccountUpdateRequest) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ManagedAccountUpdateRequest) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *ManagedAccountUpdateRequest) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetType

`func (o *ManagedAccountUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagedAccountUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagedAccountUpdateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManagedAccountUpdateRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


