/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fuglego

import (
	"encoding/json"
	"fmt"
)

// BSFlag the model 'BSFlag'
type BSFlag string

// List of BSFlag
const (
	FOK BSFlag = "F"
	IOC BSFlag = "I"
	ROD BSFlag = "R"
)

// All allowed values of BSFlag enum
var AllowedBSFlagEnumValues = []BSFlag{
	"F",
	"I",
	"R",
}

func (v *BSFlag) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BSFlag(value)
	for _, existing := range AllowedBSFlagEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BSFlag", value)
}

// NewBSFlagFromValue returns a pointer to a valid BSFlag
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBSFlagFromValue(v string) (*BSFlag, error) {
	ev := BSFlag(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BSFlag: valid values are %v", v, AllowedBSFlagEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BSFlag) IsValid() bool {
	for _, existing := range AllowedBSFlagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BSFlag value
func (v BSFlag) Ptr() *BSFlag {
	return &v
}

type NullableBSFlag struct {
	value *BSFlag
	isSet bool
}

func (v NullableBSFlag) Get() *BSFlag {
	return v.value
}

func (v *NullableBSFlag) Set(val *BSFlag) {
	v.value = val
	v.isSet = true
}

func (v NullableBSFlag) IsSet() bool {
	return v.isSet
}

func (v *NullableBSFlag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBSFlag(val *BSFlag) *NullableBSFlag {
	return &NullableBSFlag{value: val, isSet: true}
}

func (v NullableBSFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBSFlag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

