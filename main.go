package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pisondev/parking-system-api/app"
	"github.com/pisondev/parking-system-api/controller"
	"github.com/pisondev/parking-system-api/exception"
	"github.com/pisondev/parking-system-api/repository"
	"github.com/pisondev/parking-system-api/service"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	if err := godotenv.Load(); err != nil {
		log.Warn("no .env file found")
	}

	db := app.NewDB(log)
	validate := validator.New()

	parkingRepo := repository.NewParkingRepository()
	parkingService := service.NewParkingService(parkingRepo, db)
	parkingController := controller.NewParkingController(parkingService, validate)

	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	app.SetupRouter(fiberApp, parkingController)

	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Mall Parking System API is running",
		})
	})

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}

	listenAddr := fmt.Sprintf(":%s", appPort)
	log.Infof("server starting on port %s...", appPort)

	if err := fiberApp.Listen(listenAddr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
