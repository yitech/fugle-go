# \OrderAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrderEndpointApiV1OrderPost**](OrderAPI.md#CreateOrderEndpointApiV1OrderPost) | **Post** /api/v1/order | Create Order Endpoint
[**DeleteOrderEndpointApiV1OrderOrdNoDelete**](OrderAPI.md#DeleteOrderEndpointApiV1OrderOrdNoDelete) | **Delete** /api/v1/order/{ord_no} | Delete Order Endpoint
[**GetMarketStatusEndpointApiV1MarketStatusGet**](OrderAPI.md#GetMarketStatusEndpointApiV1MarketStatusGet) | **Get** /api/v1/market_status | Get Market Status Endpoint
[**GetOrdersEndpointApiV1OrdersGet**](OrderAPI.md#GetOrdersEndpointApiV1OrdersGet) | **Get** /api/v1/orders | Get Orders Endpoint



## CreateOrderEndpointApiV1OrderPost

> OrderResponse CreateOrderEndpointApiV1OrderPost(ctx).CreateOrder(createOrder).Execute()

Create Order Endpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createOrder := *openapiclient.NewCreateOrder(openapiclient.Action("B"), openapiclient.APCode("1"), openapiclient.PriceFlag("0"), openapiclient.BSFlag("F"), openapiclient.Trade("0"), "StockNo_example", int32(123), float32(123)) // CreateOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CreateOrderEndpointApiV1OrderPost(context.Background()).CreateOrder(createOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CreateOrderEndpointApiV1OrderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrderEndpointApiV1OrderPost`: OrderResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CreateOrderEndpointApiV1OrderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderEndpointApiV1OrderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrder** | [**CreateOrder**](CreateOrder.md) |  | 

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrderEndpointApiV1OrderOrdNoDelete

> CancelResponse DeleteOrderEndpointApiV1OrderOrdNoDelete(ctx, ordNo).Execute()

Delete Order Endpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ordNo := "ordNo_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.DeleteOrderEndpointApiV1OrderOrdNoDelete(context.Background(), ordNo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.DeleteOrderEndpointApiV1OrderOrdNoDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrderEndpointApiV1OrderOrdNoDelete`: CancelResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.DeleteOrderEndpointApiV1OrderOrdNoDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ordNo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderEndpointApiV1OrderOrdNoDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CancelResponse**](CancelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketStatusEndpointApiV1MarketStatusGet

> MarketStatusResponse GetMarketStatusEndpointApiV1MarketStatusGet(ctx).Execute()

Get Market Status Endpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.GetMarketStatusEndpointApiV1MarketStatusGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetMarketStatusEndpointApiV1MarketStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketStatusEndpointApiV1MarketStatusGet`: MarketStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetMarketStatusEndpointApiV1MarketStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketStatusEndpointApiV1MarketStatusGetRequest struct via the builder pattern


### Return type

[**MarketStatusResponse**](MarketStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersEndpointApiV1OrdersGet

> []OrderResultResponse GetOrdersEndpointApiV1OrdersGet(ctx).Execute()

Get Orders Endpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.GetOrdersEndpointApiV1OrdersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetOrdersEndpointApiV1OrdersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrdersEndpointApiV1OrdersGet`: []OrderResultResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetOrdersEndpointApiV1OrdersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersEndpointApiV1OrdersGetRequest struct via the builder pattern


### Return type

[**[]OrderResultResponse**](OrderResultResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

