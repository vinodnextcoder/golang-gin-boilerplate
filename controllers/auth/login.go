package login

import (
	"fmt"
	"golang-gin-boilerplate/database"
	_ "golang-gin-boilerplate/docs"
	"golang-gin-boilerplate/models"
	responses "golang-gin-boilerplate/services/apiresponse"
	"golang-gin-boilerplate/services/logservice"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary login
// @Schemes
// @Description Login
// @Tags Login
// @Accept json
// @Produce json
// @Success 200 {string} login api
// @Router /login [get]
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		logservice.Info("Login endpoint called", logrus.Fields{"status": "200", "functionName": "Login", "controller": "logincontroller"})
		var input models.UserLogin
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Error: Invalid input", Data: ""})
			return
		}
		var email = input.Email
		fmt.Println(input)
		var existingUser models.User
		if err := database.Db.Where("Email = ?", email).First(&existingUser).Error; err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Error: User not found", Data: ""})
			return
		}
		fmt.Println(existingUser)
		err1 := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(input.Password))
		if err1 != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Error: Invalid email or password", Data: ""})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	}
}
