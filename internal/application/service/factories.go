package service

import "Runlet/internal/domain/repository"

func NewCourseService(courseRepo repository.CourseRepositoryInterface) *CourseService {
	return &CourseService{
		courseRepo: courseRepo,
	}
}

func NewStudentAuthService(classRepo repository.ClassRepositoryInterface, studentRepo repository.StudentRepositoryInterface) *StudentAuthService {
	return &StudentAuthService{
		studentRepository: studentRepo,
		classRepository:   classRepo,
	}
}

func NewStudentCourseService(courseRepo repository.CourseRepositoryInterface) *StudentCourseService {
	return &StudentCourseService{
		courseRepository: courseRepo,
	}
}
