package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getEnvvars() {
	err := godotenv.Load("credencial.env")
	if err != nil {
		fmt.Println(err)
	}
}

var DBcon *gorm.DB

//DBconnection ...
func DBconnection() {
	//var users model.Users
	getEnvvars()
	database := os.Getenv("DATABASE")

	dsn := database
	var err error
	DBcon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("connected")
	}
	//db.AutoMigrate(&users)

}
