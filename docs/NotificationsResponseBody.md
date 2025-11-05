# NotificationsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | Pointer to [**[]NotificationResponse**](NotificationResponse.md) |  | [optional] 

## Methods

### NewNotificationsResponseBody

`func NewNotificationsResponseBody() *NotificationsResponseBody`

NewNotificationsResponseBody instantiates a new NotificationsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsResponseBodyWithDefaults

`func NewNotificationsResponseBodyWithDefaults() *NotificationsResponseBody`

NewNotificationsResponseBodyWithDefaults instantiates a new NotificationsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *NotificationsResponseBody) GetNotifications() []NotificationResponse`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *NotificationsResponseBody) GetNotificationsOk() (*[]NotificationResponse, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *NotificationsResponseBody) SetNotifications(v []NotificationResponse)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *NotificationsResponseBody) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


