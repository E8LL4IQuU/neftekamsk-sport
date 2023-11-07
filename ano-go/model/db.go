package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() {
	// Export dns to dotenv
	dsn := "root:miumimiclairno@tcp(vue-db-1:3307)/ano"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		println(err)
		panic("Couldn't connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&User{})
}
