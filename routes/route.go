package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(route *gin.Engine) {

	authRoutes(route)
	eventRoutes(route)

}