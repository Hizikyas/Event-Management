package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest_api/models"
	"github.com/gin-gonic/gin"
)

// Get all events
func GetEventHandler(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the event. Try again hi"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": events}) //or it can be context.JSON(200 , gin.H{"message" : "Hello"})
}

// POST Events
func CreateEventHandler(context *gin.Context) {
	event := models.Event{} // literal notation, or can be var event models.Event
	err := context.ShouldBindJSON(&event) // this is to recieve value from the body it is like req.body 

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error in binding data"})
		return
	}
	id := context.GetInt64("UserId")
	event.UserID = id
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event. Try again"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

// GET Events by ID
func GetEventByIdHandler(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

// Update Event by ID
 func UpdateEventHandler (context *gin.Context)  {

	var updateEvents models.Event
    eventId , err := strconv.ParseInt(context.Param("id") , 10 , 64)

	if err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"message" : "Could not parse the integer"})
		return
	}
    
	event, err := models.GetEventById(eventId)
	   if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message" : "Could not fetch the event"})
	   }
	   
   err = context.ShouldBindJSON(&updateEvents)
    if err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"message" : "Could not parse the requested data"})
		return
	}

   updateEvents.ID = event.ID
   id := context.GetInt64("UserId")
   if event.UserID != id {
	context.JSON(http.StatusUnauthorized , gin.H{"message" : "Unauthorized to perform this action"})
	 return
   }

   err = updateEvents.UpdateEvent()
   if err != nil {
	    fmt.Println(err)
		context.JSON(http.StatusInternalServerError , gin.H{"message" : "Could not Update the event"})
		return
   }

   context.JSON(http.StatusOK , gin.H{"message" : "Event Updated successfuly"})
 }

 // Delete Events by ID
 func DeleteEvent (context *gin.Context) {
	context.Param("id")
	eventId , err := strconv.ParseInt(context.Param("id") , 10 , 64)
		if err != nil {
			context.JSON(http.StatusBadRequest , gin.H{"message" : "Could not parse the event ID"})
			return
		}
	
	events , err := models.GetEventById(eventId)
		if err != nil {
			context.JSON(http.StatusInternalServerError , gin.H{"message": "Could not fetch the event"})
			return
		}
        userId := context.GetInt64("UserId")
		    if userId != events.UserID {
				context.JSON(http.StatusUnauthorized , gin.H{"message" : "Unauthorized to perform this action"})
				return
			}

	err = events.DeleteEvent()	
	if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message" : "Could not delete the event"})
		return
	}
    
	context.JSON(http.StatusOK , gin.H{"message": "Events deleted successfully"})
 }