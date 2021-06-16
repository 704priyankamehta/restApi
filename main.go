package main

import (
	"api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	routes.Setroutes(app)

	app.Listen(":3000")
}
