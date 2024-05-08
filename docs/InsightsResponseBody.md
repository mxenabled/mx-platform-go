# InsightsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Insights** | Pointer to [**[]InsightResponse**](InsightResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewInsightsResponseBody

`func NewInsightsResponseBody() *InsightsResponseBody`

NewInsightsResponseBody instantiates a new InsightsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsResponseBodyWithDefaults

`func NewInsightsResponseBodyWithDefaults() *InsightsResponseBody`

NewInsightsResponseBodyWithDefaults instantiates a new InsightsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsights

`func (o *InsightsResponseBody) GetInsights() []InsightResponse`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *InsightsResponseBody) GetInsightsOk() (*[]InsightResponse, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsights

`func (o *InsightsResponseBody) SetInsights(v []InsightResponse)`

SetInsights sets Insights field to given value.

### HasInsights

`func (o *InsightsResponseBody) HasInsights() bool`

HasInsights returns a boolean if a field has been set.

### GetPagination

`func (o *InsightsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InsightsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InsightsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *InsightsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


