/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fugle-go

import (
	"encoding/json"
	"fmt"
)

// Action the model 'Action'
type Action string

// List of Action
const (
	Buy Action = "B"
	Sell Action = "S"
)

// All allowed values of Action enum
var AllowedActionEnumValues = []Action{
	"B",
	"S",
}

func (v *Action) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Action(value)
	for _, existing := range AllowedActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Action", value)
}

// NewActionFromValue returns a pointer to a valid Action
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionFromValue(v string) (*Action, error) {
	ev := Action(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Action: valid values are %v", v, AllowedActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Action) IsValid() bool {
	for _, existing := range AllowedActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Action value
func (v Action) Ptr() *Action {
	return &v
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

