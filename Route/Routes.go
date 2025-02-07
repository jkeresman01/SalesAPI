package Route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/SalesAPI/Controller"
)

func Setup(app *fiber.App) {
	SetupAuthRoutes(app)
	SetupCashierRoutes(app)
	SetupProductRoutes(app)
	SetupOrderRoutes(app)
}

func SetupOrderRoutes(app *fiber.App) {
	app.Get("/orders", Controller.GetOrders)
	app.Get("/orders/:orderId", Controller.GetOrderWithId)
	app.Post("/orders/subtotal", Controller.SubTotalOrders)
	app.Get("/orders/:orderId/download", Controller.DownloadOrder)
	app.Get("/orders/:orderId/check-download", Controller.CheckOrder)
}

func SetupProductRoutes(app *fiber.App) {
	app.Get("/categories", Controller.GetCategories)
	app.Get("/categories/:categoryId", Controller.GetCategoryWithId)
	app.Post("/categories", Controller.CreateCategory)
	app.Delete("/categories:categoryId", Controller.DeleteCategoryWithId)
	app.Put("/categories:categoryId", Controller.UpdateCategoryWithId)
}

func SetupCashierRoutes(app *fiber.App) {
	app.Post("/cashiers", Controller.CreateCashier)
	app.Get("/cashiers", Controller.GetCashiers)
	app.Get("/cashiers/:cashierId", Controller.GetCashierWithId)
	app.Delete("/cashiers/:cashierId", Controller.DeleteCashierWithId)
	app.Put("/cashiers/:cashierId", Controller.UpdateCashierWithId)
}

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/cashiers/:cashierId/login", Controller.Login)
	app.Get("/cashiers/:cashierId/logout", Controller.Logout)
	app.Get("/cashiers/:cashierId/passcode", Controller.Passcode)
}
