# HoldingsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holdings** | Pointer to [**[]HoldingResponse**](HoldingResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewHoldingsResponseBody

`func NewHoldingsResponseBody() *HoldingsResponseBody`

NewHoldingsResponseBody instantiates a new HoldingsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldingsResponseBodyWithDefaults

`func NewHoldingsResponseBodyWithDefaults() *HoldingsResponseBody`

NewHoldingsResponseBodyWithDefaults instantiates a new HoldingsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoldings

`func (o *HoldingsResponseBody) GetHoldings() []HoldingResponse`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *HoldingsResponseBody) GetHoldingsOk() (*[]HoldingResponse, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *HoldingsResponseBody) SetHoldings(v []HoldingResponse)`

SetHoldings sets Holdings field to given value.

### HasHoldings

`func (o *HoldingsResponseBody) HasHoldings() bool`

HasHoldings returns a boolean if a field has been set.

### GetPagination

`func (o *HoldingsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *HoldingsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *HoldingsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *HoldingsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


