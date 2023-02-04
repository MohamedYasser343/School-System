package Database

import (
	"Snap/Models"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

var DBAddress string = "<user>:<password>@tcp(<ip>:3306)/<database_name>"

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("<user>:<password>@tcp(<ip>:3306)/<database_name>"), &gorm.Config{})
	if err != nil {
		panic("Couldn't Connect To The Database.")
	}
	DB = connection
	connection.AutoMigrate(&Models.User{})
}
