# \CategoriesAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCategory**](CategoriesAPI.md#CreateCategory) | **Post** /users/{user_guid}/categories | Create category
[**DeleteCategory**](CategoriesAPI.md#DeleteCategory) | **Delete** /users/{user_guid}/categories/{category_guid} | Delete category
[**ListCategories**](CategoriesAPI.md#ListCategories) | **Get** /users/{user_guid}/categories | List categories
[**ListDefaultCategories**](CategoriesAPI.md#ListDefaultCategories) | **Get** /categories/default | List default categories
[**ListDefaultCategoriesByUser**](CategoriesAPI.md#ListDefaultCategoriesByUser) | **Get** /users/{user_guid}/categories/default | List default categories by user
[**ReadCategory**](CategoriesAPI.md#ReadCategory) | **Get** /users/{user_guid}/categories/{category_guid} | Read a custom category
[**ReadDefaultCategory**](CategoriesAPI.md#ReadDefaultCategory) | **Get** /categories/{category_guid} | Read a default category
[**UpdateCategory**](CategoriesAPI.md#UpdateCategory) | **Put** /users/{user_guid}/categories/{category_guid} | Update category



## CreateCategory

> CategoryResponseBody CreateCategory(ctx, userGuid).CategoryCreateRequestBody(categoryCreateRequestBody).Execute()

Create category



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
    categoryCreateRequestBody := *openapiclient.NewCategoryCreateRequestBody() // CategoryCreateRequestBody | Custom category object to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesAPI.CreateCategory(context.Background(), userGuid).CategoryCreateRequestBody(categoryCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.CreateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.CreateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryCreateRequestBody** | [**CategoryCreateRequestBody**](CategoryCreateRequestBody.md) | Custom category object to be created | 

### Return type

[**CategoryResponseBody**](CategoryResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCategory

> DeleteCategory(ctx, categoryGuid, userGuid).Execute()

Delete category



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
    categoryGuid := "CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874" // string | The unique id for a `category`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CategoriesAPI.DeleteCategory(context.Background(), categoryGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.DeleteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryGuid** | **string** | The unique id for a &#x60;category&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCategoryRequest struct via the builder pattern


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


## ListCategories

> CategoriesResponseBody ListCategories(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List categories



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
    resp, r, err := apiClient.CategoriesAPI.ListCategories(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.ListCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCategories`: CategoriesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.ListCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

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


## ListDefaultCategories

> CategoriesResponseBody ListDefaultCategories(ctx).Page(page).RecordsPerPage(recordsPerPage).Execute()

List default categories



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
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesAPI.ListDefaultCategories(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.ListDefaultCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDefaultCategories`: CategoriesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.ListDefaultCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDefaultCategoriesRequest struct via the builder pattern


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


## ListDefaultCategoriesByUser

> CategoriesResponseBody ListDefaultCategoriesByUser(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List default categories by user



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
    resp, r, err := apiClient.CategoriesAPI.ListDefaultCategoriesByUser(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.ListDefaultCategoriesByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDefaultCategoriesByUser`: CategoriesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.ListDefaultCategoriesByUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDefaultCategoriesByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

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


## ReadCategory

> CategoryResponseBody ReadCategory(ctx, categoryGuid, userGuid).Execute()

Read a custom category



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
    categoryGuid := "CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874" // string | The unique id for a `category`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesAPI.ReadCategory(context.Background(), categoryGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.ReadCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.ReadCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryGuid** | **string** | The unique id for a &#x60;category&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CategoryResponseBody**](CategoryResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDefaultCategory

> CategoryResponseBody ReadDefaultCategory(ctx, categoryGuid).Execute()

Read a default category



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
    categoryGuid := "CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874" // string | The unique id for a `category`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesAPI.ReadDefaultCategory(context.Background(), categoryGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.ReadDefaultCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDefaultCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.ReadDefaultCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryGuid** | **string** | The unique id for a &#x60;category&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadDefaultCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CategoryResponseBody**](CategoryResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCategory

> CategoryResponseBody UpdateCategory(ctx, categoryGuid, userGuid).CategoryUpdateRequestBody(categoryUpdateRequestBody).Execute()

Update category



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
    categoryGuid := "CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874" // string | The unique id for a `category`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    categoryUpdateRequestBody := *openapiclient.NewCategoryUpdateRequestBody() // CategoryUpdateRequestBody | Category object to be updated (While no single parameter is required, the `category` object cannot be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesAPI.UpdateCategory(context.Background(), categoryGuid, userGuid).CategoryUpdateRequestBody(categoryUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.UpdateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.UpdateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryGuid** | **string** | The unique id for a &#x60;category&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **categoryUpdateRequestBody** | [**CategoryUpdateRequestBody**](CategoryUpdateRequestBody.md) | Category object to be updated (While no single parameter is required, the &#x60;category&#x60; object cannot be empty) | 

### Return type

[**CategoryResponseBody**](CategoryResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

