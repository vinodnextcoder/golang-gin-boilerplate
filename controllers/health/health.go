package health

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func Health() gin.HandlerFunc {
	return func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Ok",
    })
	}
}