# \ProcessorTokenAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckRealTimeAccountBalance**](ProcessorTokenAPI.md#CheckRealTimeAccountBalance) | **Post** /account/check_balance | Check Real Time Account Balance (Processors Only)
[**DeprecatedRequestPaymentProcessorAuthorizationCode**](ProcessorTokenAPI.md#DeprecatedRequestPaymentProcessorAuthorizationCode) | **Post** /payment_processor_authorization_code | (Deprecated) Request an authorization code
[**GetAccountOwnerInfo**](ProcessorTokenAPI.md#GetAccountOwnerInfo) | **Get** /account/transactions | Get account owner information (Processors Only)
[**ListTokens**](ProcessorTokenAPI.md#ListTokens) | **Get** /tokens | View a List of Tokens
[**ReadAccountBalance**](ProcessorTokenAPI.md#ReadAccountBalance) | **Get** /payment_account | Read the account balance (Processors Only)
[**RequestAccountNumber**](ProcessorTokenAPI.md#RequestAccountNumber) | **Get** /account/account_numbers | Request an account number (Processors Only)
[**RequestAuthorizationCode**](ProcessorTokenAPI.md#RequestAuthorizationCode) | **Post** /authorization_code | Request an authorization code



## CheckRealTimeAccountBalance

> MemberResponseBody CheckRealTimeAccountBalance(ctx).Execute()

Check Real Time Account Balance (Processors Only)



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProcessorTokenAPI.CheckRealTimeAccountBalance(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorTokenAPI.CheckRealTimeAccountBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckRealTimeAccountBalance`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `ProcessorTokenAPI.CheckRealTimeAccountBalance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckRealTimeAccountBalanceRequest struct via the builder pattern


### Return type

[**MemberResponseBody**](MemberResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeprecatedRequestPaymentProcessorAuthorizationCode

> PaymentProcessorAuthorizationCodeResponseBody DeprecatedRequestPaymentProcessorAuthorizationCode(ctx).PaymentProcessorAuthorizationCodeRequestBody(paymentProcessorAuthorizationCodeRequestBody).Execute()

(Deprecated) Request an authorization code



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
    paymentProcessorAuthorizationCodeRequestBody := *openapiclient.NewPaymentProcessorAuthorizationCodeRequestBody() // PaymentProcessorAuthorizationCodeRequestBody | The scope for the authorization code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProcessorTokenAPI.DeprecatedRequestPaymentProcessorAuthorizationCode(context.Background()).PaymentProcessorAuthorizationCodeRequestBody(paymentProcessorAuthorizationCodeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorTokenAPI.DeprecatedRequestPaymentProcessorAuthorizationCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedRequestPaymentProcessorAuthorizationCode`: PaymentProcessorAuthorizationCodeResponseBody
    fmt.Fprintf(os.Stdout, "Response from `ProcessorTokenAPI.DeprecatedRequestPaymentProcessorAuthorizationCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedRequestPaymentProcessorAuthorizationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentProcessorAuthorizationCodeRequestBody** | [**PaymentProcessorAuthorizationCodeRequestBody**](PaymentProcessorAuthorizationCodeRequestBody.md) | The scope for the authorization code. | 

### Return type

[**PaymentProcessorAuthorizationCodeResponseBody**](PaymentProcessorAuthorizationCodeResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountOwnerInfo

> ProcessorOwnerBody GetAccountOwnerInfo(ctx).Execute()

Get account owner information (Processors Only)



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProcessorTokenAPI.GetAccountOwnerInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorTokenAPI.GetAccountOwnerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountOwnerInfo`: ProcessorOwnerBody
    fmt.Fprintf(os.Stdout, "Response from `ProcessorTokenAPI.GetAccountOwnerInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountOwnerInfoRequest struct via the builder pattern


### Return type

[**ProcessorOwnerBody**](ProcessorOwnerBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokens

> TokenResponseBody ListTokens(ctx).TokenRequestBody(tokenRequestBody).Execute()

View a List of Tokens



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
    tokenRequestBody := *openapiclient.NewTokenRequestBody() // TokenRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProcessorTokenAPI.ListTokens(context.Background()).TokenRequestBody(tokenRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorTokenAPI.ListTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokens`: TokenResponseBody
    fmt.Fprintf(os.Stdout, "Response from `ProcessorTokenAPI.ListTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequestBody** | [**TokenRequestBody**](TokenRequestBody.md) |  | 

### Return type

[**TokenResponseBody**](TokenResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccountBalance

> PaymentAccountBody ReadAccountBalance(ctx).Execute()

Read the account balance (Processors Only)



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProcessorTokenAPI.ReadAccountBalance(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorTokenAPI.ReadAccountBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAccountBalance`: PaymentAccountBody
    fmt.Fprintf(os.Stdout, "Response from `ProcessorTokenAPI.ReadAccountBalance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadAccountBalanceRequest struct via the builder pattern


### Return type

[**PaymentAccountBody**](PaymentAccountBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestAccountNumber

> ProcessorAccountNumberBody RequestAccountNumber(ctx).Execute()

Request an account number (Processors Only)



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProcessorTokenAPI.RequestAccountNumber(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorTokenAPI.RequestAccountNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestAccountNumber`: ProcessorAccountNumberBody
    fmt.Fprintf(os.Stdout, "Response from `ProcessorTokenAPI.RequestAccountNumber`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRequestAccountNumberRequest struct via the builder pattern


### Return type

[**ProcessorAccountNumberBody**](ProcessorAccountNumberBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestAuthorizationCode

> AuthorizationCodeResponseBody RequestAuthorizationCode(ctx).AuthorizationCodeRequestBody(authorizationCodeRequestBody).Execute()

Request an authorization code



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
    authorizationCodeRequestBody := *openapiclient.NewAuthorizationCodeRequestBody() // AuthorizationCodeRequestBody | The scope for the authorization code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProcessorTokenAPI.RequestAuthorizationCode(context.Background()).AuthorizationCodeRequestBody(authorizationCodeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorTokenAPI.RequestAuthorizationCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestAuthorizationCode`: AuthorizationCodeResponseBody
    fmt.Fprintf(os.Stdout, "Response from `ProcessorTokenAPI.RequestAuthorizationCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestAuthorizationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationCodeRequestBody** | [**AuthorizationCodeRequestBody**](AuthorizationCodeRequestBody.md) | The scope for the authorization code. | 

### Return type

[**AuthorizationCodeResponseBody**](AuthorizationCodeResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

