# ConnectWidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientRedirectUrl** | Pointer to **string** |  | [optional] 
**ColorScheme** | Pointer to **string** |  | [optional] 
**CurrentInstitutionCode** | Pointer to **string** |  | [optional] 
**CurrentMemberGuid** | Pointer to **string** |  | [optional] 
**DisableInstitutionSearch** | Pointer to **bool** |  | [optional] 
**IncludeTransactions** | Pointer to **bool** |  | [optional] 
**IsMobileWebview** | Pointer to **bool** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**UiMessageVersion** | Pointer to **int32** |  | [optional] 
**UiMessageWebviewUrlScheme** | Pointer to **string** |  | [optional] 
**UpdateCredentials** | Pointer to **bool** |  | [optional] 

## Methods

### NewConnectWidgetRequest

`func NewConnectWidgetRequest() *ConnectWidgetRequest`

NewConnectWidgetRequest instantiates a new ConnectWidgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectWidgetRequestWithDefaults

`func NewConnectWidgetRequestWithDefaults() *ConnectWidgetRequest`

NewConnectWidgetRequestWithDefaults instantiates a new ConnectWidgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientRedirectUrl

`func (o *ConnectWidgetRequest) GetClientRedirectUrl() string`

GetClientRedirectUrl returns the ClientRedirectUrl field if non-nil, zero value otherwise.

### GetClientRedirectUrlOk

`func (o *ConnectWidgetRequest) GetClientRedirectUrlOk() (*string, bool)`

GetClientRedirectUrlOk returns a tuple with the ClientRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRedirectUrl

`func (o *ConnectWidgetRequest) SetClientRedirectUrl(v string)`

SetClientRedirectUrl sets ClientRedirectUrl field to given value.

### HasClientRedirectUrl

`func (o *ConnectWidgetRequest) HasClientRedirectUrl() bool`

HasClientRedirectUrl returns a boolean if a field has been set.

### GetColorScheme

`func (o *ConnectWidgetRequest) GetColorScheme() string`

GetColorScheme returns the ColorScheme field if non-nil, zero value otherwise.

### GetColorSchemeOk

`func (o *ConnectWidgetRequest) GetColorSchemeOk() (*string, bool)`

GetColorSchemeOk returns a tuple with the ColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorScheme

`func (o *ConnectWidgetRequest) SetColorScheme(v string)`

SetColorScheme sets ColorScheme field to given value.

### HasColorScheme

`func (o *ConnectWidgetRequest) HasColorScheme() bool`

HasColorScheme returns a boolean if a field has been set.

### GetCurrentInstitutionCode

`func (o *ConnectWidgetRequest) GetCurrentInstitutionCode() string`

GetCurrentInstitutionCode returns the CurrentInstitutionCode field if non-nil, zero value otherwise.

### GetCurrentInstitutionCodeOk

`func (o *ConnectWidgetRequest) GetCurrentInstitutionCodeOk() (*string, bool)`

GetCurrentInstitutionCodeOk returns a tuple with the CurrentInstitutionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentInstitutionCode

`func (o *ConnectWidgetRequest) SetCurrentInstitutionCode(v string)`

SetCurrentInstitutionCode sets CurrentInstitutionCode field to given value.

### HasCurrentInstitutionCode

`func (o *ConnectWidgetRequest) HasCurrentInstitutionCode() bool`

HasCurrentInstitutionCode returns a boolean if a field has been set.

### GetCurrentMemberGuid

`func (o *ConnectWidgetRequest) GetCurrentMemberGuid() string`

GetCurrentMemberGuid returns the CurrentMemberGuid field if non-nil, zero value otherwise.

### GetCurrentMemberGuidOk

`func (o *ConnectWidgetRequest) GetCurrentMemberGuidOk() (*string, bool)`

GetCurrentMemberGuidOk returns a tuple with the CurrentMemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMemberGuid

`func (o *ConnectWidgetRequest) SetCurrentMemberGuid(v string)`

SetCurrentMemberGuid sets CurrentMemberGuid field to given value.

### HasCurrentMemberGuid

`func (o *ConnectWidgetRequest) HasCurrentMemberGuid() bool`

HasCurrentMemberGuid returns a boolean if a field has been set.

### GetDisableInstitutionSearch

`func (o *ConnectWidgetRequest) GetDisableInstitutionSearch() bool`

GetDisableInstitutionSearch returns the DisableInstitutionSearch field if non-nil, zero value otherwise.

### GetDisableInstitutionSearchOk

`func (o *ConnectWidgetRequest) GetDisableInstitutionSearchOk() (*bool, bool)`

