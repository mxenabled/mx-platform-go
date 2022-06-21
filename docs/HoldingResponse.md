# HoldingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | Pointer to **NullableString** |  | [optional] 
**CostBasis** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**Cusip** | Pointer to **NullableString** |  | [optional] 
**DailyChange** | Pointer to **NullableFloat32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Guid** | Pointer to **NullableString** |  | [optional] 
**HoldingType** | Pointer to **NullableString** |  | [optional] 
**HoldingTypeId** | Pointer to **NullableInt32** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**MarketValue** | Pointer to **NullableFloat32** |  | [optional] 
**MemberGuid** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**PurchasePrice** | Pointer to **NullableFloat32** |  | [optional] 
**Shares** | Pointer to **NullableFloat32** |  | [optional] 
**Symbol** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**UserGuid** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHoldingResponse

`func NewHoldingResponse() *HoldingResponse`

NewHoldingResponse instantiates a new HoldingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldingResponseWithDefaults

`func NewHoldingResponseWithDefaults() *HoldingResponse`

NewHoldingResponseWithDefaults instantiates a new HoldingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *HoldingResponse) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *HoldingResponse) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *HoldingResponse) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.

### HasAccountGuid

`func (o *HoldingResponse) HasAccountGuid() bool`

HasAccountGuid returns a boolean if a field has been set.

### SetAccountGuidNil

`func (o *HoldingResponse) SetAccountGuidNil(b bool)`

 SetAccountGuidNil sets the value for AccountGuid to be an explicit nil

### UnsetAccountGuid
`func (o *HoldingResponse) UnsetAccountGuid()`

UnsetAccountGuid ensures that no value is present for AccountGuid, not even an explicit nil
### GetCostBasis

`func (o *HoldingResponse) GetCostBasis() float32`

GetCostBasis returns the CostBasis field if non-nil, zero value otherwise.

### GetCostBasisOk

`func (o *HoldingResponse) GetCostBasisOk() (*float32, bool)`

GetCostBasisOk returns a tuple with the CostBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBasis

`func (o *HoldingResponse) SetCostBasis(v float32)`

SetCostBasis sets CostBasis field to given value.

### HasCostBasis

`func (o *HoldingResponse) HasCostBasis() bool`

HasCostBasis returns a boolean if a field has been set.

### SetCostBasisNil

`func (o *HoldingResponse) SetCostBasisNil(b bool)`

 SetCostBasisNil sets the value for CostBasis to be an explicit nil

### UnsetCostBasis
`func (o *HoldingResponse) UnsetCostBasis()`

UnsetCostBasis ensures that no value is present for CostBasis, not even an explicit nil
### GetCreatedAt

`func (o *HoldingResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HoldingResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HoldingResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HoldingResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *HoldingResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *HoldingResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCurrencyCode

`func (o *HoldingResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *HoldingResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *HoldingResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *HoldingResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *HoldingResponse) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *HoldingResponse) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetCusip

