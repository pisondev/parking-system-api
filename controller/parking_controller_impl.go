package controller

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pisondev/parking-system-api/model/web"
	"github.com/pisondev/parking-system-api/service"
)

type ParkingControllerImpl struct {
	ParkingService service.ParkingService
	Validator      *validator.Validate
}

func NewParkingController(parkingService service.ParkingService, validator *validator.Validate) ParkingController {
	return &ParkingControllerImpl{
		ParkingService: parkingService,
		Validator:      validator,
	}
}

func (controller *ParkingControllerImpl) Create(c *fiber.Ctx) error {
	var request web.ParkingCreateRequest

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := controller.Validator.Struct(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := controller.ParkingService.Create(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":   fiber.StatusCreated,
		"status": "CREATED",
		"data":   response,
	})
}

func (controller *ParkingControllerImpl) FindAll(c *fiber.Ctx) error {
	responses, err := controller.ParkingService.FindAll(c.Context())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   responses,
	})
}

func (controller *ParkingControllerImpl) FindById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid transaction id")
	}

	response, err := controller.ParkingService.FindById(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   response,
	})
}

func (controller *ParkingControllerImpl) UpdateCheckout(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid transaction id")
	}

	response, err := controller.ParkingService.UpdateCheckout(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   response,
	})
}

func (controller *ParkingControllerImpl) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid transaction id")
	}

	err = controller.ParkingService.Delete(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   nil,
	})
}
