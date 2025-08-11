package api

import (
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/interfaces/api/routers"

	_ "Runlet/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Runlet API
// @version 1.0
// @description API documentation for Runlet
// @host localhost:8080
// @BasePath /api
func GetRouter(dbClient *ent.Client) *gin.Engine {
	router := gin.Default()
	apiRouter := router.Group("/api")
	routers.ConnectStudentRouter(apiRouter, dbClient)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
