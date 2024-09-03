package main

import (
	"annotator-backend/internal/app"
	"annotator-backend/pkg/db/mysql"
	"database/sql"
	"github.com/labstack/gommon/log"

	"annotator-backend/config"
)

const (
	ConfigFile = "config"
)

func main() {
	log.Info("Starting annotator backend...")

	configFile, err := config.LoadConfig(ConfigFile)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	parsedConfigFile, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	mySqlDB, err := mysql.NewMySqlDB(parsedConfigFile)
	if err != nil {
		log.Fatalf("MySql init: %s", err)
	} else {
		log.Printf("MySql connected, Status: %#v\n", mySqlDB.Stats())
	}
	defer func(mySqlDB *sql.DB) {
		err := mySqlDB.Close()
		if err != nil {

		}
	}(mySqlDB)

	s := app.NewAnnotatorApp(parsedConfigFile, mySqlDB)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
