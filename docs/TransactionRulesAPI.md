# \TransactionRulesAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransactionRule**](TransactionRulesAPI.md#CreateTransactionRule) | **Post** /users/{user_guid}/transaction_rules | Create transaction rule
[**ListTransactionRules**](TransactionRulesAPI.md#ListTransactionRules) | **Get** /users/{user_guid}/transaction_rules | List transaction rules
[**ReadTransactionRule**](TransactionRulesAPI.md#ReadTransactionRule) | **Get** /users/{user_guid}/transaction_rules/{transaction_rule_guid} | Read transaction rule
[**UpdateTransactionRule**](TransactionRulesAPI.md#UpdateTransactionRule) | **Put** /users/{user_guid}/transaction_rules/{transaction_rule_guid} | Update transaction rule



## CreateTransactionRule

> TransactionRuleResponseBody CreateTransactionRule(ctx, userGuid).TransactionRuleCreateRequestBody(transactionRuleCreateRequestBody).Execute()

Create transaction rule



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
    transactionRuleCreateRequestBody := *openapiclient.NewTransactionRuleCreateRequestBody() // TransactionRuleCreateRequestBody | TransactionRule object to be created with optional parameters (description) and required parameters (category_guid and match_description)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionRulesAPI.CreateTransactionRule(context.Background(), userGuid).TransactionRuleCreateRequestBody(transactionRuleCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionRulesAPI.CreateTransactionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransactionRule`: TransactionRuleResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionRulesAPI.CreateTransactionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionRuleCreateRequestBody** | [**TransactionRuleCreateRequestBody**](TransactionRuleCreateRequestBody.md) | TransactionRule object to be created with optional parameters (description) and required parameters (category_guid and match_description) | 

### Return type

[**TransactionRuleResponseBody**](TransactionRuleResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionRules

> TransactionRulesResponseBody ListTransactionRules(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List transaction rules



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
    resp, r, err := apiClient.TransactionRulesAPI.ListTransactionRules(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionRulesAPI.ListTransactionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionRules`: TransactionRulesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionRulesAPI.ListTransactionRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**TransactionRulesResponseBody**](TransactionRulesResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTransactionRule

> TransactionRuleResponseBody ReadTransactionRule(ctx, transactionRuleGuid, userGuid).Execute()

Read transaction rule



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
    resp, r, err := apiClient.TransactionRulesAPI.ReadTransactionRule(context.Background(), transactionRuleGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionRulesAPI.ReadTransactionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTransactionRule`: TransactionRuleResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionRulesAPI.ReadTransactionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionRuleGuid** | **string** | The unique id for a &#x60;transaction_rule&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTransactionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TransactionRuleResponseBody**](TransactionRuleResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransactionRule

> TransactionRuleResponseBody UpdateTransactionRule(ctx, transactionRuleGuid, userGuid).TransactionRuleUpdateRequestBody(transactionRuleUpdateRequestBody).Execute()

Update transaction rule



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
    transactionRuleUpdateRequestBody := *openapiclient.NewTransactionRuleUpdateRequestBody() // TransactionRuleUpdateRequestBody | TransactionRule object to be updated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionRulesAPI.UpdateTransactionRule(context.Background(), transactionRuleGuid, userGuid).TransactionRuleUpdateRequestBody(transactionRuleUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionRulesAPI.UpdateTransactionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransactionRule`: TransactionRuleResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionRulesAPI.UpdateTransactionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionRuleGuid** | **string** | The unique id for a &#x60;transaction_rule&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransactionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionRuleUpdateRequestBody** | [**TransactionRuleUpdateRequestBody**](TransactionRuleUpdateRequestBody.md) | TransactionRule object to be updated | 

### Return type

[**TransactionRuleResponseBody**](TransactionRuleResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

