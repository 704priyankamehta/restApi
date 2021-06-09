package routes

import (
	"api/database"
	"api/model"

	"github.com/gofiber/fiber/v2"
)

//GetBooks ...
func GetBooks(c *fiber.Ctx) error {
	db := database.DBconnection()
	var usersData []model.Users
	db.Find(&usersData)
	return c.JSON(usersData)
}

//GetBook ...
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconnection()
	var userData model.Users
	db.Find(&userData, id)
	return c.JSON(userData)
}

//Create ...
func Create(c *fiber.Ctx) error {

	db := database.DBconnection()
	newUser := new(model.Users)
	if err := c.BodyParser(newUser); err != nil {
		c.Status(503).SendString("error")

	}
	db.Create(&newUser)
	return c.JSON(newUser)

}

//Update ...
func Update(c *fiber.Ctx) error {
	db := database.DBconnection()
	id := c.Params("id")
	user := new(model.Users)
	c.BodyParser(user)

	var data model.Users
	db.First(&data, id)
	if data.Name == "" {

		return c.Status(500).SendString("No user Found with ID")
	}
	db.Model(&data).Updates(model.Users{Name: user.Name, Phone: user.Phone})

	return c.SendString(" user successfully updated")
}

//Delete ...
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBconnection()

	var user model.Users
	db.First(&user, id)
	if user.Name == "" {
		return c.Status(500).SendString("No user Found with ID")

	}
	db.Delete(&user)
	return c.SendString("Successfully deleted")
}
