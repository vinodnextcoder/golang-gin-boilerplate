package routes

import (
	login "golang-gin-boilerplate/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/auth/login", login.Login())
}
