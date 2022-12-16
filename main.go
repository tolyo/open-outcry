package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
)

func main() {
	fmt.Print("Test\n")

	conf.LoadConfig("DEV")

	db.SetupInstance()

	if err := db.MigrateUp(); err != nil {
		log.Fatal(err)
	}

}
