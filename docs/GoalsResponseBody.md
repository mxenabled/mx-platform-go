# GoalsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Goals** | Pointer to [**[]GoalsResponse**](GoalsResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewGoalsResponseBody

`func NewGoalsResponseBody() *GoalsResponseBody`

NewGoalsResponseBody instantiates a new GoalsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalsResponseBodyWithDefaults

`func NewGoalsResponseBodyWithDefaults() *GoalsResponseBody`

NewGoalsResponseBodyWithDefaults instantiates a new GoalsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoals

`func (o *GoalsResponseBody) GetGoals() []GoalsResponse`

GetGoals returns the Goals field if non-nil, zero value otherwise.

### GetGoalsOk

`func (o *GoalsResponseBody) GetGoalsOk() (*[]GoalsResponse, bool)`

GetGoalsOk returns a tuple with the Goals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoals

`func (o *GoalsResponseBody) SetGoals(v []GoalsResponse)`

SetGoals sets Goals field to given value.

### HasGoals

`func (o *GoalsResponseBody) HasGoals() bool`

HasGoals returns a boolean if a field has been set.

### GetPagination

`func (o *GoalsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GoalsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GoalsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GoalsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


