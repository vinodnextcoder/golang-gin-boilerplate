package main

import (
  "github.com/gin-gonic/gin"
  health "golang-gin-boilerplate/controllers/health"
)

func main() {
  r := gin.Default()
  r.GET("/health", health.Health())
  r.Run() 
}