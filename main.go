package main

import (
	"baro-backend/config"
	"baro-backend/docs"
	_ "baro-backend/docs"
	"baro-backend/middlewares"
	"baro-backend/routes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8421"
	}

	// DB setup
	client, err := config.InitMongoDB()
	if err != nil {
		log.Fatalf("Mongo Connection Error: %v\n", err)
	}
	db := client.Database(os.Getenv("MONGODB_DB"))

	// Initialize Gin router
	router := gin.Default()
	router.Use(middlewares.TimeoutMiddleware(5 * time.Second))
	routes.SetupRoutes(router, db)

	// Initialize Swagger
	if os.Getenv("APP_ENV") == "dev" {
		docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("PORT"))
		docs.SwaggerInfo.Title = "Baro Backend"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Run the server
	router.Run(":" + port)
}
