package main

import (
  "os"
  "log"
  "github.com/lpernett/godotenv"
  "github.com/gin-gonic/gin"
  "golang-gin-boilerplate/routes"
  "golang-gin-boilerplate/services/logservice"
  // "github.com/sirupsen/logrus"
)

func main() {
  logservice.InitLogger()
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  logservice.Info("golang app started")
  port := os.Getenv("PORT")
  router := gin.Default()
  routes.HealthRoute(router)
  
  router.Run("0.0.0.0:" + port)
}