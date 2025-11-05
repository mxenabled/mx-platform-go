# WidgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientRedirectUrl** | Pointer to **string** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. This determines the redirect destination at the end of OAuth when used with &#x60;is_mobile_webview: true&#x60; or &#x60;oauth_referral_source: &#39;APP&#39;&#x60;.  | [optional] 
**ColorScheme** | Pointer to **string** | This option can be passed to any &#x60;widget_type&#x60; but will not affect [legacy PFM widgets](products/experience/pfm/legacy-widget-overviews/). Load the widget with the specified &#x60;color_scheme&#x60;; options are &#x60;light&#x60;, &#x60;browser&#x60; (respects user&#39;s browser setting), and &#x60;dark&#x60;. Defaults to &#x60;light&#x60;. | [optional] 
**ConnectionsUseCaseFilter** | Pointer to **bool** | To use this parameter, you must also set &#x60;use_cases&#x60; in the same request. If &#x60;connections_use_case_filter&#x60; is set to &#x60;true&#x60;, the Connections Widget will only show connections (members) with the &#x60;use_cases&#x60; you set in the same request. For some examples, see [Filter Connections](/products/experience/pfm/widget-overviews/connections-widget#example-1). | [optional] 
**CurrentInstitutionCode** | Pointer to **string** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. Load the widget into the credential view for the specified institution.  | [optional] 
**CurrentInstitutionGuid** | Pointer to **string** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. Load the widget into the credential view for the specified institution.  | [optional] 
**CurrentMemberGuid** | Pointer to **string** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. Load the widget into a specific member that contains an error or requires multifactor authentication. The widget will determine the best view to load based on the member&#39;s current state. &#x60;current_member_guid&#x60; takes precedence over &#x60;current_institution_code&#x60; and &#x60;current_institution_guid&#x60;.  | [optional] 
**DisableBackgroundAgg** | Pointer to **bool** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. This determines whether background aggregation is enabled or disabled for the member created by the Connect Widget. Defaults to &#x60;false&#x60; in &#x60;aggregation&#x60; mode and &#x60;true&#x60; in &#x60;verification&#x60; mode. A global default for all members can be set by reaching out to MX.  | [optional] 
**DisableInstitutionSearch** | Pointer to **bool** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. This determines whether the institution search is displayed within the Connect Widget. This option must be used with &#x60;current_institution_code&#x60;, &#x60;current_instituion_guid&#x60;, or &#x60;current_member_guid&#x60;. When set to &#x60;true&#x60;, the institution search feature will be disabled and end users will not be able to navigate to it. Defaults to &#x60;false&#x60;.  | [optional] 
**EnableApp2app** | Pointer to **bool** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. This indicates whether OAuth app2app behavior is enabled for institutions that support it. Defaults to &#x60;true&#x60;. When set to &#x60;false&#x60;, the widget will **not** direct the end user to the institution&#39;s mobile application. This setting is not persistent. This setting currently only affects Chase institutions.  | [optional] 
**IncludeIdentity** | Pointer to **bool** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. This determines whether an account owner identification (AOI, previously called identity verification) is run in addition to the process specified by the &#x60;mode&#x60;. Defaults to &#x60;false&#x60;. This can be set in either &#x60;aggregation&#x60; or &#x60;verification&#x60; mode. The AOI runs after the primary process is complete.  | [optional] 
**IncludeTransactions** | Pointer to **bool** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. This determines whether transaction data are retrieved. Defaults to &#x60;true&#x60; in aggregation mode and &#x60;false&#x60; in verification mode. This can be set in either &#x60;aggregation&#x60; or &#x60;verification&#x60; mode. This option does not affect future foreground or background aggregations.  | [optional] 
**InsightGuid** | Pointer to **NullableString** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;pulse_widget&#x60;. Set this to the insight guid you want to appear at the top of the insights feed.  | [optional] 
**IsoCountryCode** | Pointer to **[]string** | An array of strings that filters institutions in the widget by the specified country code. Acceptable codes include &#x60;US&#x60;, &#x60;CA&#x60;, and &#x60;MX&#x60; (Mexico).  | [optional] 
**IsMobileWebview** | Pointer to **bool** | This option is for all &#x60;widget_type&#x60;s. This configures the widget to render in a mobile WebView. JavaScript event postMessages are replaced with URL updates.  | [optional] 
**MicrowidgetInstanceId** | Pointer to **NullableString** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;micro_pulse_carousel_widget&#x60;. Set this to a unique value for each instance of the Micro Widget. This lets us collect unique data for each instance of the widget.  | [optional] 
**Mode** | Pointer to **string** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. &#x60;mode&#x60; is the most important option for the Connect Widget. This determines what kind of process Connect will run, which affects how you should set many other options. Defaults to &#x60;aggregation&#x60;. &#x60;aggregation&#x60; mode retrieves account and transaction data; in other words, this runs a standard aggregation. &#x60;verification&#x60; mode retrieves account numbers and routing/transit numbers; in other words, it runs an Instant Account Verification (IAV). By default, verification mode does not retrieve transaction data; this default can be modified with secondary options. By default, background aggregation is disabled for all members created in verification mode; this default can be modified with secondary options.  | [optional] 
**OauthReferralSource** | Pointer to **string** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. This determines how MX will respond to the result of an OAuth flow. When set to &#x60;APP&#x60;, MX will redirect to the URI specified in the &#x60;ui_message_webview_url_scheme&#x60;. When set to &#x60;BROWSER&#x60;, MX will send a postMessage but not redirect. If &#x60;is_mobile_webview&#x60; is &#x60;true&#x60;, this defaults to &#x60;APP&#x60;. If false, it defaults to &#x60;BROWSER&#x60;.  | [optional] 
**UiMessageVersion** | Pointer to **int32** | This option is for all &#x60;widget_type&#x60;s. This determines which version of postMessage events are triggered. Defaults to 4. All new implementations must use version 4. Prior versions are deprecated.  | [optional] 
**UiMessageWebviewUrlScheme** | Pointer to **string** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. This is a client-defined scheme used in OAuth redirects in WebViews; also used in URL updates when these replace postMessages in WebViews. Defaults to &#x60;mx&#x60;.  | [optional] 
**UpdateCredentials** | Pointer to **bool** | Only use this option if the &#x60;widget_type&#x60; is set to &#x60;connect_widget&#x60;. Load the widget into a view that allows them to update the current member. Optionally used with &#x60;current_member_guid&#x60;. This option should be used sparingly. The best practice is to use &#x60;current_member_guid&#x60; and let the widget resolve the issue.  | [optional] 
**UseCases** | Pointer to **[]string** | The use case that will be associated with any members created through the widget. Valid values are &#x60;PFM&#x60; and/or &#x60;MONEY_MOVEMENT&#x60;. This is **required** if you&#39;ve met with MX, opted in to using this field, and are requesting a widget with a &#x60;widget_type&#x60; of &#x60;connect_widget&#x60; or &#x60;connections_widget&#x60;. | [optional] 
**WidgetType** | **string** | This determines which widget URL you&#39;ll receive.  See [Widget Types](/api-reference/platform-api/reference/widget-types) for a list of potential values. Additional request parameters may only apply to some widget types.  | 

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

