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

//DBconnection ...
func DBconnection() *gorm.DB {
	//var users model.Users
	getEnvvars()
	database := os.Getenv("DATABASE")

	dsn := database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("connected")
	}
	//db.AutoMigrate(&users)
	return db
}
