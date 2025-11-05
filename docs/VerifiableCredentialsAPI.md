# \VerifiableCredentialsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountsData**](VerifiableCredentialsAPI.md#GetAccountsData) | **Get** /vc/users/{user_guid}/members/{member_guid}/accounts | Get Accounts Data
[**GetIdentityData**](VerifiableCredentialsAPI.md#GetIdentityData) | **Get** /vc/users/{user_guid}/members/{member_guid}/customers | Get Identity Data
[**GetTransactionsData**](VerifiableCredentialsAPI.md#GetTransactionsData) | **Get** /vc/users/{user_guid}/accounts/{account_guid}/transactions | Get Transactions Data



## GetAccountsData

> VCResponse GetAccountsData(ctx, userGuid, memberGuid).Execute()

Get Accounts Data



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
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerifiableCredentialsAPI.GetAccountsData(context.Background(), userGuid, memberGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerifiableCredentialsAPI.GetAccountsData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountsData`: VCResponse
    fmt.Fprintf(os.Stdout, "Response from `VerifiableCredentialsAPI.GetAccountsData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VCResponse**](VCResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v2beta+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityData

> VCResponse GetIdentityData(ctx, userGuid, memberGuid).Execute()

Get Identity Data



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
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerifiableCredentialsAPI.GetIdentityData(context.Background(), userGuid, memberGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerifiableCredentialsAPI.GetIdentityData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityData`: VCResponse
    fmt.Fprintf(os.Stdout, "Response from `VerifiableCredentialsAPI.GetIdentityData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VCResponse**](VCResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v2beta+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsData

> VCResponse GetTransactionsData(ctx, userGuid, accountGuid).StartTime(startTime).EndTime(endTime).Execute()

Get Transactions Data



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
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    startTime := "2015-09-20" // string | Filter transactions from this date. Must be in the format YYYY-MM-DD. The range is limited to 1 year. (optional)
    endTime := "2015-09-20" // string | Filter transactions to this date. Must be in the format YYYY-MM-DD. The range is limited to 1 year. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerifiableCredentialsAPI.GetTransactionsData(context.Background(), userGuid, accountGuid).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerifiableCredentialsAPI.GetTransactionsData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionsData`: VCResponse
    fmt.Fprintf(os.Stdout, "Response from `VerifiableCredentialsAPI.GetTransactionsData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **string** | Filter transactions from this date. Must be in the format YYYY-MM-DD. The range is limited to 1 year. | 
 **endTime** | **string** | Filter transactions to this date. Must be in the format YYYY-MM-DD. The range is limited to 1 year. | 

### Return type

[**VCResponse**](VCResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v2beta+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

