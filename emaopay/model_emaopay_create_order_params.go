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

// checks if the EmaopayCreateOrderParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmaopayCreateOrderParams{}

// EmaopayCreateOrderParams struct for EmaopayCreateOrderParams
type EmaopayCreateOrderParams struct {
	// 金额
	Amount *string `json:"amount,omitempty"`
	// 商户 Id
	MerchantId *string `json:"merchantId,omitempty"`
	// 商户订单号, 对商户而言是唯一的
	MerchantTradeNo *string `json:"merchantTradeNo,omitempty"`
	// 回调地址
	NotifyURL *string `json:"notifyURL,omitempty"`
	// 支付方式
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	// 产品 Id
	ProductId *string `json:"productId,omitempty"`
	// 支付成功后的跳转地址
	ReturnURL *string `json:"returnURL,omitempty"`
	// 签名
	Signature *string `json:"signature,omitempty"`
	// 用户 Id, 可为空，如果是匿名用户，那么就为空
	UserId *string `json:"userId,omitempty"`
}

// NewEmaopayCreateOrderParams instantiates a new EmaopayCreateOrderParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmaopayCreateOrderParams() *EmaopayCreateOrderParams {
	this := EmaopayCreateOrderParams{}
	return &this
}

// NewEmaopayCreateOrderParamsWithDefaults instantiates a new EmaopayCreateOrderParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmaopayCreateOrderParamsWithDefaults() *EmaopayCreateOrderParams {
	this := EmaopayCreateOrderParams{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *EmaopayCreateOrderParams) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateOrderParams) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *EmaopayCreateOrderParams) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *EmaopayCreateOrderParams) SetAmount(v string) {
	o.Amount = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *EmaopayCreateOrderParams) GetMerchantId() string {
	if o == nil || IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateOrderParams) GetMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *EmaopayCreateOrderParams) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *EmaopayCreateOrderParams) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetMerchantTradeNo returns the MerchantTradeNo field value if set, zero value otherwise.
func (o *EmaopayCreateOrderParams) GetMerchantTradeNo() string {
	if o == nil || IsNil(o.MerchantTradeNo) {
		var ret string
		return ret
	}
	return *o.MerchantTradeNo
}

// GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateOrderParams) GetMerchantTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantTradeNo) {
		return nil, false
	}
	return o.MerchantTradeNo, true
}

// HasMerchantTradeNo returns a boolean if a field has been set.
func (o *EmaopayCreateOrderParams) HasMerchantTradeNo() bool {
	if o != nil && !IsNil(o.MerchantTradeNo) {
		return true
	}

	return false
}

// SetMerchantTradeNo gets a reference to the given string and assigns it to the MerchantTradeNo field.
func (o *EmaopayCreateOrderParams) SetMerchantTradeNo(v string) {
	o.MerchantTradeNo = &v
}

// GetNotifyURL returns the NotifyURL field value if set, zero value otherwise.
func (o *EmaopayCreateOrderParams) GetNotifyURL() string {
	if o == nil || IsNil(o.NotifyURL) {
		var ret string
		return ret
	}
	return *o.NotifyURL
}

// GetNotifyURLOk returns a tuple with the NotifyURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateOrderParams) GetNotifyURLOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyURL) {
		return nil, false
	}
	return o.NotifyURL, true
}

// HasNotifyURL returns a boolean if a field has been set.
func (o *EmaopayCreateOrderParams) HasNotifyURL() bool {
	if o != nil && !IsNil(o.NotifyURL) {
		return true
	}

	return false
}

// SetNotifyURL gets a reference to the given string and assigns it to the NotifyURL field.
func (o *EmaopayCreateOrderParams) SetNotifyURL(v string) {
	o.NotifyURL = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *EmaopayCreateOrderParams) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateOrderParams) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *EmaopayCreateOrderParams) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *EmaopayCreateOrderParams) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *EmaopayCreateOrderParams) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateOrderParams) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *EmaopayCreateOrderParams) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *EmaopayCreateOrderParams) SetProductId(v string) {
	o.ProductId = &v
}

// GetReturnURL returns the ReturnURL field value if set, zero value otherwise.
func (o *EmaopayCreateOrderParams) GetReturnURL() string {
	if o == nil || IsNil(o.ReturnURL) {
		var ret string
		return ret
	}
	return *o.ReturnURL
}

// GetReturnURLOk returns a tuple with the ReturnURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateOrderParams) GetReturnURLOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnURL) {
		return nil, false
	}
	return o.ReturnURL, true
}

// HasReturnURL returns a boolean if a field has been set.
func (o *EmaopayCreateOrderParams) HasReturnURL() bool {
	if o != nil && !IsNil(o.ReturnURL) {
		return true
	}

	return false
}

// SetReturnURL gets a reference to the given string and assigns it to the ReturnURL field.
func (o *EmaopayCreateOrderParams) SetReturnURL(v string) {
	o.ReturnURL = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *EmaopayCreateOrderParams) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateOrderParams) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *EmaopayCreateOrderParams) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *EmaopayCreateOrderParams) SetSignature(v string) {
	o.Signature = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *EmaopayCreateOrderParams) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmaopayCreateOrderParams) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EmaopayCreateOrderParams) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EmaopayCreateOrderParams) SetUserId(v string) {
	o.UserId = &v
}

func (o EmaopayCreateOrderParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmaopayCreateOrderParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.MerchantTradeNo) {
		toSerialize["merchantTradeNo"] = o.MerchantTradeNo
	}
	if !IsNil(o.NotifyURL) {
		toSerialize["notifyURL"] = o.NotifyURL
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.ReturnURL) {
		toSerialize["returnURL"] = o.ReturnURL
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableEmaopayCreateOrderParams struct {
	value *EmaopayCreateOrderParams
	isSet bool
}

func (v NullableEmaopayCreateOrderParams) Get() *EmaopayCreateOrderParams {
	return v.value
}

func (v *NullableEmaopayCreateOrderParams) Set(val *EmaopayCreateOrderParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEmaopayCreateOrderParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEmaopayCreateOrderParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmaopayCreateOrderParams(val *EmaopayCreateOrderParams) *NullableEmaopayCreateOrderParams {
	return &NullableEmaopayCreateOrderParams{value: val, isSet: true}
}

func (v NullableEmaopayCreateOrderParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmaopayCreateOrderParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


