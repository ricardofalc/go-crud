package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ricardofalc/go-crud/initializers"
)

func init() { // loads before main
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("Hello, World!123")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
