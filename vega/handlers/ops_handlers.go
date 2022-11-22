package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func HandleHealthCheck(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Up with 💚 by @zcubbs!")
	return c.SendString(msg)
}
