# InsightResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAt** | Pointer to **NullableString** |  | [optional] 
**ClientGuid** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**CtaClickedAt** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Guid** | Pointer to **NullableString** |  | [optional] 
**HasAssociatedAccounts** | Pointer to **NullableBool** |  | [optional] 
**HasAssociatedMerchants** | Pointer to **NullableBool** |  | [optional] 
**HasAssociatedScheduledPayments** | Pointer to **NullableBool** |  | [optional] 
**HasAssociatedTransactions** | Pointer to **NullableBool** |  | [optional] 
**HasBeenDisplayed** | Pointer to **NullableBool** |  | [optional] 
**IsDismissed** | Pointer to **NullableBool** |  | [optional] 
**MicroCallToAction** | Pointer to **NullableString** |  | [optional] 
**MicroDescription** | Pointer to **NullableString** |  | [optional] 
**MicroTitle** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**UserGuid** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewInsightResponse

`func NewInsightResponse() *InsightResponse`

NewInsightResponse instantiates a new InsightResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightResponseWithDefaults

`func NewInsightResponseWithDefaults() *InsightResponse`

NewInsightResponseWithDefaults instantiates a new InsightResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveAt

`func (o *InsightResponse) GetActiveAt() string`

GetActiveAt returns the ActiveAt field if non-nil, zero value otherwise.

### GetActiveAtOk

`func (o *InsightResponse) GetActiveAtOk() (*string, bool)`

GetActiveAtOk returns a tuple with the ActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAt

`func (o *InsightResponse) SetActiveAt(v string)`

SetActiveAt sets ActiveAt field to given value.

### HasActiveAt

`func (o *InsightResponse) HasActiveAt() bool`

HasActiveAt returns a boolean if a field has been set.

### SetActiveAtNil

`func (o *InsightResponse) SetActiveAtNil(b bool)`

 SetActiveAtNil sets the value for ActiveAt to be an explicit nil

### UnsetActiveAt
`func (o *InsightResponse) UnsetActiveAt()`

UnsetActiveAt ensures that no value is present for ActiveAt, not even an explicit nil
### GetClientGuid

`func (o *InsightResponse) GetClientGuid() string`

GetClientGuid returns the ClientGuid field if non-nil, zero value otherwise.

### GetClientGuidOk

`func (o *InsightResponse) GetClientGuidOk() (*string, bool)`

GetClientGuidOk returns a tuple with the ClientGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGuid

`func (o *InsightResponse) SetClientGuid(v string)`

SetClientGuid sets ClientGuid field to given value.

### HasClientGuid

`func (o *InsightResponse) HasClientGuid() bool`

HasClientGuid returns a boolean if a field has been set.

### SetClientGuidNil

`func (o *InsightResponse) SetClientGuidNil(b bool)`

 SetClientGuidNil sets the value for ClientGuid to be an explicit nil

### UnsetClientGuid
`func (o *InsightResponse) UnsetClientGuid()`

UnsetClientGuid ensures that no value is present for ClientGuid, not even an explicit nil
### GetCreatedAt

