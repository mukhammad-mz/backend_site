package main

import (
	"fmt"
	"log"

	"site_backend/db"
	"site_backend/middlewares"
	"site_backend/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	logr "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/natefinch/lumberjack.v2"
)

const logFilePath = "logs/logError.log"

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
	lumberjackLogRotate := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    2,
		MaxAge:     60,
		MaxBackups: 500,
		LocalTime:  true,
		Compress:   true,
	}
	logr.SetOutput(lumberjackLogRotate)
	//print("Hello")
	if err := db.ConnectDB(); err == nil {
		//gin.SetMode(gin.ReleaseMode)
		router := gin.Default()
		//pprof.Register(router, "/test")
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.Use(middlewares.ReteLimitter)
		//router.Use(middlewares.Logger())

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
		log.Fatal(err)
	}
}
