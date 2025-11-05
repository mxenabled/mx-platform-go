# \WidgetsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestConnectWidgetURL**](WidgetsAPI.md#RequestConnectWidgetURL) | **Post** /users/{user_guid}/connect_widget_url | (Deprecated) Request connect widget URL
[**RequestOAuthWindowURI**](WidgetsAPI.md#RequestOAuthWindowURI) | **Get** /users/{user_guid}/members/{member_guid}/oauth_window_uri | Request oauth window uri
[**RequestWidgetURL**](WidgetsAPI.md#RequestWidgetURL) | **Post** /users/{user_guid}/widget_urls | Request widget URL



## RequestConnectWidgetURL

> ConnectWidgetResponseBody RequestConnectWidgetURL(ctx, userGuid).ConnectWidgetRequestBody(connectWidgetRequestBody).Execute()

(Deprecated) Request connect widget URL



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    connectWidgetRequestBody := *openapiclient.NewConnectWidgetRequestBody() // ConnectWidgetRequestBody | Optional config options for WebView (is_mobile_webview, current_institution_code, current_member_guid, update_credentials)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WidgetsAPI.RequestConnectWidgetURL(context.Background(), userGuid).ConnectWidgetRequestBody(connectWidgetRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WidgetsAPI.RequestConnectWidgetURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestConnectWidgetURL`: ConnectWidgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `WidgetsAPI.RequestConnectWidgetURL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestConnectWidgetURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectWidgetRequestBody** | [**ConnectWidgetRequestBody**](ConnectWidgetRequestBody.md) | Optional config options for WebView (is_mobile_webview, current_institution_code, current_member_guid, update_credentials) | 

### Return type

[**ConnectWidgetResponseBody**](ConnectWidgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestOAuthWindowURI

> OAuthWindowResponseBody RequestOAuthWindowURI(ctx, memberGuid, userGuid).ClientRedirectUrl(clientRedirectUrl).EnableApp2app(enableApp2app).ReferralSource(referralSource).SkipAggregation(skipAggregation).UiMessageWebviewUrlScheme(uiMessageWebviewUrlScheme).Execute()

Request oauth window uri



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    clientRedirectUrl := "https://{yoursite.com}" // string | A URL that MX will redirect to at the end of OAuth with additional query parameters. Only available with `referral_source=APP`. (optional)
    enableApp2app := "false" // string | This indicates whether OAuth app2app behavior is enabled for institutions that support it. Defaults to `true`. When set to `false`, any `oauth_window_uri` generated will **not** direct the end user to the institution's mobile application. This setting is not persistent. This setting currently only affects Chase institutions. (optional)
    referralSource := "APP" // string | Must be either `BROWSER` or `APP` depending on the implementation. Defaults to `BROWSER`. (optional)
    skipAggregation := false // bool | Setting this parameter to `true` will prevent the member from automatically aggregating after being redirected from the authorization page. (optional)
    uiMessageWebviewUrlScheme := "uiMessageWebviewUrlScheme_example" // string | A scheme for routing the user back to the application state they were previously in. Only available with `referral_source=APP`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WidgetsAPI.RequestOAuthWindowURI(context.Background(), memberGuid, userGuid).ClientRedirectUrl(clientRedirectUrl).EnableApp2app(enableApp2app).ReferralSource(referralSource).SkipAggregation(skipAggregation).UiMessageWebviewUrlScheme(uiMessageWebviewUrlScheme).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WidgetsAPI.RequestOAuthWindowURI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestOAuthWindowURI`: OAuthWindowResponseBody
    fmt.Fprintf(os.Stdout, "Response from `WidgetsAPI.RequestOAuthWindowURI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestOAuthWindowURIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRedirectUrl** | **string** | A URL that MX will redirect to at the end of OAuth with additional query parameters. Only available with &#x60;referral_source&#x3D;APP&#x60;. | 
 **enableApp2app** | **string** | This indicates whether OAuth app2app behavior is enabled for institutions that support it. Defaults to &#x60;true&#x60;. When set to &#x60;false&#x60;, any &#x60;oauth_window_uri&#x60; generated will **not** direct the end user to the institution&#39;s mobile application. This setting is not persistent. This setting currently only affects Chase institutions. | 
 **referralSource** | **string** | Must be either &#x60;BROWSER&#x60; or &#x60;APP&#x60; depending on the implementation. Defaults to &#x60;BROWSER&#x60;. | 
 **skipAggregation** | **bool** | Setting this parameter to &#x60;true&#x60; will prevent the member from automatically aggregating after being redirected from the authorization page. | 
 **uiMessageWebviewUrlScheme** | **string** | A scheme for routing the user back to the application state they were previously in. Only available with &#x60;referral_source&#x3D;APP&#x60;. | 

### Return type

[**OAuthWindowResponseBody**](OAuthWindowResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestWidgetURL

> WidgetResponseBody RequestWidgetURL(ctx, userGuid).WidgetRequestBody(widgetRequestBody).AcceptLanguage(acceptLanguage).XCALLBACKPAYLOAD(xCALLBACKPAYLOAD).Execute()

Request widget URL



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    widgetRequestBody := *openapiclient.NewWidgetRequestBody() // WidgetRequestBody | The widget url configuration options.
    acceptLanguage := "en-US" // string | The desired language of the widget. (optional)
    xCALLBACKPAYLOAD := "813e50bd-4a7e-4517-b6bb-9eef65a68cbd" // string | The base64 encoded string defined in this header will be returned in the [Member](/resources/webhooks/member/) and [Member Data Updated](/resources/webhooks/member#member-data-updated) webhooks. This allows you to trace user interactions and workflows initiated externally and internally in the MX Platform. Max 1024 characters. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WidgetsAPI.RequestWidgetURL(context.Background(), userGuid).WidgetRequestBody(widgetRequestBody).AcceptLanguage(acceptLanguage).XCALLBACKPAYLOAD(xCALLBACKPAYLOAD).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WidgetsAPI.RequestWidgetURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestWidgetURL`: WidgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `WidgetsAPI.RequestWidgetURL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestWidgetURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **widgetRequestBody** | [**WidgetRequestBody**](WidgetRequestBody.md) | The widget url configuration options. | 
 **acceptLanguage** | **string** | The desired language of the widget. | 
 **xCALLBACKPAYLOAD** | **string** | The base64 encoded string defined in this header will be returned in the [Member](/resources/webhooks/member/) and [Member Data Updated](/resources/webhooks/member#member-data-updated) webhooks. This allows you to trace user interactions and workflows initiated externally and internally in the MX Platform. Max 1024 characters. | 

### Return type

[**WidgetResponseBody**](WidgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

