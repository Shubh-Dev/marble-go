package main

import (
	"github.com/gofiber/fiber"
	// "github.com/gofiber/fiber/v2"
)


func helloMarbel(c *fiber.Ctx) {
	c.Send("Hello, Marbel")

}



func main() {
	app :=fiber.New()

	app.Get("/", helloMarbel)
	
	app.Listen(3000)

}