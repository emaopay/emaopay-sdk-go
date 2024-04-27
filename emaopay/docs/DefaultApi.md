# \DefaultApi

All URIs are relative to *http://emaopay.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckServiceStatus**](DefaultApi.md#CheckServiceStatus) | **Get** /api/monitors/messages | 检查服务状态
[**CreateMerchant**](DefaultApi.md#CreateMerchant) | **Post** /api/merchants | 创建商户
[**CreateMonitor**](DefaultApi.md#CreateMonitor) | **Post** /api/monitors | 创建监控
[**CreateNewMerchant**](DefaultApi.md#CreateNewMerchant) | **Post** /api/merchants/new | 创建新商户
[**CreateNewMonitor**](DefaultApi.md#CreateNewMonitor) | **Post** /api/monitors/new | 创建新监控
[**CreateOrder**](DefaultApi.md#CreateOrder) | **Post** /api/orders | 创建订单
[**CreateQrCode**](DefaultApi.md#CreateQrCode) | **Post** /api/qrcodes | 创建二维码
[**DeleteMonitorById**](DefaultApi.md#DeleteMonitorById) | **Delete** /api/monitors/{id} | 删除监控
[**DeleteQrCodeById**](DefaultApi.md#DeleteQrCodeById) | **Delete** /api/qrcodes | 删除二维码
[**FinishedOrderById**](DefaultApi.md#FinishedOrderById) | **Post** /api/orders/finish/{id} | 手动完成订单
[**GetMerchantById**](DefaultApi.md#GetMerchantById) | **Get** /api/merchants/{id} | 获取商户信息
[**GetMerchantListByUserId**](DefaultApi.md#GetMerchantListByUserId) | **Get** /api/merchants | 获取商户列表
[**GetMonitorById**](DefaultApi.md#GetMonitorById) | **Get** /api/monitors/{id} | 获取监控
[**GetMonitorListByMerchantId**](DefaultApi.md#GetMonitorListByMerchantId) | **Get** /api/monitors | 获取监控列表
[**GetMonitorMessageListByMerchantId**](DefaultApi.md#GetMonitorMessageListByMerchantId) | **Get** /api/merchants/:merchantId/messages | 获取监控消息列表（商户ID）
[**GetMonitorMessageListByMonitorId**](DefaultApi.md#GetMonitorMessageListByMonitorId) | **Get** /api/monitors/:monitorId/messages | 获取监控消息列表（商户监控）
[**GetOrderById**](DefaultApi.md#GetOrderById) | **Get** /api/orders/{id} | 获取订单信息
[**GetOrderListByUserId**](DefaultApi.md#GetOrderListByUserId) | **Get** /api/orders/user | 获取订单列表
[**GetPagedOrderList**](DefaultApi.md#GetPagedOrderList) | **Get** /api/orders | 获取订单列表
[**GetPaymentProviders**](DefaultApi.md#GetPaymentProviders) | **Get** /api/orders/payment-providers | 获取支付方式
[**GetQrCodePagedListByMonitorId**](DefaultApi.md#GetQrCodePagedListByMonitorId) | **Get** /api/qrcodes | 获取二维码列表
[**GetQrCodeUploadCredits**](DefaultApi.md#GetQrCodeUploadCredits) | **Get** /api/qrcodes/upload-credits | 获取二维码上传凭证
[**GetUserInfo**](DefaultApi.md#GetUserInfo) | **Get** /api/user/info | 获取用户信息
[**Login**](DefaultApi.md#Login) | **Post** /api/user/login | 登录
[**NotifyOrderById**](DefaultApi.md#NotifyOrderById) | **Post** /api/orders/notify/{id} | 通知订单回调
[**PushMessage**](DefaultApi.md#PushMessage) | **Post** /api/monitors/messages | 推送消息
[**RefreshMonitorApiToken**](DefaultApi.md#RefreshMonitorApiToken) | **Put** /api/monitors/{monitorId}/refresh-api-token | 刷新ApiToken
[**Register**](DefaultApi.md#Register) | **Post** /api/user/register | 注册
[**UpdateMerchant**](DefaultApi.md#UpdateMerchant) | **Put** /api/merchants/{id} | 更新商户
[**UpdateMonitor**](DefaultApi.md#UpdateMonitor) | **Put** /api/monitors/{id} | 更新监控



## CheckServiceStatus

> CheckServiceStatus(ctx).XMonitorId(xMonitorId).XToken(xToken).Execute()

检查服务状态



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	xMonitorId := "xMonitorId_example" // string | 监控Id
	xToken := "xToken_example" // string | ApiToken

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.CheckServiceStatus(context.Background()).XMonitorId(xMonitorId).XToken(xToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CheckServiceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckServiceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xMonitorId** | **string** | 监控Id | 
 **xToken** | **string** | ApiToken | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMerchant

> EmaopayMerchant CreateMerchant(ctx).Merchant(merchant).Execute()

创建商户



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	merchant := *openapiclient.NewEmaopayCreateMerchantParams("Name_example") // EmaopayCreateMerchantParams | 商户

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateMerchant(context.Background()).Merchant(merchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMerchant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMerchant`: EmaopayMerchant
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMerchant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMerchantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchant** | [**EmaopayCreateMerchantParams**](EmaopayCreateMerchantParams.md) | 商户 | 

### Return type

[**EmaopayMerchant**](EmaopayMerchant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMonitor

> EmaopayMonitor CreateMonitor(ctx).Monitor(monitor).Execute()

创建监控



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	monitor := *openapiclient.NewEmaopayCreateMonitorParams("MerchantId_example", "Name_example", int32(123)) // EmaopayCreateMonitorParams | 监控

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateMonitor(context.Background()).Monitor(monitor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMonitor`: EmaopayMonitor
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitor** | [**EmaopayCreateMonitorParams**](EmaopayCreateMonitorParams.md) | 监控 | 

### Return type

[**EmaopayMonitor**](EmaopayMonitor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewMerchant

> EmaopayMerchant CreateNewMerchant(ctx).Execute()

创建新商户



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateNewMerchant(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNewMerchant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewMerchant`: EmaopayMerchant
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNewMerchant`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewMerchantRequest struct via the builder pattern


### Return type

[**EmaopayMerchant**](EmaopayMerchant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewMonitor

> EmaopayMonitor CreateNewMonitor(ctx).Monitor(monitor).Execute()

创建新监控



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	monitor := *openapiclient.NewEmaopayCreateNewMonitorParams("MerchantId_example") // EmaopayCreateNewMonitorParams | 监控

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateNewMonitor(context.Background()).Monitor(monitor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNewMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewMonitor`: EmaopayMonitor
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNewMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitor** | [**EmaopayCreateNewMonitorParams**](EmaopayCreateNewMonitorParams.md) | 监控 | 

### Return type

[**EmaopayMonitor**](EmaopayMonitor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrder

> EmaopayOrder CreateOrder(ctx).XSignature(xSignature).Params(params).Execute()

创建订单



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	xSignature := "xSignature_example" // string | 签名
	params := *openapiclient.NewEmaopayCreateOrderParams() // EmaopayCreateOrderParams | 订单

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateOrder(context.Background()).XSignature(xSignature).Params(params).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrder`: EmaopayOrder
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSignature** | **string** | 签名 | 
 **params** | [**EmaopayCreateOrderParams**](EmaopayCreateOrderParams.md) | 订单 | 

### Return type

[**EmaopayOrder**](EmaopayOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQrCode

> EmaopayQrCode CreateQrCode(ctx).QrCode(qrCode).Execute()

创建二维码



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	qrCode := *openapiclient.NewEmaopayCreateQrCodeParams() // EmaopayCreateQrCodeParams | 二维码

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateQrCode(context.Background()).QrCode(qrCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateQrCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQrCode`: EmaopayQrCode
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateQrCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQrCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qrCode** | [**EmaopayCreateQrCodeParams**](EmaopayCreateQrCodeParams.md) | 二维码 | 

### Return type

[**EmaopayQrCode**](EmaopayQrCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitorById

> DeleteMonitorById(ctx, id).Execute()

删除监控



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	id := "id_example" // string | 监控Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteMonitorById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMonitorById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 监控Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQrCodeById

> DeleteQrCodeById(ctx).QrCode(qrCode).Execute()

删除二维码



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	qrCode := *openapiclient.NewEmaopayDeleteQrCodeParams() // EmaopayDeleteQrCodeParams | 二维码

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteQrCodeById(context.Background()).QrCode(qrCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteQrCodeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQrCodeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qrCode** | [**EmaopayDeleteQrCodeParams**](EmaopayDeleteQrCodeParams.md) | 二维码 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinishedOrderById

> EmaopayOrder FinishedOrderById(ctx, id).Execute()

手动完成订单



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	id := "id_example" // string | 订单Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.FinishedOrderById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FinishedOrderById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinishedOrderById`: EmaopayOrder
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FinishedOrderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 订单Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinishedOrderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmaopayOrder**](EmaopayOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantById

> EmaopayMerchant GetMerchantById(ctx, id).Execute()

获取商户信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	id := "id_example" // string | 商户Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetMerchantById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMerchantById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMerchantById`: EmaopayMerchant
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMerchantById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 商户Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmaopayMerchant**](EmaopayMerchant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantListByUserId

> []EmaopayMerchant GetMerchantListByUserId(ctx).Execute()

获取商户列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetMerchantListByUserId(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMerchantListByUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMerchantListByUserId`: []EmaopayMerchant
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMerchantListByUserId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantListByUserIdRequest struct via the builder pattern


### Return type

[**[]EmaopayMerchant**](EmaopayMerchant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorById

> EmaopayMonitor GetMonitorById(ctx, id).Execute()

获取监控



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	id := "id_example" // string | 监控Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetMonitorById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMonitorById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonitorById`: EmaopayMonitor
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMonitorById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 监控Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmaopayMonitor**](EmaopayMonitor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorListByMerchantId

> []EmaopayMonitor GetMonitorListByMerchantId(ctx).MerchantId(merchantId).Execute()

获取监控列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	merchantId := "merchantId_example" // string | 商户Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetMonitorListByMerchantId(context.Background()).MerchantId(merchantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMonitorListByMerchantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonitorListByMerchantId`: []EmaopayMonitor
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMonitorListByMerchantId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorListByMerchantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantId** | **string** | 商户Id | 

### Return type

[**[]EmaopayMonitor**](EmaopayMonitor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorMessageListByMerchantId

> EmaopayPagedResponseEmaopayMonitorMessage GetMonitorMessageListByMerchantId(ctx, merchantId).PageIndex(pageIndex).PageSize(pageSize).Execute()

获取监控消息列表（商户ID）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	merchantId := "merchantId_example" // string | 商户Id
	pageIndex := int32(56) // int32 | 页码
	pageSize := int32(56) // int32 | 每页数量

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetMonitorMessageListByMerchantId(context.Background(), merchantId).PageIndex(pageIndex).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMonitorMessageListByMerchantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonitorMessageListByMerchantId`: EmaopayPagedResponseEmaopayMonitorMessage
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMonitorMessageListByMerchantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | 商户Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorMessageListByMerchantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **int32** | 页码 | 
 **pageSize** | **int32** | 每页数量 | 

### Return type

[**EmaopayPagedResponseEmaopayMonitorMessage**](EmaopayPagedResponseEmaopayMonitorMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorMessageListByMonitorId

> EmaopayPagedResponseEmaopayMonitorMessage GetMonitorMessageListByMonitorId(ctx, monitorId).PageIndex(pageIndex).PageSize(pageSize).Execute()

获取监控消息列表（商户监控）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	monitorId := "monitorId_example" // string | 监控Id
	pageIndex := int32(56) // int32 | 页码
	pageSize := int32(56) // int32 | 每页数量

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetMonitorMessageListByMonitorId(context.Background(), monitorId).PageIndex(pageIndex).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMonitorMessageListByMonitorId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonitorMessageListByMonitorId`: EmaopayPagedResponseEmaopayMonitorMessage
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMonitorMessageListByMonitorId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** | 监控Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorMessageListByMonitorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **int32** | 页码 | 
 **pageSize** | **int32** | 每页数量 | 

### Return type

[**EmaopayPagedResponseEmaopayMonitorMessage**](EmaopayPagedResponseEmaopayMonitorMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderById

> EmaopayOrder GetOrderById(ctx, id).Execute()

获取订单信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	id := "id_example" // string | 订单Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetOrderById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOrderById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderById`: EmaopayOrder
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOrderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 订单Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmaopayOrder**](EmaopayOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderListByUserId

> []EmaopayOrder GetOrderListByUserId(ctx).Execute()

获取订单列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetOrderListByUserId(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOrderListByUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderListByUserId`: []EmaopayOrder
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOrderListByUserId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderListByUserIdRequest struct via the builder pattern


### Return type

[**[]EmaopayOrder**](EmaopayOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagedOrderList

> []EmaopayPagedResponseEmaopayOrder GetPagedOrderList(ctx, status, merchantId, productId).PageIndex(pageIndex).PageSize(pageSize).Execute()

获取订单列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	status := "status_example" // string | 订单状态
	merchantId := "merchantId_example" // string | 商户Id
	productId := "productId_example" // string | 产品Id
	pageIndex := int32(56) // int32 | 页码
	pageSize := int32(56) // int32 | 每页数量

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetPagedOrderList(context.Background(), status, merchantId, productId).PageIndex(pageIndex).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPagedOrderList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagedOrderList`: []EmaopayPagedResponseEmaopayOrder
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPagedOrderList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**status** | **string** | 订单状态 | 
**merchantId** | **string** | 商户Id | 
**productId** | **string** | 产品Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagedOrderListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageIndex** | **int32** | 页码 | 
 **pageSize** | **int32** | 每页数量 | 

### Return type

[**[]EmaopayPagedResponseEmaopayOrder**](EmaopayPagedResponseEmaopayOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentProviders

> []string GetPaymentProviders(ctx).MerchantId(merchantId).Execute()

获取支付方式



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	merchantId := "merchantId_example" // string | 订单Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetPaymentProviders(context.Background()).MerchantId(merchantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPaymentProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentProviders`: []string
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPaymentProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantId** | **string** | 订单Id | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQrCodePagedListByMonitorId

> EmaopayPagedResponseEmaopayQrCode GetQrCodePagedListByMonitorId(ctx).MonitorId(monitorId).PageIndex(pageIndex).PageSize(pageSize).Execute()

获取二维码列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	monitorId := "monitorId_example" // string | 商户Id
	pageIndex := int32(56) // int32 | 页码 (optional)
	pageSize := int32(56) // int32 | 每页数量 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetQrCodePagedListByMonitorId(context.Background()).MonitorId(monitorId).PageIndex(pageIndex).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetQrCodePagedListByMonitorId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQrCodePagedListByMonitorId`: EmaopayPagedResponseEmaopayQrCode
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetQrCodePagedListByMonitorId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQrCodePagedListByMonitorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitorId** | **string** | 商户Id | 
 **pageIndex** | **int32** | 页码 | 
 **pageSize** | **int32** | 每页数量 | 

### Return type

[**EmaopayPagedResponseEmaopayQrCode**](EmaopayPagedResponseEmaopayQrCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQrCodeUploadCredits

> EmaopayUploadCredits GetQrCodeUploadCredits(ctx).MerchantId(merchantId).Execute()

获取二维码上传凭证



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	merchantId := "merchantId_example" // string | 商户Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetQrCodeUploadCredits(context.Background()).MerchantId(merchantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetQrCodeUploadCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQrCodeUploadCredits`: EmaopayUploadCredits
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetQrCodeUploadCredits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQrCodeUploadCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantId** | **string** | 商户Id | 

### Return type

[**EmaopayUploadCredits**](EmaopayUploadCredits.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInfo

> EmaopayUser GetUserInfo(ctx).Execute()

获取用户信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInfo`: EmaopayUser
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoRequest struct via the builder pattern


### Return type

[**EmaopayUser**](EmaopayUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> EmaopayLoginResponse Login(ctx).Params(params).Execute()

登录



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	params := *openapiclient.NewEmaopayLoginParams() // EmaopayLoginParams | 登录参数

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.Login(context.Background()).Params(params).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: EmaopayLoginResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | [**EmaopayLoginParams**](EmaopayLoginParams.md) | 登录参数 | 

### Return type

[**EmaopayLoginResponse**](EmaopayLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotifyOrderById

> EmaopayOrder NotifyOrderById(ctx, id).Execute()

通知订单回调



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	id := "id_example" // string | 订单Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.NotifyOrderById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.NotifyOrderById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotifyOrderById`: EmaopayOrder
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.NotifyOrderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 订单Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotifyOrderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmaopayOrder**](EmaopayOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PushMessage

> PushMessage(ctx).XMonitorId(xMonitorId).XToken(xToken).MonitorMessage(monitorMessage).Execute()

推送消息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	xMonitorId := "xMonitorId_example" // string | 监控Id
	xToken := "xToken_example" // string | ApiToken
	monitorMessage := *openapiclient.NewEmaopayMonitorMessageParams() // EmaopayMonitorMessageParams | 监控消息

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.PushMessage(context.Background()).XMonitorId(xMonitorId).XToken(xToken).MonitorMessage(monitorMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PushMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xMonitorId** | **string** | 监控Id | 
 **xToken** | **string** | ApiToken | 
 **monitorMessage** | [**EmaopayMonitorMessageParams**](EmaopayMonitorMessageParams.md) | 监控消息 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshMonitorApiToken

> RefreshMonitorApiToken(ctx, monitorId).Execute()

刷新ApiToken



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	monitorId := "monitorId_example" // string | 监控Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.RefreshMonitorApiToken(context.Background(), monitorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RefreshMonitorApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** | 监控Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshMonitorApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> Register(ctx).Params(params).Execute()

注册



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	params := *openapiclient.NewEmaopayRegisterParams() // EmaopayRegisterParams | 注册参数

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.Register(context.Background()).Params(params).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Register``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | [**EmaopayRegisterParams**](EmaopayRegisterParams.md) | 注册参数 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMerchant

> EmaopayMerchant UpdateMerchant(ctx, id).Merchant(merchant).Execute()

更新商户



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	id := "id_example" // string | 商户Id
	merchant := *openapiclient.NewEmaopayCreateMerchantParams("Name_example") // EmaopayCreateMerchantParams | 商户

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.UpdateMerchant(context.Background(), id).Merchant(merchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMerchant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMerchant`: EmaopayMerchant
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMerchant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 商户Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMerchantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchant** | [**EmaopayCreateMerchantParams**](EmaopayCreateMerchantParams.md) | 商户 | 

### Return type

[**EmaopayMerchant**](EmaopayMerchant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitor

> EmaopayMonitor UpdateMonitor(ctx, id).Monitor(monitor).Execute()

更新监控



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/emaopay/emaopay-sdk-go/emaopay"
)

func main() {
	id := "id_example" // string | 监控Id
	monitor := *openapiclient.NewEmaopayCreateMonitorParams("MerchantId_example", "Name_example", int32(123)) // EmaopayCreateMonitorParams | 监控

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.UpdateMonitor(context.Background(), id).Monitor(monitor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMonitor`: EmaopayMonitor
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 监控Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitor** | [**EmaopayCreateMonitorParams**](EmaopayCreateMonitorParams.md) | 监控 | 

### Return type

[**EmaopayMonitor**](EmaopayMonitor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

