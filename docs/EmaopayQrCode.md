# EmaopayQrCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MerchantId** | Pointer to **string** |  | [optional] 
**MonitorId** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** | 二维码提供商, 如: wechat, alipay | [optional] 
**UnlockAt** | Pointer to **int32** | 解锁时间, unix时间戳, 0表示未锁定 | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEmaopayQrCode

`func NewEmaopayQrCode() *EmaopayQrCode`

NewEmaopayQrCode instantiates a new EmaopayQrCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmaopayQrCodeWithDefaults

`func NewEmaopayQrCodeWithDefaults() *EmaopayQrCode`

NewEmaopayQrCodeWithDefaults instantiates a new EmaopayQrCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EmaopayQrCode) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EmaopayQrCode) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EmaopayQrCode) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EmaopayQrCode) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetContent

`func (o *EmaopayQrCode) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EmaopayQrCode) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EmaopayQrCode) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *EmaopayQrCode) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmaopayQrCode) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmaopayQrCode) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmaopayQrCode) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmaopayQrCode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *EmaopayQrCode) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EmaopayQrCode) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EmaopayQrCode) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EmaopayQrCode) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *EmaopayQrCode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmaopayQrCode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmaopayQrCode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmaopayQrCode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *EmaopayQrCode) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *EmaopayQrCode) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *EmaopayQrCode) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *EmaopayQrCode) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMonitorId

`func (o *EmaopayQrCode) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *EmaopayQrCode) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *EmaopayQrCode) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *EmaopayQrCode) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetProvider

`func (o *EmaopayQrCode) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EmaopayQrCode) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EmaopayQrCode) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EmaopayQrCode) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetUnlockAt

`func (o *EmaopayQrCode) GetUnlockAt() int32`

GetUnlockAt returns the UnlockAt field if non-nil, zero value otherwise.

### GetUnlockAtOk

`func (o *EmaopayQrCode) GetUnlockAtOk() (*int32, bool)`

GetUnlockAtOk returns a tuple with the UnlockAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockAt

`func (o *EmaopayQrCode) SetUnlockAt(v int32)`

SetUnlockAt sets UnlockAt field to given value.

### HasUnlockAt

`func (o *EmaopayQrCode) HasUnlockAt() bool`

HasUnlockAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EmaopayQrCode) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EmaopayQrCode) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EmaopayQrCode) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EmaopayQrCode) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


