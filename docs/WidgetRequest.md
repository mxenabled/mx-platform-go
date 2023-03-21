# WidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientRedirectUrl** | Pointer to **string** |  | [optional] 
**ColorScheme** | Pointer to **string** |  | [optional] 
**CurrentInstitutionCode** | Pointer to **string** |  | [optional] 
**CurrentInstitutionGuid** | Pointer to **string** |  | [optional] 
**CurrentMemberGuid** | Pointer to **string** |  | [optional] 
**DisableBackgroundAgg** | Pointer to **bool** |  | [optional] 
**DisableInstitutionSearch** | Pointer to **bool** |  | [optional] 
**IncludeIdentity** | Pointer to **bool** |  | [optional] 
**IncludeTransactions** | Pointer to **bool** |  | [optional] 
**IsMobileWebview** | Pointer to **bool** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**OauthReferralSource** | Pointer to **string** |  | [optional] 
**UiMessageVersion** | Pointer to **int32** |  | [optional] 
**UiMessageWebviewUrlScheme** | Pointer to **string** |  | [optional] 
**UpdateCredentials** | Pointer to **bool** |  | [optional] 
**WidgetType** | **string** |  | 

## Methods

### NewWidgetRequest

`func NewWidgetRequest(widgetType string, ) *WidgetRequest`

NewWidgetRequest instantiates a new WidgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetRequestWithDefaults

`func NewWidgetRequestWithDefaults() *WidgetRequest`

NewWidgetRequestWithDefaults instantiates a new WidgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientRedirectUrl

`func (o *WidgetRequest) GetClientRedirectUrl() string`

GetClientRedirectUrl returns the ClientRedirectUrl field if non-nil, zero value otherwise.

### GetClientRedirectUrlOk

`func (o *WidgetRequest) GetClientRedirectUrlOk() (*string, bool)`

GetClientRedirectUrlOk returns a tuple with the ClientRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRedirectUrl

`func (o *WidgetRequest) SetClientRedirectUrl(v string)`

SetClientRedirectUrl sets ClientRedirectUrl field to given value.

### HasClientRedirectUrl

`func (o *WidgetRequest) HasClientRedirectUrl() bool`

HasClientRedirectUrl returns a boolean if a field has been set.

### GetColorScheme

`func (o *WidgetRequest) GetColorScheme() string`

GetColorScheme returns the ColorScheme field if non-nil, zero value otherwise.

### GetColorSchemeOk

`func (o *WidgetRequest) GetColorSchemeOk() (*string, bool)`

GetColorSchemeOk returns a tuple with the ColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorScheme

`func (o *WidgetRequest) SetColorScheme(v string)`

SetColorScheme sets ColorScheme field to given value.

### HasColorScheme

`func (o *WidgetRequest) HasColorScheme() bool`

HasColorScheme returns a boolean if a field has been set.

### GetCurrentInstitutionCode

`func (o *WidgetRequest) GetCurrentInstitutionCode() string`

GetCurrentInstitutionCode returns the CurrentInstitutionCode field if non-nil, zero value otherwise.

### GetCurrentInstitutionCodeOk

`func (o *WidgetRequest) GetCurrentInstitutionCodeOk() (*string, bool)`

GetCurrentInstitutionCodeOk returns a tuple with the CurrentInstitutionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentInstitutionCode

`func (o *WidgetRequest) SetCurrentInstitutionCode(v string)`

SetCurrentInstitutionCode sets CurrentInstitutionCode field to given value.

### HasCurrentInstitutionCode

`func (o *WidgetRequest) HasCurrentInstitutionCode() bool`

HasCurrentInstitutionCode returns a boolean if a field has been set.

### GetCurrentInstitutionGuid

`func (o *WidgetRequest) GetCurrentInstitutionGuid() string`

GetCurrentInstitutionGuid returns the CurrentInstitutionGuid field if non-nil, zero value otherwise.

### GetCurrentInstitutionGuidOk

`func (o *WidgetRequest) GetCurrentInstitutionGuidOk() (*string, bool)`

GetCurrentInstitutionGuidOk returns a tuple with the CurrentInstitutionGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentInstitutionGuid

`func (o *WidgetRequest) SetCurrentInstitutionGuid(v string)`

