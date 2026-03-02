package controller

import "github.com/gofiber/fiber/v2"

type ParkingController interface {
	Create(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	UpdateCheckout(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
