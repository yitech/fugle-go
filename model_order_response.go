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

// checks if the OrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponse{}

// OrderResponse struct for OrderResponse
type OrderResponse struct {
	OrdDate string `json:"ord_date"`
	OrdTime string `json:"ord_time"`
	OrdType string `json:"ord_type"`
	OrdNo string `json:"ord_no"`
	RetCode string `json:"ret_code"`
	RetMsg string `json:"ret_msg"`
	WorkDate string `json:"work_date"`
}

type _OrderResponse OrderResponse

// NewOrderResponse instantiates a new OrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponse(ordDate string, ordTime string, ordType string, ordNo string, retCode string, retMsg string, workDate string) *OrderResponse {
	this := OrderResponse{}
	this.OrdDate = ordDate
	this.OrdTime = ordTime
	this.OrdType = ordType
	this.OrdNo = ordNo
	this.RetCode = retCode
	this.RetMsg = retMsg
	this.WorkDate = workDate
	return &this
}

// NewOrderResponseWithDefaults instantiates a new OrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseWithDefaults() *OrderResponse {
	this := OrderResponse{}
	return &this
}

// GetOrdDate returns the OrdDate field value
func (o *OrderResponse) GetOrdDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrdDate
}

// GetOrdDateOk returns a tuple with the OrdDate field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetOrdDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrdDate, true
}

// SetOrdDate sets field value
func (o *OrderResponse) SetOrdDate(v string) {
	o.OrdDate = v
}

// GetOrdTime returns the OrdTime field value
func (o *OrderResponse) GetOrdTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrdTime
}

// GetOrdTimeOk returns a tuple with the OrdTime field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetOrdTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrdTime, true
}

// SetOrdTime sets field value
func (o *OrderResponse) SetOrdTime(v string) {
	o.OrdTime = v
}

// GetOrdType returns the OrdType field value
func (o *OrderResponse) GetOrdType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrdType
}

// GetOrdTypeOk returns a tuple with the OrdType field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetOrdTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrdType, true
}

// SetOrdType sets field value
func (o *OrderResponse) SetOrdType(v string) {
	o.OrdType = v
}

// GetOrdNo returns the OrdNo field value
func (o *OrderResponse) GetOrdNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrdNo
}

// GetOrdNoOk returns a tuple with the OrdNo field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetOrdNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrdNo, true
}

// SetOrdNo sets field value
func (o *OrderResponse) SetOrdNo(v string) {
	o.OrdNo = v
}

// GetRetCode returns the RetCode field value
func (o *OrderResponse) GetRetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RetCode
}

// GetRetCodeOk returns a tuple with the RetCode field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetRetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetCode, true
}

// SetRetCode sets field value
func (o *OrderResponse) SetRetCode(v string) {
	o.RetCode = v
}

// GetRetMsg returns the RetMsg field value
func (o *OrderResponse) GetRetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RetMsg
}

// GetRetMsgOk returns a tuple with the RetMsg field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetRetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetMsg, true
}

// SetRetMsg sets field value
func (o *OrderResponse) SetRetMsg(v string) {
	o.RetMsg = v
}

// GetWorkDate returns the WorkDate field value
func (o *OrderResponse) GetWorkDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkDate
}

// GetWorkDateOk returns a tuple with the WorkDate field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetWorkDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkDate, true
}

// SetWorkDate sets field value
func (o *OrderResponse) SetWorkDate(v string) {
	o.WorkDate = v
}

func (o OrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ord_date"] = o.OrdDate
	toSerialize["ord_time"] = o.OrdTime
	toSerialize["ord_type"] = o.OrdType
	toSerialize["ord_no"] = o.OrdNo
	toSerialize["ret_code"] = o.RetCode
	toSerialize["ret_msg"] = o.RetMsg
	toSerialize["work_date"] = o.WorkDate
	return toSerialize, nil
}

func (o *OrderResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ord_date",
		"ord_time",
		"ord_type",
		"ord_no",
		"ret_code",
		"ret_msg",
		"work_date",
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

	varOrderResponse := _OrderResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderResponse)

	if err != nil {
		return err
	}

	*o = OrderResponse(varOrderResponse)

	return err
}

type NullableOrderResponse struct {
	value *OrderResponse
	isSet bool
}

func (v NullableOrderResponse) Get() *OrderResponse {
	return v.value
}

func (v *NullableOrderResponse) Set(val *OrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponse(val *OrderResponse) *NullableOrderResponse {
	return &NullableOrderResponse{value: val, isSet: true}
}

func (v NullableOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


