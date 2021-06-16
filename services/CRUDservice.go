package services

import (
	"api/database"
	"api/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func pwd(user string) string {
	pwd := user
	b := []byte(pwd)

	user = hashing(b)
	return user
}
func hashing(pwd []byte) string {

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
func ShowusersService() []model.Users {
	db := database.DBconnection()
	var usersData []model.Users
	db.Find(&usersData)
	return usersData

}

//Showuser ...
func ShowuserService(id string) model.Users {
	userid := id
	db := database.DBconnection()
	var userData model.Users
	db.Find(&userData, userid)
	return userData
}

//Create ...
func CreateService(newUser *model.Users) model.Users {

	db := database.DBconnection()

	newUser.Password = pwd(newUser.Password)

	db.Create(&newUser)
	return *newUser

}
func LoginService(user model.Users) string {
	db := database.DBconnection()

	var data model.Users
	Email := user.Email
	db.Where("email = ?", Email).Find(&data)
	fmt.Println(user.Email)
	fmt.Println(data.Email)
	if data.Email == "" {

		return "No user Found with Email"
	}
	if false == comparePass(data.Password, []byte(user.Password)) {
		return "Incorrect Password"

	}
	return "logged in"
}

//Update ...
func UpdateService(id string, user model.Users) string {
	db := database.DBconnection()
	var data model.Users
	db.First(&data, id)
	if data.Name == "" {

		return "No user Found with ID"
	}
	db.Model(&data).Updates(model.Users{Name: user.Name, Email: user.Email})

	return " user successfully updated"
}

//Delete ...
func DeleteService(id string) string {

	db := database.DBconnection()

	var user model.Users
	db.First(&user, id)
	if user.Name == "" {
		return "No user Found with ID"

	}
	db.Delete(&user)
	return "Successfully deleted"
}
