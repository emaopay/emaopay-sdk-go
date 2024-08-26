# Go API client for emaopay

EmaoPay API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.5.0
- Generator version: 7.6.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://emaopay.com/support](https://emaopay.com/support)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import emaopay "github.com/emaopay/emaopay-sdk-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `emaopay.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), emaopay.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `emaopay.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), emaopay.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `emaopay.ContextOperationServerIndices` and `emaopay.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), emaopay.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), emaopay.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://emaopay.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CheckServiceStatus**](docs/DefaultApi.md#checkservicestatus) | **Get** /api/monitors/messages | 检查服务状态
*DefaultApi* | [**CreateMerchant**](docs/DefaultApi.md#createmerchant) | **Post** /api/merchants | 创建商户
*DefaultApi* | [**CreateMonitor**](docs/DefaultApi.md#createmonitor) | **Post** /api/monitors | 创建监控
*DefaultApi* | [**CreateNewMerchant**](docs/DefaultApi.md#createnewmerchant) | **Post** /api/merchants/new | 创建新商户
*DefaultApi* | [**CreateNewMonitor**](docs/DefaultApi.md#createnewmonitor) | **Post** /api/monitors/new | 创建新监控
*DefaultApi* | [**CreateOrder**](docs/DefaultApi.md#createorder) | **Post** /api/orders | 创建订单
*DefaultApi* | [**CreateQrCode**](docs/DefaultApi.md#createqrcode) | **Post** /api/qrcodes | 创建二维码
*DefaultApi* | [**DeleteMonitorById**](docs/DefaultApi.md#deletemonitorbyid) | **Delete** /api/monitors/{id} | 删除监控
*DefaultApi* | [**DeleteQrCodeById**](docs/DefaultApi.md#deleteqrcodebyid) | **Delete** /api/qrcodes | 删除二维码
*DefaultApi* | [**FinishedOrderById**](docs/DefaultApi.md#finishedorderbyid) | **Post** /api/orders/finish/{id} | 手动完成订单
*DefaultApi* | [**GetMerchantById**](docs/DefaultApi.md#getmerchantbyid) | **Get** /api/merchants/{id} | 获取商户信息
*DefaultApi* | [**GetMerchantListByUserId**](docs/DefaultApi.md#getmerchantlistbyuserid) | **Get** /api/merchants | 获取商户列表
*DefaultApi* | [**GetMonitorById**](docs/DefaultApi.md#getmonitorbyid) | **Get** /api/monitors/{id} | 获取监控
*DefaultApi* | [**GetMonitorListByMerchantId**](docs/DefaultApi.md#getmonitorlistbymerchantid) | **Get** /api/monitors | 获取监控列表
*DefaultApi* | [**GetMonitorMessageListByMerchantId**](docs/DefaultApi.md#getmonitormessagelistbymerchantid) | **Get** /api/merchants/:merchantId/messages | 获取监控消息列表（商户ID）
*DefaultApi* | [**GetMonitorMessageListByMonitorId**](docs/DefaultApi.md#getmonitormessagelistbymonitorid) | **Get** /api/monitors/:monitorId/messages | 获取监控消息列表（商户监控）
*DefaultApi* | [**GetOrderById**](docs/DefaultApi.md#getorderbyid) | **Get** /api/orders/{id} | 获取订单信息
*DefaultApi* | [**GetOrderListByUserId**](docs/DefaultApi.md#getorderlistbyuserid) | **Get** /api/orders/user | 获取订单列表
*DefaultApi* | [**GetPagedOrderList**](docs/DefaultApi.md#getpagedorderlist) | **Get** /api/orders | 获取订单列表
*DefaultApi* | [**GetPaymentProviders**](docs/DefaultApi.md#getpaymentproviders) | **Get** /api/orders/payment-providers | 获取支付方式
*DefaultApi* | [**GetQrCodePagedListByMonitorId**](docs/DefaultApi.md#getqrcodepagedlistbymonitorid) | **Get** /api/qrcodes | 获取二维码列表
*DefaultApi* | [**GetQrCodeUploadCredits**](docs/DefaultApi.md#getqrcodeuploadcredits) | **Get** /api/qrcodes/upload-credits | 获取二维码上传凭证
*DefaultApi* | [**GetUserInfo**](docs/DefaultApi.md#getuserinfo) | **Get** /api/user/info | 获取用户信息
*DefaultApi* | [**Login**](docs/DefaultApi.md#login) | **Post** /api/user/login | 登录
*DefaultApi* | [**NotifyOrderById**](docs/DefaultApi.md#notifyorderbyid) | **Post** /api/orders/notify/{id} | 通知订单回调
*DefaultApi* | [**PushMessage**](docs/DefaultApi.md#pushmessage) | **Post** /api/monitors/messages | 推送消息
*DefaultApi* | [**RefreshMonitorApiToken**](docs/DefaultApi.md#refreshmonitorapitoken) | **Put** /api/monitors/{monitorId}/refresh-api-token | 刷新ApiToken
*DefaultApi* | [**Register**](docs/DefaultApi.md#register) | **Post** /api/user/register | 注册
*DefaultApi* | [**UpdateMerchant**](docs/DefaultApi.md#updatemerchant) | **Put** /api/merchants/{id} | 更新商户
*DefaultApi* | [**UpdateMonitor**](docs/DefaultApi.md#updatemonitor) | **Put** /api/monitors/{id} | 更新监控


## Documentation For Models

 - [EmaopayCreateMerchantParams](docs/EmaopayCreateMerchantParams.md)
 - [EmaopayCreateMonitorParams](docs/EmaopayCreateMonitorParams.md)
 - [EmaopayCreateNewMonitorParams](docs/EmaopayCreateNewMonitorParams.md)
 - [EmaopayCreateOrderParams](docs/EmaopayCreateOrderParams.md)
 - [EmaopayCreateQrCodeParams](docs/EmaopayCreateQrCodeParams.md)
 - [EmaopayCredentials](docs/EmaopayCredentials.md)
 - [EmaopayDeleteQrCodeParams](docs/EmaopayDeleteQrCodeParams.md)
 - [EmaopayLoginParams](docs/EmaopayLoginParams.md)
 - [EmaopayLoginResponse](docs/EmaopayLoginResponse.md)
 - [EmaopayMerchant](docs/EmaopayMerchant.md)
 - [EmaopayMonitor](docs/EmaopayMonitor.md)
 - [EmaopayMonitorMessage](docs/EmaopayMonitorMessage.md)
 - [EmaopayMonitorMessageParams](docs/EmaopayMonitorMessageParams.md)
 - [EmaopayOrder](docs/EmaopayOrder.md)
 - [EmaopayPagedResponseEmaopayMonitorMessage](docs/EmaopayPagedResponseEmaopayMonitorMessage.md)
 - [EmaopayPagedResponseEmaopayOrder](docs/EmaopayPagedResponseEmaopayOrder.md)
 - [EmaopayPagedResponseEmaopayQrCode](docs/EmaopayPagedResponseEmaopayQrCode.md)
 - [EmaopayQrCode](docs/EmaopayQrCode.md)
 - [EmaopayRegisterParams](docs/EmaopayRegisterParams.md)
 - [EmaopayUploadCredits](docs/EmaopayUploadCredits.md)
 - [EmaopayUser](docs/EmaopayUser.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@emaopay.com

