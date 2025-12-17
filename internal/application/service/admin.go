package service

import "Runlet/internal/domain/interfaces"

type AdminService struct {
	CourseRepository  interfaces.CourseRepository
	ClassRepository   interfaces.ClassRepository
	StudentRepository interfaces.StudentRepository
	TeacherRepository interfaces.TeacherRepository
	ProblemRepository interfaces.ProblemRepository
	AttemptRepository interfaces.AttemptRepository
}
