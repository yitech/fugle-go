# OrderResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApCode** | [**APCode**](APCode.md) |  | 
**AvgPrice** | **float32** |  | 
**BsFlag** | [**BSFlag**](BSFlag.md) |  | 
**BuySell** | [**Action**](Action.md) |  | 
**CelQty** | **float32** |  | 
**CelQtyShare** | **int32** |  | 
**Celable** | **string** |  | 
**ErrCode** | **string** |  | 
**ErrMsg** | **string** |  | 
**MatQty** | **float32** |  | 
**MatQtyShare** | **int32** |  | 
**OdPrice** | **float32** |  | 
**OrdDate** | **string** |  | 
**OrdNo** | **string** |  | 
**OrdStatus** | **string** |  | 
**OrdTime** | **string** |  | 
**OrgQty** | **float32** |  | 
**OrgQtyShare** | **int32** |  | 
**PreOrdNo** | **string** |  | 
**PriceFlag** | [**PriceFlag**](PriceFlag.md) |  | 
**StockNo** | **string** |  | 
**Trade** | [**Trade**](Trade.md) |  | 
**WorkDate** | **string** |  | 
**UserDef** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewOrderResultResponse

`func NewOrderResultResponse(apCode APCode, avgPrice float32, bsFlag BSFlag, buySell Action, celQty float32, celQtyShare int32, celable string, errCode string, errMsg string, matQty float32, matQtyShare int32, odPrice float32, ordDate string, ordNo string, ordStatus string, ordTime string, orgQty float32, orgQtyShare int32, preOrdNo string, priceFlag PriceFlag, stockNo string, trade Trade, workDate string, ) *OrderResultResponse`

NewOrderResultResponse instantiates a new OrderResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResultResponseWithDefaults

`func NewOrderResultResponseWithDefaults() *OrderResultResponse`

NewOrderResultResponseWithDefaults instantiates a new OrderResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApCode

`func (o *OrderResultResponse) GetApCode() APCode`

GetApCode returns the ApCode field if non-nil, zero value otherwise.

### GetApCodeOk

`func (o *OrderResultResponse) GetApCodeOk() (*APCode, bool)`

GetApCodeOk returns a tuple with the ApCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApCode

`func (o *OrderResultResponse) SetApCode(v APCode)`

SetApCode sets ApCode field to given value.


### GetAvgPrice

