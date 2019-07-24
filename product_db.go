package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	driverName := os.Getenv("SQL_DRIVER_NAME")
	dataSourceName := os.Getenv("SQL_DATA_SOURCE_NAME")

	var err error
	db, err = gorm.Open(driverName, dataSourceName)

	if err != nil {
		log.Println("ERROR    Gorm open DB failed")
	} else {
		log.Println("INFO    Gorm open DB OK")
	}
}
