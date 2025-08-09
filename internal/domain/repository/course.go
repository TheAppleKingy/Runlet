package repository

import (
	"Runlet/internal/infrastructure/ent"

	"context"
)

type CourseRepositoryInterface interface {
	GetCourseById(ctx context.Context, id int) (*ent.Course, error)
	GetAllCourses(ctx context.Context) ([]*ent.Course, error)
	CreateCourse(ctx context.Context, title string, description string, classesIds []int) (*ent.Course, error)
	DeleteCourse(ctx context.Context, id int) error
	AddClasses(ctx context.Context, courseId int, classesIds []int) ([]*ent.Class, error)
	DeleteClasses(ctx context.Context, courseId int, classesIds []int) ([]*ent.Student, error)
}
