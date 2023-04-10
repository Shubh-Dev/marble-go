package main

import (
	"fmt"

	"github.com/Marbleture/api/database"
	"github.com/Marbleture/api/marble"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/marble", marble.GetMarbles)
	app.Get("/api/v1/marble/:id", marble.GetMarble)
	app.Post("/api/v1/marble", marble.NewMarble)
	app.Delete("/api/v1/marble/:id", marble.DeleteMarble)
}

func initDatabase() {
	var err error
	 database.DBConnect, err = gorm.Open("postgreSQL", "marbles.db")
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connection successfully opened!")
}

func main() {
	app := fiber.New()
	initDatabase()
	setUpRoutes(app)
	app.Listen(":3000")
}


// Path: app_test.go
