package users

import "github.com/gin-gonic/gin"

func GetRoutes(r *gin.RouterGroup, hf ...gin.HandlerFunc) {
	user := r.Group("/user")
	user.Use(hf...)
	user.GET("", GetUser)
	user.GET("/users", GetUsers)
	user.POST("", PostUser)
	user.PUT("", PutUser)
	user.PUT("/password", UserChangePassword)
	user.DELETE(":uid", DeleteUser)
	user.GET("/login",chenckLogin)
}
