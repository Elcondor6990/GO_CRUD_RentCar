package controller

import (
	"Go_Rent_Car/src/dto"
	"Go_Rent_Car/src/service"

	"github.com/gofiber/fiber/v3"
)

type ReservationCarController struct {
	service service.ReservationCarService
}

func NewReservationCarController(service *service.ReservationCarService) *ReservationCarController{
	return &ReservationCarController{service: *service}
}

func (rc *ReservationCarController) CreateReservation(c fiber.Ctx)error{
	var input dto.ReservationCarDto

	if err := c.Bind().JSON(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error: ": err.Error()})
	}

	reservation, err := rc.service.Create(c.Context(), input)
	if err != nil{
		return c.SendStatus(fiber.StatusInternalServerError)
	}	

	return c.Status(fiber.StatusCreated).JSON(reservation)
}

func (rc *ReservationCarController) GetByID(c fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error:": "ID required"})
	}

	reservation, err := rc.service.GetByID(c.Context(), id)
	if  err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error: ": "Reservation not found"})
	}

	return c.JSON(reservation)
}

func (rc *ReservationCarController) DeleteReservation(c fiber.Ctx) error {
	id := c.Params("id")
	
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID required"})
	}

	err := rc.service.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Reservation not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}



