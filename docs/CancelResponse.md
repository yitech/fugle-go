# CancelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetCode** | **string** |  | 
**RetMsg** | **string** |  | 
**OrdDate** | **string** |  | 
**OrdTime** | **string** |  | 

## Methods

### NewCancelResponse

`func NewCancelResponse(retCode string, retMsg string, ordDate string, ordTime string, ) *CancelResponse`

NewCancelResponse instantiates a new CancelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelResponseWithDefaults

`func NewCancelResponseWithDefaults() *CancelResponse`

NewCancelResponseWithDefaults instantiates a new CancelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetCode

`func (o *CancelResponse) GetRetCode() string`

GetRetCode returns the RetCode field if non-nil, zero value otherwise.

### GetRetCodeOk

`func (o *CancelResponse) GetRetCodeOk() (*string, bool)`

GetRetCodeOk returns a tuple with the RetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetCode

`func (o *CancelResponse) SetRetCode(v string)`

SetRetCode sets RetCode field to given value.


### GetRetMsg

`func (o *CancelResponse) GetRetMsg() string`

GetRetMsg returns the RetMsg field if non-nil, zero value otherwise.

### GetRetMsgOk

`func (o *CancelResponse) GetRetMsgOk() (*string, bool)`

GetRetMsgOk returns a tuple with the RetMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetMsg

`func (o *CancelResponse) SetRetMsg(v string)`

SetRetMsg sets RetMsg field to given value.


### GetOrdDate

`func (o *CancelResponse) GetOrdDate() string`

GetOrdDate returns the OrdDate field if non-nil, zero value otherwise.

### GetOrdDateOk

`func (o *CancelResponse) GetOrdDateOk() (*string, bool)`

GetOrdDateOk returns a tuple with the OrdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdDate

`func (o *CancelResponse) SetOrdDate(v string)`

SetOrdDate sets OrdDate field to given value.


### GetOrdTime

`func (o *CancelResponse) GetOrdTime() string`

GetOrdTime returns the OrdTime field if non-nil, zero value otherwise.

### GetOrdTimeOk

`func (o *CancelResponse) GetOrdTimeOk() (*string, bool)`

GetOrdTimeOk returns a tuple with the OrdTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdTime

`func (o *CancelResponse) SetOrdTime(v string)`

SetOrdTime sets OrdTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


