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

// checks if the MatDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatDetail{}

// MatDetail struct for MatDetail
type MatDetail struct {
	BuySell string `json:"buy_sell"`
	CDate string `json:"c_date"`
	DbFee *string `json:"db_fee,omitempty"`
	Fee string `json:"fee"`
	Make int32 `json:"make"`
	MakePer float32 `json:"make_per"`
	OrderNo string `json:"order_no"`
	PayN string `json:"pay_n"`
	Price string `json:"price"`
	PriceQty string `json:"price_qty"`
	Qty string `json:"qty"`
	SType string `json:"s_type"`
	StkNa string `json:"stk_na"`
	StkNo string `json:"stk_no"`
	TDate string `json:"t_date"`
	TTime *string `json:"t_time,omitempty"`
	Tax string `json:"tax"`
	TaxG string `json:"tax_g"`
	Trade int32 `json:"trade"`
	UserDef *string `json:"user_def,omitempty"`
}

type _MatDetail MatDetail

// NewMatDetail instantiates a new MatDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatDetail(buySell string, cDate string, fee string, make int32, makePer float32, orderNo string, payN string, price string, priceQty string, qty string, sType string, stkNa string, stkNo string, tDate string, tax string, taxG string, trade int32) *MatDetail {
	this := MatDetail{}
	this.BuySell = buySell
	this.CDate = cDate
	var dbFee string = "0"
	this.DbFee = &dbFee
	this.Fee = fee
	this.Make = make
	this.MakePer = makePer
	this.OrderNo = orderNo
	this.PayN = payN
	this.Price = price
	this.PriceQty = priceQty
	this.Qty = qty
	this.SType = sType
	this.StkNa = stkNa
	this.StkNo = stkNo
	this.TDate = tDate
	var tTime string = ""
	this.TTime = &tTime
	this.Tax = tax
	this.TaxG = taxG
	this.Trade = trade
	var userDef string = ""
	this.UserDef = &userDef
	return &this
}

// NewMatDetailWithDefaults instantiates a new MatDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatDetailWithDefaults() *MatDetail {
	this := MatDetail{}
	var dbFee string = "0"
	this.DbFee = &dbFee
	var tTime string = ""
	this.TTime = &tTime
	var userDef string = ""
	this.UserDef = &userDef
	return &this
}

// GetBuySell returns the BuySell field value
func (o *MatDetail) GetBuySell() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuySell
}

// GetBuySellOk returns a tuple with the BuySell field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetBuySellOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuySell, true
}

// SetBuySell sets field value
func (o *MatDetail) SetBuySell(v string) {
	o.BuySell = v
}

// GetCDate returns the CDate field value
func (o *MatDetail) GetCDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CDate
}

// GetCDateOk returns a tuple with the CDate field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetCDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CDate, true
}

// SetCDate sets field value
func (o *MatDetail) SetCDate(v string) {
	o.CDate = v
}

// GetDbFee returns the DbFee field value if set, zero value otherwise.
func (o *MatDetail) GetDbFee() string {
	if o == nil || IsNil(o.DbFee) {
		var ret string
		return ret
	}
	return *o.DbFee
}

// GetDbFeeOk returns a tuple with the DbFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatDetail) GetDbFeeOk() (*string, bool) {
	if o == nil || IsNil(o.DbFee) {
		return nil, false
	}
	return o.DbFee, true
}

// HasDbFee returns a boolean if a field has been set.
func (o *MatDetail) HasDbFee() bool {
	if o != nil && !IsNil(o.DbFee) {
		return true
	}

	return false
}

// SetDbFee gets a reference to the given string and assigns it to the DbFee field.
func (o *MatDetail) SetDbFee(v string) {
	o.DbFee = &v
}

// GetFee returns the Fee field value
func (o *MatDetail) GetFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *MatDetail) SetFee(v string) {
	o.Fee = v
}

// GetMake returns the Make field value
func (o *MatDetail) GetMake() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Make
}

// GetMakeOk returns a tuple with the Make field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetMakeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Make, true
}

// SetMake sets field value
func (o *MatDetail) SetMake(v int32) {
	o.Make = v
}

// GetMakePer returns the MakePer field value
func (o *MatDetail) GetMakePer() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MakePer
}

// GetMakePerOk returns a tuple with the MakePer field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetMakePerOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MakePer, true
}

// SetMakePer sets field value
func (o *MatDetail) SetMakePer(v float32) {
	o.MakePer = v
}

// GetOrderNo returns the OrderNo field value
func (o *MatDetail) GetOrderNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetOrderNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderNo, true
}

// SetOrderNo sets field value
func (o *MatDetail) SetOrderNo(v string) {
	o.OrderNo = v
}

// GetPayN returns the PayN field value
func (o *MatDetail) GetPayN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayN
}

// GetPayNOk returns a tuple with the PayN field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetPayNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayN, true
}

// SetPayN sets field value
func (o *MatDetail) SetPayN(v string) {
	o.PayN = v
}

// GetPrice returns the Price field value
func (o *MatDetail) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *MatDetail) SetPrice(v string) {
	o.Price = v
}

// GetPriceQty returns the PriceQty field value
func (o *MatDetail) GetPriceQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PriceQty
}

// GetPriceQtyOk returns a tuple with the PriceQty field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetPriceQtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceQty, true
}

// SetPriceQty sets field value
func (o *MatDetail) SetPriceQty(v string) {
	o.PriceQty = v
}

// GetQty returns the Qty field value
func (o *MatDetail) GetQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetQtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *MatDetail) SetQty(v string) {
	o.Qty = v
}

// GetSType returns the SType field value
func (o *MatDetail) GetSType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SType
}

