# InvestmentHoldingsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestmentHoldings** | Pointer to [**[]InvestmentHoldingResponse**](InvestmentHoldingResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewInvestmentHoldingsResponseBody

`func NewInvestmentHoldingsResponseBody() *InvestmentHoldingsResponseBody`

NewInvestmentHoldingsResponseBody instantiates a new InvestmentHoldingsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentHoldingsResponseBodyWithDefaults

`func NewInvestmentHoldingsResponseBodyWithDefaults() *InvestmentHoldingsResponseBody`

NewInvestmentHoldingsResponseBodyWithDefaults instantiates a new InvestmentHoldingsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestmentHoldings

`func (o *InvestmentHoldingsResponseBody) GetInvestmentHoldings() []InvestmentHoldingResponse`

GetInvestmentHoldings returns the InvestmentHoldings field if non-nil, zero value otherwise.

### GetInvestmentHoldingsOk

`func (o *InvestmentHoldingsResponseBody) GetInvestmentHoldingsOk() (*[]InvestmentHoldingResponse, bool)`

GetInvestmentHoldingsOk returns a tuple with the InvestmentHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentHoldings

`func (o *InvestmentHoldingsResponseBody) SetInvestmentHoldings(v []InvestmentHoldingResponse)`

SetInvestmentHoldings sets InvestmentHoldings field to given value.

### HasInvestmentHoldings

`func (o *InvestmentHoldingsResponseBody) HasInvestmentHoldings() bool`

HasInvestmentHoldings returns a boolean if a field has been set.

### GetPagination

`func (o *InvestmentHoldingsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InvestmentHoldingsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InvestmentHoldingsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *InvestmentHoldingsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


