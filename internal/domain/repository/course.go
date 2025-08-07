package repository

import (
	"Runlet/internal/infrastructure/ent"

	"context"
)

type CourseRepositoryInterface interface {
	GetCourseById(ctx context.Context, id int) (*ent.Course, error)
	GetCourseByTitle(ctx context.Context, title string) (*ent.Course, error)
	GetCourses(ctx context.Context) ([]*ent.Course, error)
	CreateCourse(ctx context.Context, title string, description string) (*ent.Course, error)
}
