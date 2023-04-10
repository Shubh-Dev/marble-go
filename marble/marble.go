package marble

import (
	"github.com/gofiber/fiber/v2"
)

func getMarbles(c *fiber.Ctx) error {
	return c.SendString("Get all marbles")
}

func getMarble(c *fiber.Ctx) error {
	return c.SendString("Get a single marble")
}

func newMarble(c *fiber.Ctx) error {
	return c.SendString("create New marble")
}

func deleteMarble(c *fiber.Ctx) error {
	return c.SendString("Delete a single marble")
}

// Path: main.go

