# \MicrodepositsAPI

All URIs are relative to *https://api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MicroDepositsMicrodepositGuidVerifyPut**](MicrodepositsAPI.md#MicroDepositsMicrodepositGuidVerifyPut) | **Put** /micro_deposits/{microdeposit_guid}/verify | Verify a Microdeposit
[**UsersUserGuidMicroDepositsGet**](MicrodepositsAPI.md#UsersUserGuidMicroDepositsGet) | **Get** /users/{user_guid}/micro_deposits | List all microdeposits for a user
[**UsersUserGuidMicroDepositsMicroDepositGuidDelete**](MicrodepositsAPI.md#UsersUserGuidMicroDepositsMicroDepositGuidDelete) | **Delete** /users/{user_guid}/micro_deposits/{micro_deposit_guid} | Delete a microdeposit
[**UsersUserGuidMicroDepositsMicroDepositGuidGet**](MicrodepositsAPI.md#UsersUserGuidMicroDepositsMicroDepositGuidGet) | **Get** /users/{user_guid}/micro_deposits/{micro_deposit_guid} | Read a microdeposit for a user
[**UsersUserGuidMicroDepositsPost**](MicrodepositsAPI.md#UsersUserGuidMicroDepositsPost) | **Post** /users/{user_guid}/micro_deposits | Create a microdeposit



## MicroDepositsMicrodepositGuidVerifyPut

> MicrodepositResponseBody MicroDepositsMicrodepositGuidVerifyPut(ctx, microdepositGuid).MicrodepositVerifyRequestBody(microdepositVerifyRequestBody).Execute()

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
    microdepositGuid := "microdepositGuid_example" // string | The unique identifier for the microdeposit. Defined by MX.
    microdepositVerifyRequestBody := *openapiclient.NewMicrodepositVerifyRequestBody() // MicrodepositVerifyRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicrodepositsAPI.MicroDepositsMicrodepositGuidVerifyPut(context.Background(), microdepositGuid).MicrodepositVerifyRequestBody(microdepositVerifyRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicrodepositsAPI.MicroDepositsMicrodepositGuidVerifyPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MicroDepositsMicrodepositGuidVerifyPut`: MicrodepositResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MicrodepositsAPI.MicroDepositsMicrodepositGuidVerifyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microdepositGuid** | **string** | The unique identifier for the microdeposit. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMicroDepositsMicrodepositGuidVerifyPutRequest struct via the builder pattern


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


## UsersUserGuidMicroDepositsGet

> MicrodepositsResponseBody UsersUserGuidMicroDepositsGet(ctx, userGuid).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for the user. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicrodepositsAPI.UsersUserGuidMicroDepositsGet(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicrodepositsAPI.UsersUserGuidMicroDepositsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidMicroDepositsGet`: MicrodepositsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MicrodepositsAPI.UsersUserGuidMicroDepositsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the user. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidMicroDepositsGetRequest struct via the builder pattern


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


## UsersUserGuidMicroDepositsMicroDepositGuidDelete

> UsersUserGuidMicroDepositsMicroDepositGuidDelete(ctx, microDepositGuid, userGuid).Execute()

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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MicrodepositsAPI.UsersUserGuidMicroDepositsMicroDepositGuidDelete(context.Background(), microDepositGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicrodepositsAPI.UsersUserGuidMicroDepositsMicroDepositGuidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microDepositGuid** | **string** | The unique identifier for the microdeposit. Defined by MX. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidMicroDepositsMicroDepositGuidDeleteRequest struct via the builder pattern


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


## UsersUserGuidMicroDepositsMicroDepositGuidGet

> MicrodepositResponseBody UsersUserGuidMicroDepositsMicroDepositGuidGet(ctx, userGuid, microDepositGuid).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for the user. Defined by MX.
    microDepositGuid := "microDepositGuid_example" // string | The unique identifier for the microdeposit. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicrodepositsAPI.UsersUserGuidMicroDepositsMicroDepositGuidGet(context.Background(), userGuid, microDepositGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicrodepositsAPI.UsersUserGuidMicroDepositsMicroDepositGuidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidMicroDepositsMicroDepositGuidGet`: MicrodepositResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MicrodepositsAPI.UsersUserGuidMicroDepositsMicroDepositGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the user. Defined by MX. | 
**microDepositGuid** | **string** | The unique identifier for the microdeposit. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidMicroDepositsMicroDepositGuidGetRequest struct via the builder pattern


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


## UsersUserGuidMicroDepositsPost

> MicrodepositResponseBody UsersUserGuidMicroDepositsPost(ctx, userGuid).MicrodepositRequestBody(microdepositRequestBody).Execute()

Create a microdeposit



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
    userGuid := "userGuid_example" // string | The unique identifier for the user. Defined by MX.
    microdepositRequestBody := *openapiclient.NewMicrodepositRequestBody() // MicrodepositRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MicrodepositsAPI.UsersUserGuidMicroDepositsPost(context.Background(), userGuid).MicrodepositRequestBody(microdepositRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MicrodepositsAPI.UsersUserGuidMicroDepositsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidMicroDepositsPost`: MicrodepositResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MicrodepositsAPI.UsersUserGuidMicroDepositsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the user. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidMicroDepositsPostRequest struct via the builder pattern


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

