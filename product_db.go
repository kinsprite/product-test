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
	Books     []Book
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

	if isEmptyBooks() {
		log.Println("INFO    Gorm DB books is Empty")
		createUserBooks()
	}
}

func getUserBooks() (userBooks *[]UserBook, err error) {
	userID := uint(10001)
	var userBooksGot []UserBook

	err = db.Raw(`SELECT C.id, C.code, C.name, C.price, B.isbn, B.comment
		FROM (SELECT product_id FROM user_buys WHERE user_id = ?) AS A
		INNER JOIN books B ON A.product_id = B.product_id
		LEFT JOIN products C ON B.product_id = C.id`, userID).Scan(&userBooksGot).Error

	if err == nil {
		userBooks = &userBooksGot
	}

	return
}

func isEmptyBooks() bool {
	var count int
	err := db.Model(&Book{}).Count(&count).Error
	return err == nil && count == 0
}

func createUserBooks() {
	userID := uint(10001)
	authors := []Author{
		{
			FirstName: "Ken",
			LastName:  "Bench",
			PenName:   "KB",
			Birthday:  "1953-2-11",
			Books: []Book{
				{
					Product: Product{
						Code:  "P-10001",
						Name:  "Book Name 1",
						Price: 46.33,
					},
					ISBN:    "543-233-33",
					Comment: "c",
				},
				{
					Product: Product{
						Code:  "P-10002",
						Name:  "Book Name 2",
						Price: 26.67,
					},
					ISBN:    "543-233-34",
					Comment: "abc",
				},
			},
		},
		{
			FirstName: "Martin",
			LastName:  "Fowler",
			PenName:   "MF",
			Birthday:  "1973-12-13",
			Books: []Book{
				{
					Product: Product{
						Code:  "P-10021",
						Name:  "Book Name 21",
						Price: 76.33,
					},
					ISBN:    "543-233-43",
					Comment: "a",
				},
				{
					Product: Product{
						Code:  "P-10022",
						Name:  "Book Name 22",
						Price: 76.67,
					},
					ISBN:    "543-233-44",
					Comment: "123",
				},
			},
		},
	}

	for i := range authors {
		db.Create(&authors[i])
	}

	// UserBuys
	var products []Product
	db.Find(&products)

	userBuys := make([]UserBuy, len(products))

	for i, product := range products {
		userBuys[i] = UserBuy{
			UserID:    userID,
			ProductID: product.ID,
		}
	}

	for i := range userBuys {
		db.Create(&userBuys[i])
	}
}
