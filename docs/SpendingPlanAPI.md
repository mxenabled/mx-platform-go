# \SpendingPlanAPI

All URIs are relative to *https://api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpendingPlan**](SpendingPlanAPI.md#CreateSpendingPlan) | **Post** /users/{user_guid}/spending_plans | Create spending plan
[**CreateSpendingPlanIterationItem**](SpendingPlanAPI.md#CreateSpendingPlanIterationItem) | **Post** /users/{user_guid}/spending_plans/{spending_plan_guid}/iterations/current/iteration_items | Create spending plan iteration item
[**DeleteSpendingPlan**](SpendingPlanAPI.md#DeleteSpendingPlan) | **Delete** /users/{user_guid}/spending_plans/{spending_plan_guid} | Delete spending plan
[**DeleteSpendingPlanAccount**](SpendingPlanAPI.md#DeleteSpendingPlanAccount) | **Delete** /users/{user_guid}/spending_plans/{spending_plan_guid}/spending_plan_accounts/{spending_plan_account_guid} | Delete spending plan account
[**DeleteSpendingPlanIterationItem**](SpendingPlanAPI.md#DeleteSpendingPlanIterationItem) | **Delete** /users/{user_guid}/spending_plans/{spending_plan_guid}/iterations/current/iteration_items/{iteration_item_guid} | Delete spending plan iteration item
[**ListSpendingPlanAccounts**](SpendingPlanAPI.md#ListSpendingPlanAccounts) | **Get** /users/{user_guid}/spending_plans/{spending_plan_guid}/spending_plan_accounts | List spending plan accounts
[**ListSpendingPlanIterationItems**](SpendingPlanAPI.md#ListSpendingPlanIterationItems) | **Get** /users/{user_guid}/spending_plans/{spending_plan_guid}/iterations/current/iteration_items | List spending plan iteration items
[**ListSpendingPlanIterations**](SpendingPlanAPI.md#ListSpendingPlanIterations) | **Get** /users/{user_guid}/spending_plans/{spending_plan_guid}/iterations | List spending plan iterations
[**ListSpendingPlans**](SpendingPlanAPI.md#ListSpendingPlans) | **Get** /users/{user_guid}/spending_plans | List spending plans
[**ReadSpendingPlanAccount**](SpendingPlanAPI.md#ReadSpendingPlanAccount) | **Get** /users/{user_guid}/spending_plans/{spending_plan_guid}/spending_plan_accounts/{spending_plan_account_guid} | Read spending plan account
[**ReadSpendingPlanIteration**](SpendingPlanAPI.md#ReadSpendingPlanIteration) | **Get** /users/{user_guid}/spending_plans/{spending_plan_guid}/iterations/{iteration_number} | Read a spending plan iteration
[**ReadSpendingPlanIterationItem**](SpendingPlanAPI.md#ReadSpendingPlanIterationItem) | **Get** /users/{user_guid}/spending_plans/{spending_plan_guid}/iterations/current/iteration_items/{iteration_item_guid} | Read a spending plan iteration item
[**ReadSpendingPlanUser**](SpendingPlanAPI.md#ReadSpendingPlanUser) | **Get** /users/{user_guid}/spending_plans/{spending_plan_guid} | Read a spending plan for a user
[**UpdateSpendingPlanIterationItem**](SpendingPlanAPI.md#UpdateSpendingPlanIterationItem) | **Put** /users/{user_guid}/spending_plans/{spending_plan_guid}/iterations/current/iteration_items/{iteration_item_guid} | Update a spending plan iteration item



## CreateSpendingPlan

> SpendingPlanResponse CreateSpendingPlan(ctx, userGuid).Execute()

Create spending plan



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendingPlanAPI.CreateSpendingPlan(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.CreateSpendingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpendingPlan`: SpendingPlanResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendingPlanAPI.CreateSpendingPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpendingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpendingPlanResponse**](SpendingPlanResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpendingPlanIterationItem

> SpendingPlanIterationItemResponse CreateSpendingPlanIterationItem(ctx, spendingPlanGuid, userGuid).SpendingPlanIterationItemCreateRequestBody(spendingPlanIterationItemCreateRequestBody).Execute()

Create spending plan iteration item



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
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    spendingPlanIterationItemCreateRequestBody := *openapiclient.NewSpendingPlanIterationItemCreateRequestBody(float32(110)) // SpendingPlanIterationItemCreateRequestBody | Iteration item to be created with required parameters (planned_amount)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendingPlanAPI.CreateSpendingPlanIterationItem(context.Background(), spendingPlanGuid, userGuid).SpendingPlanIterationItemCreateRequestBody(spendingPlanIterationItemCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.CreateSpendingPlanIterationItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpendingPlanIterationItem`: SpendingPlanIterationItemResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendingPlanAPI.CreateSpendingPlanIterationItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpendingPlanIterationItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **spendingPlanIterationItemCreateRequestBody** | [**SpendingPlanIterationItemCreateRequestBody**](SpendingPlanIterationItemCreateRequestBody.md) | Iteration item to be created with required parameters (planned_amount) | 

### Return type

[**SpendingPlanIterationItemResponse**](SpendingPlanIterationItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpendingPlan

> DeleteSpendingPlan(ctx, userGuid, spendingPlanGuid).Execute()

Delete spending plan



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique ID for a `user`.
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpendingPlanAPI.DeleteSpendingPlan(context.Background(), userGuid, spendingPlanGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.DeleteSpendingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique ID for a &#x60;user&#x60;. | 
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpendingPlanRequest struct via the builder pattern


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


## DeleteSpendingPlanAccount

> DeleteSpendingPlanAccount(ctx, userGuid, spendingPlanGuid, spendingPlanAccountGuid).Execute()

Delete spending plan account



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique ID for a `user`.
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.
    spendingPlanAccountGuid := "ACT-e9f80fee-84da-7s7r-9a5e-0346g4279b4c" // string | The unique ID for the specified account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpendingPlanAPI.DeleteSpendingPlanAccount(context.Background(), userGuid, spendingPlanGuid, spendingPlanAccountGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.DeleteSpendingPlanAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique ID for a &#x60;user&#x60;. | 
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 
**spendingPlanAccountGuid** | **string** | The unique ID for the specified account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpendingPlanAccountRequest struct via the builder pattern


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


## DeleteSpendingPlanIterationItem

> DeleteSpendingPlanIterationItem(ctx, userGuid, spendingPlanGuid, iterationItemGuid).Execute()

Delete spending plan iteration item



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique ID for a `user`.
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.
    iterationItemGuid := "SII-a4dc1549-da28-1245-9c9c-53eee4cdfbe3" // string | The unique ID for the `iteration_item`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpendingPlanAPI.DeleteSpendingPlanIterationItem(context.Background(), userGuid, spendingPlanGuid, iterationItemGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.DeleteSpendingPlanIterationItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique ID for a &#x60;user&#x60;. | 
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 
**iterationItemGuid** | **string** | The unique ID for the &#x60;iteration_item&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpendingPlanIterationItemRequest struct via the builder pattern


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


## ListSpendingPlanAccounts

> SpendingPlanAccountsResponse ListSpendingPlanAccounts(ctx, userGuid, spendingPlanGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List spending plan accounts



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendingPlanAPI.ListSpendingPlanAccounts(context.Background(), userGuid, spendingPlanGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.ListSpendingPlanAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpendingPlanAccounts`: SpendingPlanAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendingPlanAPI.ListSpendingPlanAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpendingPlanAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**SpendingPlanAccountsResponse**](SpendingPlanAccountsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpendingPlanIterationItems

> SpendingPlanIterationItemsResponseBody ListSpendingPlanIterationItems(ctx, userGuid, spendingPlanGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List spending plan iteration items



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendingPlanAPI.ListSpendingPlanIterationItems(context.Background(), userGuid, spendingPlanGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.ListSpendingPlanIterationItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpendingPlanIterationItems`: SpendingPlanIterationItemsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `SpendingPlanAPI.ListSpendingPlanIterationItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpendingPlanIterationItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**SpendingPlanIterationItemsResponseBody**](SpendingPlanIterationItemsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpendingPlanIterations

> SpendingPlanIterationsResponse ListSpendingPlanIterations(ctx, userGuid, spendingPlanGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List spending plan iterations



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendingPlanAPI.ListSpendingPlanIterations(context.Background(), userGuid, spendingPlanGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.ListSpendingPlanIterations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpendingPlanIterations`: SpendingPlanIterationsResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendingPlanAPI.ListSpendingPlanIterations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpendingPlanIterationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**SpendingPlanIterationsResponse**](SpendingPlanIterationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpendingPlans

> SpendingPlansResponseBody ListSpendingPlans(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List spending plans



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendingPlanAPI.ListSpendingPlans(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.ListSpendingPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpendingPlans`: SpendingPlansResponseBody
    fmt.Fprintf(os.Stdout, "Response from `SpendingPlanAPI.ListSpendingPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpendingPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**SpendingPlansResponseBody**](SpendingPlansResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSpendingPlanAccount

> SpendingPlanAccountResponse ReadSpendingPlanAccount(ctx, userGuid, spendingPlanGuid, spendingPlanAccountGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

Read spending plan account



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.
    spendingPlanAccountGuid := "ACT-e9f80fee-84da-7s7r-9a5e-0346g4279b4c" // string | The unique ID for the specified account.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendingPlanAPI.ReadSpendingPlanAccount(context.Background(), userGuid, spendingPlanGuid, spendingPlanAccountGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.ReadSpendingPlanAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSpendingPlanAccount`: SpendingPlanAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendingPlanAPI.ReadSpendingPlanAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 
**spendingPlanAccountGuid** | **string** | The unique ID for the specified account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSpendingPlanAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**SpendingPlanAccountResponse**](SpendingPlanAccountResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSpendingPlanIteration

> SpendingPlanIterationResponse ReadSpendingPlanIteration(ctx, userGuid, spendingPlanGuid, iterationNumber).Page(page).RecordsPerPage(recordsPerPage).Execute()

Read a spending plan iteration



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.
    iterationNumber := int32(1) // int32 | The current iteration number for the spending plan `iteration``.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendingPlanAPI.ReadSpendingPlanIteration(context.Background(), userGuid, spendingPlanGuid, iterationNumber).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.ReadSpendingPlanIteration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSpendingPlanIteration`: SpendingPlanIterationResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendingPlanAPI.ReadSpendingPlanIteration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 
**iterationNumber** | **int32** | The current iteration number for the spending plan &#x60;iteration&#x60;&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSpendingPlanIterationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**SpendingPlanIterationResponse**](SpendingPlanIterationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSpendingPlanIterationItem

> SpendingPlanIterationItemResponse ReadSpendingPlanIterationItem(ctx, userGuid, spendingPlanGuid, iterationItemGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

Read a spending plan iteration item



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.
    iterationItemGuid := "SII-a4dc1549-da28-1245-9c9c-53eee4cdfbe3" // string | The unique ID for the `iteration_item`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendingPlanAPI.ReadSpendingPlanIterationItem(context.Background(), userGuid, spendingPlanGuid, iterationItemGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.ReadSpendingPlanIterationItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSpendingPlanIterationItem`: SpendingPlanIterationItemResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendingPlanAPI.ReadSpendingPlanIterationItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 
**iterationItemGuid** | **string** | The unique ID for the &#x60;iteration_item&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSpendingPlanIterationItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**SpendingPlanIterationItemResponse**](SpendingPlanIterationItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSpendingPlanUser

> SpendingPlanResponse ReadSpendingPlanUser(ctx, userGuid, spendingPlanGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

Read a spending plan for a user



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendingPlanAPI.ReadSpendingPlanUser(context.Background(), userGuid, spendingPlanGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.ReadSpendingPlanUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSpendingPlanUser`: SpendingPlanResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendingPlanAPI.ReadSpendingPlanUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSpendingPlanUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**SpendingPlanResponse**](SpendingPlanResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpendingPlanIterationItem

> SpendingPlanIterationItemResponse UpdateSpendingPlanIterationItem(ctx, userGuid, spendingPlanGuid, iterationItemGuid).SpendingPlanIterationItemCreateRequestBody(spendingPlanIterationItemCreateRequestBody).Execute()

Update a spending plan iteration item



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    spendingPlanGuid := "SPL-e5f9a5bd-c5b3-4901-bdc0-62119b9db262" // string | The unique ID for the `spending_plan`.
    iterationItemGuid := "SII-a4dc1549-da28-1245-9c9c-53eee4cdfbe3" // string | The unique ID for the `iteration_item`.
    spendingPlanIterationItemCreateRequestBody := *openapiclient.NewSpendingPlanIterationItemCreateRequestBody(float32(110)) // SpendingPlanIterationItemCreateRequestBody | Iteration item object to be updated with required parameter (iteration_item_guid)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpendingPlanAPI.UpdateSpendingPlanIterationItem(context.Background(), userGuid, spendingPlanGuid, iterationItemGuid).SpendingPlanIterationItemCreateRequestBody(spendingPlanIterationItemCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpendingPlanAPI.UpdateSpendingPlanIterationItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpendingPlanIterationItem`: SpendingPlanIterationItemResponse
    fmt.Fprintf(os.Stdout, "Response from `SpendingPlanAPI.UpdateSpendingPlanIterationItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**spendingPlanGuid** | **string** | The unique ID for the &#x60;spending_plan&#x60;. | 
**iterationItemGuid** | **string** | The unique ID for the &#x60;iteration_item&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpendingPlanIterationItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **spendingPlanIterationItemCreateRequestBody** | [**SpendingPlanIterationItemCreateRequestBody**](SpendingPlanIterationItemCreateRequestBody.md) | Iteration item object to be updated with required parameter (iteration_item_guid) | 

### Return type

[**SpendingPlanIterationItemResponse**](SpendingPlanIterationItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

