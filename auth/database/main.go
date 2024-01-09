package database

import (
	"auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) (){
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil{
		log.Fatal(dbError)
		panic("cannot connect to database")
	}
	log.Println("database connected successfully")
}

func Migrate() {
	err := Instance.AutoMigrate(&models.User{})
	if err != nil {
		return 
	}
	log.Println("database migration completed")
}

