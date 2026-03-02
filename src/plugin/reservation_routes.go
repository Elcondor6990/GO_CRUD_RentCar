package plugin

import (
	"Go_Rent_Car/src/controller"
	"Go_Rent_Car/src/service"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App, reserveService *service.ReservationCarService) {

	app.Get("/healthz", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	api := app.Group("/api/internal/reservations")

	reserveController := controller.NewReservationCarController(reserveService)
	api.Post("/v1", reserveController.CreateReservation)
	api.Get("/:id/v1", reserveController.GetByID)
	api.Delete("/:id/v1", reserveController.DeleteReservation)
}