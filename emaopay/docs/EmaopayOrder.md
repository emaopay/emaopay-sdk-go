# EmaopayOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | 金额 | [optional] 
**Created** | Pointer to **int32** | 创建时间戳 | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiredAt** | Pointer to **int32** | 超时时间 | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MerchantId** | Pointer to **string** | 商户 Id | [optional] 
**MerchantTradeNo** | Pointer to **string** | 商户订单号, 对商户而言是唯一的 | [optional] 
**MonitorId** | Pointer to **string** | 监控 Id | [optional] 
**NotifyFailedCount** | Pointer to **int32** | 通知失败次数 | [optional] 
**NotifySuccessAt** | Pointer to **int32** | 通知成功时间 | [optional] 
**NotifyURL** | Pointer to **string** | 回调地址 | [optional] 
**PaidAt** | Pointer to **int32** | 更新时间戳 | [optional] 
**PaymentMethod** | Pointer to **string** | 支付方式 | [optional] 
**ProductId** | Pointer to **string** | 产品 Id, 可为空 | [optional] 
**QrCode** | Pointer to **string** | 二维码 | [optional] 
**RealAmount** | Pointer to **string** | 真实金额 | [optional] 
**ReturnURL** | Pointer to **string** | 支付成功后的跳转地址 | [optional] 
**Signature** | Pointer to **string** | 签名 | [optional] 
**Status** | Pointer to **string** | 状态 | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** | 用户 Id, 可为空，如果是匿名用户，那么就为空 | [optional] 

## Methods

### NewEmaopayOrder

`func NewEmaopayOrder() *EmaopayOrder`

NewEmaopayOrder instantiates a new EmaopayOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmaopayOrderWithDefaults

`func NewEmaopayOrderWithDefaults() *EmaopayOrder`

NewEmaopayOrderWithDefaults instantiates a new EmaopayOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EmaopayOrder) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EmaopayOrder) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EmaopayOrder) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EmaopayOrder) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreated

`func (o *EmaopayOrder) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EmaopayOrder) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EmaopayOrder) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EmaopayOrder) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmaopayOrder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmaopayOrder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmaopayOrder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmaopayOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EmaopayOrder) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EmaopayOrder) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EmaopayOrder) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EmaopayOrder) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetExpiredAt

`func (o *EmaopayOrder) GetExpiredAt() int32`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *EmaopayOrder) GetExpiredAtOk() (*int32, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *EmaopayOrder) SetExpiredAt(v int32)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *EmaopayOrder) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetId

`func (o *EmaopayOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmaopayOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmaopayOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmaopayOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *EmaopayOrder) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *EmaopayOrder) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *EmaopayOrder) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *EmaopayOrder) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMerchantTradeNo

`func (o *EmaopayOrder) GetMerchantTradeNo() string`

GetMerchantTradeNo returns the MerchantTradeNo field if non-nil, zero value otherwise.

### GetMerchantTradeNoOk

`func (o *EmaopayOrder) GetMerchantTradeNoOk() (*string, bool)`

GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeNo

`func (o *EmaopayOrder) SetMerchantTradeNo(v string)`

SetMerchantTradeNo sets MerchantTradeNo field to given value.

### HasMerchantTradeNo

`func (o *EmaopayOrder) HasMerchantTradeNo() bool`

HasMerchantTradeNo returns a boolean if a field has been set.

### GetMonitorId

