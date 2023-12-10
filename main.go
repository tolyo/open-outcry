package main

import (
	_ "embed"
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
	"open-outcry/pkg/rest"
	"os"

	log "github.com/sirupsen/logrus"
)

//go:embed fees.csv
var fees string

func main() {

	envVarValue := os.Getenv("ENV")
	if envVarValue == "" {
		envVarValue = "DEV"
	}

	conf.LoadConfig(envVarValue)
	log.SetOutput(os.Stdout)
	db.SetupInstance()

	if conf.Get().UpdateFees {
		models.LoadFees(fees)
	}

	// Sample seed for debugging
	//_, tradingAccount1 := services.Acc("test")
	//services.ProcessTradeOrder(tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 1, "GTC")

	server := rest.NewServer()
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
