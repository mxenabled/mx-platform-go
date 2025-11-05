# MemberCreateRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientRedirectUrl** | Pointer to **string** |  | [optional] 
**EnableApp2app** | Pointer to **bool** | This indicates whether OAuth app2app behavior is enabled for institutions that support it. Defaults to &#x60;true&#x60;. When set to &#x60;false&#x60;, any &#x60;oauth_window_uri&#x60; generated will **not** direct the end user to the institution&#39;s mobile application. This setting is not persistent. This setting currently only affects Chase institutions.  | [optional] 
**Member** | Pointer to [**MemberCreateRequest**](MemberCreateRequest.md) |  | [optional] 
**ReferralSource** | Pointer to **string** |  | [optional] 
**UiMessageWebviewUrlScheme** | Pointer to **string** |  | [optional] 

## Methods

### NewMemberCreateRequestBody

`func NewMemberCreateRequestBody() *MemberCreateRequestBody`

NewMemberCreateRequestBody instantiates a new MemberCreateRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberCreateRequestBodyWithDefaults

`func NewMemberCreateRequestBodyWithDefaults() *MemberCreateRequestBody`

NewMemberCreateRequestBodyWithDefaults instantiates a new MemberCreateRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientRedirectUrl

`func (o *MemberCreateRequestBody) GetClientRedirectUrl() string`

GetClientRedirectUrl returns the ClientRedirectUrl field if non-nil, zero value otherwise.

### GetClientRedirectUrlOk

`func (o *MemberCreateRequestBody) GetClientRedirectUrlOk() (*string, bool)`

GetClientRedirectUrlOk returns a tuple with the ClientRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRedirectUrl

`func (o *MemberCreateRequestBody) SetClientRedirectUrl(v string)`

SetClientRedirectUrl sets ClientRedirectUrl field to given value.

### HasClientRedirectUrl

`func (o *MemberCreateRequestBody) HasClientRedirectUrl() bool`

HasClientRedirectUrl returns a boolean if a field has been set.

### GetEnableApp2app

`func (o *MemberCreateRequestBody) GetEnableApp2app() bool`

GetEnableApp2app returns the EnableApp2app field if non-nil, zero value otherwise.

### GetEnableApp2appOk

`func (o *MemberCreateRequestBody) GetEnableApp2appOk() (*bool, bool)`

GetEnableApp2appOk returns a tuple with the EnableApp2app field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableApp2app

`func (o *MemberCreateRequestBody) SetEnableApp2app(v bool)`

SetEnableApp2app sets EnableApp2app field to given value.

### HasEnableApp2app

`func (o *MemberCreateRequestBody) HasEnableApp2app() bool`

HasEnableApp2app returns a boolean if a field has been set.

### GetMember

`func (o *MemberCreateRequestBody) GetMember() MemberCreateRequest`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *MemberCreateRequestBody) GetMemberOk() (*MemberCreateRequest, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *MemberCreateRequestBody) SetMember(v MemberCreateRequest)`

SetMember sets Member field to given value.

### HasMember

`func (o *MemberCreateRequestBody) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetReferralSource

`func (o *MemberCreateRequestBody) GetReferralSource() string`

GetReferralSource returns the ReferralSource field if non-nil, zero value otherwise.

### GetReferralSourceOk

`func (o *MemberCreateRequestBody) GetReferralSourceOk() (*string, bool)`

GetReferralSourceOk returns a tuple with the ReferralSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralSource

`func (o *MemberCreateRequestBody) SetReferralSource(v string)`

SetReferralSource sets ReferralSource field to given value.

### HasReferralSource

`func (o *MemberCreateRequestBody) HasReferralSource() bool`

HasReferralSource returns a boolean if a field has been set.

### GetUiMessageWebviewUrlScheme

`func (o *MemberCreateRequestBody) GetUiMessageWebviewUrlScheme() string`

GetUiMessageWebviewUrlScheme returns the UiMessageWebviewUrlScheme field if non-nil, zero value otherwise.

### GetUiMessageWebviewUrlSchemeOk

`func (o *MemberCreateRequestBody) GetUiMessageWebviewUrlSchemeOk() (*string, bool)`

GetUiMessageWebviewUrlSchemeOk returns a tuple with the UiMessageWebviewUrlScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMessageWebviewUrlScheme

`func (o *MemberCreateRequestBody) SetUiMessageWebviewUrlScheme(v string)`

SetUiMessageWebviewUrlScheme sets UiMessageWebviewUrlScheme field to given value.

### HasUiMessageWebviewUrlScheme

`func (o *MemberCreateRequestBody) HasUiMessageWebviewUrlScheme() bool`

HasUiMessageWebviewUrlScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


