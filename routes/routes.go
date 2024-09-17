package routes

import (
	"github.com/Fidel-wole/go-based-chat-application/controller"
	"github.com/Fidel-wole/go-based-chat-application/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Public routes
	server.POST("/signup", controller.CreateUser)
	server.POST("/login", controller.LoginUser)
    
	// Protected routes
	authorized := server.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/create-room", controller.CreateRoom)
		authorized.POST("/message/:roomId", controller.SendMessage)
		authorized.GET("/room/:roomId", controller.GetMessagesByRoom)
	}
}
