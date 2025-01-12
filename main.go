package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/jkeresman01/SalesAPI/Config"
	routes "github.com/jkeresman01/SalesAPI/Route"
)

func main() {
	db.Connect()

	app := fiber.New()
	app.Use(app)
	routes.Setup(app)

	app.Listen(":3000")
}
