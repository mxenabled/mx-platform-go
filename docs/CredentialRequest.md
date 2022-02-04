# CredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewCredentialRequest

`func NewCredentialRequest() *CredentialRequest`

NewCredentialRequest instantiates a new CredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialRequestWithDefaults

`func NewCredentialRequestWithDefaults() *CredentialRequest`

NewCredentialRequestWithDefaults instantiates a new CredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *CredentialRequest) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *CredentialRequest) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *CredentialRequest) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *CredentialRequest) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetValue

`func (o *CredentialRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CredentialRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CredentialRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CredentialRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


