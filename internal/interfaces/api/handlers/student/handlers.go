package student

import (
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/repositoryimpl"

	"github.com/gin-gonic/gin"
)

func ConnectStudentCourseHandler(studentCourseRouter *gin.RouterGroup, dbClient *ent.Client) {
	courseRepo := repositoryimpl.NewCourseRepository(dbClient)
	studentService := service.NewStudentCourseService(courseRepo)
	handler := &StudentCourseHandler{
		studentService: studentService,
	}
	studentCourseRouter.GET("/my_courses", handler.GetMyCourses)
}

func ConnectStudentAuthHandler(studentAuthRouter *gin.RouterGroup, dbClient *ent.Client) {
	studentRepo := repositoryimpl.NewStudentRepository(dbClient)
	classRepository := repositoryimpl.NewClassRepository(dbClient)
	studentAuthService := service.NewStudentAuthService(classRepository, studentRepo)
	handler := StudentAuthHandler{
		studentAuthService: *studentAuthService,
	}
	studentAuthRouter.POST("/login", handler.Login)
	studentAuthRouter.POST("/registration", handler.Register)
	studentAuthRouter.POST("/logout", handler.Logout)
}
