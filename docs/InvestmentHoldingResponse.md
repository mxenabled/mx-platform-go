# InvestmentHoldingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGuid** | Pointer to **NullableString** |  | [optional] 
**CostBasis** | Pointer to **NullableFloat32** |  | [optional] 
**CouponYield** | Pointer to **NullableString** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**CurrentPrice** | Pointer to **NullableFloat32** |  | [optional] 
**DailyChange** | Pointer to **NullableFloat32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Expiration** | Pointer to **NullableString** |  | [optional] 
**FaceValue** | Pointer to **NullableFloat32** |  | [optional] 
**Frequency** | Pointer to **NullableString** |  | [optional] 
**Guid** | Pointer to **NullableString** |  | [optional] 
**MarketValue** | Pointer to **NullableFloat32** |  | [optional] 
**MaturityDate** | Pointer to **NullableString** |  | [optional] 
**PercentageChange** | Pointer to **NullableFloat32** |  | [optional] 
**PurchasePrice** | Pointer to **NullableFloat32** |  | [optional] 
**Quantity** | Pointer to **NullableString** |  | [optional] 
**Rate** | Pointer to **NullableFloat32** |  | [optional] 
**StrikePrice** | Pointer to **NullableFloat32** |  | [optional] 
**Symbol** | Pointer to **NullableString** |  | [optional] 
**Term** | Pointer to **NullableString** |  | [optional] 
**TodayUglAmount** | Pointer to **NullableFloat32** |  | [optional] 
**TodayUglPercentage** | Pointer to **NullableFloat32** |  | [optional] 
**TotalUglAmount** | Pointer to **NullableFloat32** |  | [optional] 
**TotalUglPercentage** | Pointer to **NullableFloat32** |  | [optional] 
**UnvestedQuantity** | Pointer to **NullableFloat32** |  | [optional] 
**UnvestedValue** | Pointer to **NullableFloat32** |  | [optional] 
**UserGuid** | Pointer to **NullableString** |  | [optional] 
**VestedQuantity** | Pointer to **NullableFloat32** |  | [optional] 
**VestedValue** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**CurrentPriceAsOf** | Pointer to **NullableString** |  | [optional] 
**IssueDate** | Pointer to **NullableString** |  | [optional] 
**VestingStartDate** | Pointer to **NullableString** |  | [optional] 
**VestingEndDate** | Pointer to **NullableString** |  | [optional] 
**PutOrCall** | Pointer to **NullableString** |  | [optional] 
**HoldingType** | Pointer to **NullableString** |  | [optional] 
**TermUnit** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvestmentHoldingResponse

`func NewInvestmentHoldingResponse() *InvestmentHoldingResponse`

NewInvestmentHoldingResponse instantiates a new InvestmentHoldingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentHoldingResponseWithDefaults

`func NewInvestmentHoldingResponseWithDefaults() *InvestmentHoldingResponse`

NewInvestmentHoldingResponseWithDefaults instantiates a new InvestmentHoldingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGuid

`func (o *InvestmentHoldingResponse) GetAccountGuid() string`

GetAccountGuid returns the AccountGuid field if non-nil, zero value otherwise.

### GetAccountGuidOk

`func (o *InvestmentHoldingResponse) GetAccountGuidOk() (*string, bool)`

GetAccountGuidOk returns a tuple with the AccountGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGuid

`func (o *InvestmentHoldingResponse) SetAccountGuid(v string)`

SetAccountGuid sets AccountGuid field to given value.

### HasAccountGuid

`func (o *InvestmentHoldingResponse) HasAccountGuid() bool`

HasAccountGuid returns a boolean if a field has been set.

### SetAccountGuidNil

`func (o *InvestmentHoldingResponse) SetAccountGuidNil(b bool)`

 SetAccountGuidNil sets the value for AccountGuid to be an explicit nil

### UnsetAccountGuid
`func (o *InvestmentHoldingResponse) UnsetAccountGuid()`

UnsetAccountGuid ensures that no value is present for AccountGuid, not even an explicit nil
### GetCostBasis

`func (o *InvestmentHoldingResponse) GetCostBasis() float32`

