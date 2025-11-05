# ACHResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | Pointer to **string** |  | [optional] 
**AccountNumberLastFour** | Pointer to **NullableString** |  | [optional] 
**AccountType** | Pointer to **NullableString** |  | [optional] 
**AchInitiatedAt** | Pointer to **NullableString** |  | [optional] 
**ClientGuid** | Pointer to **string** |  | [optional] 
**CorrectedAccountNumber** | Pointer to **NullableString** |  | [optional] 
**CorrectedRoutingNumber** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InstitutionGuid** | Pointer to **string** |  | [optional] 
**InvestigationNotes** | Pointer to **NullableString** |  | [optional] 
**MemberGuid** | Pointer to **string** |  | [optional] 
**ProcessingErrors** | Pointer to **NullableString** |  | [optional] 
**ResolutionCode** | Pointer to **NullableString** |  | [optional] 
**ResolutionDetail** | Pointer to **NullableString** |  | [optional] 
**ResolvedStatusAt** | Pointer to **NullableString** |  | [optional] 
**ReturnCode** | Pointer to **string** |  | [optional] 
**ReturnNotes** | Pointer to **NullableString** |  | [optional] 
**ReturnAccountNumber** | Pointer to **NullableString** |  | [optional] 
**ReturnRoutingNumber** | Pointer to **NullableString** |  | [optional] 
**ReturnStatus** | Pointer to **NullableString** |  | [optional] 
**ReturnedAt** | Pointer to **NullableString** |  | [optional] 
**SecCode** | Pointer to **NullableString** |  | [optional] 
**StartedProcessingAt** | Pointer to **NullableString** |  | [optional] 
**SubmittedAt** | Pointer to **NullableString** |  | [optional] 
**TransactionAmount** | Pointer to **NullableFloat64** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserGuid** | Pointer to **string** |  | [optional] 

## Methods

### NewACHResponse

`func NewACHResponse() *ACHResponse`

NewACHResponse instantiates a new ACHResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACHResponseWithDefaults

`func NewACHResponseWithDefaults() *ACHResponse`

NewACHResponseWithDefaults instantiates a new ACHResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *ACHResponse) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *ACHResponse) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *ACHResponse) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.

### HasAccountGuid

`func (o *ACHResponse) HasAccountGuid() bool`

HasAccountGuid returns a boolean if a field has been set.

### GetAccountNumberLastFour

`func (o *ACHResponse) GetAccountNumberLastFour() string`

GetAccountNumberLastFour returns the AccountNumberLastFour field if non-nil, zero value otherwise.

### GetAccountNumberLastFourOk

`func (o *ACHResponse) GetAccountNumberLastFourOk() (*string, bool)`

GetAccountNumberLastFourOk returns a tuple with the AccountNumberLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberLastFour

`func (o *ACHResponse) SetAccountNumberLastFour(v string)`

SetAccountNumberLastFour sets AccountNumberLastFour field to given value.

### HasAccountNumberLastFour

`func (o *ACHResponse) HasAccountNumberLastFour() bool`

HasAccountNumberLastFour returns a boolean if a field has been set.

### SetAccountNumberLastFourNil

`func (o *ACHResponse) SetAccountNumberLastFourNil(b bool)`

 SetAccountNumberLastFourNil sets the value for AccountNumberLastFour to be an explicit nil

### UnsetAccountNumberLastFour
`func (o *ACHResponse) UnsetAccountNumberLastFour()`

UnsetAccountNumberLastFour ensures that no value is present for AccountNumberLastFour, not even an explicit nil
### GetAccountType

`func (o *ACHResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ACHResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ACHResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ACHResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountTypeNil

`func (o *ACHResponse) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *ACHResponse) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetAchInitiatedAt

`func (o *ACHResponse) GetAchInitiatedAt() string`

GetAchInitiatedAt returns the AchInitiatedAt field if non-nil, zero value otherwise.

### GetAchInitiatedAtOk

