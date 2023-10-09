package main

import (
	"github.com/ricardofalc/go-crud/initializers"
	"github.com/ricardofalc/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
