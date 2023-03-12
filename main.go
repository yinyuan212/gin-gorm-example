package main

import (
	"gin-gorm-example/config"
	"gin-gorm-example/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	config.RoutePost(r)
	config.RouteStudent(r)

	r.Run()
}
