package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func HandleHealthCheck(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Up with 💚 by github.com/zcubbs")
	return c.SendString(msg)
}
