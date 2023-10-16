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

	// Serve the main HTML file from the root path
	r.LoadHTMLGlob("../frontend/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Serve static files from /static
	r.Static("/static", "../frontend/")

	r.POST("/posts", controllers.PostsCreate)    // Create
	r.PUT("/posts/:id", controllers.PostsUpdate) // Update

	r.GET("/posts", controllers.PostsIndex)    // Read
	r.GET("/posts/:id", controllers.PostsShow) // Read

	r.DELETE("/posts/:id", controllers.PostsDelete) // Delete

	r.Run() // listen and serve on 0.0.0.0:&PORT
}
