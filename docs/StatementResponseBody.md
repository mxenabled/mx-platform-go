# StatementResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | Pointer to [**StatementResponse**](StatementResponse.md) |  | [optional] 

## Methods

### NewStatementResponseBody

`func NewStatementResponseBody() *StatementResponseBody`

NewStatementResponseBody instantiates a new StatementResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementResponseBodyWithDefaults

`func NewStatementResponseBodyWithDefaults() *StatementResponseBody`

NewStatementResponseBodyWithDefaults instantiates a new StatementResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *StatementResponseBody) GetStatement() StatementResponse`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *StatementResponseBody) GetStatementOk() (*StatementResponse, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *StatementResponseBody) SetStatement(v StatementResponse)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *StatementResponseBody) HasStatement() bool`

HasStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


