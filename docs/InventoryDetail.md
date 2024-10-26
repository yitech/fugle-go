# InventoryDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuySell** | **string** |  | 
**CostR** | Pointer to **NullableFloat32** |  | [optional] 
**Fee** | **float32** |  | 
**MakeA** | **float32** |  | 
**MakeAPer** | **float32** |  | 
**OrdNo** | **string** |  | 
**PayN** | **float32** |  | 
**Price** | **float32** |  | 
**PriceEvn** | **float32** |  | 
**Qty** | **int32** |  | 
**QtyC** | Pointer to **NullableInt32** |  | [optional] 
**QtyH** | Pointer to **NullableInt32** |  | [optional] 
**QtyR** | Pointer to **NullableInt32** |  | [optional] 
**TDate** | **string** |  | 
**TTime** | Pointer to **NullableString** |  | [optional] 
**Tax** | Pointer to **NullableFloat32** |  | [optional] 
**TaxG** | Pointer to **NullableFloat32** |  | [optional] 
**Trade** | Pointer to **NullableInt32** |  | [optional] 
**ValueMkt** | **float32** |  | 
**ValueNow** | **float32** |  | 
**UserDef** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInventoryDetail

`func NewInventoryDetail(buySell string, fee float32, makeA float32, makeAPer float32, ordNo string, payN float32, price float32, priceEvn float32, qty int32, tDate string, valueMkt float32, valueNow float32, ) *InventoryDetail`

NewInventoryDetail instantiates a new InventoryDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryDetailWithDefaults

`func NewInventoryDetailWithDefaults() *InventoryDetail`

NewInventoryDetailWithDefaults instantiates a new InventoryDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuySell

`func (o *InventoryDetail) GetBuySell() string`

GetBuySell returns the BuySell field if non-nil, zero value otherwise.

### GetBuySellOk

`func (o *InventoryDetail) GetBuySellOk() (*string, bool)`

GetBuySellOk returns a tuple with the BuySell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuySell

`func (o *InventoryDetail) SetBuySell(v string)`

SetBuySell sets BuySell field to given value.


### GetCostR

`func (o *InventoryDetail) GetCostR() float32`

GetCostR returns the CostR field if non-nil, zero value otherwise.

### GetCostROk

`func (o *InventoryDetail) GetCostROk() (*float32, bool)`

GetCostROk returns a tuple with the CostR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostR

`func (o *InventoryDetail) SetCostR(v float32)`

SetCostR sets CostR field to given value.

### HasCostR

`func (o *InventoryDetail) HasCostR() bool`

HasCostR returns a boolean if a field has been set.

### SetCostRNil

`func (o *InventoryDetail) SetCostRNil(b bool)`

 SetCostRNil sets the value for CostR to be an explicit nil

### UnsetCostR
`func (o *InventoryDetail) UnsetCostR()`

UnsetCostR ensures that no value is present for CostR, not even an explicit nil
### GetFee

`func (o *InventoryDetail) GetFee() float32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *InventoryDetail) GetFeeOk() (*float32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *InventoryDetail) SetFee(v float32)`

SetFee sets Fee field to given value.


### GetMakeA

`func (o *InventoryDetail) GetMakeA() float32`

GetMakeA returns the MakeA field if non-nil, zero value otherwise.

### GetMakeAOk

`func (o *InventoryDetail) GetMakeAOk() (*float32, bool)`

GetMakeAOk returns a tuple with the MakeA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeA

`func (o *InventoryDetail) SetMakeA(v float32)`

SetMakeA sets MakeA field to given value.


### GetMakeAPer

`func (o *InventoryDetail) GetMakeAPer() float32`

GetMakeAPer returns the MakeAPer field if non-nil, zero value otherwise.

### GetMakeAPerOk

`func (o *InventoryDetail) GetMakeAPerOk() (*float32, bool)`

GetMakeAPerOk returns a tuple with the MakeAPer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeAPer

`func (o *InventoryDetail) SetMakeAPer(v float32)`

SetMakeAPer sets MakeAPer field to given value.


### GetOrdNo

`func (o *InventoryDetail) GetOrdNo() string`

GetOrdNo returns the OrdNo field if non-nil, zero value otherwise.

### GetOrdNoOk

`func (o *InventoryDetail) GetOrdNoOk() (*string, bool)`

GetOrdNoOk returns a tuple with the OrdNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdNo

`func (o *InventoryDetail) SetOrdNo(v string)`

SetOrdNo sets OrdNo field to given value.


### GetPayN

