package handlers

import (
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/repositoryimpl"
)

func NewCourseHandler(dbClient *ent.Client) *CourseHandler {
	courseRepo := repositoryimpl.NewCourseRepository(dbClient)
	courseService := service.NewCourseService(courseRepo)
	return &CourseHandler{
		courseService: courseService,
	}
}
