package routes

import (
	"net/http"

	"example.com/rest_api/models"
	"example.com/rest_api/utils"
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

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	 	if err != nil {
			context.JSON(http.StatusBadRequest , gin.H{"message" : "Could not parse the request"})
			return
		}

    err = user.ValidateCredentials()	
	if err != nil {
		context.JSON(http.StatusUnauthorized , gin.H{"message" : "Could not autenticate the user"})
		return
	}
	token , err := utils.GenerateJWT(user.Id , user.Email)
	   if err != nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message" : "Could not generate the token"})
	   }

	   
	context.JSON(http.StatusOK , gin.H{"message" : "User logged in successfully" , "token" : token})
}