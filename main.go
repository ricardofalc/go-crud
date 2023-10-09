package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardofalc/go-crud/controllers"
	"github.com/ricardofalc/go-crud/initializers"
)

func init() { // loads before main
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.Static("/frontend", "./frontend")

	r.POST("/posts", controllers.PostsCreate)    // Create
	r.PUT("/posts/:id", controllers.PostsUpdate) // Update

	r.GET("/posts", controllers.PostsIndex)    // Read
	r.GET("/posts/:id", controllers.PostsShow) // Read

	r.DELETE("/posts/:id", controllers.PostsDelete) // Delete

	r.Run() // listen and serve on 0.0.0.0:&PORT
}
