# \BudgetsAPI

All URIs are relative to *https://api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersUserGuidBudgetsBudgetGuidDelete**](BudgetsAPI.md#UsersUserGuidBudgetsBudgetGuidDelete) | **Delete** /users/{user_guid}/budgets/{budget_guid} | Delete a budget
[**UsersUserGuidBudgetsBudgetGuidGet**](BudgetsAPI.md#UsersUserGuidBudgetsBudgetGuidGet) | **Get** /users/{user_guid}/budgets/{budget_guid} | Read a specific budget
[**UsersUserGuidBudgetsBudgetGuidPut**](BudgetsAPI.md#UsersUserGuidBudgetsBudgetGuidPut) | **Put** /users/{user_guid}/budgets/{budget_guid} | Update a specific budget
[**UsersUserGuidBudgetsGeneratePost**](BudgetsAPI.md#UsersUserGuidBudgetsGeneratePost) | **Post** /users/{user_guid}/budgets/generate | Auto-generate budgets
[**UsersUserGuidBudgetsGet**](BudgetsAPI.md#UsersUserGuidBudgetsGet) | **Get** /users/{user_guid}/budgets | List all budgets
[**UsersUserGuidBudgetsPost**](BudgetsAPI.md#UsersUserGuidBudgetsPost) | **Post** /users/{user_guid}/budgets | Create a budget



## UsersUserGuidBudgetsBudgetGuidDelete

> UsersUserGuidBudgetsBudgetGuidDelete(ctx, userGuid, budgetGuid).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for the budget. Defined by MX.
    budgetGuid := "budgetGuid_example" // string | The unique identifier for the budget. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BudgetsAPI.UsersUserGuidBudgetsBudgetGuidDelete(context.Background(), userGuid, budgetGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.UsersUserGuidBudgetsBudgetGuidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the budget. Defined by MX. | 
**budgetGuid** | **string** | The unique identifier for the budget. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidBudgetsBudgetGuidDeleteRequest struct via the builder pattern


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


## UsersUserGuidBudgetsBudgetGuidGet

> BudgetResponseBody UsersUserGuidBudgetsBudgetGuidGet(ctx, budgetGuid, userGuid).Execute()

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
    budgetGuid := "budgetGuid_example" // string | The unique identifier for the budget. Defined by MX.
    userGuid := "userGuid_example" // string | The unique identifier for the budget. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsAPI.UsersUserGuidBudgetsBudgetGuidGet(context.Background(), budgetGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.UsersUserGuidBudgetsBudgetGuidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidBudgetsBudgetGuidGet`: BudgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.UsersUserGuidBudgetsBudgetGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetGuid** | **string** | The unique identifier for the budget. Defined by MX. | 
**userGuid** | **string** | The unique identifier for the budget. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidBudgetsBudgetGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BudgetResponseBody**](BudgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserGuidBudgetsBudgetGuidPut

> BudgetResponseBody UsersUserGuidBudgetsBudgetGuidPut(ctx, userGuid, budgetGuid).BudgetUpdateRequestBody(budgetUpdateRequestBody).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for the budget. Defined by MX.
    budgetGuid := "budgetGuid_example" // string | The unique identifier for the budget. Defined by MX.
    budgetUpdateRequestBody := *openapiclient.NewBudgetUpdateRequestBody() // BudgetUpdateRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsAPI.UsersUserGuidBudgetsBudgetGuidPut(context.Background(), userGuid, budgetGuid).BudgetUpdateRequestBody(budgetUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.UsersUserGuidBudgetsBudgetGuidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidBudgetsBudgetGuidPut`: BudgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.UsersUserGuidBudgetsBudgetGuidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the budget. Defined by MX. | 
**budgetGuid** | **string** | The unique identifier for the budget. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidBudgetsBudgetGuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **budgetUpdateRequestBody** | [**BudgetUpdateRequestBody**](BudgetUpdateRequestBody.md) |  | 

### Return type

[**BudgetResponseBody**](BudgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserGuidBudgetsGeneratePost

> BudgetResponseBody UsersUserGuidBudgetsGeneratePost(ctx, userGuid).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for the user. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsAPI.UsersUserGuidBudgetsGeneratePost(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.UsersUserGuidBudgetsGeneratePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidBudgetsGeneratePost`: BudgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.UsersUserGuidBudgetsGeneratePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the user. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidBudgetsGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BudgetResponseBody**](BudgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserGuidBudgetsGet

> BudgetResponseBody UsersUserGuidBudgetsGet(ctx, userGuid).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for the user. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsAPI.UsersUserGuidBudgetsGet(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.UsersUserGuidBudgetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidBudgetsGet`: BudgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.UsersUserGuidBudgetsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the user. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidBudgetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BudgetResponseBody**](BudgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserGuidBudgetsPost

> BudgetResponseBody UsersUserGuidBudgetsPost(ctx, userGuid).BudgetCreateRequestBody(budgetCreateRequestBody).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for the user. Defined by MX.
    budgetCreateRequestBody := *openapiclient.NewBudgetCreateRequestBody() // BudgetCreateRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsAPI.UsersUserGuidBudgetsPost(context.Background(), userGuid).BudgetCreateRequestBody(budgetCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.UsersUserGuidBudgetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidBudgetsPost`: BudgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.UsersUserGuidBudgetsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the user. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidBudgetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **budgetCreateRequestBody** | [**BudgetCreateRequestBody**](BudgetCreateRequestBody.md) |  | 

### Return type

[**BudgetResponseBody**](BudgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

