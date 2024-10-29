# AppSchemaV1OrderTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuySell** | **string** |  | 
**CDate** | **string** |  | 
**Cost** | **int32** |  | 
**Make** | **int32** |  | 
**MakePer** | **float32** |  | 
**MatDats** | [**[]MatDetail**](MatDetail.md) |  | 
**PriceAvg** | **string** |  | 
**PriceQty** | **string** |  | 
**Qty** | **string** |  | 
**Recv** | **int32** |  | 
**SType** | **string** |  | 
**StkNa** | **string** |  | 
**StkNo** | **string** |  | 
**TDate** | **string** |  | 
**Trade** | **int32** |  | 

## Methods

### NewAppSchemaV1OrderTransactionResponse

`func NewAppSchemaV1OrderTransactionResponse(buySell string, cDate string, cost int32, make int32, makePer float32, matDats []MatDetail, priceAvg string, priceQty string, qty string, recv int32, sType string, stkNa string, stkNo string, tDate string, trade int32, ) *AppSchemaV1OrderTransactionResponse`

NewAppSchemaV1OrderTransactionResponse instantiates a new AppSchemaV1OrderTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSchemaV1OrderTransactionResponseWithDefaults

`func NewAppSchemaV1OrderTransactionResponseWithDefaults() *AppSchemaV1OrderTransactionResponse`

NewAppSchemaV1OrderTransactionResponseWithDefaults instantiates a new AppSchemaV1OrderTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuySell

`func (o *AppSchemaV1OrderTransactionResponse) GetBuySell() string`

GetBuySell returns the BuySell field if non-nil, zero value otherwise.

### GetBuySellOk

`func (o *AppSchemaV1OrderTransactionResponse) GetBuySellOk() (*string, bool)`

GetBuySellOk returns a tuple with the BuySell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuySell

`func (o *AppSchemaV1OrderTransactionResponse) SetBuySell(v string)`

SetBuySell sets BuySell field to given value.


### GetCDate

`func (o *AppSchemaV1OrderTransactionResponse) GetCDate() string`

GetCDate returns the CDate field if non-nil, zero value otherwise.

### GetCDateOk

`func (o *AppSchemaV1OrderTransactionResponse) GetCDateOk() (*string, bool)`

GetCDateOk returns a tuple with the CDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDate

`func (o *AppSchemaV1OrderTransactionResponse) SetCDate(v string)`

SetCDate sets CDate field to given value.


### GetCost

`func (o *AppSchemaV1OrderTransactionResponse) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AppSchemaV1OrderTransactionResponse) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AppSchemaV1OrderTransactionResponse) SetCost(v int32)`

SetCost sets Cost field to given value.


### GetMake

`func (o *AppSchemaV1OrderTransactionResponse) GetMake() int32`

GetMake returns the Make field if non-nil, zero value otherwise.

### GetMakeOk

`func (o *AppSchemaV1OrderTransactionResponse) GetMakeOk() (*int32, bool)`

GetMakeOk returns a tuple with the Make field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMake

`func (o *AppSchemaV1OrderTransactionResponse) SetMake(v int32)`

SetMake sets Make field to given value.


### GetMakePer

`func (o *AppSchemaV1OrderTransactionResponse) GetMakePer() float32`

GetMakePer returns the MakePer field if non-nil, zero value otherwise.

### GetMakePerOk

`func (o *AppSchemaV1OrderTransactionResponse) GetMakePerOk() (*float32, bool)`

GetMakePerOk returns a tuple with the MakePer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakePer

`func (o *AppSchemaV1OrderTransactionResponse) SetMakePer(v float32)`

SetMakePer sets MakePer field to given value.


### GetMatDats

`func (o *AppSchemaV1OrderTransactionResponse) GetMatDats() []MatDetail`

GetMatDats returns the MatDats field if non-nil, zero value otherwise.

### GetMatDatsOk

`func (o *AppSchemaV1OrderTransactionResponse) GetMatDatsOk() (*[]MatDetail, bool)`

GetMatDatsOk returns a tuple with the MatDats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatDats

`func (o *AppSchemaV1OrderTransactionResponse) SetMatDats(v []MatDetail)`

