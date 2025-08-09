package repositoryimpl

import (
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/ent/class"
	"context"
)

type CourseRepository struct {
	repository.CourseRepositoryInterface
	client *ent.Client
}

func (cr *CourseRepository) GetCourseById(ctx context.Context, id int) (*ent.Course, error) {
	return cr.client.Course.Get(ctx, id)
}

func (cr *CourseRepository) GetAllCourses(ctx context.Context) ([]*ent.Course, error) {
	return cr.client.Course.Query().WithProblems().All(ctx)
}

func (cr *CourseRepository) CreateCourse(ctx context.Context, title string, description string, classesIds []int) (*ent.Course, error) {
	return cr.client.Course.Create().SetTitle(title).SetDescription(description).AddClassIDs(classesIds...).Save(ctx)
}

func (cr *CourseRepository) DeleteCourse(ctx context.Context, id int) error {
	return cr.client.Course.DeleteOneID(id).Exec(ctx)
}

func (cr *CourseRepository) AddClasses(ctx context.Context, courseId int, classesIds []int) ([]*ent.Class, error) {
	updatedCourse, err := cr.client.Course.UpdateOneID(courseId).AddClassIDs(classesIds...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updatedCourse.QueryClasses().Where(class.IDIn(classesIds...)).All(ctx)
}

func (cr *CourseRepository) ExcludeStudents(ctx context.Context, courseId int, classesIds []int) ([]*ent.Class, error) {
	updatedCourse, err := cr.client.Course.UpdateOneID(courseId).RemoveClassIDs(classesIds...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updatedCourse.QueryClasses().Where(class.IDIn(classesIds...)).All(ctx)
}
