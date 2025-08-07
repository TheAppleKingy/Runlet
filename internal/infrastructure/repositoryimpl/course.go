package repositoryimpl

import (
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/ent/course"
	"context"
)

type CourseRepository struct {
	repository.CourseRepositoryInterface
	client *ent.Client
}

func (cr *CourseRepository) GetCourseById(ctx context.Context, id int) (*ent.Course, error) {
	return cr.client.Course.Get(ctx, id)
}

func (cr *CourseRepository) GetCourseByTitle(ctx context.Context, title string) (*ent.Course, error) {
	return cr.client.Course.Query().Where(course.TitleEQ(title)).Only(ctx)
}

func (cr *CourseRepository) GetCourses(ctx context.Context) ([]*ent.Course, error) {
	return cr.client.Course.Query().WithProblems().All(ctx)
}

func (cr *CourseRepository) CreateCourse(ctx context.Context, title string, description string) (*ent.Course, error) {
	return cr.client.Course.Create().SetTitle(title).SetDescription(description).Save(ctx)
}
