# \GoalsAPI

All URIs are relative to *https://api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersUserGuidGoalsGet**](GoalsAPI.md#UsersUserGuidGoalsGet) | **Get** /users/{user_guid}/goals | List goals
[**UsersUserGuidGoalsGoalGuidDelete**](GoalsAPI.md#UsersUserGuidGoalsGoalGuidDelete) | **Delete** /users/{user_guid}/goals/{goal_guid} | Delete a goal
[**UsersUserGuidGoalsGoalGuidGet**](GoalsAPI.md#UsersUserGuidGoalsGoalGuidGet) | **Get** /users/{user_guid}/goals/{goal_guid} | Read a goal
[**UsersUserGuidGoalsGoalGuidPut**](GoalsAPI.md#UsersUserGuidGoalsGoalGuidPut) | **Put** /users/{user_guid}/goals/{goal_guid} | Update a goal
[**UsersUserGuidGoalsPost**](GoalsAPI.md#UsersUserGuidGoalsPost) | **Post** /users/{user_guid}/goals | Create a goal
[**UsersUserGuidGoalsRepositionPut**](GoalsAPI.md#UsersUserGuidGoalsRepositionPut) | **Put** /users/{user_guid}/goals/reposition | Reposition goals



## UsersUserGuidGoalsGet

> GoalsResponseBody UsersUserGuidGoalsGet(ctx, userGuid).Page(page).RecordsPerAge(recordsPerAge).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for the user.
    page := "page_example" // string | Results are returned in paginated sets, this is the page of the results you would like to view. Defaults to page 1 if no page is specified. (optional)
    recordsPerAge := "recordsPerAge_example" // string | The supported range is from 10 to 1000. If the records_per_page parameter is not specified or is outside this range, a default of 25 records per page will be used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoalsAPI.UsersUserGuidGoalsGet(context.Background(), userGuid).Page(page).RecordsPerAge(recordsPerAge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.UsersUserGuidGoalsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidGoalsGet`: GoalsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.UsersUserGuidGoalsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidGoalsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Results are returned in paginated sets, this is the page of the results you would like to view. Defaults to page 1 if no page is specified. | 
 **recordsPerAge** | **string** | The supported range is from 10 to 1000. If the records_per_page parameter is not specified or is outside this range, a default of 25 records per page will be used. | 

### Return type

[**GoalsResponseBody**](GoalsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserGuidGoalsGoalGuidDelete

> UsersUserGuidGoalsGoalGuidDelete(ctx, goalGuid, userGuid).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for a user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GoalsAPI.UsersUserGuidGoalsGoalGuidDelete(context.Background(), goalGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.UsersUserGuidGoalsGoalGuidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**goalGuid** | **string** | The unique identifier for a goal. Defined by MX. | 
**userGuid** | **string** | The unique identifier for a user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidGoalsGoalGuidDeleteRequest struct via the builder pattern


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


## UsersUserGuidGoalsGoalGuidGet

> GoalResponseBody UsersUserGuidGoalsGoalGuidGet(ctx, goalGuid, userGuid).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for a user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoalsAPI.UsersUserGuidGoalsGoalGuidGet(context.Background(), goalGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.UsersUserGuidGoalsGoalGuidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidGoalsGoalGuidGet`: GoalResponseBody
    fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.UsersUserGuidGoalsGoalGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**goalGuid** | **string** | The unique identifier for a goal. Defined by MX. | 
**userGuid** | **string** | The unique identifier for a user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidGoalsGoalGuidGetRequest struct via the builder pattern


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


## UsersUserGuidGoalsGoalGuidPut

> GoalResponseBody UsersUserGuidGoalsGoalGuidPut(ctx, goalGuid, userGuid).UpdateGoalRequestBody(updateGoalRequestBody).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for a user.
    updateGoalRequestBody := *openapiclient.NewUpdateGoalRequestBody() // UpdateGoalRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoalsAPI.UsersUserGuidGoalsGoalGuidPut(context.Background(), goalGuid, userGuid).UpdateGoalRequestBody(updateGoalRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.UsersUserGuidGoalsGoalGuidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidGoalsGoalGuidPut`: GoalResponseBody
    fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.UsersUserGuidGoalsGoalGuidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**goalGuid** | **string** | The unique identifier for a goal. Defined by MX. | 
**userGuid** | **string** | The unique identifier for a user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidGoalsGoalGuidPutRequest struct via the builder pattern


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


## UsersUserGuidGoalsPost

> GoalResponseBody UsersUserGuidGoalsPost(ctx, userGuid).GoalRequestBody(goalRequestBody).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for the user.
    goalRequestBody := *openapiclient.NewGoalRequestBody() // GoalRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoalsAPI.UsersUserGuidGoalsPost(context.Background(), userGuid).GoalRequestBody(goalRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.UsersUserGuidGoalsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidGoalsPost`: GoalResponseBody
    fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.UsersUserGuidGoalsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidGoalsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **goalRequestBody** | [**GoalRequestBody**](GoalRequestBody.md) |  | 

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


## UsersUserGuidGoalsRepositionPut

> RepositionResponseBody UsersUserGuidGoalsRepositionPut(ctx, userGuid).RepositionRequestBody(repositionRequestBody).Execute()

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
    userGuid := "userGuid_example" // string | The unique identifier for the user.
    repositionRequestBody := *openapiclient.NewRepositionRequestBody() // RepositionRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoalsAPI.UsersUserGuidGoalsRepositionPut(context.Background(), userGuid).RepositionRequestBody(repositionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoalsAPI.UsersUserGuidGoalsRepositionPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidGoalsRepositionPut`: RepositionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `GoalsAPI.UsersUserGuidGoalsRepositionPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidGoalsRepositionPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositionRequestBody** | [**RepositionRequestBody**](RepositionRequestBody.md) |  | 

### Return type

[**RepositionResponseBody**](RepositionResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

