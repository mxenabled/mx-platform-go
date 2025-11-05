# \RewardsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreditCard**](RewardsAPI.md#CreditCard) | **Get** /credit_card_products/{credit_card_product_guid} | Read a Credit Card Product
[**FetchRewards**](RewardsAPI.md#FetchRewards) | **Post** /users/{user_guid}/members/{member_guid}/fetch_rewards | Fetch Rewards
[**ListRewards**](RewardsAPI.md#ListRewards) | **Get** /users/{user_guid}/members/{member_guid}/rewards | List Rewards
[**ReadRewards**](RewardsAPI.md#ReadRewards) | **Get** /users/{user_guid}/members/{member_guid}/rewards/{reward_guid} | Read Reward



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
    resp, r, err := apiClient.RewardsAPI.CreditCard(context.Background(), creditCardProductGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardsAPI.CreditCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreditCard`: CreditCardProductResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardsAPI.CreditCard`: %v\n", resp)
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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardsAPI.FetchRewards(context.Background(), userGuid, memberGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardsAPI.FetchRewards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRewards`: MemberResponseBody
    fmt.Fprintf(os.Stdout, "Response from `RewardsAPI.FetchRewards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 

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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardsAPI.ListRewards(context.Background(), userGuid, memberGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardsAPI.ListRewards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRewards`: RewardsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `RewardsAPI.ListRewards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 

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
    userGuid := "USR-fa7537f3-48aa-a683-a02a-b18940482f54" // string | The unique identifier for a `user`, beginning with the prefix `USR-`.
    memberGuid := "MBR-7c6f361b-e582-15b6-60c0-358f12466b4b" // string | The unique id for a `member`.
    rewardGuid := "RWD-fa7537f3-48aa-a683-a02a-b324322f54" // string | The unique identifier for the rewards. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardsAPI.ReadRewards(context.Background(), userGuid, memberGuid, rewardGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardsAPI.ReadRewards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadRewards`: RewardResponseBody
    fmt.Fprintf(os.Stdout, "Response from `RewardsAPI.ReadRewards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGuid** | **string** | The unique identifier for a &#x60;user&#x60;, beginning with the prefix &#x60;USR-&#x60;. | 
**memberGuid** | **string** | The unique id for a &#x60;member&#x60;. | 
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

