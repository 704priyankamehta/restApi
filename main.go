package main

import (
	"api/routes"

	"github.com/gofiber/fiber/v2"
)

func setroutes(app *fiber.App) {
	app.Get("/showuser/:id", routes.Showuser)
	app.Post("/register", routes.Create)
	app.Get("/showusers", routes.Showusers)
	app.Delete("/delete/:id", routes.Delete)
	app.Put("/update/:id", routes.Update)
	app.Post("/upload", routes.Uploadfile)
	app.Post("mailer", routes.Mail)
	app.Get("/login", routes.Login)
}

func main() {

	app := fiber.New()

	setroutes(app)

	app.Listen(":3000")
}