`func (o *InsightResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InsightResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InsightResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InsightResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *InsightResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *InsightResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCtaClickedAt

`func (o *InsightResponse) GetCtaClickedAt() string`

GetCtaClickedAt returns the CtaClickedAt field if non-nil, zero value otherwise.

### GetCtaClickedAtOk

`func (o *InsightResponse) GetCtaClickedAtOk() (*string, bool)`

GetCtaClickedAtOk returns a tuple with the CtaClickedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtaClickedAt

`func (o *InsightResponse) SetCtaClickedAt(v string)`

SetCtaClickedAt sets CtaClickedAt field to given value.

### HasCtaClickedAt

`func (o *InsightResponse) HasCtaClickedAt() bool`

HasCtaClickedAt returns a boolean if a field has been set.

### SetCtaClickedAtNil

`func (o *InsightResponse) SetCtaClickedAtNil(b bool)`

 SetCtaClickedAtNil sets the value for CtaClickedAt to be an explicit nil

### UnsetCtaClickedAt
`func (o *InsightResponse) UnsetCtaClickedAt()`

UnsetCtaClickedAt ensures that no value is present for CtaClickedAt, not even an explicit nil
### GetDescription

`func (o *InsightResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InsightResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InsightResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InsightResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InsightResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InsightResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGuid

`func (o *InsightResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *InsightResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *InsightResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *InsightResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *InsightResponse) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *InsightResponse) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetHasAssociatedAccounts

`func (o *InsightResponse) GetHasAssociatedAccounts() bool`

GetHasAssociatedAccounts returns the HasAssociatedAccounts field if non-nil, zero value otherwise.

### GetHasAssociatedAccountsOk

`func (o *InsightResponse) GetHasAssociatedAccountsOk() (*bool, bool)`

GetHasAssociatedAccountsOk returns a tuple with the HasAssociatedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAssociatedAccounts

`func (o *InsightResponse) SetHasAssociatedAccounts(v bool)`

SetHasAssociatedAccounts sets HasAssociatedAccounts field to given value.

### HasHasAssociatedAccounts

`func (o *InsightResponse) HasHasAssociatedAccounts() bool`

HasHasAssociatedAccounts returns a boolean if a field has been set.

### SetHasAssociatedAccountsNil

`func (o *InsightResponse) SetHasAssociatedAccountsNil(b bool)`

 SetHasAssociatedAccountsNil sets the value for HasAssociatedAccounts to be an explicit nil

### UnsetHasAssociatedAccounts
`func (o *InsightResponse) UnsetHasAssociatedAccounts()`

UnsetHasAssociatedAccounts ensures that no value is present for HasAssociatedAccounts, not even an explicit nil
### GetHasAssociatedMerchants

`func (o *InsightResponse) GetHasAssociatedMerchants() bool`

GetHasAssociatedMerchants returns the HasAssociatedMerchants field if non-nil, zero value otherwise.

### GetHasAssociatedMerchantsOk

`func (o *InsightResponse) GetHasAssociatedMerchantsOk() (*bool, bool)`

GetHasAssociatedMerchantsOk returns a tuple with the HasAssociatedMerchants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAssociatedMerchants

`func (o *InsightResponse) SetHasAssociatedMerchants(v bool)`

SetHasAssociatedMerchants sets HasAssociatedMerchants field to given value.

### HasHasAssociatedMerchants

`func (o *InsightResponse) HasHasAssociatedMerchants() bool`

HasHasAssociatedMerchants returns a boolean if a field has been set.

### SetHasAssociatedMerchantsNil

`func (o *InsightResponse) SetHasAssociatedMerchantsNil(b bool)`

 SetHasAssociatedMerchantsNil sets the value for HasAssociatedMerchants to be an explicit nil

### UnsetHasAssociatedMerchants
`func (o *InsightResponse) UnsetHasAssociatedMerchants()`

UnsetHasAssociatedMerchants ensures that no value is present for HasAssociatedMerchants, not even an explicit nil
### GetHasAssociatedScheduledPayments

`func (o *InsightResponse) GetHasAssociatedScheduledPayments() bool`

GetHasAssociatedScheduledPayments returns the HasAssociatedScheduledPayments field if non-nil, zero value otherwise.

### GetHasAssociatedScheduledPaymentsOk

`func (o *InsightResponse) GetHasAssociatedScheduledPaymentsOk() (*bool, bool)`

GetHasAssociatedScheduledPaymentsOk returns a tuple with the HasAssociatedScheduledPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAssociatedScheduledPayments

`func (o *InsightResponse) SetHasAssociatedScheduledPayments(v bool)`

SetHasAssociatedScheduledPayments sets HasAssociatedScheduledPayments field to given value.

### HasHasAssociatedScheduledPayments

`func (o *InsightResponse) HasHasAssociatedScheduledPayments() bool`

HasHasAssociatedScheduledPayments returns a boolean if a field has been set.

### SetHasAssociatedScheduledPaymentsNil

`func (o *InsightResponse) SetHasAssociatedScheduledPaymentsNil(b bool)`

 SetHasAssociatedScheduledPaymentsNil sets the value for HasAssociatedScheduledPayments to be an explicit nil

### UnsetHasAssociatedScheduledPayments
`func (o *InsightResponse) UnsetHasAssociatedScheduledPayments()`

UnsetHasAssociatedScheduledPayments ensures that no value is present for HasAssociatedScheduledPayments, not even an explicit nil
### GetHasAssociatedTransactions

`func (o *InsightResponse) GetHasAssociatedTransactions() bool`

GetHasAssociatedTransactions returns the HasAssociatedTransactions field if non-nil, zero value otherwise.

### GetHasAssociatedTransactionsOk

`func (o *InsightResponse) GetHasAssociatedTransactionsOk() (*bool, bool)`

GetHasAssociatedTransactionsOk returns a tuple with the HasAssociatedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAssociatedTransactions

`func (o *InsightResponse) SetHasAssociatedTransactions(v bool)`

SetHasAssociatedTransactions sets HasAssociatedTransactions field to given value.

### HasHasAssociatedTransactions

`func (o *InsightResponse) HasHasAssociatedTransactions() bool`

HasHasAssociatedTransactions returns a boolean if a field has been set.

### SetHasAssociatedTransactionsNil

`func (o *InsightResponse) SetHasAssociatedTransactionsNil(b bool)`

 SetHasAssociatedTransactionsNil sets the value for HasAssociatedTransactions to be an explicit nil

### UnsetHasAssociatedTransactions
`func (o *InsightResponse) UnsetHasAssociatedTransactions()`

UnsetHasAssociatedTransactions ensures that no value is present for HasAssociatedTransactions, not even an explicit nil
### GetHasBeenDisplayed

`func (o *InsightResponse) GetHasBeenDisplayed() bool`

GetHasBeenDisplayed returns the HasBeenDisplayed field if non-nil, zero value otherwise.

### GetHasBeenDisplayedOk

`func (o *InsightResponse) GetHasBeenDisplayedOk() (*bool, bool)`

GetHasBeenDisplayedOk returns a tuple with the HasBeenDisplayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBeenDisplayed

`func (o *InsightResponse) SetHasBeenDisplayed(v bool)`

SetHasBeenDisplayed sets HasBeenDisplayed field to given value.

### HasHasBeenDisplayed

`func (o *InsightResponse) HasHasBeenDisplayed() bool`

HasHasBeenDisplayed returns a boolean if a field has been set.

### SetHasBeenDisplayedNil

`func (o *InsightResponse) SetHasBeenDisplayedNil(b bool)`

 SetHasBeenDisplayedNil sets the value for HasBeenDisplayed to be an explicit nil

### UnsetHasBeenDisplayed
`func (o *InsightResponse) UnsetHasBeenDisplayed()`

UnsetHasBeenDisplayed ensures that no value is present for HasBeenDisplayed, not even an explicit nil
### GetIsDismissed

`func (o *InsightResponse) GetIsDismissed() bool`

GetIsDismissed returns the IsDismissed field if non-nil, zero value otherwise.

### GetIsDismissedOk

`func (o *InsightResponse) GetIsDismissedOk() (*bool, bool)`

GetIsDismissedOk returns a tuple with the IsDismissed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDismissed

`func (o *InsightResponse) SetIsDismissed(v bool)`

SetIsDismissed sets IsDismissed field to given value.

### HasIsDismissed

`func (o *InsightResponse) HasIsDismissed() bool`

HasIsDismissed returns a boolean if a field has been set.

### SetIsDismissedNil

`func (o *InsightResponse) SetIsDismissedNil(b bool)`

 SetIsDismissedNil sets the value for IsDismissed to be an explicit nil

### UnsetIsDismissed
`func (o *InsightResponse) UnsetIsDismissed()`

UnsetIsDismissed ensures that no value is present for IsDismissed, not even an explicit nil
### GetMicroCallToAction

`func (o *InsightResponse) GetMicroCallToAction() string`

GetMicroCallToAction returns the MicroCallToAction field if non-nil, zero value otherwise.

### GetMicroCallToActionOk

`func (o *InsightResponse) GetMicroCallToActionOk() (*string, bool)`

GetMicroCallToActionOk returns a tuple with the MicroCallToAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroCallToAction

`func (o *InsightResponse) SetMicroCallToAction(v string)`

SetMicroCallToAction sets MicroCallToAction field to given value.

### HasMicroCallToAction

`func (o *InsightResponse) HasMicroCallToAction() bool`

HasMicroCallToAction returns a boolean if a field has been set.

### SetMicroCallToActionNil

`func (o *InsightResponse) SetMicroCallToActionNil(b bool)`

 SetMicroCallToActionNil sets the value for MicroCallToAction to be an explicit nil

### UnsetMicroCallToAction
`func (o *InsightResponse) UnsetMicroCallToAction()`

UnsetMicroCallToAction ensures that no value is present for MicroCallToAction, not even an explicit nil
### GetMicroDescription

`func (o *InsightResponse) GetMicroDescription() string`

GetMicroDescription returns the MicroDescription field if non-nil, zero value otherwise.

### GetMicroDescriptionOk

`func (o *InsightResponse) GetMicroDescriptionOk() (*string, bool)`

GetMicroDescriptionOk returns a tuple with the MicroDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroDescription

`func (o *InsightResponse) SetMicroDescription(v string)`

SetMicroDescription sets MicroDescription field to given value.

### HasMicroDescription

`func (o *InsightResponse) HasMicroDescription() bool`

HasMicroDescription returns a boolean if a field has been set.

### SetMicroDescriptionNil

`func (o *InsightResponse) SetMicroDescriptionNil(b bool)`

 SetMicroDescriptionNil sets the value for MicroDescription to be an explicit nil

### UnsetMicroDescription
`func (o *InsightResponse) UnsetMicroDescription()`

UnsetMicroDescription ensures that no value is present for MicroDescription, not even an explicit nil
### GetMicroTitle

`func (o *InsightResponse) GetMicroTitle() string`

GetMicroTitle returns the MicroTitle field if non-nil, zero value otherwise.

### GetMicroTitleOk

`func (o *InsightResponse) GetMicroTitleOk() (*string, bool)`

GetMicroTitleOk returns a tuple with the MicroTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroTitle

`func (o *InsightResponse) SetMicroTitle(v string)`

SetMicroTitle sets MicroTitle field to given value.

### HasMicroTitle

`func (o *InsightResponse) HasMicroTitle() bool`

HasMicroTitle returns a boolean if a field has been set.

### SetMicroTitleNil

`func (o *InsightResponse) SetMicroTitleNil(b bool)`

 SetMicroTitleNil sets the value for MicroTitle to be an explicit nil

### UnsetMicroTitle
`func (o *InsightResponse) UnsetMicroTitle()`

UnsetMicroTitle ensures that no value is present for MicroTitle, not even an explicit nil
### GetTemplate

`func (o *InsightResponse) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *InsightResponse) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *InsightResponse) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *InsightResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *InsightResponse) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *InsightResponse) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetTitle

`func (o *InsightResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InsightResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InsightResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InsightResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *InsightResponse) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InsightResponse) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUpdatedAt

`func (o *InsightResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InsightResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InsightResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InsightResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *InsightResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *InsightResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserGuid

`func (o *InsightResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *InsightResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *InsightResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *InsightResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### GetUserId

`func (o *InsightResponse) GetUserId() interface{}`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *InsightResponse) GetUserIdOk() (*interface{}, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *InsightResponse) SetUserId(v interface{})`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *InsightResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *InsightResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *InsightResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


