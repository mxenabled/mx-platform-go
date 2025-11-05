# \StatementsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadStatementPDF**](StatementsAPI.md#DownloadStatementPDF) | **Get** /users/{user_guid}/members/{member_guid}/statements/{statement_guid}.pdf | Download statement pdf
[**FetchStatements**](StatementsAPI.md#FetchStatements) | **Post** /users/{user_guid}/members/{member_guid}/fetch_statements | Fetch statements
[**ListStatementsByMember**](StatementsAPI.md#ListStatementsByMember) | **Get** /users/{user_guid}/members/{member_guid}/statements | List statements by member
[**ReadStatementByMember**](StatementsAPI.md#ReadStatementByMember) | **Get** /users/{user_guid}/members/{member_guid}/statements/{statement_guid} | Read statement by member



## DownloadStatementPDF

> *os.File DownloadStatementPDF(ctx, memberGuid, statementGuid, userGuid).Execute()

Download statement pdf



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
    statementGuid := "STA-737a344b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for a `statement`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatementsAPI.DownloadStatementPDF(context.Background(), memberGuid, statementGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.DownloadStatementPDF``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadStatementPDF`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.DownloadStatementPDF`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**statementGuid** | **string** | The unique id for a &#x60;statement&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadStatementPDFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStatements

> MemberResponseBody FetchStatements(ctx, memberGuid, userGuid).Execute()

Fetch statements



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
    resp, r, err := apiClient.StatementsAPI.FetchStatements(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.FetchStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchStatements`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.FetchStatements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchStatementsRequest struct via the builder pattern


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


## ListStatementsByMember

> StatementsResponseBody ListStatementsByMember(ctx, memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List statements by member



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatementsAPI.ListStatementsByMember(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.ListStatementsByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStatementsByMember`: StatementsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.ListStatementsByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStatementsByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**StatementsResponseBody**](StatementsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadStatementByMember

> StatementResponseBody ReadStatementByMember(ctx, memberGuid, statementGuid, userGuid).Execute()

Read statement by member



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
    statementGuid := "STA-737a344b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for a `statement`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatementsAPI.ReadStatementByMember(context.Background(), memberGuid, statementGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.ReadStatementByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadStatementByMember`: StatementResponseBody
    fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.ReadStatementByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**statementGuid** | **string** | The unique id for a &#x60;statement&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadStatementByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StatementResponseBody**](StatementResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

