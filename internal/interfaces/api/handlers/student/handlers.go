package student

import (
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/repositoryimpl"
)

func NewStudentCourseHandler(dbClient *ent.Client) *StudentCourseHandler {
	courseRepo := repositoryimpl.NewCourseRepository(dbClient)
	courseService := service.NewCourseService(courseRepo)
	return &StudentCourseHandler{
		CourseService: courseService,
	}
}
