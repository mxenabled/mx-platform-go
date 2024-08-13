# \MxPlatformAPI

All URIs are relative to *https://api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateMember**](MxPlatformAPI.md#AggregateMember) | **Post** /users/{user_guid}/members/{member_guid}/aggregate | Aggregate member
[**CheckBalances**](MxPlatformAPI.md#CheckBalances) | **Post** /users/{user_guid}/members/{member_guid}/check_balance | Check balances
[**CreateCategory**](MxPlatformAPI.md#CreateCategory) | **Post** /users/{user_guid}/categories | Create category
[**CreateManagedAccount**](MxPlatformAPI.md#CreateManagedAccount) | **Post** /users/{user_guid}/managed_members/{member_guid}/accounts | Create managed account
[**CreateManagedMember**](MxPlatformAPI.md#CreateManagedMember) | **Post** /users/{user_guid}/managed_members | Create managed member
[**CreateManagedTransaction**](MxPlatformAPI.md#CreateManagedTransaction) | **Post** /users/{user_guid}/managed_members/{member_guid}/accounts/{account_guid}/transactions | Create managed transaction
[**CreateManualAccount**](MxPlatformAPI.md#CreateManualAccount) | **Post** /users/{user_guid}/accounts | Create manual account
[**CreateMember**](MxPlatformAPI.md#CreateMember) | **Post** /users/{user_guid}/members | Create member
[**CreateTag**](MxPlatformAPI.md#CreateTag) | **Post** /users/{user_guid}/tags | Create tag
[**CreateTagging**](MxPlatformAPI.md#CreateTagging) | **Post** /users/{user_guid}/taggings | Create tagging
[**CreateTransactionRule**](MxPlatformAPI.md#CreateTransactionRule) | **Post** /users/{user_guid}/transaction_rules | Create transaction rule
[**CreateUser**](MxPlatformAPI.md#CreateUser) | **Post** /users | Create user
[**CreditCard**](MxPlatformAPI.md#CreditCard) | **Get** /credit_card_products/{credit_card_product_guid} | Read a Credit Card Product
[**DeleteCategory**](MxPlatformAPI.md#DeleteCategory) | **Delete** /users/{user_guid}/categories/{category_guid} | Delete category
[**DeleteManagedAccount**](MxPlatformAPI.md#DeleteManagedAccount) | **Delete** /users/{user_guid}/managed_members/{member_guid}/accounts/{account_guid} | Delete managed account
[**DeleteManagedMember**](MxPlatformAPI.md#DeleteManagedMember) | **Delete** /users/{user_guid}/managed_members/{member_guid} | Delete managed member
[**DeleteManagedTransaction**](MxPlatformAPI.md#DeleteManagedTransaction) | **Delete** /users/{user_guid}/managed_members/{member_guid}/accounts/{account_guid}/transactions/{transaction_guid} | Delete managed transaction
[**DeleteManualAccount**](MxPlatformAPI.md#DeleteManualAccount) | **Delete** /users/{user_guid}/accounts/{account_guid} | Delete manual account
[**DeleteMember**](MxPlatformAPI.md#DeleteMember) | **Delete** /users/{user_guid}/members/{member_guid} | Delete member
[**DeleteTag**](MxPlatformAPI.md#DeleteTag) | **Delete** /users/{user_guid}/tags/{tag_guid} | Delete tag
[**DeleteTagging**](MxPlatformAPI.md#DeleteTagging) | **Delete** /users/{user_guid}/taggings/{tagging_guid} | Delete tagging
[**DeleteTransactionRule**](MxPlatformAPI.md#DeleteTransactionRule) | **Delete** /users/{user_guid}/transaction_rules/{transaction_rule_guid} | Delete transaction rule
[**DeleteUser**](MxPlatformAPI.md#DeleteUser) | **Delete** /users/{user_guid} | Delete user
[**DeprecatedRequestPaymentProcessorAuthorizationCode**](MxPlatformAPI.md#DeprecatedRequestPaymentProcessorAuthorizationCode) | **Post** /payment_processor_authorization_code | (Deprecated) Request an authorization code.
[**DownloadStatementPDF**](MxPlatformAPI.md#DownloadStatementPDF) | **Get** /users/{user_guid}/members/{member_guid}/statements/{statement_guid}.pdf | Download statement pdf
[**DownloadTaxDocument**](MxPlatformAPI.md#DownloadTaxDocument) | **Get** /users/{user_guid}/members/{member_guid}/tax_documents/{tax_document_guid}.pdf | Download a Tax Document PDF
[**EnhanceTransactions**](MxPlatformAPI.md#EnhanceTransactions) | **Post** /transactions/enhance | Enhance transactions
[**ExtendHistory**](MxPlatformAPI.md#ExtendHistory) | **Post** /users/{user_guid}/members/{member_guid}/extend_history | Extend history
[**FetchRewards**](MxPlatformAPI.md#FetchRewards) | **Post** /users/{user_guid}/members/{member_guid}/fetch_rewards | Fetch Rewards
[**FetchStatements**](MxPlatformAPI.md#FetchStatements) | **Post** /users/{user_guid}/members/{member_guid}/fetch_statements | Fetch statements
[**FetchTaxDocuments**](MxPlatformAPI.md#FetchTaxDocuments) | **Post** /users/{user_guid}/members/{member_guid}/fetch_tax_documents | Fetch Tax Documents
[**IdentifyMember**](MxPlatformAPI.md#IdentifyMember) | **Post** /users/{user_guid}/members/{member_guid}/identify | Identify member
[**ListAccountNumbersByAccount**](MxPlatformAPI.md#ListAccountNumbersByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/account_numbers | List account numbers by account
[**ListAccountNumbersByMember**](MxPlatformAPI.md#ListAccountNumbersByMember) | **Get** /users/{user_guid}/members/{member_guid}/account_numbers | List account numbers by member
[**ListAccountOwnersByMember**](MxPlatformAPI.md#ListAccountOwnersByMember) | **Get** /users/{user_guid}/members/{member_guid}/account_owners | List account owners by member
[**ListCategories**](MxPlatformAPI.md#ListCategories) | **Get** /users/{user_guid}/categories | List categories
[**ListDefaultCategories**](MxPlatformAPI.md#ListDefaultCategories) | **Get** /categories/default | List default categories
[**ListDefaultCategoriesByUser**](MxPlatformAPI.md#ListDefaultCategoriesByUser) | **Get** /users/{user_guid}/categories/default | List default categories by user
[**ListFavoriteInstitutions**](MxPlatformAPI.md#ListFavoriteInstitutions) | **Get** /institutions/favorites | List favorite institutions
[**ListHoldings**](MxPlatformAPI.md#ListHoldings) | **Get** /users/{user_guid}/holdings | List holdings
[**ListHoldingsByAccount**](MxPlatformAPI.md#ListHoldingsByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/holdings | List holdings by account
[**ListHoldingsByMember**](MxPlatformAPI.md#ListHoldingsByMember) | **Get** /users/{user_guid}/members/{member_guid}/holdings | List holdings by member
[**ListInstitutionCredentials**](MxPlatformAPI.md#ListInstitutionCredentials) | **Get** /institutions/{institution_code}/credentials | List institution credentials
[**ListInstitutions**](MxPlatformAPI.md#ListInstitutions) | **Get** /institutions | List institutions
[**ListManagedAccounts**](MxPlatformAPI.md#ListManagedAccounts) | **Get** /users/{user_guid}/managed_members/{member_guid}/accounts | List managed accounts
[**ListManagedInstitutions**](MxPlatformAPI.md#ListManagedInstitutions) | **Get** /managed_institutions | List managed institutions
[**ListManagedMembers**](MxPlatformAPI.md#ListManagedMembers) | **Get** /users/{user_guid}/managed_members | List managed members
[**ListManagedTransactions**](MxPlatformAPI.md#ListManagedTransactions) | **Get** /users/{user_guid}/managed_members/{member_guid}/accounts/{account_guid}/transactions | List managed transactions
[**ListMemberAccounts**](MxPlatformAPI.md#ListMemberAccounts) | **Get** /users/{user_guid}/members/{member_guid}/accounts | List accounts by member
[**ListMemberChallenges**](MxPlatformAPI.md#ListMemberChallenges) | **Get** /users/{user_guid}/members/{member_guid}/challenges | List member challenges
[**ListMemberCredentials**](MxPlatformAPI.md#ListMemberCredentials) | **Get** /users/{user_guid}/members/{member_guid}/credentials | List member credentials
[**ListMembers**](MxPlatformAPI.md#ListMembers) | **Get** /users/{user_guid}/members | List members
[**ListMerchants**](MxPlatformAPI.md#ListMerchants) | **Get** /merchants | List merchants
[**ListRewards**](MxPlatformAPI.md#ListRewards) | **Get** /users/{user_guid}/members/{member_guid}/rewards | List Rewards
[**ListStatementsByMember**](MxPlatformAPI.md#ListStatementsByMember) | **Get** /users/{user_guid}/members/{member_guid}/statements | List statements by member
[**ListTaggings**](MxPlatformAPI.md#ListTaggings) | **Get** /users/{user_guid}/taggings | List taggings
[**ListTags**](MxPlatformAPI.md#ListTags) | **Get** /users/{user_guid}/tags | List tags
[**ListTaxDocuments**](MxPlatformAPI.md#ListTaxDocuments) | **Get** /users/{user_guid}/members/{member_guid}/tax_documents | List Tax Documents
[**ListTransactionRules**](MxPlatformAPI.md#ListTransactionRules) | **Get** /users/{user_guid}/transaction_rules | List transaction rules
[**ListTransactions**](MxPlatformAPI.md#ListTransactions) | **Get** /users/{user_guid}/transactions | List transactions
[**ListTransactionsByAccount**](MxPlatformAPI.md#ListTransactionsByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/transactions | List transactions by account
[**ListTransactionsByMember**](MxPlatformAPI.md#ListTransactionsByMember) | **Get** /users/{user_guid}/members/{member_guid}/transactions | List transactions by member
[**ListTransactionsByTag**](MxPlatformAPI.md#ListTransactionsByTag) | **Get** /users/{user_guid}/tags/{tag_guid}/transactions | List transactions by tag
[**ListUserAccounts**](MxPlatformAPI.md#ListUserAccounts) | **Get** /users/{user_guid}/accounts | List accounts
[**ListUsers**](MxPlatformAPI.md#ListUsers) | **Get** /users | List users
[**ReadAccount**](MxPlatformAPI.md#ReadAccount) | **Get** /users/{user_guid}/accounts/{account_guid} | Read account
[**ReadAccountByMember**](MxPlatformAPI.md#ReadAccountByMember) | **Get** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Read account by member
[**ReadCategory**](MxPlatformAPI.md#ReadCategory) | **Get** /users/{user_guid}/categories/{category_guid} | Read a custom category
[**ReadDefaultCategory**](MxPlatformAPI.md#ReadDefaultCategory) | **Get** /categories/{category_guid} | Read a default category
[**ReadHolding**](MxPlatformAPI.md#ReadHolding) | **Get** /users/{user_guid}/holdings/{holding_guid} | Read holding
[**ReadInstitution**](MxPlatformAPI.md#ReadInstitution) | **Get** /institutions/{institution_code} | Read institution
[**ReadManagedAccount**](MxPlatformAPI.md#ReadManagedAccount) | **Get** /users/{user_guid}/managed_members/{member_guid}/accounts/{account_guid} | Read managed account
[**ReadManagedMember**](MxPlatformAPI.md#ReadManagedMember) | **Get** /users/{user_guid}/managed_members/{member_guid} | Read managed member
[**ReadManagedTransaction**](MxPlatformAPI.md#ReadManagedTransaction) | **Get** /users/{user_guid}/managed_members/{member_guid}/accounts/{account_guid}/transactions/{transaction_guid} | Read managed transaction
[**ReadMember**](MxPlatformAPI.md#ReadMember) | **Get** /users/{user_guid}/members/{member_guid} | Read member
[**ReadMemberStatus**](MxPlatformAPI.md#ReadMemberStatus) | **Get** /users/{user_guid}/members/{member_guid}/status | Read member status
[**ReadMerchant**](MxPlatformAPI.md#ReadMerchant) | **Get** /merchants/{merchant_guid} | Read merchant
[**ReadMerchantLocation**](MxPlatformAPI.md#ReadMerchantLocation) | **Get** /merchant_locations/{merchant_location_guid} | Read merchant location
[**ReadRewards**](MxPlatformAPI.md#ReadRewards) | **Get** /users/{user_guid}/members/{member_guid}/rewards/{reward_guid} | Read Reward
[**ReadStatementByMember**](MxPlatformAPI.md#ReadStatementByMember) | **Get** /users/{user_guid}/members/{member_guid}/statements/{statement_guid} | Read statement by member
[**ReadTag**](MxPlatformAPI.md#ReadTag) | **Get** /users/{user_guid}/tags/{tag_guid} | Read tag
[**ReadTagging**](MxPlatformAPI.md#ReadTagging) | **Get** /users/{user_guid}/taggings/{tagging_guid} | Read tagging
[**ReadTaxDocument**](MxPlatformAPI.md#ReadTaxDocument) | **Get** /users/{user_guid}/members/{member_guid}/tax_documents/{tax_document_guid} | Read a Tax Document
[**ReadTransaction**](MxPlatformAPI.md#ReadTransaction) | **Get** /users/{user_guid}/transactions/{transaction_guid} | Read transaction
[**ReadTransactionRule**](MxPlatformAPI.md#ReadTransactionRule) | **Get** /users/{user_guid}/transaction_rules/{transaction_rule_guid} | Read transaction rule
[**ReadUser**](MxPlatformAPI.md#ReadUser) | **Get** /users/{user_guid} | Read user
[**RequestAuthorizationCode**](MxPlatformAPI.md#RequestAuthorizationCode) | **Post** /authorization_code | Request an authorization code.
[**RequestConnectWidgetURL**](MxPlatformAPI.md#RequestConnectWidgetURL) | **Post** /users/{user_guid}/connect_widget_url | Request connect widget url
[**RequestOAuthWindowURI**](MxPlatformAPI.md#RequestOAuthWindowURI) | **Get** /users/{user_guid}/members/{member_guid}/oauth_window_uri | Request oauth window uri
[**RequestWidgetURL**](MxPlatformAPI.md#RequestWidgetURL) | **Post** /users/{user_guid}/widget_urls | Request widget url
[**ResumeAggregation**](MxPlatformAPI.md#ResumeAggregation) | **Put** /users/{user_guid}/members/{member_guid}/resume | Resume aggregation
[**UpdateAccountByMember**](MxPlatformAPI.md#UpdateAccountByMember) | **Put** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Update account by member
[**UpdateCategory**](MxPlatformAPI.md#UpdateCategory) | **Put** /users/{user_guid}/categories/{category_guid} | Update category
[**UpdateManagedAccount**](MxPlatformAPI.md#UpdateManagedAccount) | **Put** /users/{user_guid}/managed_members/{member_guid}/accounts/{account_guid} | Update managed account
[**UpdateManagedMember**](MxPlatformAPI.md#UpdateManagedMember) | **Put** /users/{user_guid}/managed_members/{member_guid} | Update managed member
[**UpdateManagedTransaction**](MxPlatformAPI.md#UpdateManagedTransaction) | **Put** /users/{user_guid}/managed_members/{member_guid}/accounts/{account_guid}/transactions/{transaction_guid} | Update managed transaction
[**UpdateMember**](MxPlatformAPI.md#UpdateMember) | **Put** /users/{user_guid}/members/{member_guid} | Update member
[**UpdateTag**](MxPlatformAPI.md#UpdateTag) | **Put** /users/{user_guid}/tags/{tag_guid} | Update tag
[**UpdateTagging**](MxPlatformAPI.md#UpdateTagging) | **Put** /users/{user_guid}/taggings/{tagging_guid} | Update tagging
[**UpdateTransaction**](MxPlatformAPI.md#UpdateTransaction) | **Put** /users/{user_guid}/transactions/{transaction_guid} | Update transaction
[**UpdateTransactionRule**](MxPlatformAPI.md#UpdateTransactionRule) | **Put** /users/{user_guid}/transaction_rules/{transaction_rule_guid} | Update transaction_rule
[**UpdateUser**](MxPlatformAPI.md#UpdateUser) | **Put** /users/{user_guid} | Update user
[**VerifyMember**](MxPlatformAPI.md#VerifyMember) | **Post** /users/{user_guid}/members/{member_guid}/verify | Verify member



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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.AggregateMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.AggregateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AggregateMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.AggregateMember`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CheckBalances(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CheckBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckBalances`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CheckBalances`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    categoryCreateRequestBody := *openapiclient.NewCategoryCreateRequestBody() // CategoryCreateRequestBody | Custom category object to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CreateCategory(context.Background(), userGuid).CategoryCreateRequestBody(categoryCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CreateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CreateCategory`: %v\n", resp)
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

> AccountResponseBody CreateManagedAccount(ctx, memberGuid, userGuid).ManagedAccountCreateRequestBody(managedAccountCreateRequestBody).Execute()

Create managed account



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    managedAccountCreateRequestBody := *openapiclient.NewManagedAccountCreateRequestBody() // ManagedAccountCreateRequestBody | Managed account to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CreateManagedAccount(context.Background(), memberGuid, userGuid).ManagedAccountCreateRequestBody(managedAccountCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CreateManagedAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagedAccount`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CreateManagedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    managedMemberCreateRequestBody := *openapiclient.NewManagedMemberCreateRequestBody() // ManagedMemberCreateRequestBody | Managed member to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CreateManagedMember(context.Background(), userGuid).ManagedMemberCreateRequestBody(managedMemberCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CreateManagedMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagedMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CreateManagedMember`: %v\n", resp)
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

> TransactionResponseBody CreateManagedTransaction(ctx, accountGuid, memberGuid, userGuid).ManagedTransactionCreateRequestBody(managedTransactionCreateRequestBody).Execute()

Create managed transaction



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    managedTransactionCreateRequestBody := *openapiclient.NewManagedTransactionCreateRequestBody() // ManagedTransactionCreateRequestBody | Managed transaction to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CreateManagedTransaction(context.Background(), accountGuid, memberGuid, userGuid).ManagedTransactionCreateRequestBody(managedTransactionCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CreateManagedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagedTransaction`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CreateManagedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    accountCreateRequestBody := *openapiclient.NewAccountCreateRequestBody() // AccountCreateRequestBody | Manual account object to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CreateManualAccount(context.Background(), userGuid).AccountCreateRequestBody(accountCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CreateManualAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManualAccount`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CreateManualAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberCreateRequestBody := *openapiclient.NewMemberCreateRequestBody() // MemberCreateRequestBody | Member object to be created with optional parameters (id and metadata) and required parameters (credentials and institution_code)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CreateMember(context.Background(), userGuid).MemberCreateRequestBody(memberCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CreateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CreateMember`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    tagCreateRequestBody := *openapiclient.NewTagCreateRequestBody() // TagCreateRequestBody | Tag object to be created with required parameters (tag_guid)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CreateTag(context.Background(), userGuid).TagCreateRequestBody(tagCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CreateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTag`: TagResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CreateTag`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    taggingCreateRequestBody := *openapiclient.NewTaggingCreateRequestBody() // TaggingCreateRequestBody | Tagging object to be created with required parameters (tag_guid and transaction_guid)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CreateTagging(context.Background(), userGuid).TaggingCreateRequestBody(taggingCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CreateTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagging`: TaggingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CreateTagging`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    transactionRuleCreateRequestBody := *openapiclient.NewTransactionRuleCreateRequestBody() // TransactionRuleCreateRequestBody | TransactionRule object to be created with optional parameters (description) and required parameters (category_guid and match_description)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CreateTransactionRule(context.Background(), userGuid).TransactionRuleCreateRequestBody(transactionRuleCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CreateTransactionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransactionRule`: TransactionRuleResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CreateTransactionRule`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userCreateRequestBody := *openapiclient.NewUserCreateRequestBody() // UserCreateRequestBody | User object to be created. (None of these parameters are required, but the user object cannot be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CreateUser(context.Background()).UserCreateRequestBody(userCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: UserResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CreateUser`: %v\n", resp)
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


## CreditCard

> CreditCardProductResponse CreditCard(ctx, creditCardProductGuid).Execute()

Read a Credit Card Product



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
    creditCardProductGuid := "credit_card_product_guid" // string | The required `credit_card_product_guid` can be found on the `account` object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.CreditCard(context.Background(), creditCardProductGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.CreditCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreditCard`: CreditCardProductResponse
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.CreditCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditCardProductGuid** | **string** | The required &#x60;credit_card_product_guid&#x60; can be found on the &#x60;account&#x60; object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreditCardProductResponse**](CreditCardProductResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MxPlatformAPI.DeleteCategory(context.Background(), categoryGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DeleteCategory``: %v\n", err)
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

> DeleteManagedAccount(ctx, accountGuid, memberGuid, userGuid).Execute()

Delete managed account



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MxPlatformAPI.DeleteManagedAccount(context.Background(), accountGuid, memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DeleteManagedAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MxPlatformAPI.DeleteManagedMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DeleteManagedMember``: %v\n", err)
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

> DeleteManagedTransaction(ctx, accountGuid, memberGuid, transactionGuid, userGuid).Execute()

Delete managed transaction



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
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MxPlatformAPI.DeleteManagedTransaction(context.Background(), accountGuid, memberGuid, transactionGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DeleteManagedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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


## DeleteManualAccount

> DeleteManualAccount(ctx, accountGuid, userGuid).Execute()

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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MxPlatformAPI.DeleteManualAccount(context.Background(), accountGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DeleteManualAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManualAccountRequest struct via the builder pattern


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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MxPlatformAPI.DeleteMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DeleteMember``: %v\n", err)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    tagGuid := "TAG-aef36e72-6294-4c38-844d-e573e80aed52" // string | The unique id for a `tag`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MxPlatformAPI.DeleteTag(context.Background(), tagGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DeleteTag``: %v\n", err)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    taggingGuid := "TGN-007f5486-17e1-45fc-8b87-8f03984430fe" // string | The unique id for a `tagging`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MxPlatformAPI.DeleteTagging(context.Background(), taggingGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DeleteTagging``: %v\n", err)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    transactionRuleGuid := "TXR-a080e0f9-a2d4-4d6f-9e03-672cc357a4d3" // string | The unique id for a `transaction_rule`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MxPlatformAPI.DeleteTransactionRule(context.Background(), transactionRuleGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DeleteTransactionRule``: %v\n", err)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MxPlatformAPI.DeleteUser(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DeleteUser``: %v\n", err)
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


## DeprecatedRequestPaymentProcessorAuthorizationCode

> PaymentProcessorAuthorizationCodeResponseBody DeprecatedRequestPaymentProcessorAuthorizationCode(ctx).PaymentProcessorAuthorizationCodeRequestBody(paymentProcessorAuthorizationCodeRequestBody).Execute()

(Deprecated) Request an authorization code.



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
    paymentProcessorAuthorizationCodeRequestBody := *openapiclient.NewPaymentProcessorAuthorizationCodeRequestBody() // PaymentProcessorAuthorizationCodeRequestBody | The scope for the authorization code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.DeprecatedRequestPaymentProcessorAuthorizationCode(context.Background()).PaymentProcessorAuthorizationCodeRequestBody(paymentProcessorAuthorizationCodeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DeprecatedRequestPaymentProcessorAuthorizationCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedRequestPaymentProcessorAuthorizationCode`: PaymentProcessorAuthorizationCodeResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.DeprecatedRequestPaymentProcessorAuthorizationCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedRequestPaymentProcessorAuthorizationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentProcessorAuthorizationCodeRequestBody** | [**PaymentProcessorAuthorizationCodeRequestBody**](PaymentProcessorAuthorizationCodeRequestBody.md) | The scope for the authorization code. | 

### Return type

[**PaymentProcessorAuthorizationCodeResponseBody**](PaymentProcessorAuthorizationCodeResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    statementGuid := "STA-737a344b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for a `statement`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.DownloadStatementPDF(context.Background(), memberGuid, statementGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DownloadStatementPDF``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadStatementPDF`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.DownloadStatementPDF`: %v\n", resp)
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


## DownloadTaxDocument

> *os.File DownloadTaxDocument(ctx, taxDocumentGuid, memberGuid, userGuid).Execute()

Download a Tax Document PDF



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
    taxDocumentGuid := "TAX-987dfds1b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `tax_document`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.DownloadTaxDocument(context.Background(), taxDocumentGuid, memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.DownloadTaxDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadTaxDocument`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.DownloadTaxDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxDocumentGuid** | **string** | The unique id for a &#x60;tax_document&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadTaxDocumentRequest struct via the builder pattern


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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    enhanceTransactionsRequestBody := *openapiclient.NewEnhanceTransactionsRequestBody() // EnhanceTransactionsRequestBody | Transaction object to be enhanced

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.EnhanceTransactions(context.Background()).EnhanceTransactionsRequestBody(enhanceTransactionsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.EnhanceTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnhanceTransactions`: EnhanceTransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.EnhanceTransactions`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique identifier for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ExtendHistory(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ExtendHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtendHistory`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ExtendHistory`: %v\n", resp)
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


## FetchRewards

> MemberResponseBody FetchRewards(ctx, userGuid, memberGuid).Execute()

Fetch Rewards



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
    memberGuid := "MBR-fa7537f3-48aa-a683-a02a-b18345562f54" // string | The unique identifier for the member. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.FetchRewards(context.Background(), userGuid, memberGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.FetchRewards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRewards`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.FetchRewards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**memberGuid** | **string** | The unique identifier for the member. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchRewardsRequest struct via the builder pattern


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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.FetchStatements(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.FetchStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchStatements`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.FetchStatements`: %v\n", resp)
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


## FetchTaxDocuments

> MemberResponseBody FetchTaxDocuments(ctx, memberGuid, userGuid).Execute()

Fetch Tax Documents



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.FetchTaxDocuments(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.FetchTaxDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTaxDocuments`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.FetchTaxDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTaxDocumentsRequest struct via the builder pattern


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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.IdentifyMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.IdentifyMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentifyMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.IdentifyMember`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    accountGuid := "ACT-06d7f44b-caae-0f6e-1384-01f52e75dcb1" // string | The unique id for an `account`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListAccountNumbersByAccount(context.Background(), accountGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListAccountNumbersByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountNumbersByAccount`: AccountNumbersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListAccountNumbersByAccount`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListAccountNumbersByMember(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListAccountNumbersByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountNumbersByMember`: AccountNumbersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListAccountNumbersByMember`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListAccountOwnersByMember(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListAccountOwnersByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountOwnersByMember`: AccountOwnersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListAccountOwnersByMember`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListCategories(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCategories`: CategoriesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListCategories`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListDefaultCategories(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListDefaultCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDefaultCategories`: CategoriesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListDefaultCategories`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListDefaultCategoriesByUser(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListDefaultCategoriesByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDefaultCategoriesByUser`: CategoriesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListDefaultCategoriesByUser`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListFavoriteInstitutions(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListFavoriteInstitutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFavoriteInstitutions`: InstitutionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListFavoriteInstitutions`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    fromDate := "2015-09-20" // string | Filter holdings from this date. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)
    toDate := "2019-10-20" // string | Filter holdings to this date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListHoldings(context.Background(), userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListHoldings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHoldings`: HoldingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListHoldings`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
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
    resp, r, err := apiClient.MxPlatformAPI.ListHoldingsByAccount(context.Background(), accountGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListHoldingsByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHoldingsByAccount`: HoldingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListHoldingsByAccount`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
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
    resp, r, err := apiClient.MxPlatformAPI.ListHoldingsByMember(context.Background(), memberGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListHoldingsByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHoldingsByMember`: HoldingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListHoldingsByMember`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    institutionCode := "chase" // string | The institution_code of the institution.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListInstitutionCredentials(context.Background(), institutionCode).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListInstitutionCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstitutionCredentials`: CredentialsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListInstitutionCredentials`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
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
    resp, r, err := apiClient.MxPlatformAPI.ListInstitutions(context.Background()).Name(name).Page(page).RecordsPerPage(recordsPerPage).SupportsAccountIdentification(supportsAccountIdentification).SupportsAccountStatement(supportsAccountStatement).SupportsAccountVerification(supportsAccountVerification).SupportsTransactionHistory(supportsTransactionHistory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListInstitutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstitutions`: InstitutionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListInstitutions`: %v\n", resp)
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

> AccountsResponseBody ListManagedAccounts(ctx, memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List managed accounts



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListManagedAccounts(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListManagedAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManagedAccounts`: AccountsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListManagedAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListManagedInstitutions(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListManagedInstitutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManagedInstitutions`: InstitutionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListManagedInstitutions`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListManagedMembers(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListManagedMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManagedMembers`: MembersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListManagedMembers`: %v\n", resp)
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

> TransactionsResponseBody ListManagedTransactions(ctx, accountGuid, memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List managed transactions



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListManagedTransactions(context.Background(), accountGuid, memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListManagedTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManagedTransactions`: TransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListManagedTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    memberIsManagedByUser := true // bool | List only accounts whose member is managed by the user. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListMemberAccounts(context.Background(), userGuid, memberGuid).MemberIsManagedByUser(memberIsManagedByUser).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListMemberAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMemberAccounts`: AccountsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListMemberAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMemberAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **memberIsManagedByUser** | **bool** | List only accounts whose member is managed by the user. | 
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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListMemberChallenges(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListMemberChallenges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMemberChallenges`: ChallengesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListMemberChallenges`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListMemberCredentials(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListMemberCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMemberCredentials`: CredentialsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListMemberCredentials`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListMembers(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMembers`: MembersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListMembers`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListMerchants(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListMerchants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMerchants`: MerchantsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListMerchants`: %v\n", resp)
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


## ListRewards

> RewardsResponseBody ListRewards(ctx, userGuid, memberGuid).Execute()

List Rewards



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
    memberGuid := "MBR-fa7537f3-48aa-a683-a02a-b18345562f54" // string | The unique identifier for the member. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListRewards(context.Background(), userGuid, memberGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListRewards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRewards`: RewardsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListRewards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**memberGuid** | **string** | The unique identifier for the member. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRewardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RewardsResponseBody**](RewardsResponseBody.md)

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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListStatementsByMember(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListStatementsByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStatementsByMember`: StatementsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListStatementsByMember`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListTaggings(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListTaggings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaggings`: TaggingsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListTaggings`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListTags(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTags`: TagsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListTags`: %v\n", resp)
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


## ListTaxDocuments

> TaxDocumentsResponseBody ListTaxDocuments(ctx, memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()

List Tax Documents



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListTaxDocuments(context.Background(), memberGuid, userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListTaxDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaxDocuments`: TaxDocumentsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListTaxDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTaxDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 

### Return type

[**TaxDocumentsResponseBody**](TaxDocumentsResponseBody.md)

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListTransactionRules(context.Background(), userGuid).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListTransactionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionRules`: TransactionRulesResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListTransactionRules`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    fromDate := "2015-09-20" // string | Filter transactions from this date. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)
    toDate := "2019-10-20" // string | Filter transactions to this date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListTransactions(context.Background(), userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactions`: TransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListTransactions`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
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
    resp, r, err := apiClient.MxPlatformAPI.ListTransactionsByAccount(context.Background(), accountGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListTransactionsByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByAccount`: TransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListTransactionsByAccount`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
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
    resp, r, err := apiClient.MxPlatformAPI.ListTransactionsByMember(context.Background(), memberGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListTransactionsByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByMember`: TransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListTransactionsByMember`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
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
    resp, r, err := apiClient.MxPlatformAPI.ListTransactionsByTag(context.Background(), tagGuid, userGuid).FromDate(fromDate).Page(page).RecordsPerPage(recordsPerPage).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListTransactionsByTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionsByTag`: TransactionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListTransactionsByTag`: %v\n", resp)
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

> AccountsResponseBody ListUserAccounts(ctx, userGuid).MemberIsManagedByUser(memberIsManagedByUser).Page(page).IsManual(isManual).RecordsPerPage(recordsPerPage).Execute()

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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberIsManagedByUser := true // bool | List only accounts whose member is managed by the user. (optional)
    page := int32(1) // int32 | Specify current page. (optional)
    isManual := true // bool | List only accounts that were manually created. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListUserAccounts(context.Background(), userGuid).MemberIsManagedByUser(memberIsManagedByUser).Page(page).IsManual(isManual).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListUserAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserAccounts`: AccountsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListUserAccounts`: %v\n", resp)
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

 **memberIsManagedByUser** | **bool** | List only accounts whose member is managed by the user. | 
 **page** | **int32** | Specify current page. | 
 **isManual** | **bool** | List only accounts that were manually created. | 
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
    page := int32(1) // int32 | Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | Specify records per page. (optional)
    id := "u-12324-abdc" // string | The user `id` to search for. (optional)
    email := "example@example.com" // string | The user `email` to search for. (optional)
    isDisabled := true // bool | Search for users that are diabled. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ListUsers(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Id(id).Email(email).IsDisabled(isDisabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: UsersResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Specify current page. | 
 **recordsPerPage** | **int32** | Specify records per page. | 
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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadAccount(context.Background(), accountGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAccount`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadAccount`: %v\n", resp)
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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadAccountByMember(context.Background(), accountGuid, memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadAccountByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAccountByMember`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadAccountByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadCategory(context.Background(), categoryGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadCategory`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    categoryGuid := "CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874" // string | The unique id for a `category`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadDefaultCategory(context.Background(), categoryGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadDefaultCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDefaultCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadDefaultCategory`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    holdingGuid := "HOL-d65683e8-9eab-26bb-bcfd-ced159c9abe2" // string | The unique id for a `holding`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadHolding(context.Background(), holdingGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadHolding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadHolding`: HoldingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadHolding`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    institutionCode := "chase" // string | The institution_code of the institution.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadInstitution(context.Background(), institutionCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadInstitution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadInstitution`: InstitutionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadInstitution`: %v\n", resp)
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

> AccountResponseBody ReadManagedAccount(ctx, accountGuid, memberGuid, userGuid).Execute()

Read managed account



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadManagedAccount(context.Background(), accountGuid, memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadManagedAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadManagedAccount`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadManagedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadManagedMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadManagedMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadManagedMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadManagedMember`: %v\n", resp)
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

> TransactionResponseBody ReadManagedTransaction(ctx, accountGuid, memberGuid, transactionGuid, userGuid).Execute()

Read managed transaction



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
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadManagedTransaction(context.Background(), accountGuid, memberGuid, transactionGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadManagedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadManagedTransaction`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadManagedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadMember`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadMemberStatus(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadMemberStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMemberStatus`: MemberStatusResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadMemberStatus`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    merchantGuid := "MCH-7ed79542-884d-2b1b-dd74-501c5cc9d25b" // string | The unique id for a `merchant`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadMerchant(context.Background(), merchantGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadMerchant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMerchant`: MerchantResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadMerchant`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    merchantLocationGuid := "MCH-09466f0a-fb58-9d1a-bae2-2af0afbea621" // string | The unique id for a `merchant_location`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadMerchantLocation(context.Background(), merchantLocationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadMerchantLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMerchantLocation`: MerchantLocationResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadMerchantLocation`: %v\n", resp)
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


## ReadRewards

> RewardResponseBody ReadRewards(ctx, userGuid, memberGuid, rewardGuid).Execute()

Read Reward



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
    memberGuid := "MBR-fa7537f3-48aa-a683-a02a-b18345562f54" // string | The unique identifier for the member. Defined by MX.
    rewardGuid := "RWD-fa7537f3-48aa-a683-a02a-b324322f54" // string | The unique identifier for the rewards. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadRewards(context.Background(), userGuid, memberGuid, rewardGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadRewards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadRewards`: RewardResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadRewards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 
**memberGuid** | **string** | The unique identifier for the member. Defined by MX. | 
**rewardGuid** | **string** | The unique identifier for the rewards. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadRewardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RewardResponseBody**](RewardResponseBody.md)

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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadStatementByMember(context.Background(), memberGuid, statementGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadStatementByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadStatementByMember`: StatementResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadStatementByMember`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    tagGuid := "TAG-aef36e72-6294-4c38-844d-e573e80aed52" // string | The unique id for a `tag`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadTag(context.Background(), tagGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTag`: TagResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadTag`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    taggingGuid := "TGN-007f5486-17e1-45fc-8b87-8f03984430fe" // string | The unique id for a `tagging`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadTagging(context.Background(), taggingGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTagging`: TaggingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadTagging`: %v\n", resp)
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


## ReadTaxDocument

> TaxDocumentResponseBody ReadTaxDocument(ctx, taxDocumentGuid, memberGuid, userGuid).Execute()

Read a Tax Document



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
    taxDocumentGuid := "TAX-987dfds1b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `tax_document`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadTaxDocument(context.Background(), taxDocumentGuid, memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadTaxDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTaxDocument`: TaxDocumentResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadTaxDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxDocumentGuid** | **string** | The unique id for a &#x60;tax_document&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTaxDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TaxDocumentResponseBody**](TaxDocumentResponseBody.md)

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadTransaction(context.Background(), transactionGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTransaction`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadTransaction`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    transactionRuleGuid := "TXR-a080e0f9-a2d4-4d6f-9e03-672cc357a4d3" // string | The unique id for a `transaction_rule`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadTransactionRule(context.Background(), transactionRuleGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadTransactionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTransactionRule`: TransactionRuleResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadTransactionRule`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ReadUser(context.Background(), userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ReadUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUser`: UserResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ReadUser`: %v\n", resp)
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


## RequestAuthorizationCode

> AuthorizationCodeResponseBody RequestAuthorizationCode(ctx).AuthorizationCodeRequestBody(authorizationCodeRequestBody).Execute()

Request an authorization code.



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
    authorizationCodeRequestBody := *openapiclient.NewAuthorizationCodeRequestBody() // AuthorizationCodeRequestBody | The scope for the authorization code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.RequestAuthorizationCode(context.Background()).AuthorizationCodeRequestBody(authorizationCodeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.RequestAuthorizationCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestAuthorizationCode`: AuthorizationCodeResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.RequestAuthorizationCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestAuthorizationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationCodeRequestBody** | [**AuthorizationCodeRequestBody**](AuthorizationCodeRequestBody.md) | The scope for the authorization code. | 

### Return type

[**AuthorizationCodeResponseBody**](AuthorizationCodeResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    connectWidgetRequestBody := *openapiclient.NewConnectWidgetRequestBody() // ConnectWidgetRequestBody | Optional config options for WebView (is_mobile_webview, current_institution_code, current_member_guid, update_credentials)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.RequestConnectWidgetURL(context.Background(), userGuid).ConnectWidgetRequestBody(connectWidgetRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.RequestConnectWidgetURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestConnectWidgetURL`: ConnectWidgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.RequestConnectWidgetURL`: %v\n", resp)
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

> OAuthWindowResponseBody RequestOAuthWindowURI(ctx, memberGuid, userGuid).ClientRedirectUrl(clientRedirectUrl).EnableApp2app(enableApp2app).ReferralSource(referralSource).SkipAggregation(skipAggregation).UiMessageWebviewUrlScheme(uiMessageWebviewUrlScheme).Execute()

Request oauth window uri



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    clientRedirectUrl := "https://mx.com" // string | A URL that MX will redirect to at the end of OAuth with additional query parameters. Only available with `referral_source=APP`. (optional)
    enableApp2app := "false" // string | This indicates whether OAuth app2app behavior is enabled for institutions that support it. Defaults to `true`. This setting is not persistent. (optional)
    referralSource := "APP" // string | Must be either `BROWSER` or `APP` depending on the implementation. Defaults to `BROWSER`. (optional)
    skipAggregation := false // bool | Setting this parameter to `true` will prevent the member from automatically aggregating after being redirected from the authorization page. (optional)
    uiMessageWebviewUrlScheme := "mx" // string | A scheme for routing the user back to the application state they were previously in. Only available with `referral_source=APP`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.RequestOAuthWindowURI(context.Background(), memberGuid, userGuid).ClientRedirectUrl(clientRedirectUrl).EnableApp2app(enableApp2app).ReferralSource(referralSource).SkipAggregation(skipAggregation).UiMessageWebviewUrlScheme(uiMessageWebviewUrlScheme).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.RequestOAuthWindowURI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestOAuthWindowURI`: OAuthWindowResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.RequestOAuthWindowURI`: %v\n", resp)
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


 **clientRedirectUrl** | **string** | A URL that MX will redirect to at the end of OAuth with additional query parameters. Only available with &#x60;referral_source&#x3D;APP&#x60;. | 
 **enableApp2app** | **string** | This indicates whether OAuth app2app behavior is enabled for institutions that support it. Defaults to &#x60;true&#x60;. This setting is not persistent. | 
 **referralSource** | **string** | Must be either &#x60;BROWSER&#x60; or &#x60;APP&#x60; depending on the implementation. Defaults to &#x60;BROWSER&#x60;. | 
 **skipAggregation** | **bool** | Setting this parameter to &#x60;true&#x60; will prevent the member from automatically aggregating after being redirected from the authorization page. | 
 **uiMessageWebviewUrlScheme** | **string** | A scheme for routing the user back to the application state they were previously in. Only available with &#x60;referral_source&#x3D;APP&#x60;. | 

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    widgetRequestBody := *openapiclient.NewWidgetRequestBody() // WidgetRequestBody | The widget url configuration options.
    acceptLanguage := "en-US" // string | The desired language of the widget. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.RequestWidgetURL(context.Background(), userGuid).WidgetRequestBody(widgetRequestBody).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.RequestWidgetURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestWidgetURL`: WidgetResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.RequestWidgetURL`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberResumeRequestBody := *openapiclient.NewMemberResumeRequestBody() // MemberResumeRequestBody | Member object with MFA challenge answers

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.ResumeAggregation(context.Background(), memberGuid, userGuid).MemberResumeRequestBody(memberResumeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.ResumeAggregation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeAggregation`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.ResumeAggregation`: %v\n", resp)
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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    accountUpdateRequestBody := *openapiclient.NewAccountUpdateRequestBody() // AccountUpdateRequestBody | Account object to be created with optional parameters (is_hidden)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.UpdateAccountByMember(context.Background(), accountGuid, memberGuid, userGuid).AccountUpdateRequestBody(accountUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.UpdateAccountByMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountByMember`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.UpdateAccountByMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    categoryGuid := "CAT-7829f71c-2e8c-afa5-2f55-fa3634b89874" // string | The unique id for a `category`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    categoryUpdateRequestBody := *openapiclient.NewCategoryUpdateRequestBody() // CategoryUpdateRequestBody | Category object to be updated (While no single parameter is required, the `category` object cannot be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.UpdateCategory(context.Background(), categoryGuid, userGuid).CategoryUpdateRequestBody(categoryUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.UpdateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCategory`: CategoryResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.UpdateCategory`: %v\n", resp)
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

> AccountResponseBody UpdateManagedAccount(ctx, accountGuid, memberGuid, userGuid).ManagedAccountUpdateRequestBody(managedAccountUpdateRequestBody).Execute()

Update managed account



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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    managedAccountUpdateRequestBody := *openapiclient.NewManagedAccountUpdateRequestBody() // ManagedAccountUpdateRequestBody | Managed account object to be updated (While no single parameter is required, the request body can't be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.UpdateManagedAccount(context.Background(), accountGuid, memberGuid, userGuid).ManagedAccountUpdateRequestBody(managedAccountUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.UpdateManagedAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedAccount`: AccountResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.UpdateManagedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    managedMemberUpdateRequestBody := *openapiclient.NewManagedMemberUpdateRequestBody() // ManagedMemberUpdateRequestBody | Managed member object to be updated (While no single parameter is required, the request body can't be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.UpdateManagedMember(context.Background(), memberGuid, userGuid).ManagedMemberUpdateRequestBody(managedMemberUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.UpdateManagedMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.UpdateManagedMember`: %v\n", resp)
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

> TransactionResponseBody UpdateManagedTransaction(ctx, accountGuid, memberGuid, transactionGuid, userGuid).ManagedTransactionUpdateRequestBody(managedTransactionUpdateRequestBody).Execute()

Update managed transaction



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
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    managedTransactionUpdateRequestBody := *openapiclient.NewManagedTransactionUpdateRequestBody() // ManagedTransactionUpdateRequestBody | Managed transaction object to be updated (While no single parameter is required, the request body can't be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.UpdateManagedTransaction(context.Background(), accountGuid, memberGuid, transactionGuid, userGuid).ManagedTransactionUpdateRequestBody(managedTransactionUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.UpdateManagedTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagedTransaction`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.UpdateManagedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountGuid** | **string** | The unique id for an &#x60;account&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
**transactionGuid** | **string** | The unique id for a &#x60;transaction&#x60;. | 
**userGuid** | **string** | The unique id for a &#x60;user&#x60;. | 

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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    memberUpdateRequestBody := *openapiclient.NewMemberUpdateRequestBody() // MemberUpdateRequestBody | Member object to be updated (While no single parameter is required, the request body can't be empty)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.UpdateMember(context.Background(), memberGuid, userGuid).MemberUpdateRequestBody(memberUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.UpdateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.UpdateMember`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    tagGuid := "TAG-aef36e72-6294-4c38-844d-e573e80aed52" // string | The unique id for a `tag`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    tagUpdateRequestBody := *openapiclient.NewTagUpdateRequestBody() // TagUpdateRequestBody | Tag object to be updated with required parameter (tag_guid)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.UpdateTag(context.Background(), tagGuid, userGuid).TagUpdateRequestBody(tagUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.UpdateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTag`: TagResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.UpdateTag`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    taggingGuid := "TGN-007f5486-17e1-45fc-8b87-8f03984430fe" // string | The unique id for a `tagging`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    taggingUpdateRequestBody := *openapiclient.NewTaggingUpdateRequestBody() // TaggingUpdateRequestBody | Tagging object to be updated with required parameter (tag_guid)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.UpdateTagging(context.Background(), taggingGuid, userGuid).TaggingUpdateRequestBody(taggingUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.UpdateTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagging`: TaggingResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.UpdateTagging`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    transactionGuid := "TRN-810828b0-5210-4878-9bd3-f4ce514f90c4" // string | The unique id for a `transaction`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    transactionUpdateRequestBody := *openapiclient.NewTransactionUpdateRequestBody() // TransactionUpdateRequestBody | Transaction object to be updated with a new description

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.UpdateTransaction(context.Background(), transactionGuid, userGuid).TransactionUpdateRequestBody(transactionUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.UpdateTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransaction`: TransactionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.UpdateTransaction`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    transactionRuleGuid := "TXR-a080e0f9-a2d4-4d6f-9e03-672cc357a4d3" // string | The unique id for a `transaction_rule`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    transactionRuleUpdateRequestBody := *openapiclient.NewTransactionRuleUpdateRequestBody() // TransactionRuleUpdateRequestBody | TransactionRule object to be updated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.UpdateTransactionRule(context.Background(), transactionRuleGuid, userGuid).TransactionRuleUpdateRequestBody(transactionRuleUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.UpdateTransactionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransactionRule`: TransactionRuleResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.UpdateTransactionRule`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.
    userUpdateRequestBody := *openapiclient.NewUserUpdateRequestBody() // UserUpdateRequestBody | User object to be updated (None of these parameters are required, but the user object cannot be empty.)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.UpdateUser(context.Background(), userGuid).UserUpdateRequestBody(userUpdateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: UserResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.UpdateUser`: %v\n", resp)
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
    openapiclient "github.com/mxenabled/mx-platform-go"
)

func main() {
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique id for a `user`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MxPlatformAPI.VerifyMember(context.Background(), memberGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformAPI.VerifyMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyMember`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MxPlatformAPI.VerifyMember`: %v\n", resp)
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

