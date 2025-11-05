# \MicrodepositsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMicrodeposit**](MicrodepositsAPI.md#CreateMicrodeposit) | **Post** /users/{user_guid}/micro_deposits | Create or pre-initiate a microdeposit
[**DeleteMicrodeposit**](MicrodepositsAPI.md#DeleteMicrodeposit) | **Delete** /users/{user_guid}/micro_deposits/{micro_deposit_guid} | Delete a microdeposit
[**ListUserMicrodeposits**](MicrodepositsAPI.md#ListUserMicrodeposits) | **Get** /users/{user_guid}/micro_deposits | List all microdeposits for a user
[**ListUserVerifications**](MicrodepositsAPI.md#ListUserVerifications) | **Get** /users/{user_guid}/account_verifications | List all verifications for a user
[**ReadUserMicrodeposit**](MicrodepositsAPI.md#ReadUserMicrodeposit) | **Get** /users/{user_guid}/micro_deposits/{micro_deposit_guid} | Read a microdeposit for a user
[**VerifyMicrodeposit**](MicrodepositsAPI.md#VerifyMicrodeposit) | **Put** /micro_deposits/{micro_deposit_guid}/verify | Verify a Microdeposit



## CreateMicrodeposit

> MicrodepositResponseBody CreateMicrodeposit(ctx, userGuid).MicrodepositRequestBody(microdepositRequestBody).Execute()

Create or pre-initiate a microdeposit



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
    microdepositRequestBody := *openapiclient.NewMicrodepositRequestBody() // MicrodepositRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicrodepositsAPI.CreateMicrodeposit(context.Background(), userGuid).MicrodepositRequestBody(microdepositRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicrodepositsAPI.CreateMicrodeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMicrodeposit`: MicrodepositResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MicrodepositsAPI.CreateMicrodeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMicrodepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microdepositRequestBody** | [**MicrodepositRequestBody**](MicrodepositRequestBody.md) |  | 

### Return type

[**MicrodepositResponseBody**](MicrodepositResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMicrodeposit

> DeleteMicrodeposit(ctx, microDepositGuid, userGuid).Execute()

Delete a microdeposit



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
    microDepositGuid := "MIC-09ba578e-8448-4f7f-89e1-b62ff2517edb" // string | The unique identifier for the microdeposit. Defined by MX.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MicrodepositsAPI.DeleteMicrodeposit(context.Background(), microDepositGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicrodepositsAPI.DeleteMicrodeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microDepositGuid** | **string** | The unique identifier for the microdeposit. Defined by MX. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMicrodepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserMicrodeposits

> MicrodepositsResponseBody ListUserMicrodeposits(ctx, userGuid).Execute()

List all microdeposits for a user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicrodepositsAPI.ListUserMicrodeposits(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicrodepositsAPI.ListUserMicrodeposits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserMicrodeposits`: MicrodepositsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MicrodepositsAPI.ListUserMicrodeposits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserMicrodepositsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MicrodepositsResponseBody**](MicrodepositsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserVerifications

> MicrodepositResponseBody ListUserVerifications(ctx, userGuid).Execute()

List all verifications for a user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicrodepositsAPI.ListUserVerifications(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicrodepositsAPI.ListUserVerifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserVerifications`: MicrodepositResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MicrodepositsAPI.ListUserVerifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MicrodepositResponseBody**](MicrodepositResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserMicrodeposit

> MicrodepositResponseBody ReadUserMicrodeposit(ctx, microDepositGuid, userGuid).Execute()

Read a microdeposit for a user



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
    microDepositGuid := "MIC-09ba578e-8448-4f7f-89e1-b62ff2517edb" // string | The unique identifier for the microdeposit. Defined by MX.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicrodepositsAPI.ReadUserMicrodeposit(context.Background(), microDepositGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicrodepositsAPI.ReadUserMicrodeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserMicrodeposit`: MicrodepositResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MicrodepositsAPI.ReadUserMicrodeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microDepositGuid** | **string** | The unique identifier for the microdeposit. Defined by MX. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadUserMicrodepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MicrodepositResponseBody**](MicrodepositResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyMicrodeposit

> MicrodepositResponseBody VerifyMicrodeposit(ctx, microDepositGuid).MicrodepositVerifyRequestBody(microdepositVerifyRequestBody).Execute()

Verify a Microdeposit



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
    microDepositGuid := "MIC-09ba578e-8448-4f7f-89e1-b62ff2517edb" // string | The unique identifier for the microdeposit. Defined by MX.
    microdepositVerifyRequestBody := *openapiclient.NewMicrodepositVerifyRequestBody() // MicrodepositVerifyRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicrodepositsAPI.VerifyMicrodeposit(context.Background(), microDepositGuid).MicrodepositVerifyRequestBody(microdepositVerifyRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicrodepositsAPI.VerifyMicrodeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyMicrodeposit`: MicrodepositResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MicrodepositsAPI.VerifyMicrodeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microDepositGuid** | **string** | The unique identifier for the microdeposit. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyMicrodepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microdepositVerifyRequestBody** | [**MicrodepositVerifyRequestBody**](MicrodepositVerifyRequestBody.md) |  | 

### Return type

[**MicrodepositResponseBody**](MicrodepositResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

