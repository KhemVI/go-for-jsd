package controllers

import (
	"fmt"
	"net/http"
	"time"

	"baro-backend/common"
	"baro-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	Collection *mongo.Collection
}

// @Summary      Create a new user
// @Description  This endpoint creates a new user with the provided details.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body common.CreateUserRequest true "User details"
// @Success      200  {object}  models.User
// @Failure      400  {object}  common.ErrorResponse
// @Router       /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var request common.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user := models.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := uc.Collection.InsertOne(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Get Users
func (uc *UserController) GetUsers(c *gin.Context) {
	var users []models.User

	cursor, err := uc.Collection.Find(c.Request.Context(), bson.M{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

// @Summary      Get User by ID
// @Description  This is description
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      404  {object}  common.ErrorResponse
// @Router       /users/{id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var user models.User

	err := uc.Collection.FindOne(c.Request.Context(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary      Delete user
// @Description  This endpoint deletes.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  common.SuccessResponse
// @Failure      400  {object}  common.ErrorResponse
// @Router       /users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	result, err := uc.Collection.DeleteOne(c.Request.Context(), bson.M{"_id": objID})
	if err != nil || result.DeletedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User deleted id: %s", id)})
}