SetCurrentInstitutionGuid sets CurrentInstitutionGuid field to given value.

### HasCurrentInstitutionGuid

`func (o *WidgetRequest) HasCurrentInstitutionGuid() bool`

HasCurrentInstitutionGuid returns a boolean if a field has been set.

### GetCurrentMemberGuid

`func (o *WidgetRequest) GetCurrentMemberGuid() string`

GetCurrentMemberGuid returns the CurrentMemberGuid field if non-nil, zero value otherwise.

### GetCurrentMemberGuidOk

`func (o *WidgetRequest) GetCurrentMemberGuidOk() (*string, bool)`

GetCurrentMemberGuidOk returns a tuple with the CurrentMemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMemberGuid

`func (o *WidgetRequest) SetCurrentMemberGuid(v string)`

SetCurrentMemberGuid sets CurrentMemberGuid field to given value.

### HasCurrentMemberGuid

`func (o *WidgetRequest) HasCurrentMemberGuid() bool`

HasCurrentMemberGuid returns a boolean if a field has been set.

### GetDisableBackgroundAgg

`func (o *WidgetRequest) GetDisableBackgroundAgg() bool`

GetDisableBackgroundAgg returns the DisableBackgroundAgg field if non-nil, zero value otherwise.

### GetDisableBackgroundAggOk

`func (o *WidgetRequest) GetDisableBackgroundAggOk() (*bool, bool)`

GetDisableBackgroundAggOk returns a tuple with the DisableBackgroundAgg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBackgroundAgg

`func (o *WidgetRequest) SetDisableBackgroundAgg(v bool)`

SetDisableBackgroundAgg sets DisableBackgroundAgg field to given value.

### HasDisableBackgroundAgg

`func (o *WidgetRequest) HasDisableBackgroundAgg() bool`

HasDisableBackgroundAgg returns a boolean if a field has been set.

### GetDisableInstitutionSearch

`func (o *WidgetRequest) GetDisableInstitutionSearch() bool`

GetDisableInstitutionSearch returns the DisableInstitutionSearch field if non-nil, zero value otherwise.

### GetDisableInstitutionSearchOk

`func (o *WidgetRequest) GetDisableInstitutionSearchOk() (*bool, bool)`

GetDisableInstitutionSearchOk returns a tuple with the DisableInstitutionSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableInstitutionSearch

`func (o *WidgetRequest) SetDisableInstitutionSearch(v bool)`

SetDisableInstitutionSearch sets DisableInstitutionSearch field to given value.

### HasDisableInstitutionSearch

`func (o *WidgetRequest) HasDisableInstitutionSearch() bool`

HasDisableInstitutionSearch returns a boolean if a field has been set.

### GetIncludeIdentity

`func (o *WidgetRequest) GetIncludeIdentity() bool`

GetIncludeIdentity returns the IncludeIdentity field if non-nil, zero value otherwise.

### GetIncludeIdentityOk

`func (o *WidgetRequest) GetIncludeIdentityOk() (*bool, bool)`

GetIncludeIdentityOk returns a tuple with the IncludeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIdentity

`func (o *WidgetRequest) SetIncludeIdentity(v bool)`

SetIncludeIdentity sets IncludeIdentity field to given value.

### HasIncludeIdentity

`func (o *WidgetRequest) HasIncludeIdentity() bool`

HasIncludeIdentity returns a boolean if a field has been set.

### GetIncludeTransactions

`func (o *WidgetRequest) GetIncludeTransactions() bool`

GetIncludeTransactions returns the IncludeTransactions field if non-nil, zero value otherwise.

### GetIncludeTransactionsOk

`func (o *WidgetRequest) GetIncludeTransactionsOk() (*bool, bool)`

GetIncludeTransactionsOk returns a tuple with the IncludeTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTransactions

`func (o *WidgetRequest) SetIncludeTransactions(v bool)`

SetIncludeTransactions sets IncludeTransactions field to given value.

### HasIncludeTransactions

`func (o *WidgetRequest) HasIncludeTransactions() bool`

HasIncludeTransactions returns a boolean if a field has been set.

### GetIsMobileWebview

`func (o *WidgetRequest) GetIsMobileWebview() bool`

