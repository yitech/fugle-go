# \WalletAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalanceEndpointApiV1BalanceGet**](WalletAPI.md#GetBalanceEndpointApiV1BalanceGet) | **Get** /api/v1/balance | Get Balance Endpoint
[**GetInventoriesEndpointApiV1InventoriesGet**](WalletAPI.md#GetInventoriesEndpointApiV1InventoriesGet) | **Get** /api/v1/inventories | Get Inventories Endpoint
[**GetSettlementsEndpointApiV1SettlementsGet**](WalletAPI.md#GetSettlementsEndpointApiV1SettlementsGet) | **Get** /api/v1/settlements | Get Settlements Endpoint



## GetBalanceEndpointApiV1BalanceGet

> BalanceResponse GetBalanceEndpointApiV1BalanceGet(ctx).Execute()

Get Balance Endpoint

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
	resp, r, err := apiClient.WalletAPI.GetBalanceEndpointApiV1BalanceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetBalanceEndpointApiV1BalanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalanceEndpointApiV1BalanceGet`: BalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetBalanceEndpointApiV1BalanceGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceEndpointApiV1BalanceGetRequest struct via the builder pattern


### Return type

[**BalanceResponse**](BalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInventoriesEndpointApiV1InventoriesGet

> []InventoryResponse GetInventoriesEndpointApiV1InventoriesGet(ctx).Execute()

Get Inventories Endpoint

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
	resp, r, err := apiClient.WalletAPI.GetInventoriesEndpointApiV1InventoriesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetInventoriesEndpointApiV1InventoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInventoriesEndpointApiV1InventoriesGet`: []InventoryResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetInventoriesEndpointApiV1InventoriesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoriesEndpointApiV1InventoriesGetRequest struct via the builder pattern


### Return type

[**[]InventoryResponse**](InventoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettlementsEndpointApiV1SettlementsGet

> []SettlementResponse GetSettlementsEndpointApiV1SettlementsGet(ctx).Execute()

Get Settlements Endpoint

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
	resp, r, err := apiClient.WalletAPI.GetSettlementsEndpointApiV1SettlementsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetSettlementsEndpointApiV1SettlementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettlementsEndpointApiV1SettlementsGet`: []SettlementResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetSettlementsEndpointApiV1SettlementsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettlementsEndpointApiV1SettlementsGetRequest struct via the builder pattern


### Return type

[**[]SettlementResponse**](SettlementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

