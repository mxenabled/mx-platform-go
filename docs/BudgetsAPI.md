# \BudgetsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoGenerateBudgets**](BudgetsAPI.md#AutoGenerateBudgets) | **Post** /users/{user_guid}/budgets/generate | Auto-generate budgets
[**CreateBudget**](BudgetsAPI.md#CreateBudget) | **Post** /users/{user_guid}/budgets | Create a budget
[**DeleteBudget**](BudgetsAPI.md#DeleteBudget) | **Delete** /users/{user_guid}/budgets/{budget_guid} | Delete a budget
[**ListAllBudgets**](BudgetsAPI.md#ListAllBudgets) | **Get** /users/{user_guid}/budgets | List all budgets
[**ReadSpecificBudget**](BudgetsAPI.md#ReadSpecificBudget) | **Get** /users/{user_guid}/budgets/{budget_guid} | Read a specific budget
[**UpdateSpecificBudget**](BudgetsAPI.md#UpdateSpecificBudget) | **Put** /users/{user_guid}/budgets/{budget_guid} | Update a specific budget



## AutoGenerateBudgets

> BudgetResponseBody AutoGenerateBudgets(ctx, userGuid).Execute()

Auto-generate budgets



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
    resp, r, err := apiClient.BudgetsAPI.AutoGenerateBudgets(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.AutoGenerateBudgets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutoGenerateBudgets`: BudgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.AutoGenerateBudgets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoGenerateBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BudgetResponseBody**](BudgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBudget

> BudgetResponseBody CreateBudget(ctx, userGuid).BudgetCreateRequestBody(budgetCreateRequestBody).Execute()

Create a budget



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
    budgetCreateRequestBody := *openapiclient.NewBudgetCreateRequestBody() // BudgetCreateRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsAPI.CreateBudget(context.Background(), userGuid).BudgetCreateRequestBody(budgetCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.CreateBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBudget`: BudgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.CreateBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **budgetCreateRequestBody** | [**BudgetCreateRequestBody**](BudgetCreateRequestBody.md) |  | 

### Return type

[**BudgetResponseBody**](BudgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/vnd.mx.api.v1+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBudget

> DeleteBudget(ctx, userGuid, budgetGuid).Execute()

Delete a budget



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
    budgetGuid := "budgetGuid_example" // string | The unique identifier for the budget. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BudgetsAPI.DeleteBudget(context.Background(), userGuid, budgetGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.DeleteBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**budgetGuid** | **string** | The unique identifier for the budget. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBudgetRequest struct via the builder pattern


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


## ListAllBudgets

> BudgetResponseBody ListAllBudgets(ctx, userGuid).Execute()

List all budgets



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
    resp, r, err := apiClient.BudgetsAPI.ListAllBudgets(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.ListAllBudgets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllBudgets`: BudgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.ListAllBudgets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BudgetResponseBody**](BudgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSpecificBudget

> BudgetResponseBody ReadSpecificBudget(ctx, userGuid, budgetGuid).Execute()

Read a specific budget



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
    budgetGuid := "budgetGuid_example" // string | The unique identifier for the budget. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsAPI.ReadSpecificBudget(context.Background(), userGuid, budgetGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.ReadSpecificBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSpecificBudget`: BudgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.ReadSpecificBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**budgetGuid** | **string** | The unique identifier for the budget. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSpecificBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BudgetResponseBody**](BudgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpecificBudget

> BudgetResponseBody UpdateSpecificBudget(ctx, userGuid, budgetGuid).BudgetUpdateRequestBody(budgetUpdateRequestBody).Execute()

Update a specific budget



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
    budgetGuid := "budgetGuid_example" // string | The unique identifier for the budget. Defined by MX.
    budgetUpdateRequestBody := *openapiclient.NewBudgetUpdateRequestBody() // BudgetUpdateRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsAPI.UpdateSpecificBudget(context.Background(), userGuid, budgetGuid).BudgetUpdateRequestBody(budgetUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.UpdateSpecificBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpecificBudget`: BudgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.UpdateSpecificBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**budgetGuid** | **string** | The unique identifier for the budget. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpecificBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **budgetUpdateRequestBody** | [**BudgetUpdateRequestBody**](BudgetUpdateRequestBody.md) |  | 

### Return type

[**BudgetResponseBody**](BudgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/vnd.mx.api.v1+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

