# \InstitutionsAPI

All URIs are relative to *https://int-api.mx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFavoriteInstitutions**](InstitutionsAPI.md#ListFavoriteInstitutions) | **Get** /institutions/favorites | List favorite institutions
[**ListInstitutionCredentials**](InstitutionsAPI.md#ListInstitutionCredentials) | **Get** /institutions/{institution_code}/credentials | List institution credentials
[**ListInstitutions**](InstitutionsAPI.md#ListInstitutions) | **Get** /institutions | List institutions
[**ReadInstitution**](InstitutionsAPI.md#ReadInstitution) | **Get** /institutions/{institution_code} | Read institution



## ListFavoriteInstitutions

> InstitutionsResponseBody ListFavoriteInstitutions(ctx).IsoCountryCode(isoCountryCode).Page(page).RecordsPerPage(recordsPerPage).Execute()

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
    isoCountryCode := []string{"Inner_example"} // []string | An array of strings that filters institutions in the widget by the specified country code. Acceptable codes include `US`, `CA`, and `MX` (Mexico). (optional)
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstitutionsAPI.ListFavoriteInstitutions(context.Background()).IsoCountryCode(isoCountryCode).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.ListFavoriteInstitutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFavoriteInstitutions`: InstitutionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.ListFavoriteInstitutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFavoriteInstitutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isoCountryCode** | **[]string** | An array of strings that filters institutions in the widget by the specified country code. Acceptable codes include &#x60;US&#x60;, &#x60;CA&#x60;, and &#x60;MX&#x60; (Mexico). | 
 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;100&#x60;. If the value exceeds &#x60;100&#x60;, the default value of &#x60;25&#x60; will be used instead. | 

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
    institutionCode := "mxbank" // string | The institution_code of the institution.
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `1000`. If the value exceeds `1000`, the default value of `25` will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstitutionsAPI.ListInstitutionCredentials(context.Background(), institutionCode).Page(page).RecordsPerPage(recordsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.ListInstitutionCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstitutionCredentials`: CredentialsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.ListInstitutionCredentials`: %v\n", resp)
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


## ListInstitutions

> InstitutionsResponseBody ListInstitutions(ctx).Name(name).IsoCountryCode(isoCountryCode).Page(page).RecordsPerPage(recordsPerPage).SupportsAccountIdentification(supportsAccountIdentification).SupportsAccountStatement(supportsAccountStatement).SupportsAccountVerification(supportsAccountVerification).SupportsTransactionHistory(supportsTransactionHistory).Execute()

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
    name := "mxbank" // string | This will list only institutions in which the appended string appears. (optional)
    isoCountryCode := []string{"Inner_example"} // []string | An array of strings that filters institutions in the widget by the specified country code. Acceptable codes include `US`, `CA`, and `MX` (Mexico). (optional)
    page := int32(1) // int32 | Results are paginated. Specify current page. (optional)
    recordsPerPage := int32(10) // int32 | This specifies the number of records to be returned on each page. Defaults to `25`. The valid range is from `10` to `100`. If the value exceeds `100`, the default value of `25` will be used instead. (optional)
    supportsAccountIdentification := true // bool | Filter only institutions which support account identification. (optional)
    supportsAccountStatement := true // bool | Filter only institutions which support account statements. (optional)
    supportsAccountVerification := true // bool | Filter only institutions which support account verification. (optional)
    supportsTransactionHistory := true // bool | Filter only institutions which support extended transaction history. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstitutionsAPI.ListInstitutions(context.Background()).Name(name).IsoCountryCode(isoCountryCode).Page(page).RecordsPerPage(recordsPerPage).SupportsAccountIdentification(supportsAccountIdentification).SupportsAccountStatement(supportsAccountStatement).SupportsAccountVerification(supportsAccountVerification).SupportsTransactionHistory(supportsTransactionHistory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.ListInstitutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstitutions`: InstitutionsResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.ListInstitutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstitutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | This will list only institutions in which the appended string appears. | 
 **isoCountryCode** | **[]string** | An array of strings that filters institutions in the widget by the specified country code. Acceptable codes include &#x60;US&#x60;, &#x60;CA&#x60;, and &#x60;MX&#x60; (Mexico). | 
 **page** | **int32** | Results are paginated. Specify current page. | 
 **recordsPerPage** | **int32** | This specifies the number of records to be returned on each page. Defaults to &#x60;25&#x60;. The valid range is from &#x60;10&#x60; to &#x60;100&#x60;. If the value exceeds &#x60;100&#x60;, the default value of &#x60;25&#x60; will be used instead. | 
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
    institutionCode := "mxbank" // string | The institution_code of the institution.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstitutionsAPI.ReadInstitution(context.Background(), institutionCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.ReadInstitution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadInstitution`: InstitutionResponseBody
    fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.ReadInstitution`: %v\n", resp)
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

