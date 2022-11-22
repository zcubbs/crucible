package handlers

import (
	"crucible/vega/configs"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func Login(c *fiber.Ctx) error {
	user := c.FormValue("username")
	pass := c.FormValue("password")

	// Throws Unauthorized error
	if user != configs.Config.API.Username || pass != configs.Config.API.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": fmt.Sprintf("%s", configs.Config.API.Username),
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(configs.Config.API.TokenSecret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
