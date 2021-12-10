# MerchantsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchants** | Pointer to [**[]MerchantResponse**](MerchantResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewMerchantsResponseBody

`func NewMerchantsResponseBody() *MerchantsResponseBody`

NewMerchantsResponseBody instantiates a new MerchantsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantsResponseBodyWithDefaults

`func NewMerchantsResponseBodyWithDefaults() *MerchantsResponseBody`

NewMerchantsResponseBodyWithDefaults instantiates a new MerchantsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchants

`func (o *MerchantsResponseBody) GetMerchants() []MerchantResponse`

GetMerchants returns the Merchants field if non-nil, zero value otherwise.

### GetMerchantsOk

`func (o *MerchantsResponseBody) GetMerchantsOk() (*[]MerchantResponse, bool)`

GetMerchantsOk returns a tuple with the Merchants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchants

`func (o *MerchantsResponseBody) SetMerchants(v []MerchantResponse)`

SetMerchants sets Merchants field to given value.

### HasMerchants

`func (o *MerchantsResponseBody) HasMerchants() bool`

HasMerchants returns a boolean if a field has been set.

### GetPagination

`func (o *MerchantsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *MerchantsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *MerchantsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *MerchantsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


