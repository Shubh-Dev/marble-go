package main

import (

	"github.com/Marbleture/api/database"
	"github.com/Marbleture/api/marble"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/marbleture/api/docs/doc.go"
)

//	@title			marbleture API
//	@version		1.0
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

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


// Path: app_test.go
