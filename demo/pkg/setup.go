package pkg

import (
	"context"
	"open-outcry/demo/pkg/api"
)

type App struct {
	Ctx       context.Context
	ApiClient *api.APIClient
}

var app *App

func SetupAll() error {
	ctx := context.Background()
	app = &App{Ctx: ctx}

	configuration := api.NewConfiguration()
	app.ApiClient = api.NewAPIClient(configuration)

	return nil
}

func GetApp() *App {
	return app
}
