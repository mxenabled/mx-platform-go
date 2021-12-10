# UserCreateRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UserCreateRequest**](UserCreateRequest.md) |  | [optional] 

## Methods

### NewUserCreateRequestBody

`func NewUserCreateRequestBody() *UserCreateRequestBody`

NewUserCreateRequestBody instantiates a new UserCreateRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateRequestBodyWithDefaults

`func NewUserCreateRequestBodyWithDefaults() *UserCreateRequestBody`

NewUserCreateRequestBodyWithDefaults instantiates a new UserCreateRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserCreateRequestBody) GetUser() UserCreateRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserCreateRequestBody) GetUserOk() (*UserCreateRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserCreateRequestBody) SetUser(v UserCreateRequest)`

SetUser sets User field to given value.

### HasUser

`func (o *UserCreateRequestBody) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


