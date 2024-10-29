# MatDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuySell** | **string** |  | 
**CDate** | **string** |  | 
**DbFee** | Pointer to **string** |  | [optional] [default to "0"]
**Fee** | **string** |  | 
**Make** | **int32** |  | 
**MakePer** | **float32** |  | 
**OrderNo** | **string** |  | 
**PayN** | **string** |  | 
**Price** | **string** |  | 
**PriceQty** | **string** |  | 
**Qty** | **string** |  | 
**SType** | **string** |  | 
**StkNa** | **string** |  | 
**StkNo** | **string** |  | 
**TDate** | **string** |  | 
**TTime** | Pointer to **string** |  | [optional] [default to ""]
**Tax** | **string** |  | 
**TaxG** | **string** |  | 
**Trade** | **int32** |  | 
**UserDef** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewMatDetail

`func NewMatDetail(buySell string, cDate string, fee string, make int32, makePer float32, orderNo string, payN string, price string, priceQty string, qty string, sType string, stkNa string, stkNo string, tDate string, tax string, taxG string, trade int32, ) *MatDetail`

NewMatDetail instantiates a new MatDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatDetailWithDefaults

`func NewMatDetailWithDefaults() *MatDetail`

NewMatDetailWithDefaults instantiates a new MatDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuySell

`func (o *MatDetail) GetBuySell() string`

GetBuySell returns the BuySell field if non-nil, zero value otherwise.

### GetBuySellOk

`func (o *MatDetail) GetBuySellOk() (*string, bool)`

GetBuySellOk returns a tuple with the BuySell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuySell

`func (o *MatDetail) SetBuySell(v string)`

SetBuySell sets BuySell field to given value.


### GetCDate

`func (o *MatDetail) GetCDate() string`

GetCDate returns the CDate field if non-nil, zero value otherwise.

### GetCDateOk

`func (o *MatDetail) GetCDateOk() (*string, bool)`

GetCDateOk returns a tuple with the CDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDate

`func (o *MatDetail) SetCDate(v string)`

SetCDate sets CDate field to given value.


### GetDbFee

`func (o *MatDetail) GetDbFee() string`

GetDbFee returns the DbFee field if non-nil, zero value otherwise.

### GetDbFeeOk

`func (o *MatDetail) GetDbFeeOk() (*string, bool)`

GetDbFeeOk returns a tuple with the DbFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbFee

`func (o *MatDetail) SetDbFee(v string)`

SetDbFee sets DbFee field to given value.

### HasDbFee

`func (o *MatDetail) HasDbFee() bool`

HasDbFee returns a boolean if a field has been set.

### GetFee

`func (o *MatDetail) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *MatDetail) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *MatDetail) SetFee(v string)`

SetFee sets Fee field to given value.


### GetMake

`func (o *MatDetail) GetMake() int32`

GetMake returns the Make field if non-nil, zero value otherwise.

### GetMakeOk

`func (o *MatDetail) GetMakeOk() (*int32, bool)`

GetMakeOk returns a tuple with the Make field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMake

`func (o *MatDetail) SetMake(v int32)`

SetMake sets Make field to given value.


### GetMakePer

`func (o *MatDetail) GetMakePer() float32`

GetMakePer returns the MakePer field if non-nil, zero value otherwise.

### GetMakePerOk

`func (o *MatDetail) GetMakePerOk() (*float32, bool)`

GetMakePerOk returns a tuple with the MakePer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakePer

`func (o *MatDetail) SetMakePer(v float32)`

SetMakePer sets MakePer field to given value.


### GetOrderNo

`func (o *MatDetail) GetOrderNo() string`

GetOrderNo returns the OrderNo field if non-nil, zero value otherwise.

### GetOrderNoOk

`func (o *MatDetail) GetOrderNoOk() (*string, bool)`

GetOrderNoOk returns a tuple with the OrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNo

`func (o *MatDetail) SetOrderNo(v string)`

SetOrderNo sets OrderNo field to given value.


### GetPayN

`func (o *MatDetail) GetPayN() string`

GetPayN returns the PayN field if non-nil, zero value otherwise.

### GetPayNOk

`func (o *MatDetail) GetPayNOk() (*string, bool)`

GetPayNOk returns a tuple with the PayN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayN

`func (o *MatDetail) SetPayN(v string)`

