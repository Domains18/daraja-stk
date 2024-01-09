package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func connectDb() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database\n")
		os.Exit(2)
	}
	log.Println("connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations") //to add migrations

	Database = DbInstance{Db: db}
}
