package middlewares

import (
	"site_backend/authorization"
	"site_backend/users"
	"strings"

	"github.com/gin-gonic/gin"
)

func SiteAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer authorization.Recoverd(c, "SiteAuthorization: ")
		handlerName := c.HandlerName()
		loc := handlerName[strings.Index(handlerName, "/")+1:]
		userUID, ext := c.Get("userUID")
		if !ext {
			redirectToAccessDenied(c)
			return
		} else if !users.CheckPermission(userUID.(string), loc) {
			redirectToAccessDenied(c)
			return
		}
	}

}
