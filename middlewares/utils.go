package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func redirectToAccessDenied(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "forbidden"})
}
