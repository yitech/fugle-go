# CreateOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuySell** | [**Action**](Action.md) |  | 
**ApCode** | [**APCode**](APCode.md) |  | 
**PriceFlag** | [**PriceFlag**](PriceFlag.md) |  | 
**BsFlag** | [**BSFlag**](BSFlag.md) |  | 
**Trade** | [**Trade**](Trade.md) |  | 
**StockNo** | **string** |  | 
**Quantity** | **int32** |  | 
**Price** | **float32** |  | 

## Methods

### NewCreateOrder

`func NewCreateOrder(buySell Action, apCode APCode, priceFlag PriceFlag, bsFlag BSFlag, trade Trade, stockNo string, quantity int32, price float32, ) *CreateOrder`

NewCreateOrder instantiates a new CreateOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderWithDefaults

`func NewCreateOrderWithDefaults() *CreateOrder`

NewCreateOrderWithDefaults instantiates a new CreateOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuySell

`func (o *CreateOrder) GetBuySell() Action`

GetBuySell returns the BuySell field if non-nil, zero value otherwise.

### GetBuySellOk

`func (o *CreateOrder) GetBuySellOk() (*Action, bool)`

GetBuySellOk returns a tuple with the BuySell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuySell

`func (o *CreateOrder) SetBuySell(v Action)`

SetBuySell sets BuySell field to given value.


### GetApCode

`func (o *CreateOrder) GetApCode() APCode`

GetApCode returns the ApCode field if non-nil, zero value otherwise.

### GetApCodeOk

`func (o *CreateOrder) GetApCodeOk() (*APCode, bool)`

GetApCodeOk returns a tuple with the ApCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApCode

`func (o *CreateOrder) SetApCode(v APCode)`

SetApCode sets ApCode field to given value.


### GetPriceFlag

`func (o *CreateOrder) GetPriceFlag() PriceFlag`

GetPriceFlag returns the PriceFlag field if non-nil, zero value otherwise.

### GetPriceFlagOk

`func (o *CreateOrder) GetPriceFlagOk() (*PriceFlag, bool)`

GetPriceFlagOk returns a tuple with the PriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFlag

`func (o *CreateOrder) SetPriceFlag(v PriceFlag)`

SetPriceFlag sets PriceFlag field to given value.


### GetBsFlag

`func (o *CreateOrder) GetBsFlag() BSFlag`

GetBsFlag returns the BsFlag field if non-nil, zero value otherwise.

### GetBsFlagOk

`func (o *CreateOrder) GetBsFlagOk() (*BSFlag, bool)`

GetBsFlagOk returns a tuple with the BsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsFlag

`func (o *CreateOrder) SetBsFlag(v BSFlag)`

SetBsFlag sets BsFlag field to given value.


### GetTrade

`func (o *CreateOrder) GetTrade() Trade`

GetTrade returns the Trade field if non-nil, zero value otherwise.

### GetTradeOk

`func (o *CreateOrder) GetTradeOk() (*Trade, bool)`

GetTradeOk returns a tuple with the Trade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrade

`func (o *CreateOrder) SetTrade(v Trade)`

SetTrade sets Trade field to given value.


### GetStockNo

`func (o *CreateOrder) GetStockNo() string`

GetStockNo returns the StockNo field if non-nil, zero value otherwise.

### GetStockNoOk

`func (o *CreateOrder) GetStockNoOk() (*string, bool)`

GetStockNoOk returns a tuple with the StockNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockNo

`func (o *CreateOrder) SetStockNo(v string)`

SetStockNo sets StockNo field to given value.


### GetQuantity

`func (o *CreateOrder) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateOrder) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateOrder) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *CreateOrder) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateOrder) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateOrder) SetPrice(v float32)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


