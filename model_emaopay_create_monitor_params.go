/*
EmaoPay API

EmaoPay API.

API version: 1.0
Contact: support@emaopay.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emaopay

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EmaopayCreateMonitorParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmaopayCreateMonitorParams{}

// EmaopayCreateMonitorParams struct for EmaopayCreateMonitorParams
type EmaopayCreateMonitorParams struct {
	MerchantId string `json:"merchantId"`
	Name string `json:"name"`
	Weight int32 `json:"weight"`
}

type _EmaopayCreateMonitorParams EmaopayCreateMonitorParams

// NewEmaopayCreateMonitorParams instantiates a new EmaopayCreateMonitorParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmaopayCreateMonitorParams(merchantId string, name string, weight int32) *EmaopayCreateMonitorParams {
	this := EmaopayCreateMonitorParams{}
	this.MerchantId = merchantId
	this.Name = name
	this.Weight = weight
	return &this
}

// NewEmaopayCreateMonitorParamsWithDefaults instantiates a new EmaopayCreateMonitorParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmaopayCreateMonitorParamsWithDefaults() *EmaopayCreateMonitorParams {
	this := EmaopayCreateMonitorParams{}
	return &this
}

// GetMerchantId returns the MerchantId field value
func (o *EmaopayCreateMonitorParams) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *EmaopayCreateMonitorParams) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *EmaopayCreateMonitorParams) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetName returns the Name field value
func (o *EmaopayCreateMonitorParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmaopayCreateMonitorParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EmaopayCreateMonitorParams) SetName(v string) {
	o.Name = v
}

// GetWeight returns the Weight field value
func (o *EmaopayCreateMonitorParams) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *EmaopayCreateMonitorParams) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *EmaopayCreateMonitorParams) SetWeight(v int32) {
	o.Weight = v
}

func (o EmaopayCreateMonitorParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmaopayCreateMonitorParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantId"] = o.MerchantId
	toSerialize["name"] = o.Name
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

func (o *EmaopayCreateMonitorParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"merchantId",
		"name",
		"weight",
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

	varEmaopayCreateMonitorParams := _EmaopayCreateMonitorParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmaopayCreateMonitorParams)

	if err != nil {
		return err
	}

	*o = EmaopayCreateMonitorParams(varEmaopayCreateMonitorParams)

	return err
}

type NullableEmaopayCreateMonitorParams struct {
	value *EmaopayCreateMonitorParams
	isSet bool
}

func (v NullableEmaopayCreateMonitorParams) Get() *EmaopayCreateMonitorParams {
	return v.value
}

func (v *NullableEmaopayCreateMonitorParams) Set(val *EmaopayCreateMonitorParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEmaopayCreateMonitorParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEmaopayCreateMonitorParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmaopayCreateMonitorParams(val *EmaopayCreateMonitorParams) *NullableEmaopayCreateMonitorParams {
	return &NullableEmaopayCreateMonitorParams{value: val, isSet: true}
}

func (v NullableEmaopayCreateMonitorParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmaopayCreateMonitorParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


