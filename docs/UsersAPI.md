# \UsersAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersAPI.md#CreateUser) | **Post** /users | Create user
[**DeleteUser**](UsersAPI.md#DeleteUser) | **Delete** /users/{user_guid} | Delete user
[**ListUsers**](UsersAPI.md#ListUsers) | **Get** /users | List users
[**ReadUser**](UsersAPI.md#ReadUser) | **Get** /users/{user_guid} | Read user
[**UpdateUser**](UsersAPI.md#UpdateUser) | **Put** /users/{user_guid} | Update user



## CreateUser

> UserResponseBody CreateUser(ctx).UserCreateRequestBody(userCreateRequestBody).Execute()

Create user



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
    userCreateRequestBody := *openapiclient.NewUserCreateRequestBody() // UserCreateRequestBody | User object to be created. (None of these parameters are required, but the user object cannot be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.CreateUser(context.Background()).UserCreateRequestBody(userCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: UserResponseBody
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userCreateRequestBody** | [**UserCreateRequestBody**](UserCreateRequestBody.md) | User object to be created. (None of these parameters are required, but the user object cannot be empty) | 

### Return type

[**UserResponseBody**](UserResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userGuid).Accept(accept).Execute()

Delete user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.DeleteUser(context.Background(), userGuid).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## ListUsers

> UsersResponseBody ListUsers(ctx).Page(page).RecordsPerPage(recordsPerPage).Id(id).Email(email).IsDisabled(isDisabled).Execute()

List users



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
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `1000`. If the value exceeds `1000`, the default value of `25` will be used instead. (optional)
    id := "u-12324-abdc" // string | The user `id` to search for. (optional)
    email := "example@example.com" // string | The user `email` to search for. (optional)
    isDisabled := true // bool | Search for users that are diabled. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.ListUsers(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Id(id).Email(email).IsDisabled(isDisabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: UsersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 
 **id** | **string** | The user &#x60;id&#x60; to search for. | 
 **email** | **string** | The user &#x60;email&#x60; to search for. | 
 **isDisabled** | **bool** | Search for users that are diabled. | 

### Return type

[**UsersResponseBody**](UsersResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUser

> UserResponseBody ReadUser(ctx, userGuid).Execute()

Read user



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
    resp, r, err := apiClient.UsersAPI.ReadUser(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ReadUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUser`: UserResponseBody
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ReadUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserResponseBody**](UserResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UserResponseBody UpdateUser(ctx, userGuid).UserUpdateRequestBody(userUpdateRequestBody).Execute()

Update user



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
    userUpdateRequestBody := *openapiclient.NewUserUpdateRequestBody() // UserUpdateRequestBody | User object to be updated (None of these parameters are required, but the user object cannot be empty.)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.UpdateUser(context.Background(), userGuid).UserUpdateRequestBody(userUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: UserResponseBody
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUpdateRequestBody** | [**UserUpdateRequestBody**](UserUpdateRequestBody.md) | User object to be updated (None of these parameters are required, but the user object cannot be empty.) | 

### Return type

[**UserResponseBody**](UserResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

