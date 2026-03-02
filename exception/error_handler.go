package exception

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := err.Error()

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"code":   code,
		"status": getHttpStatusText(code),
		"data":   message,
	})
}

func getHttpStatusText(code int) string {
	switch code {
	case fiber.StatusBadRequest:
		return "BAD_REQUEST"
	case fiber.StatusNotFound:
		return "NOT_FOUND"
	case fiber.StatusInternalServerError:
		return "INTERNAL_SERVER_ERROR"
	default:
		return "ERROR"
	}
}
