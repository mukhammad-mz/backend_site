package files

import (
	"github.com/gin-gonic/gin"
)


func GetRoutes(r *gin.RouterGroup, hf ...gin.HandlerFunc) {
	music := r.Group("/music")
	music.GET("/list", GetListMusic)
	music.GET("", GetMisicInfo)
	music.GET("/download/:filename", download)

	file := r.Group("/file")
	file.Use(hf...)
	file.GET("/files", GetFile)
	file.POST("", PostFile)
	file.GET("", GetUserFile)
	file.PUT("", PutFile)
	file.DELETE(":id", DeleteFile)


}
