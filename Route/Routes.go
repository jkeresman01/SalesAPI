package Route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/SalesAPI/Controller"
)

func Setup(app *fiber.App) {
	app.Post("/cashiers/:cashierId/login", Controller.Login)
	app.Get("/cashiers/:cashierId", Controller.Logout)
	app.Get("/cashiers/:cashierId/passcode", Controller.Passcode)

	app.Post("/cashiers", Controller.CreateCashier)
	app.Get("/cashiers", Controller.GetCashiers)
	app.Get("/cashiers/:cashierId", Controller.GetCashierWithId)
	app.Delete("/cashiers/:cashierId", Controller.DeleteCashierWithId)
	app.Put("/cashiers/:cashierId", Controller.UpdateCashierWithId)

}
