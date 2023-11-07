package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() {
	// TODO: Export dns to dotenv, including database name, user, pass, port
	dsn := "root:Kikuri@tcp(sport_db:3306)/database"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		println(err)
		panic("Couldn't connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&User{})
}
