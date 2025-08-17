package repository

import (
	"Runlet/internal/domain/entities"

	"context"
)

type CourseRepositoryInterface interface {
	GetCourseById(ctx context.Context, id int) (entities.Course, error)
	GetAllCourses(ctx context.Context) ([]entities.Course, error)
	GetAllStudentCourses(ctx context.Context, studentId int) ([]entities.Course, error)
	CreateCourse(ctx context.Context, title string, description string, classesIds []int, teachersIds []int) (entities.Course, error)
	DeleteCourse(ctx context.Context, id int) error
	AddClasses(ctx context.Context, courseId int, classesIds []int) ([]entities.Class, error)
	DeleteClasses(ctx context.Context, courseId int, classesIds []int) ([]entities.Class, error)
}
