package database

import (
	"api/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DBconnection ...
func DBconnection() *gorm.DB {
	var users model.Users

	dsn := "+++:****@tcp(127.0.0.1:3306)/Data"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("connected")
	}
	db.AutoMigrate(&users)
	return db
}
