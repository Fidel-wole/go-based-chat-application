package main

import (
	"github.com/Fidel-wole/go-based-chat-application/db"
	//"github.com/Fidel-wole/go-based-chat-application/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db.InitDB()
	server := gin.Default()
	
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to chat-app backend",
		})
	})
	
	server.Run(":8080") 
}