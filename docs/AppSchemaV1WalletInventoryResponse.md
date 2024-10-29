# AppSchemaV1WalletInventoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApCode** | Pointer to **string** |  | [optional] 
**CostQty** | **float32** |  | 
**CostSum** | **float32** |  | 
**MakeAPer** | **float32** |  | 
**MakeASum** | **float32** |  | 
**PriceAvg** | **float32** |  | 
**PriceEvn** | **float32** |  | 
**PriceMkt** | **float32** |  | 
**PriceNow** | **float32** |  | 
**PriceQtySum** | **float32** |  | 
**QtyB** | **int32** |  | 
**QtyBm** | **int32** |  | 
**QtyC** | **int32** |  | 
**QtyL** | **int32** |  | 
**QtyS** | **int32** |  | 
**QtySm** | **int32** |  | 
**RecVaSum** | **float32** |  | 
**SType** | **string** |  | 
**StkDats** | [**[]InventoryDetail**](InventoryDetail.md) |  | 
**StkNa** | **string** |  | 
**StkNo** | **string** |  | 
**Trade** | Pointer to **int32** |  | [optional] 
**ValueMkt** | **float32** |  | 
**ValueNow** | **float32** |  | 

## Methods

### NewAppSchemaV1WalletInventoryResponse

`func NewAppSchemaV1WalletInventoryResponse(costQty float32, costSum float32, makeAPer float32, makeASum float32, priceAvg float32, priceEvn float32, priceMkt float32, priceNow float32, priceQtySum float32, qtyB int32, qtyBm int32, qtyC int32, qtyL int32, qtyS int32, qtySm int32, recVaSum float32, sType string, stkDats []InventoryDetail, stkNa string, stkNo string, valueMkt float32, valueNow float32, ) *AppSchemaV1WalletInventoryResponse`

NewAppSchemaV1WalletInventoryResponse instantiates a new AppSchemaV1WalletInventoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSchemaV1WalletInventoryResponseWithDefaults

`func NewAppSchemaV1WalletInventoryResponseWithDefaults() *AppSchemaV1WalletInventoryResponse`

NewAppSchemaV1WalletInventoryResponseWithDefaults instantiates a new AppSchemaV1WalletInventoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApCode

`func (o *AppSchemaV1WalletInventoryResponse) GetApCode() string`

GetApCode returns the ApCode field if non-nil, zero value otherwise.

### GetApCodeOk

`func (o *AppSchemaV1WalletInventoryResponse) GetApCodeOk() (*string, bool)`

GetApCodeOk returns a tuple with the ApCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApCode

`func (o *AppSchemaV1WalletInventoryResponse) SetApCode(v string)`

SetApCode sets ApCode field to given value.

### HasApCode

`func (o *AppSchemaV1WalletInventoryResponse) HasApCode() bool`

HasApCode returns a boolean if a field has been set.

### GetCostQty

`func (o *AppSchemaV1WalletInventoryResponse) GetCostQty() float32`

GetCostQty returns the CostQty field if non-nil, zero value otherwise.

### GetCostQtyOk

`func (o *AppSchemaV1WalletInventoryResponse) GetCostQtyOk() (*float32, bool)`

GetCostQtyOk returns a tuple with the CostQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostQty

`func (o *AppSchemaV1WalletInventoryResponse) SetCostQty(v float32)`

SetCostQty sets CostQty field to given value.


### GetCostSum

`func (o *AppSchemaV1WalletInventoryResponse) GetCostSum() float32`

GetCostSum returns the CostSum field if non-nil, zero value otherwise.

### GetCostSumOk

`func (o *AppSchemaV1WalletInventoryResponse) GetCostSumOk() (*float32, bool)`

GetCostSumOk returns a tuple with the CostSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostSum

`func (o *AppSchemaV1WalletInventoryResponse) SetCostSum(v float32)`

SetCostSum sets CostSum field to given value.


### GetMakeAPer

`func (o *AppSchemaV1WalletInventoryResponse) GetMakeAPer() float32`

GetMakeAPer returns the MakeAPer field if non-nil, zero value otherwise.

### GetMakeAPerOk

`func (o *AppSchemaV1WalletInventoryResponse) GetMakeAPerOk() (*float32, bool)`

GetMakeAPerOk returns a tuple with the MakeAPer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeAPer

`func (o *AppSchemaV1WalletInventoryResponse) SetMakeAPer(v float32)`

SetMakeAPer sets MakeAPer field to given value.


### GetMakeASum

`func (o *AppSchemaV1WalletInventoryResponse) GetMakeASum() float32`

GetMakeASum returns the MakeASum field if non-nil, zero value otherwise.

### GetMakeASumOk

`func (o *AppSchemaV1WalletInventoryResponse) GetMakeASumOk() (*float32, bool)`

GetMakeASumOk returns a tuple with the MakeASum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeASum

`func (o *AppSchemaV1WalletInventoryResponse) SetMakeASum(v float32)`

SetMakeASum sets MakeASum field to given value.


### GetPriceAvg

`func (o *AppSchemaV1WalletInventoryResponse) GetPriceAvg() float32`

