package routes

import (
	"api/database"
	"api/model"
	"api/services"
	"fmt"
	"net/mail"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validateUserData(newUser model.Users) string {

	if newUser.Name == "" {
		return "Name is provide"
	}
	if newUser.Email == "" {
		return "Name is provide"
	}
	if newUser.Password == "" {
		return "Name is provide"
	}
	return "validated"
}

//Showusers ...
func Showusers(c *fiber.Ctx) error {

	result := services.ShowusersService()

	return c.JSON(result)

}

func Showuser(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	if id == " " {
		return c.SendString("invalid id")
	}
	result := services.ShowuserService(id)
	if result.Name == " " {
		return c.SendString("no user found for this ID")
	}
	value := []map[string]interface{}{{"name": result.Name, "email": result.Email, "ID": strconv.Itoa(result.ID)}}
	return c.JSON(value)

}
func Create(c *fiber.Ctx) error {

	newUser := new(model.Users)

	if err := c.BodyParser(newUser); err != nil {
		c.Status(503).SendString("error")

	}
	if !valid(newUser.Email) {
		return c.SendString("invalid Email")
	}
	validate := validateUserData(*newUser)
	if validate != "validated" {
		return c.SendString(validate)
	}
	result := services.CreateService(*newUser)
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
	result := services.UpdateService(id, *user)
	return c.JSON(result)

}
func Login(c *fiber.Ctx) error {
	user := new(model.Users)
	c.BodyParser(user)
	result := services.LoginService(*user)
	return c.JSON(result)

}
func Getusers(c *fiber.Ctx) error {
	var users []model.Users

	sql := "SELECT id,name,email FROM USERS"
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendString("error")
	}
	perPage := 5
	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)

	database.DBcon.Raw(sql).Scan(&users)
	return c.JSON(users)
}
