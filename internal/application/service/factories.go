package service

import "Runlet/internal/domain/repository"

func NewCourseService(courseRepo repository.CourseRepositoryInterface) *CourseService {
	return &CourseService{
		courseRepo: courseRepo,
	}
}
