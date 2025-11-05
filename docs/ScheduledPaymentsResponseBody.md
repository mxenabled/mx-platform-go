# ScheduledPaymentsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledPayments** | Pointer to [**[]ScheduledPaymentResponse**](ScheduledPaymentResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewScheduledPaymentsResponseBody

`func NewScheduledPaymentsResponseBody() *ScheduledPaymentsResponseBody`

NewScheduledPaymentsResponseBody instantiates a new ScheduledPaymentsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledPaymentsResponseBodyWithDefaults

`func NewScheduledPaymentsResponseBodyWithDefaults() *ScheduledPaymentsResponseBody`

NewScheduledPaymentsResponseBodyWithDefaults instantiates a new ScheduledPaymentsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledPayments

`func (o *ScheduledPaymentsResponseBody) GetScheduledPayments() []ScheduledPaymentResponse`

GetScheduledPayments returns the ScheduledPayments field if non-nil, zero value otherwise.

### GetScheduledPaymentsOk

`func (o *ScheduledPaymentsResponseBody) GetScheduledPaymentsOk() (*[]ScheduledPaymentResponse, bool)`

GetScheduledPaymentsOk returns a tuple with the ScheduledPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPayments

`func (o *ScheduledPaymentsResponseBody) SetScheduledPayments(v []ScheduledPaymentResponse)`

SetScheduledPayments sets ScheduledPayments field to given value.

### HasScheduledPayments

`func (o *ScheduledPaymentsResponseBody) HasScheduledPayments() bool`

HasScheduledPayments returns a boolean if a field has been set.

### GetPagination

`func (o *ScheduledPaymentsResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ScheduledPaymentsResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ScheduledPaymentsResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ScheduledPaymentsResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


