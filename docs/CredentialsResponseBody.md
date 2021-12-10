# CredentialsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**[]CredentialResponse**](CredentialResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewCredentialsResponseBody

`func NewCredentialsResponseBody() *CredentialsResponseBody`

NewCredentialsResponseBody instantiates a new CredentialsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsResponseBodyWithDefaults

`func NewCredentialsResponseBodyWithDefaults() *CredentialsResponseBody`

NewCredentialsResponseBodyWithDefaults instantiates a new CredentialsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *CredentialsResponseBody) GetCredentials() []CredentialResponse`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CredentialsResponseBody) GetCredentialsOk() (*[]CredentialResponse, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CredentialsResponseBody) SetCredentials(v []CredentialResponse)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CredentialsResponseBody) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetPagination

`func (o *CredentialsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CredentialsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CredentialsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *CredentialsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


