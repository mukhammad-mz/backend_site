package middlewares

import (
	"net/http"
	"site_backend/response"

	"github.com/gin-gonic/gin"
)

func redirectToAccessDenied(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, response.Forbidden())
}