GetCostBasis returns the CostBasis field if non-nil, zero value otherwise.

### GetCostBasisOk

`func (o *InvestmentHoldingResponse) GetCostBasisOk() (*float32, bool)`

GetCostBasisOk returns a tuple with the CostBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBasis

`func (o *InvestmentHoldingResponse) SetCostBasis(v float32)`

SetCostBasis sets CostBasis field to given value.

### HasCostBasis

`func (o *InvestmentHoldingResponse) HasCostBasis() bool`

HasCostBasis returns a boolean if a field has been set.

### SetCostBasisNil

`func (o *InvestmentHoldingResponse) SetCostBasisNil(b bool)`

 SetCostBasisNil sets the value for CostBasis to be an explicit nil

### UnsetCostBasis
`func (o *InvestmentHoldingResponse) UnsetCostBasis()`

UnsetCostBasis ensures that no value is present for CostBasis, not even an explicit nil
### GetCouponYield

`func (o *InvestmentHoldingResponse) GetCouponYield() string`

GetCouponYield returns the CouponYield field if non-nil, zero value otherwise.

### GetCouponYieldOk

`func (o *InvestmentHoldingResponse) GetCouponYieldOk() (*string, bool)`

GetCouponYieldOk returns a tuple with the CouponYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponYield

`func (o *InvestmentHoldingResponse) SetCouponYield(v string)`

SetCouponYield sets CouponYield field to given value.

### HasCouponYield

`func (o *InvestmentHoldingResponse) HasCouponYield() bool`

HasCouponYield returns a boolean if a field has been set.

### SetCouponYieldNil

`func (o *InvestmentHoldingResponse) SetCouponYieldNil(b bool)`

 SetCouponYieldNil sets the value for CouponYield to be an explicit nil

### UnsetCouponYield
`func (o *InvestmentHoldingResponse) UnsetCouponYield()`

UnsetCouponYield ensures that no value is present for CouponYield, not even an explicit nil
### GetCurrencyCode

`func (o *InvestmentHoldingResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InvestmentHoldingResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InvestmentHoldingResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *InvestmentHoldingResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *InvestmentHoldingResponse) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *InvestmentHoldingResponse) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetCurrentPrice

`func (o *InvestmentHoldingResponse) GetCurrentPrice() float32`

GetCurrentPrice returns the CurrentPrice field if non-nil, zero value otherwise.

### GetCurrentPriceOk

`func (o *InvestmentHoldingResponse) GetCurrentPriceOk() (*float32, bool)`

GetCurrentPriceOk returns a tuple with the CurrentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPrice

`func (o *InvestmentHoldingResponse) SetCurrentPrice(v float32)`

SetCurrentPrice sets CurrentPrice field to given value.

### HasCurrentPrice

`func (o *InvestmentHoldingResponse) HasCurrentPrice() bool`

HasCurrentPrice returns a boolean if a field has been set.

### SetCurrentPriceNil

`func (o *InvestmentHoldingResponse) SetCurrentPriceNil(b bool)`

 SetCurrentPriceNil sets the value for CurrentPrice to be an explicit nil

### UnsetCurrentPrice
`func (o *InvestmentHoldingResponse) UnsetCurrentPrice()`

UnsetCurrentPrice ensures that no value is present for CurrentPrice, not even an explicit nil
### GetDailyChange

`func (o *InvestmentHoldingResponse) GetDailyChange() float32`

GetDailyChange returns the DailyChange field if non-nil, zero value otherwise.

### GetDailyChangeOk

`func (o *InvestmentHoldingResponse) GetDailyChangeOk() (*float32, bool)`

GetDailyChangeOk returns a tuple with the DailyChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyChange

`func (o *InvestmentHoldingResponse) SetDailyChange(v float32)`

SetDailyChange sets DailyChange field to given value.

### HasDailyChange

`func (o *InvestmentHoldingResponse) HasDailyChange() bool`

HasDailyChange returns a boolean if a field has been set.

### SetDailyChangeNil

`func (o *InvestmentHoldingResponse) SetDailyChangeNil(b bool)`

 SetDailyChangeNil sets the value for DailyChange to be an explicit nil

### UnsetDailyChange
`func (o *InvestmentHoldingResponse) UnsetDailyChange()`

UnsetDailyChange ensures that no value is present for DailyChange, not even an explicit nil
### GetDescription

`func (o *InvestmentHoldingResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvestmentHoldingResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvestmentHoldingResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvestmentHoldingResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InvestmentHoldingResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvestmentHoldingResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpiration

`func (o *InvestmentHoldingResponse) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *InvestmentHoldingResponse) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *InvestmentHoldingResponse) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *InvestmentHoldingResponse) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *InvestmentHoldingResponse) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *InvestmentHoldingResponse) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetFaceValue

`func (o *InvestmentHoldingResponse) GetFaceValue() float32`

GetFaceValue returns the FaceValue field if non-nil, zero value otherwise.

### GetFaceValueOk

`func (o *InvestmentHoldingResponse) GetFaceValueOk() (*float32, bool)`

GetFaceValueOk returns a tuple with the FaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceValue

`func (o *InvestmentHoldingResponse) SetFaceValue(v float32)`

SetFaceValue sets FaceValue field to given value.

### HasFaceValue

`func (o *InvestmentHoldingResponse) HasFaceValue() bool`

HasFaceValue returns a boolean if a field has been set.

### SetFaceValueNil

`func (o *InvestmentHoldingResponse) SetFaceValueNil(b bool)`

 SetFaceValueNil sets the value for FaceValue to be an explicit nil

### UnsetFaceValue
`func (o *InvestmentHoldingResponse) UnsetFaceValue()`

UnsetFaceValue ensures that no value is present for FaceValue, not even an explicit nil
### GetFrequency

`func (o *InvestmentHoldingResponse) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *InvestmentHoldingResponse) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *InvestmentHoldingResponse) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *InvestmentHoldingResponse) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *InvestmentHoldingResponse) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *InvestmentHoldingResponse) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetGuid

