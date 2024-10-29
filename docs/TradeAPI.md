# \TradeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransactionEndpointApiV2TransactionsGet**](TradeAPI.md#GetTransactionEndpointApiV2TransactionsGet) | **Get** /api/v2/transactions | Get Transaction Endpoint



## GetTransactionEndpointApiV2TransactionsGet

> []AppSchemaV2TradeTransactionResponse GetTransactionEndpointApiV2TransactionsGet(ctx).QueryRange(queryRange).Execute()

Get Transaction Endpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yitech/fugle-go"
)

func main() {
	queryRange := "queryRange_example" // string |  (optional) (default to "0d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTransactionEndpointApiV2TransactionsGet(context.Background()).QueryRange(queryRange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTransactionEndpointApiV2TransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionEndpointApiV2TransactionsGet`: []AppSchemaV2TradeTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTransactionEndpointApiV2TransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionEndpointApiV2TransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryRange** | **string** |  | [default to &quot;0d&quot;]

### Return type

[**[]AppSchemaV2TradeTransactionResponse**](AppSchemaV2TradeTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

