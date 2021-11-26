package main

import (
	"fmt"

	"site_backend/db"
	_ "site_backend/docs"
	"site_backend/middlewares"
	"site_backend/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/natefinch/lumberjack.v2"
)

const logFilePath = "logs/logError.log"
const DOWNLOADS_PATH = "D:/projekts/tests/"

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	if err := db.ConnectDB(); err == nil {
		lumberjackLogRotate := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    5,   // Max megabytes before log is rotated
			MaxBackups: 500, // Max number of old log files to keep
			MaxAge:     60,  // Max number of days to retain log files
			Compress:   true,
		}
		log.SetOutput(lumberjackLogRotate)
		//gin.SetMode(gin.ReleaseMode)
		router := gin.Default()
		//pprof.Register(router, "/test")
		router.Use(middlewares.ReteLimitter)
		router.Use(middlewares.Logger())
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.Use(static.Serve("/file", static.LocalFile(DOWNLOADS_PATH, false)))

		if gin.Mode() == gin.DebugMode {
			router.Use(cors.New(cors.Config{
				AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS", "GET"},
				AllowCredentials: true,
				AllowAllOrigins:  true,
				AllowHeaders:     []string{"*"},
				AllowWildcard:    true,
				MaxAge:           12 * time.Hour,
			}))
		}
		routes.GetRoutes(router)
		if err := router.Run(":8080"); err != nil {
			fmt.Println(err)
			panic("err")
		}
	} else {
		log.Warnf("%v", err)
	}
}
