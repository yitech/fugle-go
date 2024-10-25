/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LastTrial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastTrial{}

// LastTrial struct for LastTrial
type LastTrial struct {
	Bid float32 `json:"bid"`
	Ask float32 `json:"ask"`
	Price float32 `json:"price"`
	Size int32 `json:"size"`
	Time int32 `json:"time"`
	Serial int32 `json:"serial"`
}

type _LastTrial LastTrial

// NewLastTrial instantiates a new LastTrial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastTrial(bid float32, ask float32, price float32, size int32, time int32, serial int32) *LastTrial {
	this := LastTrial{}
	this.Bid = bid
	this.Ask = ask
	this.Price = price
	this.Size = size
	this.Time = time
	this.Serial = serial
	return &this
}

// NewLastTrialWithDefaults instantiates a new LastTrial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastTrialWithDefaults() *LastTrial {
	this := LastTrial{}
	return &this
}

// GetBid returns the Bid field value
func (o *LastTrial) GetBid() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Bid
}

// GetBidOk returns a tuple with the Bid field value
// and a boolean to check if the value has been set.
func (o *LastTrial) GetBidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bid, true
}

// SetBid sets field value
func (o *LastTrial) SetBid(v float32) {
	o.Bid = v
}

// GetAsk returns the Ask field value
func (o *LastTrial) GetAsk() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Ask
}

// GetAskOk returns a tuple with the Ask field value
// and a boolean to check if the value has been set.
func (o *LastTrial) GetAskOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ask, true
}

// SetAsk sets field value
func (o *LastTrial) SetAsk(v float32) {
	o.Ask = v
}

// GetPrice returns the Price field value
func (o *LastTrial) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *LastTrial) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *LastTrial) SetPrice(v float32) {
	o.Price = v
}

// GetSize returns the Size field value
func (o *LastTrial) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *LastTrial) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *LastTrial) SetSize(v int32) {
	o.Size = v
}

// GetTime returns the Time field value
func (o *LastTrial) GetTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *LastTrial) GetTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *LastTrial) SetTime(v int32) {
	o.Time = v
}

// GetSerial returns the Serial field value
func (o *LastTrial) GetSerial() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *LastTrial) GetSerialOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *LastTrial) SetSerial(v int32) {
	o.Serial = v
}

func (o LastTrial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastTrial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bid"] = o.Bid
	toSerialize["ask"] = o.Ask
	toSerialize["price"] = o.Price
	toSerialize["size"] = o.Size
	toSerialize["time"] = o.Time
	toSerialize["serial"] = o.Serial
	return toSerialize, nil
}

func (o *LastTrial) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bid",
		"ask",
		"price",
		"size",
		"time",
		"serial",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLastTrial := _LastTrial{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLastTrial)

	if err != nil {
		return err
	}

	*o = LastTrial(varLastTrial)

	return err
}

type NullableLastTrial struct {
	value *LastTrial
	isSet bool
}

func (v NullableLastTrial) Get() *LastTrial {
	return v.value
}

func (v *NullableLastTrial) Set(val *LastTrial) {
	v.value = val
	v.isSet = true
}

func (v NullableLastTrial) IsSet() bool {
	return v.isSet
}

func (v *NullableLastTrial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastTrial(val *LastTrial) *NullableLastTrial {
	return &NullableLastTrial{value: val, isSet: true}
}

func (v NullableLastTrial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastTrial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

