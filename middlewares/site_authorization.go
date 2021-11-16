package middlewares

import (
	"site_backend/users"
	"strings"

	"github.com/gin-gonic/gin"
)

func SiteAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		handlerName := c.HandlerName()
		loc := handlerName[strings.Index(handlerName, "/")+1:]
		userID := c.GetHeader("ID")
		
		if !users.CheckPermission(userID, loc) {
			redirectToAccessDenied(c)
			return
		}

	}

}
