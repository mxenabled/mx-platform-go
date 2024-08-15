# RepositionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | **string** | The unique identifier for the goal. Defined by MX. | 
**Position** | **int32** | The priority of the goal in relation to multiple goals. | 

## Methods

### NewRepositionRequest

`func NewRepositionRequest(guid string, position int32, ) *RepositionRequest`

NewRepositionRequest instantiates a new RepositionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositionRequestWithDefaults

`func NewRepositionRequestWithDefaults() *RepositionRequest`

NewRepositionRequestWithDefaults instantiates a new RepositionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *RepositionRequest) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *RepositionRequest) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *RepositionRequest) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetPosition

`func (o *RepositionRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *RepositionRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *RepositionRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


