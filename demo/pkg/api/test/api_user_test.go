/*
OPEN OUTCRY API

Testing UserAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	"testing"

	openapiclient "open-outcry/demo/pkg/api"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_UserAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserAPIService CreateTrade", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tradingAccountId string

		resp, httpRes, err := apiClient.UserAPI.CreateTrade(context.Background(), tradingAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService DeleteTradeOrderById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tradingAccountId string
		var tradeOrderId string

		httpRes, err := apiClient.UserAPI.DeleteTradeOrderById(context.Background(), tradingAccountId, tradeOrderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetBookOrders", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tradingAccountId string

		resp, httpRes, err := apiClient.UserAPI.GetBookOrders(context.Background(), tradingAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetPaymentAccounts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appEntityId string

		resp, httpRes, err := apiClient.UserAPI.GetPaymentAccounts(context.Background(), appEntityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetTradeById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tradingAccountId string
		var tradeId string

		resp, httpRes, err := apiClient.UserAPI.GetTradeById(context.Background(), tradingAccountId, tradeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetTradeOrderById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tradingAccountId string
		var tradeOrderId string

		resp, httpRes, err := apiClient.UserAPI.GetTradeOrderById(context.Background(), tradingAccountId, tradeOrderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetTradeOrders", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tradingAccountId string

		resp, httpRes, err := apiClient.UserAPI.GetTradeOrders(context.Background(), tradingAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetTrades", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tradingAccountId string

		resp, httpRes, err := apiClient.UserAPI.GetTrades(context.Background(), tradingAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetTradingAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tradingAccountId string

		resp, httpRes, err := apiClient.UserAPI.GetTradingAccount(context.Background(), tradingAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
