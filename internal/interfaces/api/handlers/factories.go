package handlers

import (
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/interfaces/api/handlers/student"

	"github.com/gin-gonic/gin"
)

func ConnectStudentRouter(apiRouter *gin.RouterGroup, dbClient *ent.Client) {
	studentCourseHandler := student.NewStudentCourseHandler(dbClient)
	studentRouter := apiRouter.Group("/student")
	studentRouter.GET("/my_courses", studentCourseHandler.GetCourses)
	studentRouter.POST("/create_course", studentCourseHandler.CreateCourse)
}