SetPayN sets PayN field to given value.


### GetPrice

`func (o *MatDetail) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MatDetail) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MatDetail) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetPriceQty

`func (o *MatDetail) GetPriceQty() string`

GetPriceQty returns the PriceQty field if non-nil, zero value otherwise.

### GetPriceQtyOk

`func (o *MatDetail) GetPriceQtyOk() (*string, bool)`

GetPriceQtyOk returns a tuple with the PriceQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceQty

`func (o *MatDetail) SetPriceQty(v string)`

SetPriceQty sets PriceQty field to given value.


### GetQty

`func (o *MatDetail) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *MatDetail) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *MatDetail) SetQty(v string)`

SetQty sets Qty field to given value.


### GetSType

`func (o *MatDetail) GetSType() string`

GetSType returns the SType field if non-nil, zero value otherwise.

### GetSTypeOk

`func (o *MatDetail) GetSTypeOk() (*string, bool)`

GetSTypeOk returns a tuple with the SType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSType

`func (o *MatDetail) SetSType(v string)`

SetSType sets SType field to given value.


### GetStkNa

`func (o *MatDetail) GetStkNa() string`

GetStkNa returns the StkNa field if non-nil, zero value otherwise.

### GetStkNaOk

`func (o *MatDetail) GetStkNaOk() (*string, bool)`

GetStkNaOk returns a tuple with the StkNa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkNa

`func (o *MatDetail) SetStkNa(v string)`

SetStkNa sets StkNa field to given value.


### GetStkNo

`func (o *MatDetail) GetStkNo() string`

GetStkNo returns the StkNo field if non-nil, zero value otherwise.

### GetStkNoOk

`func (o *MatDetail) GetStkNoOk() (*string, bool)`

GetStkNoOk returns a tuple with the StkNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkNo

`func (o *MatDetail) SetStkNo(v string)`

SetStkNo sets StkNo field to given value.


### GetTDate

`func (o *MatDetail) GetTDate() string`

GetTDate returns the TDate field if non-nil, zero value otherwise.

### GetTDateOk

`func (o *MatDetail) GetTDateOk() (*string, bool)`

GetTDateOk returns a tuple with the TDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTDate

`func (o *MatDetail) SetTDate(v string)`

SetTDate sets TDate field to given value.


### GetTTime

`func (o *MatDetail) GetTTime() string`

GetTTime returns the TTime field if non-nil, zero value otherwise.

### GetTTimeOk

`func (o *MatDetail) GetTTimeOk() (*string, bool)`

GetTTimeOk returns a tuple with the TTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTime

`func (o *MatDetail) SetTTime(v string)`

SetTTime sets TTime field to given value.

### HasTTime

`func (o *MatDetail) HasTTime() bool`

HasTTime returns a boolean if a field has been set.

### GetTax

`func (o *MatDetail) GetTax() string`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *MatDetail) GetTaxOk() (*string, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *MatDetail) SetTax(v string)`

SetTax sets Tax field to given value.


### GetTaxG

`func (o *MatDetail) GetTaxG() string`

GetTaxG returns the TaxG field if non-nil, zero value otherwise.

### GetTaxGOk

`func (o *MatDetail) GetTaxGOk() (*string, bool)`

GetTaxGOk returns a tuple with the TaxG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxG

`func (o *MatDetail) SetTaxG(v string)`

SetTaxG sets TaxG field to given value.


### GetTrade

`func (o *MatDetail) GetTrade() int32`

GetTrade returns the Trade field if non-nil, zero value otherwise.

### GetTradeOk

`func (o *MatDetail) GetTradeOk() (*int32, bool)`

GetTradeOk returns a tuple with the Trade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrade

`func (o *MatDetail) SetTrade(v int32)`

SetTrade sets Trade field to given value.


### GetUserDef

`func (o *MatDetail) GetUserDef() string`

GetUserDef returns the UserDef field if non-nil, zero value otherwise.

### GetUserDefOk

`func (o *MatDetail) GetUserDefOk() (*string, bool)`

GetUserDefOk returns a tuple with the UserDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDef

`func (o *MatDetail) SetUserDef(v string)`

SetUserDef sets UserDef field to given value.

### HasUserDef

`func (o *MatDetail) HasUserDef() bool`

HasUserDef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


