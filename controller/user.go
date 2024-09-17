package controller

import (
	
	"net/http"

	db1 "github.com/Fidel-wole/go-based-chat-application/db"
	db "github.com/Fidel-wole/go-based-chat-application/db/sqlc"
	"github.com/Fidel-wole/go-based-chat-application/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user db.CreateUserParams 
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data", "error": err.Error()})
		return
	}

	queries := db1.GetQueries()

	// Check if user already exists
	existingUser, err := queries.GetUserByUsername(c, user.Username)
	if err == nil && existingUser.Username != "" {
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password", "error": err.Error()})
		return
	}
	user.Password = hashedPassword

	createdUser, err := queries.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": createdUser})
}

func LoginUser(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data", "error": err.Error()})
		return
	}

	queries := db1.GetQueries()

	// Retrieve user by username
	user, err := queries.GetUserByUsername(c, loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	// Compare provided password with stored hash
	passwordIsValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !passwordIsValid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
    token, err := utils.GenerateToken(user.Username, user.ID)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user", "error": err.Error()})
        return
    }
	// Authentication successful
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}