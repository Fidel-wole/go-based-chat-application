package controller

import (
	"net/http"

	db "github.com/Fidel-wole/go-based-chat-application/db"
	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context){
	var roomData struct{
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&roomData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data", "error": err.Error()})
		return
	} 

	queries := db.GetQueries()
   room, err := queries.CreateRoom(c, roomData.Name)
   if err !=nil{
	c.JSON(http.StatusInternalServerError, gin.H{"message":"Error creating room"})
   }
   c.JSON(http.StatusCreated, gin.H{"message":"Room created", "room":room})
}