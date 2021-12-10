# InstitutionsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Institutions** | Pointer to [**[]InstitutionResponse**](InstitutionResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewInstitutionsResponseBody

`func NewInstitutionsResponseBody() *InstitutionsResponseBody`

NewInstitutionsResponseBody instantiates a new InstitutionsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsResponseBodyWithDefaults

`func NewInstitutionsResponseBodyWithDefaults() *InstitutionsResponseBody`

NewInstitutionsResponseBodyWithDefaults instantiates a new InstitutionsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutions

`func (o *InstitutionsResponseBody) GetInstitutions() []InstitutionResponse`

GetInstitutions returns the Institutions field if non-nil, zero value otherwise.

### GetInstitutionsOk

`func (o *InstitutionsResponseBody) GetInstitutionsOk() (*[]InstitutionResponse, bool)`

GetInstitutionsOk returns a tuple with the Institutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutions

`func (o *InstitutionsResponseBody) SetInstitutions(v []InstitutionResponse)`

SetInstitutions sets Institutions field to given value.

### HasInstitutions

`func (o *InstitutionsResponseBody) HasInstitutions() bool`

HasInstitutions returns a boolean if a field has been set.

### GetPagination

`func (o *InstitutionsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InstitutionsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InstitutionsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *InstitutionsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


