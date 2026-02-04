package routes

import (
	"net/http"

	"example.com/rest_api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
		if err != nil {
			context.JSON(http.StatusBadRequest , gin.H{"message": "Could not parse the request"})
			return
		}
    
    err = user.Save()	
		if err != nil {
			context.JSON(http.StatusInternalServerError , gin.H{"message" : "Could not create the user"})
			return
		}	

	context.JSON(http.StatusCreated , gin.H{"message" : "User created successfully"})	
}