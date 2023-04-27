# TaxDocumentsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 
**TaxDocuments** | Pointer to [**[]TaxDocumentResponse**](TaxDocumentResponse.md) |  | [optional] 

## Methods

### NewTaxDocumentsResponseBody

`func NewTaxDocumentsResponseBody() *TaxDocumentsResponseBody`

NewTaxDocumentsResponseBody instantiates a new TaxDocumentsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxDocumentsResponseBodyWithDefaults

`func NewTaxDocumentsResponseBodyWithDefaults() *TaxDocumentsResponseBody`

NewTaxDocumentsResponseBodyWithDefaults instantiates a new TaxDocumentsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *TaxDocumentsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TaxDocumentsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TaxDocumentsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TaxDocumentsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetTaxDocuments

`func (o *TaxDocumentsResponseBody) GetTaxDocuments() []TaxDocumentResponse`

GetTaxDocuments returns the TaxDocuments field if non-nil, zero value otherwise.

### GetTaxDocumentsOk

`func (o *TaxDocumentsResponseBody) GetTaxDocumentsOk() (*[]TaxDocumentResponse, bool)`

GetTaxDocumentsOk returns a tuple with the TaxDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDocuments

`func (o *TaxDocumentsResponseBody) SetTaxDocuments(v []TaxDocumentResponse)`

SetTaxDocuments sets TaxDocuments field to given value.

### HasTaxDocuments

`func (o *TaxDocumentsResponseBody) HasTaxDocuments() bool`

HasTaxDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


