package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setroutes(app *fiber.App) {
	app.Get("api/v1/showuser/:id", Showuser)

	//register User
	app.Post("api/v1/register", Create)

	//paging paerpage 2
	app.Get("api/v1/getusers", Getusers)
	app.Get("api/v1/showusers", Showusers)
	app.Delete("api/v1/delete/:id", Delete)
	app.Put("api/v1/update/:id", Update)
	//app.Post("/upload", services.Uploadfile)
	//app.Post("mailer", services.Mail)
	app.Get("/login", Login)

}
