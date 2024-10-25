# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 1.0.0
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MarketAPI* | [**GetCandlesApiV1HistoricalCandlesGet**](docs/MarketAPI.md#getcandlesapiv1historicalcandlesget) | **Get** /api/v1/historical/candles | Get Candles
*MarketAPI* | [**GetQuoteApiV1IntradayQuoteGet**](docs/MarketAPI.md#getquoteapiv1intradayquoteget) | **Get** /api/v1/intraday/quote | Get Quote
*OrderAPI* | [**CreateOrderEndpointApiV1OrderPost**](docs/OrderAPI.md#createorderendpointapiv1orderpost) | **Post** /api/v1/order | Create Order Endpoint
*OrderAPI* | [**DeleteOrderEndpointApiV1OrderOrdNoDelete**](docs/OrderAPI.md#deleteorderendpointapiv1orderordnodelete) | **Delete** /api/v1/order/{ord_no} | Delete Order Endpoint
*OrderAPI* | [**GetMarketStatusEndpointApiV1MarketStatusGet**](docs/OrderAPI.md#getmarketstatusendpointapiv1marketstatusget) | **Get** /api/v1/market_status | Get Market Status Endpoint
*OrderAPI* | [**GetOrdersEndpointApiV1OrdersGet**](docs/OrderAPI.md#getordersendpointapiv1ordersget) | **Get** /api/v1/orders | Get Orders Endpoint
*SystemAPI* | [**PingApiV1PingGet**](docs/SystemAPI.md#pingapiv1pingget) | **Get** /api/v1/ping | Ping
*WalletAPI* | [**GetBalanceEndpointApiV1BalanceGet**](docs/WalletAPI.md#getbalanceendpointapiv1balanceget) | **Get** /api/v1/balance | Get Balance Endpoint
*WalletAPI* | [**GetInventoriesEndpointApiV1InventoriesGet**](docs/WalletAPI.md#getinventoriesendpointapiv1inventoriesget) | **Get** /api/v1/inventories | Get Inventories Endpoint
*WalletAPI* | [**GetSettlementsEndpointApiV1SettlementsGet**](docs/WalletAPI.md#getsettlementsendpointapiv1settlementsget) | **Get** /api/v1/settlements | Get Settlements Endpoint


## Documentation For Models

 - [APCode](docs/APCode.md)
 - [Action](docs/Action.md)
 - [BSFlag](docs/BSFlag.md)
 - [BalanceResponse](docs/BalanceResponse.md)
 - [BidAsk](docs/BidAsk.md)
 - [CancelResponse](docs/CancelResponse.md)
 - [CreateOrder](docs/CreateOrder.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [InventoryDetail](docs/InventoryDetail.md)
 - [InventoryResponse](docs/InventoryResponse.md)
 - [KLine](docs/KLine.md)
 - [KLinesResponse](docs/KLinesResponse.md)
 - [LastTrade](docs/LastTrade.md)
 - [LastTrial](docs/LastTrial.md)
 - [MarketStatusResponse](docs/MarketStatusResponse.md)
 - [OrderResponse](docs/OrderResponse.md)
 - [OrderResultResponse](docs/OrderResultResponse.md)
 - [PriceFlag](docs/PriceFlag.md)
 - [QuoteResponse](docs/QuoteResponse.md)
 - [SettlementResponse](docs/SettlementResponse.md)
 - [Total](docs/Total.md)
 - [Trade](docs/Trade.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationErrorLocInner](docs/ValidationErrorLocInner.md)


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


