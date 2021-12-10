# TaggingCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagGuid** | **string** |  | 
**TransactionGuid** | **string** |  | 

## Methods

### NewTaggingCreateRequest

`func NewTaggingCreateRequest(tagGuid string, transactionGuid string, ) *TaggingCreateRequest`

NewTaggingCreateRequest instantiates a new TaggingCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaggingCreateRequestWithDefaults

`func NewTaggingCreateRequestWithDefaults() *TaggingCreateRequest`

NewTaggingCreateRequestWithDefaults instantiates a new TaggingCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagGuid

`func (o *TaggingCreateRequest) GetTagGuid() string`

GetTagGuid returns the TagGuid field if non-nil, zero value otherwise.

### GetTagGuidOk

`func (o *TaggingCreateRequest) GetTagGuidOk() (*string, bool)`

GetTagGuidOk returns a tuple with the TagGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagGuid

`func (o *TaggingCreateRequest) SetTagGuid(v string)`

SetTagGuid sets TagGuid field to given value.


### GetTransactionGuid

`func (o *TaggingCreateRequest) GetTransactionGuid() string`

GetTransactionGuid returns the TransactionGuid field if non-nil, zero value otherwise.

### GetTransactionGuidOk

`func (o *TaggingCreateRequest) GetTransactionGuidOk() (*string, bool)`

GetTransactionGuidOk returns a tuple with the TransactionGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionGuid

`func (o *TaggingCreateRequest) SetTransactionGuid(v string)`

SetTransactionGuid sets TransactionGuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


