package main
import (
	"github.com/gofiber/fiber/v2"
	"github.com/Marbleture/api/marble"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/marble", marble.GetMarbles)
	app.Get("/api/v1/marble/:id", marble.GetMarble)
	app.Post("/api/v1/marble", marble.NewMarble)
	app.Delete("/api/v1/marble/:id", marble.DeleteMarble)
}

func main() {
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(":3000")
}


// Path: app_test.go
