package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	err := godotenv.Load()
	if err != nil {
		log.Warn(".env file is not found")
	} else {
		log.Info("success load .env config")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Mall Parking System API!",
			"status":  "Server is running",
		})
	})

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}

	listenAddr := fmt.Sprintf(":%s", appPort)
	log.Infof("Server starting on port %s...", appPort)

	if err := app.Listen(listenAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
