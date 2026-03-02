package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pisondev/parking-system-api/controller"
)

func SetupRouter(app *fiber.App, parkingController controller.ParkingController) {
	api := app.Group("/api")

	api.Post("/parking", parkingController.Create)
	api.Get("/parking", parkingController.FindAll)
	api.Get("/parking/:id", parkingController.FindById)
	api.Patch("/parking/:id", parkingController.UpdateCheckout)
	api.Delete("/parking/:id", parkingController.Delete)
}
