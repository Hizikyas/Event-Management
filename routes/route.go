package routes

import "github.com/gin-gonic/gin"

func RegisterRoute(route *gin.Engine) {
	route.GET("/events", getEventHandler)
	route.POST("/events", createEventHandler)
	route.GET("/events/:id", getEventByIdHandler)
	route.PUT("/events/:id" , updateEventHandler)
	route.DELETE("/events/:id" , DeleteEvent)
	route.POST("/signup" , signup)
}