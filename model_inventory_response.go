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

// checks if the InventoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryResponse{}

// InventoryResponse struct for InventoryResponse
type InventoryResponse struct {
	ApCode *string `json:"ap_code,omitempty"`
	CostQty float32 `json:"cost_qty"`
	CostSum float32 `json:"cost_sum"`
	MakeAPer float32 `json:"make_a_per"`
	MakeASum float32 `json:"make_a_sum"`
	PriceAvg float32 `json:"price_avg"`
	PriceEvn float32 `json:"price_evn"`
	PriceMkt float32 `json:"price_mkt"`
	PriceNow float32 `json:"price_now"`
	PriceQtySum float32 `json:"price_qty_sum"`
	QtyB int32 `json:"qty_b"`
	QtyBm int32 `json:"qty_bm"`
	QtyC int32 `json:"qty_c"`
	QtyL int32 `json:"qty_l"`
	QtyS int32 `json:"qty_s"`
	QtySm int32 `json:"qty_sm"`
	RecVaSum float32 `json:"rec_va_sum"`
	SType string `json:"s_type"`
	StkDats []InventoryDetail `json:"stk_dats"`
	StkNa string `json:"stk_na"`
	StkNo string `json:"stk_no"`
	Trade *int32 `json:"trade,omitempty"`
	ValueMkt float32 `json:"value_mkt"`
	ValueNow float32 `json:"value_now"`
}

type _InventoryResponse InventoryResponse

// NewInventoryResponse instantiates a new InventoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryResponse(costQty float32, costSum float32, makeAPer float32, makeASum float32, priceAvg float32, priceEvn float32, priceMkt float32, priceNow float32, priceQtySum float32, qtyB int32, qtyBm int32, qtyC int32, qtyL int32, qtyS int32, qtySm int32, recVaSum float32, sType string, stkDats []InventoryDetail, stkNa string, stkNo string, valueMkt float32, valueNow float32) *InventoryResponse {
	this := InventoryResponse{}
	this.CostQty = costQty
	this.CostSum = costSum
	this.MakeAPer = makeAPer
	this.MakeASum = makeASum
	this.PriceAvg = priceAvg
	this.PriceEvn = priceEvn
	this.PriceMkt = priceMkt
	this.PriceNow = priceNow
	this.PriceQtySum = priceQtySum
	this.QtyB = qtyB
	this.QtyBm = qtyBm
	this.QtyC = qtyC
	this.QtyL = qtyL
	this.QtyS = qtyS
	this.QtySm = qtySm
	this.RecVaSum = recVaSum
	this.SType = sType
	this.StkDats = stkDats
	this.StkNa = stkNa
	this.StkNo = stkNo
	this.ValueMkt = valueMkt
	this.ValueNow = valueNow
	return &this
}

// NewInventoryResponseWithDefaults instantiates a new InventoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryResponseWithDefaults() *InventoryResponse {
	this := InventoryResponse{}
	return &this
}

// GetApCode returns the ApCode field value if set, zero value otherwise.
func (o *InventoryResponse) GetApCode() string {
	if o == nil || IsNil(o.ApCode) {
		var ret string
		return ret
	}
	return *o.ApCode
}

// GetApCodeOk returns a tuple with the ApCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetApCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ApCode) {
		return nil, false
	}
	return o.ApCode, true
}

// HasApCode returns a boolean if a field has been set.
func (o *InventoryResponse) HasApCode() bool {
	if o != nil && !IsNil(o.ApCode) {
		return true
	}

	return false
}

// SetApCode gets a reference to the given string and assigns it to the ApCode field.
func (o *InventoryResponse) SetApCode(v string) {
	o.ApCode = &v
}

// GetCostQty returns the CostQty field value
func (o *InventoryResponse) GetCostQty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CostQty
}

// GetCostQtyOk returns a tuple with the CostQty field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetCostQtyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostQty, true
}

// SetCostQty sets field value
func (o *InventoryResponse) SetCostQty(v float32) {
	o.CostQty = v
}

// GetCostSum returns the CostSum field value
func (o *InventoryResponse) GetCostSum() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CostSum
}

// GetCostSumOk returns a tuple with the CostSum field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetCostSumOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostSum, true
}

// SetCostSum sets field value
func (o *InventoryResponse) SetCostSum(v float32) {
	o.CostSum = v
}

// GetMakeAPer returns the MakeAPer field value
func (o *InventoryResponse) GetMakeAPer() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MakeAPer
}

// GetMakeAPerOk returns a tuple with the MakeAPer field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetMakeAPerOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MakeAPer, true
}

// SetMakeAPer sets field value
func (o *InventoryResponse) SetMakeAPer(v float32) {
	o.MakeAPer = v
}

// GetMakeASum returns the MakeASum field value
func (o *InventoryResponse) GetMakeASum() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MakeASum
}

// GetMakeASumOk returns a tuple with the MakeASum field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetMakeASumOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MakeASum, true
}

