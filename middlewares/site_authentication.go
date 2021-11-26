package middlewares

import (
	"site_backend/authorization"

	"github.com/gin-gonic/gin"
)

func SiteAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		userID := c.GetHeader("ID")
		t := authorization.ChekToken(token)
		if !t || token == "" {
			redirectToAccessDenied(c)
			return
		}
		c.Set("userID", userID)
	}
}
