# \TransactionsAPI

All URIs are relative to *https://api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersUserGuidAccountsAccountGuidTransactionsPost**](TransactionsAPI.md#UsersUserGuidAccountsAccountGuidTransactionsPost) | **Post** /users/{user_guid}/accounts/{account_guid}/transactions | Create manual transaction



## UsersUserGuidAccountsAccountGuidTransactionsPost

> TransactionCreateResponseBody UsersUserGuidAccountsAccountGuidTransactionsPost(ctx, userGuid, accountGuid).TransactionCreateRequestBody(transactionCreateRequestBody).Execute()

Create manual transaction



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
    accountGuid := "accountGuid_example" // string | The unique identifier for the account.
    transactionCreateRequestBody := *openapiclient.NewTransactionCreateRequestBody() // TransactionCreateRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.UsersUserGuidAccountsAccountGuidTransactionsPost(context.Background(), userGuid, accountGuid).TransactionCreateRequestBody(transactionCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.UsersUserGuidAccountsAccountGuidTransactionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUserGuidAccountsAccountGuidTransactionsPost`: TransactionCreateResponseBody
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.UsersUserGuidAccountsAccountGuidTransactionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for the user. | 
**accountGuid** | **string** | The unique identifier for the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserGuidAccountsAccountGuidTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionCreateRequestBody** | [**TransactionCreateRequestBody**](TransactionCreateRequestBody.md) |  | 

### Return type

[**TransactionCreateResponseBody**](TransactionCreateResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

