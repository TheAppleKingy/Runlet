package api

import (
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/interfaces/api/handlers"

	"github.com/gin-gonic/gin"
)

func GetRouter(dbClient *ent.Client) *gin.Engine {
	courseHandler := handlers.NewCourseHandler(dbClient)
	router := gin.Default()
	router.GET("/api/courses", courseHandler.GetCourses)
	router.POST("/api/courses", courseHandler.CreateCourse)
	return router
}
