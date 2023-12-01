package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitializeDB() {
	// variables declared in .env
	var user string = "root"
	var password string = os.Getenv("DATABASE_ROOT_PASSWORD")
	var db_host string = os.Getenv("DATABASE_CONTAINER_NAME")
	var db_name string = os.Getenv("DATABASE_NAME")
	var db_port string = os.Getenv("DATABASE_PORT")

	// username:password@protocol(address)/dbname?param=value
	var dsn string = user + ":" + password + "@tcp(" + db_host + ":" + db_port + ")" + "/" + db_name

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		println(err)
		panic("Couldn't connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&User{})
	connection.AutoMigrate(&Event{})
}