`func (o *InvestmentHoldingResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *InvestmentHoldingResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *InvestmentHoldingResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *InvestmentHoldingResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *InvestmentHoldingResponse) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *InvestmentHoldingResponse) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetMarketValue

`func (o *InvestmentHoldingResponse) GetMarketValue() float32`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *InvestmentHoldingResponse) GetMarketValueOk() (*float32, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *InvestmentHoldingResponse) SetMarketValue(v float32)`

SetMarketValue sets MarketValue field to given value.

### HasMarketValue

`func (o *InvestmentHoldingResponse) HasMarketValue() bool`

HasMarketValue returns a boolean if a field has been set.

### SetMarketValueNil

`func (o *InvestmentHoldingResponse) SetMarketValueNil(b bool)`

 SetMarketValueNil sets the value for MarketValue to be an explicit nil

### UnsetMarketValue
`func (o *InvestmentHoldingResponse) UnsetMarketValue()`

UnsetMarketValue ensures that no value is present for MarketValue, not even an explicit nil
### GetMaturityDate

`func (o *InvestmentHoldingResponse) GetMaturityDate() string`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *InvestmentHoldingResponse) GetMaturityDateOk() (*string, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *InvestmentHoldingResponse) SetMaturityDate(v string)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *InvestmentHoldingResponse) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### SetMaturityDateNil

`func (o *InvestmentHoldingResponse) SetMaturityDateNil(b bool)`

 SetMaturityDateNil sets the value for MaturityDate to be an explicit nil

### UnsetMaturityDate
`func (o *InvestmentHoldingResponse) UnsetMaturityDate()`

UnsetMaturityDate ensures that no value is present for MaturityDate, not even an explicit nil
### GetPercentageChange

`func (o *InvestmentHoldingResponse) GetPercentageChange() float32`

GetPercentageChange returns the PercentageChange field if non-nil, zero value otherwise.

### GetPercentageChangeOk

`func (o *InvestmentHoldingResponse) GetPercentageChangeOk() (*float32, bool)`

GetPercentageChangeOk returns a tuple with the PercentageChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageChange

`func (o *InvestmentHoldingResponse) SetPercentageChange(v float32)`