GetPriceAvg returns the PriceAvg field if non-nil, zero value otherwise.

### GetPriceAvgOk

`func (o *AppSchemaV1WalletInventoryResponse) GetPriceAvgOk() (*float32, bool)`

GetPriceAvgOk returns a tuple with the PriceAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAvg

`func (o *AppSchemaV1WalletInventoryResponse) SetPriceAvg(v float32)`

SetPriceAvg sets PriceAvg field to given value.


### GetPriceEvn

`func (o *AppSchemaV1WalletInventoryResponse) GetPriceEvn() float32`

GetPriceEvn returns the PriceEvn field if non-nil, zero value otherwise.

### GetPriceEvnOk

`func (o *AppSchemaV1WalletInventoryResponse) GetPriceEvnOk() (*float32, bool)`

GetPriceEvnOk returns a tuple with the PriceEvn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceEvn

`func (o *AppSchemaV1WalletInventoryResponse) SetPriceEvn(v float32)`

SetPriceEvn sets PriceEvn field to given value.


### GetPriceMkt

`func (o *AppSchemaV1WalletInventoryResponse) GetPriceMkt() float32`

GetPriceMkt returns the PriceMkt field if non-nil, zero value otherwise.

### GetPriceMktOk

`func (o *AppSchemaV1WalletInventoryResponse) GetPriceMktOk() (*float32, bool)`

GetPriceMktOk returns a tuple with the PriceMkt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMkt

`func (o *AppSchemaV1WalletInventoryResponse) SetPriceMkt(v float32)`

SetPriceMkt sets PriceMkt field to given value.


### GetPriceNow

`func (o *AppSchemaV1WalletInventoryResponse) GetPriceNow() float32`

GetPriceNow returns the PriceNow field if non-nil, zero value otherwise.

### GetPriceNowOk

`func (o *AppSchemaV1WalletInventoryResponse) GetPriceNowOk() (*float32, bool)`

GetPriceNowOk returns a tuple with the PriceNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNow

`func (o *AppSchemaV1WalletInventoryResponse) SetPriceNow(v float32)`

SetPriceNow sets PriceNow field to given value.


### GetPriceQtySum

`func (o *AppSchemaV1WalletInventoryResponse) GetPriceQtySum() float32`

GetPriceQtySum returns the PriceQtySum field if non-nil, zero value otherwise.

### GetPriceQtySumOk

`func (o *AppSchemaV1WalletInventoryResponse) GetPriceQtySumOk() (*float32, bool)`

GetPriceQtySumOk returns a tuple with the PriceQtySum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceQtySum

`func (o *AppSchemaV1WalletInventoryResponse) SetPriceQtySum(v float32)`

SetPriceQtySum sets PriceQtySum field to given value.


### GetQtyB

`func (o *AppSchemaV1WalletInventoryResponse) GetQtyB() int32`

GetQtyB returns the QtyB field if non-nil, zero value otherwise.

### GetQtyBOk

`func (o *AppSchemaV1WalletInventoryResponse) GetQtyBOk() (*int32, bool)`

GetQtyBOk returns a tuple with the QtyB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyB

`func (o *AppSchemaV1WalletInventoryResponse) SetQtyB(v int32)`

SetQtyB sets QtyB field to given value.


### GetQtyBm

`func (o *AppSchemaV1WalletInventoryResponse) GetQtyBm() int32`

GetQtyBm returns the QtyBm field if non-nil, zero value otherwise.

### GetQtyBmOk

`func (o *AppSchemaV1WalletInventoryResponse) GetQtyBmOk() (*int32, bool)`

GetQtyBmOk returns a tuple with the QtyBm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyBm

`func (o *AppSchemaV1WalletInventoryResponse) SetQtyBm(v int32)`

SetQtyBm sets QtyBm field to given value.


### GetQtyC

`func (o *AppSchemaV1WalletInventoryResponse) GetQtyC() int32`

GetQtyC returns the QtyC field if non-nil, zero value otherwise.

### GetQtyCOk

`func (o *AppSchemaV1WalletInventoryResponse) GetQtyCOk() (*int32, bool)`

GetQtyCOk returns a tuple with the QtyC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyC

`func (o *AppSchemaV1WalletInventoryResponse) SetQtyC(v int32)`

SetQtyC sets QtyC field to given value.


### GetQtyL

`func (o *AppSchemaV1WalletInventoryResponse) GetQtyL() int32`

GetQtyL returns the QtyL field if non-nil, zero value otherwise.

### GetQtyLOk

`func (o *AppSchemaV1WalletInventoryResponse) GetQtyLOk() (*int32, bool)`

GetQtyLOk returns a tuple with the QtyL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyL

`func (o *AppSchemaV1WalletInventoryResponse) SetQtyL(v int32)`

SetQtyL sets QtyL field to given value.


### GetQtyS

`func (o *AppSchemaV1WalletInventoryResponse) GetQtyS() int32`

GetQtyS returns the QtyS field if non-nil, zero value otherwise.

### GetQtySOk

