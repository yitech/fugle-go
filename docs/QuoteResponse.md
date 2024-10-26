# QuoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Type** | **string** |  | 
**Exchange** | **string** |  | 
**Market** | **string** |  | 
**Symbol** | **string** |  | 
**Name** | **string** |  | 
**ReferencePrice** | **float32** |  | 
**PreviousClose** | **float32** |  | 
**OpenPrice** | **float32** |  | 
**OpenTime** | **int32** |  | 
**HighPrice** | **float32** |  | 
**HighTime** | **int32** |  | 
**LowPrice** | **float32** |  | 
**LowTime** | **int32** |  | 
**ClosePrice** | **float32** |  | 
**CloseTime** | **int32** |  | 
**AvgPrice** | **float32** |  | 
**Change** | **float32** |  | 
**ChangePercent** | **float32** |  | 
**Amplitude** | **float32** |  | 
**LastPrice** | **float32** |  | 
**LastSize** | **int32** |  | 
**Bids** | [**[]BidAsk**](BidAsk.md) |  | 
**Asks** | [**[]BidAsk**](BidAsk.md) |  | 
**Total** | [**Total**](Total.md) |  | 
**LastTrade** | [**LastTrade**](LastTrade.md) |  | 
**LastTrial** | [**LastTrial**](LastTrial.md) |  | 
**IsTrial** | Pointer to **NullableBool** |  | [optional] 
**IsDelayedOpen** | Pointer to **NullableBool** |  | [optional] 
**IsDelayedClose** | Pointer to **NullableBool** |  | [optional] 
**IsOpen** | Pointer to **NullableBool** |  | [optional] 
**IsClose** | Pointer to **NullableBool** |  | [optional] 
**LastUpdated** | **int32** |  | 
**Serial** | **int32** |  | 

## Methods

### NewQuoteResponse

`func NewQuoteResponse(date string, type_ string, exchange string, market string, symbol string, name string, referencePrice float32, previousClose float32, openPrice float32, openTime int32, highPrice float32, highTime int32, lowPrice float32, lowTime int32, closePrice float32, closeTime int32, avgPrice float32, change float32, changePercent float32, amplitude float32, lastPrice float32, lastSize int32, bids []BidAsk, asks []BidAsk, total Total, lastTrade LastTrade, lastTrial LastTrial, lastUpdated int32, serial int32, ) *QuoteResponse`

NewQuoteResponse instantiates a new QuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteResponseWithDefaults

`func NewQuoteResponseWithDefaults() *QuoteResponse`

NewQuoteResponseWithDefaults instantiates a new QuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *QuoteResponse) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *QuoteResponse) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *QuoteResponse) SetDate(v string)`

SetDate sets Date field to given value.


### GetType

`func (o *QuoteResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuoteResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuoteResponse) SetType(v string)`

SetType sets Type field to given value.


### GetExchange

`func (o *QuoteResponse) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *QuoteResponse) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *QuoteResponse) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMarket

`func (o *QuoteResponse) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *QuoteResponse) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *QuoteResponse) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetSymbol

`func (o *QuoteResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QuoteResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QuoteResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *QuoteResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuoteResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuoteResponse) SetName(v string)`

SetName sets Name field to given value.


### GetReferencePrice

`func (o *QuoteResponse) GetReferencePrice() float32`

GetReferencePrice returns the ReferencePrice field if non-nil, zero value otherwise.

### GetReferencePriceOk

`func (o *QuoteResponse) GetReferencePriceOk() (*float32, bool)`

GetReferencePriceOk returns a tuple with the ReferencePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePrice

`func (o *QuoteResponse) SetReferencePrice(v float32)`

SetReferencePrice sets ReferencePrice field to given value.


### GetPreviousClose

`func (o *QuoteResponse) GetPreviousClose() float32`

GetPreviousClose returns the PreviousClose field if non-nil, zero value otherwise.

### GetPreviousCloseOk

`func (o *QuoteResponse) GetPreviousCloseOk() (*float32, bool)`

GetPreviousCloseOk returns a tuple with the PreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousClose

`func (o *QuoteResponse) SetPreviousClose(v float32)`

SetPreviousClose sets PreviousClose field to given value.


### GetOpenPrice

`func (o *QuoteResponse) GetOpenPrice() float32`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *QuoteResponse) GetOpenPriceOk() (*float32, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *QuoteResponse) SetOpenPrice(v float32)`

SetOpenPrice sets OpenPrice field to given value.


### GetOpenTime

`func (o *QuoteResponse) GetOpenTime() int32`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *QuoteResponse) GetOpenTimeOk() (*int32, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *QuoteResponse) SetOpenTime(v int32)`

SetOpenTime sets OpenTime field to given value.


### GetHighPrice

`func (o *QuoteResponse) GetHighPrice() float32`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *QuoteResponse) GetHighPriceOk() (*float32, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *QuoteResponse) SetHighPrice(v float32)`

SetHighPrice sets HighPrice field to given value.


### GetHighTime

`func (o *QuoteResponse) GetHighTime() int32`

