package student

import (
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/repositoryimpl"

	"github.com/gin-gonic/gin"
)

func ConnectStudentCourseHandler(studentRouterGroup *gin.RouterGroup, dbClient *ent.Client) {
	courseRepo := repositoryimpl.NewCourseRepository(dbClient)
	studentService := service.NewStudentCourseService(courseRepo)
	handler := &StudentCourseHandler{
		studentService: studentService,
	}
	studentRouterGroup.GET("/my_courses", handler.GetMyCourses)
}

func ConnectStudentAuthHandler(studentRouterGroup *gin.RouterGroup, dbClient *ent.Client) {
	studentRepo := repositoryimpl.NewStudentRepository(dbClient)
	classRepository := repositoryimpl.NewClassRepository(dbClient)
	studentAuthService := service.NewStudentAuthService(classRepository, studentRepo)
	handler := StudentAuthHandler{
		studentAuthService: *studentAuthService,
	}
	studentRouterGroup.POST("/login", handler.Login)
	studentRouterGroup.POST("/registration", handler.Register)
}
