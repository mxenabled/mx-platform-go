# ChallengesResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Challenges** | Pointer to [**[]ChallengeResponse**](ChallengeResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewChallengesResponseBody

`func NewChallengesResponseBody() *ChallengesResponseBody`

NewChallengesResponseBody instantiates a new ChallengesResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengesResponseBodyWithDefaults

`func NewChallengesResponseBodyWithDefaults() *ChallengesResponseBody`

NewChallengesResponseBodyWithDefaults instantiates a new ChallengesResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenges

`func (o *ChallengesResponseBody) GetChallenges() []ChallengeResponse`

GetChallenges returns the Challenges field if non-nil, zero value otherwise.

### GetChallengesOk

`func (o *ChallengesResponseBody) GetChallengesOk() (*[]ChallengeResponse, bool)`

GetChallengesOk returns a tuple with the Challenges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenges

`func (o *ChallengesResponseBody) SetChallenges(v []ChallengeResponse)`

SetChallenges sets Challenges field to given value.

### HasChallenges

`func (o *ChallengesResponseBody) HasChallenges() bool`

HasChallenges returns a boolean if a field has been set.

### GetPagination

`func (o *ChallengesResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ChallengesResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ChallengesResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ChallengesResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


