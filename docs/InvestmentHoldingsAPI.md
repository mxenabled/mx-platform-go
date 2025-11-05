# \InvestmentHoldingsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeactivateUser**](InvestmentHoldingsAPI.md#DeactivateUser) | **Get** /users/{user_guid}/investment_holdings_deactivate | Deactivate user from Investment Holdings
[**ListHoldings**](InvestmentHoldingsAPI.md#ListHoldings) | **Get** /users/{user_guid}/investment_holdings | List holdings by user
[**ListHoldingsByAccount**](InvestmentHoldingsAPI.md#ListHoldingsByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/investment_holdings | List holdings by account
[**ListHoldingsByMember**](InvestmentHoldingsAPI.md#ListHoldingsByMember) | **Get** /users/{user_guid}/members/{member_guid}/investment_holdings | List holdings by member
[**ReadHolding**](InvestmentHoldingsAPI.md#ReadHolding) | **Get** /users/{user_guid}/investment_holdings/{holding_guid} | Read holding



## DeactivateUser

> InvestmentHoldingsDeactivation DeactivateUser(ctx, userGuid).Execute()

Deactivate user from Investment Holdings



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
    resp, r, err := apiClient.InvestmentHoldingsAPI.DeactivateUser(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestmentHoldingsAPI.DeactivateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateUser`: InvestmentHoldingsDeactivation
    fmt.Fprintf(os.Stdout, "Response from `InvestmentHoldingsAPI.DeactivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvestmentHoldingsDeactivation**](InvestmentHoldingsDeactivation.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHoldings

> InvestmentHoldingsResponseBody ListHoldings(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List holdings by user



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
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `1000`. If the value exceeds `1000`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvestmentHoldingsAPI.ListHoldings(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestmentHoldingsAPI.ListHoldings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHoldings`: InvestmentHoldingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InvestmentHoldingsAPI.ListHoldings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**InvestmentHoldingsResponseBody**](InvestmentHoldingsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHoldingsByAccount

> InvestmentHoldingsResponseBody ListHoldingsByAccount(ctx, accountGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List holdings by account



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
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `1000`. If the value exceeds `1000`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvestmentHoldingsAPI.ListHoldingsByAccount(context.Background(), accountGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestmentHoldingsAPI.ListHoldingsByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHoldingsByAccount`: InvestmentHoldingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InvestmentHoldingsAPI.ListHoldingsByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHoldingsByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**InvestmentHoldingsResponseBody**](InvestmentHoldingsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHoldingsByMember

> InvestmentHoldingsResponseBody ListHoldingsByMember(ctx, memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List holdings by member



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
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `1000`. If the value exceeds `1000`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvestmentHoldingsAPI.ListHoldingsByMember(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestmentHoldingsAPI.ListHoldingsByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHoldingsByMember`: InvestmentHoldingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InvestmentHoldingsAPI.ListHoldingsByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHoldingsByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**InvestmentHoldingsResponseBody**](InvestmentHoldingsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadHolding

> InvestmentHoldingResponseBody ReadHolding(ctx, holdingGuid, userGuid).Execute()

Read holding



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
    holdingGuid := "HOL-d65683e8-9eab-26bb-bcfd-ced159c9abe2" // string | The unique id for a `holding`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvestmentHoldingsAPI.ReadHolding(context.Background(), holdingGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvestmentHoldingsAPI.ReadHolding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadHolding`: InvestmentHoldingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InvestmentHoldingsAPI.ReadHolding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holdingGuid** | **string** | The unique id for a &#x60;holding&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadHoldingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InvestmentHoldingResponseBody**](InvestmentHoldingResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

