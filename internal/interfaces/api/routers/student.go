package routers

import (
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/interfaces/api/handlers/student"
	"Runlet/internal/interfaces/api/middlewares"

	"github.com/gin-gonic/gin"
)

func ConnectStudentAuthRouter(studentRouter *gin.RouterGroup, dbClient *ent.Client) {
	authRouter := studentRouter.Group("/auth")
	student.ConnectStudentAuthHandler(authRouter, dbClient)
}

func ConnectStudentCourseRouter(studentRouter *gin.RouterGroup, dbClient *ent.Client) {
	courseRouter := studentRouter.Group("/course", middlewares.AuthMiddleware())
	student.ConnectStudentCourseHandler(courseRouter, dbClient)
}

func ConnectStudentRouter(apiRouter *gin.RouterGroup, dbClient *ent.Client) {
	studentRouter := apiRouter.Group("/student")
	ConnectStudentAuthRouter(studentRouter, dbClient)
	ConnectStudentCourseRouter(studentRouter, dbClient)
}
