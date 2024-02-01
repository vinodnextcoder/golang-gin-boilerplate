package health

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
  "golang-gin-boilerplate/services/logservice"
)

func Health() gin.HandlerFunc {
	return func(c *gin.Context) {
    logservice.Info("error code", logrus.Fields{"status":"200"})
    c.JSON(http.StatusOK, gin.H{
      "message": "Ok",
    })
	}
}