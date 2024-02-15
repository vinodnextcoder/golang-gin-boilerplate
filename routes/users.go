package routes

import (
	users "golang-gin-boilerplate/controllers/users"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	v1 := router.Group("/users/v1")
	{
		v1.POST("/create", users.CreateUser())
		v1.PUT("/update/:id", users.UpdateUser())
		v1.GET("/users", users.GetUsers())
		v1.DELETE("/delete/:id", users.DeleteUser())
	}
}