SetPercentageChange sets PercentageChange field to given value.

### HasPercentageChange

`func (o *InvestmentHoldingResponse) HasPercentageChange() bool`

HasPercentageChange returns a boolean if a field has been set.

### SetPercentageChangeNil

`func (o *InvestmentHoldingResponse) SetPercentageChangeNil(b bool)`

 SetPercentageChangeNil sets the value for PercentageChange to be an explicit nil

### UnsetPercentageChange
`func (o *InvestmentHoldingResponse) UnsetPercentageChange()`

UnsetPercentageChange ensures that no value is present for PercentageChange, not even an explicit nil
### GetPurchasePrice

`func (o *InvestmentHoldingResponse) GetPurchasePrice() float32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *InvestmentHoldingResponse) GetPurchasePriceOk() (*float32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *InvestmentHoldingResponse) SetPurchasePrice(v float32)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *InvestmentHoldingResponse) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### SetPurchasePriceNil

`func (o *InvestmentHoldingResponse) SetPurchasePriceNil(b bool)`

 SetPurchasePriceNil sets the value for PurchasePrice to be an explicit nil

### UnsetPurchasePrice
`func (o *InvestmentHoldingResponse) UnsetPurchasePrice()`

UnsetPurchasePrice ensures that no value is present for PurchasePrice, not even an explicit nil
### GetQuantity

`func (o *InvestmentHoldingResponse) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvestmentHoldingResponse) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvestmentHoldingResponse) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvestmentHoldingResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *InvestmentHoldingResponse) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *InvestmentHoldingResponse) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetRate

`func (o *InvestmentHoldingResponse) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *InvestmentHoldingResponse) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *InvestmentHoldingResponse) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *InvestmentHoldingResponse) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRateNil

`func (o *InvestmentHoldingResponse) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *InvestmentHoldingResponse) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetStrikePrice

`func (o *InvestmentHoldingResponse) GetStrikePrice() float32`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *InvestmentHoldingResponse) GetStrikePriceOk() (*float32, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *InvestmentHoldingResponse) SetStrikePrice(v float32)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *InvestmentHoldingResponse) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### SetStrikePriceNil

`func (o *InvestmentHoldingResponse) SetStrikePriceNil(b bool)`

 SetStrikePriceNil sets the value for StrikePrice to be an explicit nil

### UnsetStrikePrice
`func (o *InvestmentHoldingResponse) UnsetStrikePrice()`

UnsetStrikePrice ensures that no value is present for StrikePrice, not even an explicit nil
### GetSymbol

`func (o *InvestmentHoldingResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InvestmentHoldingResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InvestmentHoldingResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InvestmentHoldingResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### SetSymbolNil

`func (o *InvestmentHoldingResponse) SetSymbolNil(b bool)`

 SetSymbolNil sets the value for Symbol to be an explicit nil

### UnsetSymbol
`func (o *InvestmentHoldingResponse) UnsetSymbol()`

UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil
### GetTerm

`func (o *InvestmentHoldingResponse) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *InvestmentHoldingResponse) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *InvestmentHoldingResponse) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *InvestmentHoldingResponse) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### SetTermNil

`func (o *InvestmentHoldingResponse) SetTermNil(b bool)`

 SetTermNil sets the value for Term to be an explicit nil

### UnsetTerm
`func (o *InvestmentHoldingResponse) UnsetTerm()`

UnsetTerm ensures that no value is present for Term, not even an explicit nil
### GetTodayUglAmount

`func (o *InvestmentHoldingResponse) GetTodayUglAmount() float32`

GetTodayUglAmount returns the TodayUglAmount field if non-nil, zero value otherwise.

### GetTodayUglAmountOk

`func (o *InvestmentHoldingResponse) GetTodayUglAmountOk() (*float32, bool)`

GetTodayUglAmountOk returns a tuple with the TodayUglAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodayUglAmount

`func (o *InvestmentHoldingResponse) SetTodayUglAmount(v float32)`

SetTodayUglAmount sets TodayUglAmount field to given value.

### HasTodayUglAmount