GetHighTime returns the HighTime field if non-nil, zero value otherwise.

### GetHighTimeOk

`func (o *QuoteResponse) GetHighTimeOk() (*int32, bool)`

GetHighTimeOk returns a tuple with the HighTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighTime

`func (o *QuoteResponse) SetHighTime(v int32)`

SetHighTime sets HighTime field to given value.


### GetLowPrice

`func (o *QuoteResponse) GetLowPrice() float32`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *QuoteResponse) GetLowPriceOk() (*float32, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *QuoteResponse) SetLowPrice(v float32)`

SetLowPrice sets LowPrice field to given value.


### GetLowTime

`func (o *QuoteResponse) GetLowTime() int32`

GetLowTime returns the LowTime field if non-nil, zero value otherwise.

### GetLowTimeOk

`func (o *QuoteResponse) GetLowTimeOk() (*int32, bool)`

GetLowTimeOk returns a tuple with the LowTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowTime

`func (o *QuoteResponse) SetLowTime(v int32)`

SetLowTime sets LowTime field to given value.


### GetClosePrice

`func (o *QuoteResponse) GetClosePrice() float32`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *QuoteResponse) GetClosePriceOk() (*float32, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *QuoteResponse) SetClosePrice(v float32)`

SetClosePrice sets ClosePrice field to given value.


### GetCloseTime

`func (o *QuoteResponse) GetCloseTime() int32`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *QuoteResponse) GetCloseTimeOk() (*int32, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *QuoteResponse) SetCloseTime(v int32)`

SetCloseTime sets CloseTime field to given value.


### GetAvgPrice

`func (o *QuoteResponse) GetAvgPrice() float32`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *QuoteResponse) GetAvgPriceOk() (*float32, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *QuoteResponse) SetAvgPrice(v float32)`

SetAvgPrice sets AvgPrice field to given value.


### GetChange

`func (o *QuoteResponse) GetChange() float32`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *QuoteResponse) GetChangeOk() (*float32, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *QuoteResponse) SetChange(v float32)`

SetChange sets Change field to given value.


### GetChangePercent

`func (o *QuoteResponse) GetChangePercent() float32`

GetChangePercent returns the ChangePercent field if non-nil, zero value otherwise.

### GetChangePercentOk

`func (o *QuoteResponse) GetChangePercentOk() (*float32, bool)`

GetChangePercentOk returns a tuple with the ChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePercent

`func (o *QuoteResponse) SetChangePercent(v float32)`

SetChangePercent sets ChangePercent field to given value.


### GetAmplitude

`func (o *QuoteResponse) GetAmplitude() float32`

GetAmplitude returns the Amplitude field if non-nil, zero value otherwise.

### GetAmplitudeOk

`func (o *QuoteResponse) GetAmplitudeOk() (*float32, bool)`

GetAmplitudeOk returns a tuple with the Amplitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmplitude

`func (o *QuoteResponse) SetAmplitude(v float32)`

SetAmplitude sets Amplitude field to given value.


### GetLastPrice

`func (o *QuoteResponse) GetLastPrice() float32`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *QuoteResponse) GetLastPriceOk() (*float32, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *QuoteResponse) SetLastPrice(v float32)`

SetLastPrice sets LastPrice field to given value.


### GetLastSize

`func (o *QuoteResponse) GetLastSize() int32`

GetLastSize returns the LastSize field if non-nil, zero value otherwise.

### GetLastSizeOk

`func (o *QuoteResponse) GetLastSizeOk() (*int32, bool)`

GetLastSizeOk returns a tuple with the LastSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSize

`func (o *QuoteResponse) SetLastSize(v int32)`

SetLastSize sets LastSize field to given value.


### GetBids

`func (o *QuoteResponse) GetBids() []BidAsk`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *QuoteResponse) GetBidsOk() (*[]BidAsk, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *QuoteResponse) SetBids(v []BidAsk)`

SetBids sets Bids field to given value.


### GetAsks

`func (o *QuoteResponse) GetAsks() []BidAsk`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *QuoteResponse) GetAsksOk() (*[]BidAsk, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *QuoteResponse) SetAsks(v []BidAsk)`

SetAsks sets Asks field to given value.


### GetTotal

`func (o *QuoteResponse) GetTotal() Total`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QuoteResponse) GetTotalOk() (*Total, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QuoteResponse) SetTotal(v Total)`

SetTotal sets Total field to given value.


### GetLastTrade

`func (o *QuoteResponse) GetLastTrade() LastTrade`

GetLastTrade returns the LastTrade field if non-nil, zero value otherwise.

### GetLastTradeOk

`func (o *QuoteResponse) GetLastTradeOk() (*LastTrade, bool)`

GetLastTradeOk returns a tuple with the LastTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrade

`func (o *QuoteResponse) SetLastTrade(v LastTrade)`

SetLastTrade sets LastTrade field to given value.


### GetLastTrial

`func (o *QuoteResponse) GetLastTrial() LastTrial`

GetLastTrial returns the LastTrial field if non-nil, zero value otherwise.

### GetLastTrialOk

`func (o *QuoteResponse) GetLastTrialOk() (*LastTrial, bool)`

GetLastTrialOk returns a tuple with the LastTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrial

`func (o *QuoteResponse) SetLastTrial(v LastTrial)`

SetLastTrial sets LastTrial field to given value.


### GetIsTrial

`func (o *QuoteResponse) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *QuoteResponse) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *QuoteResponse) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.

