package main

import (
	_ "embed"
	log "github.com/sirupsen/logrus"
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
	"os"
)

//go:embed fees.csv
var fees string

func main() {

	conf.LoadConfig("DEV")
	log.SetOutput(os.Stdout)
	db.SetupInstance()

	if err := db.MigrateUp(); err != nil {
		log.Fatal(err)
	}

	if conf.Get().UpdateFees {
		models.LoadFees(fees)
	}

}
