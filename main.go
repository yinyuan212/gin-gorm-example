package main

import (
	"gin-gorm-example/controllers"
	"gin-gorm-example/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.POST("", controllers.Post)
	r.GET("", controllers.GetAll)
	r.GET(":id", controllers.GetById)
	r.PUT(":id", controllers.Update)
	r.DELETE(":id", controllers.Delete)

	r.Run()
}