### HasIsTrial

`func (o *QuoteResponse) HasIsTrial() bool`

HasIsTrial returns a boolean if a field has been set.

### SetIsTrialNil

`func (o *QuoteResponse) SetIsTrialNil(b bool)`

 SetIsTrialNil sets the value for IsTrial to be an explicit nil

### UnsetIsTrial
`func (o *QuoteResponse) UnsetIsTrial()`

UnsetIsTrial ensures that no value is present for IsTrial, not even an explicit nil
### GetIsDelayedOpen

`func (o *QuoteResponse) GetIsDelayedOpen() bool`

GetIsDelayedOpen returns the IsDelayedOpen field if non-nil, zero value otherwise.

### GetIsDelayedOpenOk

`func (o *QuoteResponse) GetIsDelayedOpenOk() (*bool, bool)`

GetIsDelayedOpenOk returns a tuple with the IsDelayedOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelayedOpen

`func (o *QuoteResponse) SetIsDelayedOpen(v bool)`

SetIsDelayedOpen sets IsDelayedOpen field to given value.

### HasIsDelayedOpen

`func (o *QuoteResponse) HasIsDelayedOpen() bool`

HasIsDelayedOpen returns a boolean if a field has been set.

### SetIsDelayedOpenNil

`func (o *QuoteResponse) SetIsDelayedOpenNil(b bool)`

 SetIsDelayedOpenNil sets the value for IsDelayedOpen to be an explicit nil

### UnsetIsDelayedOpen
`func (o *QuoteResponse) UnsetIsDelayedOpen()`

UnsetIsDelayedOpen ensures that no value is present for IsDelayedOpen, not even an explicit nil
### GetIsDelayedClose

`func (o *QuoteResponse) GetIsDelayedClose() bool`

GetIsDelayedClose returns the IsDelayedClose field if non-nil, zero value otherwise.

### GetIsDelayedCloseOk

`func (o *QuoteResponse) GetIsDelayedCloseOk() (*bool, bool)`

GetIsDelayedCloseOk returns a tuple with the IsDelayedClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelayedClose

`func (o *QuoteResponse) SetIsDelayedClose(v bool)`

SetIsDelayedClose sets IsDelayedClose field to given value.

### HasIsDelayedClose

`func (o *QuoteResponse) HasIsDelayedClose() bool`

HasIsDelayedClose returns a boolean if a field has been set.

### SetIsDelayedCloseNil

`func (o *QuoteResponse) SetIsDelayedCloseNil(b bool)`

 SetIsDelayedCloseNil sets the value for IsDelayedClose to be an explicit nil

### UnsetIsDelayedClose
`func (o *QuoteResponse) UnsetIsDelayedClose()`

UnsetIsDelayedClose ensures that no value is present for IsDelayedClose, not even an explicit nil
### GetIsOpen

`func (o *QuoteResponse) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *QuoteResponse) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *QuoteResponse) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.

### HasIsOpen

`func (o *QuoteResponse) HasIsOpen() bool`

HasIsOpen returns a boolean if a field has been set.

### SetIsOpenNil

`func (o *QuoteResponse) SetIsOpenNil(b bool)`

 SetIsOpenNil sets the value for IsOpen to be an explicit nil

### UnsetIsOpen
`func (o *QuoteResponse) UnsetIsOpen()`

UnsetIsOpen ensures that no value is present for IsOpen, not even an explicit nil
### GetIsClose

`func (o *QuoteResponse) GetIsClose() bool`

GetIsClose returns the IsClose field if non-nil, zero value otherwise.

### GetIsCloseOk

`func (o *QuoteResponse) GetIsCloseOk() (*bool, bool)`

GetIsCloseOk returns a tuple with the IsClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClose

`func (o *QuoteResponse) SetIsClose(v bool)`

SetIsClose sets IsClose field to given value.

### HasIsClose

`func (o *QuoteResponse) HasIsClose() bool`

HasIsClose returns a boolean if a field has been set.

### SetIsCloseNil

`func (o *QuoteResponse) SetIsCloseNil(b bool)`

 SetIsCloseNil sets the value for IsClose to be an explicit nil

### UnsetIsClose
`func (o *QuoteResponse) UnsetIsClose()`

UnsetIsClose ensures that no value is present for IsClose, not even an explicit nil
### GetLastUpdated

`func (o *QuoteResponse) GetLastUpdated() int32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *QuoteResponse) GetLastUpdatedOk() (*int32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *QuoteResponse) SetLastUpdated(v int32)`

SetLastUpdated sets LastUpdated field to given value.


### GetSerial

`func (o *QuoteResponse) GetSerial() int32`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *QuoteResponse) GetSerialOk() (*int32, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *QuoteResponse) SetSerial(v int32)`

SetSerial sets Serial field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


