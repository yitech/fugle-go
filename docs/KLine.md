# KLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Open** | **float32** |  | 
**High** | **float32** |  | 
**Low** | **float32** |  | 
**Close** | **float32** |  | 
**Volume** | **int32** |  | 

## Methods

### NewKLine

`func NewKLine(date string, open float32, high float32, low float32, close float32, volume int32, ) *KLine`

NewKLine instantiates a new KLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKLineWithDefaults

`func NewKLineWithDefaults() *KLine`

NewKLineWithDefaults instantiates a new KLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *KLine) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *KLine) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *KLine) SetDate(v string)`

SetDate sets Date field to given value.


### GetOpen

`func (o *KLine) GetOpen() float32`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *KLine) GetOpenOk() (*float32, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *KLine) SetOpen(v float32)`

SetOpen sets Open field to given value.


### GetHigh

`func (o *KLine) GetHigh() float32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *KLine) GetHighOk() (*float32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *KLine) SetHigh(v float32)`

SetHigh sets High field to given value.


### GetLow

`func (o *KLine) GetLow() float32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *KLine) GetLowOk() (*float32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *KLine) SetLow(v float32)`

SetLow sets Low field to given value.


### GetClose

`func (o *KLine) GetClose() float32`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *KLine) GetCloseOk() (*float32, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *KLine) SetClose(v float32)`

SetClose sets Close field to given value.


### GetVolume

`func (o *KLine) GetVolume() int32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *KLine) GetVolumeOk() (*int32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *KLine) SetVolume(v int32)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


