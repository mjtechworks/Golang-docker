package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/go-api/controllers"
	"github.com/go-api/models"
)

func main() {
	r := gin.Default()

	db := models.SetupModels() // new

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", controllers.GetAllUser)
	r.Run()
}
