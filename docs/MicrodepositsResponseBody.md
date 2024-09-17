# MicrodepositsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MicroDeposits** | Pointer to [**[]MicrodepositResponse**](MicrodepositResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewMicrodepositsResponseBody

`func NewMicrodepositsResponseBody() *MicrodepositsResponseBody`

NewMicrodepositsResponseBody instantiates a new MicrodepositsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrodepositsResponseBodyWithDefaults

`func NewMicrodepositsResponseBodyWithDefaults() *MicrodepositsResponseBody`

NewMicrodepositsResponseBodyWithDefaults instantiates a new MicrodepositsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMicroDeposits

`func (o *MicrodepositsResponseBody) GetMicroDeposits() []MicrodepositResponse`

GetMicroDeposits returns the MicroDeposits field if non-nil, zero value otherwise.

### GetMicroDepositsOk

`func (o *MicrodepositsResponseBody) GetMicroDepositsOk() (*[]MicrodepositResponse, bool)`

GetMicroDepositsOk returns a tuple with the MicroDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroDeposits

`func (o *MicrodepositsResponseBody) SetMicroDeposits(v []MicrodepositResponse)`

SetMicroDeposits sets MicroDeposits field to given value.

### HasMicroDeposits

`func (o *MicrodepositsResponseBody) HasMicroDeposits() bool`

HasMicroDeposits returns a boolean if a field has been set.

### GetPagination

`func (o *MicrodepositsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *MicrodepositsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *MicrodepositsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *MicrodepositsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


