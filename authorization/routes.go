package authorization

import (
	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.RouterGroup, hf ...gin.HandlerFunc) {
	auth := r.Group("/auth")
	//auth.Use(hf...)
	auth.POST("/token", GetToken)
}
