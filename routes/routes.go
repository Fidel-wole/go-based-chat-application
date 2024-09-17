package routes

import (
	"github.com/Fidel-wole/go-based-chat-application/controller"
	"github.com/gin-gonic/gin"
	//"github.com/Fidel-wolee/event-booking/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	// Public routes
	server.POST("/signup", controller.CreateUser)
	server.POST("/login", controller.LoginUser)

	// Protected routes
	//authorized := server.Group("/")

}

