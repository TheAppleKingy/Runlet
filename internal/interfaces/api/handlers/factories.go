package handlers

import (
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/interfaces/api/handlers/student"

	"github.com/gin-gonic/gin"
)

func ConnectStudentRouter(apiRouter *gin.RouterGroup, dbClient *ent.Client) {
	studentRouter := apiRouter.Group("/student")
	student.ConnectStudentAuthHandler(studentRouter, dbClient)
	student.ConnectStudentCourseHandler(studentRouter, dbClient)
}
