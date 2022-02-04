# ChallengeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **NullableString** |  | [optional] 
**Guid** | Pointer to **NullableString** |  | [optional] 
**ImageData** | Pointer to **NullableString** |  | [optional] 
**ImageOptions** | Pointer to [**[]ImageOptionResponse**](ImageOptionResponse.md) |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to [**[]OptionResponse**](OptionResponse.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChallengeResponse

`func NewChallengeResponse() *ChallengeResponse`

NewChallengeResponse instantiates a new ChallengeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeResponseWithDefaults

`func NewChallengeResponseWithDefaults() *ChallengeResponse`

NewChallengeResponseWithDefaults instantiates a new ChallengeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *ChallengeResponse) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ChallengeResponse) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ChallengeResponse) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ChallengeResponse) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### SetFieldNameNil

`func (o *ChallengeResponse) SetFieldNameNil(b bool)`

 SetFieldNameNil sets the value for FieldName to be an explicit nil

### UnsetFieldName
`func (o *ChallengeResponse) UnsetFieldName()`

UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
### GetGuid

`func (o *ChallengeResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ChallengeResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ChallengeResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ChallengeResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *ChallengeResponse) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *ChallengeResponse) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetImageData

`func (o *ChallengeResponse) GetImageData() string`

GetImageData returns the ImageData field if non-nil, zero value otherwise.

### GetImageDataOk

`func (o *ChallengeResponse) GetImageDataOk() (*string, bool)`

GetImageDataOk returns a tuple with the ImageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageData

`func (o *ChallengeResponse) SetImageData(v string)`

SetImageData sets ImageData field to given value.

### HasImageData

`func (o *ChallengeResponse) HasImageData() bool`

HasImageData returns a boolean if a field has been set.

### SetImageDataNil

`func (o *ChallengeResponse) SetImageDataNil(b bool)`

 SetImageDataNil sets the value for ImageData to be an explicit nil

### UnsetImageData
`func (o *ChallengeResponse) UnsetImageData()`

UnsetImageData ensures that no value is present for ImageData, not even an explicit nil
### GetImageOptions

`func (o *ChallengeResponse) GetImageOptions() []ImageOptionResponse`

GetImageOptions returns the ImageOptions field if non-nil, zero value otherwise.

### GetImageOptionsOk

`func (o *ChallengeResponse) GetImageOptionsOk() (*[]ImageOptionResponse, bool)`

GetImageOptionsOk returns a tuple with the ImageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOptions

`func (o *ChallengeResponse) SetImageOptions(v []ImageOptionResponse)`

SetImageOptions sets ImageOptions field to given value.

### HasImageOptions

`func (o *ChallengeResponse) HasImageOptions() bool`

HasImageOptions returns a boolean if a field has been set.

### GetLabel

`func (o *ChallengeResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ChallengeResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ChallengeResponse) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ChallengeResponse) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ChallengeResponse) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ChallengeResponse) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOptions

`func (o *ChallengeResponse) GetOptions() []OptionResponse`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ChallengeResponse) GetOptionsOk() (*[]OptionResponse, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ChallengeResponse) SetOptions(v []OptionResponse)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ChallengeResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetType

`func (o *ChallengeResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChallengeResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChallengeResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChallengeResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ChallengeResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ChallengeResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


