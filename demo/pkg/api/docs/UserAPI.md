# \UserAPI

All URIs are relative to *http://localhost:4000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrade**](UserAPI.md#CreateTrade) | **Post** /trade-orders/{trading_account_id} | Create trade order
[**DeleteTradeOrderById**](UserAPI.md#DeleteTradeOrderById) | **Delete** /trade-orders/{trading_account_id}/id/{trade_order_id} | Cancel trade order
[**GetBookOrders**](UserAPI.md#GetBookOrders) | **Get** /book-orders/{trading_account_id} | Get book orders
[**GetPaymentAccounts**](UserAPI.md#GetPaymentAccounts) | **Get** /payment-accounts/{app_entity_id} | Get payment accounts
[**GetTradeById**](UserAPI.md#GetTradeById) | **Get** /trades/{trading_account_id}/id/{trade_id} | Get trade
[**GetTradeOrderById**](UserAPI.md#GetTradeOrderById) | **Get** /trade-orders/{trading_account_id}/id/{trade_order_id} | Get trade order
[**GetTradeOrders**](UserAPI.md#GetTradeOrders) | **Get** /trade-orders/{trading_account_id} | Get trade orders
[**GetTrades**](UserAPI.md#GetTrades) | **Get** /trades/{trading_account_id} | Trades list
[**GetTradingAccount**](UserAPI.md#GetTradingAccount) | **Get** /trading-accounts/{trading_account_id} | Get trading account



## CreateTrade

> TradeOrder CreateTrade(ctx, tradingAccountId).CreateTradeRequest(createTradeRequest).Execute()

Create trade order



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
    tradingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createTradeRequest := *openapiclient.NewCreateTradeRequest("BTC-EUR", openapiclient.TradeOrderSide("SELL"), openapiclient.TradeOrderType("LIMIT"), openapiclient.TradeOrderTimeInForce("GTC"), "100.50") // CreateTradeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.CreateTrade(context.Background(), tradingAccountId).CreateTradeRequest(createTradeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateTrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrade`: TradeOrder
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateTrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradingAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTradeRequest** | [**CreateTradeRequest**](CreateTradeRequest.md) |  | 

### Return type

[**TradeOrder**](TradeOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTradeOrderById

> DeleteTradeOrderById(ctx, tradingAccountId, tradeOrderId).Execute()

Cancel trade order



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
    tradingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    tradeOrderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.DeleteTradeOrderById(context.Background(), tradingAccountId, tradeOrderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteTradeOrderById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradingAccountId** | **string** |  | 
**tradeOrderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTradeOrderByIdRequest struct via the builder pattern


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


## GetBookOrders

> []TradeOrder GetBookOrders(ctx, tradingAccountId).Execute()

Get book orders



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
    tradingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetBookOrders(context.Background(), tradingAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetBookOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookOrders`: []TradeOrder
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetBookOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradingAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TradeOrder**](TradeOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentAccounts

> PaymentAccountList GetPaymentAccounts(ctx, appEntityId).Execute()

Get payment accounts



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
    appEntityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetPaymentAccounts(context.Background(), appEntityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetPaymentAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentAccounts`: PaymentAccountList
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetPaymentAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appEntityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentAccountList**](PaymentAccountList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeById

> Trade GetTradeById(ctx, tradingAccountId, tradeId).Execute()

Get trade



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
    tradingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    tradeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetTradeById(context.Background(), tradingAccountId, tradeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTradeById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradeById`: Trade
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTradeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradingAccountId** | **string** |  | 
**tradeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Trade**](Trade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOrderById

> TradeOrder GetTradeOrderById(ctx, tradingAccountId, tradeOrderId).Execute()

Get trade order



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
    tradingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    tradeOrderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetTradeOrderById(context.Background(), tradingAccountId, tradeOrderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTradeOrderById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradeOrderById`: TradeOrder
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTradeOrderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradingAccountId** | **string** |  | 
**tradeOrderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOrderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TradeOrder**](TradeOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOrders

> []TradeOrder GetTradeOrders(ctx, tradingAccountId).Execute()

Get trade orders



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
    tradingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetTradeOrders(context.Background(), tradingAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTradeOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradeOrders`: []TradeOrder
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTradeOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradingAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TradeOrder**](TradeOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrades

> []Trade GetTrades(ctx, tradingAccountId).Execute()

Trades list



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
    tradingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetTrades(context.Background(), tradingAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrades`: []Trade
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradingAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Trade**](Trade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingAccount

> TradingAccount GetTradingAccount(ctx, tradingAccountId).Execute()

Get trading account



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
    tradingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetTradingAccount(context.Background(), tradingAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTradingAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradingAccount`: TradingAccount
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTradingAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tradingAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TradingAccount**](TradingAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

