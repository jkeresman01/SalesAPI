package Route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/SalesAPI/Controller"
)

func Setup(app *fiber.App) {
	AuthRoutes(app)
	CashierRoutes(app)
	ProductRoutes(app)
}

func ProductRoutes(app *fiber.App) {
	app.Get("/categories", Controller.GetCategories)
	app.Get("/categories/:categoryId", Controller.GetCategoryWithId)
	app.Post("/categories", Controller.CreateCategory)
	app.Delete("/categories:categoryId", Controller.DeleteCategoryWithId)
	app.Put("/categories:categoryId", Controller.UpdateCategoryWithId)
}

func CashierRoutes(app *fiber.App) {
	app.Post("/cashiers", Controller.CreateCashier)
	app.Get("/cashiers", Controller.GetCashiers)
	app.Get("/cashiers/:cashierId", Controller.GetCashierWithId)
	app.Delete("/cashiers/:cashierId", Controller.DeleteCashierWithId)
	app.Put("/cashiers/:cashierId", Controller.UpdateCashierWithId)
}

func AuthRoutes(app *fiber.App) {
	app.Post("/cashiers/:cashierId/login", Controller.Login)
	app.Get("/cashiers/:cashierId/logout", Controller.Logout)
	app.Get("/cashiers/:cashierId/passcode", Controller.Passcode)
}
