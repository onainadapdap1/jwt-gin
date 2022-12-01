package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/jwt-gin/models"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	//set objek input 
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// deklarasi dan inisialisasi objek User
	u := models.User{Username: input.Username, Password: input.Password}
	// u.Username = input.Username
	// u.Password = input.Password

	u.Prepare()

	if err := u.BeforeSave(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := u.SaveUser();
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration success", "data": res})
}