`func (o *InvestmentHoldingResponse) HasTodayUglAmount() bool`

HasTodayUglAmount returns a boolean if a field has been set.

### SetTodayUglAmountNil

`func (o *InvestmentHoldingResponse) SetTodayUglAmountNil(b bool)`

 SetTodayUglAmountNil sets the value for TodayUglAmount to be an explicit nil

### UnsetTodayUglAmount
`func (o *InvestmentHoldingResponse) UnsetTodayUglAmount()`

UnsetTodayUglAmount ensures that no value is present for TodayUglAmount, not even an explicit nil
### GetTodayUglPercentage

`func (o *InvestmentHoldingResponse) GetTodayUglPercentage() float32`

GetTodayUglPercentage returns the TodayUglPercentage field if non-nil, zero value otherwise.

### GetTodayUglPercentageOk

`func (o *InvestmentHoldingResponse) GetTodayUglPercentageOk() (*float32, bool)`

GetTodayUglPercentageOk returns a tuple with the TodayUglPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodayUglPercentage

`func (o *InvestmentHoldingResponse) SetTodayUglPercentage(v float32)`

SetTodayUglPercentage sets TodayUglPercentage field to given value.

### HasTodayUglPercentage

`func (o *InvestmentHoldingResponse) HasTodayUglPercentage() bool`

HasTodayUglPercentage returns a boolean if a field has been set.

### SetTodayUglPercentageNil

`func (o *InvestmentHoldingResponse) SetTodayUglPercentageNil(b bool)`

 SetTodayUglPercentageNil sets the value for TodayUglPercentage to be an explicit nil

### UnsetTodayUglPercentage
`func (o *InvestmentHoldingResponse) UnsetTodayUglPercentage()`

UnsetTodayUglPercentage ensures that no value is present for TodayUglPercentage, not even an explicit nil
### GetTotalUglAmount

`func (o *InvestmentHoldingResponse) GetTotalUglAmount() float32`

GetTotalUglAmount returns the TotalUglAmount field if non-nil, zero value otherwise.

### GetTotalUglAmountOk

`func (o *InvestmentHoldingResponse) GetTotalUglAmountOk() (*float32, bool)`

GetTotalUglAmountOk returns a tuple with the TotalUglAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUglAmount

`func (o *InvestmentHoldingResponse) SetTotalUglAmount(v float32)`

SetTotalUglAmount sets TotalUglAmount field to given value.

### HasTotalUglAmount

`func (o *InvestmentHoldingResponse) HasTotalUglAmount() bool`

HasTotalUglAmount returns a boolean if a field has been set.

### SetTotalUglAmountNil

`func (o *InvestmentHoldingResponse) SetTotalUglAmountNil(b bool)`

 SetTotalUglAmountNil sets the value for TotalUglAmount to be an explicit nil

### UnsetTotalUglAmount
`func (o *InvestmentHoldingResponse) UnsetTotalUglAmount()`

UnsetTotalUglAmount ensures that no value is present for TotalUglAmount, not even an explicit nil
### GetTotalUglPercentage

`func (o *InvestmentHoldingResponse) GetTotalUglPercentage() float32`

GetTotalUglPercentage returns the TotalUglPercentage field if non-nil, zero value otherwise.

### GetTotalUglPercentageOk

`func (o *InvestmentHoldingResponse) GetTotalUglPercentageOk() (*float32, bool)`

GetTotalUglPercentageOk returns a tuple with the TotalUglPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUglPercentage

`func (o *InvestmentHoldingResponse) SetTotalUglPercentage(v float32)`

SetTotalUglPercentage sets TotalUglPercentage field to given value.

### HasTotalUglPercentage

`func (o *InvestmentHoldingResponse) HasTotalUglPercentage() bool`

HasTotalUglPercentage returns a boolean if a field has been set.

### SetTotalUglPercentageNil

`func (o *InvestmentHoldingResponse) SetTotalUglPercentageNil(b bool)`

 SetTotalUglPercentageNil sets the value for TotalUglPercentage to be an explicit nil

### UnsetTotalUglPercentage
`func (o *InvestmentHoldingResponse) UnsetTotalUglPercentage()`

