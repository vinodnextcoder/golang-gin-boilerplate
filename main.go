package main

import (
	"golang-gin-boilerplate/routes"
	"golang-gin-boilerplate/services/logservice"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"

	// "github.com/sirupsen/logrus"
	"golang-gin-boilerplate/database"
	_ "golang-gin-boilerplate/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample gin web server

// @contact.name   vinod
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:3001
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	logservice.InitLogger()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDb()
	logservice.Info("golang app started")
	port := os.Getenv("PORT")
	router := gin.Default()
	routes.HealthRoute(router)
	routes.UserRoute(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("0.0.0.0:" + port)
}
