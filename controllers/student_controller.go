package controllers

import (
	"gin-gorm-example/initializers"
	"gin-gorm-example/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {

	var request struct {
		Name   string
		Gender string
		Age    uint
	}

	c.Bind(&request)

	student := models.Student{Name: request.Name, Gender: request.Gender, Age: request.Age}
	result := initializers.DB.Create(&student)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "success",
		"data":   student,
	})
}

func GetAllStudents(c *gin.Context) {
	students := make([]models.Student, 0)
	initializers.DB.Find(&students)

	c.JSON(http.StatusOK, gin.H{
		"students": students,
	})

}