`func (o *ACHResponse) GetAchInitiatedAtOk() (*string, bool)`

GetAchInitiatedAtOk returns a tuple with the AchInitiatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchInitiatedAt

`func (o *ACHResponse) SetAchInitiatedAt(v string)`

SetAchInitiatedAt sets AchInitiatedAt field to given value.

### HasAchInitiatedAt

`func (o *ACHResponse) HasAchInitiatedAt() bool`

HasAchInitiatedAt returns a boolean if a field has been set.

### SetAchInitiatedAtNil

`func (o *ACHResponse) SetAchInitiatedAtNil(b bool)`

 SetAchInitiatedAtNil sets the value for AchInitiatedAt to be an explicit nil

### UnsetAchInitiatedAt
`func (o *ACHResponse) UnsetAchInitiatedAt()`

UnsetAchInitiatedAt ensures that no value is present for AchInitiatedAt, not even an explicit nil
### GetClientGuid

`func (o *ACHResponse) GetClientGuid() string`

GetClientGuid returns the ClientGuid field if non-nil, zero value otherwise.

### GetClientGuidOk

`func (o *ACHResponse) GetClientGuidOk() (*string, bool)`

GetClientGuidOk returns a tuple with the ClientGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGuid

`func (o *ACHResponse) SetClientGuid(v string)`

SetClientGuid sets ClientGuid field to given value.

### HasClientGuid

`func (o *ACHResponse) HasClientGuid() bool`

HasClientGuid returns a boolean if a field has been set.

### GetCorrectedAccountNumber

`func (o *ACHResponse) GetCorrectedAccountNumber() string`

GetCorrectedAccountNumber returns the CorrectedAccountNumber field if non-nil, zero value otherwise.

### GetCorrectedAccountNumberOk

`func (o *ACHResponse) GetCorrectedAccountNumberOk() (*string, bool)`

GetCorrectedAccountNumberOk returns a tuple with the CorrectedAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectedAccountNumber

`func (o *ACHResponse) SetCorrectedAccountNumber(v string)`

SetCorrectedAccountNumber sets CorrectedAccountNumber field to given value.

### HasCorrectedAccountNumber

`func (o *ACHResponse) HasCorrectedAccountNumber() bool`

HasCorrectedAccountNumber returns a boolean if a field has been set.

### SetCorrectedAccountNumberNil

`func (o *ACHResponse) SetCorrectedAccountNumberNil(b bool)`

 SetCorrectedAccountNumberNil sets the value for CorrectedAccountNumber to be an explicit nil

### UnsetCorrectedAccountNumber
`func (o *ACHResponse) UnsetCorrectedAccountNumber()`

UnsetCorrectedAccountNumber ensures that no value is present for CorrectedAccountNumber, not even an explicit nil
### GetCorrectedRoutingNumber

`func (o *ACHResponse) GetCorrectedRoutingNumber() string`

GetCorrectedRoutingNumber returns the CorrectedRoutingNumber field if non-nil, zero value otherwise.

### GetCorrectedRoutingNumberOk

`func (o *ACHResponse) GetCorrectedRoutingNumberOk() (*string, bool)`

GetCorrectedRoutingNumberOk returns a tuple with the CorrectedRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectedRoutingNumber

`func (o *ACHResponse) SetCorrectedRoutingNumber(v string)`

SetCorrectedRoutingNumber sets CorrectedRoutingNumber field to given value.

### HasCorrectedRoutingNumber

`func (o *ACHResponse) HasCorrectedRoutingNumber() bool`

HasCorrectedRoutingNumber returns a boolean if a field has been set.

### SetCorrectedRoutingNumberNil

`func (o *ACHResponse) SetCorrectedRoutingNumberNil(b bool)`

 SetCorrectedRoutingNumberNil sets the value for CorrectedRoutingNumber to be an explicit nil

### UnsetCorrectedRoutingNumber
`func (o *ACHResponse) UnsetCorrectedRoutingNumber()`

