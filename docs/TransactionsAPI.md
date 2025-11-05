# \TransactionsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManualTransaction**](TransactionsAPI.md#CreateManualTransaction) | **Post** /users/{user_guid}/accounts/{account_guid}/transactions | Create manual transaction
[**CreateSplitTransactions**](TransactionsAPI.md#CreateSplitTransactions) | **Post** /users/{user_guid}/transactions/{transaction_guid}/split | Create split transactions
[**DeleteManualTransactions**](TransactionsAPI.md#DeleteManualTransactions) | **Delete** /users/{user_guid}/transactions/{transaction_guid} | Delete manual transactions
[**DeleteSplitTransactions**](TransactionsAPI.md#DeleteSplitTransactions) | **Delete** /users/{user_guid}/transactions/{transaction_guid}/split | Delete split transactions
[**DeleteTransactionRule**](TransactionsAPI.md#DeleteTransactionRule) | **Delete** /users/{user_guid}/transaction_rules/{transaction_rule_guid} | Delete transaction rule
[**EnhanceTransactions**](TransactionsAPI.md#EnhanceTransactions) | **Post** /transactions/enhance | Enhance transactions
[**ExtendHistory**](TransactionsAPI.md#ExtendHistory) | **Post** /users/{user_guid}/members/{member_guid}/extend_history | Extend history
[**ListTransactions**](TransactionsAPI.md#ListTransactions) | **Get** /users/{user_guid}/transactions | List transactions
[**ListTransactionsByAccount**](TransactionsAPI.md#ListTransactionsByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/transactions | List transactions by account
[**ListTransactionsByMember**](TransactionsAPI.md#ListTransactionsByMember) | **Get** /users/{user_guid}/members/{member_guid}/transactions | List transactions by member
[**ListTransactionsByTag**](TransactionsAPI.md#ListTransactionsByTag) | **Get** /users/{user_guid}/tags/{tag_guid}/transactions | List transactions by tag
[**ReadTransaction**](TransactionsAPI.md#ReadTransaction) | **Get** /users/{user_guid}/transactions/{transaction_guid} | Read transaction
[**RepeatingTransactions**](TransactionsAPI.md#RepeatingTransactions) | **Get** /users/{user_guid}/repeating_transactions | List Repeating Transactions
[**SpecificRepeatingTransaction**](TransactionsAPI.md#SpecificRepeatingTransaction) | **Get** /users/{user_guid}/repeating_transactions/{repeating_transaction_guid} | Get a Repeating Transaction
[**UpdateTransaction**](TransactionsAPI.md#UpdateTransaction) | **Put** /users/{user_guid}/transactions/{transaction_guid} | Update transaction
[**UpdateTransactionByAccount**](TransactionsAPI.md#UpdateTransactionByAccount) | **Put** /users/{user_guid}/members/{member_guid}/accounts/{account_guid}/transactions/{transaction_guid} | Update Transaction by Account



## CreateManualTransaction

> TransactionCreateResponseBody CreateManualTransaction(ctx, userGuid, accountGuid).TransactionCreateRequestBody(transactionCreateRequestBody).Execute()

Create manual transaction



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
    transactionCreateRequestBody := *openapiclient.NewTransactionCreateRequestBody() // TransactionCreateRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.CreateManualTransaction(context.Background(), userGuid, accountGuid).TransactionCreateRequestBody(transactionCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.CreateManualTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManualTransaction`: TransactionCreateResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.CreateManualTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateManualTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionCreateRequestBody** | [**TransactionCreateRequestBody**](TransactionCreateRequestBody.md) |  | 

### Return type

[**TransactionCreateResponseBody**](TransactionCreateResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSplitTransactions

> SplitTransactionsResponseBody CreateSplitTransactions(ctx, transactionGuid, userGuid).SplitTransactionRequestBody(splitTransactionRequestBody).Execute()

Create split transactions



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
    splitTransactionRequestBody := *openapiclient.NewSplitTransactionRequestBody(*openapiclient.NewSplitTransactionRequest(float32(54.19))) // SplitTransactionRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.CreateSplitTransactions(context.Background(), transactionGuid, userGuid).SplitTransactionRequestBody(splitTransactionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.CreateSplitTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSplitTransactions`: SplitTransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.CreateSplitTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSplitTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **splitTransactionRequestBody** | [**SplitTransactionRequestBody**](SplitTransactionRequestBody.md) |  | 

### Return type

[**SplitTransactionsResponseBody**](SplitTransactionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManualTransactions

> DeleteManualTransactions(ctx, userGuid, transactionGuid).Execute()

Delete manual transactions



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
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransactionsAPI.DeleteManualTransactions(context.Background(), userGuid, transactionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.DeleteManualTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManualTransactionsRequest struct via the builder pattern


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


## DeleteSplitTransactions

> DeleteSplitTransactions(ctx, transactionGuid, userGuid).Execute()

Delete split transactions



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransactionsAPI.DeleteSplitTransactions(context.Background(), transactionGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.DeleteSplitTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSplitTransactionsRequest struct via the builder pattern


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


## DeleteTransactionRule

> DeleteTransactionRule(ctx, transactionRuleGuid, userGuid).Execute()

Delete transaction rule



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
    transactionRuleGuid := "TXR-a080e0f9-a2d4-4d6f-9e03-672cc357a4d3" // string | The unique id for a `transaction_rule`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransactionsAPI.DeleteTransactionRule(context.Background(), transactionRuleGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.DeleteTransactionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionRuleGuid** | **string** | The unique id for a &#x60;transaction_rule&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransactionRuleRequest struct via the builder pattern


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


## EnhanceTransactions

> EnhanceTransactionsResponseBody EnhanceTransactions(ctx).EnhanceTransactionsRequestBody(enhanceTransactionsRequestBody).Execute()

Enhance transactions



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
    enhanceTransactionsRequestBody := *openapiclient.NewEnhanceTransactionsRequestBody() // EnhanceTransactionsRequestBody | Transaction object to be enhanced

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.EnhanceTransactions(context.Background()).EnhanceTransactionsRequestBody(enhanceTransactionsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.EnhanceTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnhanceTransactions`: EnhanceTransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.EnhanceTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnhanceTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enhanceTransactionsRequestBody** | [**EnhanceTransactionsRequestBody**](EnhanceTransactionsRequestBody.md) | Transaction object to be enhanced | 

### Return type

[**EnhanceTransactionsResponseBody**](EnhanceTransactionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendHistory

> MemberResponseBody ExtendHistory(ctx, memberGuid, userGuid).Execute()

Extend history



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.ExtendHistory(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ExtendHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtendHistory`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ExtendHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MemberResponseBody**](MemberResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> TransactionsResponseBodyIncludes ListTransactions(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).FromDate(fromDate).ToDate(toDate).FromCreatedAt(fromCreatedAt).ToCreatedAt(toCreatedAt).FromUpdatedAt(fromUpdatedAt).ToUpdatedAt(toUpdatedAt).CategoryGuid(categoryGuid).CategoryGuid2(categoryGuid2).TopLevelCategoryGuid(topLevelCategoryGuid).TopLevelCategoryGuid2(topLevelCategoryGuid2).UseCase(useCase).Includes(includes).Execute()

List transactions



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
    fromDate := "2024-01-01" // string | Filter transactions from this date. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 120 days ago if not provided. (optional)
    toDate := "2024-03-31" // string | Filter transactions to this date (at midnight). This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 5 days forward from the day the request is made to capture pending transactions. (optional)
    fromCreatedAt := "2024-01-01" // string | Filter transactions from the date the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    toCreatedAt := "2024-03-31" // string | Filter transaction to the date in which the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    fromUpdatedAt := "2024-01-01" // string | Filter transactions from the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    toUpdatedAt := "2024-03-31" // string | Filter transactions to the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    categoryGuid := "categoryGuid_example" // string | Filter transactions belonging to specified `category_guid`.  For example, `?category_guid=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    categoryGuid2 := []string{"Inner_example"} // []string | Filter transactions belonging to any specified `category_guid[]` in url.  For example, `?category_guid[]=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    topLevelCategoryGuid := "topLevelCategoryGuid_example" // string | Filter transactions belonging to specified `top_level_category_guid`. This must be top level category guid, use `category_guid` for subcategory guid.  For example, `?top_level_category_guid=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    topLevelCategoryGuid2 := []string{"Inner_example"} // []string | Filter transactions belonging to any specified `top_level_category_guid[]` in url. This must be top level category guid(s), use `category_guid` for subcategory guid(s).  For example, `?top_level_category_guid[]=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    useCase := "MONEY_MOVEMENT" // string | The use case associated with the member. Valid values are `PFM` and `MONEY_MOVEMENT`. For example, you can append either `?use_case=PFM` or `?use_case=MONEY_MOVEMENT`. (optional)
    includes := "repeating_transactions,merchants,classifications,geolocations" // string | Options for enhanced transactions. This query parameter is optional. Possible additional metadata: `repeating_transactions`, `merchants`, `classifications`, `geolocations`. The query value is format sensitive. To retrieve all available enhancements, append:  `?includes=repeating_transactions,merchants,classifications,geolocations`.    The query options may be combined to specific enhancements. For example, to request Repeating Transactions and Geolocation data, use:   `?includes=repeating_transactions,geolocations`.  - Repeating Transactions: Identifies transactions with predictable recurrence patterns (e.g., Bill, Income, Subscription). - Merchants: Enriches transactions with merchant name. - Classifications: Provides more insight into the type of money movement that is occurring on the transaction, whether it be retail or investments. - Geolocation: Provides geographic metadata.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.ListTransactions(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).FromDate(fromDate).ToDate(toDate).FromCreatedAt(fromCreatedAt).ToCreatedAt(toCreatedAt).FromUpdatedAt(fromUpdatedAt).ToUpdatedAt(toUpdatedAt).CategoryGuid(categoryGuid).CategoryGuid2(categoryGuid2).TopLevelCategoryGuid(topLevelCategoryGuid).TopLevelCategoryGuid2(topLevelCategoryGuid2).UseCase(useCase).Includes(includes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactions`: TransactionsResponseBodyIncludes
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 
 **fromDate** | **string** | Filter transactions from this date. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 120 days ago if not provided. | 
 **toDate** | **string** | Filter transactions to this date (at midnight). This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 5 days forward from the day the request is made to capture pending transactions. | 
 **fromCreatedAt** | **string** | Filter transactions from the date the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **toCreatedAt** | **string** | Filter transaction to the date in which the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **fromUpdatedAt** | **string** | Filter transactions from the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **toUpdatedAt** | **string** | Filter transactions to the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **categoryGuid** | **string** | Filter transactions belonging to specified &#x60;category_guid&#x60;.  For example, &#x60;?category_guid&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **categoryGuid2** | **[]string** | Filter transactions belonging to any specified &#x60;category_guid[]&#x60; in url.  For example, &#x60;?category_guid[]&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **topLevelCategoryGuid** | **string** | Filter transactions belonging to specified &#x60;top_level_category_guid&#x60;. This must be top level category guid, use &#x60;category_guid&#x60; for subcategory guid.  For example, &#x60;?top_level_category_guid&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **topLevelCategoryGuid2** | **[]string** | Filter transactions belonging to any specified &#x60;top_level_category_guid[]&#x60; in url. This must be top level category guid(s), use &#x60;category_guid&#x60; for subcategory guid(s).  For example, &#x60;?top_level_category_guid[]&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **useCase** | **string** | The use case associated with the member. Valid values are &#x60;PFM&#x60; and &#x60;MONEY_MOVEMENT&#x60;. For example, you can append either &#x60;?use_case&#x3D;PFM&#x60; or &#x60;?use_case&#x3D;MONEY_MOVEMENT&#x60;. | 
 **includes** | **string** | Options for enhanced transactions. This query parameter is optional. Possible additional metadata: &#x60;repeating_transactions&#x60;, &#x60;merchants&#x60;, &#x60;classifications&#x60;, &#x60;geolocations&#x60;. The query value is format sensitive. To retrieve all available enhancements, append:  &#x60;?includes&#x3D;repeating_transactions,merchants,classifications,geolocations&#x60;.    The query options may be combined to specific enhancements. For example, to request Repeating Transactions and Geolocation data, use:   &#x60;?includes&#x3D;repeating_transactions,geolocations&#x60;.  - Repeating Transactions: Identifies transactions with predictable recurrence patterns (e.g., Bill, Income, Subscription). - Merchants: Enriches transactions with merchant name. - Classifications: Provides more insight into the type of money movement that is occurring on the transaction, whether it be retail or investments. - Geolocation: Provides geographic metadata.  | 

### Return type

[**TransactionsResponseBodyIncludes**](TransactionsResponseBodyIncludes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsByAccount

> TransactionsResponseBodyIncludes ListTransactionsByAccount(ctx, accountGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).FromDate(fromDate).ToDate(toDate).FromCreatedAt(fromCreatedAt).ToCreatedAt(toCreatedAt).FromUpdatedAt(fromUpdatedAt).ToUpdatedAt(toUpdatedAt).CategoryGuid(categoryGuid).CategoryGuid2(categoryGuid2).TopLevelCategoryGuid(topLevelCategoryGuid).TopLevelCategoryGuid2(topLevelCategoryGuid2).Includes(includes).Execute()

List transactions by account



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
    fromDate := "2024-01-01" // string | Filter transactions from this date. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 120 days ago if not provided. (optional)
    toDate := "2024-03-31" // string | Filter transactions to this date (at midnight). This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 5 days forward from the day the request is made to capture pending transactions. (optional)
    fromCreatedAt := "2024-01-01" // string | Filter transactions from the date the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    toCreatedAt := "2024-03-31" // string | Filter transaction to the date in which the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    fromUpdatedAt := "2024-01-01" // string | Filter transactions from the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    toUpdatedAt := "2024-03-31" // string | Filter transactions to the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    categoryGuid := "categoryGuid_example" // string | Filter transactions belonging to specified `category_guid`.  For example, `?category_guid=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    categoryGuid2 := []string{"Inner_example"} // []string | Filter transactions belonging to any specified `category_guid[]` in url.  For example, `?category_guid[]=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    topLevelCategoryGuid := "topLevelCategoryGuid_example" // string | Filter transactions belonging to specified `top_level_category_guid`. This must be top level category guid, use `category_guid` for subcategory guid.  For example, `?top_level_category_guid=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    topLevelCategoryGuid2 := []string{"Inner_example"} // []string | Filter transactions belonging to any specified `top_level_category_guid[]` in url. This must be top level category guid(s), use `category_guid` for subcategory guid(s).  For example, `?top_level_category_guid[]=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    includes := "repeating_transactions,merchants,classifications,geolocations" // string | Options for enhanced transactions. This query parameter is optional. Possible additional metadata: `repeating_transactions`, `merchants`, `classifications`, `geolocations`. The query value is format sensitive. To retrieve all available enhancements, append:  `?includes=repeating_transactions,merchants,classifications,geolocations`.    The query options may be combined to specific enhancements. For example, to request Repeating Transactions and Geolocation data, use:   `?includes=repeating_transactions,geolocations`.  - Repeating Transactions: Identifies transactions with predictable recurrence patterns (e.g., Bill, Income, Subscription). - Merchants: Enriches transactions with merchant name. - Classifications: Provides more insight into the type of money movement that is occurring on the transaction, whether it be retail or investments. - Geolocation: Provides geographic metadata.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.ListTransactionsByAccount(context.Background(), accountGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).FromDate(fromDate).ToDate(toDate).FromCreatedAt(fromCreatedAt).ToCreatedAt(toCreatedAt).FromUpdatedAt(fromUpdatedAt).ToUpdatedAt(toUpdatedAt).CategoryGuid(categoryGuid).CategoryGuid2(categoryGuid2).TopLevelCategoryGuid(topLevelCategoryGuid).TopLevelCategoryGuid2(topLevelCategoryGuid2).Includes(includes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactionsByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByAccount`: TransactionsResponseBodyIncludes
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactionsByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 
 **fromDate** | **string** | Filter transactions from this date. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 120 days ago if not provided. | 
 **toDate** | **string** | Filter transactions to this date (at midnight). This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 5 days forward from the day the request is made to capture pending transactions. | 
 **fromCreatedAt** | **string** | Filter transactions from the date the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **toCreatedAt** | **string** | Filter transaction to the date in which the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **fromUpdatedAt** | **string** | Filter transactions from the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **toUpdatedAt** | **string** | Filter transactions to the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **categoryGuid** | **string** | Filter transactions belonging to specified &#x60;category_guid&#x60;.  For example, &#x60;?category_guid&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **categoryGuid2** | **[]string** | Filter transactions belonging to any specified &#x60;category_guid[]&#x60; in url.  For example, &#x60;?category_guid[]&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **topLevelCategoryGuid** | **string** | Filter transactions belonging to specified &#x60;top_level_category_guid&#x60;. This must be top level category guid, use &#x60;category_guid&#x60; for subcategory guid.  For example, &#x60;?top_level_category_guid&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **topLevelCategoryGuid2** | **[]string** | Filter transactions belonging to any specified &#x60;top_level_category_guid[]&#x60; in url. This must be top level category guid(s), use &#x60;category_guid&#x60; for subcategory guid(s).  For example, &#x60;?top_level_category_guid[]&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **includes** | **string** | Options for enhanced transactions. This query parameter is optional. Possible additional metadata: &#x60;repeating_transactions&#x60;, &#x60;merchants&#x60;, &#x60;classifications&#x60;, &#x60;geolocations&#x60;. The query value is format sensitive. To retrieve all available enhancements, append:  &#x60;?includes&#x3D;repeating_transactions,merchants,classifications,geolocations&#x60;.    The query options may be combined to specific enhancements. For example, to request Repeating Transactions and Geolocation data, use:   &#x60;?includes&#x3D;repeating_transactions,geolocations&#x60;.  - Repeating Transactions: Identifies transactions with predictable recurrence patterns (e.g., Bill, Income, Subscription). - Merchants: Enriches transactions with merchant name. - Classifications: Provides more insight into the type of money movement that is occurring on the transaction, whether it be retail or investments. - Geolocation: Provides geographic metadata.  | 

### Return type

[**TransactionsResponseBodyIncludes**](TransactionsResponseBodyIncludes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsByMember

> TransactionsResponseBodyIncludes ListTransactionsByMember(ctx, memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).FromDate(fromDate).ToDate(toDate).FromCreatedAt(fromCreatedAt).ToCreatedAt(toCreatedAt).FromUpdatedAt(fromUpdatedAt).ToUpdatedAt(toUpdatedAt).CategoryGuid(categoryGuid).CategoryGuid2(categoryGuid2).TopLevelCategoryGuid(topLevelCategoryGuid).TopLevelCategoryGuid2(topLevelCategoryGuid2).Includes(includes).Execute()

List transactions by member



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
    fromDate := "2024-01-01" // string | Filter transactions from this date. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 120 days ago if not provided. (optional)
    toDate := "2024-03-31" // string | Filter transactions to this date (at midnight). This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 5 days forward from the day the request is made to capture pending transactions. (optional)
    fromCreatedAt := "2024-01-01" // string | Filter transactions from the date the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    toCreatedAt := "2024-03-31" // string | Filter transaction to the date in which the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    fromUpdatedAt := "2024-01-01" // string | Filter transactions from the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    toUpdatedAt := "2024-03-31" // string | Filter transactions to the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    categoryGuid := "categoryGuid_example" // string | Filter transactions belonging to specified `category_guid`.  For example, `?category_guid=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    categoryGuid2 := []string{"Inner_example"} // []string | Filter transactions belonging to any specified `category_guid[]` in url.  For example, `?category_guid[]=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    topLevelCategoryGuid := "topLevelCategoryGuid_example" // string | Filter transactions belonging to specified `top_level_category_guid`. This must be top level category guid, use `category_guid` for subcategory guid.  For example, `?top_level_category_guid=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    topLevelCategoryGuid2 := []string{"Inner_example"} // []string | Filter transactions belonging to any specified `top_level_category_guid[]` in url. This must be top level category guid(s), use `category_guid` for subcategory guid(s).  For example, `?top_level_category_guid[]=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    includes := "repeating_transactions,merchants,classifications,geolocations" // string | Options for enhanced transactions. This query parameter is optional. Possible additional metadata: `repeating_transactions`, `merchants`, `classifications`, `geolocations`. The query value is format sensitive. To retrieve all available enhancements, append:  `?includes=repeating_transactions,merchants,classifications,geolocations`.    The query options may be combined to specific enhancements. For example, to request Repeating Transactions and Geolocation data, use:   `?includes=repeating_transactions,geolocations`.  - Repeating Transactions: Identifies transactions with predictable recurrence patterns (e.g., Bill, Income, Subscription). - Merchants: Enriches transactions with merchant name. - Classifications: Provides more insight into the type of money movement that is occurring on the transaction, whether it be retail or investments. - Geolocation: Provides geographic metadata.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.ListTransactionsByMember(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).FromDate(fromDate).ToDate(toDate).FromCreatedAt(fromCreatedAt).ToCreatedAt(toCreatedAt).FromUpdatedAt(fromUpdatedAt).ToUpdatedAt(toUpdatedAt).CategoryGuid(categoryGuid).CategoryGuid2(categoryGuid2).TopLevelCategoryGuid(topLevelCategoryGuid).TopLevelCategoryGuid2(topLevelCategoryGuid2).Includes(includes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactionsByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByMember`: TransactionsResponseBodyIncludes
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactionsByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 
 **fromDate** | **string** | Filter transactions from this date. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 120 days ago if not provided. | 
 **toDate** | **string** | Filter transactions to this date (at midnight). This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 5 days forward from the day the request is made to capture pending transactions. | 
 **fromCreatedAt** | **string** | Filter transactions from the date the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **toCreatedAt** | **string** | Filter transaction to the date in which the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **fromUpdatedAt** | **string** | Filter transactions from the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **toUpdatedAt** | **string** | Filter transactions to the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **categoryGuid** | **string** | Filter transactions belonging to specified &#x60;category_guid&#x60;.  For example, &#x60;?category_guid&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **categoryGuid2** | **[]string** | Filter transactions belonging to any specified &#x60;category_guid[]&#x60; in url.  For example, &#x60;?category_guid[]&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **topLevelCategoryGuid** | **string** | Filter transactions belonging to specified &#x60;top_level_category_guid&#x60;. This must be top level category guid, use &#x60;category_guid&#x60; for subcategory guid.  For example, &#x60;?top_level_category_guid&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **topLevelCategoryGuid2** | **[]string** | Filter transactions belonging to any specified &#x60;top_level_category_guid[]&#x60; in url. This must be top level category guid(s), use &#x60;category_guid&#x60; for subcategory guid(s).  For example, &#x60;?top_level_category_guid[]&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **includes** | **string** | Options for enhanced transactions. This query parameter is optional. Possible additional metadata: &#x60;repeating_transactions&#x60;, &#x60;merchants&#x60;, &#x60;classifications&#x60;, &#x60;geolocations&#x60;. The query value is format sensitive. To retrieve all available enhancements, append:  &#x60;?includes&#x3D;repeating_transactions,merchants,classifications,geolocations&#x60;.    The query options may be combined to specific enhancements. For example, to request Repeating Transactions and Geolocation data, use:   &#x60;?includes&#x3D;repeating_transactions,geolocations&#x60;.  - Repeating Transactions: Identifies transactions with predictable recurrence patterns (e.g., Bill, Income, Subscription). - Merchants: Enriches transactions with merchant name. - Classifications: Provides more insight into the type of money movement that is occurring on the transaction, whether it be retail or investments. - Geolocation: Provides geographic metadata.  | 

### Return type

[**TransactionsResponseBodyIncludes**](TransactionsResponseBodyIncludes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsByTag

> TransactionsResponseBodyIncludes ListTransactionsByTag(ctx, userGuid, tagGuid).Page(page).RecordsPerPage(recordsPerPage).FromDate(fromDate).ToDate(toDate).FromCreatedAt(fromCreatedAt).ToCreatedAt(toCreatedAt).FromUpdatedAt(fromUpdatedAt).ToUpdatedAt(toUpdatedAt).CategoryGuid(categoryGuid).CategoryGuid2(categoryGuid2).TopLevelCategoryGuid(topLevelCategoryGuid).TopLevelCategoryGuid2(topLevelCategoryGuid2).UseCase(useCase).Includes(includes).Execute()

List transactions by tag



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
    tagGuid := "TAG-aef36e72-6294-4c38-844d-e573e80aed52" // string | The unique id for a `tag`.
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `1000`. If the value exceeds `1000`, the default value of `25` will be used instead. (optional)
    fromDate := "2024-01-01" // string | Filter transactions from this date. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 120 days ago if not provided. (optional)
    toDate := "2024-03-31" // string | Filter transactions to this date (at midnight). This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 5 days forward from the day the request is made to capture pending transactions. (optional)
    fromCreatedAt := "2024-01-01" // string | Filter transactions from the date the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    toCreatedAt := "2024-03-31" // string | Filter transaction to the date in which the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    fromUpdatedAt := "2024-01-01" // string | Filter transactions from the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    toUpdatedAt := "2024-03-31" // string | Filter transactions to the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. (optional)
    categoryGuid := "categoryGuid_example" // string | Filter transactions belonging to specified `category_guid`.  For example, `?category_guid=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    categoryGuid2 := []string{"Inner_example"} // []string | Filter transactions belonging to any specified `category_guid[]` in url.  For example, `?category_guid[]=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    topLevelCategoryGuid := "topLevelCategoryGuid_example" // string | Filter transactions belonging to specified `top_level_category_guid`. This must be top level category guid, use `category_guid` for subcategory guid.  For example, `?top_level_category_guid=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    topLevelCategoryGuid2 := []string{"Inner_example"} // []string | Filter transactions belonging to any specified `top_level_category_guid[]` in url. This must be top level category guid(s), use `category_guid` for subcategory guid(s).  For example, `?top_level_category_guid[]=CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874`. (optional)
    useCase := "MONEY_MOVEMENT" // string | The use case associated with the member. Valid values are `PFM` and `MONEY_MOVEMENT`. For example, you can append either `?use_case=PFM` or `?use_case=MONEY_MOVEMENT`. (optional)
    includes := "repeating_transactions,merchants,classifications,geolocations" // string | Options for enhanced transactions. This query parameter is optional. Possible additional metadata: `repeating_transactions`, `merchants`, `classifications`, `geolocations`. The query value is format sensitive. To retrieve all available enhancements, append:  `?includes=repeating_transactions,merchants,classifications,geolocations`.    The query options may be combined to specific enhancements. For example, to request Repeating Transactions and Geolocation data, use:   `?includes=repeating_transactions,geolocations`.  - Repeating Transactions: Identifies transactions with predictable recurrence patterns (e.g., Bill, Income, Subscription). - Merchants: Enriches transactions with merchant name. - Classifications: Provides more insight into the type of money movement that is occurring on the transaction, whether it be retail or investments. - Geolocation: Provides geographic metadata.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.ListTransactionsByTag(context.Background(), userGuid, tagGuid).Page(page).RecordsPerPage(recordsPerPage).FromDate(fromDate).ToDate(toDate).FromCreatedAt(fromCreatedAt).ToCreatedAt(toCreatedAt).FromUpdatedAt(fromUpdatedAt).ToUpdatedAt(toUpdatedAt).CategoryGuid(categoryGuid).CategoryGuid2(categoryGuid2).TopLevelCategoryGuid(topLevelCategoryGuid).TopLevelCategoryGuid2(topLevelCategoryGuid2).UseCase(useCase).Includes(includes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactionsByTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByTag`: TransactionsResponseBodyIncludes
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactionsByTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**tagGuid** | **string** | The unique id for a &#x60;tag&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsByTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 
 **fromDate** | **string** | Filter transactions from this date. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 120 days ago if not provided. | 
 **toDate** | **string** | Filter transactions to this date (at midnight). This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Defaults to 5 days forward from the day the request is made to capture pending transactions. | 
 **fromCreatedAt** | **string** | Filter transactions from the date the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **toCreatedAt** | **string** | Filter transaction to the date in which the transaction was created. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **fromUpdatedAt** | **string** | Filter transactions from the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **toUpdatedAt** | **string** | Filter transactions to the date in which the transaction was updated. This only supports ISO 8601 format without timestamp (YYYY-MM-DD). Maximum date range limit is 6 months. | 
 **categoryGuid** | **string** | Filter transactions belonging to specified &#x60;category_guid&#x60;.  For example, &#x60;?category_guid&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **categoryGuid2** | **[]string** | Filter transactions belonging to any specified &#x60;category_guid[]&#x60; in url.  For example, &#x60;?category_guid[]&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **topLevelCategoryGuid** | **string** | Filter transactions belonging to specified &#x60;top_level_category_guid&#x60;. This must be top level category guid, use &#x60;category_guid&#x60; for subcategory guid.  For example, &#x60;?top_level_category_guid&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **topLevelCategoryGuid2** | **[]string** | Filter transactions belonging to any specified &#x60;top_level_category_guid[]&#x60; in url. This must be top level category guid(s), use &#x60;category_guid&#x60; for subcategory guid(s).  For example, &#x60;?top_level_category_guid[]&#x3D;CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874&#x60;. | 
 **useCase** | **string** | The use case associated with the member. Valid values are &#x60;PFM&#x60; and &#x60;MONEY_MOVEMENT&#x60;. For example, you can append either &#x60;?use_case&#x3D;PFM&#x60; or &#x60;?use_case&#x3D;MONEY_MOVEMENT&#x60;. | 
 **includes** | **string** | Options for enhanced transactions. This query parameter is optional. Possible additional metadata: &#x60;repeating_transactions&#x60;, &#x60;merchants&#x60;, &#x60;classifications&#x60;, &#x60;geolocations&#x60;. The query value is format sensitive. To retrieve all available enhancements, append:  &#x60;?includes&#x3D;repeating_transactions,merchants,classifications,geolocations&#x60;.    The query options may be combined to specific enhancements. For example, to request Repeating Transactions and Geolocation data, use:   &#x60;?includes&#x3D;repeating_transactions,geolocations&#x60;.  - Repeating Transactions: Identifies transactions with predictable recurrence patterns (e.g., Bill, Income, Subscription). - Merchants: Enriches transactions with merchant name. - Classifications: Provides more insight into the type of money movement that is occurring on the transaction, whether it be retail or investments. - Geolocation: Provides geographic metadata.  | 

### Return type

[**TransactionsResponseBodyIncludes**](TransactionsResponseBodyIncludes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTransaction

> TransactionsResponseBodyIncludes ReadTransaction(ctx, userGuid, transactionGuid).Includes(includes).Execute()

Read transaction



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
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    includes := "repeating_transactions,merchants,classifications,geolocations" // string | Options for enhanced transactions. This query parameter is optional. Possible additional metadata: `repeating_transactions`, `merchants`, `classifications`, `geolocations`. The query value is format sensitive. To retrieve all available enhancements, append:  `?includes=repeating_transactions,merchants,classifications,geolocations`.    The query options may be combined to specific enhancements. For example, to request Repeating Transactions and Geolocation data, use:   `?includes=repeating_transactions,geolocations`.  - Repeating Transactions: Identifies transactions with predictable recurrence patterns (e.g., Bill, Income, Subscription). - Merchants: Enriches transactions with merchant name. - Classifications: Provides more insight into the type of money movement that is occurring on the transaction, whether it be retail or investments. - Geolocation: Provides geographic metadata.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.ReadTransaction(context.Background(), userGuid, transactionGuid).Includes(includes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ReadTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTransaction`: TransactionsResponseBodyIncludes
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ReadTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includes** | **string** | Options for enhanced transactions. This query parameter is optional. Possible additional metadata: &#x60;repeating_transactions&#x60;, &#x60;merchants&#x60;, &#x60;classifications&#x60;, &#x60;geolocations&#x60;. The query value is format sensitive. To retrieve all available enhancements, append:  &#x60;?includes&#x3D;repeating_transactions,merchants,classifications,geolocations&#x60;.    The query options may be combined to specific enhancements. For example, to request Repeating Transactions and Geolocation data, use:   &#x60;?includes&#x3D;repeating_transactions,geolocations&#x60;.  - Repeating Transactions: Identifies transactions with predictable recurrence patterns (e.g., Bill, Income, Subscription). - Merchants: Enriches transactions with merchant name. - Classifications: Provides more insight into the type of money movement that is occurring on the transaction, whether it be retail or investments. - Geolocation: Provides geographic metadata.  | 

### Return type

[**TransactionsResponseBodyIncludes**](TransactionsResponseBodyIncludes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepeatingTransactions

> RepeatingTransactionsResponseBody RepeatingTransactions(ctx, userGuid).Execute()

List Repeating Transactions



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
    resp, r, err := apiClient.TransactionsAPI.RepeatingTransactions(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.RepeatingTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepeatingTransactions`: RepeatingTransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.RepeatingTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepeatingTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RepeatingTransactionsResponseBody**](RepeatingTransactionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpecificRepeatingTransaction

> RepeatingTransactionsResponseBody SpecificRepeatingTransaction(ctx, userGuid, repeatingTransactionGuid).Execute()

Get a Repeating Transaction



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
    repeatingTransactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a recurring transaction.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.SpecificRepeatingTransaction(context.Background(), userGuid, repeatingTransactionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.SpecificRepeatingTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpecificRepeatingTransaction`: RepeatingTransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.SpecificRepeatingTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**repeatingTransactionGuid** | **string** | The unique id for a recurring transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpecificRepeatingTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepeatingTransactionsResponseBody**](RepeatingTransactionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransaction

> TransactionResponseBody UpdateTransaction(ctx, userGuid, transactionGuid).TransactionUpdateRequestBody(transactionUpdateRequestBody).Execute()

Update transaction



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
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    transactionUpdateRequestBody := *openapiclient.NewTransactionUpdateRequestBody() // TransactionUpdateRequestBody | Transaction object with the fields to be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.UpdateTransaction(context.Background(), userGuid, transactionGuid).TransactionUpdateRequestBody(transactionUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.UpdateTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransaction`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.UpdateTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionUpdateRequestBody** | [**TransactionUpdateRequestBody**](TransactionUpdateRequestBody.md) | Transaction object with the fields to be updated. | 

### Return type

[**TransactionResponseBody**](TransactionResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransactionByAccount

> TransactionResponseBody UpdateTransactionByAccount(ctx, userGuid, memberGuid, accountGuid, transactionGuid).TransactionUpdateRequestBody(transactionUpdateRequestBody).Execute()

Update Transaction by Account



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
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    transactionUpdateRequestBody := *openapiclient.NewTransactionUpdateRequestBody() // TransactionUpdateRequestBody | Transaction object to be updated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.UpdateTransactionByAccount(context.Background(), userGuid, memberGuid, accountGuid, transactionGuid).TransactionUpdateRequestBody(transactionUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.UpdateTransactionByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransactionByAccount`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.UpdateTransactionByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransactionByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionUpdateRequestBody** | [**TransactionUpdateRequestBody**](TransactionUpdateRequestBody.md) | Transaction object to be updated | 

### Return type

[**TransactionResponseBody**](TransactionResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

