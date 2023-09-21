# SpendingPlanIterationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iterations** | Pointer to [**[]SpendingPlanIterationResponse**](SpendingPlanIterationResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewSpendingPlanIterationsResponse

`func NewSpendingPlanIterationsResponse() *SpendingPlanIterationsResponse`

NewSpendingPlanIterationsResponse instantiates a new SpendingPlanIterationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendingPlanIterationsResponseWithDefaults

`func NewSpendingPlanIterationsResponseWithDefaults() *SpendingPlanIterationsResponse`

NewSpendingPlanIterationsResponseWithDefaults instantiates a new SpendingPlanIterationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIterations

`func (o *SpendingPlanIterationsResponse) GetIterations() []SpendingPlanIterationResponse`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *SpendingPlanIterationsResponse) GetIterationsOk() (*[]SpendingPlanIterationResponse, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *SpendingPlanIterationsResponse) SetIterations(v []SpendingPlanIterationResponse)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *SpendingPlanIterationsResponse) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### GetPagination

`func (o *SpendingPlanIterationsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SpendingPlanIterationsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SpendingPlanIterationsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *SpendingPlanIterationsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