// SetMakeASum sets field value
func (o *InventoryResponse) SetMakeASum(v float32) {
	o.MakeASum = v
}

// GetPriceAvg returns the PriceAvg field value
func (o *InventoryResponse) GetPriceAvg() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PriceAvg
}

// GetPriceAvgOk returns a tuple with the PriceAvg field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetPriceAvgOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceAvg, true
}

// SetPriceAvg sets field value
func (o *InventoryResponse) SetPriceAvg(v float32) {
	o.PriceAvg = v
}

// GetPriceEvn returns the PriceEvn field value
func (o *InventoryResponse) GetPriceEvn() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PriceEvn
}

// GetPriceEvnOk returns a tuple with the PriceEvn field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetPriceEvnOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceEvn, true
}

// SetPriceEvn sets field value
func (o *InventoryResponse) SetPriceEvn(v float32) {
	o.PriceEvn = v
}

// GetPriceMkt returns the PriceMkt field value
func (o *InventoryResponse) GetPriceMkt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PriceMkt
}

// GetPriceMktOk returns a tuple with the PriceMkt field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetPriceMktOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceMkt, true
}

// SetPriceMkt sets field value
func (o *InventoryResponse) SetPriceMkt(v float32) {
	o.PriceMkt = v
}

// GetPriceNow returns the PriceNow field value
func (o *InventoryResponse) GetPriceNow() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PriceNow
}

// GetPriceNowOk returns a tuple with the PriceNow field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetPriceNowOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceNow, true
}

// SetPriceNow sets field value
func (o *InventoryResponse) SetPriceNow(v float32) {
	o.PriceNow = v
}

// GetPriceQtySum returns the PriceQtySum field value
func (o *InventoryResponse) GetPriceQtySum() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PriceQtySum
}

// GetPriceQtySumOk returns a tuple with the PriceQtySum field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetPriceQtySumOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceQtySum, true
}

// SetPriceQtySum sets field value
func (o *InventoryResponse) SetPriceQtySum(v float32) {
	o.PriceQtySum = v
}

// GetQtyB returns the QtyB field value
func (o *InventoryResponse) GetQtyB() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QtyB
}

// GetQtyBOk returns a tuple with the QtyB field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetQtyBOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QtyB, true
}

// SetQtyB sets field value
func (o *InventoryResponse) SetQtyB(v int32) {
	o.QtyB = v
}

// GetQtyBm returns the QtyBm field value
func (o *InventoryResponse) GetQtyBm() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QtyBm
}

// GetQtyBmOk returns a tuple with the QtyBm field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetQtyBmOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QtyBm, true
}

// SetQtyBm sets field value
func (o *InventoryResponse) SetQtyBm(v int32) {
	o.QtyBm = v
}

// GetQtyC returns the QtyC field value
func (o *InventoryResponse) GetQtyC() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QtyC
}

// GetQtyCOk returns a tuple with the QtyC field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetQtyCOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QtyC, true
}

// SetQtyC sets field value
func (o *InventoryResponse) SetQtyC(v int32) {
	o.QtyC = v
}

// GetQtyL returns the QtyL field value
func (o *InventoryResponse) GetQtyL() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QtyL
}

// GetQtyLOk returns a tuple with the QtyL field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetQtyLOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QtyL, true
}

// SetQtyL sets field value
func (o *InventoryResponse) SetQtyL(v int32) {
	o.QtyL = v
}

// GetQtyS returns the QtyS field value
func (o *InventoryResponse) GetQtyS() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QtyS
}

// GetQtySOk returns a tuple with the QtyS field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetQtySOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QtyS, true
}

// SetQtyS sets field value
func (o *InventoryResponse) SetQtyS(v int32) {
	o.QtyS = v
}

// GetQtySm returns the QtySm field value
func (o *InventoryResponse) GetQtySm() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QtySm
}

// GetQtySmOk returns a tuple with the QtySm field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetQtySmOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QtySm, true
}

// SetQtySm sets field value
func (o *InventoryResponse) SetQtySm(v int32) {
	o.QtySm = v
}

// GetRecVaSum returns the RecVaSum field value
func (o *InventoryResponse) GetRecVaSum() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RecVaSum
}

// GetRecVaSumOk returns a tuple with the RecVaSum field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetRecVaSumOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecVaSum, true
}

// SetRecVaSum sets field value
func (o *InventoryResponse) SetRecVaSum(v float32) {
	o.RecVaSum = v
}

// GetSType returns the SType field value
func (o *InventoryResponse) GetSType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SType
}

// GetSTypeOk returns a tuple with the SType field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetSTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SType, true
}

// SetSType sets field value
func (o *InventoryResponse) SetSType(v string) {
	o.SType = v
}

// GetStkDats returns the StkDats field value
func (o *InventoryResponse) GetStkDats() []InventoryDetail {
	if o == nil {
		var ret []InventoryDetail
		return ret
	}

	return o.StkDats
}

