package models

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306/go_restapi_gin)"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&product{})

	DB = database
}
