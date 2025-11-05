# \InsightsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccountsInsight**](InsightsAPI.md#ListAccountsInsight) | **Get** /users/{user_guid}/insights/{insight_guid}/accounts | List all accounts associated with an insight
[**ListCategoriesInsight**](InsightsAPI.md#ListCategoriesInsight) | **Get** /users/{user_guid}/insights/{insight_guid}/categories | List all categories associated with an insight
[**ListInsightsByAccount**](InsightsAPI.md#ListInsightsByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/insights | List insights by account
[**ListInsightsByTransaction**](InsightsAPI.md#ListInsightsByTransaction) | **Get** /users/{user_guid}/transactions/{transaction_guid}/insights | List insights by transaction
[**ListInsightsUser**](InsightsAPI.md#ListInsightsUser) | **Get** /users/{user_guid}/insights | List all insights for a user
[**ListMerchantsInsight**](InsightsAPI.md#ListMerchantsInsight) | **Get** /users/{user_guid}/insights/{insight_guid}/merchants | List all merchants associated with an insight
[**ListScheduledPaymentsInsight**](InsightsAPI.md#ListScheduledPaymentsInsight) | **Get** /users/{user_guid}/insights/{insight_guid}/scheduled_payments | List all scheduled payments associated with an insight
[**ListTransactionsInsight**](InsightsAPI.md#ListTransactionsInsight) | **Get** /users/{user_guid}/insights/{insight_guid}/transactions | List all transactions associated with an insight
[**ReadInsightUser**](InsightsAPI.md#ReadInsightUser) | **Get** /users/{user_guid}/insights/{insight_guid} | Read insight
[**UpdateInsight**](InsightsAPI.md#UpdateInsight) | **Put** /users/{user_guid}/insights/{insight_guid} | Update insight



## ListAccountsInsight

> AccountsResponseBody ListAccountsInsight(ctx, userGuid, insightGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List all accounts associated with an insight



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
    insightGuid := "BET-1234-abcd" // string | The unique identifier for the insight. Defined by MX.
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsAPI.ListAccountsInsight(context.Background(), userGuid, insightGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.ListAccountsInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountsInsight`: AccountsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.ListAccountsInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**insightGuid** | **string** | The unique identifier for the insight. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;100&#x60;. If the value exceeds &#x60;100&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**AccountsResponseBody**](AccountsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCategoriesInsight

> CategoriesResponseBody ListCategoriesInsight(ctx, userGuid, insightGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List all categories associated with an insight



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
    insightGuid := "BET-1234-abcd" // string | The unique identifier for the insight. Defined by MX.
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsAPI.ListCategoriesInsight(context.Background(), userGuid, insightGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.ListCategoriesInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCategoriesInsight`: CategoriesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.ListCategoriesInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**insightGuid** | **string** | The unique identifier for the insight. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCategoriesInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;100&#x60;. If the value exceeds &#x60;100&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**CategoriesResponseBody**](CategoriesResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInsightsByAccount

> InsightsResponseBody ListInsightsByAccount(ctx, accountGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List insights by account



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
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsAPI.ListInsightsByAccount(context.Background(), accountGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.ListInsightsByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInsightsByAccount`: InsightsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.ListInsightsByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInsightsByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;100&#x60;. If the value exceeds &#x60;100&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**InsightsResponseBody**](InsightsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInsightsByTransaction

> InsightsResponseBody ListInsightsByTransaction(ctx, transactionGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List insights by transaction



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
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsAPI.ListInsightsByTransaction(context.Background(), transactionGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.ListInsightsByTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInsightsByTransaction`: InsightsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.ListInsightsByTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInsightsByTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;100&#x60;. If the value exceeds &#x60;100&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**InsightsResponseBody**](InsightsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInsightsUser

> InsightsResponseBody ListInsightsUser(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List all insights for a user



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
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsAPI.ListInsightsUser(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.ListInsightsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInsightsUser`: InsightsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.ListInsightsUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInsightsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;100&#x60;. If the value exceeds &#x60;100&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**InsightsResponseBody**](InsightsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMerchantsInsight

> MerchantsResponseBody ListMerchantsInsight(ctx, userGuid, insightGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List all merchants associated with an insight



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
    insightGuid := "BET-1234-abcd" // string | The unique identifier for the insight. Defined by MX.
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsAPI.ListMerchantsInsight(context.Background(), userGuid, insightGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.ListMerchantsInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMerchantsInsight`: MerchantsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.ListMerchantsInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**insightGuid** | **string** | The unique identifier for the insight. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMerchantsInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;100&#x60;. If the value exceeds &#x60;100&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**MerchantsResponseBody**](MerchantsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScheduledPaymentsInsight

> ScheduledPaymentsResponseBody ListScheduledPaymentsInsight(ctx, userGuid, insightGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List all scheduled payments associated with an insight



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
    insightGuid := "BET-1234-abcd" // string | The unique identifier for the insight. Defined by MX.
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsAPI.ListScheduledPaymentsInsight(context.Background(), userGuid, insightGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.ListScheduledPaymentsInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScheduledPaymentsInsight`: ScheduledPaymentsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.ListScheduledPaymentsInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**insightGuid** | **string** | The unique identifier for the insight. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListScheduledPaymentsInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;100&#x60;. If the value exceeds &#x60;100&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**ScheduledPaymentsResponseBody**](ScheduledPaymentsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsInsight

> TransactionsResponseBody ListTransactionsInsight(ctx, userGuid, insightGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List all transactions associated with an insight



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
    insightGuid := "BET-1234-abcd" // string | The unique identifier for the insight. Defined by MX.
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsAPI.ListTransactionsInsight(context.Background(), userGuid, insightGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.ListTransactionsInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsInsight`: TransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.ListTransactionsInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**insightGuid** | **string** | The unique identifier for the insight. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;100&#x60;. If the value exceeds &#x60;100&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**TransactionsResponseBody**](TransactionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInsightUser

> InsightResponseBody ReadInsightUser(ctx, userGuid, insightGuid).Execute()

Read insight



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
    insightGuid := "BET-1234-abcd" // string | The unique identifier for the insight. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsAPI.ReadInsightUser(context.Background(), userGuid, insightGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.ReadInsightUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadInsightUser`: InsightResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.ReadInsightUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**insightGuid** | **string** | The unique identifier for the insight. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadInsightUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InsightResponseBody**](InsightResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsight

> InsightResponse UpdateInsight(ctx, userGuid, insightGuid).InsightUpdateRequestBody(insightUpdateRequestBody).Execute()

Update insight



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
    insightGuid := "BET-1234-abcd" // string | The unique identifier for the insight. Defined by MX.
    insightUpdateRequestBody := *openapiclient.NewInsightUpdateRequestBody() // InsightUpdateRequestBody | The insight to be updated (None of these parameters are required, but the user object cannot be empty.)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsAPI.UpdateInsight(context.Background(), userGuid, insightGuid).InsightUpdateRequestBody(insightUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.UpdateInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsight`: InsightResponse
    fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.UpdateInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**insightGuid** | **string** | The unique identifier for the insight. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **insightUpdateRequestBody** | [**InsightUpdateRequestBody**](InsightUpdateRequestBody.md) | The insight to be updated (None of these parameters are required, but the user object cannot be empty.) | 

### Return type

[**InsightResponse**](InsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

