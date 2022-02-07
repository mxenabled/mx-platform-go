# \MxPlatformApi

All URIs are relative to *https://api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateMember**](MxPlatformApi.md#AggregateMember) | **Post** /users/{user_guid}/members/{member_guid}/aggregate | Aggregate member
[**CheckBalances**](MxPlatformApi.md#CheckBalances) | **Post** /users/{user_guid}/members/{member_guid}/check_balance | Check balances
[**CreateCategory**](MxPlatformApi.md#CreateCategory) | **Post** /users/{user_guid}/categories | Create category
[**CreateManagedAccount**](MxPlatformApi.md#CreateManagedAccount) | **Post** /users/{user_guid}/managed_members/{member_guid}/accounts | Create managed account
[**CreateManagedMember**](MxPlatformApi.md#CreateManagedMember) | **Post** /users/{user_guid}/managed_members | Create managed member
[**CreateManagedTransaction**](MxPlatformApi.md#CreateManagedTransaction) | **Post** /users/{user_guid}/managed_members/{member_guid}/transactions | Create managed transaction
[**CreateMember**](MxPlatformApi.md#CreateMember) | **Post** /users/{user_guid}/members | Create member
[**CreateTag**](MxPlatformApi.md#CreateTag) | **Post** /users/{user_guid}/tags | Create tag
[**CreateTagging**](MxPlatformApi.md#CreateTagging) | **Post** /users/{user_guid}/taggings | Create tagging
[**CreateTransactionRule**](MxPlatformApi.md#CreateTransactionRule) | **Post** /users/{user_guid}/transaction_rules | Create transaction rule
[**CreateUser**](MxPlatformApi.md#CreateUser) | **Post** /users | Create user
[**DeleteCategory**](MxPlatformApi.md#DeleteCategory) | **Delete** /users/{user_guid}/categories/{category_guid} | Delete category
[**DeleteManagedAccount**](MxPlatformApi.md#DeleteManagedAccount) | **Delete** /users/{user_guid}/managed_members/{member_guid}/accounts/{account_guid} | Delete managed account
[**DeleteManagedMember**](MxPlatformApi.md#DeleteManagedMember) | **Delete** /users/{user_guid}/managed_members/{member_guid} | Delete managed member
[**DeleteManagedTransaction**](MxPlatformApi.md#DeleteManagedTransaction) | **Delete** /users/{user_guid}/managed_members/{member_guid}/transactions/{transaction_guid} | Delete managed transaction
[**DeleteMember**](MxPlatformApi.md#DeleteMember) | **Delete** /users/{user_guid}/members/{member_guid} | Delete member
[**DeleteTag**](MxPlatformApi.md#DeleteTag) | **Delete** /users/{user_guid}/tags/{tag_guid} | Delete tag
[**DeleteTagging**](MxPlatformApi.md#DeleteTagging) | **Delete** /users/{user_guid}/taggings/{tagging_guid} | Delete tagging
[**DeleteTransactionRule**](MxPlatformApi.md#DeleteTransactionRule) | **Delete** /users/{user_guid}/transaction_rules/{transaction_rule_guid} | Delete transaction rule
[**DeleteUser**](MxPlatformApi.md#DeleteUser) | **Delete** /users/{user_guid} | Delete user
[**DownloadStatementPDF**](MxPlatformApi.md#DownloadStatementPDF) | **Get** /users/{user_guid}/members/{member_guid}/statements/{statement_guid}.pdf | Download statement pdf
[**EnhanceTransactions**](MxPlatformApi.md#EnhanceTransactions) | **Post** /transactions/enhance | Enhance transactions
[**ExtendHistory**](MxPlatformApi.md#ExtendHistory) | **Post** /users/{user_guid}/members/{member_guid}/extend_history | Extend history
[**FetchStatements**](MxPlatformApi.md#FetchStatements) | **Post** /users/{user_guid}/members/{member_guid}/fetch_statements | Fetch statements
[**IdentifyMember**](MxPlatformApi.md#IdentifyMember) | **Post** /users/{user_guid}/members/{member_guid}/identify | Identify member
[**ListAccountNumbersByAccount**](MxPlatformApi.md#ListAccountNumbersByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/account_numbers | List account numbers by account
[**ListAccountNumbersByMember**](MxPlatformApi.md#ListAccountNumbersByMember) | **Get** /users/{user_guid}/members/{member_guid}/account_numbers | List account numbers by member
[**ListAccountOwnersByMember**](MxPlatformApi.md#ListAccountOwnersByMember) | **Get** /users/{user_guid}/members/{member_guid}/account_owners | List account owners by member
[**ListCategories**](MxPlatformApi.md#ListCategories) | **Get** /users/{user_guid}/categories | List categories
[**ListDefaultCategories**](MxPlatformApi.md#ListDefaultCategories) | **Get** /categories/default | List default categories
[**ListDefaultCategoriesByUser**](MxPlatformApi.md#ListDefaultCategoriesByUser) | **Get** /users/{user_guid}/categories/default | List default categories by user
[**ListFavoriteInstitutions**](MxPlatformApi.md#ListFavoriteInstitutions) | **Get** /institutions/favorites | List favorite institutions
[**ListHoldings**](MxPlatformApi.md#ListHoldings) | **Get** /users/{user_guid}/holdings | List holdings
[**ListHoldingsByAccount**](MxPlatformApi.md#ListHoldingsByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/holdings | List holdings by account
[**ListHoldingsByMember**](MxPlatformApi.md#ListHoldingsByMember) | **Get** /users/{user_guid}/members/{member_guid}/holdings | List holdings by member
[**ListInstitutionCredentials**](MxPlatformApi.md#ListInstitutionCredentials) | **Get** /institutions/{institution_code}/credentials | List institution credentials
[**ListInstitutions**](MxPlatformApi.md#ListInstitutions) | **Get** /institutions | List institutions
[**ListManagedAccounts**](MxPlatformApi.md#ListManagedAccounts) | **Get** /users/{user_guid}/managed_members/{member_guid}/accounts | List managed accounts
[**ListManagedInstitutions**](MxPlatformApi.md#ListManagedInstitutions) | **Get** /managed_institutions | List managed institutions
[**ListManagedMembers**](MxPlatformApi.md#ListManagedMembers) | **Get** /users/{user_guid}/managed_members | List managed members
[**ListManagedTransactions**](MxPlatformApi.md#ListManagedTransactions) | **Get** /users/{user_guid}/managed_members/{member_guid}/transactions | List managed transactions
[**ListMemberChallenges**](MxPlatformApi.md#ListMemberChallenges) | **Get** /users/{user_guid}/members/{member_guid}/challenges | List member challenges
[**ListMemberCredentials**](MxPlatformApi.md#ListMemberCredentials) | **Get** /users/{user_guid}/members/{member_guid}/credentials | List member credentials
[**ListMembers**](MxPlatformApi.md#ListMembers) | **Get** /users/{user_guid}/members | List members
[**ListMerchants**](MxPlatformApi.md#ListMerchants) | **Get** /merchants | List merchants
[**ListStatementsByMember**](MxPlatformApi.md#ListStatementsByMember) | **Get** /users/{user_guid}/members/{member_guid}/statements | List statements by member
[**ListTaggings**](MxPlatformApi.md#ListTaggings) | **Get** /users/{user_guid}/taggings | List taggings
[**ListTags**](MxPlatformApi.md#ListTags) | **Get** /users/{user_guid}/tags | List tags
[**ListTransactionRules**](MxPlatformApi.md#ListTransactionRules) | **Get** /users/{user_guid}/transaction_rules | List transaction rules
[**ListTransactions**](MxPlatformApi.md#ListTransactions) | **Get** /users/{user_guid}/transactions | List transactions
[**ListTransactionsByAccount**](MxPlatformApi.md#ListTransactionsByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/transactions | List transactions by account
[**ListTransactionsByMember**](MxPlatformApi.md#ListTransactionsByMember) | **Get** /users/{user_guid}/members/{member_guid}/transactions | List transactions by member
[**ListTransactionsByTag**](MxPlatformApi.md#ListTransactionsByTag) | **Get** /users/{user_guid}/tags/{tag_guid}/transactions | List transactions by tag
[**ListUserAccounts**](MxPlatformApi.md#ListUserAccounts) | **Get** /users/{user_guid}/accounts | List accounts
[**ListUsers**](MxPlatformApi.md#ListUsers) | **Get** /users | List users
[**ReadAccount**](MxPlatformApi.md#ReadAccount) | **Get** /users/{user_guid}/accounts/{account_guid} | Read account
[**ReadCategory**](MxPlatformApi.md#ReadCategory) | **Get** /users/{user_guid}/categories/{category_guid} | Read a custom category
[**ReadDefaultCategory**](MxPlatformApi.md#ReadDefaultCategory) | **Get** /categories/{category_guid} | Read a default category
[**ReadHolding**](MxPlatformApi.md#ReadHolding) | **Get** /users/{user_guid}/holdings/{holding_guid} | Read holding
[**ReadInstitution**](MxPlatformApi.md#ReadInstitution) | **Get** /institutions/{institution_code} | Read institution
[**ReadManagedAccount**](MxPlatformApi.md#ReadManagedAccount) | **Get** /users/{user_guid}/managed_members/{member_guid}/accounts/{account_guid} | Read managed account
[**ReadManagedMember**](MxPlatformApi.md#ReadManagedMember) | **Get** /users/{user_guid}/managed_members/{member_guid} | Read managed member
[**ReadManagedTransaction**](MxPlatformApi.md#ReadManagedTransaction) | **Get** /users/{user_guid}/managed_members/{member_guid}/transactions/{transaction_guid} | Read managed transaction
[**ReadMember**](MxPlatformApi.md#ReadMember) | **Get** /users/{user_guid}/members/{member_guid} | Read member
[**ReadMemberStatus**](MxPlatformApi.md#ReadMemberStatus) | **Get** /users/{user_guid}/members/{member_guid}/status | Read member status
[**ReadMerchant**](MxPlatformApi.md#ReadMerchant) | **Get** /merchants/{merchant_guid} | Read merchant
[**ReadMerchantLocation**](MxPlatformApi.md#ReadMerchantLocation) | **Get** /merchant_locations/{merchant_location_guid} | Read merchant location
[**ReadStatementByMember**](MxPlatformApi.md#ReadStatementByMember) | **Get** /users/{user_guid}/members/{member_guid}/statements/{statement_guid} | Read statement by member
[**ReadTag**](MxPlatformApi.md#ReadTag) | **Get** /users/{user_guid}/tags/{tag_guid} | Read tag
[**ReadTagging**](MxPlatformApi.md#ReadTagging) | **Get** /users/{user_guid}/taggings/{tagging_guid} | Read tagging
[**ReadTransaction**](MxPlatformApi.md#ReadTransaction) | **Get** /users/{user_guid}/transactions/{transaction_guid} | Read transaction
[**ReadTransactionRule**](MxPlatformApi.md#ReadTransactionRule) | **Get** /users/{user_guid}/transaction_rules/{transaction_rule_guid} | Read transaction rule
[**ReadUser**](MxPlatformApi.md#ReadUser) | **Get** /users/{user_guid} | Read user
[**RequestConnectWidgetURL**](MxPlatformApi.md#RequestConnectWidgetURL) | **Post** /users/{user_guid}/connect_widget_url | Request connect widget url
[**RequestOAuthWindowURI**](MxPlatformApi.md#RequestOAuthWindowURI) | **Get** /users/{user_guid}/members/{member_guid}/oauth_window_uri | Request oauth window uri
[**RequestWidgetURL**](MxPlatformApi.md#RequestWidgetURL) | **Post** /users/{user_guid}/widget_urls | Request widget url
[**ResumeAggregation**](MxPlatformApi.md#ResumeAggregation) | **Put** /users/{user_guid}/members/{member_guid}/resume | Resume aggregation
[**UpdateAccountByMember**](MxPlatformApi.md#UpdateAccountByMember) | **Put** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Update account by member
[**UpdateCategory**](MxPlatformApi.md#UpdateCategory) | **Put** /users/{user_guid}/categories/{category_guid} | Update category
[**UpdateManagedAccount**](MxPlatformApi.md#UpdateManagedAccount) | **Put** /users/{user_guid}/managed_members/{member_guid}/accounts/{account_guid} | Update managed account
[**UpdateManagedMember**](MxPlatformApi.md#UpdateManagedMember) | **Put** /users/{user_guid}/managed_members/{member_guid} | Update managed member
[**UpdateManagedTransaction**](MxPlatformApi.md#UpdateManagedTransaction) | **Put** /users/{user_guid}/managed_members/{member_guid}/transactions/{transaction_guid} | Update managed transaction
[**UpdateMember**](MxPlatformApi.md#UpdateMember) | **Put** /users/{user_guid}/members/{member_guid} | Update member
[**UpdateTag**](MxPlatformApi.md#UpdateTag) | **Put** /users/{user_guid}/tags/{tag_guid} | Update tag
[**UpdateTagging**](MxPlatformApi.md#UpdateTagging) | **Put** /users/{user_guid}/taggings/{tagging_guid} | Update tagging
[**UpdateTransaction**](MxPlatformApi.md#UpdateTransaction) | **Put** /users/{user_guid}/transactions/{transaction_guid} | Update transaction
[**UpdateTransactionRule**](MxPlatformApi.md#UpdateTransactionRule) | **Put** /users/{user_guid}/transaction_rules/{transaction_rule_guid} | Update transaction_rule
[**UpdateUser**](MxPlatformApi.md#UpdateUser) | **Put** /users/{user_guid} | Update user
[**VerifyMember**](MxPlatformApi.md#VerifyMember) | **Post** /users/{user_guid}/members/{member_guid}/verify | Verify member



## AggregateMember

> MemberResponseBody AggregateMember(ctx, memberGuid, userGuid).Execute()

Aggregate member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.AggregateMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.AggregateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AggregateMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.AggregateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAggregateMemberRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.CheckBalances(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.CheckBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckBalances`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.CheckBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    categoryCreateRequestBody := *openapiclient.NewCategoryCreateRequestBody() // CategoryCreateRequestBody | Custom category object to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.CreateCategory(context.Background(), userGuid).CategoryCreateRequestBody(categoryCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.CreateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.CreateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## CreateManagedAccount

> AccountResponseBody CreateManagedAccount(ctx, userGuid, memberGuid).ManagedAccountCreateRequestBody(managedAccountCreateRequestBody).Execute()

Create managed account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    managedAccountCreateRequestBody := *openapiclient.NewManagedAccountCreateRequestBody() // ManagedAccountCreateRequestBody | Managed account to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.CreateManagedAccount(context.Background(), userGuid, memberGuid).ManagedAccountCreateRequestBody(managedAccountCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.CreateManagedAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagedAccount`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.CreateManagedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **managedAccountCreateRequestBody** | [**ManagedAccountCreateRequestBody**](ManagedAccountCreateRequestBody.md) | Managed account to be created. | 

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


## CreateManagedMember

> MemberResponseBody CreateManagedMember(ctx, userGuid).ManagedMemberCreateRequestBody(managedMemberCreateRequestBody).Execute()

Create managed member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    managedMemberCreateRequestBody := *openapiclient.NewManagedMemberCreateRequestBody() // ManagedMemberCreateRequestBody | Managed member to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.CreateManagedMember(context.Background(), userGuid).ManagedMemberCreateRequestBody(managedMemberCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.CreateManagedMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagedMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.CreateManagedMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedMemberCreateRequestBody** | [**ManagedMemberCreateRequestBody**](ManagedMemberCreateRequestBody.md) | Managed member to be created. | 

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


## CreateManagedTransaction

> TransactionResponseBody CreateManagedTransaction(ctx, userGuid, memberGuid).ManagedTransactionCreateRequestBody(managedTransactionCreateRequestBody).Execute()

Create managed transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    managedTransactionCreateRequestBody := *openapiclient.NewManagedTransactionCreateRequestBody() // ManagedTransactionCreateRequestBody | Managed transaction to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.CreateManagedTransaction(context.Background(), userGuid, memberGuid).ManagedTransactionCreateRequestBody(managedTransactionCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.CreateManagedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagedTransaction`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.CreateManagedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **managedTransactionCreateRequestBody** | [**ManagedTransactionCreateRequestBody**](ManagedTransactionCreateRequestBody.md) | Managed transaction to be created. | 

### Return type

[**TransactionResponseBody**](TransactionResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMember

> MemberResponseBody CreateMember(ctx, userGuid).MemberCreateRequestBody(memberCreateRequestBody).Execute()

Create member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberCreateRequestBody := *openapiclient.NewMemberCreateRequestBody() // MemberCreateRequestBody | Member object to be created with optional parameters (id and metadata) and required parameters (credentials and institution_code)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.CreateMember(context.Background(), userGuid).MemberCreateRequestBody(memberCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.CreateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.CreateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memberCreateRequestBody** | [**MemberCreateRequestBody**](MemberCreateRequestBody.md) | Member object to be created with optional parameters (id and metadata) and required parameters (credentials and institution_code) | 

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


## CreateTag

> TagResponseBody CreateTag(ctx, userGuid).TagCreateRequestBody(tagCreateRequestBody).Execute()

Create tag



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    tagCreateRequestBody := *openapiclient.NewTagCreateRequestBody() // TagCreateRequestBody | Tag object to be created with required parameters (tag_guid)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.CreateTag(context.Background(), userGuid).TagCreateRequestBody(tagCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.CreateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTag`: TagResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.CreateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagCreateRequestBody** | [**TagCreateRequestBody**](TagCreateRequestBody.md) | Tag object to be created with required parameters (tag_guid) | 

### Return type

[**TagResponseBody**](TagResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    taggingCreateRequestBody := *openapiclient.NewTaggingCreateRequestBody() // TaggingCreateRequestBody | Tagging object to be created with required parameters (tag_guid and transaction_guid)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.CreateTagging(context.Background(), userGuid).TaggingCreateRequestBody(taggingCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.CreateTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagging`: TaggingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.CreateTagging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## CreateTransactionRule

> TransactionRuleResponseBody CreateTransactionRule(ctx, userGuid).TransactionRuleCreateRequestBody(transactionRuleCreateRequestBody).Execute()

Create transaction rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    transactionRuleCreateRequestBody := *openapiclient.NewTransactionRuleCreateRequestBody() // TransactionRuleCreateRequestBody | TransactionRule object to be created with optional parameters (description) and required parameters (category_guid and match_description)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.CreateTransactionRule(context.Background(), userGuid).TransactionRuleCreateRequestBody(transactionRuleCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.CreateTransactionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransactionRule`: TransactionRuleResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.CreateTransactionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionRuleCreateRequestBody** | [**TransactionRuleCreateRequestBody**](TransactionRuleCreateRequestBody.md) | TransactionRule object to be created with optional parameters (description) and required parameters (category_guid and match_description) | 

### Return type

[**TransactionRuleResponseBody**](TransactionRuleResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "./openapi"
)

func main() {
    userCreateRequestBody := *openapiclient.NewUserCreateRequestBody() // UserCreateRequestBody | User object to be created. (None of these parameters are required, but the user object cannot be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.CreateUser(context.Background()).UserCreateRequestBody(userCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: UserResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.CreateUser`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    categoryGuid := "CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874" // string | The unique id for a `category`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.DeleteCategory(context.Background(), categoryGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.DeleteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryGuid** | **string** | The unique id for a &#x60;category&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## DeleteManagedAccount

> DeleteManagedAccount(ctx, memberGuid, userGuid, accountGuid).Execute()

Delete managed account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.DeleteManagedAccount(context.Background(), memberGuid, userGuid, accountGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.DeleteManagedAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedAccountRequest struct via the builder pattern


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


## DeleteManagedMember

> DeleteManagedMember(ctx, memberGuid, userGuid).Execute()

Delete managed member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.DeleteManagedMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.DeleteManagedMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedMemberRequest struct via the builder pattern


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


## DeleteManagedTransaction

> DeleteManagedTransaction(ctx, memberGuid, userGuid, transactionGuid).Execute()

Delete managed transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.DeleteManagedTransaction(context.Background(), memberGuid, userGuid, transactionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.DeleteManagedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedTransactionRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.DeleteMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.DeleteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## DeleteTag

> DeleteTag(ctx, tagGuid, userGuid).Execute()

Delete tag



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tagGuid := "TAG-aef36e72-6294-4c38-844d-e573e80aed52" // string | The unique id for a `tag`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.DeleteTag(context.Background(), tagGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagGuid** | **string** | The unique id for a &#x60;tag&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    taggingGuid := "TGN-007f5486-17e1-45fc-8b87-8f03984430fe" // string | The unique id for a `tagging`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.DeleteTagging(context.Background(), taggingGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.DeleteTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taggingGuid** | **string** | The unique id for a &#x60;tagging&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## DeleteTransactionRule

> DeleteTransactionRule(ctx, transactionRuleGuid, userGuid).Execute()

Delete transaction rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionRuleGuid := "TXR-a080e0f9-a2d4-4d6f-9e03-672cc357a4d3" // string | The unique id for a `transaction_rule`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.DeleteTransactionRule(context.Background(), transactionRuleGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.DeleteTransactionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionRuleGuid** | **string** | The unique id for a &#x60;transaction_rule&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransactionRuleRequest struct via the builder pattern


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


## DeleteUser

> DeleteUser(ctx, userGuid).Execute()

Delete user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.DeleteUser(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    statementGuid := "STA-737a344b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for a `statement`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.DownloadStatementPDF(context.Background(), memberGuid, statementGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.DownloadStatementPDF``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadStatementPDF`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.DownloadStatementPDF`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**statementGuid** | **string** | The unique id for a &#x60;statement&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## EnhanceTransactions

> EnhanceTransactionsResponseBody EnhanceTransactions(ctx).EnhanceTransactionsRequestBody(enhanceTransactionsRequestBody).Execute()

Enhance transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enhanceTransactionsRequestBody := *openapiclient.NewEnhanceTransactionsRequestBody() // EnhanceTransactionsRequestBody | Transaction object to be enhanced

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.EnhanceTransactions(context.Background()).EnhanceTransactionsRequestBody(enhanceTransactionsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.EnhanceTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnhanceTransactions`: EnhanceTransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.EnhanceTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnhanceTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enhanceTransactionsRequestBody** | [**EnhanceTransactionsRequestBody**](EnhanceTransactionsRequestBody.md) | Transaction object to be enhanced | 

### Return type

[**EnhanceTransactionsResponseBody**](EnhanceTransactionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendHistory

> MemberResponseBody ExtendHistory(ctx, memberGuid, userGuid).Execute()

Extend history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique identifier for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ExtendHistory(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ExtendHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtendHistory`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ExtendHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique identifier for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendHistoryRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.FetchStatements(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.FetchStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchStatements`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.FetchStatements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.IdentifyMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.IdentifyMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentifyMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.IdentifyMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "./openapi"
)

func main() {
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListAccountNumbersByAccount(context.Background(), accountGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListAccountNumbersByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountNumbersByAccount`: AccountNumbersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListAccountNumbersByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountNumbersByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListAccountNumbersByMember(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListAccountNumbersByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountNumbersByMember`: AccountNumbersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListAccountNumbersByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountNumbersByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListAccountOwnersByMember(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListAccountOwnersByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountOwnersByMember`: AccountOwnersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListAccountOwnersByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountOwnersByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListCategories(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCategories`: CategoriesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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
    openapiclient "./openapi"
)

func main() {
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListDefaultCategories(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListDefaultCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDefaultCategories`: CategoriesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListDefaultCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDefaultCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListDefaultCategoriesByUser(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListDefaultCategoriesByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDefaultCategoriesByUser`: CategoriesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListDefaultCategoriesByUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDefaultCategoriesByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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


## ListFavoriteInstitutions

> InstitutionsResponseBody ListFavoriteInstitutions(ctx).Page(page).RecordsPerPage(recordsPerPage).Execute()

List favorite institutions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListFavoriteInstitutions(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListFavoriteInstitutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFavoriteInstitutions`: InstitutionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListFavoriteInstitutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFavoriteInstitutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**InstitutionsResponseBody**](InstitutionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHoldings

> HoldingsResponseBody ListHoldings(ctx, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()

List holdings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    fromDate := "2015-09-20" // string | Filter holdings from this date. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)
    toDate := "2019-10-20" // string | Filter holdings to this date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListHoldings(context.Background(), userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListHoldings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHoldings`: HoldingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListHoldings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | Filter holdings from this date. | 
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 
 **toDate** | **string** | Filter holdings to this date. | 

### Return type

[**HoldingsResponseBody**](HoldingsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHoldingsByAccount

> HoldingsResponseBody ListHoldingsByAccount(ctx, accountGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()

List holdings by account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountGuid := "ACT-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for the `account`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for the `user`.
    fromDate := "2015-09-20" // string | Filter holdings from this date. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)
    toDate := "2019-10-20" // string | Filter holdings to this date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListHoldingsByAccount(context.Background(), accountGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListHoldingsByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHoldingsByAccount`: HoldingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListHoldingsByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for the &#x60;account&#x60;. | 
**userGuid** | **string** | The unique id for the &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHoldingsByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDate** | **string** | Filter holdings from this date. | 
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 
 **toDate** | **string** | Filter holdings to this date. | 

### Return type

[**HoldingsResponseBody**](HoldingsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHoldingsByMember

> HoldingsResponseBody ListHoldingsByMember(ctx, memberGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()

List holdings by member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    fromDate := "2015-09-20" // string | Filter holdings from this date. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)
    toDate := "2019-10-20" // string | Filter holdings to this date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListHoldingsByMember(context.Background(), memberGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListHoldingsByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHoldingsByMember`: HoldingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListHoldingsByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHoldingsByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDate** | **string** | Filter holdings from this date. | 
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 
 **toDate** | **string** | Filter holdings to this date. | 

### Return type

[**HoldingsResponseBody**](HoldingsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstitutionCredentials

> CredentialsResponseBody ListInstitutionCredentials(ctx, institutionCode).Page(page).RecordsPerPage(recordsPerPage).Execute()

List institution credentials



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    institutionCode := "chase" // string | The institution_code of the institution.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListInstitutionCredentials(context.Background(), institutionCode).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListInstitutionCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstitutionCredentials`: CredentialsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListInstitutionCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**institutionCode** | **string** | The institution_code of the institution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstitutionCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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


## ListInstitutions

> InstitutionsResponseBody ListInstitutions(ctx).Name(name).Page(page).RecordsPerPage(recordsPerPage).SupportsAccountIdentification(supportsAccountIdentification).SupportsAccountStatement(supportsAccountStatement).SupportsAccountVerification(supportsAccountVerification).SupportsTransactionHistory(supportsTransactionHistory).Execute()

List institutions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "chase" // string | This will list only institutions in which the appended string appears. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)
    supportsAccountIdentification := true // bool | Filter only institutions which support account identification. (optional)
    supportsAccountStatement := true // bool | Filter only institutions which support account statements. (optional)
    supportsAccountVerification := true // bool | Filter only institutions which support account verification. (optional)
    supportsTransactionHistory := true // bool | Filter only institutions which support extended transaction history. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListInstitutions(context.Background()).Name(name).Page(page).RecordsPerPage(recordsPerPage).SupportsAccountIdentification(supportsAccountIdentification).SupportsAccountStatement(supportsAccountStatement).SupportsAccountVerification(supportsAccountVerification).SupportsTransactionHistory(supportsTransactionHistory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListInstitutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstitutions`: InstitutionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListInstitutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstitutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | This will list only institutions in which the appended string appears. | 
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 
 **supportsAccountIdentification** | **bool** | Filter only institutions which support account identification. | 
 **supportsAccountStatement** | **bool** | Filter only institutions which support account statements. | 
 **supportsAccountVerification** | **bool** | Filter only institutions which support account verification. | 
 **supportsTransactionHistory** | **bool** | Filter only institutions which support extended transaction history. | 

### Return type

[**InstitutionsResponseBody**](InstitutionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListManagedAccounts

> AccountsResponseBody ListManagedAccounts(ctx, userGuid, memberGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List managed accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListManagedAccounts(context.Background(), userGuid, memberGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListManagedAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManagedAccounts`: AccountsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListManagedAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListManagedAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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


## ListManagedInstitutions

> InstitutionsResponseBody ListManagedInstitutions(ctx).Page(page).RecordsPerPage(recordsPerPage).Execute()

List managed institutions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListManagedInstitutions(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListManagedInstitutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManagedInstitutions`: InstitutionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListManagedInstitutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListManagedInstitutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**InstitutionsResponseBody**](InstitutionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListManagedMembers

> MembersResponseBody ListManagedMembers(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List managed members



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListManagedMembers(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListManagedMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManagedMembers`: MembersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListManagedMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListManagedMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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


## ListManagedTransactions

> TransactionsResponseBody ListManagedTransactions(ctx, userGuid, memberGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List managed transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListManagedTransactions(context.Background(), userGuid, memberGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListManagedTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManagedTransactions`: TransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListManagedTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListManagedTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**TransactionsResponseBody**](TransactionsResponseBody.md)

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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListMemberChallenges(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListMemberChallenges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMemberChallenges`: ChallengesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListMemberChallenges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMemberChallengesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListMemberCredentials(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListMemberCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMemberCredentials`: CredentialsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListMemberCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMemberCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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

> MembersResponseBody ListMembers(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List members



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListMembers(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMembers`: MembersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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


## ListMerchants

> MerchantsResponseBody ListMerchants(ctx).Page(page).RecordsPerPage(recordsPerPage).Execute()

List merchants



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListMerchants(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListMerchants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMerchants`: MerchantsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListMerchants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMerchantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**MerchantsResponseBody**](MerchantsResponseBody.md)

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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListStatementsByMember(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListStatementsByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStatementsByMember`: StatementsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListStatementsByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStatementsByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListTaggings(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListTaggings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaggings`: TaggingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListTaggings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTaggingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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


## ListTags

> TagsResponseBody ListTags(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List tags



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListTags(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTags`: TagsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**TagsResponseBody**](TagsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionRules

> TransactionRulesResponseBody ListTransactionRules(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List transaction rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListTransactionRules(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListTransactionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionRules`: TransactionRulesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListTransactionRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**TransactionRulesResponseBody**](TransactionRulesResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> TransactionsResponseBody ListTransactions(ctx, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()

List transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    fromDate := "2015-09-20" // string | Filter transactions from this date. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)
    toDate := "2019-10-20" // string | Filter transactions to this date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListTransactions(context.Background(), userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactions`: TransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | Filter transactions from this date. | 
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 
 **toDate** | **string** | Filter transactions to this date. | 

### Return type

[**TransactionsResponseBody**](TransactionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsByAccount

> TransactionsResponseBody ListTransactionsByAccount(ctx, accountGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()

List transactions by account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    fromDate := "2015-09-20" // string | Filter transactions from this date. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)
    toDate := "2019-10-20" // string | Filter transactions to this date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListTransactionsByAccount(context.Background(), accountGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListTransactionsByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByAccount`: TransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListTransactionsByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDate** | **string** | Filter transactions from this date. | 
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 
 **toDate** | **string** | Filter transactions to this date. | 

### Return type

[**TransactionsResponseBody**](TransactionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsByMember

> TransactionsResponseBody ListTransactionsByMember(ctx, memberGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()

List transactions by member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    fromDate := "2015-09-20" // string | Filter transactions from this date. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)
    toDate := "2019-10-20" // string | Filter transactions to this date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListTransactionsByMember(context.Background(), memberGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListTransactionsByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByMember`: TransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListTransactionsByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDate** | **string** | Filter transactions from this date. | 
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 
 **toDate** | **string** | Filter transactions to this date. | 

### Return type

[**TransactionsResponseBody**](TransactionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionsByTag

> TransactionsResponseBody ListTransactionsByTag(ctx, tagGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()

List transactions by tag



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tagGuid := "TAG-aef36e72-6294-4c38-844d-e573e80aed52" // string | The unique id for a `tag`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    fromDate := "2015-09-20" // string | Filter transactions from this date. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)
    toDate := "2019-10-20" // string | Filter transactions to this date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListTransactionsByTag(context.Background(), tagGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListTransactionsByTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByTag`: TransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListTransactionsByTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagGuid** | **string** | The unique id for a &#x60;tag&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsByTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDate** | **string** | Filter transactions from this date. | 
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 
 **toDate** | **string** | Filter transactions to this date. | 

### Return type

[**TransactionsResponseBody**](TransactionsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserAccounts

> AccountsResponseBody ListUserAccounts(ctx, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListUserAccounts(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListUserAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserAccounts`: AccountsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListUserAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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


## ListUsers

> UsersResponseBody ListUsers(ctx).Page(page).RecordsPerPage(recordsPerPage).Execute()

List users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ListUsers(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: UsersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

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
    openapiclient "./openapi"
)

func main() {
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadAccount(context.Background(), accountGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAccount`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "./openapi"
)

func main() {
    categoryGuid := "CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874" // string | The unique id for a `category`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadCategory(context.Background(), categoryGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryGuid** | **string** | The unique id for a &#x60;category&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "./openapi"
)

func main() {
    categoryGuid := "CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874" // string | The unique id for a `category`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadDefaultCategory(context.Background(), categoryGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadDefaultCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDefaultCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadDefaultCategory`: %v\n", resp)
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


## ReadHolding

> HoldingResponseBody ReadHolding(ctx, holdingGuid, userGuid).Execute()

Read holding



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    holdingGuid := "HOL-d65683e8-9eab-26bb-bcfd-ced159c9abe2" // string | The unique id for a `holding`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadHolding(context.Background(), holdingGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadHolding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadHolding`: HoldingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadHolding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holdingGuid** | **string** | The unique id for a &#x60;holding&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadHoldingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HoldingResponseBody**](HoldingResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInstitution

> InstitutionResponseBody ReadInstitution(ctx, institutionCode).Execute()

Read institution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    institutionCode := "chase" // string | The institution_code of the institution.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadInstitution(context.Background(), institutionCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadInstitution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadInstitution`: InstitutionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadInstitution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**institutionCode** | **string** | The institution_code of the institution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadInstitutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstitutionResponseBody**](InstitutionResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadManagedAccount

> AccountResponseBody ReadManagedAccount(ctx, memberGuid, userGuid, accountGuid).Execute()

Read managed account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadManagedAccount(context.Background(), memberGuid, userGuid, accountGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadManagedAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadManagedAccount`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadManagedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadManagedAccountRequest struct via the builder pattern


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


## ReadManagedMember

> MemberResponseBody ReadManagedMember(ctx, memberGuid, userGuid).Execute()

Read managed member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadManagedMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadManagedMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadManagedMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadManagedMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadManagedMemberRequest struct via the builder pattern


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


## ReadManagedTransaction

> TransactionResponseBody ReadManagedTransaction(ctx, memberGuid, userGuid, transactionGuid).Execute()

Read managed transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadManagedTransaction(context.Background(), memberGuid, userGuid, transactionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadManagedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadManagedTransaction`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadManagedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadManagedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TransactionResponseBody**](TransactionResponseBody.md)

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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadMemberStatus(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadMemberStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMemberStatus`: MemberStatusResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadMemberStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## ReadMerchant

> MerchantResponseBody ReadMerchant(ctx, merchantGuid).Execute()

Read merchant



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    merchantGuid := "MCH-7ed79542-884d-2b1b-dd74-501c5cc9d25b" // string | The unique id for a `merchant`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadMerchant(context.Background(), merchantGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadMerchant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMerchant`: MerchantResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadMerchant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantGuid** | **string** | The unique id for a &#x60;merchant&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadMerchantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MerchantResponseBody**](MerchantResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMerchantLocation

> MerchantLocationResponseBody ReadMerchantLocation(ctx, merchantLocationGuid).Execute()

Read merchant location



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    merchantLocationGuid := "MCH-09466f0a-fb58-9d1a-bae2-2af0afbea621" // string | The unique id for a `merchant_location`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadMerchantLocation(context.Background(), merchantLocationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadMerchantLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMerchantLocation`: MerchantLocationResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadMerchantLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantLocationGuid** | **string** | The unique id for a &#x60;merchant_location&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadMerchantLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MerchantLocationResponseBody**](MerchantLocationResponseBody.md)

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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    statementGuid := "STA-737a344b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for a `statement`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadStatementByMember(context.Background(), memberGuid, statementGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadStatementByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadStatementByMember`: StatementResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadStatementByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**statementGuid** | **string** | The unique id for a &#x60;statement&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## ReadTag

> TagResponseBody ReadTag(ctx, tagGuid, userGuid).Execute()

Read tag



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tagGuid := "TAG-aef36e72-6294-4c38-844d-e573e80aed52" // string | The unique id for a `tag`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadTag(context.Background(), tagGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTag`: TagResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagGuid** | **string** | The unique id for a &#x60;tag&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TagResponseBody**](TagResponseBody.md)

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
    openapiclient "./openapi"
)

func main() {
    taggingGuid := "TGN-007f5486-17e1-45fc-8b87-8f03984430fe" // string | The unique id for a `tagging`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadTagging(context.Background(), taggingGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTagging`: TaggingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadTagging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taggingGuid** | **string** | The unique id for a &#x60;tagging&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## ReadTransaction

> TransactionResponseBody ReadTransaction(ctx, transactionGuid, userGuid).Execute()

Read transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadTransaction(context.Background(), transactionGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTransaction`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TransactionResponseBody**](TransactionResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTransactionRule

> TransactionRuleResponseBody ReadTransactionRule(ctx, transactionRuleGuid, userGuid).Execute()

Read transaction rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionRuleGuid := "TXR-a080e0f9-a2d4-4d6f-9e03-672cc357a4d3" // string | The unique id for a `transaction_rule`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadTransactionRule(context.Background(), transactionRuleGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadTransactionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTransactionRule`: TransactionRuleResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadTransactionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionRuleGuid** | **string** | The unique id for a &#x60;transaction_rule&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTransactionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TransactionRuleResponseBody**](TransactionRuleResponseBody.md)

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
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ReadUser(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ReadUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUser`: UserResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ReadUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## RequestConnectWidgetURL

> ConnectWidgetResponseBody RequestConnectWidgetURL(ctx, userGuid).ConnectWidgetRequestBody(connectWidgetRequestBody).Execute()

Request connect widget url



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    connectWidgetRequestBody := *openapiclient.NewConnectWidgetRequestBody() // ConnectWidgetRequestBody | Optional config options for WebView (is_mobile_webview, current_institution_code, current_member_guid, update_credentials)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.RequestConnectWidgetURL(context.Background(), userGuid).ConnectWidgetRequestBody(connectWidgetRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.RequestConnectWidgetURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestConnectWidgetURL`: ConnectWidgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.RequestConnectWidgetURL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestConnectWidgetURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectWidgetRequestBody** | [**ConnectWidgetRequestBody**](ConnectWidgetRequestBody.md) | Optional config options for WebView (is_mobile_webview, current_institution_code, current_member_guid, update_credentials) | 

### Return type

[**ConnectWidgetResponseBody**](ConnectWidgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestOAuthWindowURI

> OAuthWindowResponseBody RequestOAuthWindowURI(ctx, memberGuid, userGuid).ReferralSource(referralSource).UiMessageWebviewUrlScheme(uiMessageWebviewUrlScheme).SkipAggregation(skipAggregation).Execute()

Request oauth window uri



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    referralSource := "APP" // string | Must be either `BROWSER` or `APP` depending on the implementation. Defaults to `BROWSER`. (optional)
    uiMessageWebviewUrlScheme := "mx" // string | A scheme for routing the user back to the application state they were previously in. (optional)
    skipAggregation := false // bool | Setting this parameter to `true` will prevent the member from automatically aggregating after being redirected from the authorization page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.RequestOAuthWindowURI(context.Background(), memberGuid, userGuid).ReferralSource(referralSource).UiMessageWebviewUrlScheme(uiMessageWebviewUrlScheme).SkipAggregation(skipAggregation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.RequestOAuthWindowURI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestOAuthWindowURI`: OAuthWindowResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.RequestOAuthWindowURI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestOAuthWindowURIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **referralSource** | **string** | Must be either &#x60;BROWSER&#x60; or &#x60;APP&#x60; depending on the implementation. Defaults to &#x60;BROWSER&#x60;. | 
 **uiMessageWebviewUrlScheme** | **string** | A scheme for routing the user back to the application state they were previously in. | 
 **skipAggregation** | **bool** | Setting this parameter to &#x60;true&#x60; will prevent the member from automatically aggregating after being redirected from the authorization page. | 

### Return type

[**OAuthWindowResponseBody**](OAuthWindowResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestWidgetURL

> WidgetResponseBody RequestWidgetURL(ctx, userGuid).WidgetRequestBody(widgetRequestBody).AcceptLanguage(acceptLanguage).Execute()

Request widget url



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    widgetRequestBody := *openapiclient.NewWidgetRequestBody() // WidgetRequestBody | The widget url configuration options.
    acceptLanguage := "en-US" // string | The desired language of the widget. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.RequestWidgetURL(context.Background(), userGuid).WidgetRequestBody(widgetRequestBody).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.RequestWidgetURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestWidgetURL`: WidgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.RequestWidgetURL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestWidgetURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **widgetRequestBody** | [**WidgetRequestBody**](WidgetRequestBody.md) | The widget url configuration options. | 
 **acceptLanguage** | **string** | The desired language of the widget. | 

### Return type

[**WidgetResponseBody**](WidgetResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberResumeRequestBody := *openapiclient.NewMemberResumeRequestBody() // MemberResumeRequestBody | Member object with MFA challenge answers

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.ResumeAggregation(context.Background(), memberGuid, userGuid).MemberResumeRequestBody(memberResumeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.ResumeAggregation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeAggregation`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.ResumeAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## UpdateAccountByMember

> AccountResponseBody UpdateAccountByMember(ctx, userGuid, memberGuid, accountGuid).AccountUpdateRequestBody(accountUpdateRequestBody).Execute()

Update account by member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    accountUpdateRequestBody := *openapiclient.NewAccountUpdateRequestBody() // AccountUpdateRequestBody | Account object to be created with optional parameters (is_hidden)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.UpdateAccountByMember(context.Background(), userGuid, memberGuid, accountGuid).AccountUpdateRequestBody(accountUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.UpdateAccountByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountByMember`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.UpdateAccountByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountByMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accountUpdateRequestBody** | [**AccountUpdateRequestBody**](AccountUpdateRequestBody.md) | Account object to be created with optional parameters (is_hidden) | 

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
    openapiclient "./openapi"
)

func main() {
    categoryGuid := "CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874" // string | The unique id for a `category`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    categoryUpdateRequestBody := *openapiclient.NewCategoryUpdateRequestBody() // CategoryUpdateRequestBody | Category object to be updated (While no single parameter is required, the `category` object cannot be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.UpdateCategory(context.Background(), categoryGuid, userGuid).CategoryUpdateRequestBody(categoryUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.UpdateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.UpdateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryGuid** | **string** | The unique id for a &#x60;category&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## UpdateManagedAccount

> AccountResponseBody UpdateManagedAccount(ctx, memberGuid, userGuid, accountGuid).ManagedAccountUpdateRequestBody(managedAccountUpdateRequestBody).Execute()

Update managed account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    managedAccountUpdateRequestBody := *openapiclient.NewManagedAccountUpdateRequestBody() // ManagedAccountUpdateRequestBody | Managed account object to be updated (While no single parameter is required, the request body can't be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.UpdateManagedAccount(context.Background(), memberGuid, userGuid, accountGuid).ManagedAccountUpdateRequestBody(managedAccountUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.UpdateManagedAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedAccount`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.UpdateManagedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **managedAccountUpdateRequestBody** | [**ManagedAccountUpdateRequestBody**](ManagedAccountUpdateRequestBody.md) | Managed account object to be updated (While no single parameter is required, the request body can&#39;t be empty) | 

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


## UpdateManagedMember

> MemberResponseBody UpdateManagedMember(ctx, memberGuid, userGuid).ManagedMemberUpdateRequestBody(managedMemberUpdateRequestBody).Execute()

Update managed member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    managedMemberUpdateRequestBody := *openapiclient.NewManagedMemberUpdateRequestBody() // ManagedMemberUpdateRequestBody | Managed member object to be updated (While no single parameter is required, the request body can't be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.UpdateManagedMember(context.Background(), memberGuid, userGuid).ManagedMemberUpdateRequestBody(managedMemberUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.UpdateManagedMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.UpdateManagedMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **managedMemberUpdateRequestBody** | [**ManagedMemberUpdateRequestBody**](ManagedMemberUpdateRequestBody.md) | Managed member object to be updated (While no single parameter is required, the request body can&#39;t be empty) | 

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


## UpdateManagedTransaction

> TransactionResponseBody UpdateManagedTransaction(ctx, memberGuid, userGuid, transactionGuid).ManagedTransactionUpdateRequestBody(managedTransactionUpdateRequestBody).Execute()

Update managed transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    managedTransactionUpdateRequestBody := *openapiclient.NewManagedTransactionUpdateRequestBody() // ManagedTransactionUpdateRequestBody | Managed transaction object to be updated (While no single parameter is required, the request body can't be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.UpdateManagedTransaction(context.Background(), memberGuid, userGuid, transactionGuid).ManagedTransactionUpdateRequestBody(managedTransactionUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.UpdateManagedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedTransaction`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.UpdateManagedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **managedTransactionUpdateRequestBody** | [**ManagedTransactionUpdateRequestBody**](ManagedTransactionUpdateRequestBody.md) | Managed transaction object to be updated (While no single parameter is required, the request body can&#39;t be empty) | 

### Return type

[**TransactionResponseBody**](TransactionResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> MemberResponseBody UpdateMember(ctx, memberGuid, userGuid).MemberUpdateRequestBody(memberUpdateRequestBody).Execute()

Update member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberUpdateRequestBody := *openapiclient.NewMemberUpdateRequestBody() // MemberUpdateRequestBody | Member object to be updated (While no single parameter is required, the request body can't be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.UpdateMember(context.Background(), memberGuid, userGuid).MemberUpdateRequestBody(memberUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.UpdateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.UpdateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **memberUpdateRequestBody** | [**MemberUpdateRequestBody**](MemberUpdateRequestBody.md) | Member object to be updated (While no single parameter is required, the request body can&#39;t be empty) | 

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


## UpdateTag

> TagResponseBody UpdateTag(ctx, tagGuid, userGuid).TagUpdateRequestBody(tagUpdateRequestBody).Execute()

Update tag



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tagGuid := "TAG-aef36e72-6294-4c38-844d-e573e80aed52" // string | The unique id for a `tag`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    tagUpdateRequestBody := *openapiclient.NewTagUpdateRequestBody() // TagUpdateRequestBody | Tag object to be updated with required parameter (tag_guid)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.UpdateTag(context.Background(), tagGuid, userGuid).TagUpdateRequestBody(tagUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.UpdateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTag`: TagResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagGuid** | **string** | The unique id for a &#x60;tag&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagUpdateRequestBody** | [**TagUpdateRequestBody**](TagUpdateRequestBody.md) | Tag object to be updated with required parameter (tag_guid) | 

### Return type

[**TagResponseBody**](TagResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "./openapi"
)

func main() {
    taggingGuid := "TGN-007f5486-17e1-45fc-8b87-8f03984430fe" // string | The unique id for a `tagging`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    taggingUpdateRequestBody := *openapiclient.NewTaggingUpdateRequestBody() // TaggingUpdateRequestBody | Tagging object to be updated with required parameter (tag_guid)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.UpdateTagging(context.Background(), taggingGuid, userGuid).TaggingUpdateRequestBody(taggingUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.UpdateTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagging`: TaggingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.UpdateTagging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taggingGuid** | **string** | The unique id for a &#x60;tagging&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## UpdateTransaction

> TransactionResponseBody UpdateTransaction(ctx, transactionGuid, userGuid).TransactionUpdateRequestBody(transactionUpdateRequestBody).Execute()

Update transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    transactionUpdateRequestBody := *openapiclient.NewTransactionUpdateRequestBody() // TransactionUpdateRequestBody | Transaction object to be updated with a new description

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.UpdateTransaction(context.Background(), transactionGuid, userGuid).TransactionUpdateRequestBody(transactionUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.UpdateTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransaction`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.UpdateTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionUpdateRequestBody** | [**TransactionUpdateRequestBody**](TransactionUpdateRequestBody.md) | Transaction object to be updated with a new description | 

### Return type

[**TransactionResponseBody**](TransactionResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransactionRule

> TransactionRuleResponseBody UpdateTransactionRule(ctx, transactionRuleGuid, userGuid).TransactionRuleUpdateRequestBody(transactionRuleUpdateRequestBody).Execute()

Update transaction_rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionRuleGuid := "TXR-a080e0f9-a2d4-4d6f-9e03-672cc357a4d3" // string | The unique id for a `transaction_rule`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    transactionRuleUpdateRequestBody := *openapiclient.NewTransactionRuleUpdateRequestBody() // TransactionRuleUpdateRequestBody | TransactionRule object to be updated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.UpdateTransactionRule(context.Background(), transactionRuleGuid, userGuid).TransactionRuleUpdateRequestBody(transactionRuleUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.UpdateTransactionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransactionRule`: TransactionRuleResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.UpdateTransactionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionRuleGuid** | **string** | The unique id for a &#x60;transaction_rule&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransactionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionRuleUpdateRequestBody** | [**TransactionRuleUpdateRequestBody**](TransactionRuleUpdateRequestBody.md) | TransactionRule object to be updated | 

### Return type

[**TransactionRuleResponseBody**](TransactionRuleResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "./openapi"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    userUpdateRequestBody := *openapiclient.NewUserUpdateRequestBody() // UserUpdateRequestBody | User object to be updated (None of these parameters are required, but the user object cannot be empty.)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.UpdateUser(context.Background(), userGuid).UserUpdateRequestBody(userUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: UserResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## VerifyMember

> MemberResponseBody VerifyMember(ctx, memberGuid, userGuid).Execute()

Verify member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformApi.VerifyMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.VerifyMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformApi.VerifyMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyMemberRequest struct via the builder pattern


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

