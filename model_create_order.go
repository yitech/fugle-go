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

// checks if the CreateOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrder{}

// CreateOrder struct for CreateOrder
type CreateOrder struct {
	BuySell Action `json:"buy_sell"`
	ApCode APCode `json:"ap_code"`
	PriceFlag PriceFlag `json:"price_flag"`
	BsFlag BSFlag `json:"bs_flag"`
	Trade Trade `json:"trade"`
	StockNo string `json:"stock_no"`
	Quantity int32 `json:"quantity"`
	Price float32 `json:"price"`
}

type _CreateOrder CreateOrder

// NewCreateOrder instantiates a new CreateOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrder(buySell Action, apCode APCode, priceFlag PriceFlag, bsFlag BSFlag, trade Trade, stockNo string, quantity int32, price float32) *CreateOrder {
	this := CreateOrder{}
	this.BuySell = buySell
	this.ApCode = apCode
	this.PriceFlag = priceFlag
	this.BsFlag = bsFlag
	this.Trade = trade
	this.StockNo = stockNo
	this.Quantity = quantity
	this.Price = price
	return &this
}

// NewCreateOrderWithDefaults instantiates a new CreateOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrderWithDefaults() *CreateOrder {
	this := CreateOrder{}
	return &this
}

// GetBuySell returns the BuySell field value
func (o *CreateOrder) GetBuySell() Action {
	if o == nil {
		var ret Action
		return ret
	}

	return o.BuySell
}

// GetBuySellOk returns a tuple with the BuySell field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetBuySellOk() (*Action, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuySell, true
}

// SetBuySell sets field value
func (o *CreateOrder) SetBuySell(v Action) {
	o.BuySell = v
}

// GetApCode returns the ApCode field value
func (o *CreateOrder) GetApCode() APCode {
	if o == nil {
		var ret APCode
		return ret
	}

	return o.ApCode
}

// GetApCodeOk returns a tuple with the ApCode field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetApCodeOk() (*APCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApCode, true
}

// SetApCode sets field value
func (o *CreateOrder) SetApCode(v APCode) {
	o.ApCode = v
}

// GetPriceFlag returns the PriceFlag field value
func (o *CreateOrder) GetPriceFlag() PriceFlag {
	if o == nil {
		var ret PriceFlag
		return ret
	}

	return o.PriceFlag
}

// GetPriceFlagOk returns a tuple with the PriceFlag field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetPriceFlagOk() (*PriceFlag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceFlag, true
}

// SetPriceFlag sets field value
func (o *CreateOrder) SetPriceFlag(v PriceFlag) {
	o.PriceFlag = v
}

// GetBsFlag returns the BsFlag field value
func (o *CreateOrder) GetBsFlag() BSFlag {
	if o == nil {
		var ret BSFlag
		return ret
	}

	return o.BsFlag
}

// GetBsFlagOk returns a tuple with the BsFlag field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetBsFlagOk() (*BSFlag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BsFlag, true
}

// SetBsFlag sets field value
func (o *CreateOrder) SetBsFlag(v BSFlag) {
	o.BsFlag = v
}

// GetTrade returns the Trade field value
func (o *CreateOrder) GetTrade() Trade {
	if o == nil {
		var ret Trade
		return ret
	}

	return o.Trade
}

// GetTradeOk returns a tuple with the Trade field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetTradeOk() (*Trade, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trade, true
}

// SetTrade sets field value
func (o *CreateOrder) SetTrade(v Trade) {
	o.Trade = v
}

// GetStockNo returns the StockNo field value
func (o *CreateOrder) GetStockNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StockNo
}

// GetStockNoOk returns a tuple with the StockNo field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetStockNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StockNo, true
}

// SetStockNo sets field value
func (o *CreateOrder) SetStockNo(v string) {
	o.StockNo = v
}

// GetQuantity returns the Quantity field value
func (o *CreateOrder) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CreateOrder) SetQuantity(v int32) {
	o.Quantity = v
}

// GetPrice returns the Price field value
func (o *CreateOrder) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *CreateOrder) SetPrice(v float32) {
	o.Price = v
}

func (o CreateOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buy_sell"] = o.BuySell
	toSerialize["ap_code"] = o.ApCode
	toSerialize["price_flag"] = o.PriceFlag
	toSerialize["bs_flag"] = o.BsFlag
	toSerialize["trade"] = o.Trade
	toSerialize["stock_no"] = o.StockNo
	toSerialize["quantity"] = o.Quantity
	toSerialize["price"] = o.Price
	return toSerialize, nil
}

func (o *CreateOrder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"buy_sell",
		"ap_code",
		"price_flag",
		"bs_flag",
		"trade",
		"stock_no",
		"quantity",
		"price",
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

	varCreateOrder := _CreateOrder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrder)

	if err != nil {
		return err
	}

	*o = CreateOrder(varCreateOrder)

	return err
}

type NullableCreateOrder struct {
	value *CreateOrder
	isSet bool
}

func (v NullableCreateOrder) Get() *CreateOrder {
	return v.value
}

func (v *NullableCreateOrder) Set(val *CreateOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrder(val *CreateOrder) *NullableCreateOrder {
	return &NullableCreateOrder{value: val, isSet: true}
}

func (v NullableCreateOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


