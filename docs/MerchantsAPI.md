# \MerchantsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMerchants**](MerchantsAPI.md#ListMerchants) | **Get** /merchants | List merchants
[**ReadMerchant**](MerchantsAPI.md#ReadMerchant) | **Get** /merchants/{merchant_guid} | Read merchant
[**ReadMerchantLocation**](MerchantsAPI.md#ReadMerchantLocation) | **Get** /merchant_locations/{merchant_location_guid} | Read merchant location



## ListMerchants

> MerchantsResponseBody ListMerchants(ctx).Name(name).Page(page).RecordsPerPage(recordsPerPage).Execute()

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
    name := "Comcast" // string | This will list only merchants in which the appended string appears. (optional)
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `1000`. If the value exceeds `1000`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerchantsAPI.ListMerchants(context.Background()).Name(name).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantsAPI.ListMerchants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMerchants`: MerchantsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MerchantsAPI.ListMerchants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMerchantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | This will list only merchants in which the appended string appears. | 
 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;1000&#x60;. If the value exceeds &#x60;1000&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

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
    resp, r, err := apiClient.MerchantsAPI.ReadMerchant(context.Background(), merchantGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantsAPI.ReadMerchant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMerchant`: MerchantResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MerchantsAPI.ReadMerchant`: %v\n", resp)
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
    resp, r, err := apiClient.MerchantsAPI.ReadMerchantLocation(context.Background(), merchantLocationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantsAPI.ReadMerchantLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMerchantLocation`: MerchantLocationResponseBody
    fmt.Fprintf(os.Stdout, "Response from `MerchantsAPI.ReadMerchantLocation`: %v\n", resp)
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

