# \MonthlyCashFlowProfileAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadMonthlyCashFlowProfile**](MonthlyCashFlowProfileAPI.md#ReadMonthlyCashFlowProfile) | **Get** /users/{user_guid}/monthly_cash_flow_profile | Read monthly cash flow profile
[**UpdateMonthlyCashFlowProfile**](MonthlyCashFlowProfileAPI.md#UpdateMonthlyCashFlowProfile) | **Put** /users/{user_guid}/monthly_cash_flow_profile | Update monthly cash flow profile



## ReadMonthlyCashFlowProfile

> MonthlyCashFlowResponseBody ReadMonthlyCashFlowProfile(ctx, userGuid).Execute()

Read monthly cash flow profile



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
    resp, r, err := apiClient.MonthlyCashFlowProfileAPI.ReadMonthlyCashFlowProfile(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonthlyCashFlowProfileAPI.ReadMonthlyCashFlowProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMonthlyCashFlowProfile`: MonthlyCashFlowResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MonthlyCashFlowProfileAPI.ReadMonthlyCashFlowProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadMonthlyCashFlowProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonthlyCashFlowResponseBody**](MonthlyCashFlowResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonthlyCashFlowProfile

> MonthlyCashFlowResponseBody UpdateMonthlyCashFlowProfile(ctx, userGuid).MonthlyCashFlowProfileRequestBody(monthlyCashFlowProfileRequestBody).Execute()

Update monthly cash flow profile



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
    monthlyCashFlowProfileRequestBody := *openapiclient.NewMonthlyCashFlowProfileRequestBody() // MonthlyCashFlowProfileRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonthlyCashFlowProfileAPI.UpdateMonthlyCashFlowProfile(context.Background(), userGuid).MonthlyCashFlowProfileRequestBody(monthlyCashFlowProfileRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonthlyCashFlowProfileAPI.UpdateMonthlyCashFlowProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonthlyCashFlowProfile`: MonthlyCashFlowResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MonthlyCashFlowProfileAPI.UpdateMonthlyCashFlowProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonthlyCashFlowProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monthlyCashFlowProfileRequestBody** | [**MonthlyCashFlowProfileRequestBody**](MonthlyCashFlowProfileRequestBody.md) |  | 

### Return type

[**MonthlyCashFlowResponseBody**](MonthlyCashFlowResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

