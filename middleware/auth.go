package middleware

import (
	"net/http"

	"example.com/rest_api/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(context *gin.Context)  {
	token := context.Request.Header.Get("Authorization")
	if token == "" { // the AbortWithStatusJSON  when this json is send the next handler will not execute and it will return the response immediately , that why we not used JSON
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return 
	}

	id, err := utils.VerifyJWT(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return 
	}
	context.Set("UserId" , id)
	context.Next()
}