// package login

// import (
// 	"golang-gin-boilerplate/services/logservice"
// 	"net/http"

// 	_ "golang-gin-boilerplate/docs"

// 	"github.com/gin-gonic/gin"
// 	"github.com/sirupsen/logrus"
// )

// // helloCall godoc
// // @Summary hellow example
// // @Schemes
// // @Description Hello
// // @Tags example
// // @Accept json
// // @Produce json
// // @Success 200 {string} Hello, You created a Web App!
// // @Router /health [get]
// func Health() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		logservice.Info("Health endpoint called", logrus.Fields{"status": "200", "functionName": "Health", "controller": "healthcontroller"})
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Ok",
// 		})
// 	}
// }
