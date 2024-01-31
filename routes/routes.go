package routes

import (
	"github.com/gin-gonic/gin"
	health "golang-gin-boilerplate/controllers/health"
)

func HealthRoute(router *gin.Engine) {
	router.GET("/health", health.Health())
}