// GetStkDatsOk returns a tuple with the StkDats field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetStkDatsOk() ([]InventoryDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.StkDats, true
}

// SetStkDats sets field value
func (o *InventoryResponse) SetStkDats(v []InventoryDetail) {
	o.StkDats = v
}

// GetStkNa returns the StkNa field value
func (o *InventoryResponse) GetStkNa() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StkNa
}

// GetStkNaOk returns a tuple with the StkNa field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetStkNaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StkNa, true
}

// SetStkNa sets field value
func (o *InventoryResponse) SetStkNa(v string) {
	o.StkNa = v
}

// GetStkNo returns the StkNo field value
func (o *InventoryResponse) GetStkNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StkNo
}

// GetStkNoOk returns a tuple with the StkNo field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetStkNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StkNo, true
}

// SetStkNo sets field value
func (o *InventoryResponse) SetStkNo(v string) {
	o.StkNo = v
}

// GetTrade returns the Trade field value if set, zero value otherwise.
func (o *InventoryResponse) GetTrade() int32 {
	if o == nil || IsNil(o.Trade) {
		var ret int32
		return ret
	}
	return *o.Trade
}

// GetTradeOk returns a tuple with the Trade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetTradeOk() (*int32, bool) {
	if o == nil || IsNil(o.Trade) {
		return nil, false
	}
	return o.Trade, true
}

// HasTrade returns a boolean if a field has been set.
func (o *InventoryResponse) HasTrade() bool {
	if o != nil && !IsNil(o.Trade) {
		return true
	}

	return false
}

// SetTrade gets a reference to the given int32 and assigns it to the Trade field.
func (o *InventoryResponse) SetTrade(v int32) {
	o.Trade = &v
}

// GetValueMkt returns the ValueMkt field value
func (o *InventoryResponse) GetValueMkt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ValueMkt
}

// GetValueMktOk returns a tuple with the ValueMkt field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetValueMktOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueMkt, true
}

// SetValueMkt sets field value
func (o *InventoryResponse) SetValueMkt(v float32) {
	o.ValueMkt = v
}

// GetValueNow returns the ValueNow field value
func (o *InventoryResponse) GetValueNow() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ValueNow
}

// GetValueNowOk returns a tuple with the ValueNow field value
// and a boolean to check if the value has been set.
func (o *InventoryResponse) GetValueNowOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueNow, true
}

// SetValueNow sets field value
func (o *InventoryResponse) SetValueNow(v float32) {
	o.ValueNow = v
}

func (o InventoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApCode) {
		toSerialize["ap_code"] = o.ApCode
	}
	toSerialize["cost_qty"] = o.CostQty
	toSerialize["cost_sum"] = o.CostSum
	toSerialize["make_a_per"] = o.MakeAPer
	toSerialize["make_a_sum"] = o.MakeASum
	toSerialize["price_avg"] = o.PriceAvg
	toSerialize["price_evn"] = o.PriceEvn
	toSerialize["price_mkt"] = o.PriceMkt
	toSerialize["price_now"] = o.PriceNow
	toSerialize["price_qty_sum"] = o.PriceQtySum
	toSerialize["qty_b"] = o.QtyB
	toSerialize["qty_bm"] = o.QtyBm
	toSerialize["qty_c"] = o.QtyC
	toSerialize["qty_l"] = o.QtyL
	toSerialize["qty_s"] = o.QtyS
	toSerialize["qty_sm"] = o.QtySm
	toSerialize["rec_va_sum"] = o.RecVaSum
	toSerialize["s_type"] = o.SType
	toSerialize["stk_dats"] = o.StkDats
	toSerialize["stk_na"] = o.StkNa
	toSerialize["stk_no"] = o.StkNo
	if !IsNil(o.Trade) {
		toSerialize["trade"] = o.Trade
	}
	toSerialize["value_mkt"] = o.ValueMkt
	toSerialize["value_now"] = o.ValueNow
	return toSerialize, nil
}

func (o *InventoryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cost_qty",
		"cost_sum",
		"make_a_per",
		"make_a_sum",
		"price_avg",
		"price_evn",
		"price_mkt",
		"price_now",
		"price_qty_sum",
		"qty_b",
		"qty_bm",
		"qty_c",
		"qty_l",
		"qty_s",
		"qty_sm",
		"rec_va_sum",
		"s_type",
		"stk_dats",
		"stk_na",
		"stk_no",
		"value_mkt",
		"value_now",
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

	varInventoryResponse := _InventoryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInventoryResponse)

	if err != nil {
		return err
	}

	*o = InventoryResponse(varInventoryResponse)

	return err
}

type NullableInventoryResponse struct {
	value *InventoryResponse
	isSet bool
}

func (v NullableInventoryResponse) Get() *InventoryResponse {
	return v.value
}

func (v *NullableInventoryResponse) Set(val *InventoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryResponse(val *InventoryResponse) *NullableInventoryResponse {
	return &NullableInventoryResponse{value: val, isSet: true}
}

func (v NullableInventoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


