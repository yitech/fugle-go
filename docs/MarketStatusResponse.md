# MarketStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsTradingDay** | **bool** |  | 
**LastTradingDay** | **string** |  | 
**NextTradingDay** | **string** |  | 

## Methods

### NewMarketStatusResponse

`func NewMarketStatusResponse(isTradingDay bool, lastTradingDay string, nextTradingDay string, ) *MarketStatusResponse`

NewMarketStatusResponse instantiates a new MarketStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketStatusResponseWithDefaults

`func NewMarketStatusResponseWithDefaults() *MarketStatusResponse`

NewMarketStatusResponseWithDefaults instantiates a new MarketStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsTradingDay

`func (o *MarketStatusResponse) GetIsTradingDay() bool`

GetIsTradingDay returns the IsTradingDay field if non-nil, zero value otherwise.

### GetIsTradingDayOk

`func (o *MarketStatusResponse) GetIsTradingDayOk() (*bool, bool)`

GetIsTradingDayOk returns a tuple with the IsTradingDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTradingDay

`func (o *MarketStatusResponse) SetIsTradingDay(v bool)`

SetIsTradingDay sets IsTradingDay field to given value.


### GetLastTradingDay

`func (o *MarketStatusResponse) GetLastTradingDay() string`

GetLastTradingDay returns the LastTradingDay field if non-nil, zero value otherwise.

### GetLastTradingDayOk

`func (o *MarketStatusResponse) GetLastTradingDayOk() (*string, bool)`

GetLastTradingDayOk returns a tuple with the LastTradingDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTradingDay

`func (o *MarketStatusResponse) SetLastTradingDay(v string)`

SetLastTradingDay sets LastTradingDay field to given value.


### GetNextTradingDay

`func (o *MarketStatusResponse) GetNextTradingDay() string`

GetNextTradingDay returns the NextTradingDay field if non-nil, zero value otherwise.

### GetNextTradingDayOk

`func (o *MarketStatusResponse) GetNextTradingDayOk() (*string, bool)`

GetNextTradingDayOk returns a tuple with the NextTradingDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTradingDay

`func (o *MarketStatusResponse) SetNextTradingDay(v string)`

SetNextTradingDay sets NextTradingDay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


