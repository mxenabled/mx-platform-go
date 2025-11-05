# ProcessorOwnerBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountOwners** | Pointer to [**[]ProcessorOwnerBodyAccountOwnersInner**](ProcessorOwnerBodyAccountOwnersInner.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewProcessorOwnerBody

`func NewProcessorOwnerBody() *ProcessorOwnerBody`

NewProcessorOwnerBody instantiates a new ProcessorOwnerBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorOwnerBodyWithDefaults

`func NewProcessorOwnerBodyWithDefaults() *ProcessorOwnerBody`

NewProcessorOwnerBodyWithDefaults instantiates a new ProcessorOwnerBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountOwners

`func (o *ProcessorOwnerBody) GetAccountOwners() []ProcessorOwnerBodyAccountOwnersInner`

GetAccountOwners returns the AccountOwners field if non-nil, zero value otherwise.

### GetAccountOwnersOk

`func (o *ProcessorOwnerBody) GetAccountOwnersOk() (*[]ProcessorOwnerBodyAccountOwnersInner, bool)`

GetAccountOwnersOk returns a tuple with the AccountOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwners

`func (o *ProcessorOwnerBody) SetAccountOwners(v []ProcessorOwnerBodyAccountOwnersInner)`

SetAccountOwners sets AccountOwners field to given value.

### HasAccountOwners

`func (o *ProcessorOwnerBody) HasAccountOwners() bool`

HasAccountOwners returns a boolean if a field has been set.

### GetPagination

`func (o *ProcessorOwnerBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ProcessorOwnerBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ProcessorOwnerBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ProcessorOwnerBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


