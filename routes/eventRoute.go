package routes

import (
	"example.com/rest_api/controllers"
	"example.com/rest_api/middleware"
	"github.com/gin-gonic/gin"
)

func eventRoutes(route *gin.Engine) {
	
 	route.GET("/events", controllers.GetEventHandler)
	route.GET("/events/:id", controllers.GetEventByIdHandler)	

	authenticate :=route.Group("/")
		authenticate.Use(middleware.AuthMiddleware)
		authenticate.POST("/events", controllers.CreateEventHandler)
		authenticate.PUT("/events/:id" , controllers.UpdateEventHandler)
		authenticate.DELETE("/events/:id" , controllers.DeleteEvent)
}