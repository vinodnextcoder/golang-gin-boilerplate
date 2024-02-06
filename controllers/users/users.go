package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang-gin-boilerplate/models"
	"golang-gin-boilerplate/database"
	"golang-gin-boilerplate/services/apiresponse"
	
)



func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
     	var input models.User
	    if err := c.ShouldBindJSON(&input); err != nil {
	   	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	 	return
	  }
	   fmt.Println(input)
	   userCreate := models.User(input)
	   result := database.Db.Create(&userCreate)
	   if result.Error != nil {
		   fmt.Println("something went wrong in db query")
		   return
	   }
	   fmt.Println("record inserted successfully ", result.RowsAffected)
	   
	   c.JSON(http.StatusOK, responses.SuccesResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": input}})
	}
}
