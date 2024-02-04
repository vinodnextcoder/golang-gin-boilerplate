package routes

import (
	"github.com/gin-gonic/gin"
	users "golang-gin-boilerplate/controllers/users"
)

func UserRoute(router *gin.Engine) {
	router.POST("/create", users.CreateUser())
}
