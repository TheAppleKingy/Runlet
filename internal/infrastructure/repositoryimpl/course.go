package repositoryimpl

import (
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/repository"
	textdata "Runlet/internal/infrastructure/text_data"
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
	var course entities.Course
	found, err := r.db.From(textdata.CourseTable).Select().Where(goqu.Ex{"id": id}).ScanStructContext(ctx, &course)
	if err != nil || !found {
		return entities.Course{}, err
	}
	return course, nil
}

func (r *CourseRepository) GetAllCourses(ctx context.Context) ([]entities.Course, error) {
	var courses []entities.Course
	if err := r.db.From(textdata.CourseTable).Select().ScanStructsContext(ctx, &courses); err != nil {
		return []entities.Course{}, err
	}
	return courses, nil
}

func (r *CourseRepository) GetAllStudentCourses(ctx context.Context, studentId int) ([]entities.Course, error) {
	var courses []entities.Course
	if err := r.db.From(goqu.T(textdata.CourseTable).As("c")).Select(goqu.I("c.*")).
		Join(
			goqu.T("classes_courses").As("cc"),
			goqu.On(goqu.I("c.id").Eq(goqu.I("cc.course_id"))),
		).Join(
		goqu.T(textdata.ClassTable).As("cls"),
		goqu.On(goqu.I("cls.id").Eq(goqu.I("cc.class_id"))),
	).
		Join(
			goqu.T(textdata.StudentTable).As("s"),
			goqu.On(goqu.I("s.class_id").Eq(goqu.I("cls.id"))),
		).
		Where(goqu.I("s.id").Eq(studentId)).ScanStructsContext(ctx, &courses); err != nil || len(courses) == 0 {
		return []entities.Course{}, err
	}
	return courses, nil
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