GetIsMobileWebview returns the IsMobileWebview field if non-nil, zero value otherwise.

### GetIsMobileWebviewOk

`func (o *WidgetRequest) GetIsMobileWebviewOk() (*bool, bool)`

GetIsMobileWebviewOk returns a tuple with the IsMobileWebview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMobileWebview

`func (o *WidgetRequest) SetIsMobileWebview(v bool)`

SetIsMobileWebview sets IsMobileWebview field to given value.

### HasIsMobileWebview

`func (o *WidgetRequest) HasIsMobileWebview() bool`

HasIsMobileWebview returns a boolean if a field has been set.

### GetMode

`func (o *WidgetRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WidgetRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WidgetRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WidgetRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOauthReferralSource

`func (o *WidgetRequest) GetOauthReferralSource() string`

GetOauthReferralSource returns the OauthReferralSource field if non-nil, zero value otherwise.

### GetOauthReferralSourceOk

`func (o *WidgetRequest) GetOauthReferralSourceOk() (*string, bool)`

GetOauthReferralSourceOk returns a tuple with the OauthReferralSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthReferralSource

`func (o *WidgetRequest) SetOauthReferralSource(v string)`

SetOauthReferralSource sets OauthReferralSource field to given value.

### HasOauthReferralSource

`func (o *WidgetRequest) HasOauthReferralSource() bool`

HasOauthReferralSource returns a boolean if a field has been set.

### GetUiMessageVersion

`func (o *WidgetRequest) GetUiMessageVersion() int32`

GetUiMessageVersion returns the UiMessageVersion field if non-nil, zero value otherwise.

### GetUiMessageVersionOk

`func (o *WidgetRequest) GetUiMessageVersionOk() (*int32, bool)`

GetUiMessageVersionOk returns a tuple with the UiMessageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMessageVersion

`func (o *WidgetRequest) SetUiMessageVersion(v int32)`

SetUiMessageVersion sets UiMessageVersion field to given value.

### HasUiMessageVersion

`func (o *WidgetRequest) HasUiMessageVersion() bool`

HasUiMessageVersion returns a boolean if a field has been set.

### GetUiMessageWebviewUrlScheme

`func (o *WidgetRequest) GetUiMessageWebviewUrlScheme() string`

GetUiMessageWebviewUrlScheme returns the UiMessageWebviewUrlScheme field if non-nil, zero value otherwise.

### GetUiMessageWebviewUrlSchemeOk

`func (o *WidgetRequest) GetUiMessageWebviewUrlSchemeOk() (*string, bool)`

GetUiMessageWebviewUrlSchemeOk returns a tuple with the UiMessageWebviewUrlScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMessageWebviewUrlScheme

`func (o *WidgetRequest) SetUiMessageWebviewUrlScheme(v string)`

SetUiMessageWebviewUrlScheme sets UiMessageWebviewUrlScheme field to given value.

### HasUiMessageWebviewUrlScheme

`func (o *WidgetRequest) HasUiMessageWebviewUrlScheme() bool`

HasUiMessageWebviewUrlScheme returns a boolean if a field has been set.

### GetUpdateCredentials

`func (o *WidgetRequest) GetUpdateCredentials() bool`

GetUpdateCredentials returns the UpdateCredentials field if non-nil, zero value otherwise.

### GetUpdateCredentialsOk

`func (o *WidgetRequest) GetUpdateCredentialsOk() (*bool, bool)`

GetUpdateCredentialsOk returns a tuple with the UpdateCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCredentials

`func (o *WidgetRequest) SetUpdateCredentials(v bool)`

SetUpdateCredentials sets UpdateCredentials field to given value.

### HasUpdateCredentials

`func (o *WidgetRequest) HasUpdateCredentials() bool`

HasUpdateCredentials returns a boolean if a field has been set.

### GetWidgetType

`func (o *WidgetRequest) GetWidgetType() string`

GetWidgetType returns the WidgetType field if non-nil, zero value otherwise.

### GetWidgetTypeOk

`func (o *WidgetRequest) GetWidgetTypeOk() (*string, bool)`

GetWidgetTypeOk returns a tuple with the WidgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetType

`func (o *WidgetRequest) SetWidgetType(v string)`

SetWidgetType sets WidgetType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


