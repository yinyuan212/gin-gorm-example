package config

import (
	"gin-gorm-example/controllers"

	"github.com/gin-gonic/gin"
)

func RoutePost(r *gin.Engine) {
	r.POST("", controllers.Post)
	r.GET("", controllers.GetAll)
	r.GET(":id", controllers.GetById)
	r.PUT(":id", controllers.Update)
	r.DELETE(":id", controllers.Delete)
}

func RouteStudent(r *gin.Engine) {
	students := r.Group("/students")
	{
		students.POST("/", controllers.CreateStudent)
		students.GET("/", controllers.GetAllStudents)
	}
}
