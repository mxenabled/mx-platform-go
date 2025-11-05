# \GoalsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGoal**](GoalsAPI.md#CreateGoal) | **Post** /users/{user_guid}/goals | Create a goal
[**DeleteGoal**](GoalsAPI.md#DeleteGoal) | **Delete** /users/{user_guid}/goals/{goal_guid} | Delete a goal
[**ListGoals**](GoalsAPI.md#ListGoals) | **Get** /users/{user_guid}/goals | List goals
[**ReadGoal**](GoalsAPI.md#ReadGoal) | **Get** /users/{user_guid}/goals/{goal_guid} | Read a goal
[**RepositionGoals**](GoalsAPI.md#RepositionGoals) | **Put** /users/{user_guid}/goals/reposition | Reposition goals
[**UpdateGoal**](GoalsAPI.md#UpdateGoal) | **Put** /users/{user_guid}/goals/{goal_guid} | Update a goal



## CreateGoal

> GoalResponseBody CreateGoal(ctx, userGuid).GoalRequestBody(goalRequestBody).Execute()

Create a goal



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
    goalRequestBody := *openapiclient.NewGoalRequestBody() // GoalRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoalsAPI.CreateGoal(context.Background(), userGuid).GoalRequestBody(goalRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.CreateGoal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGoal`: GoalResponseBody
    fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.CreateGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **goalRequestBody** | [**GoalRequestBody**](GoalRequestBody.md) |  | 

### Return type

[**GoalResponseBody**](GoalResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/vnd.mx.api.v1+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGoal

> DeleteGoal(ctx, goalGuid, userGuid).Accept(accept).Execute()

Delete a goal



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
    goalGuid := "goalGuid_example" // string | The unique identifier for a goal. Defined by MX.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    accept := "application/vnd.mx.api.v1+json" // string | Specifies the media type expected in the response.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GoalsAPI.DeleteGoal(context.Background(), goalGuid, userGuid).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.DeleteGoal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**goalGuid** | **string** | The unique identifier for a goal. Defined by MX. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** | Specifies the media type expected in the response. | 

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


## ListGoals

> GoalsResponseBody ListGoals(ctx, userGuid).Accept(accept).Page(page).RecordsPerPage(recordsPerPage).Execute()

List goals



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
    accept := "application/vnd.mx.api.v1+json" // string | Specifies the media type expected in the response.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `1000`. If the value exceeds `1000`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoalsAPI.ListGoals(context.Background(), userGuid).Accept(accept).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.ListGoals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGoals`: GoalsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.ListGoals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGoalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Specifies the media type expected in the response. | 

 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**GoalsResponseBody**](GoalsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGoal

> GoalResponseBody ReadGoal(ctx, goalGuid, userGuid).Execute()

Read a goal



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
    goalGuid := "goalGuid_example" // string | The unique identifier for a goal. Defined by MX.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoalsAPI.ReadGoal(context.Background(), goalGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.ReadGoal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadGoal`: GoalResponseBody
    fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.ReadGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**goalGuid** | **string** | The unique identifier for a goal. Defined by MX. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GoalResponseBody**](GoalResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositionGoals

> RepositionResponseBody RepositionGoals(ctx, userGuid).RepositionRequestBody(repositionRequestBody).Execute()

Reposition goals



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
    repositionRequestBody := *openapiclient.NewRepositionRequestBody() // RepositionRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoalsAPI.RepositionGoals(context.Background(), userGuid).RepositionRequestBody(repositionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.RepositionGoals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositionGoals`: RepositionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.RepositionGoals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositionGoalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositionRequestBody** | [**RepositionRequestBody**](RepositionRequestBody.md) |  | 

### Return type

[**RepositionResponseBody**](RepositionResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/vnd.mx.api.v1+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGoal

> GoalResponseBody UpdateGoal(ctx, goalGuid, userGuid).UpdateGoalRequestBody(updateGoalRequestBody).Execute()

Update a goal



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
    goalGuid := "goalGuid_example" // string | The unique identifier for a goal. Defined by MX.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    updateGoalRequestBody := *openapiclient.NewUpdateGoalRequestBody() // UpdateGoalRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoalsAPI.UpdateGoal(context.Background(), goalGuid, userGuid).UpdateGoalRequestBody(updateGoalRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.UpdateGoal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGoal`: GoalResponseBody
    fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.UpdateGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**goalGuid** | **string** | The unique identifier for a goal. Defined by MX. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGoalRequestBody** | [**UpdateGoalRequestBody**](UpdateGoalRequestBody.md) |  | 

### Return type

[**GoalResponseBody**](GoalResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

