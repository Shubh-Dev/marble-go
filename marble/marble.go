package marble

import (
	"github.com/gofiber/fiber/v2"
)

func GetMarbles(c *fiber.Ctx) error {
	return c.SendString("Get all marbles")
}

func GetMarble(c *fiber.Ctx) error {
	return c.SendString("Get a single marble")
}

func NewMarble(c *fiber.Ctx) error {
	return c.SendString("create New marble")
}

func DeleteMarble(c *fiber.Ctx) error {
	return c.SendString("Delete a single marble")
}


