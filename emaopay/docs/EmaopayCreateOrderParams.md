# EmaopayCreateOrderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | 金额 | [optional] 
**MerchantId** | Pointer to **string** | 商户 Id | [optional] 
**MerchantTradeNo** | Pointer to **string** | 商户订单号, 对商户而言是唯一的 | [optional] 
**NotifyURL** | Pointer to **string** | 回调地址 | [optional] 
**PaymentMethod** | Pointer to **string** | 支付方式 | [optional] 
**ProductId** | Pointer to **string** | 产品 Id | [optional] 
**ReturnURL** | Pointer to **string** | 支付成功后的跳转地址 | [optional] 
**Signature** | Pointer to **string** | 签名 | [optional] 
**UserId** | Pointer to **string** | 用户 Id, 可为空，如果是匿名用户，那么就为空 | [optional] 

## Methods

### NewEmaopayCreateOrderParams

`func NewEmaopayCreateOrderParams() *EmaopayCreateOrderParams`

NewEmaopayCreateOrderParams instantiates a new EmaopayCreateOrderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmaopayCreateOrderParamsWithDefaults

`func NewEmaopayCreateOrderParamsWithDefaults() *EmaopayCreateOrderParams`

NewEmaopayCreateOrderParamsWithDefaults instantiates a new EmaopayCreateOrderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EmaopayCreateOrderParams) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EmaopayCreateOrderParams) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EmaopayCreateOrderParams) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EmaopayCreateOrderParams) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMerchantId

`func (o *EmaopayCreateOrderParams) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *EmaopayCreateOrderParams) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *EmaopayCreateOrderParams) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *EmaopayCreateOrderParams) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMerchantTradeNo

`func (o *EmaopayCreateOrderParams) GetMerchantTradeNo() string`

GetMerchantTradeNo returns the MerchantTradeNo field if non-nil, zero value otherwise.

### GetMerchantTradeNoOk

`func (o *EmaopayCreateOrderParams) GetMerchantTradeNoOk() (*string, bool)`

GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeNo

`func (o *EmaopayCreateOrderParams) SetMerchantTradeNo(v string)`

SetMerchantTradeNo sets MerchantTradeNo field to given value.

### HasMerchantTradeNo

`func (o *EmaopayCreateOrderParams) HasMerchantTradeNo() bool`

HasMerchantTradeNo returns a boolean if a field has been set.

### GetNotifyURL

`func (o *EmaopayCreateOrderParams) GetNotifyURL() string`

GetNotifyURL returns the NotifyURL field if non-nil, zero value otherwise.

### GetNotifyURLOk

`func (o *EmaopayCreateOrderParams) GetNotifyURLOk() (*string, bool)`

GetNotifyURLOk returns a tuple with the NotifyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyURL

`func (o *EmaopayCreateOrderParams) SetNotifyURL(v string)`

SetNotifyURL sets NotifyURL field to given value.

### HasNotifyURL

`func (o *EmaopayCreateOrderParams) HasNotifyURL() bool`

HasNotifyURL returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *EmaopayCreateOrderParams) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *EmaopayCreateOrderParams) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *EmaopayCreateOrderParams) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *EmaopayCreateOrderParams) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetProductId

`func (o *EmaopayCreateOrderParams) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *EmaopayCreateOrderParams) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *EmaopayCreateOrderParams) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *EmaopayCreateOrderParams) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetReturnURL

`func (o *EmaopayCreateOrderParams) GetReturnURL() string`

GetReturnURL returns the ReturnURL field if non-nil, zero value otherwise.

### GetReturnURLOk

`func (o *EmaopayCreateOrderParams) GetReturnURLOk() (*string, bool)`

GetReturnURLOk returns a tuple with the ReturnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnURL

`func (o *EmaopayCreateOrderParams) SetReturnURL(v string)`

SetReturnURL sets ReturnURL field to given value.

### HasReturnURL

`func (o *EmaopayCreateOrderParams) HasReturnURL() bool`

HasReturnURL returns a boolean if a field has been set.

### GetSignature

`func (o *EmaopayCreateOrderParams) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *EmaopayCreateOrderParams) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *EmaopayCreateOrderParams) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *EmaopayCreateOrderParams) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetUserId

`func (o *EmaopayCreateOrderParams) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EmaopayCreateOrderParams) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EmaopayCreateOrderParams) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EmaopayCreateOrderParams) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