`func (o *OrderResultResponse) GetAvgPrice() float32`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *OrderResultResponse) GetAvgPriceOk() (*float32, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *OrderResultResponse) SetAvgPrice(v float32)`

SetAvgPrice sets AvgPrice field to given value.


### GetBsFlag

`func (o *OrderResultResponse) GetBsFlag() BSFlag`

GetBsFlag returns the BsFlag field if non-nil, zero value otherwise.

### GetBsFlagOk

`func (o *OrderResultResponse) GetBsFlagOk() (*BSFlag, bool)`

GetBsFlagOk returns a tuple with the BsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsFlag

`func (o *OrderResultResponse) SetBsFlag(v BSFlag)`

SetBsFlag sets BsFlag field to given value.


### GetBuySell

`func (o *OrderResultResponse) GetBuySell() Action`

GetBuySell returns the BuySell field if non-nil, zero value otherwise.

### GetBuySellOk

`func (o *OrderResultResponse) GetBuySellOk() (*Action, bool)`

GetBuySellOk returns a tuple with the BuySell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuySell

`func (o *OrderResultResponse) SetBuySell(v Action)`

SetBuySell sets BuySell field to given value.


### GetCelQty

`func (o *OrderResultResponse) GetCelQty() float32`

GetCelQty returns the CelQty field if non-nil, zero value otherwise.

### GetCelQtyOk

`func (o *OrderResultResponse) GetCelQtyOk() (*float32, bool)`

GetCelQtyOk returns a tuple with the CelQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCelQty

`func (o *OrderResultResponse) SetCelQty(v float32)`

SetCelQty sets CelQty field to given value.


### GetCelQtyShare

`func (o *OrderResultResponse) GetCelQtyShare() int32`

GetCelQtyShare returns the CelQtyShare field if non-nil, zero value otherwise.

### GetCelQtyShareOk

`func (o *OrderResultResponse) GetCelQtyShareOk() (*int32, bool)`

GetCelQtyShareOk returns a tuple with the CelQtyShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCelQtyShare

`func (o *OrderResultResponse) SetCelQtyShare(v int32)`

SetCelQtyShare sets CelQtyShare field to given value.


### GetCelable

`func (o *OrderResultResponse) GetCelable() string`

GetCelable returns the Celable field if non-nil, zero value otherwise.

### GetCelableOk

`func (o *OrderResultResponse) GetCelableOk() (*string, bool)`

GetCelableOk returns a tuple with the Celable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCelable

`func (o *OrderResultResponse) SetCelable(v string)`

SetCelable sets Celable field to given value.


### GetErrCode

`func (o *OrderResultResponse) GetErrCode() string`

GetErrCode returns the ErrCode field if non-nil, zero value otherwise.

### GetErrCodeOk

`func (o *OrderResultResponse) GetErrCodeOk() (*string, bool)`

GetErrCodeOk returns a tuple with the ErrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrCode

`func (o *OrderResultResponse) SetErrCode(v string)`

SetErrCode sets ErrCode field to given value.


### GetErrMsg

`func (o *OrderResultResponse) GetErrMsg() string`

GetErrMsg returns the ErrMsg field if non-nil, zero value otherwise.

### GetErrMsgOk

`func (o *OrderResultResponse) GetErrMsgOk() (*string, bool)`

GetErrMsgOk returns a tuple with the ErrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMsg

`func (o *OrderResultResponse) SetErrMsg(v string)`

SetErrMsg sets ErrMsg field to given value.


### GetMatQty

`func (o *OrderResultResponse) GetMatQty() float32`

GetMatQty returns the MatQty field if non-nil, zero value otherwise.

### GetMatQtyOk

`func (o *OrderResultResponse) GetMatQtyOk() (*float32, bool)`

GetMatQtyOk returns a tuple with the MatQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatQty

`func (o *OrderResultResponse) SetMatQty(v float32)`

SetMatQty sets MatQty field to given value.


### GetMatQtyShare

`func (o *OrderResultResponse) GetMatQtyShare() int32`

GetMatQtyShare returns the MatQtyShare field if non-nil, zero value otherwise.

### GetMatQtyShareOk

`func (o *OrderResultResponse) GetMatQtyShareOk() (*int32, bool)`

GetMatQtyShareOk returns a tuple with the MatQtyShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatQtyShare

`func (o *OrderResultResponse) SetMatQtyShare(v int32)`

SetMatQtyShare sets MatQtyShare field to given value.


### GetOdPrice

`func (o *OrderResultResponse) GetOdPrice() float32`

GetOdPrice returns the OdPrice field if non-nil, zero value otherwise.

### GetOdPriceOk

`func (o *OrderResultResponse) GetOdPriceOk() (*float32, bool)`

GetOdPriceOk returns a tuple with the OdPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdPrice

`func (o *OrderResultResponse) SetOdPrice(v float32)`

SetOdPrice sets OdPrice field to given value.


### GetOrdDate

`func (o *OrderResultResponse) GetOrdDate() string`

GetOrdDate returns the OrdDate field if non-nil, zero value otherwise.

### GetOrdDateOk

`func (o *OrderResultResponse) GetOrdDateOk() (*string, bool)`

GetOrdDateOk returns a tuple with the OrdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdDate

`func (o *OrderResultResponse) SetOrdDate(v string)`

SetOrdDate sets OrdDate field to given value.


### GetOrdNo

`func (o *OrderResultResponse) GetOrdNo() string`

GetOrdNo returns the OrdNo field if non-nil, zero value otherwise.

### GetOrdNoOk

`func (o *OrderResultResponse) GetOrdNoOk() (*string, bool)`

GetOrdNoOk returns a tuple with the OrdNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdNo

`func (o *OrderResultResponse) SetOrdNo(v string)`

SetOrdNo sets OrdNo field to given value.


### GetOrdStatus

`func (o *OrderResultResponse) GetOrdStatus() string`

GetOrdStatus returns the OrdStatus field if non-nil, zero value otherwise.

### GetOrdStatusOk

`func (o *OrderResultResponse) GetOrdStatusOk() (*string, bool)`

GetOrdStatusOk returns a tuple with the OrdStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdStatus

`func (o *OrderResultResponse) SetOrdStatus(v string)`

SetOrdStatus sets OrdStatus field to given value.


### GetOrdTime

`func (o *OrderResultResponse) GetOrdTime() string`

GetOrdTime returns the OrdTime field if non-nil, zero value otherwise.

### GetOrdTimeOk

`func (o *OrderResultResponse) GetOrdTimeOk() (*string, bool)`

GetOrdTimeOk returns a tuple with the OrdTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdTime

`func (o *OrderResultResponse) SetOrdTime(v string)`

SetOrdTime sets OrdTime field to given value.


### GetOrgQty

`func (o *OrderResultResponse) GetOrgQty() float32`

GetOrgQty returns the OrgQty field if non-nil, zero value otherwise.

### GetOrgQtyOk

`func (o *OrderResultResponse) GetOrgQtyOk() (*float32, bool)`

GetOrgQtyOk returns a tuple with the OrgQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgQty

`func (o *OrderResultResponse) SetOrgQty(v float32)`

SetOrgQty sets OrgQty field to given value.


### GetOrgQtyShare

`func (o *OrderResultResponse) GetOrgQtyShare() int32`

GetOrgQtyShare returns the OrgQtyShare field if non-nil, zero value otherwise.

### GetOrgQtyShareOk

`func (o *OrderResultResponse) GetOrgQtyShareOk() (*int32, bool)`

GetOrgQtyShareOk returns a tuple with the OrgQtyShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgQtyShare

`func (o *OrderResultResponse) SetOrgQtyShare(v int32)`

SetOrgQtyShare sets OrgQtyShare field to given value.


### GetPreOrdNo

`func (o *OrderResultResponse) GetPreOrdNo() string`

GetPreOrdNo returns the PreOrdNo field if non-nil, zero value otherwise.

### GetPreOrdNoOk

`func (o *OrderResultResponse) GetPreOrdNoOk() (*string, bool)`

GetPreOrdNoOk returns a tuple with the PreOrdNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOrdNo

`func (o *OrderResultResponse) SetPreOrdNo(v string)`

SetPreOrdNo sets PreOrdNo field to given value.


### GetPriceFlag

`func (o *OrderResultResponse) GetPriceFlag() PriceFlag`

GetPriceFlag returns the PriceFlag field if non-nil, zero value otherwise.

### GetPriceFlagOk

`func (o *OrderResultResponse) GetPriceFlagOk() (*PriceFlag, bool)`

GetPriceFlagOk returns a tuple with the PriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFlag

`func (o *OrderResultResponse) SetPriceFlag(v PriceFlag)`

SetPriceFlag sets PriceFlag field to given value.


### GetStockNo

`func (o *OrderResultResponse) GetStockNo() string`

GetStockNo returns the StockNo field if non-nil, zero value otherwise.

### GetStockNoOk

`func (o *OrderResultResponse) GetStockNoOk() (*string, bool)`

GetStockNoOk returns a tuple with the StockNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockNo

`func (o *OrderResultResponse) SetStockNo(v string)`

SetStockNo sets StockNo field to given value.


### GetTrade

`func (o *OrderResultResponse) GetTrade() Trade`

GetTrade returns the Trade field if non-nil, zero value otherwise.

### GetTradeOk

`func (o *OrderResultResponse) GetTradeOk() (*Trade, bool)`

GetTradeOk returns a tuple with the Trade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrade

`func (o *OrderResultResponse) SetTrade(v Trade)`

SetTrade sets Trade field to given value.


### GetWorkDate

`func (o *OrderResultResponse) GetWorkDate() string`

GetWorkDate returns the WorkDate field if non-nil, zero value otherwise.

### GetWorkDateOk

`func (o *OrderResultResponse) GetWorkDateOk() (*string, bool)`

GetWorkDateOk returns a tuple with the WorkDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkDate

`func (o *OrderResultResponse) SetWorkDate(v string)`

SetWorkDate sets WorkDate field to given value.


### GetUserDef

`func (o *OrderResultResponse) GetUserDef() string`

GetUserDef returns the UserDef field if non-nil, zero value otherwise.

### GetUserDefOk

`func (o *OrderResultResponse) GetUserDefOk() (*string, bool)`

GetUserDefOk returns a tuple with the UserDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDef

`func (o *OrderResultResponse) SetUserDef(v string)`

SetUserDef sets UserDef field to given value.

### HasUserDef

`func (o *OrderResultResponse) HasUserDef() bool`

HasUserDef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