`func (o *EmaopayOrder) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *EmaopayOrder) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *EmaopayOrder) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *EmaopayOrder) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetNotifyFailedCount

`func (o *EmaopayOrder) GetNotifyFailedCount() int32`

GetNotifyFailedCount returns the NotifyFailedCount field if non-nil, zero value otherwise.

### GetNotifyFailedCountOk

`func (o *EmaopayOrder) GetNotifyFailedCountOk() (*int32, bool)`

GetNotifyFailedCountOk returns a tuple with the NotifyFailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFailedCount

`func (o *EmaopayOrder) SetNotifyFailedCount(v int32)`

SetNotifyFailedCount sets NotifyFailedCount field to given value.

### HasNotifyFailedCount

`func (o *EmaopayOrder) HasNotifyFailedCount() bool`

HasNotifyFailedCount returns a boolean if a field has been set.

### GetNotifySuccessAt

`func (o *EmaopayOrder) GetNotifySuccessAt() int32`

GetNotifySuccessAt returns the NotifySuccessAt field if non-nil, zero value otherwise.

### GetNotifySuccessAtOk

`func (o *EmaopayOrder) GetNotifySuccessAtOk() (*int32, bool)`

GetNotifySuccessAtOk returns a tuple with the NotifySuccessAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySuccessAt

`func (o *EmaopayOrder) SetNotifySuccessAt(v int32)`

SetNotifySuccessAt sets NotifySuccessAt field to given value.

### HasNotifySuccessAt

`func (o *EmaopayOrder) HasNotifySuccessAt() bool`

HasNotifySuccessAt returns a boolean if a field has been set.

### GetNotifyURL

`func (o *EmaopayOrder) GetNotifyURL() string`

GetNotifyURL returns the NotifyURL field if non-nil, zero value otherwise.

### GetNotifyURLOk

`func (o *EmaopayOrder) GetNotifyURLOk() (*string, bool)`

GetNotifyURLOk returns a tuple with the NotifyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyURL

`func (o *EmaopayOrder) SetNotifyURL(v string)`

SetNotifyURL sets NotifyURL field to given value.

### HasNotifyURL

`func (o *EmaopayOrder) HasNotifyURL() bool`

HasNotifyURL returns a boolean if a field has been set.

### GetPaidAt

`func (o *EmaopayOrder) GetPaidAt() int32`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *EmaopayOrder) GetPaidAtOk() (*int32, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *EmaopayOrder) SetPaidAt(v int32)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *EmaopayOrder) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *EmaopayOrder) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *EmaopayOrder) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *EmaopayOrder) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *EmaopayOrder) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetProductId

`func (o *EmaopayOrder) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *EmaopayOrder) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *EmaopayOrder) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *EmaopayOrder) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetQrCode

`func (o *EmaopayOrder) GetQrCode() string`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *EmaopayOrder) GetQrCodeOk() (*string, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *EmaopayOrder) SetQrCode(v string)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *EmaopayOrder) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### GetRealAmount

`func (o *EmaopayOrder) GetRealAmount() string`

GetRealAmount returns the RealAmount field if non-nil, zero value otherwise.

### GetRealAmountOk

`func (o *EmaopayOrder) GetRealAmountOk() (*string, bool)`

GetRealAmountOk returns a tuple with the RealAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealAmount

`func (o *EmaopayOrder) SetRealAmount(v string)`

SetRealAmount sets RealAmount field to given value.

### HasRealAmount

`func (o *EmaopayOrder) HasRealAmount() bool`

HasRealAmount returns a boolean if a field has been set.

### GetReturnURL

`func (o *EmaopayOrder) GetReturnURL() string`

GetReturnURL returns the ReturnURL field if non-nil, zero value otherwise.

### GetReturnURLOk

`func (o *EmaopayOrder) GetReturnURLOk() (*string, bool)`

GetReturnURLOk returns a tuple with the ReturnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnURL

`func (o *EmaopayOrder) SetReturnURL(v string)`

SetReturnURL sets ReturnURL field to given value.

### HasReturnURL

`func (o *EmaopayOrder) HasReturnURL() bool`

HasReturnURL returns a boolean if a field has been set.

### GetSignature

`func (o *EmaopayOrder) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *EmaopayOrder) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *EmaopayOrder) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *EmaopayOrder) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetStatus

`func (o *EmaopayOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmaopayOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmaopayOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmaopayOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EmaopayOrder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EmaopayOrder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EmaopayOrder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EmaopayOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *EmaopayOrder) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EmaopayOrder) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EmaopayOrder) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EmaopayOrder) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


