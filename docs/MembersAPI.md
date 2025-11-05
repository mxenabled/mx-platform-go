# \MembersAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateMember**](MembersAPI.md#AggregateMember) | **Post** /users/{user_guid}/members/{member_guid}/aggregate | Aggregate member
[**CheckBalances**](MembersAPI.md#CheckBalances) | **Post** /users/{user_guid}/members/{member_guid}/check_balance | Check balances
[**CreateMember**](MembersAPI.md#CreateMember) | **Post** /users/{user_guid}/members | Create member
[**DeleteMember**](MembersAPI.md#DeleteMember) | **Delete** /users/{user_guid}/members/{member_guid} | Delete member
[**IdentifyMember**](MembersAPI.md#IdentifyMember) | **Post** /users/{user_guid}/members/{member_guid}/identify | Identify member
[**ListMemberChallenges**](MembersAPI.md#ListMemberChallenges) | **Get** /users/{user_guid}/members/{member_guid}/challenges | List member challenges
[**ListMemberCredentials**](MembersAPI.md#ListMemberCredentials) | **Get** /users/{user_guid}/members/{member_guid}/credentials | List member credentials
[**ListMembers**](MembersAPI.md#ListMembers) | **Get** /users/{user_guid}/members | List members
[**ReadMember**](MembersAPI.md#ReadMember) | **Get** /users/{user_guid}/members/{member_guid} | Read member
[**ReadMemberStatus**](MembersAPI.md#ReadMemberStatus) | **Get** /users/{user_guid}/members/{member_guid}/status | Read member status
[**ResumeAggregation**](MembersAPI.md#ResumeAggregation) | **Put** /users/{user_guid}/members/{member_guid}/resume | Resume aggregation
[**UpdateMember**](MembersAPI.md#UpdateMember) | **Put** /users/{user_guid}/members/{member_guid} | Update member
[**VerifyMember**](MembersAPI.md#VerifyMember) | **Post** /users/{user_guid}/members/{member_guid}/verify | Verify member



## AggregateMember

> MemberResponseBody AggregateMember(ctx, memberGuid, userGuid).XCALLBACKPAYLOAD(xCALLBACKPAYLOAD).IncludeHoldings(includeHoldings).IncludeTransactions(includeTransactions).Execute()

Aggregate member



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
    xCALLBACKPAYLOAD := "813e50bd-4a7e-4517-b6bb-9eef65a68cbd" // string | The base64 encoded string defined in this header will be returned in the [Member](/resources/webhooks/member/) and [Member Data Updated](/resources/webhooks/member#member-data-updated) webhooks. This allows you to trace user interactions and workflows initiated externally and internally in the MX Platform. Max 1024 characters. (optional)
    includeHoldings := false // bool | When set to `false`, the aggregation will not gather holdings data. Defaults to `true`. (optional)
    includeTransactions := true // bool | When set to `false`, the aggregation will not gather transactions data. Defaults to `true`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.AggregateMember(context.Background(), memberGuid, userGuid).XCALLBACKPAYLOAD(xCALLBACKPAYLOAD).IncludeHoldings(includeHoldings).IncludeTransactions(includeTransactions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.AggregateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AggregateMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.AggregateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAggregateMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCALLBACKPAYLOAD** | **string** | The base64 encoded string defined in this header will be returned in the [Member](/resources/webhooks/member/) and [Member Data Updated](/resources/webhooks/member#member-data-updated) webhooks. This allows you to trace user interactions and workflows initiated externally and internally in the MX Platform. Max 1024 characters. | 
 **includeHoldings** | **bool** | When set to &#x60;false&#x60;, the aggregation will not gather holdings data. Defaults to &#x60;true&#x60;. | 
 **includeTransactions** | **bool** | When set to &#x60;false&#x60;, the aggregation will not gather transactions data. Defaults to &#x60;true&#x60;. | 

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


## CheckBalances

> MemberResponseBody CheckBalances(ctx, memberGuid, userGuid).Execute()

Check balances



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
    resp, r, err := apiClient.MembersAPI.CheckBalances(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.CheckBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckBalances`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.CheckBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckBalancesRequest struct via the builder pattern


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


## CreateMember

> MemberResponseBody CreateMember(ctx, userGuid).MemberCreateRequestBody(memberCreateRequestBody).XCALLBACKPAYLOAD(xCALLBACKPAYLOAD).Execute()

Create member



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
    memberCreateRequestBody := *openapiclient.NewMemberCreateRequestBody() // MemberCreateRequestBody | Member object to be created with optional parameters (id and metadata) and required parameters (credentials and institution_code)
    xCALLBACKPAYLOAD := "813e50bd-4a7e-4517-b6bb-9eef65a68cbd" // string | The base64 encoded string defined in this header will be returned in the [Member](/resources/webhooks/member/) and [Member Data Updated](/resources/webhooks/member#member-data-updated) webhooks. This allows you to trace user interactions and workflows initiated externally and internally in the MX Platform. Max 1024 characters. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.CreateMember(context.Background(), userGuid).MemberCreateRequestBody(memberCreateRequestBody).XCALLBACKPAYLOAD(xCALLBACKPAYLOAD).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.CreateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.CreateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memberCreateRequestBody** | [**MemberCreateRequestBody**](MemberCreateRequestBody.md) | Member object to be created with optional parameters (id and metadata) and required parameters (credentials and institution_code) | 
 **xCALLBACKPAYLOAD** | **string** | The base64 encoded string defined in this header will be returned in the [Member](/resources/webhooks/member/) and [Member Data Updated](/resources/webhooks/member#member-data-updated) webhooks. This allows you to trace user interactions and workflows initiated externally and internally in the MX Platform. Max 1024 characters. | 

### Return type

[**MemberResponseBody**](MemberResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMember

> DeleteMember(ctx, memberGuid, userGuid).Execute()

Delete member



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
    r, err := apiClient.MembersAPI.DeleteMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.DeleteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberRequest struct via the builder pattern


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


## IdentifyMember

> MemberResponseBody IdentifyMember(ctx, memberGuid, userGuid).Execute()

Identify member



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
    resp, r, err := apiClient.MembersAPI.IdentifyMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.IdentifyMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentifyMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.IdentifyMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentifyMemberRequest struct via the builder pattern


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


## ListMemberChallenges

> ChallengesResponseBody ListMemberChallenges(ctx, memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List member challenges



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
    resp, r, err := apiClient.MembersAPI.ListMemberChallenges(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.ListMemberChallenges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMemberChallenges`: ChallengesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.ListMemberChallenges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMemberChallengesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**ChallengesResponseBody**](ChallengesResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMemberCredentials

> CredentialsResponseBody ListMemberCredentials(ctx, memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List member credentials



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
    resp, r, err := apiClient.MembersAPI.ListMemberCredentials(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.ListMemberCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMemberCredentials`: CredentialsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.ListMemberCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMemberCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**CredentialsResponseBody**](CredentialsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMembers

> MembersResponseBody ListMembers(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).UseCase(useCase).Execute()

List members



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
    useCase := "MONEY_MOVEMENT" // string | The use case associated with the member. Valid values are `PFM` and `MONEY_MOVEMENT`. For example, you can append either `?use_case=PFM` or `?use_case=MONEY_MOVEMENT`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.ListMembers(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).UseCase(useCase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.ListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMembers`: MembersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.ListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 
 **useCase** | **string** | The use case associated with the member. Valid values are &#x60;PFM&#x60; and &#x60;MONEY_MOVEMENT&#x60;. For example, you can append either &#x60;?use_case&#x3D;PFM&#x60; or &#x60;?use_case&#x3D;MONEY_MOVEMENT&#x60;. | 

### Return type

[**MembersResponseBody**](MembersResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMember

> MemberResponseBody ReadMember(ctx, memberGuid, userGuid).Execute()

Read member



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
    resp, r, err := apiClient.MembersAPI.ReadMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.ReadMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.ReadMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadMemberRequest struct via the builder pattern


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


## ReadMemberStatus

> MemberStatusResponseBody ReadMemberStatus(ctx, memberGuid, userGuid).Execute()

Read member status



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
    resp, r, err := apiClient.MembersAPI.ReadMemberStatus(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.ReadMemberStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMemberStatus`: MemberStatusResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.ReadMemberStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadMemberStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MemberStatusResponseBody**](MemberStatusResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeAggregation

> MemberResponseBody ResumeAggregation(ctx, memberGuid, userGuid).MemberResumeRequestBody(memberResumeRequestBody).Execute()

Resume aggregation



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
    memberResumeRequestBody := *openapiclient.NewMemberResumeRequestBody() // MemberResumeRequestBody | Member object with MFA challenge answers

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.ResumeAggregation(context.Background(), memberGuid, userGuid).MemberResumeRequestBody(memberResumeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.ResumeAggregation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeAggregation`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.ResumeAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **memberResumeRequestBody** | [**MemberResumeRequestBody**](MemberResumeRequestBody.md) | Member object with MFA challenge answers | 

### Return type

[**MemberResponseBody**](MemberResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> MemberResponseBody UpdateMember(ctx, memberGuid, userGuid).MemberUpdateRequestBody(memberUpdateRequestBody).XCALLBACKPAYLOAD(xCALLBACKPAYLOAD).Execute()

Update member



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
    memberUpdateRequestBody := *openapiclient.NewMemberUpdateRequestBody() // MemberUpdateRequestBody | Member object to be updated (While no single parameter is required, the request body can't be empty)
    xCALLBACKPAYLOAD := "813e50bd-4a7e-4517-b6bb-9eef65a68cbd" // string | The base64 encoded string defined in this header will be returned in the [Member](/resources/webhooks/member/) and [Member Data Updated](/resources/webhooks/member#member-data-updated) webhooks. This allows you to trace user interactions and workflows initiated externally and internally in the MX Platform. Max 1024 characters. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.UpdateMember(context.Background(), memberGuid, userGuid).MemberUpdateRequestBody(memberUpdateRequestBody).XCALLBACKPAYLOAD(xCALLBACKPAYLOAD).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.UpdateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.UpdateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **memberUpdateRequestBody** | [**MemberUpdateRequestBody**](MemberUpdateRequestBody.md) | Member object to be updated (While no single parameter is required, the request body can&#39;t be empty) | 
 **xCALLBACKPAYLOAD** | **string** | The base64 encoded string defined in this header will be returned in the [Member](/resources/webhooks/member/) and [Member Data Updated](/resources/webhooks/member#member-data-updated) webhooks. This allows you to trace user interactions and workflows initiated externally and internally in the MX Platform. Max 1024 characters. | 

### Return type

[**MemberResponseBody**](MemberResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyMember

> MemberResponseBody VerifyMember(ctx, memberGuid, userGuid).XCALLBACKPAYLOAD(xCALLBACKPAYLOAD).Execute()

Verify member



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
    xCALLBACKPAYLOAD := "813e50bd-4a7e-4517-b6bb-9eef65a68cbd" // string | The base64 encoded string defined in this header will be returned in the [Member](/resources/webhooks/member/) and [Member Data Updated](/resources/webhooks/member#member-data-updated) webhooks. This allows you to trace user interactions and workflows initiated externally and internally in the MX Platform. Max 1024 characters. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersAPI.VerifyMember(context.Background(), memberGuid, userGuid).XCALLBACKPAYLOAD(xCALLBACKPAYLOAD).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.VerifyMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MembersAPI.VerifyMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCALLBACKPAYLOAD** | **string** | The base64 encoded string defined in this header will be returned in the [Member](/resources/webhooks/member/) and [Member Data Updated](/resources/webhooks/member#member-data-updated) webhooks. This allows you to trace user interactions and workflows initiated externally and internally in the MX Platform. Max 1024 characters. | 

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