`func (o *HoldingResponse) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *HoldingResponse) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *HoldingResponse) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *HoldingResponse) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### SetCusipNil

`func (o *HoldingResponse) SetCusipNil(b bool)`

 SetCusipNil sets the value for Cusip to be an explicit nil

### UnsetCusip
`func (o *HoldingResponse) UnsetCusip()`

UnsetCusip ensures that no value is present for Cusip, not even an explicit nil
### GetDailyChange

`func (o *HoldingResponse) GetDailyChange() float32`

GetDailyChange returns the DailyChange field if non-nil, zero value otherwise.

### GetDailyChangeOk

`func (o *HoldingResponse) GetDailyChangeOk() (*float32, bool)`

GetDailyChangeOk returns a tuple with the DailyChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyChange

`func (o *HoldingResponse) SetDailyChange(v float32)`

SetDailyChange sets DailyChange field to given value.

### HasDailyChange

`func (o *HoldingResponse) HasDailyChange() bool`

HasDailyChange returns a boolean if a field has been set.

### SetDailyChangeNil

`func (o *HoldingResponse) SetDailyChangeNil(b bool)`

 SetDailyChangeNil sets the value for DailyChange to be an explicit nil

### UnsetDailyChange
`func (o *HoldingResponse) UnsetDailyChange()`

UnsetDailyChange ensures that no value is present for DailyChange, not even an explicit nil
### GetDescription

`func (o *HoldingResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HoldingResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HoldingResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HoldingResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HoldingResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HoldingResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGuid

`func (o *HoldingResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *HoldingResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *HoldingResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *HoldingResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *HoldingResponse) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *HoldingResponse) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetHoldingType

`func (o *HoldingResponse) GetHoldingType() string`

GetHoldingType returns the HoldingType field if non-nil, zero value otherwise.

### GetHoldingTypeOk

`func (o *HoldingResponse) GetHoldingTypeOk() (*string, bool)`

GetHoldingTypeOk returns a tuple with the HoldingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingType

`func (o *HoldingResponse) SetHoldingType(v string)`

SetHoldingType sets HoldingType field to given value.

### HasHoldingType

`func (o *HoldingResponse) HasHoldingType() bool`

HasHoldingType returns a boolean if a field has been set.

### SetHoldingTypeNil

`func (o *HoldingResponse) SetHoldingTypeNil(b bool)`

 SetHoldingTypeNil sets the value for HoldingType to be an explicit nil

### UnsetHoldingType
`func (o *HoldingResponse) UnsetHoldingType()`

UnsetHoldingType ensures that no value is present for HoldingType, not even an explicit nil
### GetHoldingTypeId

`func (o *HoldingResponse) GetHoldingTypeId() int32`

GetHoldingTypeId returns the HoldingTypeId field if non-nil, zero value otherwise.

### GetHoldingTypeIdOk

`func (o *HoldingResponse) GetHoldingTypeIdOk() (*int32, bool)`

GetHoldingTypeIdOk returns a tuple with the HoldingTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingTypeId

`func (o *HoldingResponse) SetHoldingTypeId(v int32)`

SetHoldingTypeId sets HoldingTypeId field to given value.

### HasHoldingTypeId

`func (o *HoldingResponse) HasHoldingTypeId() bool`

HasHoldingTypeId returns a boolean if a field has been set.

### SetHoldingTypeIdNil

`func (o *HoldingResponse) SetHoldingTypeIdNil(b bool)`

 SetHoldingTypeIdNil sets the value for HoldingTypeId to be an explicit nil

### UnsetHoldingTypeId
`func (o *HoldingResponse) UnsetHoldingTypeId()`

UnsetHoldingTypeId ensures that no value is present for HoldingTypeId, not even an explicit nil
### GetId

`func (o *HoldingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HoldingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HoldingResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HoldingResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HoldingResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HoldingResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMarketValue

`func (o *HoldingResponse) GetMarketValue() float32`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *HoldingResponse) GetMarketValueOk() (*float32, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *HoldingResponse) SetMarketValue(v float32)`

SetMarketValue sets MarketValue field to given value.

### HasMarketValue

`func (o *HoldingResponse) HasMarketValue() bool`

HasMarketValue returns a boolean if a field has been set.

### SetMarketValueNil

`func (o *HoldingResponse) SetMarketValueNil(b bool)`

 SetMarketValueNil sets the value for MarketValue to be an explicit nil

### UnsetMarketValue
`func (o *HoldingResponse) UnsetMarketValue()`

UnsetMarketValue ensures that no value is present for MarketValue, not even an explicit nil
### GetMemberGuid

`func (o *HoldingResponse) GetMemberGuid() string`

GetMemberGuid returns the MemberGuid field if non-nil, zero value otherwise.

### GetMemberGuidOk

`func (o *HoldingResponse) GetMemberGuidOk() (*string, bool)`

GetMemberGuidOk returns a tuple with the MemberGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGuid

`func (o *HoldingResponse) SetMemberGuid(v string)`

SetMemberGuid sets MemberGuid field to given value.

### HasMemberGuid

`func (o *HoldingResponse) HasMemberGuid() bool`

HasMemberGuid returns a boolean if a field has been set.

### SetMemberGuidNil

`func (o *HoldingResponse) SetMemberGuidNil(b bool)`

 SetMemberGuidNil sets the value for MemberGuid to be an explicit nil

### UnsetMemberGuid
`func (o *HoldingResponse) UnsetMemberGuid()`

UnsetMemberGuid ensures that no value is present for MemberGuid, not even an explicit nil
### GetMetadata

`func (o *HoldingResponse) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HoldingResponse) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HoldingResponse) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HoldingResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *HoldingResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *HoldingResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPurchasePrice

`func (o *HoldingResponse) GetPurchasePrice() float32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *HoldingResponse) GetPurchasePriceOk() (*float32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *HoldingResponse) SetPurchasePrice(v float32)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *HoldingResponse) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### SetPurchasePriceNil

`func (o *HoldingResponse) SetPurchasePriceNil(b bool)`

 SetPurchasePriceNil sets the value for PurchasePrice to be an explicit nil

### UnsetPurchasePrice
`func (o *HoldingResponse) UnsetPurchasePrice()`

UnsetPurchasePrice ensures that no value is present for PurchasePrice, not even an explicit nil
### GetShares

`func (o *HoldingResponse) GetShares() float32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *HoldingResponse) GetSharesOk() (*float32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *HoldingResponse) SetShares(v float32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *HoldingResponse) HasShares() bool`

HasShares returns a boolean if a field has been set.

### SetSharesNil

`func (o *HoldingResponse) SetSharesNil(b bool)`

 SetSharesNil sets the value for Shares to be an explicit nil

### UnsetShares
`func (o *HoldingResponse) UnsetShares()`

UnsetShares ensures that no value is present for Shares, not even an explicit nil
### GetSymbol

`func (o *HoldingResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *HoldingResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *HoldingResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *HoldingResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### SetSymbolNil

`func (o *HoldingResponse) SetSymbolNil(b bool)`

 SetSymbolNil sets the value for Symbol to be an explicit nil

### UnsetSymbol
`func (o *HoldingResponse) UnsetSymbol()`

UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil
### GetUpdatedAt

`func (o *HoldingResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HoldingResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HoldingResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HoldingResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *HoldingResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *HoldingResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserGuid

`func (o *HoldingResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *HoldingResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *HoldingResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *HoldingResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### SetUserGuidNil

`func (o *HoldingResponse) SetUserGuidNil(b bool)`

 SetUserGuidNil sets the value for UserGuid to be an explicit nil

### UnsetUserGuid
`func (o *HoldingResponse) UnsetUserGuid()`

UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


