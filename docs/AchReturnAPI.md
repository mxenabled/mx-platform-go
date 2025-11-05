# \AchReturnAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateACHReturn**](AchReturnAPI.md#CreateACHReturn) | **Post** /ach_returns | Create ACH Return
[**ListACHRetruns**](AchReturnAPI.md#ListACHRetruns) | **Get** /ach_returns | List ACH Returns
[**ReadACHRetrun**](AchReturnAPI.md#ReadACHRetrun) | **Get** /ach_returns/{ach_return_guid} | Read ACH Return



## CreateACHReturn

> ACHReturnResponseBody CreateACHReturn(ctx).ACHReturnCreateRequestBody(aCHReturnCreateRequestBody).Execute()

Create ACH Return



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
    aCHReturnCreateRequestBody := *openapiclient.NewACHReturnCreateRequestBody() // ACHReturnCreateRequestBody | ACH return object to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AchReturnAPI.CreateACHReturn(context.Background()).ACHReturnCreateRequestBody(aCHReturnCreateRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AchReturnAPI.CreateACHReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateACHReturn`: ACHReturnResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AchReturnAPI.CreateACHReturn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateACHReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aCHReturnCreateRequestBody** | [**ACHReturnCreateRequestBody**](ACHReturnCreateRequestBody.md) | ACH return object to be created. | 

### Return type

[**ACHReturnResponseBody**](ACHReturnResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListACHRetruns

> ACHReturnsResponseBody ListACHRetruns(ctx).InstitutionGuid(institutionGuid).ReturnedAt(returnedAt).ResolvedStatusAt(resolvedStatusAt).ReturnCode(returnCode).ReturnStatus(returnStatus).Page(page).RecordsPerPage(recordsPerPage).Execute()

List ACH Returns



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
    institutionGuid := "institutionGuid_example" // string | The identifier for the institution associated with the ACH return. Defined by MX. (optional)
    returnedAt := "2025-02-13T18:09:00+00:00" // string | The date and time when the return was reported by the Receiving Financial Depository Institution (RDFI) in ISO 8601 format without timestamp. (optional)
    resolvedStatusAt := "2025-02-13T18:09:00+00:00" // string | The date and time when the return was resolved by the Receiving Financial Depository Institution (RDFI) in ISO 8601 format without timestamp (optional)
    returnCode := "returnCode_example" // string | The associated ACH return code and notice of change code. See [Return Codes](/api-reference/platform-api/reference/ach-return-fields#return-codes) for a complete list. (optional)
    returnStatus := "SUBMITTED" // string | The status of the return. See [Return Statuses](/api-reference/platform-api/reference/ach-return-fields#return-status) for a complete list. (optional)
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AchReturnAPI.ListACHRetruns(context.Background()).InstitutionGuid(institutionGuid).ReturnedAt(returnedAt).ResolvedStatusAt(resolvedStatusAt).ReturnCode(returnCode).ReturnStatus(returnStatus).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AchReturnAPI.ListACHRetruns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListACHRetruns`: ACHReturnsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AchReturnAPI.ListACHRetruns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListACHRetrunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **institutionGuid** | **string** | The identifier for the institution associated with the ACH return. Defined by MX. | 
 **returnedAt** | **string** | The date and time when the return was reported by the Receiving Financial Depository Institution (RDFI) in ISO 8601 format without timestamp. | 
 **resolvedStatusAt** | **string** | The date and time when the return was resolved by the Receiving Financial Depository Institution (RDFI) in ISO 8601 format without timestamp | 
 **returnCode** | **string** | The associated ACH return code and notice of change code. See [Return Codes](/api-reference/platform-api/reference/ach-return-fields#return-codes) for a complete list. | 
 **returnStatus** | **string** | The status of the return. See [Return Statuses](/api-reference/platform-api/reference/ach-return-fields#return-status) for a complete list. | 
 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;100&#x60;. If the value exceeds &#x60;100&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

### Return type

[**ACHReturnsResponseBody**](ACHReturnsResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadACHRetrun

> ACHReturnResponseBody ReadACHRetrun(ctx, achReturnGuid).Execute()

Read ACH Return



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
    achReturnGuid := "achReturnGuid_example" // string | The unique identifier (`guid`) for the ACH return. Defined by MX.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AchReturnAPI.ReadACHRetrun(context.Background(), achReturnGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AchReturnAPI.ReadACHRetrun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadACHRetrun`: ACHReturnResponseBody
    fmt.Fprintf(os.Stdout, "Response from `AchReturnAPI.ReadACHRetrun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**achReturnGuid** | **string** | The unique identifier (&#x60;guid&#x60;) for the ACH return. Defined by MX. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadACHRetrunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ACHReturnResponseBody**](ACHReturnResponseBody.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.mx.api.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

