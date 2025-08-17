package repositoryimpl

import (
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/repository"
	"context"

	"github.com/doug-martin/goqu/v9"
)

type CourseRepository struct {
	repository.CourseRepositoryInterface
	db *goqu.Database
}

func NewCourseRepository(db *goqu.Database) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (r *CourseRepository) GetCourseById(ctx context.Context, id int) (entities.Course, error) {
	return entities.Course{}, nil
}

func (r *CourseRepository) GetAllCourses(ctx context.Context) ([]entities.Course, error) {
	return make([]entities.Course, 0), nil
}

func (r *CourseRepository) GetAllStudentCourses(ctx context.Context, studentId int) ([]entities.Course, error) {
	return make([]entities.Course, 0), nil
}

func (r *CourseRepository) CreateCourse(ctx context.Context, title string, description string, classesIds []int, teachersIds []int) (entities.Course, error) {
	return entities.Course{}, nil
}

func (r *CourseRepository) DeleteCourse(ctx context.Context, id int) error {
	return nil
}

func (r *CourseRepository) AddClasses(ctx context.Context, courseId int, classesIds []int) ([]entities.Class, error) {
	return make([]entities.Class, 0), nil
}

func (r *CourseRepository) ExcludeStudents(ctx context.Context, courseId int, classesIds []int) ([]entities.Class, error) {
	return make([]entities.Class, 0), nil
}
