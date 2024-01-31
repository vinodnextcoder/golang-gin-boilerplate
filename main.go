package main

import (
  "github.com/gin-gonic/gin"
  "golang-gin-boilerplate/routes"
)

func main() {
  router := gin.Default()
  routes.HealthRoute(router)
  var port = "3001"
  router.Run("0.0.0.0:" + port)
}