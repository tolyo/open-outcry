# Go API client for api

# Introduction
This API is documented in **OpenAPI 3.0 format**.

This API the following operations:
* Retrieve a list of available instruments
* Retrieve a list of executed trades

# Basics
* API calls have to be secured with HTTPS.
* All data has to be submitted UTF-8 encoded.
* The reply is sent JSON encoded.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import api "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `api.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), api.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `api.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), api.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `api.ContextOperationServerIndices` and `api.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), api.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), api.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:4000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminAPI* | [**CreateAdminPayment**](docs/AdminAPI.md#createadminpayment) | **Post** /apps/payments | Create admin payment
*AdminAPI* | [**GetAdminPaymentById**](docs/AdminAPI.md#getadminpaymentbyid) | **Post** /apps/payments/{payment_id} | Get payment
*AdminAPI* | [**GetAppEntities**](docs/AdminAPI.md#getappentities) | **Get** /apps | Get application entities
*AdminAPI* | [**GetAppEntity**](docs/AdminAPI.md#getappentity) | **Get** /apps/{app_entity_id} | Get application entity
*PublicAPI* | [**GetCurrencies**](docs/PublicAPI.md#getcurrencies) | **Get** /currencies | Currencies list
*PublicAPI* | [**GetFxInstruments**](docs/PublicAPI.md#getfxinstruments) | **Get** /fxinstruments | Fx instrument list
*PublicAPI* | [**GetInstruments**](docs/PublicAPI.md#getinstruments) | **Get** /instruments | Instrument list
*PublicAPI* | [**GetOrderBook**](docs/PublicAPI.md#getorderbook) | **Get** /order-books/{instrument_name} | Get order book
*UserAPI* | [**CreateTrade**](docs/UserAPI.md#createtrade) | **Post** /trade-orders/{trading_account_id} | Create trade order
*UserAPI* | [**DeleteTradeOrderById**](docs/UserAPI.md#deletetradeorderbyid) | **Delete** /trade-orders/{trading_account_id}/id/{trade_order_id} | Cancel trade order
*UserAPI* | [**GetBookOrders**](docs/UserAPI.md#getbookorders) | **Get** /book-orders/{trading_account_id} | Get book orders
*UserAPI* | [**GetPaymentAccounts**](docs/UserAPI.md#getpaymentaccounts) | **Get** /payment-accounts/{app_entity_id} | Get payment accounts
*UserAPI* | [**GetTradeById**](docs/UserAPI.md#gettradebyid) | **Get** /trades/{trading_account_id}/id/{trade_id} | Get trade
*UserAPI* | [**GetTradeOrderById**](docs/UserAPI.md#gettradeorderbyid) | **Get** /trade-orders/{trading_account_id}/id/{trade_order_id} | Get trade order
*UserAPI* | [**GetTradeOrders**](docs/UserAPI.md#gettradeorders) | **Get** /trade-orders/{trading_account_id} | Get trade orders
*UserAPI* | [**GetTrades**](docs/UserAPI.md#gettrades) | **Get** /trades/{trading_account_id} | Trades list
*UserAPI* | [**GetTradingAccount**](docs/UserAPI.md#gettradingaccount) | **Get** /trading-accounts/{trading_account_id} | Get trading account


## Documentation For Models

 - [AppEntity](docs/AppEntity.md)
 - [CreateTradeRequest](docs/CreateTradeRequest.md)
 - [Currency](docs/Currency.md)
 - [CurrencyList](docs/CurrencyList.md)
 - [FxInstrument](docs/FxInstrument.md)
 - [Instrument](docs/Instrument.md)
 - [OrderBook](docs/OrderBook.md)
 - [Payment](docs/Payment.md)
 - [PaymentAccount](docs/PaymentAccount.md)
 - [PaymentAccountList](docs/PaymentAccountList.md)
 - [PaymentType](docs/PaymentType.md)
 - [PriceVolume](docs/PriceVolume.md)
 - [Trade](docs/Trade.md)
 - [TradeOrder](docs/TradeOrder.md)
 - [TradeOrderSide](docs/TradeOrderSide.md)
 - [TradeOrderStatus](docs/TradeOrderStatus.md)
 - [TradeOrderTimeInForce](docs/TradeOrderTimeInForce.md)
 - [TradeOrderType](docs/TradeOrderType.md)
 - [TradingAccount](docs/TradingAccount.md)
 - [TradingAccountInstrument](docs/TradingAccountInstrument.md)


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


