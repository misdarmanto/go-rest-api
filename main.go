package main

import (
	"fmt"
	"go-rest-api/Config"
	"go-rest-api/Helpers"
	"go-rest-api/Models"
	"go-rest-api/Routes"

	"github.com/jinzhu/gorm"
)

func main() {
	Helpers.LoadEnv()

	var err error
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.UserModel{})
	Config.DB.AutoMigrate(&Models.PhotoModel{})

	Routes.UserRouter().Run()

}
