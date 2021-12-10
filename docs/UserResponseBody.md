# UserResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 

## Methods

### NewUserResponseBody

`func NewUserResponseBody() *UserResponseBody`

NewUserResponseBody instantiates a new UserResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseBodyWithDefaults

`func NewUserResponseBodyWithDefaults() *UserResponseBody`

NewUserResponseBodyWithDefaults instantiates a new UserResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserResponseBody) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserResponseBody) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserResponseBody) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *UserResponseBody) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


