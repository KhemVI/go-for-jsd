package routes

import (
	"baro-backend/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *gin.Engine, db *mongo.Database) {
	userCollection := db.Collection("users")
	userController := controllers.UserController{Collection: userCollection}

	// Set up the /users routes group
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/", userController.GetUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}
}
