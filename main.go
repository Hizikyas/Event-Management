// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"example.com/rest_api/db"
// 	"example.com/rest_api/models"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// server := gin.Default()
// db.InitDB()
// server.GET("/events" , getEventHandler )
// server.POST("/events" , createEventHandler )
// server.GET("/events/:id" , getEventByIdHandler )
// server.Run(":8080")
// }

// func getEventHandler(context *gin.Context) {
// 	events , err := models.GetAllEvents()
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError , gin.H{"message" : "Couldn't fetch the event. Try again hi"})
// 		return
// 	}
// 	context.JSON(http.StatusOK , gin.H{"events" : events }) //or it can be context.JSON(200 , gin.H{"message" : "Hello"})
// }

// func createEventHandler(context *gin.Context) {
// 	event := models.Event{} // literal notation, or can be var event models.Event
// 	err := context.ShouldBindJSON(&event)

// 		if err != nil {
// 			context.JSON(http.StatusBadRequest, gin.H{"message" : "Error in binding data" })
// 			return
// 		}
// 	event.UserID = 1
// 	err = event.Save()

// 		if err != nil {
// 			context.JSON(http.StatusInternalServerError , gin.H{"message" : "Could not save event. Try again"})
// 			return
// 		}
// 	context.JSON(http.StatusCreated , gin.H{"message" : "Event created" , "event" : event})
// }

// func getEventByIdHandler (context *gin.Context) {
// 	eventId, err := strconv.ParseInt(context.Param("id") , 10 , 64)

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest , gin.H{"message" : "could not parse the event id"})
// 		return
// 	}
//    event , err := models.GetEvent(eventId)
//     if err != nil {
// 		fmt.Println(err)
// 		context.JSON(http.StatusInternalServerError , gin.H{"message" : "Could not fetch the event."})
// 		return
// 	}

// 	context.JSON(http.StatusOK , event)
// }

package main

import (
	"example.com/rest_api/db"
	"example.com/rest_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
server := gin.Default()
db.InitDB()

routes.RegisterRoute(server)
server.Run(":8080")
}