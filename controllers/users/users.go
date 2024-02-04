package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "golang-gin-boilerplate/models"
	// "golang-gin-boilerplate/database"
)

type CreateBookInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}


// POST /users
// Create new user
// func CreateUser(c *gin.Context) {
	// Validate input
	// var input CreateBookInput
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// // Create user
	// user := models.User{Name: input.Name, Email: input.Email}
	// database.DB.Create(&user)

	// c.JSON(http.StatusOK, gin.H{"data": null})
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Ok",
// 	  })
// }

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

    c.JSON(http.StatusOK, gin.H{
      "message": "Ok",
    })
	}
}
