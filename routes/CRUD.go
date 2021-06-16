package routes

import (
	"api/services"
"api/model"
	"github.com/gofiber/fiber/v2"
)

//Showusers ...
func Showusers(c *fiber.Ctx) error {

	result := services.ShowusersService()
	return c.JSON(result)

}

func Showuser(c *fiber.Ctx) error {
	id:=c.Params("id")
	result := services.ShowuserService(id)
	return c.JSON(result)

}
func Create(c *fiber.Ctx) error {
	newUser := new(model.Users)

	if err := c.BodyParser(newUser); err != nil {
		c.Status(503).SendString("error")

	}
	result := services.CreateService(newUser)
	return c.JSON(result)

}
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	result := services.DeleteService(id)
	return c.JSON(result)

}
func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(model.Users)
	c.BodyParser(user)
	result := services.UpdateService(id,*user)
	return c.JSON(result)

}
func Login(c *fiber.Ctx) error {
	user := new(model.Users)
	c.BodyParser(user)
	result := services.LoginService(*user)
	return c.JSON(result)

}