UnsetTotalUglPercentage ensures that no value is present for TotalUglPercentage, not even an explicit nil
### GetUnvestedQuantity

`func (o *InvestmentHoldingResponse) GetUnvestedQuantity() float32`

GetUnvestedQuantity returns the UnvestedQuantity field if non-nil, zero value otherwise.

### GetUnvestedQuantityOk

`func (o *InvestmentHoldingResponse) GetUnvestedQuantityOk() (*float32, bool)`

GetUnvestedQuantityOk returns a tuple with the UnvestedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnvestedQuantity

`func (o *InvestmentHoldingResponse) SetUnvestedQuantity(v float32)`

SetUnvestedQuantity sets UnvestedQuantity field to given value.

### HasUnvestedQuantity

`func (o *InvestmentHoldingResponse) HasUnvestedQuantity() bool`

HasUnvestedQuantity returns a boolean if a field has been set.

### SetUnvestedQuantityNil

`func (o *InvestmentHoldingResponse) SetUnvestedQuantityNil(b bool)`

 SetUnvestedQuantityNil sets the value for UnvestedQuantity to be an explicit nil

### UnsetUnvestedQuantity
`func (o *InvestmentHoldingResponse) UnsetUnvestedQuantity()`

UnsetUnvestedQuantity ensures that no value is present for UnvestedQuantity, not even an explicit nil
### GetUnvestedValue

`func (o *InvestmentHoldingResponse) GetUnvestedValue() float32`

GetUnvestedValue returns the UnvestedValue field if non-nil, zero value otherwise.

### GetUnvestedValueOk

`func (o *InvestmentHoldingResponse) GetUnvestedValueOk() (*float32, bool)`

GetUnvestedValueOk returns a tuple with the UnvestedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnvestedValue

`func (o *InvestmentHoldingResponse) SetUnvestedValue(v float32)`

SetUnvestedValue sets UnvestedValue field to given value.

### HasUnvestedValue

`func (o *InvestmentHoldingResponse) HasUnvestedValue() bool`

HasUnvestedValue returns a boolean if a field has been set.

### SetUnvestedValueNil

`func (o *InvestmentHoldingResponse) SetUnvestedValueNil(b bool)`

 SetUnvestedValueNil sets the value for UnvestedValue to be an explicit nil

### UnsetUnvestedValue
`func (o *InvestmentHoldingResponse) UnsetUnvestedValue()`

UnsetUnvestedValue ensures that no value is present for UnvestedValue, not even an explicit nil
### GetUserGuid

