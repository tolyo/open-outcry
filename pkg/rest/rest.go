package rest

import (
	"net/http"
	"open-outcry/pkg/conf"
	"open-outcry/pkg/rest/api"
)

func NewServer() http.Server {
	router := api.NewRouter(
		api.NewCurrenciesAPIController(api.NewCurrenciesAPIService()),
	)
	return http.Server{
		Addr:    ":" + conf.Get().RestPort,
		Handler: router,
	}
}
