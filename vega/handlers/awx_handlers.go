package handlers

import (
	"crucible/core/awx"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

var client *awx.AWX

func NewClient(url, username, password string) {
	client = awx.NewAWX(url, username, password, nil)
}

func HandlePing(c *fiber.Ctx) error {
	result, err := client.PingService.Ping()
	if err != nil {
		log.Println("error while calling ping service", err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}
	_ = c.SendString(fmt.Sprintf("%v", result))
	return nil
}

func HandleRunTemplate(c *fiber.Ctx) error {
	templateId := c.Get("template_id")
	inventoryId := c.Get("inventory_id")

	intTemplateId, err := strconv.Atoi(templateId)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "template_id must be an integer")
	}

	// Run job
	job, err := client.JobTemplateService.Launch(intTemplateId, map[string]interface{}{
		"inventory": inventoryId,
	}, map[string]string{})
	if err != nil {
		log.Println("error while calling launch service", err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	_ = c.SendString(fmt.Sprintf("%v", job))

	return nil
}

func HandleGetJobEvents(c *fiber.Ctx) error {
	id := c.Get("id")
	params := map[string]string{}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "id must be an integer")
	}

	result, _, err := client.JobService.GetJobEvents(intId, params)
	if err != nil {
		log.Println("error while calling get job events service", err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	_ = c.SendString(fmt.Sprintf("%v", result))
	return nil
}
