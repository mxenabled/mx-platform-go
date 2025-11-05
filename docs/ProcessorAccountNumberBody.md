# ProcessorAccountNumberBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumbers** | Pointer to [**[]ProcessorAccountNumberBodyAccountNumbersInner**](ProcessorAccountNumberBodyAccountNumbersInner.md) |  | [optional] 

## Methods

### NewProcessorAccountNumberBody

`func NewProcessorAccountNumberBody() *ProcessorAccountNumberBody`

NewProcessorAccountNumberBody instantiates a new ProcessorAccountNumberBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorAccountNumberBodyWithDefaults

`func NewProcessorAccountNumberBodyWithDefaults() *ProcessorAccountNumberBody`

NewProcessorAccountNumberBodyWithDefaults instantiates a new ProcessorAccountNumberBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumbers

`func (o *ProcessorAccountNumberBody) GetAccountNumbers() []ProcessorAccountNumberBodyAccountNumbersInner`

GetAccountNumbers returns the AccountNumbers field if non-nil, zero value otherwise.

### GetAccountNumbersOk

`func (o *ProcessorAccountNumberBody) GetAccountNumbersOk() (*[]ProcessorAccountNumberBodyAccountNumbersInner, bool)`

GetAccountNumbersOk returns a tuple with the AccountNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumbers

`func (o *ProcessorAccountNumberBody) SetAccountNumbers(v []ProcessorAccountNumberBodyAccountNumbersInner)`

SetAccountNumbers sets AccountNumbers field to given value.

### HasAccountNumbers

`func (o *ProcessorAccountNumberBody) HasAccountNumbers() bool`

HasAccountNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