### GetConnectionsUseCaseFilter

`func (o *WidgetRequest) GetConnectionsUseCaseFilter() bool`

GetConnectionsUseCaseFilter returns the ConnectionsUseCaseFilter field if non-nil, zero value otherwise.

### GetConnectionsUseCaseFilterOk

`func (o *WidgetRequest) GetConnectionsUseCaseFilterOk() (*bool, bool)`

GetConnectionsUseCaseFilterOk returns a tuple with the ConnectionsUseCaseFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsUseCaseFilter

`func (o *WidgetRequest) SetConnectionsUseCaseFilter(v bool)`

SetConnectionsUseCaseFilter sets ConnectionsUseCaseFilter field to given value.

### HasConnectionsUseCaseFilter

`func (o *WidgetRequest) HasConnectionsUseCaseFilter() bool`

HasConnectionsUseCaseFilter returns a boolean if a field has been set.

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

### GetEnableApp2app

`func (o *WidgetRequest) GetEnableApp2app() bool`

GetEnableApp2app returns the EnableApp2app field if non-nil, zero value otherwise.

### GetEnableApp2appOk

`func (o *WidgetRequest) GetEnableApp2appOk() (*bool, bool)`

GetEnableApp2appOk returns a tuple with the EnableApp2app field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableApp2app

