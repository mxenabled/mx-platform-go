# ACHReturnsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchReturns** | Pointer to [**[]ACHResponse**](ACHResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewACHReturnsResponseBody

`func NewACHReturnsResponseBody() *ACHReturnsResponseBody`

NewACHReturnsResponseBody instantiates a new ACHReturnsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACHReturnsResponseBodyWithDefaults

`func NewACHReturnsResponseBodyWithDefaults() *ACHReturnsResponseBody`

NewACHReturnsResponseBodyWithDefaults instantiates a new ACHReturnsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchReturns

`func (o *ACHReturnsResponseBody) GetAchReturns() []ACHResponse`

GetAchReturns returns the AchReturns field if non-nil, zero value otherwise.

### GetAchReturnsOk

`func (o *ACHReturnsResponseBody) GetAchReturnsOk() (*[]ACHResponse, bool)`

GetAchReturnsOk returns a tuple with the AchReturns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchReturns

`func (o *ACHReturnsResponseBody) SetAchReturns(v []ACHResponse)`

SetAchReturns sets AchReturns field to given value.

### HasAchReturns

`func (o *ACHReturnsResponseBody) HasAchReturns() bool`

HasAchReturns returns a boolean if a field has been set.

### GetPagination

`func (o *ACHReturnsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ACHReturnsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ACHReturnsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ACHReturnsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


