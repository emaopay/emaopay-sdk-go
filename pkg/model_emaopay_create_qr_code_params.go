/*
EmaoPay API

EmaoPay API.

API version: 1.0
Contact: support@emaopay.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package emaopaysdk

import (
	"encoding/json"
)

// checks if the EmaopayCreateQrCodeParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmaopayCreateQrCodeParams{}

// EmaopayCreateQrCodeParams struct for EmaopayCreateQrCodeParams
type EmaopayCreateQrCodeParams struct {
	ImageKeys []string `json:"imageKeys,omitempty"`
	MerchantId *string `json:"merchantId,omitempty"`
	MonitorId *string `json:"monitorId,omitempty"`
}

// NewEmaopayCreateQrCodeParams instantiates a new EmaopayCreateQrCodeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmaopayCreateQrCodeParams() *EmaopayCreateQrCodeParams {
	this := EmaopayCreateQrCodeParams{}
	return &this
}

// NewEmaopayCreateQrCodeParamsWithDefaults instantiates a new EmaopayCreateQrCodeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmaopayCreateQrCodeParamsWithDefaults() *EmaopayCreateQrCodeParams {
	this := EmaopayCreateQrCodeParams{}
	return &this
}

// GetImageKeys returns the ImageKeys field value if set, zero value otherwise.
func (o *EmaopayCreateQrCodeParams) GetImageKeys() []string {
	if o == nil || IsNil(o.ImageKeys) {
		var ret []string
		return ret
	}
	return o.ImageKeys
}

// GetImageKeysOk returns a tuple with the ImageKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateQrCodeParams) GetImageKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.ImageKeys) {
		return nil, false
	}
	return o.ImageKeys, true
}

// HasImageKeys returns a boolean if a field has been set.
func (o *EmaopayCreateQrCodeParams) HasImageKeys() bool {
	if o != nil && !IsNil(o.ImageKeys) {
		return true
	}

	return false
}

// SetImageKeys gets a reference to the given []string and assigns it to the ImageKeys field.
func (o *EmaopayCreateQrCodeParams) SetImageKeys(v []string) {
	o.ImageKeys = v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *EmaopayCreateQrCodeParams) GetMerchantId() string {
	if o == nil || IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateQrCodeParams) GetMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *EmaopayCreateQrCodeParams) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *EmaopayCreateQrCodeParams) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *EmaopayCreateQrCodeParams) GetMonitorId() string {
	if o == nil || IsNil(o.MonitorId) {
		var ret string
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateQrCodeParams) GetMonitorIdOk() (*string, bool) {
	if o == nil || IsNil(o.MonitorId) {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *EmaopayCreateQrCodeParams) HasMonitorId() bool {
	if o != nil && !IsNil(o.MonitorId) {
		return true
	}

	return false
}

// SetMonitorId gets a reference to the given string and assigns it to the MonitorId field.
func (o *EmaopayCreateQrCodeParams) SetMonitorId(v string) {
	o.MonitorId = &v
}

func (o EmaopayCreateQrCodeParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmaopayCreateQrCodeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageKeys) {
		toSerialize["imageKeys"] = o.ImageKeys
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.MonitorId) {
		toSerialize["monitorId"] = o.MonitorId
	}
	return toSerialize, nil
}

type NullableEmaopayCreateQrCodeParams struct {
	value *EmaopayCreateQrCodeParams
	isSet bool
}

func (v NullableEmaopayCreateQrCodeParams) Get() *EmaopayCreateQrCodeParams {
	return v.value
}

func (v *NullableEmaopayCreateQrCodeParams) Set(val *EmaopayCreateQrCodeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEmaopayCreateQrCodeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEmaopayCreateQrCodeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmaopayCreateQrCodeParams(val *EmaopayCreateQrCodeParams) *NullableEmaopayCreateQrCodeParams {
	return &NullableEmaopayCreateQrCodeParams{value: val, isSet: true}
}

func (v NullableEmaopayCreateQrCodeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmaopayCreateQrCodeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


