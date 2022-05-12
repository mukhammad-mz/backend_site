package routes

import (
	"net/http"
	"site_backend/authorization"
	"site_backend/files"
	"site_backend/middlewares"
	"site_backend/response"
	"site_backend/users"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.NotFound())
	})
	
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// =========== Free to get token ==========
	authorization.GetRoutes(v1)

	hf := []gin.HandlerFunc{middlewares.SiteAuthentication(), middlewares.SiteAuthorization()}

	files.GetRoutes(v1, hf...)
	users.GetRoutes(v1, hf...)
}