`func (o *AppSchemaV1WalletInventoryResponse) GetQtySOk() (*int32, bool)`

GetQtySOk returns a tuple with the QtyS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyS

`func (o *AppSchemaV1WalletInventoryResponse) SetQtyS(v int32)`

SetQtyS sets QtyS field to given value.


### GetQtySm

`func (o *AppSchemaV1WalletInventoryResponse) GetQtySm() int32`

GetQtySm returns the QtySm field if non-nil, zero value otherwise.

### GetQtySmOk

`func (o *AppSchemaV1WalletInventoryResponse) GetQtySmOk() (*int32, bool)`

GetQtySmOk returns a tuple with the QtySm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtySm

`func (o *AppSchemaV1WalletInventoryResponse) SetQtySm(v int32)`

SetQtySm sets QtySm field to given value.


### GetRecVaSum

`func (o *AppSchemaV1WalletInventoryResponse) GetRecVaSum() float32`

GetRecVaSum returns the RecVaSum field if non-nil, zero value otherwise.

### GetRecVaSumOk

`func (o *AppSchemaV1WalletInventoryResponse) GetRecVaSumOk() (*float32, bool)`

GetRecVaSumOk returns a tuple with the RecVaSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecVaSum

`func (o *AppSchemaV1WalletInventoryResponse) SetRecVaSum(v float32)`

SetRecVaSum sets RecVaSum field to given value.


### GetSType

`func (o *AppSchemaV1WalletInventoryResponse) GetSType() string`

GetSType returns the SType field if non-nil, zero value otherwise.

### GetSTypeOk

`func (o *AppSchemaV1WalletInventoryResponse) GetSTypeOk() (*string, bool)`

GetSTypeOk returns a tuple with the SType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSType

`func (o *AppSchemaV1WalletInventoryResponse) SetSType(v string)`

SetSType sets SType field to given value.


### GetStkDats

`func (o *AppSchemaV1WalletInventoryResponse) GetStkDats() []InventoryDetail`

GetStkDats returns the StkDats field if non-nil, zero value otherwise.

### GetStkDatsOk

`func (o *AppSchemaV1WalletInventoryResponse) GetStkDatsOk() (*[]InventoryDetail, bool)`

GetStkDatsOk returns a tuple with the StkDats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkDats

`func (o *AppSchemaV1WalletInventoryResponse) SetStkDats(v []InventoryDetail)`

SetStkDats sets StkDats field to given value.


### GetStkNa

`func (o *AppSchemaV1WalletInventoryResponse) GetStkNa() string`

GetStkNa returns the StkNa field if non-nil, zero value otherwise.

### GetStkNaOk

`func (o *AppSchemaV1WalletInventoryResponse) GetStkNaOk() (*string, bool)`

GetStkNaOk returns a tuple with the StkNa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkNa

`func (o *AppSchemaV1WalletInventoryResponse) SetStkNa(v string)`

SetStkNa sets StkNa field to given value.


### GetStkNo

`func (o *AppSchemaV1WalletInventoryResponse) GetStkNo() string`

GetStkNo returns the StkNo field if non-nil, zero value otherwise.

### GetStkNoOk

`func (o *AppSchemaV1WalletInventoryResponse) GetStkNoOk() (*string, bool)`

GetStkNoOk returns a tuple with the StkNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkNo

`func (o *AppSchemaV1WalletInventoryResponse) SetStkNo(v string)`

SetStkNo sets StkNo field to given value.


### GetTrade

`func (o *AppSchemaV1WalletInventoryResponse) GetTrade() int32`

GetTrade returns the Trade field if non-nil, zero value otherwise.

### GetTradeOk

`func (o *AppSchemaV1WalletInventoryResponse) GetTradeOk() (*int32, bool)`

GetTradeOk returns a tuple with the Trade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrade

`func (o *AppSchemaV1WalletInventoryResponse) SetTrade(v int32)`

SetTrade sets Trade field to given value.

### HasTrade

`func (o *AppSchemaV1WalletInventoryResponse) HasTrade() bool`

HasTrade returns a boolean if a field has been set.

### GetValueMkt

`func (o *AppSchemaV1WalletInventoryResponse) GetValueMkt() float32`

GetValueMkt returns the ValueMkt field if non-nil, zero value otherwise.

### GetValueMktOk

`func (o *AppSchemaV1WalletInventoryResponse) GetValueMktOk() (*float32, bool)`

GetValueMktOk returns a tuple with the ValueMkt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMkt

`func (o *AppSchemaV1WalletInventoryResponse) SetValueMkt(v float32)`

SetValueMkt sets ValueMkt field to given value.


### GetValueNow

`func (o *AppSchemaV1WalletInventoryResponse) GetValueNow() float32`

GetValueNow returns the ValueNow field if non-nil, zero value otherwise.

### GetValueNowOk

`func (o *AppSchemaV1WalletInventoryResponse) GetValueNowOk() (*float32, bool)`

GetValueNowOk returns a tuple with the ValueNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueNow

`func (o *AppSchemaV1WalletInventoryResponse) SetValueNow(v float32)`

SetValueNow sets ValueNow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


