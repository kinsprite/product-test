package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Product Model
type Product struct {
	gorm.Model
	Code  string
	Name  string
	Price float64
}

// Book Model
type Book struct {
	gorm.Model
	ProductID uint
	Product   Product
	AuthorID  uint
	Author    Author
	ISBN      string
	Comment   string
}

// Author Model
type Author struct {
	gorm.Model
	FirstName string
	LastName  string
	PenName   string
	Birthday  string
}

// UserBuy Model
type UserBuy struct {
	gorm.Model
	UserID    uint
	ProductID uint
}

// UserBook Query result
type UserBook struct {
	ID      string
	Code    string
	Name    string
	Price   float64
	ISBN    string
	Comment string
}

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

	// Migrate the schema
	db.AutoMigrate(&Product{}, &Book{}, &Author{}, &UserBuy{})
}

func getUserBooks() (userBooks *[]UserBook, err error) {
	userID := 10001
	userBooks = &[]UserBook{}

	err = db.Raw(`SELECT (C.id, C.code, C.name, c.price, B.isbn, B.comment)
			FROM (
				SELECT (product_id) FROM user_buys WHERE user_id = ?
				) AS A
			INNER JOIN books B ON A.product_id = B.product_id
			LEFT JOIN products C ON B.product_id = C.id
		`, userID).Scan(userBooks).Error

	return
}
