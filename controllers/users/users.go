package controllers

import (
	"fmt"
	"golang-gin-boilerplate/database"
	"golang-gin-boilerplate/models"
	responses "golang-gin-boilerplate/services/apiresponse"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Error: Invalid input", Data: ""})
			return
		}
		fmt.Println(input)
		userCreate := models.User(input)
		result := database.Db.Create(&userCreate)
		if result.Error != nil {
			fmt.Println("something went wrong in db query")
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Error: Db query failed", Data: ""})
			return
		}
		fmt.Println("record inserted successfully ", result.RowsAffected)

		c.JSON(http.StatusOK, responses.SuccesResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": input}})
	}
}