`func (o *InvestmentHoldingResponse) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *InvestmentHoldingResponse) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *InvestmentHoldingResponse) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *InvestmentHoldingResponse) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### SetUserGuidNil

`func (o *InvestmentHoldingResponse) SetUserGuidNil(b bool)`

 SetUserGuidNil sets the value for UserGuid to be an explicit nil

### UnsetUserGuid
`func (o *InvestmentHoldingResponse) UnsetUserGuid()`

UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
### GetVestedQuantity

`func (o *InvestmentHoldingResponse) GetVestedQuantity() float32`

GetVestedQuantity returns the VestedQuantity field if non-nil, zero value otherwise.

### GetVestedQuantityOk

`func (o *InvestmentHoldingResponse) GetVestedQuantityOk() (*float32, bool)`

GetVestedQuantityOk returns a tuple with the VestedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVestedQuantity

`func (o *InvestmentHoldingResponse) SetVestedQuantity(v float32)`

SetVestedQuantity sets VestedQuantity field to given value.

### HasVestedQuantity

`func (o *InvestmentHoldingResponse) HasVestedQuantity() bool`

HasVestedQuantity returns a boolean if a field has been set.

### SetVestedQuantityNil

`func (o *InvestmentHoldingResponse) SetVestedQuantityNil(b bool)`

 SetVestedQuantityNil sets the value for VestedQuantity to be an explicit nil

### UnsetVestedQuantity
`func (o *InvestmentHoldingResponse) UnsetVestedQuantity()`

UnsetVestedQuantity ensures that no value is present for VestedQuantity, not even an explicit nil
### GetVestedValue

`func (o *InvestmentHoldingResponse) GetVestedValue() float32`

GetVestedValue returns the VestedValue field if non-nil, zero value otherwise.

### GetVestedValueOk

`func (o *InvestmentHoldingResponse) GetVestedValueOk() (*float32, bool)`

GetVestedValueOk returns a tuple with the VestedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVestedValue

`func (o *InvestmentHoldingResponse) SetVestedValue(v float32)`

SetVestedValue sets VestedValue field to given value.

### HasVestedValue

`func (o *InvestmentHoldingResponse) HasVestedValue() bool`

HasVestedValue returns a boolean if a field has been set.

### SetVestedValueNil

`func (o *InvestmentHoldingResponse) SetVestedValueNil(b bool)`

 SetVestedValueNil sets the value for VestedValue to be an explicit nil

### UnsetVestedValue
`func (o *InvestmentHoldingResponse) UnsetVestedValue()`

UnsetVestedValue ensures that no value is present for VestedValue, not even an explicit nil
### GetCreatedAt

`func (o *InvestmentHoldingResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvestmentHoldingResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvestmentHoldingResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InvestmentHoldingResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *InvestmentHoldingResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *InvestmentHoldingResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCurrentPriceAsOf

`func (o *InvestmentHoldingResponse) GetCurrentPriceAsOf() string`

GetCurrentPriceAsOf returns the CurrentPriceAsOf field if non-nil, zero value otherwise.

### GetCurrentPriceAsOfOk

`func (o *InvestmentHoldingResponse) GetCurrentPriceAsOfOk() (*string, bool)`

GetCurrentPriceAsOfOk returns a tuple with the CurrentPriceAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPriceAsOf

`func (o *InvestmentHoldingResponse) SetCurrentPriceAsOf(v string)`

SetCurrentPriceAsOf sets CurrentPriceAsOf field to given value.

### HasCurrentPriceAsOf

`func (o *InvestmentHoldingResponse) HasCurrentPriceAsOf() bool`

HasCurrentPriceAsOf returns a boolean if a field has been set.

### SetCurrentPriceAsOfNil

`func (o *InvestmentHoldingResponse) SetCurrentPriceAsOfNil(b bool)`

 SetCurrentPriceAsOfNil sets the value for CurrentPriceAsOf to be an explicit nil

### UnsetCurrentPriceAsOf
`func (o *InvestmentHoldingResponse) UnsetCurrentPriceAsOf()`

UnsetCurrentPriceAsOf ensures that no value is present for CurrentPriceAsOf, not even an explicit nil
### GetIssueDate

`func (o *InvestmentHoldingResponse) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *InvestmentHoldingResponse) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *InvestmentHoldingResponse) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *InvestmentHoldingResponse) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *InvestmentHoldingResponse) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *InvestmentHoldingResponse) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetVestingStartDate

`func (o *InvestmentHoldingResponse) GetVestingStartDate() string`

GetVestingStartDate returns the VestingStartDate field if non-nil, zero value otherwise.

### GetVestingStartDateOk

`func (o *InvestmentHoldingResponse) GetVestingStartDateOk() (*string, bool)`

GetVestingStartDateOk returns a tuple with the VestingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVestingStartDate

`func (o *InvestmentHoldingResponse) SetVestingStartDate(v string)`

SetVestingStartDate sets VestingStartDate field to given value.

### HasVestingStartDate

`func (o *InvestmentHoldingResponse) HasVestingStartDate() bool`

HasVestingStartDate returns a boolean if a field has been set.

### SetVestingStartDateNil

`func (o *InvestmentHoldingResponse) SetVestingStartDateNil(b bool)`

 SetVestingStartDateNil sets the value for VestingStartDate to be an explicit nil

### UnsetVestingStartDate
`func (o *InvestmentHoldingResponse) UnsetVestingStartDate()`

UnsetVestingStartDate ensures that no value is present for VestingStartDate, not even an explicit nil
### GetVestingEndDate

`func (o *InvestmentHoldingResponse) GetVestingEndDate() string`

GetVestingEndDate returns the VestingEndDate field if non-nil, zero value otherwise.

### GetVestingEndDateOk

`func (o *InvestmentHoldingResponse) GetVestingEndDateOk() (*string, bool)`

GetVestingEndDateOk returns a tuple with the VestingEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVestingEndDate

`func (o *InvestmentHoldingResponse) SetVestingEndDate(v string)`

SetVestingEndDate sets VestingEndDate field to given value.

### HasVestingEndDate

`func (o *InvestmentHoldingResponse) HasVestingEndDate() bool`

HasVestingEndDate returns a boolean if a field has been set.

### SetVestingEndDateNil

`func (o *InvestmentHoldingResponse) SetVestingEndDateNil(b bool)`

 SetVestingEndDateNil sets the value for VestingEndDate to be an explicit nil

### UnsetVestingEndDate
`func (o *InvestmentHoldingResponse) UnsetVestingEndDate()`

UnsetVestingEndDate ensures that no value is present for VestingEndDate, not even an explicit nil
### GetPutOrCall

`func (o *InvestmentHoldingResponse) GetPutOrCall() string`

GetPutOrCall returns the PutOrCall field if non-nil, zero value otherwise.

### GetPutOrCallOk

`func (o *InvestmentHoldingResponse) GetPutOrCallOk() (*string, bool)`

GetPutOrCallOk returns a tuple with the PutOrCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutOrCall

`func (o *InvestmentHoldingResponse) SetPutOrCall(v string)`

SetPutOrCall sets PutOrCall field to given value.

### HasPutOrCall

`func (o *InvestmentHoldingResponse) HasPutOrCall() bool`

HasPutOrCall returns a boolean if a field has been set.

### SetPutOrCallNil

`func (o *InvestmentHoldingResponse) SetPutOrCallNil(b bool)`

 SetPutOrCallNil sets the value for PutOrCall to be an explicit nil

### UnsetPutOrCall
`func (o *InvestmentHoldingResponse) UnsetPutOrCall()`

UnsetPutOrCall ensures that no value is present for PutOrCall, not even an explicit nil
### GetHoldingType

`func (o *InvestmentHoldingResponse) GetHoldingType() string`

GetHoldingType returns the HoldingType field if non-nil, zero value otherwise.

### GetHoldingTypeOk

`func (o *InvestmentHoldingResponse) GetHoldingTypeOk() (*string, bool)`

GetHoldingTypeOk returns a tuple with the HoldingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingType

`func (o *InvestmentHoldingResponse) SetHoldingType(v string)`

SetHoldingType sets HoldingType field to given value.

### HasHoldingType

`func (o *InvestmentHoldingResponse) HasHoldingType() bool`

HasHoldingType returns a boolean if a field has been set.

### SetHoldingTypeNil

`func (o *InvestmentHoldingResponse) SetHoldingTypeNil(b bool)`

 SetHoldingTypeNil sets the value for HoldingType to be an explicit nil

### UnsetHoldingType
`func (o *InvestmentHoldingResponse) UnsetHoldingType()`

UnsetHoldingType ensures that no value is present for HoldingType, not even an explicit nil
### GetTermUnit

`func (o *InvestmentHoldingResponse) GetTermUnit() string`

GetTermUnit returns the TermUnit field if non-nil, zero value otherwise.

### GetTermUnitOk

`func (o *InvestmentHoldingResponse) GetTermUnitOk() (*string, bool)`

GetTermUnitOk returns a tuple with the TermUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermUnit

`func (o *InvestmentHoldingResponse) SetTermUnit(v string)`

SetTermUnit sets TermUnit field to given value.

### HasTermUnit

`func (o *InvestmentHoldingResponse) HasTermUnit() bool`

HasTermUnit returns a boolean if a field has been set.

### SetTermUnitNil

`func (o *InvestmentHoldingResponse) SetTermUnitNil(b bool)`

 SetTermUnitNil sets the value for TermUnit to be an explicit nil

### UnsetTermUnit
`func (o *InvestmentHoldingResponse) UnsetTermUnit()`

UnsetTermUnit ensures that no value is present for TermUnit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


