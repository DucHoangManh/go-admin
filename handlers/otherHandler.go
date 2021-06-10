package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Others(c *fiber.Ctx) error {
	return c.SendString("THis is other handler")
}
