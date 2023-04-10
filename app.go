package main
import (
	"github.com/gofiber/fiber/v2"
	"github.com/Marbleture/api/marble"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/marbles", marble.getMarbles)
	app.Get("/marbles/:marbleId", marble.getMarble)
	app.Post("/marbles", marble.newMarble)
	app.Delete("/marbles/:marbleId", marble.deleteMarble)
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
}


// Path: app_test.go
