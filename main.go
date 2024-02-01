package main

import (
  "os"
  "log"
  "github.com/lpernett/godotenv"
  "github.com/gin-gonic/gin"
  "golang-gin-boilerplate/routes"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  port := os.Getenv("PORT")
  router := gin.Default()
  routes.HealthRoute(router)
  router.Run("0.0.0.0:" + port)
}