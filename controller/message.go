package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Fidel-wole/go-based-chat-application/db"
	sqlc "github.com/Fidel-wole/go-based-chat-application/db/sqlc"
	"github.com/gin-gonic/gin"
)

type SendMessageRequest struct {
	RoomID    int64  `json:"room_id"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
}

func SendMessage(c *gin.Context) {
	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID, err := strconv.ParseInt(c.Param("roomId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert userID to int64
	userIDInt64, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	
	queries := db.GetQueries()
	
	// Check if the room exists
	_, err = queries.GetRoom(c, roomID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	message, err := queries.CreateMessage(c, sqlc.CreateMessageParams{
		RoomID:  sql.NullInt64{Int64: roomID, Valid: true},
		UserID:  sql.NullInt64{Int64: userIDInt64, Valid: true},
		Content: req.Content,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Message sent successfully",
		"data":    message,
	})
}

func GetMessagesByRoom(c *gin.Context) {
	roomID, err := strconv.ParseInt(c.Param("roomId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	queries := db.GetQueries()

	// Check if the room exists
	_, err = queries.GetRoom(c, roomID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	messages, err := queries.GetMessagesByRoom(c, sql.NullInt64{Int64: roomID, Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}