/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// Trade the model 'Trade'
type Trade string

// List of Trade
const (
	Cash Trade = "0"
	Margin Trade = "3"
	Short Trade = "4"
	DayTradingSell Trade = "A"
)

// All allowed values of Trade enum
var AllowedTradeEnumValues = []Trade{
	"0",
	"3",
	"4",
	"A",
}

func (v *Trade) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Trade(value)
	for _, existing := range AllowedTradeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Trade", value)
}

// NewTradeFromValue returns a pointer to a valid Trade
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTradeFromValue(v string) (*Trade, error) {
	ev := Trade(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Trade: valid values are %v", v, AllowedTradeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Trade) IsValid() bool {
	for _, existing := range AllowedTradeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Trade value
func (v Trade) Ptr() *Trade {
	return &v
}

type NullableTrade struct {
	value *Trade
	isSet bool
}

func (v NullableTrade) Get() *Trade {
	return v.value
}

func (v *NullableTrade) Set(val *Trade) {
	v.value = val
	v.isSet = true
}

func (v NullableTrade) IsSet() bool {
	return v.isSet
}

func (v *NullableTrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrade(val *Trade) *NullableTrade {
	return &NullableTrade{value: val, isSet: true}
}

func (v NullableTrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
