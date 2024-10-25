# KLinesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**Type** | **string** |  | 
**Exchange** | **string** |  | 
**Market** | **string** |  | 
**Timeframe** | **string** |  | 
**Data** | [**[]KLine**](KLine.md) |  | 

## Methods

### NewKLinesResponse

`func NewKLinesResponse(symbol string, type_ string, exchange string, market string, timeframe string, data []KLine, ) *KLinesResponse`

NewKLinesResponse instantiates a new KLinesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKLinesResponseWithDefaults

`func NewKLinesResponseWithDefaults() *KLinesResponse`

NewKLinesResponseWithDefaults instantiates a new KLinesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *KLinesResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *KLinesResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *KLinesResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetType

`func (o *KLinesResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KLinesResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KLinesResponse) SetType(v string)`

SetType sets Type field to given value.


### GetExchange

`func (o *KLinesResponse) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *KLinesResponse) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *KLinesResponse) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMarket

`func (o *KLinesResponse) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *KLinesResponse) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *KLinesResponse) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetTimeframe

`func (o *KLinesResponse) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *KLinesResponse) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *KLinesResponse) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.


### GetData

`func (o *KLinesResponse) GetData() []KLine`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KLinesResponse) GetDataOk() (*[]KLine, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KLinesResponse) SetData(v []KLine)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


