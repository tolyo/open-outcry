package rest

import (
	"net/http"
	"open-outcry/pkg/conf"
	"open-outcry/pkg/rest/api"
)

func NewServer() http.Server {
	router := api.NewRouter(
		api.NewPublicAPIController(api.NewPublicAPIService()),
	)

	if conf.IsDevEnvironment() {
		router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./pkg/static/"))))
		router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./pkg/static/index.html")
		})
	}

	return http.Server{
		Addr:    ":" + conf.Get().RestPort,
		Handler: router,
	}
}
