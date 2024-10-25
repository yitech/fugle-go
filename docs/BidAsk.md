# BidAsk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **float32** |  | 
**Size** | **int32** |  | 

## Methods

### NewBidAsk

`func NewBidAsk(price float32, size int32, ) *BidAsk`

NewBidAsk instantiates a new BidAsk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBidAskWithDefaults

`func NewBidAskWithDefaults() *BidAsk`

NewBidAskWithDefaults instantiates a new BidAsk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *BidAsk) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BidAsk) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BidAsk) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetSize

`func (o *BidAsk) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BidAsk) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BidAsk) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


