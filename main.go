package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jkeresman01/SalesAPI/Model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func main() {
	godotenv.Load()

	dbHost := os.Getenv("MYSQL_HOST ")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbUser := os.Getenv("MYSQL_USER")
	dbName := os.Getenv("MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	var db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("MYSQL databse connection has failed")
	}

	DB = db
	fmt.Printf("Successful connection to %s", dbName)

	err = AutoMigrate(DB)

	if err != nil {
		log.Error("DB auto migration has failed!")
	}
}

func AutoMigrate(connection *gorm.DB) error {
	err := connection.Debug().AutoMigrate(
		&Model.Cashier{},
		&Model.Category{},
		&Model.Discount{},
		&Model.Order{},
		&Model.Payment{},
		&Model.PaymentType{},
		&Model.Product{},
	)

	return err
}
