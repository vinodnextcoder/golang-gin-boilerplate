package controllers

import (
	"fmt"
	"golang-gin-boilerplate/database"
	"golang-gin-boilerplate/models"
	responses "golang-gin-boilerplate/services/apiresponse"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser create
// @Summary create user
// @Schemes
// @Description  create user
// @Tags userCreate
// @Accept json
// @Param       user body     models.User true "User Data"
// @Produce json
// @Success 201 {string} Success
// @Router /users/v1/create [post]
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Error: Invalid input", Data: ""})
			return
		}
		fmt.Println(input.Password)
		userCreate := models.User(input)
		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
		if err != nil {
			fmt.Println("something went wrong in db query")
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Error: Password error", Data: ""})
			panic(err)
		}
		fmt.Println("j", hash)
		input.Password = string(hash)
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

// @Summary update user
// @Schemes
// @Description  update user
// @Tags userUpdate
// @Param       id path  int  true "User Id"
// @Param       user body     models.User true "User Data"
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Router /users/v1/update/:id [put]
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := database.Db.Where("Id = ?", c.Param("id")).First(&user).Error; err != nil {

			return
		}
		// Validate input
		var input models.UserUpdate
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		database.Db.Model(&user).Updates(input)

		c.JSON(http.StatusOK, responses.SuccesResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

// @Summary get users list
// @Schemes
// @Description  get users list
// @Tags userList
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Router /users/v1/users [get]
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user []models.User

		err := database.Db.Find(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		}

		c.JSON(http.StatusOK, responses.SuccesResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

// @Summary delete user
// @Schemes
// @Description  delete user
// @Tags deleteUser
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Router /users/v1/delete/:id [delete]
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		result := database.Db.Where("Id = ?", c.Param("id")).Delete(&user)

		if result.Error != nil {
			// Handle database error
			fmt.Println("something went wrong in db query", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		if result.RowsAffected == 0 {
			// No record was deleted, return a not found response
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}

		// Record deleted successfully
		c.JSON(http.StatusOK, responses.SuccesResponse{
			Status:  http.StatusOK,
			Message: "User deleted successfully",
			Data:    map[string]interface{}{"data": nil},
		})
	}
}
