package main

import (
	"api/database"
	"api/routes"

	"github.com/gofiber/fiber/v2"
)

func setroutes(app *fiber.App) {
	app.Get("/showuser/:id", routes.GetBook)
	app.Post("/create", routes.Create)
	app.Get("/showusers", routes.GetBooks)
	app.Delete("/delete/:id", routes.Delete)
	app.Put("/update/:id", routes.Update)
}

func main() {
	database.DBconnection()

	app := fiber.New()

	setroutes(app)

	app.Listen(":3000")
}