`func (o *InventoryDetail) GetPayN() float32`

GetPayN returns the PayN field if non-nil, zero value otherwise.

### GetPayNOk

`func (o *InventoryDetail) GetPayNOk() (*float32, bool)`

GetPayNOk returns a tuple with the PayN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayN

`func (o *InventoryDetail) SetPayN(v float32)`

SetPayN sets PayN field to given value.


### GetPrice

`func (o *InventoryDetail) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InventoryDetail) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InventoryDetail) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetPriceEvn

`func (o *InventoryDetail) GetPriceEvn() float32`

GetPriceEvn returns the PriceEvn field if non-nil, zero value otherwise.

### GetPriceEvnOk

`func (o *InventoryDetail) GetPriceEvnOk() (*float32, bool)`

GetPriceEvnOk returns a tuple with the PriceEvn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceEvn

`func (o *InventoryDetail) SetPriceEvn(v float32)`

SetPriceEvn sets PriceEvn field to given value.


### GetQty

`func (o *InventoryDetail) GetQty() int32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *InventoryDetail) GetQtyOk() (*int32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *InventoryDetail) SetQty(v int32)`

SetQty sets Qty field to given value.


### GetQtyC

`func (o *InventoryDetail) GetQtyC() int32`

GetQtyC returns the QtyC field if non-nil, zero value otherwise.

### GetQtyCOk

`func (o *InventoryDetail) GetQtyCOk() (*int32, bool)`

GetQtyCOk returns a tuple with the QtyC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyC

`func (o *InventoryDetail) SetQtyC(v int32)`

SetQtyC sets QtyC field to given value.

### HasQtyC

`func (o *InventoryDetail) HasQtyC() bool`

HasQtyC returns a boolean if a field has been set.

### SetQtyCNil

`func (o *InventoryDetail) SetQtyCNil(b bool)`

 SetQtyCNil sets the value for QtyC to be an explicit nil

### UnsetQtyC
`func (o *InventoryDetail) UnsetQtyC()`

UnsetQtyC ensures that no value is present for QtyC, not even an explicit nil
### GetQtyH

`func (o *InventoryDetail) GetQtyH() int32`

GetQtyH returns the QtyH field if non-nil, zero value otherwise.

### GetQtyHOk

`func (o *InventoryDetail) GetQtyHOk() (*int32, bool)`

GetQtyHOk returns a tuple with the QtyH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyH

`func (o *InventoryDetail) SetQtyH(v int32)`

SetQtyH sets QtyH field to given value.

### HasQtyH

`func (o *InventoryDetail) HasQtyH() bool`

HasQtyH returns a boolean if a field has been set.

### SetQtyHNil

`func (o *InventoryDetail) SetQtyHNil(b bool)`

 SetQtyHNil sets the value for QtyH to be an explicit nil

### UnsetQtyH
`func (o *InventoryDetail) UnsetQtyH()`

UnsetQtyH ensures that no value is present for QtyH, not even an explicit nil
### GetQtyR

`func (o *InventoryDetail) GetQtyR() int32`

GetQtyR returns the QtyR field if non-nil, zero value otherwise.

### GetQtyROk

`func (o *InventoryDetail) GetQtyROk() (*int32, bool)`

GetQtyROk returns a tuple with the QtyR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyR

`func (o *InventoryDetail) SetQtyR(v int32)`

SetQtyR sets QtyR field to given value.

### HasQtyR

`func (o *InventoryDetail) HasQtyR() bool`

HasQtyR returns a boolean if a field has been set.

### SetQtyRNil

`func (o *InventoryDetail) SetQtyRNil(b bool)`

 SetQtyRNil sets the value for QtyR to be an explicit nil

### UnsetQtyR
`func (o *InventoryDetail) UnsetQtyR()`

UnsetQtyR ensures that no value is present for QtyR, not even an explicit nil
### GetTDate

`func (o *InventoryDetail) GetTDate() string`

GetTDate returns the TDate field if non-nil, zero value otherwise.

### GetTDateOk

`func (o *InventoryDetail) GetTDateOk() (*string, bool)`

GetTDateOk returns a tuple with the TDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTDate

`func (o *InventoryDetail) SetTDate(v string)`

SetTDate sets TDate field to given value.


### GetTTime

`func (o *InventoryDetail) GetTTime() string`

GetTTime returns the TTime field if non-nil, zero value otherwise.

### GetTTimeOk

`func (o *InventoryDetail) GetTTimeOk() (*string, bool)`

GetTTimeOk returns a tuple with the TTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTime

`func (o *InventoryDetail) SetTTime(v string)`

SetTTime sets TTime field to given value.

### HasTTime

`func (o *InventoryDetail) HasTTime() bool`

HasTTime returns a boolean if a field has been set.

### SetTTimeNil

`func (o *InventoryDetail) SetTTimeNil(b bool)`

 SetTTimeNil sets the value for TTime to be an explicit nil

### UnsetTTime
`func (o *InventoryDetail) UnsetTTime()`

UnsetTTime ensures that no value is present for TTime, not even an explicit nil
### GetTax

`func (o *InventoryDetail) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *InventoryDetail) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *InventoryDetail) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *InventoryDetail) HasTax() bool`

HasTax returns a boolean if a field has been set.

### SetTaxNil

`func (o *InventoryDetail) SetTaxNil(b bool)`

 SetTaxNil sets the value for Tax to be an explicit nil

### UnsetTax
`func (o *InventoryDetail) UnsetTax()`

UnsetTax ensures that no value is present for Tax, not even an explicit nil
### GetTaxG

`func (o *InventoryDetail) GetTaxG() float32`

GetTaxG returns the TaxG field if non-nil, zero value otherwise.

### GetTaxGOk

`func (o *InventoryDetail) GetTaxGOk() (*float32, bool)`

GetTaxGOk returns a tuple with the TaxG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxG

`func (o *InventoryDetail) SetTaxG(v float32)`

SetTaxG sets TaxG field to given value.

### HasTaxG

`func (o *InventoryDetail) HasTaxG() bool`

HasTaxG returns a boolean if a field has been set.

### SetTaxGNil

`func (o *InventoryDetail) SetTaxGNil(b bool)`

 SetTaxGNil sets the value for TaxG to be an explicit nil

### UnsetTaxG
`func (o *InventoryDetail) UnsetTaxG()`

UnsetTaxG ensures that no value is present for TaxG, not even an explicit nil
### GetTrade

`func (o *InventoryDetail) GetTrade() int32`

GetTrade returns the Trade field if non-nil, zero value otherwise.

### GetTradeOk

`func (o *InventoryDetail) GetTradeOk() (*int32, bool)`

GetTradeOk returns a tuple with the Trade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrade

`func (o *InventoryDetail) SetTrade(v int32)`

SetTrade sets Trade field to given value.

### HasTrade

`func (o *InventoryDetail) HasTrade() bool`

HasTrade returns a boolean if a field has been set.

### SetTradeNil

`func (o *InventoryDetail) SetTradeNil(b bool)`

 SetTradeNil sets the value for Trade to be an explicit nil

### UnsetTrade
`func (o *InventoryDetail) UnsetTrade()`

UnsetTrade ensures that no value is present for Trade, not even an explicit nil
### GetValueMkt

`func (o *InventoryDetail) GetValueMkt() float32`

GetValueMkt returns the ValueMkt field if non-nil, zero value otherwise.

### GetValueMktOk

`func (o *InventoryDetail) GetValueMktOk() (*float32, bool)`

GetValueMktOk returns a tuple with the ValueMkt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMkt

`func (o *InventoryDetail) SetValueMkt(v float32)`

SetValueMkt sets ValueMkt field to given value.


### GetValueNow

`func (o *InventoryDetail) GetValueNow() float32`

GetValueNow returns the ValueNow field if non-nil, zero value otherwise.

### GetValueNowOk

`func (o *InventoryDetail) GetValueNowOk() (*float32, bool)`

GetValueNowOk returns a tuple with the ValueNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueNow

`func (o *InventoryDetail) SetValueNow(v float32)`

SetValueNow sets ValueNow field to given value.


### GetUserDef

`func (o *InventoryDetail) GetUserDef() string`

GetUserDef returns the UserDef field if non-nil, zero value otherwise.

### GetUserDefOk

`func (o *InventoryDetail) GetUserDefOk() (*string, bool)`

GetUserDefOk returns a tuple with the UserDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDef

`func (o *InventoryDetail) SetUserDef(v string)`

SetUserDef sets UserDef field to given value.

### HasUserDef

`func (o *InventoryDetail) HasUserDef() bool`

HasUserDef returns a boolean if a field has been set.

### SetUserDefNil

`func (o *InventoryDetail) SetUserDefNil(b bool)`

 SetUserDefNil sets the value for UserDef to be an explicit nil

### UnsetUserDef
`func (o *InventoryDetail) UnsetUserDef()`

UnsetUserDef ensures that no value is present for UserDef, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


