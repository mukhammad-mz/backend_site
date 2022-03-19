package routes

import (
	"site_backend/authorization"
	"site_backend/files"
	"site_backend/middlewares"
	"site_backend/users"

	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// =========== Free to get token ==========
	authorization.GetRoutes(v1)

	hf := []gin.HandlerFunc{middlewares.SiteAuthentication(), middlewares.SiteAuthorization()}

	files.GetRoutes(v1, hf...)
	users.GetRoutes(v1, hf...)
}
