package main

import (
	log "github.com/sirupsen/logrus"
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	"os"
)

func main() {

	conf.LoadConfig("DEV")
	log.SetOutput(os.Stdout)
	db.SetupInstance()

	if err := db.MigrateUp(); err != nil {
		log.Fatal(err)
	}

}
