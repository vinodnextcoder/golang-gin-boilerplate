package login

import (
	"fmt"
	"golang-gin-boilerplate/database"
	_ "golang-gin-boilerplate/docs"
	"golang-gin-boilerplate/models"
	responses "golang-gin-boilerplate/services/apiresponse"
	"golang-gin-boilerplate/services/logservice"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": existingUser.Email,
			"id":   existingUser.Id,
			"sub":  existingUser.Id,
			"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		tokenStr, err := token.SignedString([]byte("supersaucysecret"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create token",
			})
			return
		}

		Refreshtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": existingUser.Email,
			"id":   existingUser.Id,
			"sub":  existingUser.Id,
			"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		refreshtoken, err := Refreshtoken.SignedString([]byte("supersaucysecret"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create token",
			})
			return
		}
		var res models.SignedResponse
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("access_token", tokenStr, 3600*24*30, "", "", false, true)
		c.SetCookie("refresh_token", refreshtoken, 3600*24*30, "", "", false, true)
		res.Access_token = tokenStr
		res.Refresh_token = refreshtoken
		c.JSON(http.StatusOK, responses.SuccesResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}})

	}
}