`func (o *WidgetRequest) SetEnableApp2app(v bool)`

SetEnableApp2app sets EnableApp2app field to given value.

### HasEnableApp2app

`func (o *WidgetRequest) HasEnableApp2app() bool`

HasEnableApp2app returns a boolean if a field has been set.

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

### GetInsightGuid

`func (o *WidgetRequest) GetInsightGuid() string`

GetInsightGuid returns the InsightGuid field if non-nil, zero value otherwise.

### GetInsightGuidOk

`func (o *WidgetRequest) GetInsightGuidOk() (*string, bool)`

GetInsightGuidOk returns a tuple with the InsightGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightGuid

`func (o *WidgetRequest) SetInsightGuid(v string)`

SetInsightGuid sets InsightGuid field to given value.

### HasInsightGuid

`func (o *WidgetRequest) HasInsightGuid() bool`

HasInsightGuid returns a boolean if a field has been set.

### SetInsightGuidNil

`func (o *WidgetRequest) SetInsightGuidNil(b bool)`

 SetInsightGuidNil sets the value for InsightGuid to be an explicit nil

### UnsetInsightGuid
`func (o *WidgetRequest) UnsetInsightGuid()`

UnsetInsightGuid ensures that no value is present for InsightGuid, not even an explicit nil
### GetIsoCountryCode

`func (o *WidgetRequest) GetIsoCountryCode() []string`

GetIsoCountryCode returns the IsoCountryCode field if non-nil, zero value otherwise.

### GetIsoCountryCodeOk

`func (o *WidgetRequest) GetIsoCountryCodeOk() (*[]string, bool)`

GetIsoCountryCodeOk returns a tuple with the IsoCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountryCode

`func (o *WidgetRequest) SetIsoCountryCode(v []string)`

SetIsoCountryCode sets IsoCountryCode field to given value.

### HasIsoCountryCode

`func (o *WidgetRequest) HasIsoCountryCode() bool`

HasIsoCountryCode returns a boolean if a field has been set.

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

### GetMicrowidgetInstanceId

`func (o *WidgetRequest) GetMicrowidgetInstanceId() string`

GetMicrowidgetInstanceId returns the MicrowidgetInstanceId field if non-nil, zero value otherwise.

### GetMicrowidgetInstanceIdOk

`func (o *WidgetRequest) GetMicrowidgetInstanceIdOk() (*string, bool)`

GetMicrowidgetInstanceIdOk returns a tuple with the MicrowidgetInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrowidgetInstanceId

`func (o *WidgetRequest) SetMicrowidgetInstanceId(v string)`

SetMicrowidgetInstanceId sets MicrowidgetInstanceId field to given value.

### HasMicrowidgetInstanceId

`func (o *WidgetRequest) HasMicrowidgetInstanceId() bool`

HasMicrowidgetInstanceId returns a boolean if a field has been set.

### SetMicrowidgetInstanceIdNil

`func (o *WidgetRequest) SetMicrowidgetInstanceIdNil(b bool)`

 SetMicrowidgetInstanceIdNil sets the value for MicrowidgetInstanceId to be an explicit nil

### UnsetMicrowidgetInstanceId
`func (o *WidgetRequest) UnsetMicrowidgetInstanceId()`

UnsetMicrowidgetInstanceId ensures that no value is present for MicrowidgetInstanceId, not even an explicit nil
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

### GetUseCases

`func (o *WidgetRequest) GetUseCases() []string`

GetUseCases returns the UseCases field if non-nil, zero value otherwise.

### GetUseCasesOk

`func (o *WidgetRequest) GetUseCasesOk() (*[]string, bool)`

GetUseCasesOk returns a tuple with the UseCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCases

`func (o *WidgetRequest) SetUseCases(v []string)`

SetUseCases sets UseCases field to given value.

### HasUseCases

`func (o *WidgetRequest) HasUseCases() bool`

HasUseCases returns a boolean if a field has been set.

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


