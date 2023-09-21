# SpendingPlanIterationItemsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IterationItems** | Pointer to [**[]SpendingPlanIterationItemResponse**](SpendingPlanIterationItemResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewSpendingPlanIterationItemsResponseBody

`func NewSpendingPlanIterationItemsResponseBody() *SpendingPlanIterationItemsResponseBody`

NewSpendingPlanIterationItemsResponseBody instantiates a new SpendingPlanIterationItemsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendingPlanIterationItemsResponseBodyWithDefaults

`func NewSpendingPlanIterationItemsResponseBodyWithDefaults() *SpendingPlanIterationItemsResponseBody`

NewSpendingPlanIterationItemsResponseBodyWithDefaults instantiates a new SpendingPlanIterationItemsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIterationItems

`func (o *SpendingPlanIterationItemsResponseBody) GetIterationItems() []SpendingPlanIterationItemResponse`

GetIterationItems returns the IterationItems field if non-nil, zero value otherwise.

### GetIterationItemsOk

`func (o *SpendingPlanIterationItemsResponseBody) GetIterationItemsOk() (*[]SpendingPlanIterationItemResponse, bool)`

GetIterationItemsOk returns a tuple with the IterationItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationItems

`func (o *SpendingPlanIterationItemsResponseBody) SetIterationItems(v []SpendingPlanIterationItemResponse)`

SetIterationItems sets IterationItems field to given value.

### HasIterationItems

`func (o *SpendingPlanIterationItemsResponseBody) HasIterationItems() bool`

HasIterationItems returns a boolean if a field has been set.

### GetPagination

`func (o *SpendingPlanIterationItemsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SpendingPlanIterationItemsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SpendingPlanIterationItemsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *SpendingPlanIterationItemsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


