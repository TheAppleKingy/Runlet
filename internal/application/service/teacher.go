package service

import "Runlet/internal/domain/interfaces"

type TeacherService struct {
	TeacherRepository interfaces.TeacherRepository
	StudentRepository interfaces.StudentRepository
	ClassRepository   interfaces.ClassRepository
	CourseRepository  interfaces.CourseRepository
}
