# \TaggingsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTagging**](TaggingsAPI.md#CreateTagging) | **Post** /users/{user_guid}/taggings | Create tagging
[**DeleteTagging**](TaggingsAPI.md#DeleteTagging) | **Delete** /users/{user_guid}/taggings/{tagging_guid} | Delete tagging
[**ListTaggings**](TaggingsAPI.md#ListTaggings) | **Get** /users/{user_guid}/taggings | List taggings
[**ReadTagging**](TaggingsAPI.md#ReadTagging) | **Get** /users/{user_guid}/taggings/{tagging_guid} | Read tagging
[**UpdateTagging**](TaggingsAPI.md#UpdateTagging) | **Put** /users/{user_guid}/taggings/{tagging_guid} | Update tagging



## CreateTagging

> TaggingResponseBody CreateTagging(ctx, userGuid).TaggingCreateRequestBody(taggingCreateRequestBody).Execute()

Create tagging



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
    taggingCreateRequestBody := *openapiclient.NewTaggingCreateRequestBody() // TaggingCreateRequestBody | Tagging object to be created with required parameters (tag_guid and transaction_guid)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaggingsAPI.CreateTagging(context.Background(), userGuid).TaggingCreateRequestBody(taggingCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggingsAPI.CreateTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagging`: TaggingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TaggingsAPI.CreateTagging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taggingCreateRequestBody** | [**TaggingCreateRequestBody**](TaggingCreateRequestBody.md) | Tagging object to be created with required parameters (tag_guid and transaction_guid) | 

### Return type

[**TaggingResponseBody**](TaggingResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagging

> DeleteTagging(ctx, taggingGuid, userGuid).Execute()

Delete tagging



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
    taggingGuid := "TGN-007f5486-17e1-45fc-8b87-8f03984430fe" // string | The unique id for a `tagging`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaggingsAPI.DeleteTagging(context.Background(), taggingGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggingsAPI.DeleteTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taggingGuid** | **string** | The unique id for a &#x60;tagging&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaggingRequest struct via the builder pattern


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


## ListTaggings

> TaggingsResponseBody ListTaggings(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List taggings



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
    resp, r, err := apiClient.TaggingsAPI.ListTaggings(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggingsAPI.ListTaggings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaggings`: TaggingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TaggingsAPI.ListTaggings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTaggingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**TaggingsResponseBody**](TaggingsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTagging

> TaggingResponseBody ReadTagging(ctx, taggingGuid, userGuid).Execute()

Read tagging



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
    taggingGuid := "TGN-007f5486-17e1-45fc-8b87-8f03984430fe" // string | The unique id for a `tagging`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaggingsAPI.ReadTagging(context.Background(), taggingGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggingsAPI.ReadTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTagging`: TaggingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TaggingsAPI.ReadTagging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taggingGuid** | **string** | The unique id for a &#x60;tagging&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTaggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TaggingResponseBody**](TaggingResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagging

> TaggingResponseBody UpdateTagging(ctx, taggingGuid, userGuid).TaggingUpdateRequestBody(taggingUpdateRequestBody).Execute()

Update tagging



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
    taggingGuid := "TGN-007f5486-17e1-45fc-8b87-8f03984430fe" // string | The unique id for a `tagging`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    taggingUpdateRequestBody := *openapiclient.NewTaggingUpdateRequestBody() // TaggingUpdateRequestBody | Tagging object to be updated with required parameter (tag_guid)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaggingsAPI.UpdateTagging(context.Background(), taggingGuid, userGuid).TaggingUpdateRequestBody(taggingUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggingsAPI.UpdateTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagging`: TaggingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TaggingsAPI.UpdateTagging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taggingGuid** | **string** | The unique id for a &#x60;tagging&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **taggingUpdateRequestBody** | [**TaggingUpdateRequestBody**](TaggingUpdateRequestBody.md) | Tagging object to be updated with required parameter (tag_guid) | 

### Return type

[**TaggingResponseBody**](TaggingResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

