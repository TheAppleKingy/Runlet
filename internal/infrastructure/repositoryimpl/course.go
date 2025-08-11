package repositoryimpl

import (
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/ent/class"
	"Runlet/internal/infrastructure/ent/course"
	"context"
)

type CourseRepository struct {
	repository.CourseRepositoryInterface
	client *ent.Client
}

func (r *CourseRepository) GetCourseById(ctx context.Context, id int) (*ent.Course, error) {
	return r.client.Course.Get(ctx, id)
}

func (r *CourseRepository) GetAllCourses(ctx context.Context) ([]*ent.Course, error) {
	return r.client.Course.Query().WithProblems().All(ctx)
}

func (r *CourseRepository) GetAllStudentCourses(ctx context.Context, studentId int) ([]*ent.Course, error) {
	st := r.client.Student.GetX(ctx, studentId)
	return r.client.Course.Query().WithProblems().Where(course.HasClassesWith(class.IDEQ(st.ClassID))).All(ctx)
}

func (r *CourseRepository) CreateCourse(ctx context.Context, title string, description string, classesIds []int, teachersIds []int) (*ent.Course, error) {
	return r.client.Course.Create().SetTitle(title).SetDescription(description).AddClassIDs(classesIds...).AddTeacherIDs(teachersIds...).Save(ctx)
}

func (r *CourseRepository) DeleteCourse(ctx context.Context, id int) error {
	return r.client.Course.DeleteOneID(id).Exec(ctx)
}

func (r *CourseRepository) AddClasses(ctx context.Context, courseId int, classesIds []int) ([]*ent.Class, error) {
	updatedCourse, err := r.client.Course.UpdateOneID(courseId).AddClassIDs(classesIds...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updatedCourse.QueryClasses().Where(class.IDIn(classesIds...)).All(ctx)
}

func (r *CourseRepository) ExcludeStudents(ctx context.Context, courseId int, classesIds []int) ([]*ent.Class, error) {
	updatedCourse, err := r.client.Course.UpdateOneID(courseId).RemoveClassIDs(classesIds...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updatedCourse.QueryClasses().Where(class.IDIn(classesIds...)).All(ctx)
}
