package main

import (
	"github.com/Fidel-wole/go-based-chat-application/db"
	"github.com/Fidel-wole/go-based-chat-application/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to chat-app backend",
		})
	})
	routes.RegisterRoutes(server)
	server.Run(":8000") 
}