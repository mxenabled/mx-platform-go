# RewardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | Pointer to **string** |  | [optional] 
**BalanceType** | Pointer to **string** |  | [optional] 
**Balance** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExpiresOn** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**MemberGuid** | Pointer to **string** |  | [optional] 
**UnitType** | Pointer to **string** |  | [optional] 
**UserGuid** | Pointer to **string** |  | [optional] 

## Methods

### NewRewardResponse

`func NewRewardResponse() *RewardResponse`

NewRewardResponse instantiates a new RewardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardResponseWithDefaults

`func NewRewardResponseWithDefaults() *RewardResponse`

NewRewardResponseWithDefaults instantiates a new RewardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *RewardResponse) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *RewardResponse) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *RewardResponse) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.

### HasAccountGuid

`func (o *RewardResponse) HasAccountGuid() bool`

HasAccountGuid returns a boolean if a field has been set.

### GetBalanceType

`func (o *RewardResponse) GetBalanceType() string`

GetBalanceType returns the BalanceType field if non-nil, zero value otherwise.

### GetBalanceTypeOk

`func (o *RewardResponse) GetBalanceTypeOk() (*string, bool)`

GetBalanceTypeOk returns a tuple with the BalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceType

`func (o *RewardResponse) SetBalanceType(v string)`

SetBalanceType sets BalanceType field to given value.

### HasBalanceType

`func (o *RewardResponse) HasBalanceType() bool`

HasBalanceType returns a boolean if a field has been set.

### GetBalance

`func (o *RewardResponse) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *RewardResponse) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *RewardResponse) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *RewardResponse) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RewardResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RewardResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RewardResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RewardResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *RewardResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RewardResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RewardResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RewardResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresOn

`func (o *RewardResponse) GetExpiresOn() string`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *RewardResponse) GetExpiresOnOk() (*string, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *RewardResponse) SetExpiresOn(v string)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *RewardResponse) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### GetGuid

`func (o *RewardResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *RewardResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *RewardResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *RewardResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetMemberGuid

`func (o *RewardResponse) GetMemberGuid() string`

GetMemberGuid returns the MemberGuid field if non-nil, zero value otherwise.

### GetMemberGuidOk

`func (o *RewardResponse) GetMemberGuidOk() (*string, bool)`

GetMemberGuidOk returns a tuple with the MemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGuid

`func (o *RewardResponse) SetMemberGuid(v string)`

SetMemberGuid sets MemberGuid field to given value.

### HasMemberGuid

`func (o *RewardResponse) HasMemberGuid() bool`

HasMemberGuid returns a boolean if a field has been set.

### GetUnitType

`func (o *RewardResponse) GetUnitType() string`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *RewardResponse) GetUnitTypeOk() (*string, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *RewardResponse) SetUnitType(v string)`

SetUnitType sets UnitType field to given value.

### HasUnitType

`func (o *RewardResponse) HasUnitType() bool`

HasUnitType returns a boolean if a field has been set.

### GetUserGuid

`func (o *RewardResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *RewardResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *RewardResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *RewardResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


