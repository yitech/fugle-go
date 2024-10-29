# AppSchemaV2TradeTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Side** | **string** |  | 
**Datetime** | **time.Time** |  | 
**Fee** | **int32** |  | 
**OrderNo** | **string** |  | 
**Price** | **float32** |  | 
**Qty** | **int32** |  | 
**StkNo** | **string** |  | 
**Tax** | **int32** |  | 
**Trade** | **int32** |  | 

## Methods

### NewAppSchemaV2TradeTransactionResponse

`func NewAppSchemaV2TradeTransactionResponse(side string, datetime time.Time, fee int32, orderNo string, price float32, qty int32, stkNo string, tax int32, trade int32, ) *AppSchemaV2TradeTransactionResponse`

NewAppSchemaV2TradeTransactionResponse instantiates a new AppSchemaV2TradeTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSchemaV2TradeTransactionResponseWithDefaults

`func NewAppSchemaV2TradeTransactionResponseWithDefaults() *AppSchemaV2TradeTransactionResponse`

NewAppSchemaV2TradeTransactionResponseWithDefaults instantiates a new AppSchemaV2TradeTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSide

`func (o *AppSchemaV2TradeTransactionResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *AppSchemaV2TradeTransactionResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *AppSchemaV2TradeTransactionResponse) SetSide(v string)`

SetSide sets Side field to given value.


### GetDatetime

`func (o *AppSchemaV2TradeTransactionResponse) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *AppSchemaV2TradeTransactionResponse) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *AppSchemaV2TradeTransactionResponse) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetFee

`func (o *AppSchemaV2TradeTransactionResponse) GetFee() int32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *AppSchemaV2TradeTransactionResponse) GetFeeOk() (*int32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *AppSchemaV2TradeTransactionResponse) SetFee(v int32)`

SetFee sets Fee field to given value.


### GetOrderNo

`func (o *AppSchemaV2TradeTransactionResponse) GetOrderNo() string`

GetOrderNo returns the OrderNo field if non-nil, zero value otherwise.

### GetOrderNoOk

`func (o *AppSchemaV2TradeTransactionResponse) GetOrderNoOk() (*string, bool)`

GetOrderNoOk returns a tuple with the OrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNo

`func (o *AppSchemaV2TradeTransactionResponse) SetOrderNo(v string)`

SetOrderNo sets OrderNo field to given value.


### GetPrice

`func (o *AppSchemaV2TradeTransactionResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AppSchemaV2TradeTransactionResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AppSchemaV2TradeTransactionResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetQty

`func (o *AppSchemaV2TradeTransactionResponse) GetQty() int32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *AppSchemaV2TradeTransactionResponse) GetQtyOk() (*int32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *AppSchemaV2TradeTransactionResponse) SetQty(v int32)`

SetQty sets Qty field to given value.


### GetStkNo

`func (o *AppSchemaV2TradeTransactionResponse) GetStkNo() string`

GetStkNo returns the StkNo field if non-nil, zero value otherwise.

### GetStkNoOk

`func (o *AppSchemaV2TradeTransactionResponse) GetStkNoOk() (*string, bool)`

GetStkNoOk returns a tuple with the StkNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkNo

`func (o *AppSchemaV2TradeTransactionResponse) SetStkNo(v string)`

SetStkNo sets StkNo field to given value.


### GetTax

`func (o *AppSchemaV2TradeTransactionResponse) GetTax() int32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *AppSchemaV2TradeTransactionResponse) GetTaxOk() (*int32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *AppSchemaV2TradeTransactionResponse) SetTax(v int32)`

SetTax sets Tax field to given value.


### GetTrade

`func (o *AppSchemaV2TradeTransactionResponse) GetTrade() int32`

GetTrade returns the Trade field if non-nil, zero value otherwise.

### GetTradeOk

`func (o *AppSchemaV2TradeTransactionResponse) GetTradeOk() (*int32, bool)`

GetTradeOk returns a tuple with the Trade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrade

`func (o *AppSchemaV2TradeTransactionResponse) SetTrade(v int32)`

SetTrade sets Trade field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


