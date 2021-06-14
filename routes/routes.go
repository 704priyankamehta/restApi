package routes

import (
	"api/database"
	"api/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hash(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	return string(hash)
}
func comparePass(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

//Showusers ...
func Showusers(c *fiber.Ctx) error {
	db := database.DBconnection()
	var usersData []model.Users
	db.Find(&usersData)
	return c.JSON(usersData)
}

//Showuser ...
func Showuser(c *fiber.Ctx) error {
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
	pwd := newUser.Password
	b := []byte(pwd)

	newUser.Password = hash(b)

	db.Create(&newUser)
	return c.JSON(newUser)

}
func Login(c *fiber.Ctx) error {
	db := database.DBconnection()
	user := new(model.Users)
	c.BodyParser(user)

	var data model.Users
	Email := user.Email
	db.Where("email = ?", Email).Find(&data)
	fmt.Println(user.Email)
	fmt.Println(data.Email)
	if data.Email == "" {

		return c.Status(500).SendString("No user Found with Email")
	}
	if false == comparePass(data.Password, []byte(user.Password)) {
		return c.Status(500).SendString("Incorrect Password")

	}
	return c.SendString("logged in")
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
	db.Model(&data).Updates(model.Users{Name: user.Name, Email: user.Email})

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
