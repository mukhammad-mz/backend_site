package middlewares

import (
	"site_backend/authorization"

	"github.com/gin-gonic/gin"
)

func SiteAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer authorization.Recoverd(c, "SiteAuthentication: ")

		token := c.GetHeader("Authorization")
		t, err := authorization.ChekToken(token)
		if err != nil {
			redirectToAccessDenied(c)
			return
		}
		c.Set("userUID", t)
	}
}
