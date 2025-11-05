# \AccountsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManualAccount**](AccountsAPI.md#CreateManualAccount) | **Post** /users/{user_guid}/accounts | Create manual account
[**DeleteManualAccount**](AccountsAPI.md#DeleteManualAccount) | **Delete** /users/{user_guid}/accounts/{account_guid} | Delete manual account
[**ListAccountNumbersByAccount**](AccountsAPI.md#ListAccountNumbersByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/account_numbers | List account numbers by account
[**ListAccountNumbersByMember**](AccountsAPI.md#ListAccountNumbersByMember) | **Get** /users/{user_guid}/members/{member_guid}/account_numbers | List account numbers by member
[**ListAccountOwnersByMember**](AccountsAPI.md#ListAccountOwnersByMember) | **Get** /users/{user_guid}/members/{member_guid}/account_owners | List account owners by member
[**ListMemberAccounts**](AccountsAPI.md#ListMemberAccounts) | **Get** /users/{user_guid}/members/{member_guid}/accounts | List accounts by member
[**ListUserAccounts**](AccountsAPI.md#ListUserAccounts) | **Get** /users/{user_guid}/accounts | List accounts
[**ReadAccount**](AccountsAPI.md#ReadAccount) | **Get** /users/{user_guid}/accounts/{account_guid} | Read account
[**ReadAccountByMember**](AccountsAPI.md#ReadAccountByMember) | **Get** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Read account by member
[**UpdateAccountByMember**](AccountsAPI.md#UpdateAccountByMember) | **Put** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Update account by member



## CreateManualAccount

> AccountResponseBody CreateManualAccount(ctx, userGuid).AccountCreateRequestBody(accountCreateRequestBody).Execute()

Create manual account



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
    accountCreateRequestBody := *openapiclient.NewAccountCreateRequestBody() // AccountCreateRequestBody | Manual account object to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.CreateManualAccount(context.Background(), userGuid).AccountCreateRequestBody(accountCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.CreateManualAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManualAccount`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.CreateManualAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateManualAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountCreateRequestBody** | [**AccountCreateRequestBody**](AccountCreateRequestBody.md) | Manual account object to be created. | 

### Return type

[**AccountResponseBody**](AccountResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManualAccount

> DeleteManualAccount(ctx, accountGuid, userGuid).Accept(accept).Execute()

Delete manual account



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
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    accept := "application/vnd.mx.api.v1+json" // string | Specifies the media type expected in the response.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountsAPI.DeleteManualAccount(context.Background(), accountGuid, userGuid).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DeleteManualAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManualAccountRequest struct via the builder pattern


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


## ListAccountNumbersByAccount

> AccountNumbersResponseBody ListAccountNumbersByAccount(ctx, accountGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List account numbers by account



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
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `1000`. If the value exceeds `1000`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.ListAccountNumbersByAccount(context.Background(), accountGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccountNumbersByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountNumbersByAccount`: AccountNumbersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccountNumbersByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountNumbersByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**AccountNumbersResponseBody**](AccountNumbersResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountNumbersByMember

> AccountNumbersResponseBody ListAccountNumbersByMember(ctx, memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List account numbers by member



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
    resp, r, err := apiClient.AccountsAPI.ListAccountNumbersByMember(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccountNumbersByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountNumbersByMember`: AccountNumbersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccountNumbersByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountNumbersByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**AccountNumbersResponseBody**](AccountNumbersResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountOwnersByMember

> AccountOwnersResponseBody ListAccountOwnersByMember(ctx, memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List account owners by member



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
    resp, r, err := apiClient.AccountsAPI.ListAccountOwnersByMember(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccountOwnersByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountOwnersByMember`: AccountOwnersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccountOwnersByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountOwnersByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**AccountOwnersResponseBody**](AccountOwnersResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMemberAccounts

> AccountsResponseBody ListMemberAccounts(ctx, userGuid, memberGuid).MemberIsManagedByUser(memberIsManagedByUser).Page(page).RecordsPerPage(recordsPerPage).Execute()

List accounts by member



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
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    memberIsManagedByUser := true // bool | List only accounts whose member is managed by the user. (optional)
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `1000`. If the value exceeds `1000`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.ListMemberAccounts(context.Background(), userGuid, memberGuid).MemberIsManagedByUser(memberIsManagedByUser).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListMemberAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMemberAccounts`: AccountsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListMemberAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMemberAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **memberIsManagedByUser** | **bool** | List only accounts whose member is managed by the user. | 
 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**AccountsResponseBody**](AccountsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserAccounts

> AccountsResponseBody ListUserAccounts(ctx, userGuid).Page(page).MemberIsManagedByUser(memberIsManagedByUser).IsManual(isManual).RecordsPerPage(recordsPerPage).UseCase(useCase).Execute()

List accounts



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
    memberIsManagedByUser := true // bool | List only accounts whose member is managed by the user. (optional)
    isManual := true // bool | List only accounts that were manually created. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `1000`. If the value exceeds `1000`, the default value of `25` will be used instead. (optional)
    useCase := "MONEY_MOVEMENT" // string | The use case associated with the member. Valid values are `PFM` and `MONEY_MOVEMENT`. For example, you can append either `?use_case=PFM` or `?use_case=MONEY_MOVEMENT`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.ListUserAccounts(context.Background(), userGuid).Page(page).MemberIsManagedByUser(memberIsManagedByUser).IsManual(isManual).RecordsPerPage(recordsPerPage).UseCase(useCase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListUserAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserAccounts`: AccountsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListUserAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results are paginated. Specify current page. | 
 **memberIsManagedByUser** | **bool** | List only accounts whose member is managed by the user. | 
 **isManual** | **bool** | List only accounts that were manually created. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 
 **useCase** | **string** | The use case associated with the member. Valid values are &#x60;PFM&#x60; and &#x60;MONEY_MOVEMENT&#x60;. For example, you can append either &#x60;?use_case&#x3D;PFM&#x60; or &#x60;?use_case&#x3D;MONEY_MOVEMENT&#x60;. | 

### Return type

[**AccountsResponseBody**](AccountsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccount

> AccountResponseBody ReadAccount(ctx, accountGuid, userGuid).Execute()

Read account



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
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.ReadAccount(context.Background(), accountGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ReadAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAccount`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ReadAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccountResponseBody**](AccountResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccountByMember

> AccountResponseBody ReadAccountByMember(ctx, accountGuid, memberGuid, userGuid).Execute()

Read account by member



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
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.ReadAccountByMember(context.Background(), accountGuid, memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ReadAccountByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAccountByMember`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ReadAccountByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAccountByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AccountResponseBody**](AccountResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountByMember

> AccountResponseBody UpdateAccountByMember(ctx, accountGuid, memberGuid, userGuid).AccountUpdateRequestBody(accountUpdateRequestBody).Execute()

Update account by member



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
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    accountUpdateRequestBody := *openapiclient.NewAccountUpdateRequestBody() // AccountUpdateRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.UpdateAccountByMember(context.Background(), accountGuid, memberGuid, userGuid).AccountUpdateRequestBody(accountUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UpdateAccountByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountByMember`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.UpdateAccountByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accountUpdateRequestBody** | [**AccountUpdateRequestBody**](AccountUpdateRequestBody.md) |  | 

### Return type

[**AccountResponseBody**](AccountResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

