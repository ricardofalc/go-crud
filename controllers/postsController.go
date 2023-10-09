package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ricardofalc/go-crud/initializers"
	"github.com/ricardofalc/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return post
	c.JSON(200, gin.H{
		"message": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get all posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get post by id
	var post models.Post
	initializers.DB.First(&post, id)

	// Return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get data off request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find post to update
	var post models.Post
	initializers.DB.First(&post, id)

	// Update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Return post
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Find post to delete
	var post models.Post
	initializers.DB.First(&post, id)

	// Delete post
	initializers.DB.Delete(&post)

	// Return post
	c.JSON(200, gin.H{
		"message": "Post " + strconv.Itoa(int(post.ID)) + " deleted",
	})

}
