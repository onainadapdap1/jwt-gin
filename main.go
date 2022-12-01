package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// set gin router
	r := gin.Default()

	//set group of public route, endpoints that will be used as authentication 
	public := r.Group("/api") 

	// public.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "successfull running"})
	// })
	
	public.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint"})
	})
	
	r.Run(":8080")
}