GetDisableInstitutionSearchOk returns a tuple with the DisableInstitutionSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableInstitutionSearch

`func (o *ConnectWidgetRequest) SetDisableInstitutionSearch(v bool)`

SetDisableInstitutionSearch sets DisableInstitutionSearch field to given value.

### HasDisableInstitutionSearch

`func (o *ConnectWidgetRequest) HasDisableInstitutionSearch() bool`

HasDisableInstitutionSearch returns a boolean if a field has been set.

### GetIncludeTransactions

`func (o *ConnectWidgetRequest) GetIncludeTransactions() bool`

GetIncludeTransactions returns the IncludeTransactions field if non-nil, zero value otherwise.

### GetIncludeTransactionsOk

`func (o *ConnectWidgetRequest) GetIncludeTransactionsOk() (*bool, bool)`

GetIncludeTransactionsOk returns a tuple with the IncludeTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTransactions

`func (o *ConnectWidgetRequest) SetIncludeTransactions(v bool)`

SetIncludeTransactions sets IncludeTransactions field to given value.

### HasIncludeTransactions

`func (o *ConnectWidgetRequest) HasIncludeTransactions() bool`

HasIncludeTransactions returns a boolean if a field has been set.

### GetIsMobileWebview

`func (o *ConnectWidgetRequest) GetIsMobileWebview() bool`

GetIsMobileWebview returns the IsMobileWebview field if non-nil, zero value otherwise.

### GetIsMobileWebviewOk

`func (o *ConnectWidgetRequest) GetIsMobileWebviewOk() (*bool, bool)`

GetIsMobileWebviewOk returns a tuple with the IsMobileWebview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMobileWebview

`func (o *ConnectWidgetRequest) SetIsMobileWebview(v bool)`

SetIsMobileWebview sets IsMobileWebview field to given value.

### HasIsMobileWebview

`func (o *ConnectWidgetRequest) HasIsMobileWebview() bool`

HasIsMobileWebview returns a boolean if a field has been set.

### GetMode

`func (o *ConnectWidgetRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ConnectWidgetRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ConnectWidgetRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ConnectWidgetRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetUiMessageVersion

`func (o *ConnectWidgetRequest) GetUiMessageVersion() int32`

GetUiMessageVersion returns the UiMessageVersion field if non-nil, zero value otherwise.

### GetUiMessageVersionOk

`func (o *ConnectWidgetRequest) GetUiMessageVersionOk() (*int32, bool)`

GetUiMessageVersionOk returns a tuple with the UiMessageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMessageVersion

`func (o *ConnectWidgetRequest) SetUiMessageVersion(v int32)`

SetUiMessageVersion sets UiMessageVersion field to given value.

### HasUiMessageVersion

`func (o *ConnectWidgetRequest) HasUiMessageVersion() bool`

HasUiMessageVersion returns a boolean if a field has been set.

### GetUiMessageWebviewUrlScheme

`func (o *ConnectWidgetRequest) GetUiMessageWebviewUrlScheme() string`

GetUiMessageWebviewUrlScheme returns the UiMessageWebviewUrlScheme field if non-nil, zero value otherwise.

### GetUiMessageWebviewUrlSchemeOk

`func (o *ConnectWidgetRequest) GetUiMessageWebviewUrlSchemeOk() (*string, bool)`

GetUiMessageWebviewUrlSchemeOk returns a tuple with the UiMessageWebviewUrlScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMessageWebviewUrlScheme

`func (o *ConnectWidgetRequest) SetUiMessageWebviewUrlScheme(v string)`

SetUiMessageWebviewUrlScheme sets UiMessageWebviewUrlScheme field to given value.

### HasUiMessageWebviewUrlScheme

`func (o *ConnectWidgetRequest) HasUiMessageWebviewUrlScheme() bool`

HasUiMessageWebviewUrlScheme returns a boolean if a field has been set.

### GetUpdateCredentials

`func (o *ConnectWidgetRequest) GetUpdateCredentials() bool`

GetUpdateCredentials returns the UpdateCredentials field if non-nil, zero value otherwise.

### GetUpdateCredentialsOk

`func (o *ConnectWidgetRequest) GetUpdateCredentialsOk() (*bool, bool)`

GetUpdateCredentialsOk returns a tuple with the UpdateCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCredentials

`func (o *ConnectWidgetRequest) SetUpdateCredentials(v bool)`

SetUpdateCredentials sets UpdateCredentials field to given value.

### HasUpdateCredentials

`func (o *ConnectWidgetRequest) HasUpdateCredentials() bool`

HasUpdateCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


