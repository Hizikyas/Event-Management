package routes

import (
	"example.com/rest_api/controllers"
	"github.com/gin-gonic/gin"
)

func authRoutes(route *gin.Engine) {

	route.POST("/signup" , controllers.Signup)
	route.POST("/login" , controllers.Login)
	
}