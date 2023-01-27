package main

import (
	"gin-gorm-example/initializers"
	"gin-gorm-example/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// 用來建立DB table的，需要另外執行這隻程式
	initializers.DB.AutoMigrate(&models.Post{})

}
