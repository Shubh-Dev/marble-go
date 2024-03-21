package main

import (
	"github.com/Shubh-Dev/marble-go/database"
	"github.com/Shubh-Dev/marble-go/marble"
	"github.com/gofiber/fiber/v2"
)

// @title			marbleture API
// @version		1.0
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/marble", marble.GetMarbles)
	app.Get("/api/v1/marble/:id", marble.GetMarble)
	app.Post("/api/v1/marble", marble.NewMarble)
	app.Delete("/api/v1/marble/:id", marble.DeleteMarble)
}

func main() {
	app := fiber.New()
	database.ConnectDB()
	setUpRoutes(app)

	app.Listen(":3000")
}
