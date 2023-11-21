# SpendingPlansResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpendingPlans** | Pointer to [**[]SpendingPlanResponse**](SpendingPlanResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewSpendingPlansResponseBody

`func NewSpendingPlansResponseBody() *SpendingPlansResponseBody`

NewSpendingPlansResponseBody instantiates a new SpendingPlansResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendingPlansResponseBodyWithDefaults

`func NewSpendingPlansResponseBodyWithDefaults() *SpendingPlansResponseBody`

NewSpendingPlansResponseBodyWithDefaults instantiates a new SpendingPlansResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpendingPlans

`func (o *SpendingPlansResponseBody) GetSpendingPlans() []SpendingPlanResponse`

GetSpendingPlans returns the SpendingPlans field if non-nil, zero value otherwise.

### GetSpendingPlansOk

`func (o *SpendingPlansResponseBody) GetSpendingPlansOk() (*[]SpendingPlanResponse, bool)`

GetSpendingPlansOk returns a tuple with the SpendingPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendingPlans

`func (o *SpendingPlansResponseBody) SetSpendingPlans(v []SpendingPlanResponse)`

SetSpendingPlans sets SpendingPlans field to given value.

### HasSpendingPlans

`func (o *SpendingPlansResponseBody) HasSpendingPlans() bool`

HasSpendingPlans returns a boolean if a field has been set.

### GetPagination

`func (o *SpendingPlansResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SpendingPlansResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SpendingPlansResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *SpendingPlansResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


