package main

import (
	"github.com/Marbleture/api/database"
	_ "github.com/Marbleture/api/docs"
	"github.com/Marbleture/api/marble"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
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
    app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: "http://localhost:3000/swagger/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
	}))

	app.Listen(":3000")
}


// Path: app_test.go
