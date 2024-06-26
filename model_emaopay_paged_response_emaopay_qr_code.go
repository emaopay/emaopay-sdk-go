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
)

// checks if the EmaopayPagedResponseEmaopayQrCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmaopayPagedResponseEmaopayQrCode{}

// EmaopayPagedResponseEmaopayQrCode struct for EmaopayPagedResponseEmaopayQrCode
type EmaopayPagedResponseEmaopayQrCode struct {
	Data []EmaopayQrCode `json:"data,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewEmaopayPagedResponseEmaopayQrCode instantiates a new EmaopayPagedResponseEmaopayQrCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmaopayPagedResponseEmaopayQrCode() *EmaopayPagedResponseEmaopayQrCode {
	this := EmaopayPagedResponseEmaopayQrCode{}
	return &this
}

// NewEmaopayPagedResponseEmaopayQrCodeWithDefaults instantiates a new EmaopayPagedResponseEmaopayQrCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmaopayPagedResponseEmaopayQrCodeWithDefaults() *EmaopayPagedResponseEmaopayQrCode {
	this := EmaopayPagedResponseEmaopayQrCode{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmaopayPagedResponseEmaopayQrCode) GetData() []EmaopayQrCode {
	if o == nil || IsNil(o.Data) {
		var ret []EmaopayQrCode
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayPagedResponseEmaopayQrCode) GetDataOk() ([]EmaopayQrCode, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmaopayPagedResponseEmaopayQrCode) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []EmaopayQrCode and assigns it to the Data field.
func (o *EmaopayPagedResponseEmaopayQrCode) SetData(v []EmaopayQrCode) {
	o.Data = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *EmaopayPagedResponseEmaopayQrCode) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayPagedResponseEmaopayQrCode) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *EmaopayPagedResponseEmaopayQrCode) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *EmaopayPagedResponseEmaopayQrCode) SetTotal(v int32) {
	o.Total = &v
}

func (o EmaopayPagedResponseEmaopayQrCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmaopayPagedResponseEmaopayQrCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableEmaopayPagedResponseEmaopayQrCode struct {
	value *EmaopayPagedResponseEmaopayQrCode
	isSet bool
}

func (v NullableEmaopayPagedResponseEmaopayQrCode) Get() *EmaopayPagedResponseEmaopayQrCode {
	return v.value
}

func (v *NullableEmaopayPagedResponseEmaopayQrCode) Set(val *EmaopayPagedResponseEmaopayQrCode) {
	v.value = val
	v.isSet = true
}

func (v NullableEmaopayPagedResponseEmaopayQrCode) IsSet() bool {
	return v.isSet
}

func (v *NullableEmaopayPagedResponseEmaopayQrCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmaopayPagedResponseEmaopayQrCode(val *EmaopayPagedResponseEmaopayQrCode) *NullableEmaopayPagedResponseEmaopayQrCode {
	return &NullableEmaopayPagedResponseEmaopayQrCode{value: val, isSet: true}
}

func (v NullableEmaopayPagedResponseEmaopayQrCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmaopayPagedResponseEmaopayQrCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


