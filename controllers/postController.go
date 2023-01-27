package controllers

import (
	"gin-gorm-example/initializers"
	"gin-gorm-example/models"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	// Get data of req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAll(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetById(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Response
	c.JSON(200, gin.H{
		"result": post,
	})

}

func Update(c *gin.Context) {
	// Get the id off url
	id := c.Param("id")
	// Get the data off req body
	var req struct {
		Title string
		Body  string
	}

	c.Bind(&req)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: req.Title,
		Body:  req.Body,
	})

	// Response
	c.JSON(200, gin.H{
		"result": req,
	})

}

func Delete(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	var post models.Post

	initializers.DB.Delete(&post, id)

	c.JSON(200, gin.H{
		"message": "success",
	})

}
