package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
server := gin.Default()

server.GET("/exact" , exactHandler )
server.Run(":8080")
}

func exactHandler(context *gin.Context) {
	context.JSON(http.StatusOK , gin.H{"message" : "Hello"}) //or it can be context.JSON(200 , gin.H{"message" : "Hello"})
}