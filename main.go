package main

import (
	"api/database"
	"api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DBconnection()
	app := fiber.New()

	routes.Setroutes(app)

	app.Listen(":3000")
}
