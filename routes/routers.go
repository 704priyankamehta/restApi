package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setroutes(app *fiber.App) {
	app.Get("/showuser/:id", Showuser)
	app.Post("/register", Create)
	app.Get("/showusers", Showusers)
	app.Delete("/delete/:id", Delete)
	app.Put("/update/:id", Update)
	//app.Post("/upload", services.Uploadfile)
	//app.Post("mailer", services.Mail)
	app.Get("/login", Login)

}
