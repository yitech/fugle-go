/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fuglego

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AppSchemaV2WalletInventoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppSchemaV2WalletInventoryResponse{}

// AppSchemaV2WalletInventoryResponse struct for AppSchemaV2WalletInventoryResponse
type AppSchemaV2WalletInventoryResponse struct {
	Symbol string `json:"symbol"`
	QtyShare int32 `json:"qty_share"`
}

type _AppSchemaV2WalletInventoryResponse AppSchemaV2WalletInventoryResponse

// NewAppSchemaV2WalletInventoryResponse instantiates a new AppSchemaV2WalletInventoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSchemaV2WalletInventoryResponse(symbol string, qtyShare int32) *AppSchemaV2WalletInventoryResponse {
	this := AppSchemaV2WalletInventoryResponse{}
	this.Symbol = symbol
	this.QtyShare = qtyShare
	return &this
}

// NewAppSchemaV2WalletInventoryResponseWithDefaults instantiates a new AppSchemaV2WalletInventoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSchemaV2WalletInventoryResponseWithDefaults() *AppSchemaV2WalletInventoryResponse {
	this := AppSchemaV2WalletInventoryResponse{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *AppSchemaV2WalletInventoryResponse) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *AppSchemaV2WalletInventoryResponse) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *AppSchemaV2WalletInventoryResponse) SetSymbol(v string) {
	o.Symbol = v
}

// GetQtyShare returns the QtyShare field value
func (o *AppSchemaV2WalletInventoryResponse) GetQtyShare() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QtyShare
}

// GetQtyShareOk returns a tuple with the QtyShare field value
// and a boolean to check if the value has been set.
func (o *AppSchemaV2WalletInventoryResponse) GetQtyShareOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QtyShare, true
}

// SetQtyShare sets field value
func (o *AppSchemaV2WalletInventoryResponse) SetQtyShare(v int32) {
	o.QtyShare = v
}

func (o AppSchemaV2WalletInventoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppSchemaV2WalletInventoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["qty_share"] = o.QtyShare
	return toSerialize, nil
}

func (o *AppSchemaV2WalletInventoryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"qty_share",
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

	varAppSchemaV2WalletInventoryResponse := _AppSchemaV2WalletInventoryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppSchemaV2WalletInventoryResponse)

	if err != nil {
		return err
	}

	*o = AppSchemaV2WalletInventoryResponse(varAppSchemaV2WalletInventoryResponse)

	return err
}

type NullableAppSchemaV2WalletInventoryResponse struct {
	value *AppSchemaV2WalletInventoryResponse
	isSet bool
}

func (v NullableAppSchemaV2WalletInventoryResponse) Get() *AppSchemaV2WalletInventoryResponse {
	return v.value
}

func (v *NullableAppSchemaV2WalletInventoryResponse) Set(val *AppSchemaV2WalletInventoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSchemaV2WalletInventoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSchemaV2WalletInventoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSchemaV2WalletInventoryResponse(val *AppSchemaV2WalletInventoryResponse) *NullableAppSchemaV2WalletInventoryResponse {
	return &NullableAppSchemaV2WalletInventoryResponse{value: val, isSet: true}
}

func (v NullableAppSchemaV2WalletInventoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSchemaV2WalletInventoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


