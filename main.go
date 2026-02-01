package main

import (
	"net/http"

	"example.com/rest_api/db"
	"example.com/rest_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
server := gin.Default()
db.InitDB()
db.CreateTable()
server.GET("/events" , exactHandler )
server.POST("/events" , createEventHandler )
server.Run(":8080")
}

func exactHandler(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK , gin.H{"event" : events }) //or it can be context.JSON(200 , gin.H{"message" : "Hello"})
}

func createEventHandler(context *gin.Context) {
event := models.Event{} // literal notation, or can be var event models.Event
err := context.ShouldBindJSON(&event)

if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message" : "Error in binding data" })
	return
}
event.ID = 1
event.UserID = 1
event.Save()
context.JSON(http.StatusCreated , gin.H{"message" : "Event created" , "event" : event})
}