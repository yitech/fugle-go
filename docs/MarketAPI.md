# \MarketAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCandlesApiV1HistoricalCandlesGet**](MarketAPI.md#GetCandlesApiV1HistoricalCandlesGet) | **Get** /api/v1/historical/candles | Get Candles
[**GetQuoteApiV1IntradayQuoteGet**](MarketAPI.md#GetQuoteApiV1IntradayQuoteGet) | **Get** /api/v1/intraday/quote | Get Quote



## GetCandlesApiV1HistoricalCandlesGet

> KLinesResponse GetCandlesApiV1HistoricalCandlesGet(ctx).Symbol(symbol).FromDate(fromDate).ToDate(toDate).Resolution(resolution).Execute()

Get Candles

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
	symbol := "symbol_example" // string | 
	fromDate := "fromDate_example" // string | 
	toDate := "toDate_example" // string | 
	resolution := "resolution_example" // string |  (optional) (default to "D")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.GetCandlesApiV1HistoricalCandlesGet(context.Background()).Symbol(symbol).FromDate(fromDate).ToDate(toDate).Resolution(resolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.GetCandlesApiV1HistoricalCandlesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCandlesApiV1HistoricalCandlesGet`: KLinesResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.GetCandlesApiV1HistoricalCandlesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCandlesApiV1HistoricalCandlesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 
 **resolution** | **string** |  | [default to &quot;D&quot;]

### Return type

[**KLinesResponse**](KLinesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteApiV1IntradayQuoteGet

> QuoteResponse GetQuoteApiV1IntradayQuoteGet(ctx).Symbol(symbol).Type_(type_).Execute()

Get Quote

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
	symbol := "symbol_example" // string | 
	type_ := "type__example" // string |  (optional) (default to "EQUITY")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.GetQuoteApiV1IntradayQuoteGet(context.Background()).Symbol(symbol).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.GetQuoteApiV1IntradayQuoteGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteApiV1IntradayQuoteGet`: QuoteResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.GetQuoteApiV1IntradayQuoteGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteApiV1IntradayQuoteGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **type_** | **string** |  | [default to &quot;EQUITY&quot;]

### Return type

[**QuoteResponse**](QuoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

