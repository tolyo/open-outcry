package main

import (
	_ "embed"
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
	"open-outcry/pkg/services"
	"os"
	"sync"

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

	// sample code to generate 100 trades
	_, tradingAccount1 := services.Acc("test")
	_, tradingAccount2 := services.Acc("test2")

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {

		go func() {
			mu.Lock()
			log.Info(services.ProcessTradeOrder(tradingAccount1, "BTC_EUR", "LIMIT", "SELL", 1, 10, "GTC"))
			log.Info(services.ProcessTradeOrder(tradingAccount2, "BTC_EUR", "MARKET", "BUY", 0, 10, "GTC"))
			mu.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

}