// GetSTypeOk returns a tuple with the SType field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetSTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SType, true
}

// SetSType sets field value
func (o *MatDetail) SetSType(v string) {
	o.SType = v
}

// GetStkNa returns the StkNa field value
func (o *MatDetail) GetStkNa() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StkNa
}

// GetStkNaOk returns a tuple with the StkNa field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetStkNaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StkNa, true
}

// SetStkNa sets field value
func (o *MatDetail) SetStkNa(v string) {
	o.StkNa = v
}

// GetStkNo returns the StkNo field value
func (o *MatDetail) GetStkNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StkNo
}

// GetStkNoOk returns a tuple with the StkNo field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetStkNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StkNo, true
}

// SetStkNo sets field value
func (o *MatDetail) SetStkNo(v string) {
	o.StkNo = v
}

// GetTDate returns the TDate field value
func (o *MatDetail) GetTDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TDate
}

// GetTDateOk returns a tuple with the TDate field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetTDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TDate, true
}

// SetTDate sets field value
func (o *MatDetail) SetTDate(v string) {
	o.TDate = v
}

// GetTTime returns the TTime field value if set, zero value otherwise.
func (o *MatDetail) GetTTime() string {
	if o == nil || IsNil(o.TTime) {
		var ret string
		return ret
	}
	return *o.TTime
}

// GetTTimeOk returns a tuple with the TTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatDetail) GetTTimeOk() (*string, bool) {
	if o == nil || IsNil(o.TTime) {
		return nil, false
	}
	return o.TTime, true
}

// HasTTime returns a boolean if a field has been set.
func (o *MatDetail) HasTTime() bool {
	if o != nil && !IsNil(o.TTime) {
		return true
	}

	return false
}

// SetTTime gets a reference to the given string and assigns it to the TTime field.
func (o *MatDetail) SetTTime(v string) {
	o.TTime = &v
}

// GetTax returns the Tax field value
func (o *MatDetail) GetTax() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tax
}

// GetTaxOk returns a tuple with the Tax field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetTaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tax, true
}

// SetTax sets field value
func (o *MatDetail) SetTax(v string) {
	o.Tax = v
}

// GetTaxG returns the TaxG field value
func (o *MatDetail) GetTaxG() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxG
}

// GetTaxGOk returns a tuple with the TaxG field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetTaxGOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxG, true
}

// SetTaxG sets field value
func (o *MatDetail) SetTaxG(v string) {
	o.TaxG = v
}

// GetTrade returns the Trade field value
func (o *MatDetail) GetTrade() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Trade
}

// GetTradeOk returns a tuple with the Trade field value
// and a boolean to check if the value has been set.
func (o *MatDetail) GetTradeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trade, true
}

// SetTrade sets field value
func (o *MatDetail) SetTrade(v int32) {
	o.Trade = v
}

// GetUserDef returns the UserDef field value if set, zero value otherwise.
func (o *MatDetail) GetUserDef() string {
	if o == nil || IsNil(o.UserDef) {
		var ret string
		return ret
	}
	return *o.UserDef
}

// GetUserDefOk returns a tuple with the UserDef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatDetail) GetUserDefOk() (*string, bool) {
	if o == nil || IsNil(o.UserDef) {
		return nil, false
	}
	return o.UserDef, true
}

// HasUserDef returns a boolean if a field has been set.
func (o *MatDetail) HasUserDef() bool {
	if o != nil && !IsNil(o.UserDef) {
		return true
	}

	return false
}

// SetUserDef gets a reference to the given string and assigns it to the UserDef field.
func (o *MatDetail) SetUserDef(v string) {
	o.UserDef = &v
}

func (o MatDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buy_sell"] = o.BuySell
	toSerialize["c_date"] = o.CDate
	if !IsNil(o.DbFee) {
		toSerialize["db_fee"] = o.DbFee
	}
	toSerialize["fee"] = o.Fee
	toSerialize["make"] = o.Make
	toSerialize["make_per"] = o.MakePer
	toSerialize["order_no"] = o.OrderNo
	toSerialize["pay_n"] = o.PayN
	toSerialize["price"] = o.Price
	toSerialize["price_qty"] = o.PriceQty
	toSerialize["qty"] = o.Qty
	toSerialize["s_type"] = o.SType
	toSerialize["stk_na"] = o.StkNa
	toSerialize["stk_no"] = o.StkNo
	toSerialize["t_date"] = o.TDate
	if !IsNil(o.TTime) {
		toSerialize["t_time"] = o.TTime
	}
	toSerialize["tax"] = o.Tax
	toSerialize["tax_g"] = o.TaxG
	toSerialize["trade"] = o.Trade
	if !IsNil(o.UserDef) {
		toSerialize["user_def"] = o.UserDef
	}
	return toSerialize, nil
}

func (o *MatDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"buy_sell",
		"c_date",
		"fee",
		"make",
		"make_per",
		"order_no",
		"pay_n",
		"price",
		"price_qty",
		"qty",
		"s_type",
		"stk_na",
		"stk_no",
		"t_date",
		"tax",
		"tax_g",
		"trade",
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

	varMatDetail := _MatDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMatDetail)

	if err != nil {
		return err
	}

	*o = MatDetail(varMatDetail)

	return err
}

type NullableMatDetail struct {
	value *MatDetail
	isSet bool
}

func (v NullableMatDetail) Get() *MatDetail {
	return v.value
}

func (v *NullableMatDetail) Set(val *MatDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMatDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMatDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatDetail(val *MatDetail) *NullableMatDetail {
	return &NullableMatDetail{value: val, isSet: true}
}

func (v NullableMatDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


