package student

import (
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"context"
	"errors"
)

type StudentCourseService struct {
	CourseRepository repository.CourseRepositoryInterface
}

func (s StudentCourseService) GetStudentCourses(ctx context.Context, studentId int) ([]*ent.Course, error) {
	courses, err := s.CourseRepository.GetAllStudentCourses(ctx, studentId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.New("error find courses")
	}
	return courses, nil
}
