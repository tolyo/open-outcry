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

	server := rest.NewServer()
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
