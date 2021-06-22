package services

import (
	database "api/database"
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
	var usersData []model.Users
	sql := "SELECT id,name,email FROM USERS"
	database.DBcon.Raw(sql).Scan(&usersData)
	return usersData

}

//ShowuserService ...
func ShowuserService(id string) model.Users {
	userid := id
	if userid == "" {
		fmt.Println("invalid ID")
	}

	var userData model.Users
	database.DBcon.Find(&userData, userid)
	fmt.Println(userData)
	return userData
}

//Create ...
func CreateService(newUser model.Users) string {

	newUser.Password = pwd(newUser.Password)

	err := database.DBcon.Create(&newUser).Error
	if err != nil {
		return "Email/ID already exist"
	}
	return "registered successfully"

}
func LoginService(user model.Users) string {

	var data model.Users
	Email := user.Email
	println(Email)
	database.DBcon.Where("email = ?", Email).Find(&data)
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

	var data model.Users
	database.DBcon.First(&data, id)
	if data.Name == "" {

		return "No user Found with ID"
	}
	database.DBcon.Model(&data).Updates(model.Users{Name: user.Name, Email: user.Email})

	return "user successfully updated"
}

//Delete ...
func DeleteService(id string) string {

	var user model.Users
	database.DBcon.First(&user, id)
	if user.Name == "" {
		return "No user Found with ID"

	}
	database.DBcon.Delete(&user)
	return "Successfully deleted"
}
