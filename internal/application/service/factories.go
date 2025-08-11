package service

import (
	"Runlet/internal/application/service/course"
	"Runlet/internal/application/service/student"
	"Runlet/internal/domain/repository"
)

func NewCourseService(courseRepo repository.CourseRepositoryInterface) *course.CourseService {
	return &course.CourseService{
		CourseRepo: courseRepo,
	}
}

func NewStudentAuthService(classRepo repository.ClassRepositoryInterface, studentRepo repository.StudentRepositoryInterface) *student.StudentAuthService {
	return &student.StudentAuthService{
		StudentRepository: studentRepo,
		ClassRepository:   classRepo,
	}
}

func NewStudentCourseService(courseRepo repository.CourseRepositoryInterface) *student.StudentCourseService {
	return &student.StudentCourseService{
		CourseRepository: courseRepo,
	}
}
