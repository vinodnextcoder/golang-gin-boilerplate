package routes

import (
	users "golang-gin-boilerplate/controllers/users"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/create", users.CreateUser())
	router.PUT("/update/:id", users.UpdateUser())
	router.GET("/users", users.GetUsers())
}
