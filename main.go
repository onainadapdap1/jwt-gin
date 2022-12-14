package main

import (
	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/jwt-gin/controllers"
	"github.com/onainadapdap1/jwt-gin/models"
)

func main() {
	// connect to database mysql
	models.ConnectDatabase()
	// set gin router
	r := gin.Default()

	//set group of public route, endpoints that will be used as authentication
	public := r.Group("/api")

	// public.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "successfull running"})
	// })

	public.POST("/register", controllers.Register)

	r.Run(":8080")
}