SetMatDats sets MatDats field to given value.


### GetPriceAvg

`func (o *AppSchemaV1OrderTransactionResponse) GetPriceAvg() string`

GetPriceAvg returns the PriceAvg field if non-nil, zero value otherwise.

### GetPriceAvgOk

`func (o *AppSchemaV1OrderTransactionResponse) GetPriceAvgOk() (*string, bool)`

GetPriceAvgOk returns a tuple with the PriceAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAvg

`func (o *AppSchemaV1OrderTransactionResponse) SetPriceAvg(v string)`

SetPriceAvg sets PriceAvg field to given value.


### GetPriceQty

`func (o *AppSchemaV1OrderTransactionResponse) GetPriceQty() string`

GetPriceQty returns the PriceQty field if non-nil, zero value otherwise.

### GetPriceQtyOk

`func (o *AppSchemaV1OrderTransactionResponse) GetPriceQtyOk() (*string, bool)`

GetPriceQtyOk returns a tuple with the PriceQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceQty

`func (o *AppSchemaV1OrderTransactionResponse) SetPriceQty(v string)`

SetPriceQty sets PriceQty field to given value.


### GetQty

`func (o *AppSchemaV1OrderTransactionResponse) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *AppSchemaV1OrderTransactionResponse) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *AppSchemaV1OrderTransactionResponse) SetQty(v string)`

SetQty sets Qty field to given value.


### GetRecv

`func (o *AppSchemaV1OrderTransactionResponse) GetRecv() int32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *AppSchemaV1OrderTransactionResponse) GetRecvOk() (*int32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *AppSchemaV1OrderTransactionResponse) SetRecv(v int32)`

SetRecv sets Recv field to given value.


### GetSType

`func (o *AppSchemaV1OrderTransactionResponse) GetSType() string`

GetSType returns the SType field if non-nil, zero value otherwise.

### GetSTypeOk

`func (o *AppSchemaV1OrderTransactionResponse) GetSTypeOk() (*string, bool)`

GetSTypeOk returns a tuple with the SType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSType

`func (o *AppSchemaV1OrderTransactionResponse) SetSType(v string)`

SetSType sets SType field to given value.


### GetStkNa

`func (o *AppSchemaV1OrderTransactionResponse) GetStkNa() string`

GetStkNa returns the StkNa field if non-nil, zero value otherwise.

### GetStkNaOk

`func (o *AppSchemaV1OrderTransactionResponse) GetStkNaOk() (*string, bool)`

GetStkNaOk returns a tuple with the StkNa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkNa

`func (o *AppSchemaV1OrderTransactionResponse) SetStkNa(v string)`

SetStkNa sets StkNa field to given value.


### GetStkNo

`func (o *AppSchemaV1OrderTransactionResponse) GetStkNo() string`

GetStkNo returns the StkNo field if non-nil, zero value otherwise.

### GetStkNoOk

`func (o *AppSchemaV1OrderTransactionResponse) GetStkNoOk() (*string, bool)`

GetStkNoOk returns a tuple with the StkNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkNo

`func (o *AppSchemaV1OrderTransactionResponse) SetStkNo(v string)`

SetStkNo sets StkNo field to given value.


### GetTDate

`func (o *AppSchemaV1OrderTransactionResponse) GetTDate() string`

GetTDate returns the TDate field if non-nil, zero value otherwise.

### GetTDateOk

`func (o *AppSchemaV1OrderTransactionResponse) GetTDateOk() (*string, bool)`

GetTDateOk returns a tuple with the TDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTDate

`func (o *AppSchemaV1OrderTransactionResponse) SetTDate(v string)`

SetTDate sets TDate field to given value.


### GetTrade

`func (o *AppSchemaV1OrderTransactionResponse) GetTrade() int32`

GetTrade returns the Trade field if non-nil, zero value otherwise.

### GetTradeOk

`func (o *AppSchemaV1OrderTransactionResponse) GetTradeOk() (*int32, bool)`

GetTradeOk returns a tuple with the Trade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrade

`func (o *AppSchemaV1OrderTransactionResponse) SetTrade(v int32)`

SetTrade sets Trade field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


