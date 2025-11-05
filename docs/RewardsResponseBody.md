# RewardsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rewards** | Pointer to [**[]RewardsResponseBodyRewardsInner**](RewardsResponseBodyRewardsInner.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewRewardsResponseBody

`func NewRewardsResponseBody() *RewardsResponseBody`

NewRewardsResponseBody instantiates a new RewardsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardsResponseBodyWithDefaults

`func NewRewardsResponseBodyWithDefaults() *RewardsResponseBody`

NewRewardsResponseBodyWithDefaults instantiates a new RewardsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRewards

`func (o *RewardsResponseBody) GetRewards() []RewardsResponseBodyRewardsInner`

GetRewards returns the Rewards field if non-nil, zero value otherwise.

### GetRewardsOk

`func (o *RewardsResponseBody) GetRewardsOk() (*[]RewardsResponseBodyRewardsInner, bool)`

GetRewardsOk returns a tuple with the Rewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewards

`func (o *RewardsResponseBody) SetRewards(v []RewardsResponseBodyRewardsInner)`

SetRewards sets Rewards field to given value.

### HasRewards

`func (o *RewardsResponseBody) HasRewards() bool`

HasRewards returns a boolean if a field has been set.

### GetPagination

`func (o *RewardsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RewardsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RewardsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RewardsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