UnsetCorrectedRoutingNumber ensures that no value is present for CorrectedRoutingNumber, not even an explicit nil
### GetCreatedAt

`func (o *ACHResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ACHResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ACHResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ACHResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *ACHResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ACHResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ACHResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ACHResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetId

`func (o *ACHResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ACHResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ACHResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ACHResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstitutionGuid

`func (o *ACHResponse) GetInstitutionGuid() string`

GetInstitutionGuid returns the InstitutionGuid field if non-nil, zero value otherwise.

### GetInstitutionGuidOk

`func (o *ACHResponse) GetInstitutionGuidOk() (*string, bool)`

GetInstitutionGuidOk returns a tuple with the InstitutionGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionGuid

`func (o *ACHResponse) SetInstitutionGuid(v string)`

SetInstitutionGuid sets InstitutionGuid field to given value.

### HasInstitutionGuid

`func (o *ACHResponse) HasInstitutionGuid() bool`

HasInstitutionGuid returns a boolean if a field has been set.

### GetInvestigationNotes

`func (o *ACHResponse) GetInvestigationNotes() string`

GetInvestigationNotes returns the InvestigationNotes field if non-nil, zero value otherwise.

### GetInvestigationNotesOk

`func (o *ACHResponse) GetInvestigationNotesOk() (*string, bool)`

GetInvestigationNotesOk returns a tuple with the InvestigationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationNotes

`func (o *ACHResponse) SetInvestigationNotes(v string)`

SetInvestigationNotes sets InvestigationNotes field to given value.

### HasInvestigationNotes

`func (o *ACHResponse) HasInvestigationNotes() bool`

HasInvestigationNotes returns a boolean if a field has been set.

### SetInvestigationNotesNil

`func (o *ACHResponse) SetInvestigationNotesNil(b bool)`

 SetInvestigationNotesNil sets the value for InvestigationNotes to be an explicit nil

### UnsetInvestigationNotes
`func (o *ACHResponse) UnsetInvestigationNotes()`

UnsetInvestigationNotes ensures that no value is present for InvestigationNotes, not even an explicit nil
### GetMemberGuid

`func (o *ACHResponse) GetMemberGuid() string`

GetMemberGuid returns the MemberGuid field if non-nil, zero value otherwise.

### GetMemberGuidOk

`func (o *ACHResponse) GetMemberGuidOk() (*string, bool)`

GetMemberGuidOk returns a tuple with the MemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGuid

`func (o *ACHResponse) SetMemberGuid(v string)`

SetMemberGuid sets MemberGuid field to given value.

### HasMemberGuid

`func (o *ACHResponse) HasMemberGuid() bool`

HasMemberGuid returns a boolean if a field has been set.

### GetProcessingErrors

`func (o *ACHResponse) GetProcessingErrors() string`

GetProcessingErrors returns the ProcessingErrors field if non-nil, zero value otherwise.

### GetProcessingErrorsOk

`func (o *ACHResponse) GetProcessingErrorsOk() (*string, bool)`

GetProcessingErrorsOk returns a tuple with the ProcessingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingErrors

`func (o *ACHResponse) SetProcessingErrors(v string)`

SetProcessingErrors sets ProcessingErrors field to given value.

### HasProcessingErrors

`func (o *ACHResponse) HasProcessingErrors() bool`

HasProcessingErrors returns a boolean if a field has been set.

### SetProcessingErrorsNil

`func (o *ACHResponse) SetProcessingErrorsNil(b bool)`

 SetProcessingErrorsNil sets the value for ProcessingErrors to be an explicit nil

### UnsetProcessingErrors
`func (o *ACHResponse) UnsetProcessingErrors()`

UnsetProcessingErrors ensures that no value is present for ProcessingErrors, not even an explicit nil
### GetResolutionCode

`func (o *ACHResponse) GetResolutionCode() string`

GetResolutionCode returns the ResolutionCode field if non-nil, zero value otherwise.

### GetResolutionCodeOk

`func (o *ACHResponse) GetResolutionCodeOk() (*string, bool)`

GetResolutionCodeOk returns a tuple with the ResolutionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionCode

`func (o *ACHResponse) SetResolutionCode(v string)`

SetResolutionCode sets ResolutionCode field to given value.

### HasResolutionCode

`func (o *ACHResponse) HasResolutionCode() bool`

HasResolutionCode returns a boolean if a field has been set.

### SetResolutionCodeNil

`func (o *ACHResponse) SetResolutionCodeNil(b bool)`

 SetResolutionCodeNil sets the value for ResolutionCode to be an explicit nil

### UnsetResolutionCode
`func (o *ACHResponse) UnsetResolutionCode()`

UnsetResolutionCode ensures that no value is present for ResolutionCode, not even an explicit nil
### GetResolutionDetail

`func (o *ACHResponse) GetResolutionDetail() string`

GetResolutionDetail returns the ResolutionDetail field if non-nil, zero value otherwise.

### GetResolutionDetailOk

`func (o *ACHResponse) GetResolutionDetailOk() (*string, bool)`

GetResolutionDetailOk returns a tuple with the ResolutionDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDetail

`func (o *ACHResponse) SetResolutionDetail(v string)`

SetResolutionDetail sets ResolutionDetail field to given value.

### HasResolutionDetail

`func (o *ACHResponse) HasResolutionDetail() bool`

HasResolutionDetail returns a boolean if a field has been set.

### SetResolutionDetailNil

`func (o *ACHResponse) SetResolutionDetailNil(b bool)`

 SetResolutionDetailNil sets the value for ResolutionDetail to be an explicit nil

### UnsetResolutionDetail
`func (o *ACHResponse) UnsetResolutionDetail()`

UnsetResolutionDetail ensures that no value is present for ResolutionDetail, not even an explicit nil
### GetResolvedStatusAt

`func (o *ACHResponse) GetResolvedStatusAt() string`

GetResolvedStatusAt returns the ResolvedStatusAt field if non-nil, zero value otherwise.

### GetResolvedStatusAtOk

`func (o *ACHResponse) GetResolvedStatusAtOk() (*string, bool)`

GetResolvedStatusAtOk returns a tuple with the ResolvedStatusAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedStatusAt

`func (o *ACHResponse) SetResolvedStatusAt(v string)`

SetResolvedStatusAt sets ResolvedStatusAt field to given value.

### HasResolvedStatusAt

`func (o *ACHResponse) HasResolvedStatusAt() bool`

HasResolvedStatusAt returns a boolean if a field has been set.

### SetResolvedStatusAtNil

`func (o *ACHResponse) SetResolvedStatusAtNil(b bool)`

 SetResolvedStatusAtNil sets the value for ResolvedStatusAt to be an explicit nil

### UnsetResolvedStatusAt
`func (o *ACHResponse) UnsetResolvedStatusAt()`

UnsetResolvedStatusAt ensures that no value is present for ResolvedStatusAt, not even an explicit nil
### GetReturnCode

`func (o *ACHResponse) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ACHResponse) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ACHResponse) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ACHResponse) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnNotes

`func (o *ACHResponse) GetReturnNotes() string`

GetReturnNotes returns the ReturnNotes field if non-nil, zero value otherwise.

### GetReturnNotesOk

`func (o *ACHResponse) GetReturnNotesOk() (*string, bool)`

GetReturnNotesOk returns a tuple with the ReturnNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnNotes

`func (o *ACHResponse) SetReturnNotes(v string)`

SetReturnNotes sets ReturnNotes field to given value.

### HasReturnNotes

`func (o *ACHResponse) HasReturnNotes() bool`

HasReturnNotes returns a boolean if a field has been set.

### SetReturnNotesNil

`func (o *ACHResponse) SetReturnNotesNil(b bool)`

 SetReturnNotesNil sets the value for ReturnNotes to be an explicit nil

### UnsetReturnNotes
`func (o *ACHResponse) UnsetReturnNotes()`

UnsetReturnNotes ensures that no value is present for ReturnNotes, not even an explicit nil
### GetReturnAccountNumber

`func (o *ACHResponse) GetReturnAccountNumber() string`

GetReturnAccountNumber returns the ReturnAccountNumber field if non-nil, zero value otherwise.

### GetReturnAccountNumberOk

`func (o *ACHResponse) GetReturnAccountNumberOk() (*string, bool)`

GetReturnAccountNumberOk returns a tuple with the ReturnAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAccountNumber

`func (o *ACHResponse) SetReturnAccountNumber(v string)`

SetReturnAccountNumber sets ReturnAccountNumber field to given value.

### HasReturnAccountNumber

`func (o *ACHResponse) HasReturnAccountNumber() bool`

HasReturnAccountNumber returns a boolean if a field has been set.

### SetReturnAccountNumberNil

`func (o *ACHResponse) SetReturnAccountNumberNil(b bool)`

 SetReturnAccountNumberNil sets the value for ReturnAccountNumber to be an explicit nil

### UnsetReturnAccountNumber
`func (o *ACHResponse) UnsetReturnAccountNumber()`

UnsetReturnAccountNumber ensures that no value is present for ReturnAccountNumber, not even an explicit nil
### GetReturnRoutingNumber

`func (o *ACHResponse) GetReturnRoutingNumber() string`

GetReturnRoutingNumber returns the ReturnRoutingNumber field if non-nil, zero value otherwise.

### GetReturnRoutingNumberOk

`func (o *ACHResponse) GetReturnRoutingNumberOk() (*string, bool)`

GetReturnRoutingNumberOk returns a tuple with the ReturnRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnRoutingNumber

`func (o *ACHResponse) SetReturnRoutingNumber(v string)`

SetReturnRoutingNumber sets ReturnRoutingNumber field to given value.

### HasReturnRoutingNumber

`func (o *ACHResponse) HasReturnRoutingNumber() bool`

HasReturnRoutingNumber returns a boolean if a field has been set.

### SetReturnRoutingNumberNil

`func (o *ACHResponse) SetReturnRoutingNumberNil(b bool)`

 SetReturnRoutingNumberNil sets the value for ReturnRoutingNumber to be an explicit nil

### UnsetReturnRoutingNumber
`func (o *ACHResponse) UnsetReturnRoutingNumber()`

UnsetReturnRoutingNumber ensures that no value is present for ReturnRoutingNumber, not even an explicit nil
### GetReturnStatus

`func (o *ACHResponse) GetReturnStatus() string`

GetReturnStatus returns the ReturnStatus field if non-nil, zero value otherwise.

### GetReturnStatusOk

`func (o *ACHResponse) GetReturnStatusOk() (*string, bool)`

GetReturnStatusOk returns a tuple with the ReturnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnStatus

`func (o *ACHResponse) SetReturnStatus(v string)`

SetReturnStatus sets ReturnStatus field to given value.

### HasReturnStatus

`func (o *ACHResponse) HasReturnStatus() bool`

HasReturnStatus returns a boolean if a field has been set.

### SetReturnStatusNil

`func (o *ACHResponse) SetReturnStatusNil(b bool)`

 SetReturnStatusNil sets the value for ReturnStatus to be an explicit nil

### UnsetReturnStatus
`func (o *ACHResponse) UnsetReturnStatus()`

UnsetReturnStatus ensures that no value is present for ReturnStatus, not even an explicit nil
### GetReturnedAt

`func (o *ACHResponse) GetReturnedAt() string`

GetReturnedAt returns the ReturnedAt field if non-nil, zero value otherwise.

### GetReturnedAtOk

`func (o *ACHResponse) GetReturnedAtOk() (*string, bool)`

GetReturnedAtOk returns a tuple with the ReturnedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedAt

`func (o *ACHResponse) SetReturnedAt(v string)`

SetReturnedAt sets ReturnedAt field to given value.

### HasReturnedAt

`func (o *ACHResponse) HasReturnedAt() bool`

HasReturnedAt returns a boolean if a field has been set.

### SetReturnedAtNil

`func (o *ACHResponse) SetReturnedAtNil(b bool)`

 SetReturnedAtNil sets the value for ReturnedAt to be an explicit nil

### UnsetReturnedAt
`func (o *ACHResponse) UnsetReturnedAt()`

UnsetReturnedAt ensures that no value is present for ReturnedAt, not even an explicit nil
### GetSecCode

`func (o *ACHResponse) GetSecCode() string`

GetSecCode returns the SecCode field if non-nil, zero value otherwise.

### GetSecCodeOk

`func (o *ACHResponse) GetSecCodeOk() (*string, bool)`

GetSecCodeOk returns a tuple with the SecCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecCode

`func (o *ACHResponse) SetSecCode(v string)`

SetSecCode sets SecCode field to given value.

### HasSecCode

`func (o *ACHResponse) HasSecCode() bool`

HasSecCode returns a boolean if a field has been set.

### SetSecCodeNil

`func (o *ACHResponse) SetSecCodeNil(b bool)`

 SetSecCodeNil sets the value for SecCode to be an explicit nil

### UnsetSecCode
`func (o *ACHResponse) UnsetSecCode()`

UnsetSecCode ensures that no value is present for SecCode, not even an explicit nil
### GetStartedProcessingAt

`func (o *ACHResponse) GetStartedProcessingAt() string`

GetStartedProcessingAt returns the StartedProcessingAt field if non-nil, zero value otherwise.

### GetStartedProcessingAtOk

`func (o *ACHResponse) GetStartedProcessingAtOk() (*string, bool)`

GetStartedProcessingAtOk returns a tuple with the StartedProcessingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedProcessingAt

`func (o *ACHResponse) SetStartedProcessingAt(v string)`

SetStartedProcessingAt sets StartedProcessingAt field to given value.

### HasStartedProcessingAt

`func (o *ACHResponse) HasStartedProcessingAt() bool`

HasStartedProcessingAt returns a boolean if a field has been set.

### SetStartedProcessingAtNil

`func (o *ACHResponse) SetStartedProcessingAtNil(b bool)`

 SetStartedProcessingAtNil sets the value for StartedProcessingAt to be an explicit nil

### UnsetStartedProcessingAt
`func (o *ACHResponse) UnsetStartedProcessingAt()`

UnsetStartedProcessingAt ensures that no value is present for StartedProcessingAt, not even an explicit nil
### GetSubmittedAt

`func (o *ACHResponse) GetSubmittedAt() string`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *ACHResponse) GetSubmittedAtOk() (*string, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *ACHResponse) SetSubmittedAt(v string)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *ACHResponse) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### SetSubmittedAtNil

`func (o *ACHResponse) SetSubmittedAtNil(b bool)`

 SetSubmittedAtNil sets the value for SubmittedAt to be an explicit nil

### UnsetSubmittedAt
`func (o *ACHResponse) UnsetSubmittedAt()`

UnsetSubmittedAt ensures that no value is present for SubmittedAt, not even an explicit nil
### GetTransactionAmount

`func (o *ACHResponse) GetTransactionAmount() float64`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *ACHResponse) GetTransactionAmountOk() (*float64, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *ACHResponse) SetTransactionAmount(v float64)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *ACHResponse) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### SetTransactionAmountNil

`func (o *ACHResponse) SetTransactionAmountNil(b bool)`

 SetTransactionAmountNil sets the value for TransactionAmount to be an explicit nil

### UnsetTransactionAmount
`func (o *ACHResponse) UnsetTransactionAmount()`

UnsetTransactionAmount ensures that no value is present for TransactionAmount, not even an explicit nil
### GetUpdatedAt

`func (o *ACHResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ACHResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ACHResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ACHResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserGuid

`func (o *ACHResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *ACHResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *ACHResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *ACHResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


