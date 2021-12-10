# UserUpdateRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UserUpdateRequest**](UserUpdateRequest.md) |  | [optional] 

## Methods

### NewUserUpdateRequestBody

`func NewUserUpdateRequestBody() *UserUpdateRequestBody`

NewUserUpdateRequestBody instantiates a new UserUpdateRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateRequestBodyWithDefaults

`func NewUserUpdateRequestBodyWithDefaults() *UserUpdateRequestBody`

NewUserUpdateRequestBodyWithDefaults instantiates a new UserUpdateRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserUpdateRequestBody) GetUser() UserUpdateRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserUpdateRequestBody) GetUserOk() (*UserUpdateRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserUpdateRequestBody) SetUser(v UserUpdateRequest)`

SetUser sets User field to given value.

### HasUser

`func (o *UserUpdateRequestBody) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


